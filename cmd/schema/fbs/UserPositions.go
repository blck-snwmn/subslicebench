// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UserPositionsT struct {
	Poss []*UserPositionT `json:"poss"`
}

func (t *UserPositionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	possOffset := flatbuffers.UOffsetT(0)
	if t.Poss != nil {
		possLength := len(t.Poss)
		possOffsets := make([]flatbuffers.UOffsetT, possLength)
		for j := 0; j < possLength; j++ {
			possOffsets[j] = t.Poss[j].Pack(builder)
		}
		UserPositionsStartPossVector(builder, possLength)
		for j := possLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(possOffsets[j])
		}
		possOffset = builder.EndVector(possLength)
	}
	UserPositionsStart(builder)
	UserPositionsAddPoss(builder, possOffset)
	return UserPositionsEnd(builder)
}

func (rcv *UserPositions) UnPackTo(t *UserPositionsT) {
	possLength := rcv.PossLength()
	t.Poss = make([]*UserPositionT, possLength)
	for j := 0; j < possLength; j++ {
		x := UserPosition{}
		rcv.Poss(&x, j)
		t.Poss[j] = x.UnPack()
	}
}

func (rcv *UserPositions) UnPack() *UserPositionsT {
	if rcv == nil { return nil }
	t := &UserPositionsT{}
	rcv.UnPackTo(t)
	return t
}

type UserPositions struct {
	_tab flatbuffers.Table
}

func GetRootAsUserPositions(buf []byte, offset flatbuffers.UOffsetT) *UserPositions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UserPositions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUserPositions(buf []byte, offset flatbuffers.UOffsetT) *UserPositions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &UserPositions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *UserPositions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UserPositions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *UserPositions) Poss(obj *UserPosition, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *UserPositions) PossLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func UserPositionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func UserPositionsAddPoss(builder *flatbuffers.Builder, poss flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(poss), 0)
}
func UserPositionsStartPossVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func UserPositionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
