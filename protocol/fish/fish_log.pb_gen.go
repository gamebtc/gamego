package fish

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *GameRound) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "i":
			z.Id, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "u":
			z.Uid, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "Uid")
				return
			}
		case "s":
			z.Start, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Start")
				return
			}
		case "e":
			z.End, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "End")
				return
			}
		case "r":
			z.Room, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "Room")
				return
			}
		case "t":
			z.Tab, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "Tab")
				return
			}
		case "c":
			z.OldCoin, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "OldCoin")
				return
			}
		case "m":
			z.Bet, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Bet")
				return
			}
		case "w":
			z.Win, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Win")
				return
			}
		case "l":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Log")
				return
			}
			if cap(z.Log) >= int(zb0002) {
				z.Log = (z.Log)[:zb0002]
			} else {
				z.Log = make([]int32, zb0002)
			}
			for za0001 := range z.Log {
				z.Log[za0001], err = dc.ReadInt32()
				if err != nil {
					err = msgp.WrapError(err, "Log", za0001)
					return
				}
			}
		case "n":
			z.Note, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Note")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *GameRound) EncodeMsg(en *msgp.Writer) (err error) {
	// omitempty: check for empty values
	zb0001Len := uint32(11)
	var zb0001Mask uint16 /* 11 bits */
	if z.Note == "" {
		zb0001Len--
		zb0001Mask |= 0x400
	}
	// variable map header, size zb0001Len
	err = en.Append(0x80 | uint8(zb0001Len))
	if err != nil {
		return
	}
	if zb0001Len == 0 {
		return
	}
	// write "i"
	err = en.Append(0xa1, 0x69)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Id)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	// write "u"
	err = en.Append(0xa1, 0x75)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Uid)
	if err != nil {
		err = msgp.WrapError(err, "Uid")
		return
	}
	// write "s"
	err = en.Append(0xa1, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Start)
	if err != nil {
		err = msgp.WrapError(err, "Start")
		return
	}
	// write "e"
	err = en.Append(0xa1, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.End)
	if err != nil {
		err = msgp.WrapError(err, "End")
		return
	}
	// write "r"
	err = en.Append(0xa1, 0x72)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Room)
	if err != nil {
		err = msgp.WrapError(err, "Room")
		return
	}
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Tab)
	if err != nil {
		err = msgp.WrapError(err, "Tab")
		return
	}
	// write "c"
	err = en.Append(0xa1, 0x63)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.OldCoin)
	if err != nil {
		err = msgp.WrapError(err, "OldCoin")
		return
	}
	// write "m"
	err = en.Append(0xa1, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Bet)
	if err != nil {
		err = msgp.WrapError(err, "Bet")
		return
	}
	// write "w"
	err = en.Append(0xa1, 0x77)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Win)
	if err != nil {
		err = msgp.WrapError(err, "Win")
		return
	}
	// write "l"
	err = en.Append(0xa1, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Log)))
	if err != nil {
		err = msgp.WrapError(err, "Log")
		return
	}
	for za0001 := range z.Log {
		err = en.WriteInt32(z.Log[za0001])
		if err != nil {
			err = msgp.WrapError(err, "Log", za0001)
			return
		}
	}
	if (zb0001Mask & 0x400) == 0 { // if not empty
		// write "n"
		err = en.Append(0xa1, 0x6e)
		if err != nil {
			return
		}
		err = en.WriteString(z.Note)
		if err != nil {
			err = msgp.WrapError(err, "Note")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameRound) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(11)
	var zb0001Mask uint16 /* 11 bits */
	if z.Note == "" {
		zb0001Len--
		zb0001Mask |= 0x400
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len == 0 {
		return
	}
	// string "i"
	o = append(o, 0xa1, 0x69)
	o = msgp.AppendInt64(o, z.Id)
	// string "u"
	o = append(o, 0xa1, 0x75)
	o = msgp.AppendInt32(o, z.Uid)
	// string "s"
	o = append(o, 0xa1, 0x73)
	o = msgp.AppendInt64(o, z.Start)
	// string "e"
	o = append(o, 0xa1, 0x65)
	o = msgp.AppendInt64(o, z.End)
	// string "r"
	o = append(o, 0xa1, 0x72)
	o = msgp.AppendInt32(o, z.Room)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendInt32(o, z.Tab)
	// string "c"
	o = append(o, 0xa1, 0x63)
	o = msgp.AppendInt64(o, z.OldCoin)
	// string "m"
	o = append(o, 0xa1, 0x6d)
	o = msgp.AppendInt64(o, z.Bet)
	// string "w"
	o = append(o, 0xa1, 0x77)
	o = msgp.AppendInt64(o, z.Win)
	// string "l"
	o = append(o, 0xa1, 0x6c)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Log)))
	for za0001 := range z.Log {
		o = msgp.AppendInt32(o, z.Log[za0001])
	}
	if (zb0001Mask & 0x400) == 0 { // if not empty
		// string "n"
		o = append(o, 0xa1, 0x6e)
		o = msgp.AppendString(o, z.Note)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameRound) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "i":
			z.Id, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "u":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Uid")
				return
			}
		case "s":
			z.Start, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Start")
				return
			}
		case "e":
			z.End, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "End")
				return
			}
		case "r":
			z.Room, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Room")
				return
			}
		case "t":
			z.Tab, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Tab")
				return
			}
		case "c":
			z.OldCoin, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "OldCoin")
				return
			}
		case "m":
			z.Bet, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bet")
				return
			}
		case "w":
			z.Win, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Win")
				return
			}
		case "l":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Log")
				return
			}
			if cap(z.Log) >= int(zb0002) {
				z.Log = (z.Log)[:zb0002]
			} else {
				z.Log = make([]int32, zb0002)
			}
			for za0001 := range z.Log {
				z.Log[za0001], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Log", za0001)
					return
				}
			}
		case "n":
			z.Note, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Note")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *GameRound) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.ArrayHeaderSize + (len(z.Log) * (msgp.Int32Size)) + 2 + msgp.StringPrefixSize + len(z.Note)
	return
}
