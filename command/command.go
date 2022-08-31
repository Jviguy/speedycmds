package command

import (
	"github.com/bwmarrin/discordgo"
)

// Command is the simple command interface for commands.
type Command interface {
	// Execute is called when a command is attempted.
	Execute(context Context, session *discordgo.Session) error
	// Name returns the given name of the command
	Name() string
	// Description returns the description of the command
	Description() string
}
