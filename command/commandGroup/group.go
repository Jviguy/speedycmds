package commandGroup

import (
	"github.com/Jviguy/GoingCommando/command"
	"github.com/Jviguy/GoingCommando/command/ctx"
	"github.com/bwmarrin/discordgo"
)

type Group interface {
	GetName() string
	AddCommand(name string,cmd command.Command)
	GetCommand(cmd string) command.Command
	GetCommands() map[string]command.Command
	CanExecute(cmd string) bool
	Execute(cmd string,ctx ctx.Ctx,session *discordgo.Session) error
}

type CmdGroup struct {
	name string
	//the registered commands added to this group this works outside the main "Command Map" while the group is in it
	commands map[string]command.Command
}

func (c CmdGroup) CanExecute(cmd string) bool {
	_,ok := c.GetCommands()[cmd]
	return ok
}

func (c CmdGroup) GetName() string {
	return c.name
}

func (c CmdGroup) AddCommand(name string,cmd command.Command) {
	c.commands[name] = cmd
}

func (c CmdGroup) GetCommand(cmd string) command.Command {
	return c.commands[cmd]
}

func (c CmdGroup) Execute(cmd string,ctx ctx.Ctx,session *discordgo.Session) error {
	return c.GetCommand(cmd).Execute(ctx,session)
}

func (c CmdGroup) GetCommands() map[string]command.Command {
	return c.commands
}

func New(name string) CmdGroup {
	return CmdGroup{name: name,commands: map[string]command.Command{}}
}

