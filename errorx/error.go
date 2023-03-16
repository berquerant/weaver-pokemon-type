package errorx

import (
	"fmt"
)

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type StatusCode -output error_stringer_generated.go

type StatusCode int

const (
	Unknown StatusCode = iota
	OK
	BadRequest
)

type Error struct {
	err        error
	statusCode StatusCode
	message    string
}

//go:generate go run github.com/berquerant/goconfig@v0.2.0 -configOption Option -option -output error_config_generated.go -field "StatusCode StatusCode|Message string"

func New(err error, opt ...Option) *Error {
	config := NewConfigBuilder().
		StatusCode(Unknown).
		Message("").
		Build()
	config.Apply(opt...)

	return &Error{
		err:        err,
		statusCode: config.StatusCode.Get(),
		message:    config.Message.Get(),
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%s %v", e.statusCode, e.message, e.err)
}

func (e *Error) Unwrap() []error {
	return []error{e.err}
}

func (e *Error) StatusCode() StatusCode {
	return e.statusCode
}
