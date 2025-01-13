package server

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/interaction_type"
	"github.com/JackHumphries9/dapper-go/managers"
	"github.com/JackHumphries9/dapper-go/verification"
)

type InteractionServerOptions struct {
	PublicKey    ed25519.PublicKey
	DefaultRoute string
}

var defaultConfig = InteractionServerOptions{
	PublicKey:    ed25519.PublicKey(""),
	DefaultRoute: "/interactions",
}

type InteractionServer struct {
	opts           InteractionServerOptions
	commandManager managers.DapperCommandManager
}

func (is *InteractionServer) handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Printf("Interaction Recieved\n")
	verify := verification.Verify(r, is.opts.PublicKey)

	if !verify {
		fmt.Println("Failed verification")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		fmt.Println("Failed to read body")
		return
	}

	fmt.Printf("Interaction %s\n", string(rawBody))

	interaction, err := discord.ParseInteraction(string(rawBody))

	if err != nil {
		fmt.Printf("Failed to parse interaction: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if interaction.IsPing() {
		discord.CreatePongResponse().ToHttpResponse().WriteResponse(w)
		return
	}

	var interactionResponse discord.InteractionResponse

	if interaction.Type == interaction_type.ApplicationCommand {
		interactionResponse, err = is.commandManager.RouteInteraction(interaction)
	}
	if interaction.Type == interaction_type.MessageComponent {

	}

	if err != nil {
		fmt.Printf("An error occured while handling the interaction")

		// TODO: Add better error handling
		w.WriteHeader(500)
		return
	}

	body, err := json.Marshal(interactionResponse)
	if err != nil {
		fmt.Printf("An error occured while responding")
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(body)

	return
}

func (is *InteractionServer) registerRoute() {
	http.HandleFunc(is.opts.DefaultRoute, is.handle)
}

func (is *InteractionServer) RegisterCommand(cmd managers.DapperCommand) {
	is.commandManager.Register(cmd)
}

func (is *InteractionServer) Listen(port int) {
	is.registerRoute()

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}

func NewInteractionServer(publicKey string) InteractionServer {
	key, err := hex.DecodeString(publicKey)

	if err != nil {
		panic("Invalid public key")
	}

	return NewInteractionServerWithOptions(InteractionServerOptions{
		PublicKey:    ed25519.PublicKey(key),
		DefaultRoute: defaultConfig.DefaultRoute,
	})
}

func NewInteractionServerWithOptions(iso InteractionServerOptions) InteractionServer {
	return InteractionServer{
		opts:           iso,
		commandManager: managers.NewDapperCommandManager(),
	}
}
