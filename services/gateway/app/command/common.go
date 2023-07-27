package command

import (
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
)

type ChannelCommandManager struct {
	commandChannels map[string]chan *command_utils.CommandMessage
}

func (c *ChannelCommandManager) SetCommandChannel(commandMessage *command_utils.CommandMessage) chan *command_utils.CommandMessage {

	ch := make(chan *command_utils.CommandMessage)
	c.commandChannels[commandMessage.Hash] = ch
	return ch
}

func (c *ChannelCommandManager) DeleteCommandChannelOLD(hash string) chan *command_utils.CommandMessage {
	ch := c.commandChannels[hash]
	delete(c.commandChannels, hash)
	return ch
}
func (c *ChannelCommandManager) DeleteCommandChannel(commandMessage *command_utils.CommandMessage) {
	ch := c.commandChannels[commandMessage.Hash]
	delete(c.commandChannels, commandMessage.Hash)
	ch <- commandMessage
	close(ch)
}

func newCommandChannelManager() *ChannelCommandManager {
	mapChannel := make(map[string]chan *command_utils.CommandMessage)
	return &ChannelCommandManager{
		commandChannels: mapChannel,
	}
}
