package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/interaction_callback_type"
	"github.com/JackHumphries9/dapper-go/discord/interaction_type"
	"github.com/JackHumphries9/dapper-go/verification"
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
	verify, err := verification.VerifyHttpRequest(is.opts.PublicKey, r)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	if !verify {
		fmt.Printf("Failed verification")
		w.WriteHeader(401)
		return
	}

	rawBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("Failed to read body")
		return
	}

	interaction, err := discord.ParseInteraction(string(rawBody))

	if err != nil {
		fmt.Printf("Failed to parse interaction")
		return
	}

	if interaction.IsPing() {
		discord.CreatePongResponse().ToHttpResponse().WriteResponse(w)
		return
	}

	if interaction.Type == interaction_type.ApplicationCommand {
		commandData := interaction.Data.(*discord.ApplicationCommandData)

		if commandData.Name == "hello" {
			response := &discord.InteractionResponse{
				Type: interaction_callback_type.UpdateMessage,
				Data: discord.MessageCallbackData{
					Embeds: []discord.Embed{
						{
							Title: "Hello World!",
						},
					},
				},
			}

			response.ToHttpResponse().WriteResponse(w)
			return
		}

	}

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
