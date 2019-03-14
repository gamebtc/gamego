package protocol

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *LoginRoomAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "kind":
			z.Kind, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "room":
			z.Room, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "tab":
			z.Tab, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "seat":
			z.Seat, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "code":
			z.Code, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "msg":
			z.Msg, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *LoginRoomAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "kind"
	err = en.Append(0x86, 0xa4, 0x6b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Kind)
	if err != nil {
		return
	}
	// write "room"
	err = en.Append(0xa4, 0x72, 0x6f, 0x6f, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Room)
	if err != nil {
		return
	}
	// write "tab"
	err = en.Append(0xa3, 0x74, 0x61, 0x62)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Tab)
	if err != nil {
		return
	}
	// write "seat"
	err = en.Append(0xa4, 0x73, 0x65, 0x61, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Seat)
	if err != nil {
		return
	}
	// write "code"
	err = en.Append(0xa4, 0x63, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Code)
	if err != nil {
		return
	}
	// write "msg"
	err = en.Append(0xa3, 0x6d, 0x73, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Msg)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *LoginRoomAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "kind"
	o = append(o, 0x86, 0xa4, 0x6b, 0x69, 0x6e, 0x64)
	o = msgp.AppendInt32(o, z.Kind)
	// string "room"
	o = append(o, 0xa4, 0x72, 0x6f, 0x6f, 0x6d)
	o = msgp.AppendInt32(o, z.Room)
	// string "tab"
	o = append(o, 0xa3, 0x74, 0x61, 0x62)
	o = msgp.AppendInt32(o, z.Tab)
	// string "seat"
	o = append(o, 0xa4, 0x73, 0x65, 0x61, 0x74)
	o = msgp.AppendInt32(o, z.Seat)
	// string "code"
	o = append(o, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt32(o, z.Code)
	// string "msg"
	o = append(o, 0xa3, 0x6d, 0x73, 0x67)
	o = msgp.AppendString(o, z.Msg)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LoginRoomAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "kind":
			z.Kind, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "room":
			z.Room, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "tab":
			z.Tab, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "seat":
			z.Seat, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "code":
			z.Code, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "msg":
			z.Msg, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *LoginRoomAck) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int32Size + 5 + msgp.Int32Size + 4 + msgp.Int32Size + 5 + msgp.Int32Size + 5 + msgp.Int32Size + 4 + msgp.StringPrefixSize + len(z.Msg)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LoginRoomReq) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbai uint32
	zbai, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbai > 0 {
		zbai--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "args":
			z.Args, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z LoginRoomReq) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "id"
	err = en.Append(0x82, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
	}
	// write "args"
	err = en.Append(0xa4, 0x61, 0x72, 0x67, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Args)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z LoginRoomReq) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "id"
	o = append(o, 0x82, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "args"
	o = append(o, 0xa4, 0x61, 0x72, 0x67, 0x73)
	o = msgp.AppendString(o, z.Args)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LoginRoomReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "args":
			z.Args, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z LoginRoomReq) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 5 + msgp.StringPrefixSize + len(z.Args)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RoomInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "id":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "kind":
			z.Kind, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "level":
			z.Level, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "cap":
			z.Cap, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "ante":
			z.Ante, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "doorMin":
			z.DoorMin, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "doorMax":
			z.DoorMax, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "stayMin":
			z.StayMin, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "stayMax":
			z.StayMax, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "playMin":
			z.PlayMin, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "playMax":
			z.PlayMax, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "coinKey":
			z.CoinKey, err = dc.ReadString()
			if err != nil {
				return
			}
		case "icon":
			z.Icon, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "online":
			z.Online, err = dc.ReadInt32()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *RoomInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 15
	// write "name"
	err = en.Append(0x8f, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "id"
	err = en.Append(0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
	}
	// write "kind"
	err = en.Append(0xa4, 0x6b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Kind)
	if err != nil {
		return
	}
	// write "level"
	err = en.Append(0xa5, 0x6c, 0x65, 0x76, 0x65, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Level)
	if err != nil {
		return
	}
	// write "cap"
	err = en.Append(0xa3, 0x63, 0x61, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Cap)
	if err != nil {
		return
	}
	// write "ante"
	err = en.Append(0xa4, 0x61, 0x6e, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Ante)
	if err != nil {
		return
	}
	// write "doorMin"
	err = en.Append(0xa7, 0x64, 0x6f, 0x6f, 0x72, 0x4d, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.DoorMin)
	if err != nil {
		return
	}
	// write "doorMax"
	err = en.Append(0xa7, 0x64, 0x6f, 0x6f, 0x72, 0x4d, 0x61, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.DoorMax)
	if err != nil {
		return
	}
	// write "stayMin"
	err = en.Append(0xa7, 0x73, 0x74, 0x61, 0x79, 0x4d, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.StayMin)
	if err != nil {
		return
	}
	// write "stayMax"
	err = en.Append(0xa7, 0x73, 0x74, 0x61, 0x79, 0x4d, 0x61, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.StayMax)
	if err != nil {
		return
	}
	// write "playMin"
	err = en.Append(0xa7, 0x70, 0x6c, 0x61, 0x79, 0x4d, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.PlayMin)
	if err != nil {
		return
	}
	// write "playMax"
	err = en.Append(0xa7, 0x70, 0x6c, 0x61, 0x79, 0x4d, 0x61, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.PlayMax)
	if err != nil {
		return
	}
	// write "coinKey"
	err = en.Append(0xa7, 0x63, 0x6f, 0x69, 0x6e, 0x4b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.CoinKey)
	if err != nil {
		return
	}
	// write "icon"
	err = en.Append(0xa4, 0x69, 0x63, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Icon)
	if err != nil {
		return
	}
	// write "online"
	err = en.Append(0xa6, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Online)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RoomInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 15
	// string "name"
	o = append(o, 0x8f, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "id"
	o = append(o, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "kind"
	o = append(o, 0xa4, 0x6b, 0x69, 0x6e, 0x64)
	o = msgp.AppendInt32(o, z.Kind)
	// string "level"
	o = append(o, 0xa5, 0x6c, 0x65, 0x76, 0x65, 0x6c)
	o = msgp.AppendInt32(o, z.Level)
	// string "cap"
	o = append(o, 0xa3, 0x63, 0x61, 0x70)
	o = msgp.AppendInt32(o, z.Cap)
	// string "ante"
	o = append(o, 0xa4, 0x61, 0x6e, 0x74, 0x65)
	o = msgp.AppendInt64(o, z.Ante)
	// string "doorMin"
	o = append(o, 0xa7, 0x64, 0x6f, 0x6f, 0x72, 0x4d, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.DoorMin)
	// string "doorMax"
	o = append(o, 0xa7, 0x64, 0x6f, 0x6f, 0x72, 0x4d, 0x61, 0x78)
	o = msgp.AppendInt64(o, z.DoorMax)
	// string "stayMin"
	o = append(o, 0xa7, 0x73, 0x74, 0x61, 0x79, 0x4d, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.StayMin)
	// string "stayMax"
	o = append(o, 0xa7, 0x73, 0x74, 0x61, 0x79, 0x4d, 0x61, 0x78)
	o = msgp.AppendInt64(o, z.StayMax)
	// string "playMin"
	o = append(o, 0xa7, 0x70, 0x6c, 0x61, 0x79, 0x4d, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.PlayMin)
	// string "playMax"
	o = append(o, 0xa7, 0x70, 0x6c, 0x61, 0x79, 0x4d, 0x61, 0x78)
	o = msgp.AppendInt64(o, z.PlayMax)
	// string "coinKey"
	o = append(o, 0xa7, 0x63, 0x6f, 0x69, 0x6e, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.CoinKey)
	// string "icon"
	o = append(o, 0xa4, 0x69, 0x63, 0x6f, 0x6e)
	o = msgp.AppendInt32(o, z.Icon)
	// string "online"
	o = append(o, 0xa6, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65)
	o = msgp.AppendInt32(o, z.Online)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RoomInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zwht uint32
	zwht, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zwht > 0 {
		zwht--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "kind":
			z.Kind, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "level":
			z.Level, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "cap":
			z.Cap, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "ante":
			z.Ante, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "doorMin":
			z.DoorMin, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "doorMax":
			z.DoorMax, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "stayMin":
			z.StayMin, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "stayMax":
			z.StayMax, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "playMin":
			z.PlayMin, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "playMax":
			z.PlayMax, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "coinKey":
			z.CoinKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "icon":
			z.Icon, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "online":
			z.Online, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *RoomInfo) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 3 + msgp.Int32Size + 5 + msgp.Int32Size + 6 + msgp.Int32Size + 4 + msgp.Int32Size + 5 + msgp.Int64Size + 8 + msgp.Int64Size + 8 + msgp.Int64Size + 8 + msgp.Int64Size + 8 + msgp.Int64Size + 8 + msgp.Int64Size + 8 + msgp.Int64Size + 8 + msgp.StringPrefixSize + len(z.CoinKey) + 5 + msgp.Int32Size + 7 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RoomListAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcua uint32
	zcua, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcua > 0 {
		zcua--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "rooms":
			var zxhx uint32
			zxhx, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rooms) >= int(zxhx) {
				z.Rooms = (z.Rooms)[:zxhx]
			} else {
				z.Rooms = make([]*RoomInfo, zxhx)
			}
			for zhct := range z.Rooms {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Rooms[zhct] = nil
				} else {
					if z.Rooms[zhct] == nil {
						z.Rooms[zhct] = new(RoomInfo)
					}
					err = z.Rooms[zhct].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *RoomListAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "rooms"
	err = en.Append(0x81, 0xa5, 0x72, 0x6f, 0x6f, 0x6d, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Rooms)))
	if err != nil {
		return
	}
	for zhct := range z.Rooms {
		if z.Rooms[zhct] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Rooms[zhct].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RoomListAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "rooms"
	o = append(o, 0x81, 0xa5, 0x72, 0x6f, 0x6f, 0x6d, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rooms)))
	for zhct := range z.Rooms {
		if z.Rooms[zhct] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Rooms[zhct].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RoomListAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zlqf uint32
	zlqf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zlqf > 0 {
		zlqf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "rooms":
			var zdaf uint32
			zdaf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rooms) >= int(zdaf) {
				z.Rooms = (z.Rooms)[:zdaf]
			} else {
				z.Rooms = make([]*RoomInfo, zdaf)
			}
			for zhct := range z.Rooms {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Rooms[zhct] = nil
				} else {
					if z.Rooms[zhct] == nil {
						z.Rooms[zhct] = new(RoomInfo)
					}
					bts, err = z.Rooms[zhct].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *RoomListAck) Msgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zhct := range z.Rooms {
		if z.Rooms[zhct] == nil {
			s += msgp.NilSize
		} else {
			s += z.Rooms[zhct].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RoomListReq) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zpks uint32
	zpks, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zpks > 0 {
		zpks--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "kind":
			z.Kind, err = dc.ReadInt32()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z RoomListReq) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "kind"
	err = en.Append(0x81, 0xa4, 0x6b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Kind)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z RoomListReq) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "kind"
	o = append(o, 0x81, 0xa4, 0x6b, 0x69, 0x6e, 0x64)
	o = msgp.AppendInt32(o, z.Kind)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RoomListReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjfb uint32
	zjfb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjfb > 0 {
		zjfb--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "kind":
			z.Kind, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z RoomListReq) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SendRoomFail) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcxo uint32
	zcxo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcxo > 0 {
		zcxo--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "code":
			z.Code, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "msg":
			z.Msg, err = dc.ReadString()
			if err != nil {
				return
			}
		case "kind":
			z.Kind, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "room":
			z.Room, err = dc.ReadInt32()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SendRoomFail) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "code"
	err = en.Append(0x84, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Code)
	if err != nil {
		return
	}
	// write "msg"
	err = en.Append(0xa3, 0x6d, 0x73, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Msg)
	if err != nil {
		return
	}
	// write "kind"
	err = en.Append(0xa4, 0x6b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Kind)
	if err != nil {
		return
	}
	// write "room"
	err = en.Append(0xa4, 0x72, 0x6f, 0x6f, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Room)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SendRoomFail) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "code"
	o = append(o, 0x84, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt32(o, z.Code)
	// string "msg"
	o = append(o, 0xa3, 0x6d, 0x73, 0x67)
	o = msgp.AppendString(o, z.Msg)
	// string "kind"
	o = append(o, 0xa4, 0x6b, 0x69, 0x6e, 0x64)
	o = msgp.AppendInt32(o, z.Kind)
	// string "room"
	o = append(o, 0xa4, 0x72, 0x6f, 0x6f, 0x6d)
	o = msgp.AppendInt32(o, z.Room)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SendRoomFail) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zeff uint32
	zeff, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zeff > 0 {
		zeff--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "code":
			z.Code, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "msg":
			z.Msg, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "kind":
			z.Kind, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "room":
			z.Room, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SendRoomFail) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int32Size + 4 + msgp.StringPrefixSize + len(z.Msg) + 5 + msgp.Int32Size + 5 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *User) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrsw uint32
	zrsw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrsw > 0 {
		zrsw--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "icon":
			z.Icon, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "vip":
			z.Vip, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "coin":
			z.Coin, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *User) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "id"
	err = en.Append(0x85, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
	}
	// write "icon"
	err = en.Append(0xa4, 0x69, 0x63, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Icon)
	if err != nil {
		return
	}
	// write "vip"
	err = en.Append(0xa3, 0x76, 0x69, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Vip)
	if err != nil {
		return
	}
	// write "coin"
	err = en.Append(0xa4, 0x63, 0x6f, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Coin)
	if err != nil {
		return
	}
	// write "name"
	err = en.Append(0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *User) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "id"
	o = append(o, 0x85, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "icon"
	o = append(o, 0xa4, 0x69, 0x63, 0x6f, 0x6e)
	o = msgp.AppendInt32(o, z.Icon)
	// string "vip"
	o = append(o, 0xa3, 0x76, 0x69, 0x70)
	o = msgp.AppendInt32(o, z.Vip)
	// string "coin"
	o = append(o, 0xa4, 0x63, 0x6f, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.Coin)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *User) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zxpk uint32
	zxpk, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxpk > 0 {
		zxpk--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "icon":
			z.Icon, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "vip":
			z.Vip, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "coin":
			z.Coin, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *User) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 5 + msgp.Int32Size + 4 + msgp.Int32Size + 5 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name)
	return
}
