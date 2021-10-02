package speedycmds

import (
	"github.com/Jviguy/SpeedyCmds/command"
	"github.com/Jviguy/SpeedyCmds/utils"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

// commandRegex is used to match a message for command arguments.
var commandRegex = regexp.MustCompile(`(?m)("[^"]+"|[^\s"]+)`)

// Handler is an abstract version of a handler.
type Handler interface {
	// Handle handles a message. It is given the session and message creation information.
	Handle(session *discordgo.Session, message *discordgo.MessageCreate)
	// Commands returns the command map.
	Commands() *command.Map
}

// BasicHandler is a basic handler with support for one prefix. You can change it up how you want because of the Handler interface.
type BasicHandler struct {
	// Regex is true if the bot should check for the arguments using regex or the strings package.
	Regex bool
	// Prefix is the bot prefix.
	Prefix string

	session  *discordgo.Session
	commands *command.Map
}

// NewBasicHandler initializes a new basic handler and returns it.
func NewBasicHandler(session *discordgo.Session, regex bool, prefix string, commands *command.Map) *BasicHandler {
	h := &BasicHandler{
		Regex:  regex,
		Prefix: prefix,

		session:  session,
		commands: commands,
	}
	session.AddHandler(h.Handle)
	return h
}

// Handle handles a message. It is given the session and message creation information.
func (h *BasicHandler) Handle(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.Bot {
		return
	}
	if strings.HasPrefix(msg.Content, h.Prefix) {
		args := strings.Split(strings.TrimPrefix(msg.Content, h.Prefix), " ")
		if h.Regex {
			args = commandRegex.FindAllString(strings.TrimPrefix(msg.Content, h.Prefix), -1)
		}

		args, cmd := utils.Shift(args, 0)
		err := h.commands.Execute(cmd, command.NewBasicContext(args, msg, s), s)
		if err != nil {
			em := &discordgo.MessageEmbed{}
			em.Title = "An Error Occurred while executing that command"
			em.Description = err.Error()
			em.Color = 16711680
			_, _ = s.ChannelMessageSendEmbed(msg.ChannelID, em)
		}
	}
}

// Commands returns the command map.
func (h *BasicHandler) Commands() *command.Map {
	return h.commands
}
