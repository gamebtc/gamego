package protocol

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
	var zlqf uint32
	zlqf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zlqf > 0 {
		zlqf--
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
			var zdaf uint32
			zdaf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zdaf) {
				z.Rich = (z.Rich)[:zdaf]
			} else {
				z.Rich = make([]*User, zdaf)
			}
			for zhct := range z.Rich {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Rich[zhct] = nil
				} else {
					if z.Rich[zhct] == nil {
						z.Rich[zhct] = new(User)
					}
					err = z.Rich[zhct].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "sum":
			var zpks uint32
			zpks, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Sum) >= int(zpks) {
				z.Sum = (z.Sum)[:zpks]
			} else {
				z.Sum = make([]int64, zpks)
			}
			for zcua := range z.Sum {
				z.Sum[zcua], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "bet":
			var zjfb uint32
			zjfb, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zjfb) {
				z.Bet = (z.Bet)[:zjfb]
			} else {
				z.Bet = make([]int64, zjfb)
			}
			for zxhx := range z.Bet {
				z.Bet[zxhx], err = dc.ReadInt64()
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
	for zhct := range z.Rich {
		if z.Rich[zhct] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Rich[zhct].EncodeMsg(en)
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
	for zcua := range z.Sum {
		err = en.WriteInt64(z.Sum[zcua])
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
	for zxhx := range z.Bet {
		err = en.WriteInt64(z.Bet[zxhx])
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
	for zhct := range z.Rich {
		if z.Rich[zhct] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Rich[zhct].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "sum"
	o = append(o, 0xa3, 0x73, 0x75, 0x6d)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Sum)))
	for zcua := range z.Sum {
		o = msgp.AppendInt64(o, z.Sum[zcua])
	}
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bet)))
	for zxhx := range z.Bet {
		o = msgp.AppendInt64(o, z.Bet[zxhx])
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
	var zcxo uint32
	zcxo, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcxo > 0 {
		zcxo--
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
			var zeff uint32
			zeff, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zeff) {
				z.Rich = (z.Rich)[:zeff]
			} else {
				z.Rich = make([]*User, zeff)
			}
			for zhct := range z.Rich {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Rich[zhct] = nil
				} else {
					if z.Rich[zhct] == nil {
						z.Rich[zhct] = new(User)
					}
					bts, err = z.Rich[zhct].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "sum":
			var zrsw uint32
			zrsw, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Sum) >= int(zrsw) {
				z.Sum = (z.Sum)[:zrsw]
			} else {
				z.Sum = make([]int64, zrsw)
			}
			for zcua := range z.Sum {
				z.Sum[zcua], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "bet":
			var zxpk uint32
			zxpk, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zxpk) {
				z.Bet = (z.Bet)[:zxpk]
			} else {
				z.Bet = make([]int64, zxpk)
			}
			for zxhx := range z.Bet {
				z.Bet[zxhx], bts, err = msgp.ReadInt64Bytes(bts)
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
	for zhct := range z.Rich {
		if z.Rich[zhct] == nil {
			s += msgp.NilSize
		} else {
			s += z.Rich[zhct].Msgsize()
		}
	}
	s += 4 + msgp.ArrayHeaderSize + (len(z.Sum) * (msgp.Int64Size)) + 4 + msgp.ArrayHeaderSize + (len(z.Bet) * (msgp.Int64Size)) + 4 + msgp.BytesPrefixSize + len(z.Log)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FolksGameResult) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zkgt uint32
	zkgt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zkgt > 0 {
		zkgt--
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
			var zema uint32
			zema, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Odd) >= int(zema) {
				z.Odd = (z.Odd)[:zema]
			} else {
				z.Odd = make([]int32, zema)
			}
			for zdnj := range z.Odd {
				z.Odd[zdnj], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "bet":
			var zpez uint32
			zpez, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zpez) {
				z.Bet = (z.Bet)[:zpez]
			} else {
				z.Bet = make([]int64, zpez)
			}
			for zobc := range z.Bet {
				z.Bet[zobc], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "rich":
			var zqke uint32
			zqke, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zqke) {
				z.Rich = (z.Rich)[:zqke]
			} else {
				z.Rich = make([]int64, zqke)
			}
			for zsnv := range z.Rich {
				z.Rich[zsnv], err = dc.ReadInt64()
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
	for zdnj := range z.Odd {
		err = en.WriteInt32(z.Odd[zdnj])
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
	for zobc := range z.Bet {
		err = en.WriteInt64(z.Bet[zobc])
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
	for zsnv := range z.Rich {
		err = en.WriteInt64(z.Rich[zsnv])
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
	for zdnj := range z.Odd {
		o = msgp.AppendInt32(o, z.Odd[zdnj])
	}
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bet)))
	for zobc := range z.Bet {
		o = msgp.AppendInt64(o, z.Bet[zobc])
	}
	// string "rich"
	o = append(o, 0xa4, 0x72, 0x69, 0x63, 0x68)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rich)))
	for zsnv := range z.Rich {
		o = msgp.AppendInt64(o, z.Rich[zsnv])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FolksGameResult) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zqyh uint32
	zqyh, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zqyh > 0 {
		zqyh--
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
			var zyzr uint32
			zyzr, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Odd) >= int(zyzr) {
				z.Odd = (z.Odd)[:zyzr]
			} else {
				z.Odd = make([]int32, zyzr)
			}
			for zdnj := range z.Odd {
				z.Odd[zdnj], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "bet":
			var zywj uint32
			zywj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zywj) {
				z.Bet = (z.Bet)[:zywj]
			} else {
				z.Bet = make([]int64, zywj)
			}
			for zobc := range z.Bet {
				z.Bet[zobc], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "rich":
			var zjpj uint32
			zjpj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zjpj) {
				z.Rich = (z.Rich)[:zjpj]
			} else {
				z.Rich = make([]int64, zjpj)
			}
			for zsnv := range z.Rich {
				z.Rich[zsnv], bts, err = msgp.ReadInt64Bytes(bts)
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
	var zrjx uint32
	zrjx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrjx > 0 {
		zrjx--
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
			var zawn uint32
			zawn, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(zawn) {
				z.Bill = (z.Bill)[:zawn]
			} else {
				z.Bill = make([]*GameBill, zawn)
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
		case "f":
			var zwel uint32
			zwel, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Flow) >= int(zwel) {
				z.Flow = (z.Flow)[:zwel]
			} else {
				z.Flow = make([]int32, zwel)
			}
			for zrfe := range z.Flow {
				z.Flow[zrfe], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "g":
			var zrbe uint32
			zrbe, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zrbe) {
				z.Group = (z.Group)[:zrbe]
			} else {
				z.Group = make([]int64, zrbe)
			}
			for zgmo := range z.Group {
				z.Group[zgmo], err = dc.ReadInt64()
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
			var zmfd uint32
			zmfd, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Odds) >= int(zmfd) {
				z.Odds = (z.Odds)[:zmfd]
			} else {
				z.Odds = make([]int32, zmfd)
			}
			for ztaf := range z.Odds {
				z.Odds[ztaf], err = dc.ReadInt32()
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
			var zzdc uint32
			zzdc, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zzdc) {
				z.Rich = (z.Rich)[:zzdc]
			} else {
				z.Rich = make([]int32, zzdc)
			}
			for zeth := range z.Rich {
				z.Rich[zeth], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "u":
			var zelx uint32
			zelx, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.BetGroup) >= int(zelx) {
				z.BetGroup = (z.BetGroup)[:zelx]
			} else {
				z.BetGroup = make([]int64, zelx)
			}
			for zsbz := range z.BetGroup {
				z.BetGroup[zsbz], err = dc.ReadInt64()
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
func (z *FolksGameRound) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 15
	// write "_id"
	err = en.Append(0x8f, 0xa3, 0x5f, 0x69, 0x64)
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
	// write "f"
	err = en.Append(0xa1, 0x66)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Flow)))
	if err != nil {
		return
	}
	for zrfe := range z.Flow {
		err = en.WriteInt32(z.Flow[zrfe])
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
	for zgmo := range z.Group {
		err = en.WriteInt64(z.Group[zgmo])
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
	for ztaf := range z.Odds {
		err = en.WriteInt32(z.Odds[ztaf])
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
	for zeth := range z.Rich {
		err = en.WriteInt32(z.Rich[zeth])
		if err != nil {
			return
		}
	}
	// write "u"
	err = en.Append(0xa1, 0x75)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.BetGroup)))
	if err != nil {
		return
	}
	for zsbz := range z.BetGroup {
		err = en.WriteInt64(z.BetGroup[zsbz])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FolksGameRound) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 15
	// string "_id"
	o = append(o, 0x8f, 0xa3, 0x5f, 0x69, 0x64)
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
	// string "f"
	o = append(o, 0xa1, 0x66)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Flow)))
	for zrfe := range z.Flow {
		o = msgp.AppendInt32(o, z.Flow[zrfe])
	}
	// string "g"
	o = append(o, 0xa1, 0x67)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Group)))
	for zgmo := range z.Group {
		o = msgp.AppendInt64(o, z.Group[zgmo])
	}
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendBytes(o, z.Poker)
	// string "o"
	o = append(o, 0xa1, 0x6f)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Odds)))
	for ztaf := range z.Odds {
		o = msgp.AppendInt32(o, z.Odds[ztaf])
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
	for zeth := range z.Rich {
		o = msgp.AppendInt32(o, z.Rich[zeth])
	}
	// string "u"
	o = append(o, 0xa1, 0x75)
	o = msgp.AppendArrayHeader(o, uint32(len(z.BetGroup)))
	for zsbz := range z.BetGroup {
		o = msgp.AppendInt64(o, z.BetGroup[zsbz])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FolksGameRound) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbal uint32
	zbal, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbal > 0 {
		zbal--
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
			var zjqz uint32
			zjqz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(zjqz) {
				z.Bill = (z.Bill)[:zjqz]
			} else {
				z.Bill = make([]*GameBill, zjqz)
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
		case "f":
			var zkct uint32
			zkct, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Flow) >= int(zkct) {
				z.Flow = (z.Flow)[:zkct]
			} else {
				z.Flow = make([]int32, zkct)
			}
			for zrfe := range z.Flow {
				z.Flow[zrfe], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "g":
			var ztmt uint32
			ztmt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Group) >= int(ztmt) {
				z.Group = (z.Group)[:ztmt]
			} else {
				z.Group = make([]int64, ztmt)
			}
			for zgmo := range z.Group {
				z.Group[zgmo], bts, err = msgp.ReadInt64Bytes(bts)
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
			var ztco uint32
			ztco, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Odds) >= int(ztco) {
				z.Odds = (z.Odds)[:ztco]
			} else {
				z.Odds = make([]int32, ztco)
			}
			for ztaf := range z.Odds {
				z.Odds[ztaf], bts, err = msgp.ReadInt32Bytes(bts)
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
			var zana uint32
			zana, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zana) {
				z.Rich = (z.Rich)[:zana]
			} else {
				z.Rich = make([]int32, zana)
			}
			for zeth := range z.Rich {
				z.Rich[zeth], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "u":
			var ztyy uint32
			ztyy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.BetGroup) >= int(ztyy) {
				z.BetGroup = (z.BetGroup)[:ztyy]
			} else {
				z.BetGroup = make([]int64, ztyy)
			}
			for zsbz := range z.BetGroup {
				z.BetGroup[zsbz], bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *FolksGameRound) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.ArrayHeaderSize
	for zzpf := range z.Bill {
		if z.Bill[zzpf] == nil {
			s += msgp.NilSize
		} else {
			s += z.Bill[zzpf].Msgsize()
		}
	}
	s += 2 + msgp.ArrayHeaderSize + (len(z.Flow) * (msgp.Int32Size)) + 2 + msgp.ArrayHeaderSize + (len(z.Group) * (msgp.Int64Size)) + 2 + msgp.BytesPrefixSize + len(z.Poker) + 2 + msgp.ArrayHeaderSize + (len(z.Odds) * (msgp.Int32Size)) + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.StringPrefixSize + len(z.Note) + 2 + msgp.ArrayHeaderSize + (len(z.Rich) * (msgp.Int32Size)) + 2 + msgp.ArrayHeaderSize + (len(z.BetGroup) * (msgp.Int64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FolksUserLog) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zare uint32
	zare, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zare > 0 {
		zare--
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
			var zljy uint32
			zljy, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zljy) {
				z.Group = (z.Group)[:zljy]
			} else {
				z.Group = make([]int64, zljy)
			}
			for zinl := range z.Group {
				z.Group[zinl], err = dc.ReadInt64()
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
	for zinl := range z.Group {
		err = en.WriteInt64(z.Group[zinl])
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
	for zinl := range z.Group {
		o = msgp.AppendInt64(o, z.Group[zinl])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FolksUserLog) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zixj uint32
	zixj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zixj > 0 {
		zixj--
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
			var zrsc uint32
			zrsc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zrsc) {
				z.Group = (z.Group)[:zrsc]
			} else {
				z.Group = make([]int64, zrsc)
			}
			for zinl := range z.Group {
				z.Group[zinl], bts, err = msgp.ReadInt64Bytes(bts)
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
			var znsg uint32
			znsg, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Group) >= int(znsg) {
				z.Group = (z.Group)[:znsg]
			} else {
				z.Group = make([]int64, znsg)
			}
			for zctn := range z.Group {
				z.Group[zctn], err = dc.ReadInt64()
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
	for zctn := range z.Group {
		err = en.WriteInt64(z.Group[zctn])
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
	for zctn := range z.Group {
		o = msgp.AppendInt64(o, z.Group[zctn])
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
			var zsvm uint32
			zsvm, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zsvm) {
				z.Group = (z.Group)[:zsvm]
			} else {
				z.Group = make([]int64, zsvm)
			}
			for zctn := range z.Group {
				z.Group[zctn], bts, err = msgp.ReadInt64Bytes(bts)
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
	var zfzb uint32
	zfzb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zfzb > 0 {
		zfzb--
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
			var zsbo uint32
			zsbo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zsbo) {
				z.Rich = (z.Rich)[:zsbo]
			} else {
				z.Rich = make([]*User, zsbo)
			}
			for zaoz := range z.Rich {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Rich[zaoz] = nil
				} else {
					if z.Rich[zaoz] == nil {
						z.Rich[zaoz] = new(User)
					}
					err = z.Rich[zaoz].DecodeMsg(dc)
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
	for zaoz := range z.Rich {
		if z.Rich[zaoz] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Rich[zaoz].EncodeMsg(en)
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
	for zaoz := range z.Rich {
		if z.Rich[zaoz] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Rich[zaoz].MarshalMsg(o)
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
	var zjif uint32
	zjif, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjif > 0 {
		zjif--
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
			var zqgz uint32
			zqgz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zqgz) {
				z.Rich = (z.Rich)[:zqgz]
			} else {
				z.Rich = make([]*User, zqgz)
			}
			for zaoz := range z.Rich {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Rich[zaoz] = nil
				} else {
					if z.Rich[zaoz] == nil {
						z.Rich[zaoz] = new(User)
					}
					bts, err = z.Rich[zaoz].UnmarshalMsg(bts)
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
	for zaoz := range z.Rich {
		if z.Rich[zaoz] == nil {
			s += msgp.NilSize
		} else {
			s += z.Rich[zaoz].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserBetAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var ztls uint32
	ztls, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ztls > 0 {
		ztls--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "item":
			var zmvo uint32
			zmvo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zmvo) {
				z.Bet = (z.Bet)[:zmvo]
			} else {
				z.Bet = make([]int32, zmvo)
			}
			for zsnw := range z.Bet {
				z.Bet[zsnw], err = dc.ReadInt32()
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
	for zsnw := range z.Bet {
		err = en.WriteInt32(z.Bet[zsnw])
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
	for zsnw := range z.Bet {
		o = msgp.AppendInt32(o, z.Bet[zsnw])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserBetAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "item":
			var zopb uint32
			zopb, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bet) >= int(zopb) {
				z.Bet = (z.Bet)[:zopb]
			} else {
				z.Bet = make([]int32, zopb)
			}
			for zsnw := range z.Bet {
				z.Bet[zsnw], bts, err = msgp.ReadInt32Bytes(bts)
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
