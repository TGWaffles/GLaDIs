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
	AssociatedComponents []DapperComponent
	CommandOptions       DapperCommandOptions
	OnCommand            DapperCommandExecutor
}

func (dc *DapperCommand) AddComponent(component DapperComponent) {
	dc.AssociatedComponents = append(dc.AssociatedComponents, component)
}

func (dc *DapperCommand) GetComponents() []DapperComponent {
	return dc.AssociatedComponents
}
