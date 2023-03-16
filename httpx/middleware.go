package httpx

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/berquerant/weaver-pokemon-type/errorx"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func (h HandlerFunc) DefaultHTTPHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			instance  = WeaverInstanceFromContext(r.Context())
			requestID = requestIDFromContext(r.Context())
		)

		instance.Logger().Info("start",
			"uuid", requestID,
			"url", r.URL,
			"method", r.Method,
		)

		err := h(w, r)
		if err == nil {
			return
		}

		instance.Logger().Error("result", err, "uuid", requestID)

		status := func() int {
			var xErr *errorx.Error

			if errors.As(err, &xErr) {
				switch xErr.StatusCode() {
				case errorx.OK:
					return http.StatusOK
				case errorx.BadRequest:
					return http.StatusBadRequest
				default:
					return http.StatusInternalServerError
				}
			}
			return http.StatusOK
		}()
		w.WriteHeader(status)
	}
}

func JSONHandler[RequestT any, ResponseT any](
	handler func(context.Context, RequestT) (ResponseT, error),
) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return errorx.New(err)
		}
		var request RequestT
		if err := json.Unmarshal(body, &request); err != nil {
			return errorx.New(err, errorx.WithStatusCode(errorx.BadRequest))
		}

		result, err := handler(r.Context(), request)
		if err != nil {
			return err
		}

		responseBody, err := json.Marshal(result)
		if err != nil {
			return errorx.New(err)
		}

		if _, err := w.Write(responseBody); err != nil {
			return errorx.New(err)
		}
		return nil
	}
}
