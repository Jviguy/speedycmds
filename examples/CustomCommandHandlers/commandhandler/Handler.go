package commandhandler

import (
	"github.com/Jviguy/GoingCommando"
	"github.com/Jviguy/GoingCommando/command/commandmap"
	"github.com/Jviguy/GoingCommando/command/ctx"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Handler struct {
	dg *discordgo.Session
	cmds commandmap.Map
	Prefixs []string
}

func (h Handler) Handle(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.Bot {
		return
	}
	for _,value := range h.Prefixs {
		if strings.HasPrefix(msg.Content, value) {
			args := strings.Split(strings.TrimPrefix(msg.Content, value), " ")
			args, cmd := GoingCommando.Shift(args, 0)
			err := h.cmds.Execute(cmd, ctx.New(args, msg, s), s)
			if err != nil {
				_, err = s.ChannelMessageSend(msg.ChannelID, "An Error Occurred while executing that command"+
					"\nDEBUG: "+err.Error())
			}
			break
		}
	}
}

func (h Handler) GetCommandHandler() commandmap.Map {
	return h.cmds
}

func New(session *discordgo.Session,cmds commandmap.Map,prefix []string) Handler {
	h := Handler{dg: session,cmds: cmds,Prefixs: prefix}
	session.AddHandler(h.Handle)
	return h
}


