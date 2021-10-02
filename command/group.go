package command

import (
	"github.com/bwmarrin/discordgo"
)

// Group is a command group used for implementing sub-categories for commands.
type Group interface {
	// Name returns the command group's name.
	Name() string
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
	name     string
	commands map[string]Command
}

// Name ...
func (c *BasicGroup) Name() string {
	return c.name
}

// AddCommand ...
func (c *BasicGroup) AddCommand(name string, cmd Command) {
	c.commands[name] = cmd
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
func NewBasicGroup(name string) *BasicGroup {
	return &BasicGroup{name: name, commands: map[string]Command{}}
}
