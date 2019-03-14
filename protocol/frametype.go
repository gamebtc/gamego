package protocol

import (
	_ "encoding/binary"
)

const (
	FrameTypeLen = 2
)

//var (
//	FrameType_MessageHead       = make([]byte, FrameTypeLen)
//	FrameType_ForwardHead       = make([]byte, FrameTypeLen)
//	FrameType_ControlHead       = make([]byte, FrameTypeLen)
//	FrameType_BroadcastUserHead = make([]byte, FrameTypeLen)
//	FrameType_BroadcastRoomHead = make([]byte, FrameTypeLen)
//	FrameType_BroadcastGameHead = make([]byte, FrameTypeLen)
//	FrameType_BroadcastAllHead  = make([]byte, FrameTypeLen)
//)
//
//func init() {
//	binary.LittleEndian.PutUint16(FrameType_MessageHead, uint16(FrameType_Message))
//	binary.LittleEndian.PutUint16(FrameType_ForwardHead, uint16(FrameType_Forward))
//	binary.LittleEndian.PutUint16(FrameType_ControlHead, uint16(FrameType_Control))
//	binary.LittleEndian.PutUint16(FrameType_BroadcastUserHead, uint16(FrameType_BroadcastUser))
//	binary.LittleEndian.PutUint16(FrameType_BroadcastRoomHead, uint16(FrameType_BroadcastRoom))
//	binary.LittleEndian.PutUint16(FrameType_BroadcastGameHead, uint16(FrameType_BroadcastGame))
//	binary.LittleEndian.PutUint16(FrameType_BroadcastAllHead, uint16(FrameType_BroadcastAll))
//}
