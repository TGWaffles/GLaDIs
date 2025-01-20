package server

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
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
	DapperLogger *DapperLogger
}

var defaultConfig = InteractionServerOptions{
	PublicKey: ed25519.PublicKey(""),
}

type InteractionHandler struct {
	opts             InteractionServerOptions
	commandManager   managers.CommandManager
	componentManager managers.ComponentManager
	modalManager     managers.ModalManager
	logger           *DapperLogger
}

func (is *InteractionHandler) Handle(w http.ResponseWriter, r *http.Request) {
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

func (is *InteractionHandler) RegisterCommand(cmd interactable.Command) {
	is.commandManager.Register(cmd)

	for _, comp := range cmd.GetComponents() {
		is.componentManager.Register(comp)
	}

	for _, modal := range cmd.GetModals() {
		is.modalManager.Register(modal)
	}
}

func (is *InteractionHandler) RegisterComponent(comp interactable.Component) {
	is.componentManager.Register(comp)
}

func (is *InteractionHandler) RegisterModal(modal interactable.Modal) {
	is.modalManager.Register(modal)
}

func (is *InteractionHandler) RegisterCommandsWithDiscord(appId discord.Snowflake, client *client.BotClient) error {
	err := is.commandManager.RegisterCommandsWithDiscord(appId, client)

	if err != nil {
		is.logger.Error(fmt.Sprintf("Failed to register discord commands: %v\n", err))
	} else {
		is.logger.Info("Successfully registered discord commands")
	}

	return err
}

func NewInteractionServer(publicKey string) InteractionHandler {
	key, err := hex.DecodeString(publicKey)

	if err != nil {
		panic("Invalid public key")
	}

	return NewInteractionHandlerWithOptions(InteractionServerOptions{
		PublicKey:    ed25519.PublicKey(key),
		DapperLogger: &DefaultLogger,
	})
}

func NewInteractionHandlerWithOptions(iso InteractionServerOptions) InteractionHandler {
	if iso.DapperLogger == nil {
		iso.DapperLogger = &DefaultLogger
	}

	return InteractionHandler{
		opts:             iso,
		commandManager:   managers.NewDapperCommandManager(),
		componentManager: managers.NewDapperComponentManager(),
		modalManager:     managers.NewDapperModalManager(),
		logger:           iso.DapperLogger,
	}
}
