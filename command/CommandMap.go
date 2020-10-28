package command

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type CmdMap interface {
	Execute(command string,args []string,s *discordgo.Session) error
	RegisterCommand(name string,command Command,override bool)
	CanRegisterCommand(name string) bool
	GetCommands() map[string]Command
}

type Map struct {
	commands map[string]Command
}

func (m Map) Execute(command string,args []string,s *discordgo.Session) error {
	return m.commands[command].Execute(args,s)
}

func (m Map) RegisterCommand(name string,command Command, override bool) {
	if m.CanRegisterCommand(name) || override{
		m.commands[strings.ToLower(name)] = command
	}
}

func (m Map) CanRegisterCommand(name string) bool {
	return m.commands[name] == nil
}

func (m Map) GetCommands() map[string]Command {
	return m.commands
}