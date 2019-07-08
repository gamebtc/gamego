package fish

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *GameRound) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "i":
			z.Id, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "u":
			z.Uid, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "s":
			z.Start, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "e":
			z.End, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "r":
			z.Room, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "t":
			z.Tab, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "c":
			z.OldCoin, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "m":
			z.Bet, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "w":
			z.Win, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "l":
			var zbai uint32
			zbai, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Log) >= int(zbai) {
				z.Log = (z.Log)[:zbai]
			} else {
				z.Log = make([]int32, zbai)
			}
			for zxvk := range z.Log {
				z.Log[zxvk], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "n":
			z.Note, err = dc.ReadString()
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
func (z *GameRound) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 11
	// write "i"
	err = en.Append(0x8b, 0xa1, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Id)
	if err != nil {
		return
	}
	// write "u"
	err = en.Append(0xa1, 0x75)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Uid)
	if err != nil {
		return
	}
	// write "s"
	err = en.Append(0xa1, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Start)
	if err != nil {
		return
	}
	// write "e"
	err = en.Append(0xa1, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.End)
	if err != nil {
		return
	}
	// write "r"
	err = en.Append(0xa1, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Room)
	if err != nil {
		return
	}
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Tab)
	if err != nil {
		return
	}
	// write "c"
	err = en.Append(0xa1, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.OldCoin)
	if err != nil {
		return
	}
	// write "m"
	err = en.Append(0xa1, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Bet)
	if err != nil {
		return
	}
	// write "w"
	err = en.Append(0xa1, 0x77)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Win)
	if err != nil {
		return
	}
	// write "l"
	err = en.Append(0xa1, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Log)))
	if err != nil {
		return
	}
	for zxvk := range z.Log {
		err = en.WriteInt32(z.Log[zxvk])
		if err != nil {
			return
		}
	}
	// write "n"
	err = en.Append(0xa1, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Note)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameRound) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 11
	// string "i"
	o = append(o, 0x8b, 0xa1, 0x69)
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
	for zxvk := range z.Log {
		o = msgp.AppendInt32(o, z.Log[zxvk])
	}
	// string "n"
	o = append(o, 0xa1, 0x6e)
	o = msgp.AppendString(o, z.Note)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameRound) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "i":
			z.Id, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "u":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "s":
			z.Start, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "e":
			z.End, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "r":
			z.Room, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "t":
			z.Tab, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "c":
			z.OldCoin, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "m":
			z.Bet, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "w":
			z.Win, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "l":
			var zajw uint32
			zajw, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Log) >= int(zajw) {
				z.Log = (z.Log)[:zajw]
			} else {
				z.Log = make([]int32, zajw)
			}
			for zxvk := range z.Log {
				z.Log[zxvk], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "n":
			z.Note, bts, err = msgp.ReadStringBytes(bts)
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
func (z *GameRound) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.ArrayHeaderSize + (len(z.Log) * (msgp.Int32Size)) + 2 + msgp.StringPrefixSize + len(z.Note)
	return
}
