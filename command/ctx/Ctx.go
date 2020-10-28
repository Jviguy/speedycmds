package ctx

import "github.com/bwmarrin/discordgo"

type Ctx interface {
	GetArgs() []string
	GetAuthor() *discordgo.User
	GetChannel() *discordgo.Channel
	GetGuild() *discordgo.Guild
	GetMessage() *discordgo.MessageCreate
}

type BaseCtx struct {
	args []string
	author *discordgo.User
	channel *discordgo.Channel
	guild *discordgo.Guild
	message *discordgo.MessageCreate
}

func (b BaseCtx) GetArgs() []string {
	return b.args
}

func (b BaseCtx) GetAuthor() *discordgo.User {
	return b.author
}

func (b BaseCtx) GetChannel() *discordgo.Channel {
	return b.channel
}

func (b BaseCtx) GetGuild() *discordgo.Guild {
	return b.guild
}

func (b BaseCtx) GetMessage() *discordgo.MessageCreate {
	return b.message
}

func New(args []string,msg *discordgo.MessageCreate,session *discordgo.Session) BaseCtx {
	ctx := BaseCtx{args: args}
	ctx.author = msg.Author
	ctx.channel,_ = session.Channel(msg.ChannelID)
	ctx.guild,_ = session.Guild(msg.GuildID)
	ctx.message = msg
	return ctx
}
