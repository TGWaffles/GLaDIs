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
	Command              client.CreateApplicationCommand
	associatedComponents []DapperComponent
	CommandOptions       DapperCommandOptions
	OnCommand            DapperCommandExecutor
}

func (dc *DapperCommand) AddComponent(component DapperComponent) {
	dc.associatedComponents = append(dc.associatedComponents, component)
}

func (dc *DapperCommand) GetComponents() []DapperComponent {
	return dc.associatedComponents
}
