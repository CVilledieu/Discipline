package user

import (
	"encoding/binary"
)

const (
	uHEAD = 14
)

type Node []byte

func setUserId(n Node, id uint32) {
	binary.LittleEndian.PutUint32(n[0:4], id)
}

func UserId(n Node) uint32 {
	return binary.LittleEndian.Uint32(n[0:4])
}
func (n *Node) getChildPtr(i uint16) uint64 {
	return binary.LittleEndian.Uint64((*n)[4:12])
}

func (n *Node) setChildPtr(i uint16, ptr uint64) {
	binary.LittleEndian.PutUint64((*n)[4:12], ptr)
}

func (n *Node) nKeys() uint16 {
	return binary.LittleEndian.Uint16((*n)[12:14])
}

func (n *Node) offsetPos(i uint16) uint16 {
	return uHEAD + (i)*2
}

func (n *Node) getOffset(i uint16) uint16 {
	if i == 0 {
		return 0
	}
	oPos := n.offsetPos(i)
	return binary.LittleEndian.Uint16((*n)[oPos:])
}

func (n *Node) setOffset(i uint16, offset uint16) {
	oPos := n.offsetPos(i)
	binary.LittleEndian.PutUint16((*n)[oPos:], offset)
}

func (n *Node) kvPos(i uint16) uint16 {
	return uHEAD + n.nKeys()*2 + n.getOffset(i)
}

func (n *Node) getKey(i uint16) uint64 {
	return binary.LittleEndian.Uint64((*n)[n.kvPos(i):])
}

func (n *Node) setKey(i uint16, key uint64) {
	binary.LittleEndian.PutUint64((*n)[n.kvPos(i):], key)
}

func (n *Node) getValueLen(i uint16) uint16 {
	pos := n.kvPos(i) + 8
	return binary.LittleEndian.Uint16((*n)[pos:])
}

func (n *Node) getValue(i uint16) []byte {
	pos := n.kvPos(i) + 8 + 2
	return (*n)[pos : pos+n.getValueLen(i)]
}

func (n *Node) getKeyRange() (lowest, highest uint64) {
	return n.getKey(0), n.getKey(n.nKeys() - 1)
}
