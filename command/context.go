package command

import "github.com/bwmarrin/discordgo"

// Context is a context for a command execute attempt. It contains the necessary information.
type Context interface {
	// Arguments ...
	Arguments() []string
	// Author ...
	Author() *discordgo.User
	// Channel ...
	Channel() *discordgo.Channel
	// Guild ...
	Guild() *discordgo.Guild
	// Message ...
	Message() *discordgo.MessageCreate
}

// BasicContext is a basic context implementation.
type BasicContext struct {
	args    []string
	author  *discordgo.User
	channel *discordgo.Channel
	guild   *discordgo.Guild
	message *discordgo.MessageCreate
}

// Arguments ...
func (b BasicContext) Arguments() []string {
	return b.args
}

// Author ...
func (b BasicContext) Author() *discordgo.User {
	return b.author
}

// Channel ...
func (b BasicContext) Channel() *discordgo.Channel {
	return b.channel
}

// Guild ...
func (b BasicContext) Guild() *discordgo.Guild {
	return b.guild
}

// Message ...
func (b BasicContext) Message() *discordgo.MessageCreate {
	return b.message
}

// NewBasicContext creates a new basic context and returns it.
func NewBasicContext(args []string, message *discordgo.MessageCreate, session *discordgo.Session) *BasicContext {
	ctx := &BasicContext{args: args}
	ctx.author = message.Author
	ctx.channel, _ = session.Channel(message.ChannelID)
	ctx.guild, _ = session.Guild(message.GuildID)
	ctx.message = message

	return ctx
}
