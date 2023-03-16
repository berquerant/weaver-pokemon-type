package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/ServiceWeaver/weaver"
	"github.com/berquerant/weaver-pokemon-type/httpx"
	"github.com/berquerant/weaver-pokemon-type/service"
	"github.com/caarlos0/env/v7"
)

type config struct {
	// App is an application name.
	App string `env:"APP" envDefault:"pokemon-type"`
	// Host is a host name that the server listens on.
	Host string `env:"HOST" envDefault:"localhost"`
	// Port is a port number that the server listens on.
	Port int `env:"PORT" envDefault:"21099"`
}

func main() {
	log.Panic(run(context.Background()))
}

func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	root := weaver.Init(ctx)

	var (
		cfg     config
		envOpts = env.Options{
			OnSet: func(tag string, value any, isDefault bool) {
				root.Logger().Info("read environment variable",
					"name", tag,
					"value", value,
					"default", isDefault,
				)
			},
		}
	)
	if err := env.Parse(&cfg, envOpts); err != nil {
		root.Logger().Error("failed to read environment variable", err)
		return err
	}

	weaverOpts := weaver.ListenerOptions{
		LocalAddress: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}
	lis, err := root.Listener(cfg.App, weaverOpts)
	if err != nil {
		root.Logger().Error("failed to listen", err, "address", weaverOpts.LocalAddress)
	}

	root.Logger().Info("listener available", "app", cfg.App, "address", lis)

	api, err := weaver.Get[service.API](root)
	if err != nil {
		root.Logger().Error("failed to construct API", err)
		return err
	}

	/* HTTP handlers */
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	handlers := map[string]http.HandlerFunc{
		"/type":    wrap(root, api.GetTypeByName),
		"/attack":  wrap(root, api.GetEffectivityListByAttack),
		"/defense": wrap(root, api.GetEffectivityListByDefenseList),
	}
	for pattern, handler := range handlers {
		http.HandleFunc(pattern, weaver.InstrumentHandlerFunc(pattern, handler).ServeHTTP)
	}
	return http.Serve(lis, nil)
}

func wrap[X any, Y any](
	instance weaver.Instance,
	f func(context.Context, X) (Y, error),
) http.HandlerFunc {
	return httpx.InjectWeaverInstance(
		instance,
		httpx.JSONHandler(f).DefaultHTTPHandlerFunc(),
	)
}
