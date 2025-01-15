package main

import (
	button_command "dapper-go/get_started/commands/button"
	ping_command "dapper-go/get_started/commands/ping"
	"os"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/server"
)

const FILENAME = "../env.json"

type Env struct {
	PublicKey string `json:"PUBLIC_KEY"`
	BotToken  string `json:"BOT_TOKEN"`
	AppId     string `json:"APP_ID"`
}

func main() {
	botServer := server.NewInteractionServer(os.Getenv("PUBLIC_KEY"))
	botClient := client.NewBot(os.Getenv("BOT_TOKEN"))

	appId, err := discord.GetSnowflake(os.Getenv("APP_ID"))

	if err != nil {
		panic("Heyo you messed up")
	}

	botServer.RegisterCommand(button_command.Command)
	botServer.RegisterCommand(ping_command.Command)

	botServer.RegisterCommandsWithDiscord(appId, botClient)

	botServer.Listen(3000)
}
