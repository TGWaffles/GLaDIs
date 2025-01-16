package main

import (
	button_command "dapper-go/get_started/commands/button"
	ping_command "dapper-go/get_started/commands/ping"
	"encoding/json"
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

func LoadJSONEnv() Env {
	plan, err := os.ReadFile(FILENAME)

	if err != nil {
		panic("no env file")
	}

	var data Env
	err = json.Unmarshal(plan, &data)

	if err != nil {
		panic("cannot unmarshal")
	}

	return data
}

func main() {
	env := LoadJSONEnv()
	botServer := server.NewInteractionServer(env.PublicKey)
	botClient := client.NewBot(env.BotToken)

	appId, err := discord.GetSnowflake(env.AppId)

	if err != nil {
		panic("Heyo you messed up")
	}

	botServer.RegisterCommand(button_command.Command)
	botServer.RegisterCommand(ping_command.Command)

	botServer.RegisterCommandsWithDiscord(appId, botClient)

	botServer.Listen(3000)
}
