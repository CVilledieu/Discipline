package messages

import "encoding/binary"

/*
byte [0:4] - user id
byte [4:8] - chat id
byte [8:12] - time stamp
byte [12:20] - messageId
byte [20: 22] - message length
byte [20:] - message
*/

type message []byte

func UserId(m message) uint32 {
	return binary.LittleEndian.Uint32(n[0:4])
}

func ChatId(m message) uint32 {
	return binary.LittleEndian.Uint32(n[4:8])
}

func (m *message) getTimeStamp() uint32 {
	return binary.LittleEndian.Uint32((*m)[8:12])
}

func (m *message) setTimeStamp(ts uint32) {
	binary.LittleEndian.PutUint32((*m)[8:12], ts)
}

func (m *message) getMessageId() uint64 {
	return binary.LittleEndian.Uint64((*m)[12:20])
}

func (m *message) setMessageId(id uint64) {
	binary.LittleEndian.PutUint64((*m)[12:20], id)
}

func setHeader(m message, userId uint32, chatId uint32) {
	binary.LittleEndian.PutUint32(m[0:4], userId)
	binary.LittleEndian.PutUint32(m[4:8], chatId)
}

func (m *message) getMessage() []byte {
	return (*m)[20:]
}
