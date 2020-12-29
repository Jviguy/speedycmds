package command

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
)
//the simple command Interface for executing code when its called.
type Command interface {
	Execute(ctx ctx.Ctx,session *discordgo.Session) error
}

