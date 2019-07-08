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
func (z *Folks_Code) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zajw int32
		zajw, err = dc.ReadInt32()
		(*z) = Folks_Code(zajw)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Folks_Code) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt32(int32(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Folks_Code) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt32(o, int32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Folks_Code) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zwht int32
		zwht, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = Folks_Code(zwht)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Folks_Code) Msgsize() (s int) {
	s = msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameDealAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zhct uint32
	zhct, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zhct > 0 {
		zhct--
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
func (z *GameDealAck) EncodeMsg(en *msgp.Writer) (err error) {
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
func (z *GameDealAck) MarshalMsg(b []byte) (o []byte, err error) {
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
func (z *GameDealAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcua uint32
	zcua, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcua > 0 {
		zcua--
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
func (z *GameDealAck) Msgsize() (s int) {
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
func (z *GameInitAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "state":
			z.State, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "time":
			z.Time, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "rich":
			var zjfb uint32
			zjfb, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zjfb) {
				z.Rich = (z.Rich)[:zjfb]
			} else {
				z.Rich = make([]*Player, zjfb)
			}
			for zxhx := range z.Rich {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Rich[zxhx] = nil
				} else {
					if z.Rich[zxhx] == nil {
						z.Rich[zxhx] = new(Player)
					}
					err = z.Rich[zxhx].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "sum":
			var zcxo uint32
			zcxo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Sum) >= int(zcxo) {
				z.Sum = (z.Sum)[:zcxo]
			} else {
				z.Sum = make([]int64, zcxo)
			}
			for zlqf := range z.Sum {
				z.Sum[zlqf], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "bet":
			var zeff uint32
			zeff, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zeff) {
				z.Bet = (z.Bet)[:zeff]
			} else {
				z.Bet = make([]int64, zeff)
			}
			for zdaf := range z.Bet {
				z.Bet[zdaf], err = dc.ReadInt64()
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
	// map header, size 9
	// write "table"
	err = en.Append(0x89, 0xa5, 0x74, 0x61, 0x62, 0x6c, 0x65)
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
	// write "state"
	err = en.Append(0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.State)
	if err != nil {
		return
	}
	// write "time"
	err = en.Append(0xa4, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Time)
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
	for zxhx := range z.Rich {
		if z.Rich[zxhx] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Rich[zxhx].EncodeMsg(en)
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
	for zlqf := range z.Sum {
		err = en.WriteInt64(z.Sum[zlqf])
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
	for zdaf := range z.Bet {
		err = en.WriteInt64(z.Bet[zdaf])
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
	// map header, size 9
	// string "table"
	o = append(o, 0x89, 0xa5, 0x74, 0x61, 0x62, 0x6c, 0x65)
	o = msgp.AppendInt32(o, z.Table)
	// string "id"
	o = append(o, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "state"
	o = append(o, 0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	o = msgp.AppendInt32(o, z.State)
	// string "time"
	o = append(o, 0xa4, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt32(o, z.Time)
	// string "rich"
	o = append(o, 0xa4, 0x72, 0x69, 0x63, 0x68)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rich)))
	for zxhx := range z.Rich {
		if z.Rich[zxhx] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Rich[zxhx].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "sum"
	o = append(o, 0xa3, 0x73, 0x75, 0x6d)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Sum)))
	for zlqf := range z.Sum {
		o = msgp.AppendInt64(o, z.Sum[zlqf])
	}
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bet)))
	for zdaf := range z.Bet {
		o = msgp.AppendInt64(o, z.Bet[zdaf])
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
	var zrsw uint32
	zrsw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrsw > 0 {
		zrsw--
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
		case "state":
			z.State, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "time":
			z.Time, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "rich":
			var zxpk uint32
			zxpk, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zxpk) {
				z.Rich = (z.Rich)[:zxpk]
			} else {
				z.Rich = make([]*Player, zxpk)
			}
			for zxhx := range z.Rich {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Rich[zxhx] = nil
				} else {
					if z.Rich[zxhx] == nil {
						z.Rich[zxhx] = new(Player)
					}
					bts, err = z.Rich[zxhx].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "sum":
			var zdnj uint32
			zdnj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Sum) >= int(zdnj) {
				z.Sum = (z.Sum)[:zdnj]
			} else {
				z.Sum = make([]int64, zdnj)
			}
			for zlqf := range z.Sum {
				z.Sum[zlqf], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "bet":
			var zobc uint32
			zobc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zobc) {
				z.Bet = (z.Bet)[:zobc]
			} else {
				z.Bet = make([]int64, zobc)
			}
			for zdaf := range z.Bet {
				z.Bet[zdaf], bts, err = msgp.ReadInt64Bytes(bts)
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
	s = 1 + 6 + msgp.Int32Size + 3 + msgp.Int32Size + 6 + msgp.Int32Size + 5 + msgp.Int32Size + 5 + msgp.ArrayHeaderSize
	for zxhx := range z.Rich {
		if z.Rich[zxhx] == nil {
			s += msgp.NilSize
		} else {
			s += z.Rich[zxhx].Msgsize()
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
	var zpez uint32
	zpez, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zpez > 0 {
		zpez--
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
		case "p":
			z.Poker, err = dc.ReadBytes(z.Poker)
			if err != nil {
				return
			}
		case "odd":
			var zqke uint32
			zqke, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Odd) >= int(zqke) {
				z.Odd = (z.Odd)[:zqke]
			} else {
				z.Odd = make([]int32, zqke)
			}
			for zsnv := range z.Odd {
				z.Odd[zsnv], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "sum":
			var zqyh uint32
			zqyh, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Sum) >= int(zqyh) {
				z.Sum = (z.Sum)[:zqyh]
			} else {
				z.Sum = make([]int64, zqyh)
			}
			for zkgt := range z.Sum {
				z.Sum[zkgt], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "rich":
			var zyzr uint32
			zyzr, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zyzr) {
				z.Rich = (z.Rich)[:zyzr]
			} else {
				z.Rich = make([]int64, zyzr)
			}
			for zema := range z.Rich {
				z.Rich[zema], err = dc.ReadInt64()
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
	err = en.WriteInt32(z.Id)
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
	for zsnv := range z.Odd {
		err = en.WriteInt32(z.Odd[zsnv])
		if err != nil {
			return
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
	for zkgt := range z.Sum {
		err = en.WriteInt64(z.Sum[zkgt])
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
	for zema := range z.Rich {
		err = en.WriteInt64(z.Rich[zema])
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
	o = msgp.AppendInt32(o, z.Id)
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendBytes(o, z.Poker)
	// string "odd"
	o = append(o, 0xa3, 0x6f, 0x64, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Odd)))
	for zsnv := range z.Odd {
		o = msgp.AppendInt32(o, z.Odd[zsnv])
	}
	// string "sum"
	o = append(o, 0xa3, 0x73, 0x75, 0x6d)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Sum)))
	for zkgt := range z.Sum {
		o = msgp.AppendInt64(o, z.Sum[zkgt])
	}
	// string "rich"
	o = append(o, 0xa4, 0x72, 0x69, 0x63, 0x68)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rich)))
	for zema := range z.Rich {
		o = msgp.AppendInt64(o, z.Rich[zema])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameResult) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "p":
			z.Poker, bts, err = msgp.ReadBytesBytes(bts, z.Poker)
			if err != nil {
				return
			}
		case "odd":
			var zjpj uint32
			zjpj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Odd) >= int(zjpj) {
				z.Odd = (z.Odd)[:zjpj]
			} else {
				z.Odd = make([]int32, zjpj)
			}
			for zsnv := range z.Odd {
				z.Odd[zsnv], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "sum":
			var zzpf uint32
			zzpf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Sum) >= int(zzpf) {
				z.Sum = (z.Sum)[:zzpf]
			} else {
				z.Sum = make([]int64, zzpf)
			}
			for zkgt := range z.Sum {
				z.Sum[zkgt], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "rich":
			var zrfe uint32
			zrfe, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zrfe) {
				z.Rich = (z.Rich)[:zrfe]
			} else {
				z.Rich = make([]int64, zrfe)
			}
			for zema := range z.Rich {
				z.Rich[zema], bts, err = msgp.ReadInt64Bytes(bts)
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
	s = 1 + 3 + msgp.Int32Size + 2 + msgp.BytesPrefixSize + len(z.Poker) + 4 + msgp.ArrayHeaderSize + (len(z.Odd) * (msgp.Int32Size)) + 4 + msgp.ArrayHeaderSize + (len(z.Sum) * (msgp.Int64Size)) + 5 + msgp.ArrayHeaderSize + (len(z.Rich) * (msgp.Int64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LeaveAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "u":
			z.Uid, err = dc.ReadInt32()
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
func (z LeaveAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "u"
	err = en.Append(0x81, 0xa1, 0x75)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Uid)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z LeaveAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "u"
	o = append(o, 0x81, 0xa1, 0x75)
	o = msgp.AppendInt32(o, z.Uid)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LeaveAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var ztaf uint32
	ztaf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztaf > 0 {
		ztaf--
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
func (z LeaveAck) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LeaveReq) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zeth uint32
	zeth, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zeth > 0 {
		zeth--
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
func (z LeaveReq) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "u"
	err = en.Append(0x81, 0xa1, 0x75)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Uid)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z LeaveReq) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "u"
	o = append(o, 0x81, 0xa1, 0x75)
	o = msgp.AppendInt32(o, z.Uid)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LeaveReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "u":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
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
func (z LeaveReq) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OpenBetAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zawn uint32
	zawn, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zawn > 0 {
		zawn--
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
		case "time":
			z.Time, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "rich":
			var zwel uint32
			zwel, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zwel) {
				z.Rich = (z.Rich)[:zwel]
			} else {
				z.Rich = make([]*Player, zwel)
			}
			for zrjx := range z.Rich {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Rich[zrjx] = nil
				} else {
					if z.Rich[zrjx] == nil {
						z.Rich[zrjx] = new(Player)
					}
					err = z.Rich[zrjx].DecodeMsg(dc)
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
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
	}
	// write "time"
	err = en.Append(0xa4, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Time)
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
	for zrjx := range z.Rich {
		if z.Rich[zrjx] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Rich[zrjx].EncodeMsg(en)
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
	o = msgp.AppendInt32(o, z.Id)
	// string "time"
	o = append(o, 0xa4, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt32(o, z.Time)
	// string "rich"
	o = append(o, 0xa4, 0x72, 0x69, 0x63, 0x68)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rich)))
	for zrjx := range z.Rich {
		if z.Rich[zrjx] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Rich[zrjx].MarshalMsg(o)
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
	var zrbe uint32
	zrbe, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrbe > 0 {
		zrbe--
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
		case "time":
			z.Time, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "rich":
			var zmfd uint32
			zmfd, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zmfd) {
				z.Rich = (z.Rich)[:zmfd]
			} else {
				z.Rich = make([]*Player, zmfd)
			}
			for zrjx := range z.Rich {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Rich[zrjx] = nil
				} else {
					if z.Rich[zrjx] == nil {
						z.Rich[zrjx] = new(Player)
					}
					bts, err = z.Rich[zrjx].UnmarshalMsg(bts)
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
	s = 1 + 3 + msgp.Int32Size + 5 + msgp.Int32Size + 5 + msgp.ArrayHeaderSize
	for zrjx := range z.Rich {
		if z.Rich[zrjx] == nil {
			s += msgp.NilSize
		} else {
			s += z.Rich[zrjx].Msgsize()
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
	var zelx uint32
	zelx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zelx > 0 {
		zelx--
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
func (z *StopBetAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
			z.Id, err = dc.ReadInt64()
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
func (z StopBetAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "id"
	err = en.Append(0x81, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Id)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z StopBetAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "id"
	o = append(o, 0x81, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Id)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StopBetAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjqz uint32
	zjqz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjqz > 0 {
		zjqz--
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
func (z StopBetAck) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserBetAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var ztmt uint32
	ztmt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ztmt > 0 {
		ztmt--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "time":
			z.Time, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "item":
			var ztco uint32
			ztco, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(ztco) {
				z.Bet = (z.Bet)[:ztco]
			} else {
				z.Bet = make([]int32, ztco)
			}
			for zkct := range z.Bet {
				z.Bet[zkct], err = dc.ReadInt32()
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
	// map header, size 2
	// write "time"
	err = en.Append(0x82, 0xa4, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Time)
	if err != nil {
		return
	}
	// write "item"
	err = en.Append(0xa4, 0x69, 0x74, 0x65, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Bet)))
	if err != nil {
		return
	}
	for zkct := range z.Bet {
		err = en.WriteInt32(z.Bet[zkct])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UserBetAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "time"
	o = append(o, 0x82, 0xa4, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt32(o, z.Time)
	// string "item"
	o = append(o, 0xa4, 0x69, 0x74, 0x65, 0x6d)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bet)))
	for zkct := range z.Bet {
		o = msgp.AppendInt32(o, z.Bet[zkct])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserBetAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "time":
			z.Time, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "item":
			var ztyy uint32
			ztyy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(ztyy) {
				z.Bet = (z.Bet)[:ztyy]
			} else {
				z.Bet = make([]int32, ztyy)
			}
			for zkct := range z.Bet {
				z.Bet[zkct], bts, err = msgp.ReadInt32Bytes(bts)
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
	s = 1 + 5 + msgp.Int32Size + 5 + msgp.ArrayHeaderSize + (len(z.Bet) * (msgp.Int32Size))
	return
}
