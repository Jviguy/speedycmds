package commandmap

import (
	"github.com/Jviguy/GoingCommando"
	"github.com/Jviguy/GoingCommando/command"
	"github.com/Jviguy/GoingCommando/command/commandGroup"
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
	groups map[string]commandGroup.Group
}
func (m Map) RegisterCommandGroup(name string,group commandGroup.Group){
	if !m.DoesGroupExist(name) && m.CanRegisterGroup(name){
		m.groups[name] = group
	}
}

func (m Map) GetGroups() map[string]commandGroup.Group {
	return m.groups
}

func (m Map) GetGroup(name string) commandGroup.Group {
	if m.DoesGroupExist(name) {
		return m.groups[name]
	}
	return nil
}

func (m Map) CanRegisterGroup(name string) bool {
	return m.commands[name] == nil && m.GetGroup(name) == nil
}

func (m Map) DoesGroupExist(name string) bool {
	_,b := m.groups[name]
	return b
}

func (m Map) Execute(command string,c ctx.Ctx,s *discordgo.Session) error {
	switch true {
	case m.CanExecute(command):
		return m.commands[strings.ToLower(command)].Execute(c,s)
	case m.DoesGroupExist(command):
		args,cmd := GoingCommando.Shift(c.GetArgs(),0)
		ct := ctx.New(args,c.GetMessage(),s)
		//custom ctx for the custom args needed
		return m.GetGroup(cmd).Execute(command,ct,s)
	default:
		em := discordgo.MessageEmbed{}
		em.Title = "Unknown Command: " + command
		em.Description = "You might have Meant: " + utils.FindClosest(command, utils.GetAllKeysCommands(m.GetAllCommands()))
		_,_ = s.ChannelMessageSendEmbed(c.GetChannel().ID,&em)
		return nil
	}
}

func (m Map) GetAllCommands() map[string]command.Command {
	cs := m.GetCommands()
	for k,g := range m.GetGroups(){
		for name,cmd := range g.GetCommands(){
			cs[k+" "+name] = cmd
		}
	}
	return cs
}

func (m Map) RegisterCommand(name string,command command.Command, override bool) {
	if m.CanRegisterCommand(name) || override{
		m.commands[strings.ToLower(name)] = command
		//for when someone doesnt put it in the struct
		if command.GetName() == ""{
			command.Setname(name)
		}
	}
}

func (m Map) CanRegisterCommand(name string) bool {
	return m.commands[name] == nil && m.GetGroup(name) == nil
}

func (m Map) GetCommands() map[string]command.Command {
	return m.commands
}

func New() Map {
	return Map{commands: map[string]command.Command{},groups: map[string]commandGroup.Group{}}
}

func (m Map) CanExecute(name string) bool {
	_,ok := m.commands[name]
	return ok
}