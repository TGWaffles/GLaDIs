package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type InteractionServerOptions struct {
	PublicKey    string
	DefaultRoute string
}

var defaultConfig = InteractionServerOptions{
	PublicKey:    "",
	DefaultRoute: "/interaction",
}

type InteractionServer struct {
	opts InteractionServerOptions
}

func (is *InteractionServer) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func (is *InteractionServer) registerRoute() {
	http.HandleFunc(is.opts.DefaultRoute, is.handle)

}

func (is *InteractionServer) Listen(port int) {
	is.registerRoute()

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}

func NewInteractionServer(publicKey string) InteractionServer {
	return NewInteractionServerWithOptions(InteractionServerOptions{
		PublicKey:    publicKey,
		DefaultRoute: defaultConfig.DefaultRoute,
	})
}

func NewInteractionServerWithOptions(iso InteractionServerOptions) InteractionServer {
	return InteractionServer{
		opts: iso,
	}
}
