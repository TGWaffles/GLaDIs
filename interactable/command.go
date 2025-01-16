package interactable

import (
	"github.com/JackHumphries9/dapper-go/client"
)

type CommandOptions struct {
	ExperimentalAutoDeferral bool
}

type Command struct {
	Command              client.CreateApplicationCommand
	AssociatedComponents []Component
	AssociatedModals     []Modal
	CommandOptions       CommandOptions
	OnCommand            InteractionHandler
}

func (dc *Command) AddComponent(component Component) {
	dc.AssociatedComponents = append(dc.AssociatedComponents, component)
}

func (dc *Command) GetComponents() []Component {
	return dc.AssociatedComponents
}

func (dc *Command) AddModal(modal Modal) {
	dc.AssociatedModals = append(dc.AssociatedModals, modal)
}

func (dc *Command) GetModals() []Modal {
	return dc.AssociatedModals
}
