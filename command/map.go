package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jviguy/speedycmds/utils"
	"strings"
)

// Map is an implementation of a command map.
type Map struct {
	commands map[string]Command
	groups   map[string]Group
}

// RegisterCommandGroup registers a command group by its name to the command map.
func (m *Map) RegisterCommandGroup(group Group) {
	if !m.GroupExists(group.Name()) && m.CanRegisterGroup(group) {
		m.groups[strings.ToLower(group.Name())] = group
	}
}

// Groups returns the command map's groups.
func (m *Map) Groups() map[string]Group {
	return m.groups
}

// Group returns a command map group by its name.
func (m *Map) Group(name string) Group {
	if m.GroupExists(name) {
		return m.groups[name]
	}
	return nil
}

// CanRegisterGroup returns true if a group can be registered under the name passed.
func (m *Map) CanRegisterGroup(group Group) bool {
	return m.commands[group.Name()] == nil && !m.GroupExists(group.Name())
}

// GroupExists returns true if the group passed exists.
func (m *Map) GroupExists(name string) bool {
	_, b := m.groups[name]
	return b
}

// Execute executes a command on the command map using the parameters passed.
func (m *Map) Execute(command string, ctx Context, session *discordgo.Session) error {
	switch true {
	case m.CanExecute(command):
		return m.commands[strings.ToLower(command)].Execute(ctx, session)
	case m.GroupExists(command):
		if len(ctx.Arguments()) > 0 {
			args, cmd := utils.Shift(ctx.Arguments(), 0)
			g := m.Group(command)
			if g.CanExecute(cmd) {
				ct := NewBasicContext(args, ctx.Message(), session, m, &g)
				//custom ctx for the custom args needed
				return m.Group(command).Execute(cmd, ct, session)
			}
		}
		fallthrough
	default:
		em := discordgo.MessageEmbed{}
		em.Title = "Unknown Command: " + command
		em.Description = "You might have Meant: " + utils.FindClosest(command, m.GetAllCommandNames())
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, &em)
		return nil
	}
}

func (m *Map) GetAllCommandNames() []string {
	s := make([]string, len(m.commands))
	for name, _ := range m.commands {
		s = append(s, name)
	}
	return s
}

// GroupedCommands returns all grouped commands in the command map.
func (m *Map) GroupedCommands() map[string]Command {
	cs := m.Commands()
	for k, g := range m.Groups() {
		for name, cmd := range g.Commands() {
			cs[k+" "+name] = cmd
		}
	}

	return cs
}

// Commands returns all commands in the command map.
func (m *Map) Commands() map[string]Command {
	return m.commands
}

// RegisterCommand registers a command to the command map using the information passed.
func (m *Map) RegisterCommand(command Command, override bool) {
	if m.CanRegisterCommand(command) || override {
		m.commands[strings.ToLower(command.Name())] = command
	}
}

// CanRegisterCommand returns true if a command can be registered under the name passed.
func (m *Map) CanRegisterCommand(cmd Command) bool {
	return m.commands[cmd.Name()] == nil && !m.GroupExists(cmd.Name())
}

// CanExecute returns true if a command can be executed by the name passed.
func (m *Map) CanExecute(name string) bool {
	_, ok := m.commands[name]
	return ok
}

// NewCommandMap returns a new implementation of a command map.
func NewCommandMap() *Map {
	return &Map{commands: map[string]Command{}, groups: map[string]Group{}}
}
