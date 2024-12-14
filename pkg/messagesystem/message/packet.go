package message

import (
	"encoding/json"
)

// ConcretePacket is the struct that implements the IPacket interface 
type ConcretePacket struct {
	Header     map[string]IValue
	Attributes map[string]IValue
	IData
}

func (g *ConcretePacket) SetHeader(key string, value IValue) {
	if g.Header == nil {
		g.Header = make(map[string]IValue)
	}
	g.Header[key] = value
}

func (g *ConcretePacket) GetHeader(key string) IValue {
	if g.Header == nil {
		return nil
	}
	return g.Header[key]
}

func (g *ConcretePacket) GetHeaders() map[string]IValue {
	return g.Header
}

func (g *ConcretePacket) SetAttribute(key string, value IValue) {
	if g.Attributes == nil {
		g.Attributes = make(map[string]IValue)
	}
	g.Attributes[key] = value
}

func (g *ConcretePacket) GetAttribute(key string) IValue {
	if g.Attributes == nil {
		return nil
	}
	return g.Attributes[key]
}

func (g *ConcretePacket) GetAttributes() map[string]IValue {
	return g.Attributes
}

func (g *ConcretePacket) SetData(data IData) {
	g.IData = data
}

func (g *ConcretePacket) GetData() IData {
	return g.IData
}

func (g *ConcretePacket) String() string {
	// Json marshal the struct
	jsonBytes, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
