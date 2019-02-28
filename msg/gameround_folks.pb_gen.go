package msg

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
		case "item":
			z.Item, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "coin":
			z.Coin, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "bet":
			var zbai uint32
			zbai, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zbai) {
				z.Bet = (z.Bet)[:zbai]
			} else {
				z.Bet = make([]int64, zbai)
			}
			for zxvk := range z.Bet {
				z.Bet[zxvk], err = dc.ReadInt64()
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
func (z *BetAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "item"
	err = en.Append(0x83, 0xa4, 0x69, 0x74, 0x65, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Item)
	if err != nil {
		return
	}
	// write "coin"
	err = en.Append(0xa4, 0x63, 0x6f, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Coin)
	if err != nil {
		return
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
	for zxvk := range z.Bet {
		err = en.WriteInt64(z.Bet[zxvk])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BetAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "item"
	o = append(o, 0x83, 0xa4, 0x69, 0x74, 0x65, 0x6d)
	o = msgp.AppendInt32(o, z.Item)
	// string "coin"
	o = append(o, 0xa4, 0x63, 0x6f, 0x69, 0x6e)
	o = msgp.AppendInt32(o, z.Coin)
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bet)))
	for zxvk := range z.Bet {
		o = msgp.AppendInt64(o, z.Bet[zxvk])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BetAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "item":
			z.Item, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "coin":
			z.Coin, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "bet":
			var zajw uint32
			zajw, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zajw) {
				z.Bet = (z.Bet)[:zajw]
			} else {
				z.Bet = make([]int64, zajw)
			}
			for zxvk := range z.Bet {
				z.Bet[zxvk], bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *BetAck) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int32Size + 5 + msgp.Int32Size + 4 + msgp.ArrayHeaderSize + (len(z.Bet) * (msgp.Int64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BetReq) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zwht uint32
	zwht, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zwht > 0 {
		zwht--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "item":
			z.Item, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "coin":
			z.Coin, err = dc.ReadInt32()
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
	// map header, size 2
	// write "item"
	err = en.Append(0x82, 0xa4, 0x69, 0x74, 0x65, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Item)
	if err != nil {
		return
	}
	// write "coin"
	err = en.Append(0xa4, 0x63, 0x6f, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Coin)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z BetReq) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "item"
	o = append(o, 0x82, 0xa4, 0x69, 0x74, 0x65, 0x6d)
	o = msgp.AppendInt32(o, z.Item)
	// string "coin"
	o = append(o, 0xa4, 0x63, 0x6f, 0x69, 0x6e)
	o = msgp.AppendInt32(o, z.Coin)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BetReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "item":
			z.Item, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "coin":
			z.Coin, bts, err = msgp.ReadInt32Bytes(bts)
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
	s = 1 + 5 + msgp.Int32Size + 5 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *CloseBetAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "r":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.R = nil
			} else {
				if z.R == nil {
					z.R = new(FolksGameResult)
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
	var zxhx uint32
	zxhx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
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
					z.R = new(FolksGameResult)
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
func (z *FolksGameInitAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zjfb uint32
	zjfb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjfb > 0 {
		zjfb--
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
			var zcxo uint32
			zcxo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zcxo) {
				z.Rich = (z.Rich)[:zcxo]
			} else {
				z.Rich = make([]*User, zcxo)
			}
			for zlqf := range z.Rich {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Rich[zlqf] = nil
				} else {
					if z.Rich[zlqf] == nil {
						z.Rich[zlqf] = new(User)
					}
					err = z.Rich[zlqf].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "sum":
			var zeff uint32
			zeff, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Sum) >= int(zeff) {
				z.Sum = (z.Sum)[:zeff]
			} else {
				z.Sum = make([]int64, zeff)
			}
			for zdaf := range z.Sum {
				z.Sum[zdaf], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "bet":
			var zrsw uint32
			zrsw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zrsw) {
				z.Bet = (z.Bet)[:zrsw]
			} else {
				z.Bet = make([]int64, zrsw)
			}
			for zpks := range z.Bet {
				z.Bet[zpks], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "log":
			z.Log, err = dc.ReadBytes(z.Log)
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
func (z *FolksGameInitAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "id"
	err = en.Append(0x86, 0xa2, 0x69, 0x64)
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
	for zlqf := range z.Rich {
		if z.Rich[zlqf] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Rich[zlqf].EncodeMsg(en)
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
	for zdaf := range z.Sum {
		err = en.WriteInt64(z.Sum[zdaf])
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
	for zpks := range z.Bet {
		err = en.WriteInt64(z.Bet[zpks])
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
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FolksGameInitAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "id"
	o = append(o, 0x86, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Id)
	// string "state"
	o = append(o, 0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	o = msgp.AppendInt32(o, z.State)
	// string "rich"
	o = append(o, 0xa4, 0x72, 0x69, 0x63, 0x68)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rich)))
	for zlqf := range z.Rich {
		if z.Rich[zlqf] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Rich[zlqf].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "sum"
	o = append(o, 0xa3, 0x73, 0x75, 0x6d)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Sum)))
	for zdaf := range z.Sum {
		o = msgp.AppendInt64(o, z.Sum[zdaf])
	}
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bet)))
	for zpks := range z.Bet {
		o = msgp.AppendInt64(o, z.Bet[zpks])
	}
	// string "log"
	o = append(o, 0xa3, 0x6c, 0x6f, 0x67)
	o = msgp.AppendBytes(o, z.Log)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FolksGameInitAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
			var zdnj uint32
			zdnj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zdnj) {
				z.Rich = (z.Rich)[:zdnj]
			} else {
				z.Rich = make([]*User, zdnj)
			}
			for zlqf := range z.Rich {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Rich[zlqf] = nil
				} else {
					if z.Rich[zlqf] == nil {
						z.Rich[zlqf] = new(User)
					}
					bts, err = z.Rich[zlqf].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "sum":
			var zobc uint32
			zobc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Sum) >= int(zobc) {
				z.Sum = (z.Sum)[:zobc]
			} else {
				z.Sum = make([]int64, zobc)
			}
			for zdaf := range z.Sum {
				z.Sum[zdaf], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "bet":
			var zsnv uint32
			zsnv, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zsnv) {
				z.Bet = (z.Bet)[:zsnv]
			} else {
				z.Bet = make([]int64, zsnv)
			}
			for zpks := range z.Bet {
				z.Bet[zpks], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "log":
			z.Log, bts, err = msgp.ReadBytesBytes(bts, z.Log)
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
func (z *FolksGameInitAck) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 6 + msgp.Int32Size + 5 + msgp.ArrayHeaderSize
	for zlqf := range z.Rich {
		if z.Rich[zlqf] == nil {
			s += msgp.NilSize
		} else {
			s += z.Rich[zlqf].Msgsize()
		}
	}
	s += 4 + msgp.ArrayHeaderSize + (len(z.Sum) * (msgp.Int64Size)) + 4 + msgp.ArrayHeaderSize + (len(z.Bet) * (msgp.Int64Size)) + 4 + msgp.BytesPrefixSize + len(z.Log)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FolksGameResult) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zqke uint32
	zqke, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zqke > 0 {
		zqke--
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
		case "a":
			z.A, err = dc.ReadBytes(z.A)
			if err != nil {
				return
			}
		case "b":
			z.B, err = dc.ReadBytes(z.B)
			if err != nil {
				return
			}
		case "odd":
			var zqyh uint32
			zqyh, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Odd) >= int(zqyh) {
				z.Odd = (z.Odd)[:zqyh]
			} else {
				z.Odd = make([]int32, zqyh)
			}
			for zkgt := range z.Odd {
				z.Odd[zkgt], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "bet":
			var zyzr uint32
			zyzr, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zyzr) {
				z.Bet = (z.Bet)[:zyzr]
			} else {
				z.Bet = make([]int64, zyzr)
			}
			for zema := range z.Bet {
				z.Bet[zema], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "rich":
			var zywj uint32
			zywj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zywj) {
				z.Rich = (z.Rich)[:zywj]
			} else {
				z.Rich = make([]int64, zywj)
			}
			for zpez := range z.Rich {
				z.Rich[zpez], err = dc.ReadInt64()
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
func (z *FolksGameResult) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "id"
	err = en.Append(0x86, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Id)
	if err != nil {
		return
	}
	// write "a"
	err = en.Append(0xa1, 0x61)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.A)
	if err != nil {
		return
	}
	// write "b"
	err = en.Append(0xa1, 0x62)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.B)
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
	for zkgt := range z.Odd {
		err = en.WriteInt32(z.Odd[zkgt])
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
	for zema := range z.Bet {
		err = en.WriteInt64(z.Bet[zema])
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
	for zpez := range z.Rich {
		err = en.WriteInt64(z.Rich[zpez])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FolksGameResult) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "id"
	o = append(o, 0x86, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Id)
	// string "a"
	o = append(o, 0xa1, 0x61)
	o = msgp.AppendBytes(o, z.A)
	// string "b"
	o = append(o, 0xa1, 0x62)
	o = msgp.AppendBytes(o, z.B)
	// string "odd"
	o = append(o, 0xa3, 0x6f, 0x64, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Odd)))
	for zkgt := range z.Odd {
		o = msgp.AppendInt32(o, z.Odd[zkgt])
	}
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bet)))
	for zema := range z.Bet {
		o = msgp.AppendInt64(o, z.Bet[zema])
	}
	// string "rich"
	o = append(o, 0xa4, 0x72, 0x69, 0x63, 0x68)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rich)))
	for zpez := range z.Rich {
		o = msgp.AppendInt64(o, z.Rich[zpez])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FolksGameResult) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjpj uint32
	zjpj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjpj > 0 {
		zjpj--
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
		case "a":
			z.A, bts, err = msgp.ReadBytesBytes(bts, z.A)
			if err != nil {
				return
			}
		case "b":
			z.B, bts, err = msgp.ReadBytesBytes(bts, z.B)
			if err != nil {
				return
			}
		case "odd":
			var zzpf uint32
			zzpf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Odd) >= int(zzpf) {
				z.Odd = (z.Odd)[:zzpf]
			} else {
				z.Odd = make([]int32, zzpf)
			}
			for zkgt := range z.Odd {
				z.Odd[zkgt], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "bet":
			var zrfe uint32
			zrfe, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zrfe) {
				z.Bet = (z.Bet)[:zrfe]
			} else {
				z.Bet = make([]int64, zrfe)
			}
			for zema := range z.Bet {
				z.Bet[zema], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "rich":
			var zgmo uint32
			zgmo, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zgmo) {
				z.Rich = (z.Rich)[:zgmo]
			} else {
				z.Rich = make([]int64, zgmo)
			}
			for zpez := range z.Rich {
				z.Rich[zpez], bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *FolksGameResult) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 2 + msgp.BytesPrefixSize + len(z.A) + 2 + msgp.BytesPrefixSize + len(z.B) + 4 + msgp.ArrayHeaderSize + (len(z.Odd) * (msgp.Int32Size)) + 4 + msgp.ArrayHeaderSize + (len(z.Bet) * (msgp.Int64Size)) + 5 + msgp.ArrayHeaderSize + (len(z.Rich) * (msgp.Int64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FolksGameRound) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zwel uint32
	zwel, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zwel > 0 {
		zwel--
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
		case "b":
			var zrbe uint32
			zrbe, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(zrbe) {
				z.Bill = (z.Bill)[:zrbe]
			} else {
				z.Bill = make([]*GameBill, zrbe)
			}
			for ztaf := range z.Bill {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Bill[ztaf] = nil
				} else {
					if z.Bill[ztaf] == nil {
						z.Bill[ztaf] = new(GameBill)
					}
					err = z.Bill[ztaf].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "f":
			var zmfd uint32
			zmfd, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Flow) >= int(zmfd) {
				z.Flow = (z.Flow)[:zmfd]
			} else {
				z.Flow = make([]int32, zmfd)
			}
			for zeth := range z.Flow {
				z.Flow[zeth], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "g":
			var zzdc uint32
			zzdc, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zzdc) {
				z.Group = (z.Group)[:zzdc]
			} else {
				z.Group = make([]int64, zzdc)
			}
			for zsbz := range z.Group {
				z.Group[zsbz], err = dc.ReadInt64()
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
			var zelx uint32
			zelx, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Odds) >= int(zelx) {
				z.Odds = (z.Odds)[:zelx]
			} else {
				z.Odds = make([]int32, zelx)
			}
			for zrjx := range z.Odds {
				z.Odds[zrjx], err = dc.ReadInt32()
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
			var zbal uint32
			zbal, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zbal) {
				z.Rich = (z.Rich)[:zbal]
			} else {
				z.Rich = make([]*RichRole, zbal)
			}
			for zawn := range z.Rich {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Rich[zawn] = nil
				} else {
					if z.Rich[zawn] == nil {
						z.Rich[zawn] = new(RichRole)
					}
					var zjqz uint32
					zjqz, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zjqz > 0 {
						zjqz--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "u":
							z.Rich[zawn].Uid, err = dc.ReadInt32()
							if err != nil {
								return
							}
						case "n":
							z.Rich[zawn].Name, err = dc.ReadString()
							if err != nil {
								return
							}
						case "c":
							z.Rich[zawn].Coin, err = dc.ReadInt64()
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
func (z *FolksGameRound) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 14
	// write "_id"
	err = en.Append(0x8e, 0xa3, 0x5f, 0x69, 0x64)
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
	for ztaf := range z.Bill {
		if z.Bill[ztaf] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Bill[ztaf].EncodeMsg(en)
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
	for zeth := range z.Flow {
		err = en.WriteInt32(z.Flow[zeth])
		if err != nil {
			return
		}
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
	for zsbz := range z.Group {
		err = en.WriteInt64(z.Group[zsbz])
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
	for zrjx := range z.Odds {
		err = en.WriteInt32(z.Odds[zrjx])
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
	for zawn := range z.Rich {
		if z.Rich[zawn] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 3
			// write "u"
			err = en.Append(0x83, 0xa1, 0x75)
			if err != nil {
				return err
			}
			err = en.WriteInt32(z.Rich[zawn].Uid)
			if err != nil {
				return
			}
			// write "n"
			err = en.Append(0xa1, 0x6e)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Rich[zawn].Name)
			if err != nil {
				return
			}
			// write "c"
			err = en.Append(0xa1, 0x63)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Rich[zawn].Coin)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FolksGameRound) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 14
	// string "_id"
	o = append(o, 0x8e, 0xa3, 0x5f, 0x69, 0x64)
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
	for ztaf := range z.Bill {
		if z.Bill[ztaf] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Bill[ztaf].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "f"
	o = append(o, 0xa1, 0x66)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Flow)))
	for zeth := range z.Flow {
		o = msgp.AppendInt32(o, z.Flow[zeth])
	}
	// string "g"
	o = append(o, 0xa1, 0x67)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Group)))
	for zsbz := range z.Group {
		o = msgp.AppendInt64(o, z.Group[zsbz])
	}
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendBytes(o, z.Poker)
	// string "o"
	o = append(o, 0xa1, 0x6f)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Odds)))
	for zrjx := range z.Odds {
		o = msgp.AppendInt32(o, z.Odds[zrjx])
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
	for zawn := range z.Rich {
		if z.Rich[zawn] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 3
			// string "u"
			o = append(o, 0x83, 0xa1, 0x75)
			o = msgp.AppendInt32(o, z.Rich[zawn].Uid)
			// string "n"
			o = append(o, 0xa1, 0x6e)
			o = msgp.AppendString(o, z.Rich[zawn].Name)
			// string "c"
			o = append(o, 0xa1, 0x63)
			o = msgp.AppendInt64(o, z.Rich[zawn].Coin)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FolksGameRound) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "b":
			var ztmt uint32
			ztmt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(ztmt) {
				z.Bill = (z.Bill)[:ztmt]
			} else {
				z.Bill = make([]*GameBill, ztmt)
			}
			for ztaf := range z.Bill {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Bill[ztaf] = nil
				} else {
					if z.Bill[ztaf] == nil {
						z.Bill[ztaf] = new(GameBill)
					}
					bts, err = z.Bill[ztaf].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "f":
			var ztco uint32
			ztco, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Flow) >= int(ztco) {
				z.Flow = (z.Flow)[:ztco]
			} else {
				z.Flow = make([]int32, ztco)
			}
			for zeth := range z.Flow {
				z.Flow[zeth], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "g":
			var zana uint32
			zana, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zana) {
				z.Group = (z.Group)[:zana]
			} else {
				z.Group = make([]int64, zana)
			}
			for zsbz := range z.Group {
				z.Group[zsbz], bts, err = msgp.ReadInt64Bytes(bts)
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
			var ztyy uint32
			ztyy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Odds) >= int(ztyy) {
				z.Odds = (z.Odds)[:ztyy]
			} else {
				z.Odds = make([]int32, ztyy)
			}
			for zrjx := range z.Odds {
				z.Odds[zrjx], bts, err = msgp.ReadInt32Bytes(bts)
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
			var zinl uint32
			zinl, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zinl) {
				z.Rich = (z.Rich)[:zinl]
			} else {
				z.Rich = make([]*RichRole, zinl)
			}
			for zawn := range z.Rich {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Rich[zawn] = nil
				} else {
					if z.Rich[zawn] == nil {
						z.Rich[zawn] = new(RichRole)
					}
					var zare uint32
					zare, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zare > 0 {
						zare--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "u":
							z.Rich[zawn].Uid, bts, err = msgp.ReadInt32Bytes(bts)
							if err != nil {
								return
							}
						case "n":
							z.Rich[zawn].Name, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "c":
							z.Rich[zawn].Coin, bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *FolksGameRound) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.ArrayHeaderSize
	for ztaf := range z.Bill {
		if z.Bill[ztaf] == nil {
			s += msgp.NilSize
		} else {
			s += z.Bill[ztaf].Msgsize()
		}
	}
	s += 2 + msgp.ArrayHeaderSize + (len(z.Flow) * (msgp.Int32Size)) + 2 + msgp.ArrayHeaderSize + (len(z.Group) * (msgp.Int64Size)) + 2 + msgp.BytesPrefixSize + len(z.Poker) + 2 + msgp.ArrayHeaderSize + (len(z.Odds) * (msgp.Int32Size)) + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.StringPrefixSize + len(z.Note) + 2 + msgp.ArrayHeaderSize
	for zawn := range z.Rich {
		if z.Rich[zawn] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 2 + msgp.Int32Size + 2 + msgp.StringPrefixSize + len(z.Rich[zawn].Name) + 2 + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FolksUserLog) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zixj uint32
	zixj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zixj > 0 {
		zixj--
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
			var zrsc uint32
			zrsc, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zrsc) {
				z.Group = (z.Group)[:zrsc]
			} else {
				z.Group = make([]int64, zrsc)
			}
			for zljy := range z.Group {
				z.Group[zljy], err = dc.ReadInt64()
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
func (z *FolksUserLog) EncodeMsg(en *msgp.Writer) (err error) {
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
	for zljy := range z.Group {
		err = en.WriteInt64(z.Group[zljy])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FolksUserLog) MarshalMsg(b []byte) (o []byte, err error) {
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
	for zljy := range z.Group {
		o = msgp.AppendInt64(o, z.Group[zljy])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FolksUserLog) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zctn uint32
	zctn, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zctn > 0 {
		zctn--
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
			var zswy uint32
			zswy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zswy) {
				z.Group = (z.Group)[:zswy]
			} else {
				z.Group = make([]int64, zswy)
			}
			for zljy := range z.Group {
				z.Group[zljy], bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *FolksUserLog) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int64Size + 2 + msgp.BytesPrefixSize + len(z.Poker) + 2 + msgp.ArrayHeaderSize + (len(z.Group) * (msgp.Int64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameBill) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrus uint32
	zrus, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrus > 0 {
		zrus--
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
			var zsvm uint32
			zsvm, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zsvm) {
				z.Group = (z.Group)[:zsvm]
			} else {
				z.Group = make([]int64, zsvm)
			}
			for znsg := range z.Group {
				z.Group[znsg], err = dc.ReadInt64()
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
	for znsg := range z.Group {
		err = en.WriteInt64(z.Group[znsg])
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
	for znsg := range z.Group {
		o = msgp.AppendInt64(o, z.Group[znsg])
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
	var zaoz uint32
	zaoz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zaoz > 0 {
		zaoz--
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
			var zfzb uint32
			zfzb, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zfzb) {
				z.Group = (z.Group)[:zfzb]
			} else {
				z.Group = make([]int64, zfzb)
			}
			for znsg := range z.Group {
				z.Group[znsg], bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *OpenBetAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
			var zqgz uint32
			zqgz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zqgz) {
				z.Rich = (z.Rich)[:zqgz]
			} else {
				z.Rich = make([]*User, zqgz)
			}
			for zsbo := range z.Rich {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Rich[zsbo] = nil
				} else {
					if z.Rich[zsbo] == nil {
						z.Rich[zsbo] = new(User)
					}
					err = z.Rich[zsbo].DecodeMsg(dc)
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
func (z *OpenBetAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "id"
	err = en.Append(0x83, 0xa2, 0x69, 0x64)
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
	for zsbo := range z.Rich {
		if z.Rich[zsbo] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Rich[zsbo].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *OpenBetAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "id"
	o = append(o, 0x83, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Id)
	// string "state"
	o = append(o, 0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	o = msgp.AppendInt32(o, z.State)
	// string "rich"
	o = append(o, 0xa4, 0x72, 0x69, 0x63, 0x68)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rich)))
	for zsbo := range z.Rich {
		if z.Rich[zsbo] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Rich[zsbo].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OpenBetAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
			var ztls uint32
			ztls, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(ztls) {
				z.Rich = (z.Rich)[:ztls]
			} else {
				z.Rich = make([]*User, ztls)
			}
			for zsbo := range z.Rich {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Rich[zsbo] = nil
				} else {
					if z.Rich[zsbo] == nil {
						z.Rich[zsbo] = new(User)
					}
					bts, err = z.Rich[zsbo].UnmarshalMsg(bts)
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
func (z *OpenBetAck) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 6 + msgp.Int32Size + 5 + msgp.ArrayHeaderSize
	for zsbo := range z.Rich {
		if z.Rich[zsbo] == nil {
			s += msgp.NilSize
		} else {
			s += z.Rich[zsbo].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RichRole) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zmvo uint32
	zmvo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zmvo > 0 {
		zmvo--
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
		case "n":
			z.Name, err = dc.ReadString()
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
func (z RichRole) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "u"
	err = en.Append(0x83, 0xa1, 0x75)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Uid)
	if err != nil {
		return
	}
	// write "n"
	err = en.Append(0xa1, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
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
func (z RichRole) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "u"
	o = append(o, 0x83, 0xa1, 0x75)
	o = msgp.AppendInt32(o, z.Uid)
	// string "n"
	o = append(o, 0xa1, 0x6e)
	o = msgp.AppendString(o, z.Name)
	// string "c"
	o = append(o, 0xa1, 0x63)
	o = msgp.AppendInt64(o, z.Coin)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RichRole) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zigk uint32
	zigk, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zigk > 0 {
		zigk--
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
		case "n":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
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
func (z RichRole) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int32Size + 2 + msgp.StringPrefixSize + len(z.Name) + 2 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserBetAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zuop uint32
	zuop, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zuop > 0 {
		zuop--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "item":
			var zedl uint32
			zedl, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zedl) {
				z.Bet = (z.Bet)[:zedl]
			} else {
				z.Bet = make([]int32, zedl)
			}
			for zopb := range z.Bet {
				z.Bet[zopb], err = dc.ReadInt32()
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
	for zopb := range z.Bet {
		err = en.WriteInt32(z.Bet[zopb])
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
	for zopb := range z.Bet {
		o = msgp.AppendInt32(o, z.Bet[zopb])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserBetAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zupd uint32
	zupd, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zupd > 0 {
		zupd--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "item":
			var zome uint32
			zome, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zome) {
				z.Bet = (z.Bet)[:zome]
			} else {
				z.Bet = make([]int32, zome)
			}
			for zopb := range z.Bet {
				z.Bet[zopb], bts, err = msgp.ReadInt32Bytes(bts)
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
