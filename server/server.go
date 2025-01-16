package server

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/interaction_type"
	"github.com/JackHumphries9/dapper-go/interactable"
	"github.com/JackHumphries9/dapper-go/managers"
	"github.com/JackHumphries9/dapper-go/verification"
)

type InteractionServerOptions struct {
	PublicKey    ed25519.PublicKey
	DefaultRoute string
	DapperLogger *DapperLogger
}

var defaultConfig = InteractionServerOptions{
	PublicKey:    ed25519.PublicKey(""),
	DefaultRoute: "/interactions",
}

type InteractionServer struct {
	opts             InteractionServerOptions
	commandManager   managers.CommandManager
	componentManager managers.ComponentManager
	modalManager     managers.ModalManager
	logger           *DapperLogger
}

func (is *InteractionServer) handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		is.logger.Error("Only POST method is supported")
		return
	}

	verify := verification.Verify(r, is.opts.PublicKey)

	if !verify {
		is.logger.Error("Recieved an invalid request")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		is.logger.Error("Failed to read body")
		return
	}

	interaction, err := discord.ParseInteraction(string(rawBody))

	if err != nil {
		is.logger.Error(fmt.Sprintf("Failed to parse interaction: %v\n", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	is.logger.OnInteractionRecieved(interaction)

	if interaction.IsPing() {
		discord.CreatePongResponse().ToHttpResponse().WriteResponse(w)
		return
	}

	var interactionResponse discord.InteractionResponse

	if interaction.Type == interaction_type.ApplicationCommand {
		interactionResponse, err = is.commandManager.RouteInteraction(interaction)
	} else if interaction.Type == interaction_type.MessageComponent {
		interactionResponse, err = is.componentManager.RouteInteraction(interaction)
	} else if interaction.Type == interaction_type.ModalSubmit {
		interactionResponse, err = is.modalManager.RouteInteraction(interaction)
	} else {
		is.logger.Error(fmt.Sprintf("Unknown interaction type: %d\n", interaction.Type))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		is.logger.Error(fmt.Sprintf("An error occured while handling the interaction: %+v", err))

		w.WriteHeader(500)
		return
	}

	body, err := json.Marshal(interactionResponse)
	if err != nil {
		is.logger.Error("An error occured while responding to interaction")
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

func (is *InteractionServer) RegisterCommand(cmd interactable.Command) {
	is.commandManager.Register(cmd)

	for _, comp := range cmd.GetComponents() {
		is.componentManager.Register(comp)
	}

	for _, modal := range cmd.GetModals() {
		is.modalManager.Register(modal)
	}
}

func (is *InteractionServer) RegisterComponent(comp interactable.Component) {
	is.componentManager.Register(comp)
}

func (is *InteractionServer) RegisterModal(modal interactable.Modal) {
	is.modalManager.Register(modal)
}

func (is *InteractionServer) RegisterCommandsWithDiscord(appId discord.Snowflake, client *client.BotClient) error {
	err := is.commandManager.RegisterCommandsWithDiscord(appId, client)

	if err != nil {
		is.logger.Error(fmt.Sprintf("Failed to register discord commands: %v\n", err))
	} else {
		is.logger.Info("Successfully registered discord commands")
	}

	return err
}

func (is *InteractionServer) Listen(port int) error {
	is.registerRoute()

	is.logger.Info(fmt.Sprintf("Serving Discord Interactions on http://localhost:%d\n", port))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if errors.Is(err, http.ErrServerClosed) {
		is.logger.Error("Server closed\n")
	} else if err != nil {
		is.logger.Error(fmt.Sprintf("Error starting server: %s\n", err))
	}

	if err != nil {
		is.logger.Error(fmt.Sprintf("Error starting server: %s\n", err))
	}

	return err
}

func NewInteractionServer(publicKey string) InteractionServer {
	key, err := hex.DecodeString(publicKey)

	if err != nil {
		panic("Invalid public key")
	}

	return NewInteractionServerWithOptions(InteractionServerOptions{
		PublicKey:    ed25519.PublicKey(key),
		DefaultRoute: defaultConfig.DefaultRoute,
		DapperLogger: &DefaultLogger,
	})
}

func NewInteractionServerWithOptions(iso InteractionServerOptions) InteractionServer {
	if iso.DapperLogger == nil {
		iso.DapperLogger = &DefaultLogger
	}

	return InteractionServer{
		opts:             iso,
		commandManager:   managers.NewDapperCommandManager(),
		componentManager: managers.NewDapperComponentManager(),
		modalManager:     managers.NewDapperModalManager(),
		logger:           iso.DapperLogger,
	}
}
