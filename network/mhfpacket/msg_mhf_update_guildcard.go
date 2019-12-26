package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfUpdateGuildcard represents the MSG_MHF_UPDATE_GUILDCARD
type MsgMhfUpdateGuildcard struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateGuildcard) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_GUILDCARD
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateGuildcard) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateGuildcard) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}