package folks

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BetAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "sn":
			z.Sn, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "i":
			z.Item, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "bet":
			z.Bet, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "coin":
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
func (z *BetAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "sn"
	err = en.Append(0x84, 0xa2, 0x73, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Sn)
	if err != nil {
		return
	}
	// write "i"
	err = en.Append(0xa1, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Item)
	if err != nil {
		return
	}
	// write "bet"
	err = en.Append(0xa3, 0x62, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Bet)
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
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BetAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "sn"
	o = append(o, 0x84, 0xa2, 0x73, 0x6e)
	o = msgp.AppendInt32(o, z.Sn)
	// string "i"
	o = append(o, 0xa1, 0x69)
	o = msgp.AppendInt32(o, z.Item)
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendInt32(o, z.Bet)
	// string "coin"
	o = append(o, 0xa4, 0x63, 0x6f, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.Coin)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BetAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "sn":
			z.Sn, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "i":
			z.Item, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "bet":
			z.Bet, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "coin":
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
func (z *BetAck) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 2 + msgp.Int32Size + 4 + msgp.Int32Size + 5 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BetReq) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "sn":
			z.Sn, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "i":
			z.Item, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "bet":
			z.Bet, err = dc.ReadInt32()
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
func (z BetReq) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "sn"
	err = en.Append(0x83, 0xa2, 0x73, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Sn)
	if err != nil {
		return
	}
	// write "i"
	err = en.Append(0xa1, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Item)
	if err != nil {
		return
	}
	// write "bet"
	err = en.Append(0xa3, 0x62, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Bet)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z BetReq) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "sn"
	o = append(o, 0x83, 0xa2, 0x73, 0x6e)
	o = msgp.AppendInt32(o, z.Sn)
	// string "i"
	o = append(o, 0xa1, 0x69)
	o = msgp.AppendInt32(o, z.Item)
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendInt32(o, z.Bet)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BetReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "sn":
			z.Sn, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "i":
			z.Item, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "bet":
			z.Bet, bts, err = msgp.ReadInt32Bytes(bts)
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
func (z BetReq) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 2 + msgp.Int32Size + 4 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *CloseBetAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "r":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.R = nil
			} else {
				if z.R == nil {
					z.R = new(GameResult)
				}
				err = z.R.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "win":
			z.Win, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "coin":
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
func (z *CloseBetAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "r"
	err = en.Append(0x83, 0xa1, 0x72)
	if err != nil {
		return err
	}
	if z.R == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.R.EncodeMsg(en)
		if err != nil {
			return
		}
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
	// write "coin"
	err = en.Append(0xa4, 0x63, 0x6f, 0x69, 0x6e)
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
func (z *CloseBetAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "r"
	o = append(o, 0x83, 0xa1, 0x72)
	if z.R == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.R.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "win"
	o = append(o, 0xa3, 0x77, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.Win)
	// string "coin"
	o = append(o, 0xa4, 0x63, 0x6f, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.Coin)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CloseBetAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "r":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.R = nil
			} else {
				if z.R == nil {
					z.R = new(GameResult)
				}
				bts, err = z.R.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "win":
			z.Win, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "coin":
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
func (z *CloseBetAck) Msgsize() (s int) {
	s = 1 + 2
	if z.R == nil {
		s += msgp.NilSize
	} else {
		s += z.R.Msgsize()
	}
	s += 4 + msgp.Int64Size + 5 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameBill) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "u":
			z.Uid, err = dc.ReadInt32()
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
		case "g":
			var zxhx uint32
			zxhx, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zxhx) {
				z.Group = (z.Group)[:zxhx]
			} else {
				z.Group = make([]int64, zxhx)
			}
			for zhct := range z.Group {
				z.Group[zhct], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "w":
			z.Win, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "t":
			z.Tax, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "j":
			z.Job, err = dc.ReadInt32()
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
func (z *GameBill) EncodeMsg(en *msgp.Writer) (err error) {
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
	// write "g"
	err = en.Append(0xa1, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Group)))
	if err != nil {
		return
	}
	for zhct := range z.Group {
		err = en.WriteInt64(z.Group[zhct])
		if err != nil {
			return
		}
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
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Tax)
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
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameBill) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "u"
	o = append(o, 0x87, 0xa1, 0x75)
	o = msgp.AppendInt32(o, z.Uid)
	// string "c"
	o = append(o, 0xa1, 0x63)
	o = msgp.AppendInt64(o, z.Coin)
	// string "b"
	o = append(o, 0xa1, 0x62)
	o = msgp.AppendInt64(o, z.Bet)
	// string "g"
	o = append(o, 0xa1, 0x67)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Group)))
	for zhct := range z.Group {
		o = msgp.AppendInt64(o, z.Group[zhct])
	}
	// string "w"
	o = append(o, 0xa1, 0x77)
	o = msgp.AppendInt64(o, z.Win)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendInt64(o, z.Tax)
	// string "j"
	o = append(o, 0xa1, 0x6a)
	o = msgp.AppendInt32(o, z.Job)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameBill) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "u":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
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
		case "g":
			var zdaf uint32
			zdaf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zdaf) {
				z.Group = (z.Group)[:zdaf]
			} else {
				z.Group = make([]int64, zdaf)
			}
			for zhct := range z.Group {
				z.Group[zhct], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "w":
			z.Win, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "t":
			z.Tax, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "j":
			z.Job, bts, err = msgp.ReadInt32Bytes(bts)
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
func (z *GameBill) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int32Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.ArrayHeaderSize + (len(z.Group) * (msgp.Int64Size)) + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameInitAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zeff uint32
	zeff, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zeff > 0 {
		zeff--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "state":
			z.State, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "rich":
			var zrsw uint32
			zrsw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zrsw) {
				z.Rich = (z.Rich)[:zrsw]
			} else {
				z.Rich = make([]*Player, zrsw)
			}
			for zpks := range z.Rich {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Rich[zpks] = nil
				} else {
					if z.Rich[zpks] == nil {
						z.Rich[zpks] = new(Player)
					}
					err = z.Rich[zpks].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "sum":
			var zxpk uint32
			zxpk, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Sum) >= int(zxpk) {
				z.Sum = (z.Sum)[:zxpk]
			} else {
				z.Sum = make([]int64, zxpk)
			}
			for zjfb := range z.Sum {
				z.Sum[zjfb], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "bet":
			var zdnj uint32
			zdnj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zdnj) {
				z.Bet = (z.Bet)[:zdnj]
			} else {
				z.Bet = make([]int64, zdnj)
			}
			for zcxo := range z.Bet {
				z.Bet[zcxo], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "log":
			z.Log, err = dc.ReadBytes(z.Log)
			if err != nil {
				return
			}
		case "bank":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Bank = nil
			} else {
				if z.Bank == nil {
					z.Bank = new(Player)
				}
				err = z.Bank.DecodeMsg(dc)
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
func (z *GameInitAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "id"
	err = en.Append(0x87, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Id)
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
	// write "rich"
	err = en.Append(0xa4, 0x72, 0x69, 0x63, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Rich)))
	if err != nil {
		return
	}
	for zpks := range z.Rich {
		if z.Rich[zpks] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Rich[zpks].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "sum"
	err = en.Append(0xa3, 0x73, 0x75, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Sum)))
	if err != nil {
		return
	}
	for zjfb := range z.Sum {
		err = en.WriteInt64(z.Sum[zjfb])
		if err != nil {
			return
		}
	}
	// write "bet"
	err = en.Append(0xa3, 0x62, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Bet)))
	if err != nil {
		return
	}
	for zcxo := range z.Bet {
		err = en.WriteInt64(z.Bet[zcxo])
		if err != nil {
			return
		}
	}
	// write "log"
	err = en.Append(0xa3, 0x6c, 0x6f, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Log)
	if err != nil {
		return
	}
	// write "bank"
	err = en.Append(0xa4, 0x62, 0x61, 0x6e, 0x6b)
	if err != nil {
		return err
	}
	if z.Bank == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Bank.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameInitAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "id"
	o = append(o, 0x87, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Id)
	// string "state"
	o = append(o, 0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	o = msgp.AppendInt32(o, z.State)
	// string "rich"
	o = append(o, 0xa4, 0x72, 0x69, 0x63, 0x68)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rich)))
	for zpks := range z.Rich {
		if z.Rich[zpks] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Rich[zpks].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "sum"
	o = append(o, 0xa3, 0x73, 0x75, 0x6d)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Sum)))
	for zjfb := range z.Sum {
		o = msgp.AppendInt64(o, z.Sum[zjfb])
	}
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bet)))
	for zcxo := range z.Bet {
		o = msgp.AppendInt64(o, z.Bet[zcxo])
	}
	// string "log"
	o = append(o, 0xa3, 0x6c, 0x6f, 0x67)
	o = msgp.AppendBytes(o, z.Log)
	// string "bank"
	o = append(o, 0xa4, 0x62, 0x61, 0x6e, 0x6b)
	if z.Bank == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Bank.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameInitAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zobc uint32
	zobc, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zobc > 0 {
		zobc--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "state":
			z.State, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "rich":
			var zsnv uint32
			zsnv, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zsnv) {
				z.Rich = (z.Rich)[:zsnv]
			} else {
				z.Rich = make([]*Player, zsnv)
			}
			for zpks := range z.Rich {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Rich[zpks] = nil
				} else {
					if z.Rich[zpks] == nil {
						z.Rich[zpks] = new(Player)
					}
					bts, err = z.Rich[zpks].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "sum":
			var zkgt uint32
			zkgt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Sum) >= int(zkgt) {
				z.Sum = (z.Sum)[:zkgt]
			} else {
				z.Sum = make([]int64, zkgt)
			}
			for zjfb := range z.Sum {
				z.Sum[zjfb], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "bet":
			var zema uint32
			zema, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zema) {
				z.Bet = (z.Bet)[:zema]
			} else {
				z.Bet = make([]int64, zema)
			}
			for zcxo := range z.Bet {
				z.Bet[zcxo], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "log":
			z.Log, bts, err = msgp.ReadBytesBytes(bts, z.Log)
			if err != nil {
				return
			}
		case "bank":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Bank = nil
			} else {
				if z.Bank == nil {
					z.Bank = new(Player)
				}
				bts, err = z.Bank.UnmarshalMsg(bts)
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
func (z *GameInitAck) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 6 + msgp.Int32Size + 5 + msgp.ArrayHeaderSize
	for zpks := range z.Rich {
		if z.Rich[zpks] == nil {
			s += msgp.NilSize
		} else {
			s += z.Rich[zpks].Msgsize()
		}
	}
	s += 4 + msgp.ArrayHeaderSize + (len(z.Sum) * (msgp.Int64Size)) + 4 + msgp.ArrayHeaderSize + (len(z.Bet) * (msgp.Int64Size)) + 4 + msgp.BytesPrefixSize + len(z.Log) + 5
	if z.Bank == nil {
		s += msgp.NilSize
	} else {
		s += z.Bank.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameResult) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zyzr uint32
	zyzr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zyzr > 0 {
		zyzr--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "p":
			z.Poker, err = dc.ReadBytes(z.Poker)
			if err != nil {
				return
			}
		case "odd":
			var zywj uint32
			zywj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Odd) >= int(zywj) {
				z.Odd = (z.Odd)[:zywj]
			} else {
				z.Odd = make([]int32, zywj)
			}
			for zpez := range z.Odd {
				z.Odd[zpez], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "bet":
			var zjpj uint32
			zjpj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zjpj) {
				z.Bet = (z.Bet)[:zjpj]
			} else {
				z.Bet = make([]int64, zjpj)
			}
			for zqke := range z.Bet {
				z.Bet[zqke], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "rich":
			var zzpf uint32
			zzpf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zzpf) {
				z.Rich = (z.Rich)[:zzpf]
			} else {
				z.Rich = make([]int64, zzpf)
			}
			for zqyh := range z.Rich {
				z.Rich[zqyh], err = dc.ReadInt64()
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
func (z *GameResult) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "id"
	err = en.Append(0x85, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Id)
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
	// write "odd"
	err = en.Append(0xa3, 0x6f, 0x64, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Odd)))
	if err != nil {
		return
	}
	for zpez := range z.Odd {
		err = en.WriteInt32(z.Odd[zpez])
		if err != nil {
			return
		}
	}
	// write "bet"
	err = en.Append(0xa3, 0x62, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Bet)))
	if err != nil {
		return
	}
	for zqke := range z.Bet {
		err = en.WriteInt64(z.Bet[zqke])
		if err != nil {
			return
		}
	}
	// write "rich"
	err = en.Append(0xa4, 0x72, 0x69, 0x63, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Rich)))
	if err != nil {
		return
	}
	for zqyh := range z.Rich {
		err = en.WriteInt64(z.Rich[zqyh])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameResult) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "id"
	o = append(o, 0x85, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Id)
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendBytes(o, z.Poker)
	// string "odd"
	o = append(o, 0xa3, 0x6f, 0x64, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Odd)))
	for zpez := range z.Odd {
		o = msgp.AppendInt32(o, z.Odd[zpez])
	}
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bet)))
	for zqke := range z.Bet {
		o = msgp.AppendInt64(o, z.Bet[zqke])
	}
	// string "rich"
	o = append(o, 0xa4, 0x72, 0x69, 0x63, 0x68)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rich)))
	for zqyh := range z.Rich {
		o = msgp.AppendInt64(o, z.Rich[zqyh])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameResult) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrfe uint32
	zrfe, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrfe > 0 {
		zrfe--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "p":
			z.Poker, bts, err = msgp.ReadBytesBytes(bts, z.Poker)
			if err != nil {
				return
			}
		case "odd":
			var zgmo uint32
			zgmo, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Odd) >= int(zgmo) {
				z.Odd = (z.Odd)[:zgmo]
			} else {
				z.Odd = make([]int32, zgmo)
			}
			for zpez := range z.Odd {
				z.Odd[zpez], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "bet":
			var ztaf uint32
			ztaf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(ztaf) {
				z.Bet = (z.Bet)[:ztaf]
			} else {
				z.Bet = make([]int64, ztaf)
			}
			for zqke := range z.Bet {
				z.Bet[zqke], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "rich":
			var zeth uint32
			zeth, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zeth) {
				z.Rich = (z.Rich)[:zeth]
			} else {
				z.Rich = make([]int64, zeth)
			}
			for zqyh := range z.Rich {
				z.Rich[zqyh], bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *GameResult) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 2 + msgp.BytesPrefixSize + len(z.Poker) + 4 + msgp.ArrayHeaderSize + (len(z.Odd) * (msgp.Int32Size)) + 4 + msgp.ArrayHeaderSize + (len(z.Bet) * (msgp.Int64Size)) + 5 + msgp.ArrayHeaderSize + (len(z.Rich) * (msgp.Int64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameRound) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zzdc uint32
	zzdc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zzdc > 0 {
		zzdc--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "_id":
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
		case "l":
			var zelx uint32
			zelx, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(zelx) {
				z.Bill = (z.Bill)[:zelx]
			} else {
				z.Bill = make([]*GameBill, zelx)
			}
			for zsbz := range z.Bill {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Bill[zsbz] = nil
				} else {
					if z.Bill[zsbz] == nil {
						z.Bill[zsbz] = new(GameBill)
					}
					err = z.Bill[zsbz].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "f":
			var zbal uint32
			zbal, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Flow) >= int(zbal) {
				z.Flow = (z.Flow)[:zbal]
			} else {
				z.Flow = make([]int32, zbal)
			}
			for zrjx := range z.Flow {
				z.Flow[zrjx], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "a":
			var zjqz uint32
			zjqz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.AllBet) >= int(zjqz) {
				z.AllBet = (z.AllBet)[:zjqz]
			} else {
				z.AllBet = make([]int64, zjqz)
			}
			for zawn := range z.AllBet {
				z.AllBet[zawn], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "p":
			z.Poker, err = dc.ReadBytes(z.Poker)
			if err != nil {
				return
			}
		case "o":
			var zkct uint32
			zkct, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Odds) >= int(zkct) {
				z.Odds = (z.Odds)[:zkct]
			} else {
				z.Odds = make([]int32, zkct)
			}
			for zwel := range z.Odds {
				z.Odds[zwel], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "x":
			z.Tax, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "w":
			z.Win, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "n":
			z.Note, err = dc.ReadString()
			if err != nil {
				return
			}
		case "h":
			var ztmt uint32
			ztmt, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(ztmt) {
				z.Rich = (z.Rich)[:ztmt]
			} else {
				z.Rich = make([]int32, ztmt)
			}
			for zrbe := range z.Rich {
				z.Rich[zrbe], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "u":
			var ztco uint32
			ztco, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.UserBet) >= int(ztco) {
				z.UserBet = (z.UserBet)[:ztco]
			} else {
				z.UserBet = make([]int64, ztco)
			}
			for zmfd := range z.UserBet {
				z.UserBet[zmfd], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "b":
			z.Bank, err = dc.ReadInt32()
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
	// map header, size 16
	// write "_id"
	err = en.Append(0xde, 0x0, 0x10, 0xa3, 0x5f, 0x69, 0x64)
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
	// write "l"
	err = en.Append(0xa1, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Bill)))
	if err != nil {
		return
	}
	for zsbz := range z.Bill {
		if z.Bill[zsbz] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Bill[zsbz].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "f"
	err = en.Append(0xa1, 0x66)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Flow)))
	if err != nil {
		return
	}
	for zrjx := range z.Flow {
		err = en.WriteInt32(z.Flow[zrjx])
		if err != nil {
			return
		}
	}
	// write "a"
	err = en.Append(0xa1, 0x61)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.AllBet)))
	if err != nil {
		return
	}
	for zawn := range z.AllBet {
		err = en.WriteInt64(z.AllBet[zawn])
		if err != nil {
			return
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
	// write "o"
	err = en.Append(0xa1, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Odds)))
	if err != nil {
		return
	}
	for zwel := range z.Odds {
		err = en.WriteInt32(z.Odds[zwel])
		if err != nil {
			return
		}
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
	// write "w"
	err = en.Append(0xa1, 0x77)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Win)
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
	// write "h"
	err = en.Append(0xa1, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Rich)))
	if err != nil {
		return
	}
	for zrbe := range z.Rich {
		err = en.WriteInt32(z.Rich[zrbe])
		if err != nil {
			return
		}
	}
	// write "u"
	err = en.Append(0xa1, 0x75)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.UserBet)))
	if err != nil {
		return
	}
	for zmfd := range z.UserBet {
		err = en.WriteInt64(z.UserBet[zmfd])
		if err != nil {
			return
		}
	}
	// write "b"
	err = en.Append(0xa1, 0x62)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Bank)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameRound) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 16
	// string "_id"
	o = append(o, 0xde, 0x0, 0x10, 0xa3, 0x5f, 0x69, 0x64)
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
	// string "l"
	o = append(o, 0xa1, 0x6c)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bill)))
	for zsbz := range z.Bill {
		if z.Bill[zsbz] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Bill[zsbz].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "f"
	o = append(o, 0xa1, 0x66)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Flow)))
	for zrjx := range z.Flow {
		o = msgp.AppendInt32(o, z.Flow[zrjx])
	}
	// string "a"
	o = append(o, 0xa1, 0x61)
	o = msgp.AppendArrayHeader(o, uint32(len(z.AllBet)))
	for zawn := range z.AllBet {
		o = msgp.AppendInt64(o, z.AllBet[zawn])
	}
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendBytes(o, z.Poker)
	// string "o"
	o = append(o, 0xa1, 0x6f)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Odds)))
	for zwel := range z.Odds {
		o = msgp.AppendInt32(o, z.Odds[zwel])
	}
	// string "x"
	o = append(o, 0xa1, 0x78)
	o = msgp.AppendInt64(o, z.Tax)
	// string "w"
	o = append(o, 0xa1, 0x77)
	o = msgp.AppendInt64(o, z.Win)
	// string "n"
	o = append(o, 0xa1, 0x6e)
	o = msgp.AppendString(o, z.Note)
	// string "h"
	o = append(o, 0xa1, 0x68)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rich)))
	for zrbe := range z.Rich {
		o = msgp.AppendInt32(o, z.Rich[zrbe])
	}
	// string "u"
	o = append(o, 0xa1, 0x75)
	o = msgp.AppendArrayHeader(o, uint32(len(z.UserBet)))
	for zmfd := range z.UserBet {
		o = msgp.AppendInt64(o, z.UserBet[zmfd])
	}
	// string "b"
	o = append(o, 0xa1, 0x62)
	o = msgp.AppendInt32(o, z.Bank)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameRound) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zana uint32
	zana, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zana > 0 {
		zana--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "_id":
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
		case "l":
			var ztyy uint32
			ztyy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(ztyy) {
				z.Bill = (z.Bill)[:ztyy]
			} else {
				z.Bill = make([]*GameBill, ztyy)
			}
			for zsbz := range z.Bill {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Bill[zsbz] = nil
				} else {
					if z.Bill[zsbz] == nil {
						z.Bill[zsbz] = new(GameBill)
					}
					bts, err = z.Bill[zsbz].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "f":
			var zinl uint32
			zinl, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Flow) >= int(zinl) {
				z.Flow = (z.Flow)[:zinl]
			} else {
				z.Flow = make([]int32, zinl)
			}
			for zrjx := range z.Flow {
				z.Flow[zrjx], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "a":
			var zare uint32
			zare, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.AllBet) >= int(zare) {
				z.AllBet = (z.AllBet)[:zare]
			} else {
				z.AllBet = make([]int64, zare)
			}
			for zawn := range z.AllBet {
				z.AllBet[zawn], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "p":
			z.Poker, bts, err = msgp.ReadBytesBytes(bts, z.Poker)
			if err != nil {
				return
			}
		case "o":
			var zljy uint32
			zljy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Odds) >= int(zljy) {
				z.Odds = (z.Odds)[:zljy]
			} else {
				z.Odds = make([]int32, zljy)
			}
			for zwel := range z.Odds {
				z.Odds[zwel], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "x":
			z.Tax, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "w":
			z.Win, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "n":
			z.Note, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "h":
			var zixj uint32
			zixj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zixj) {
				z.Rich = (z.Rich)[:zixj]
			} else {
				z.Rich = make([]int32, zixj)
			}
			for zrbe := range z.Rich {
				z.Rich[zrbe], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "u":
			var zrsc uint32
			zrsc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.UserBet) >= int(zrsc) {
				z.UserBet = (z.UserBet)[:zrsc]
			} else {
				z.UserBet = make([]int64, zrsc)
			}
			for zmfd := range z.UserBet {
				z.UserBet[zmfd], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "b":
			z.Bank, bts, err = msgp.ReadInt32Bytes(bts)
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
	s = 3 + 4 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.ArrayHeaderSize
	for zsbz := range z.Bill {
		if z.Bill[zsbz] == nil {
			s += msgp.NilSize
		} else {
			s += z.Bill[zsbz].Msgsize()
		}
	}
	s += 2 + msgp.ArrayHeaderSize + (len(z.Flow) * (msgp.Int32Size)) + 2 + msgp.ArrayHeaderSize + (len(z.AllBet) * (msgp.Int64Size)) + 2 + msgp.BytesPrefixSize + len(z.Poker) + 2 + msgp.ArrayHeaderSize + (len(z.Odds) * (msgp.Int32Size)) + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.StringPrefixSize + len(z.Note) + 2 + msgp.ArrayHeaderSize + (len(z.Rich) * (msgp.Int32Size)) + 2 + msgp.ArrayHeaderSize + (len(z.UserBet) * (msgp.Int64Size)) + 2 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OpenBetAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zswy uint32
	zswy, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zswy > 0 {
		zswy--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "state":
			z.State, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "rich":
			var znsg uint32
			znsg, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(znsg) {
				z.Rich = (z.Rich)[:znsg]
			} else {
				z.Rich = make([]*Player, znsg)
			}
			for zctn := range z.Rich {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Rich[zctn] = nil
				} else {
					if z.Rich[zctn] == nil {
						z.Rich[zctn] = new(Player)
					}
					err = z.Rich[zctn].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "bank":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Bank = nil
			} else {
				if z.Bank == nil {
					z.Bank = new(Player)
				}
				err = z.Bank.DecodeMsg(dc)
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
func (z *OpenBetAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "id"
	err = en.Append(0x84, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Id)
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
	// write "rich"
	err = en.Append(0xa4, 0x72, 0x69, 0x63, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Rich)))
	if err != nil {
		return
	}
	for zctn := range z.Rich {
		if z.Rich[zctn] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Rich[zctn].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "bank"
	err = en.Append(0xa4, 0x62, 0x61, 0x6e, 0x6b)
	if err != nil {
		return err
	}
	if z.Bank == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Bank.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *OpenBetAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "id"
	o = append(o, 0x84, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Id)
	// string "state"
	o = append(o, 0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	o = msgp.AppendInt32(o, z.State)
	// string "rich"
	o = append(o, 0xa4, 0x72, 0x69, 0x63, 0x68)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rich)))
	for zctn := range z.Rich {
		if z.Rich[zctn] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Rich[zctn].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "bank"
	o = append(o, 0xa4, 0x62, 0x61, 0x6e, 0x6b)
	if z.Bank == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Bank.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OpenBetAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrus uint32
	zrus, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrus > 0 {
		zrus--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "state":
			z.State, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "rich":
			var zsvm uint32
			zsvm, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zsvm) {
				z.Rich = (z.Rich)[:zsvm]
			} else {
				z.Rich = make([]*Player, zsvm)
			}
			for zctn := range z.Rich {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Rich[zctn] = nil
				} else {
					if z.Rich[zctn] == nil {
						z.Rich[zctn] = new(Player)
					}
					bts, err = z.Rich[zctn].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "bank":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Bank = nil
			} else {
				if z.Bank == nil {
					z.Bank = new(Player)
				}
				bts, err = z.Bank.UnmarshalMsg(bts)
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
func (z *OpenBetAck) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 6 + msgp.Int32Size + 5 + msgp.ArrayHeaderSize
	for zctn := range z.Rich {
		if z.Rich[zctn] == nil {
			s += msgp.NilSize
		} else {
			s += z.Rich[zctn].Msgsize()
		}
	}
	s += 5
	if z.Bank == nil {
		s += msgp.NilSize
	} else {
		s += z.Bank.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Player) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zaoz uint32
	zaoz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zaoz > 0 {
		zaoz--
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
func (z *Player) EncodeMsg(en *msgp.Writer) (err error) {
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
func (z *Player) MarshalMsg(b []byte) (o []byte, err error) {
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
func (z *Player) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zfzb uint32
	zfzb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zfzb > 0 {
		zfzb--
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
func (z *Player) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 5 + msgp.Int32Size + 4 + msgp.Int32Size + 5 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserBetAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zjif uint32
	zjif, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjif > 0 {
		zjif--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "item":
			var zqgz uint32
			zqgz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zqgz) {
				z.Bet = (z.Bet)[:zqgz]
			} else {
				z.Bet = make([]int32, zqgz)
			}
			for zsbo := range z.Bet {
				z.Bet[zsbo], err = dc.ReadInt32()
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
func (z *UserBetAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "item"
	err = en.Append(0x81, 0xa4, 0x69, 0x74, 0x65, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Bet)))
	if err != nil {
		return
	}
	for zsbo := range z.Bet {
		err = en.WriteInt32(z.Bet[zsbo])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UserBetAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "item"
	o = append(o, 0x81, 0xa4, 0x69, 0x74, 0x65, 0x6d)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bet)))
	for zsbo := range z.Bet {
		o = msgp.AppendInt32(o, z.Bet[zsbo])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserBetAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsnw uint32
	zsnw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsnw > 0 {
		zsnw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "item":
			var ztls uint32
			ztls, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(ztls) {
				z.Bet = (z.Bet)[:ztls]
			} else {
				z.Bet = make([]int32, ztls)
			}
			for zsbo := range z.Bet {
				z.Bet[zsbo], bts, err = msgp.ReadInt32Bytes(bts)
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
func (z *UserBetAck) Msgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize + (len(z.Bet) * (msgp.Int32Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserLog) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zigk uint32
	zigk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zigk > 0 {
		zigk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "l":
			z.Log, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "t":
			z.Tab, err = dc.ReadInt32()
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
		case "g":
			var zopb uint32
			zopb, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zopb) {
				z.Group = (z.Group)[:zopb]
			} else {
				z.Group = make([]int64, zopb)
			}
			for zmvo := range z.Group {
				z.Group[zmvo], err = dc.ReadInt64()
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
func (z *UserLog) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "l"
	err = en.Append(0x85, 0xa1, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Log)
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
	// write "g"
	err = en.Append(0xa1, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Group)))
	if err != nil {
		return
	}
	for zmvo := range z.Group {
		err = en.WriteInt64(z.Group[zmvo])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UserLog) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "l"
	o = append(o, 0x85, 0xa1, 0x6c)
	o = msgp.AppendInt64(o, z.Log)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendInt32(o, z.Tab)
	// string "b"
	o = append(o, 0xa1, 0x62)
	o = msgp.AppendInt64(o, z.Bet)
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendBytes(o, z.Poker)
	// string "g"
	o = append(o, 0xa1, 0x67)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Group)))
	for zmvo := range z.Group {
		o = msgp.AppendInt64(o, z.Group[zmvo])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserLog) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zuop uint32
	zuop, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zuop > 0 {
		zuop--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "l":
			z.Log, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "t":
			z.Tab, bts, err = msgp.ReadInt32Bytes(bts)
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
		case "g":
			var zedl uint32
			zedl, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zedl) {
				z.Group = (z.Group)[:zedl]
			} else {
				z.Group = make([]int64, zedl)
			}
			for zmvo := range z.Group {
				z.Group[zmvo], bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *UserLog) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int64Size + 2 + msgp.BytesPrefixSize + len(z.Poker) + 2 + msgp.ArrayHeaderSize + (len(z.Group) * (msgp.Int64Size))
	return
}
