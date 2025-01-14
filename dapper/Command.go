package dapper

import (
	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/discord"
)

type DapperCommandExecutor func(itx *discord.Interaction)

type DapperCommandOptions struct {
	Ephemeral bool
}

type DapperCommand struct {
	Command        client.CreateApplicationCommand
	CommandOptions DapperCommandOptions
	OnCommand       DapperCommandExecutor
}
