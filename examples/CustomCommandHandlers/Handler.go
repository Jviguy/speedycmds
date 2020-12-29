package main

import (
	"github.com/Jviguy/SpeedyCmds"
	"github.com/Jviguy/SpeedyCmds/command/commandMap"
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Handler struct {
	dg *discordgo.Session
	cmds *commandMap.Map
	Prefixs []string
}

func (h Handler) Handle(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.Bot {
		return
	}
	for _,value := range h.Prefixs {
		if strings.HasPrefix(msg.Content, value) {
			args := strings.Split(strings.TrimPrefix(msg.Content, value), " ")
			args, cmd := SpeedyCmds.Shift(args, 0)
			err := h.cmds.Execute(cmd, ctx.New(args, msg, s), s)
			if err != nil {
				_, err = s.ChannelMessageSend(msg.ChannelID, "An Error Occurred while executing that command"+
					"\nDEBUG: "+err.Error())
			}
			break
		}
	}
}

func (h Handler) GetCommandMap() *commandMap.Map {
	return h.cmds
}

func New(session *discordgo.Session,cmds *commandMap.Map,prefix []string) Handler {
	h := Handler{dg: session,cmds: cmds,Prefixs: prefix}
	session.AddHandler(h.Handle)
	return h
}


