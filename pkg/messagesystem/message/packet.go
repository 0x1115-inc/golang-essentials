package message

import (
	"encoding/json"
)

// GeneralPacket is a abstract struct that can be used to create a packet
// The packet can be used to send messages between different systems
type GeneralPacket struct {
	Header     map[string]IValue
	Attributes map[string]IValue
	IData
}

func (g *GeneralPacket) SetHeader(key string, value IValue) {
	if g.Header == nil {
		g.Header = make(map[string]IValue)
	}
	g.Header[key] = value
}

func (g *GeneralPacket) GetHeader(key string) IValue {
	if g.Header == nil {
		return nil
	}
	return g.Header[key]
}

func (g *GeneralPacket) SetAttribute(key string, value IValue) {
	if g.Attributes == nil {
		g.Attributes = make(map[string]IValue)
	}
	g.Attributes[key] = value
}

func (g *GeneralPacket) GetAttribute(key string) IValue {
	if g.Attributes == nil {
		return nil
	}
	return g.Attributes[key]
}

func (g *GeneralPacket) SetData(data IData) {
	g.IData = data
}

func (g *GeneralPacket) GetData() IData {
	return g.IData
}

func (g *GeneralPacket) String() string {
	// Json marshal the struct
	jsonBytes, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

type PacketData struct {
	// Data is the actual message data
	Data map[string]IValue
}

func (p *PacketData) GetField(key string) IValue {
	if p.Data == nil {
		return nil
	}
	return p.Data[key]
}
