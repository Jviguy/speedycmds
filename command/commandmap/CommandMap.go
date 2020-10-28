package commandmap

import (
	"github.com/Jviguy/GoingCommando/command"
	"github.com/Jviguy/GoingCommando/command/ctx"
	"github.com/Jviguy/GoingCommando/utils"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type CmdMap interface {
	Execute(command string,ctx ctx.Ctx,s *discordgo.Session) error
	RegisterCommand(name string,command command.Command,override bool)
	CanRegisterCommand(name string) bool
	GetCommands() map[string]command.Command
	CanExecute(name string) bool
}

type Map struct {
	commands map[string]command.Command
}

func (m Map) Execute(command string,ctx ctx.Ctx,s *discordgo.Session) error {
	if m.CanExecute(command) {
		return m.commands[strings.ToLower(command)].Execute(ctx,s)
	}
	var closest string
	go func() {
		closest = utils.FindClosest(command,utils.GetAllKeys(m.commands))
	}()
	em := discordgo.MessageEmbed{}
	em.Title = "Unknown Command: " + command
	em.Description = "You might have Meant: " + closest
	_,_ = s.ChannelMessageSendEmbed(ctx.GetChannel().ID,&em)
	return nil
}

func (m Map) RegisterCommand(name string,command command.Command, override bool) {
	if m.CanRegisterCommand(name) || override{
		m.commands[strings.ToLower(name)] = command
	}
}

func (m Map) CanRegisterCommand(name string) bool {
	return m.commands[name] == nil
}

func (m Map) GetCommands() map[string]command.Command {
	return m.commands
}

func New() Map {
	return Map{commands: map[string]command.Command{}}
}

func (m Map) CanExecute(name string) bool {
	_,ok := m.commands[name]
	return ok
}