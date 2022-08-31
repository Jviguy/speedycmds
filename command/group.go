package command

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// Group is a command group used for implementing sub-categories for commands.
type Group interface {
	// Name returns the command group's name.
	Name() string
	// Description returns a description about this group
	Description() string
	// AddCommand adds a command with the passed name and command.
	AddCommand(name string, command Command)
	// Command returns a command by its name.
	Command(name string) Command
	// Commands returns the command map index.
	Commands() map[string]Command
	// CanExecute returns true if the command name passed can be executed.
	CanExecute(name string) bool
	// Execute executes a command by its name, context, and session.
	Execute(name string, ctx Context, session *discordgo.Session) error
}

// BasicGroup is a basic group implementation.
type BasicGroup struct {
	name        string
	description string
	commands    map[string]Command
}

// Name ...
func (c *BasicGroup) Name() string {
	return c.name
}

// Description ...
func (c *BasicGroup) Description() string {
	return c.description
}

// AddCommand ...
func (c *BasicGroup) AddCommand(cmd Command) {
	c.commands[strings.ToLower(cmd.Name())] = cmd
}

// AddCommands adds a list of commands.
func (c *BasicGroup) AddCommands(commands []Command) {
	for _, cmd := range commands {
		c.AddCommand(cmd)
	}
}

// Command ...
func (c *BasicGroup) Command(name string) Command {
	return c.commands[name]
}

// Commands ...
func (c *BasicGroup) Commands() map[string]Command {
	return c.commands
}

// CanExecute ...
func (c *BasicGroup) CanExecute(name string) bool {
	_, ok := c.commands[name]
	return ok
}

// Execute ...
func (c *BasicGroup) Execute(name string, ctx Context, session *discordgo.Session) error {
	return c.commands[name].Execute(ctx, session)
}

// NewBasicGroup initializes a new basic group.
func NewBasicGroup(name string, description string) *BasicGroup {
	return &BasicGroup{name: name, commands: map[string]Command{}, description: description}
}
