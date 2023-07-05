// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Points3D

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Rgb struct {
	_tab flatbuffers.Struct
}

func (rcv *Rgb) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Rgb) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Rgb) R() byte {
	return rcv._tab.GetByte(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Rgb) MutateR(n byte) bool {
	return rcv._tab.MutateByte(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Rgb) G() byte {
	return rcv._tab.GetByte(rcv._tab.Pos + flatbuffers.UOffsetT(1))
}
func (rcv *Rgb) MutateG(n byte) bool {
	return rcv._tab.MutateByte(rcv._tab.Pos+flatbuffers.UOffsetT(1), n)
}

func (rcv *Rgb) B() byte {
	return rcv._tab.GetByte(rcv._tab.Pos + flatbuffers.UOffsetT(2))
}
func (rcv *Rgb) MutateB(n byte) bool {
	return rcv._tab.MutateByte(rcv._tab.Pos+flatbuffers.UOffsetT(2), n)
}

func CreateRgb(builder *flatbuffers.Builder, r byte, g byte, b byte) flatbuffers.UOffsetT {
	builder.Prep(1, 3)
	builder.PrependByte(b)
	builder.PrependByte(g)
	builder.PrependByte(r)
	return builder.Offset()
}
