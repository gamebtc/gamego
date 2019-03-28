package zjh

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *ActionAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "u":
			z.Uid, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "t":
			{
				var zbzg int32
				zbzg, err = dc.ReadInt32()
				z.Type = ActionType(zbzg)
			}
			if err != nil {
				return
			}
		case "b":
			z.Bet, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "o":
			z.Opponent, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "p":
			z.Poker, err = dc.ReadBytes(z.Poker)
			if err != nil {
				return
			}
		case "w":
			z.Win, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "c":
			z.Coin, err = dc.ReadInt64()
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
func (z *ActionAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "u"
	err = en.Append(0x87, 0xa1, 0x75)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Uid)
	if err != nil {
		return
	}
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(int32(z.Type))
	if err != nil {
		return
	}
	// write "b"
	err = en.Append(0xa1, 0x62)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Bet)
	if err != nil {
		return
	}
	// write "o"
	err = en.Append(0xa1, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Opponent)
	if err != nil {
		return
	}
	// write "p"
	err = en.Append(0xa1, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Poker)
	if err != nil {
		return
	}
	// write "w"
	err = en.Append(0xa1, 0x77)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Win)
	if err != nil {
		return
	}
	// write "c"
	err = en.Append(0xa1, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Coin)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ActionAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "u"
	o = append(o, 0x87, 0xa1, 0x75)
	o = msgp.AppendInt32(o, z.Uid)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendInt32(o, int32(z.Type))
	// string "b"
	o = append(o, 0xa1, 0x62)
	o = msgp.AppendInt32(o, z.Bet)
	// string "o"
	o = append(o, 0xa1, 0x6f)
	o = msgp.AppendInt32(o, z.Opponent)
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendBytes(o, z.Poker)
	// string "w"
	o = append(o, 0xa1, 0x77)
	o = msgp.AppendBool(o, z.Win)
	// string "c"
	o = append(o, 0xa1, 0x63)
	o = msgp.AppendInt64(o, z.Coin)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ActionAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbai uint32
	zbai, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbai > 0 {
		zbai--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "u":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "t":
			{
				var zcmr int32
				zcmr, bts, err = msgp.ReadInt32Bytes(bts)
				z.Type = ActionType(zcmr)
			}
			if err != nil {
				return
			}
		case "b":
			z.Bet, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "o":
			z.Opponent, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "p":
			z.Poker, bts, err = msgp.ReadBytesBytes(bts, z.Poker)
			if err != nil {
				return
			}
		case "w":
			z.Win, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "c":
			z.Coin, bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *ActionAck) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.BytesPrefixSize + len(z.Poker) + 2 + msgp.BoolSize + 2 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ActionLog) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "s":
			z.Start, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "u":
			z.Uid, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "t":
			{
				var zwht int32
				zwht, err = dc.ReadInt32()
				z.Type = ActionType(zwht)
			}
			if err != nil {
				return
			}
		case "c":
			z.Bet, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "o":
			z.Opponent, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "w":
			z.Win, err = dc.ReadBool()
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
func (z *ActionLog) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "s"
	err = en.Append(0x86, 0xa1, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Start)
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
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(int32(z.Type))
	if err != nil {
		return
	}
	// write "c"
	err = en.Append(0xa1, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Bet)
	if err != nil {
		return
	}
	// write "o"
	err = en.Append(0xa1, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Opponent)
	if err != nil {
		return
	}
	// write "w"
	err = en.Append(0xa1, 0x77)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Win)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ActionLog) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "s"
	o = append(o, 0x86, 0xa1, 0x73)
	o = msgp.AppendInt64(o, z.Start)
	// string "u"
	o = append(o, 0xa1, 0x75)
	o = msgp.AppendInt32(o, z.Uid)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendInt32(o, int32(z.Type))
	// string "c"
	o = append(o, 0xa1, 0x63)
	o = msgp.AppendInt32(o, z.Bet)
	// string "o"
	o = append(o, 0xa1, 0x6f)
	o = msgp.AppendInt32(o, z.Opponent)
	// string "w"
	o = append(o, 0xa1, 0x77)
	o = msgp.AppendBool(o, z.Win)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ActionLog) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zhct uint32
	zhct, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zhct > 0 {
		zhct--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "s":
			z.Start, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "u":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "t":
			{
				var zcua int32
				zcua, bts, err = msgp.ReadInt32Bytes(bts)
				z.Type = ActionType(zcua)
			}
			if err != nil {
				return
			}
		case "c":
			z.Bet, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "o":
			z.Opponent, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "w":
			z.Win, bts, err = msgp.ReadBoolBytes(bts)
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
func (z *ActionLog) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ActionReq) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxhx uint32
	zxhx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "t":
			{
				var zlqf int32
				zlqf, err = dc.ReadInt32()
				z.Type = ActionType(zlqf)
			}
			if err != nil {
				return
			}
		case "b":
			z.Bet, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "o":
			z.Opponent, err = dc.ReadInt32()
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
func (z ActionReq) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "t"
	err = en.Append(0x83, 0xa1, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(int32(z.Type))
	if err != nil {
		return
	}
	// write "b"
	err = en.Append(0xa1, 0x62)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Bet)
	if err != nil {
		return
	}
	// write "o"
	err = en.Append(0xa1, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Opponent)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ActionReq) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "t"
	o = append(o, 0x83, 0xa1, 0x74)
	o = msgp.AppendInt32(o, int32(z.Type))
	// string "b"
	o = append(o, 0xa1, 0x62)
	o = msgp.AppendInt32(o, z.Bet)
	// string "o"
	o = append(o, 0xa1, 0x6f)
	o = msgp.AppendInt32(o, z.Opponent)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ActionReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zdaf uint32
	zdaf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zdaf > 0 {
		zdaf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "t":
			{
				var zpks int32
				zpks, bts, err = msgp.ReadInt32Bytes(bts)
				z.Type = ActionType(zpks)
			}
			if err != nil {
				return
			}
		case "b":
			z.Bet, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "o":
			z.Opponent, bts, err = msgp.ReadInt32Bytes(bts)
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
func (z ActionReq) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ActionType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zjfb int32
		zjfb, err = dc.ReadInt32()
		(*z) = ActionType(zjfb)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z ActionType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt32(int32(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ActionType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt32(o, int32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ActionType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zcxo int32
		zcxo, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = ActionType(zcxo)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ActionType) Msgsize() (s int) {
	s = msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Code) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zeff int32
		zeff, err = dc.ReadInt32()
		(*z) = Code(zeff)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Code) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt32(int32(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Code) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt32(o, int32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Code) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zrsw int32
		zrsw, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = Code(zrsw)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Code) Msgsize() (s int) {
	s = msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameBill) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zdnj uint32
	zdnj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zdnj > 0 {
		zdnj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "u":
			z.Uid, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "j":
			z.Job, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "c":
			z.Coin, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "b":
			z.Bet, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "p":
			z.Poker, err = dc.ReadBytes(z.Poker)
			if err != nil {
				return
			}
		case "w":
			z.Win, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "x":
			z.Tax, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "o":
			z.Water, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "y":
			z.Lucky, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "t":
			z.Robot, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "k":
			var zobc uint32
			zobc, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Pk) >= int(zobc) {
				z.Pk = (z.Pk)[:zobc]
			} else {
				z.Pk = make([]int32, zobc)
			}
			for zxpk := range z.Pk {
				z.Pk[zxpk], err = dc.ReadInt32()
				if err != nil {
					return
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
func (z *GameBill) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 11
	// write "u"
	err = en.Append(0x8b, 0xa1, 0x75)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Uid)
	if err != nil {
		return
	}
	// write "j"
	err = en.Append(0xa1, 0x6a)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Job)
	if err != nil {
		return
	}
	// write "c"
	err = en.Append(0xa1, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Coin)
	if err != nil {
		return
	}
	// write "b"
	err = en.Append(0xa1, 0x62)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Bet)
	if err != nil {
		return
	}
	// write "p"
	err = en.Append(0xa1, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Poker)
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
	// write "x"
	err = en.Append(0xa1, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Tax)
	if err != nil {
		return
	}
	// write "o"
	err = en.Append(0xa1, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Water)
	if err != nil {
		return
	}
	// write "y"
	err = en.Append(0xa1, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Lucky)
	if err != nil {
		return
	}
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Robot)
	if err != nil {
		return
	}
	// write "k"
	err = en.Append(0xa1, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Pk)))
	if err != nil {
		return
	}
	for zxpk := range z.Pk {
		err = en.WriteInt32(z.Pk[zxpk])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameBill) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 11
	// string "u"
	o = append(o, 0x8b, 0xa1, 0x75)
	o = msgp.AppendInt32(o, z.Uid)
	// string "j"
	o = append(o, 0xa1, 0x6a)
	o = msgp.AppendInt32(o, z.Job)
	// string "c"
	o = append(o, 0xa1, 0x63)
	o = msgp.AppendInt64(o, z.Coin)
	// string "b"
	o = append(o, 0xa1, 0x62)
	o = msgp.AppendInt64(o, z.Bet)
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendBytes(o, z.Poker)
	// string "w"
	o = append(o, 0xa1, 0x77)
	o = msgp.AppendInt64(o, z.Win)
	// string "x"
	o = append(o, 0xa1, 0x78)
	o = msgp.AppendInt64(o, z.Tax)
	// string "o"
	o = append(o, 0xa1, 0x6f)
	o = msgp.AppendInt64(o, z.Water)
	// string "y"
	o = append(o, 0xa1, 0x79)
	o = msgp.AppendInt64(o, z.Lucky)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendInt64(o, z.Robot)
	// string "k"
	o = append(o, 0xa1, 0x6b)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Pk)))
	for zxpk := range z.Pk {
		o = msgp.AppendInt32(o, z.Pk[zxpk])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameBill) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsnv uint32
	zsnv, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsnv > 0 {
		zsnv--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "u":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "j":
			z.Job, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "c":
			z.Coin, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "b":
			z.Bet, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "p":
			z.Poker, bts, err = msgp.ReadBytesBytes(bts, z.Poker)
			if err != nil {
				return
			}
		case "w":
			z.Win, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "x":
			z.Tax, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "o":
			z.Water, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "y":
			z.Lucky, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "t":
			z.Robot, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "k":
			var zkgt uint32
			zkgt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Pk) >= int(zkgt) {
				z.Pk = (z.Pk)[:zkgt]
			} else {
				z.Pk = make([]int32, zkgt)
			}
			for zxpk := range z.Pk {
				z.Pk[zxpk], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
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
func (z *GameBill) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.BytesPrefixSize + len(z.Poker) + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.ArrayHeaderSize + (len(z.Pk) * (msgp.Int32Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameEndAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zema uint32
	zema, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zema > 0 {
		zema--
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
		case "win":
			z.Win, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "winner":
			z.Winner, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "coin":
			z.Coin, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "poker":
			z.Poker, err = dc.ReadBytes(z.Poker)
			if err != nil {
				return
			}
		case "lucky":
			z.Lucky, err = dc.ReadInt64()
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
func (z *GameEndAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "id"
	err = en.Append(0x86, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
	}
	// write "win"
	err = en.Append(0xa3, 0x77, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Win)
	if err != nil {
		return
	}
	// write "winner"
	err = en.Append(0xa6, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Winner)
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
	// write "poker"
	err = en.Append(0xa5, 0x70, 0x6f, 0x6b, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Poker)
	if err != nil {
		return
	}
	// write "lucky"
	err = en.Append(0xa5, 0x6c, 0x75, 0x63, 0x6b, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Lucky)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameEndAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "id"
	o = append(o, 0x86, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "win"
	o = append(o, 0xa3, 0x77, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.Win)
	// string "winner"
	o = append(o, 0xa6, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72)
	o = msgp.AppendInt32(o, z.Winner)
	// string "coin"
	o = append(o, 0xa4, 0x63, 0x6f, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.Coin)
	// string "poker"
	o = append(o, 0xa5, 0x70, 0x6f, 0x6b, 0x65, 0x72)
	o = msgp.AppendBytes(o, z.Poker)
	// string "lucky"
	o = append(o, 0xa5, 0x6c, 0x75, 0x63, 0x6b, 0x79)
	o = msgp.AppendInt64(o, z.Lucky)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameEndAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zpez uint32
	zpez, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zpez > 0 {
		zpez--
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
		case "win":
			z.Win, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "winner":
			z.Winner, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "coin":
			z.Coin, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "poker":
			z.Poker, bts, err = msgp.ReadBytesBytes(bts, z.Poker)
			if err != nil {
				return
			}
		case "lucky":
			z.Lucky, bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *GameEndAck) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 4 + msgp.Int64Size + 7 + msgp.Int32Size + 5 + msgp.Int64Size + 6 + msgp.BytesPrefixSize + len(z.Poker) + 6 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameInitAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zqyh uint32
	zqyh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zqyh > 0 {
		zqyh--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "table":
			z.Table, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "id":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "pool":
			z.Pool, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "state":
			z.State, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "ring":
			z.Ring, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "first":
			z.First, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "play":
			var zyzr uint32
			zyzr, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Player) >= int(zyzr) {
				z.Player = (z.Player)[:zyzr]
			} else {
				z.Player = make([]*Player, zyzr)
			}
			for zqke := range z.Player {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Player[zqke] = nil
				} else {
					if z.Player[zqke] == nil {
						z.Player[zqke] = new(Player)
					}
					err = z.Player[zqke].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "p":
			z.Poker, err = dc.ReadBytes(z.Poker)
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
func (z *GameInitAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 8
	// write "table"
	err = en.Append(0x88, 0xa5, 0x74, 0x61, 0x62, 0x6c, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Table)
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
	// write "pool"
	err = en.Append(0xa4, 0x70, 0x6f, 0x6f, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Pool)
	if err != nil {
		return
	}
	// write "state"
	err = en.Append(0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.State)
	if err != nil {
		return
	}
	// write "ring"
	err = en.Append(0xa4, 0x72, 0x69, 0x6e, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Ring)
	if err != nil {
		return
	}
	// write "first"
	err = en.Append(0xa5, 0x66, 0x69, 0x72, 0x73, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.First)
	if err != nil {
		return
	}
	// write "play"
	err = en.Append(0xa4, 0x70, 0x6c, 0x61, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Player)))
	if err != nil {
		return
	}
	for zqke := range z.Player {
		if z.Player[zqke] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Player[zqke].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "p"
	err = en.Append(0xa1, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Poker)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameInitAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "table"
	o = append(o, 0x88, 0xa5, 0x74, 0x61, 0x62, 0x6c, 0x65)
	o = msgp.AppendInt32(o, z.Table)
	// string "id"
	o = append(o, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "pool"
	o = append(o, 0xa4, 0x70, 0x6f, 0x6f, 0x6c)
	o = msgp.AppendInt64(o, z.Pool)
	// string "state"
	o = append(o, 0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	o = msgp.AppendInt32(o, z.State)
	// string "ring"
	o = append(o, 0xa4, 0x72, 0x69, 0x6e, 0x67)
	o = msgp.AppendInt32(o, z.Ring)
	// string "first"
	o = append(o, 0xa5, 0x66, 0x69, 0x72, 0x73, 0x74)
	o = msgp.AppendInt32(o, z.First)
	// string "play"
	o = append(o, 0xa4, 0x70, 0x6c, 0x61, 0x79)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Player)))
	for zqke := range z.Player {
		if z.Player[zqke] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Player[zqke].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendBytes(o, z.Poker)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameInitAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zywj uint32
	zywj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zywj > 0 {
		zywj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "table":
			z.Table, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "pool":
			z.Pool, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "state":
			z.State, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "ring":
			z.Ring, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "first":
			z.First, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "play":
			var zjpj uint32
			zjpj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Player) >= int(zjpj) {
				z.Player = (z.Player)[:zjpj]
			} else {
				z.Player = make([]*Player, zjpj)
			}
			for zqke := range z.Player {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Player[zqke] = nil
				} else {
					if z.Player[zqke] == nil {
						z.Player[zqke] = new(Player)
					}
					bts, err = z.Player[zqke].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "p":
			z.Poker, bts, err = msgp.ReadBytesBytes(bts, z.Poker)
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
func (z *GameInitAck) Msgsize() (s int) {
	s = 1 + 6 + msgp.Int32Size + 3 + msgp.Int32Size + 5 + msgp.Int64Size + 6 + msgp.Int32Size + 5 + msgp.Int32Size + 6 + msgp.Int32Size + 5 + msgp.ArrayHeaderSize
	for zqke := range z.Player {
		if z.Player[zqke] == nil {
			s += msgp.NilSize
		} else {
			s += z.Player[zqke].Msgsize()
		}
	}
	s += 2 + msgp.BytesPrefixSize + len(z.Poker)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameRound) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zgmo uint32
	zgmo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zgmo > 0 {
		zgmo--
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
		case "b":
			var ztaf uint32
			ztaf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(ztaf) {
				z.Bill = (z.Bill)[:ztaf]
			} else {
				z.Bill = make([]*GameBill, ztaf)
			}
			for zzpf := range z.Bill {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Bill[zzpf] = nil
				} else {
					if z.Bill[zzpf] == nil {
						z.Bill[zzpf] = new(GameBill)
					}
					err = z.Bill[zzpf].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "a":
			z.Ante, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "g":
			z.Ring, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "m":
			z.Sum, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "w":
			z.Win, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "x":
			z.Tax, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "o":
			z.Water, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "l":
			var zeth uint32
			zeth, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Log) >= int(zeth) {
				z.Log = (z.Log)[:zeth]
			} else {
				z.Log = make([]*ActionLog, zeth)
			}
			for zrfe := range z.Log {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Log[zrfe] = nil
				} else {
					if z.Log[zrfe] == nil {
						z.Log[zrfe] = new(ActionLog)
					}
					err = z.Log[zrfe].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "p":
			z.Pool, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "y":
			z.Lucky, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "n":
			z.Note, err = dc.ReadString()
			if err != nil {
				return
			}
		case "v":
			z.Cheat, err = dc.ReadBool()
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
	// map header, size 17
	// write "i"
	err = en.Append(0xde, 0x0, 0x11, 0xa1, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Id)
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
	// write "b"
	err = en.Append(0xa1, 0x62)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Bill)))
	if err != nil {
		return
	}
	for zzpf := range z.Bill {
		if z.Bill[zzpf] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Bill[zzpf].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "a"
	err = en.Append(0xa1, 0x61)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Ante)
	if err != nil {
		return
	}
	// write "g"
	err = en.Append(0xa1, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Ring)
	if err != nil {
		return
	}
	// write "m"
	err = en.Append(0xa1, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Sum)
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
	// write "x"
	err = en.Append(0xa1, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Tax)
	if err != nil {
		return
	}
	// write "o"
	err = en.Append(0xa1, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Water)
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
	for zrfe := range z.Log {
		if z.Log[zrfe] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Log[zrfe].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "p"
	err = en.Append(0xa1, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Pool)
	if err != nil {
		return
	}
	// write "y"
	err = en.Append(0xa1, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Lucky)
	if err != nil {
		return
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
	// write "v"
	err = en.Append(0xa1, 0x76)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Cheat)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameRound) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 17
	// string "i"
	o = append(o, 0xde, 0x0, 0x11, 0xa1, 0x69)
	o = msgp.AppendInt64(o, z.Id)
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
	// string "b"
	o = append(o, 0xa1, 0x62)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bill)))
	for zzpf := range z.Bill {
		if z.Bill[zzpf] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Bill[zzpf].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "a"
	o = append(o, 0xa1, 0x61)
	o = msgp.AppendInt64(o, z.Ante)
	// string "g"
	o = append(o, 0xa1, 0x67)
	o = msgp.AppendInt32(o, z.Ring)
	// string "m"
	o = append(o, 0xa1, 0x6d)
	o = msgp.AppendInt64(o, z.Sum)
	// string "w"
	o = append(o, 0xa1, 0x77)
	o = msgp.AppendInt64(o, z.Win)
	// string "x"
	o = append(o, 0xa1, 0x78)
	o = msgp.AppendInt64(o, z.Tax)
	// string "o"
	o = append(o, 0xa1, 0x6f)
	o = msgp.AppendInt64(o, z.Water)
	// string "l"
	o = append(o, 0xa1, 0x6c)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Log)))
	for zrfe := range z.Log {
		if z.Log[zrfe] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Log[zrfe].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendInt64(o, z.Pool)
	// string "y"
	o = append(o, 0xa1, 0x79)
	o = msgp.AppendInt64(o, z.Lucky)
	// string "n"
	o = append(o, 0xa1, 0x6e)
	o = msgp.AppendString(o, z.Note)
	// string "v"
	o = append(o, 0xa1, 0x76)
	o = msgp.AppendBool(o, z.Cheat)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameRound) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsbz uint32
	zsbz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsbz > 0 {
		zsbz--
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
		case "b":
			var zrjx uint32
			zrjx, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(zrjx) {
				z.Bill = (z.Bill)[:zrjx]
			} else {
				z.Bill = make([]*GameBill, zrjx)
			}
			for zzpf := range z.Bill {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Bill[zzpf] = nil
				} else {
					if z.Bill[zzpf] == nil {
						z.Bill[zzpf] = new(GameBill)
					}
					bts, err = z.Bill[zzpf].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "a":
			z.Ante, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "g":
			z.Ring, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "m":
			z.Sum, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "w":
			z.Win, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "x":
			z.Tax, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "o":
			z.Water, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "l":
			var zawn uint32
			zawn, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Log) >= int(zawn) {
				z.Log = (z.Log)[:zawn]
			} else {
				z.Log = make([]*ActionLog, zawn)
			}
			for zrfe := range z.Log {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Log[zrfe] = nil
				} else {
					if z.Log[zrfe] == nil {
						z.Log[zrfe] = new(ActionLog)
					}
					bts, err = z.Log[zrfe].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "p":
			z.Pool, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "y":
			z.Lucky, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "n":
			z.Note, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "v":
			z.Cheat, bts, err = msgp.ReadBoolBytes(bts)
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
	s = 3 + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.ArrayHeaderSize
	for zzpf := range z.Bill {
		if z.Bill[zzpf] == nil {
			s += msgp.NilSize
		} else {
			s += z.Bill[zzpf].Msgsize()
		}
	}
	s += 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.ArrayHeaderSize
	for zrfe := range z.Log {
		if z.Log[zrfe] == nil {
			s += msgp.NilSize
		} else {
			s += z.Log[zrfe].Msgsize()
		}
	}
	s += 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.StringPrefixSize + len(z.Note) + 2 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameStartAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrbe uint32
	zrbe, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrbe > 0 {
		zrbe--
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
		case "pool":
			z.Pool, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "play":
			var zmfd uint32
			zmfd, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Player) >= int(zmfd) {
				z.Player = (z.Player)[:zmfd]
			} else {
				z.Player = make([]*Player, zmfd)
			}
			for zwel := range z.Player {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Player[zwel] = nil
				} else {
					if z.Player[zwel] == nil {
						z.Player[zwel] = new(Player)
					}
					err = z.Player[zwel].DecodeMsg(dc)
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
func (z *GameStartAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "id"
	err = en.Append(0x83, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
	}
	// write "pool"
	err = en.Append(0xa4, 0x70, 0x6f, 0x6f, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Pool)
	if err != nil {
		return
	}
	// write "play"
	err = en.Append(0xa4, 0x70, 0x6c, 0x61, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Player)))
	if err != nil {
		return
	}
	for zwel := range z.Player {
		if z.Player[zwel] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Player[zwel].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameStartAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "id"
	o = append(o, 0x83, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "pool"
	o = append(o, 0xa4, 0x70, 0x6f, 0x6f, 0x6c)
	o = msgp.AppendInt64(o, z.Pool)
	// string "play"
	o = append(o, 0xa4, 0x70, 0x6c, 0x61, 0x79)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Player)))
	for zwel := range z.Player {
		if z.Player[zwel] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Player[zwel].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameStartAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zzdc uint32
	zzdc, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zzdc > 0 {
		zzdc--
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
		case "pool":
			z.Pool, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "play":
			var zelx uint32
			zelx, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Player) >= int(zelx) {
				z.Player = (z.Player)[:zelx]
			} else {
				z.Player = make([]*Player, zelx)
			}
			for zwel := range z.Player {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Player[zwel] = nil
				} else {
					if z.Player[zwel] == nil {
						z.Player[zwel] = new(Player)
					}
					bts, err = z.Player[zwel].UnmarshalMsg(bts)
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
func (z *GameStartAck) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 5 + msgp.Int64Size + 5 + msgp.ArrayHeaderSize
	for zwel := range z.Player {
		if z.Player[zwel] == nil {
			s += msgp.NilSize
		} else {
			s += z.Player[zwel].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Player) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbal uint32
	zbal, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbal > 0 {
		zbal--
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
		case "chair":
			z.Chair, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "coin":
			z.Coin, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "bet":
			z.Bet, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "state":
			{
				var zjqz int32
				zjqz, err = dc.ReadInt32()
				z.State = Player_State(zjqz)
			}
			if err != nil {
				return
			}
		case "look":
			z.Look, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "down":
			z.Down, err = dc.ReadInt32()
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
func (z *Player) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "id"
	err = en.Append(0x8a, 0xa2, 0x69, 0x64)
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
	// write "chair"
	err = en.Append(0xa5, 0x63, 0x68, 0x61, 0x69, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Chair)
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
	// write "bet"
	err = en.Append(0xa3, 0x62, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Bet)
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
	// write "state"
	err = en.Append(0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(int32(z.State))
	if err != nil {
		return
	}
	// write "look"
	err = en.Append(0xa4, 0x6c, 0x6f, 0x6f, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Look)
	if err != nil {
		return
	}
	// write "down"
	err = en.Append(0xa4, 0x64, 0x6f, 0x77, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Down)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Player) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "id"
	o = append(o, 0x8a, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "icon"
	o = append(o, 0xa4, 0x69, 0x63, 0x6f, 0x6e)
	o = msgp.AppendInt32(o, z.Icon)
	// string "vip"
	o = append(o, 0xa3, 0x76, 0x69, 0x70)
	o = msgp.AppendInt32(o, z.Vip)
	// string "chair"
	o = append(o, 0xa5, 0x63, 0x68, 0x61, 0x69, 0x72)
	o = msgp.AppendInt32(o, z.Chair)
	// string "coin"
	o = append(o, 0xa4, 0x63, 0x6f, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.Coin)
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendInt64(o, z.Bet)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "state"
	o = append(o, 0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	o = msgp.AppendInt32(o, int32(z.State))
	// string "look"
	o = append(o, 0xa4, 0x6c, 0x6f, 0x6f, 0x6b)
	o = msgp.AppendBool(o, z.Look)
	// string "down"
	o = append(o, 0xa4, 0x64, 0x6f, 0x77, 0x6e)
	o = msgp.AppendInt32(o, z.Down)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Player) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zkct uint32
	zkct, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zkct > 0 {
		zkct--
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
		case "chair":
			z.Chair, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "coin":
			z.Coin, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "bet":
			z.Bet, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "state":
			{
				var ztmt int32
				ztmt, bts, err = msgp.ReadInt32Bytes(bts)
				z.State = Player_State(ztmt)
			}
			if err != nil {
				return
			}
		case "look":
			z.Look, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "down":
			z.Down, bts, err = msgp.ReadInt32Bytes(bts)
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
func (z *Player) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 5 + msgp.Int32Size + 4 + msgp.Int32Size + 6 + msgp.Int32Size + 5 + msgp.Int64Size + 4 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 6 + msgp.Int32Size + 5 + msgp.BoolSize + 5 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Player_State) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var ztco int32
		ztco, err = dc.ReadInt32()
		(*z) = Player_State(ztco)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Player_State) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt32(int32(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Player_State) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt32(o, int32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Player_State) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zana int32
		zana, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = Player_State(zana)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Player_State) Msgsize() (s int) {
	s = msgp.Int32Size
	return
}
