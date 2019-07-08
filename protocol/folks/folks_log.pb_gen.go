package folks

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *GameBill) DecodeMsg(dc *msgp.Reader) (err error) {
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
			var zbai uint32
			zbai, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zbai) {
				z.Group = (z.Group)[:zbai]
			} else {
				z.Group = make([]int64, zbai)
			}
			for zxvk := range z.Group {
				z.Group[zxvk], err = dc.ReadInt64()
				if err != nil {
					return
				}
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
	for zxvk := range z.Group {
		err = en.WriteInt64(z.Group[zxvk])
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
	// write "x"
	err = en.Append(0xa1, 0x78)
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
	for zxvk := range z.Group {
		o = msgp.AppendInt64(o, z.Group[zxvk])
	}
	// string "w"
	o = append(o, 0xa1, 0x77)
	o = msgp.AppendInt64(o, z.Win)
	// string "x"
	o = append(o, 0xa1, 0x78)
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
			var zajw uint32
			zajw, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zajw) {
				z.Group = (z.Group)[:zajw]
			} else {
				z.Group = make([]int64, zajw)
			}
			for zxvk := range z.Group {
				z.Group[zxvk], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
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
func (z *GameRound) DecodeMsg(dc *msgp.Reader) (err error) {
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
			var zjfb uint32
			zjfb, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(zjfb) {
				z.Bill = (z.Bill)[:zjfb]
			} else {
				z.Bill = make([]*GameBill, zjfb)
			}
			for zwht := range z.Bill {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Bill[zwht] = nil
				} else {
					if z.Bill[zwht] == nil {
						z.Bill[zwht] = new(GameBill)
					}
					err = z.Bill[zwht].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "f":
			var zcxo uint32
			zcxo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Flow) >= int(zcxo) {
				z.Flow = (z.Flow)[:zcxo]
			} else {
				z.Flow = make([]int32, zcxo)
			}
			for zhct := range z.Flow {
				z.Flow[zhct], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "g":
			var zeff uint32
			zeff, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zeff) {
				z.Group = (z.Group)[:zeff]
			} else {
				z.Group = make([]int64, zeff)
			}
			for zcua := range z.Group {
				z.Group[zcua], err = dc.ReadInt64()
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
			var zrsw uint32
			zrsw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Odds) >= int(zrsw) {
				z.Odds = (z.Odds)[:zrsw]
			} else {
				z.Odds = make([]int32, zrsw)
			}
			for zxhx := range z.Odds {
				z.Odds[zxhx], err = dc.ReadInt32()
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
			var zxpk uint32
			zxpk, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zxpk) {
				z.Rich = (z.Rich)[:zxpk]
			} else {
				z.Rich = make([]int32, zxpk)
			}
			for zlqf := range z.Rich {
				z.Rich[zlqf], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "u":
			var zdnj uint32
			zdnj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.UserBet) >= int(zdnj) {
				z.UserBet = (z.UserBet)[:zdnj]
			} else {
				z.UserBet = make([]int64, zdnj)
			}
			for zdaf := range z.UserBet {
				z.UserBet[zdaf], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "k":
			z.Bank, err = dc.ReadInt32()
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
	for zwht := range z.Bill {
		if z.Bill[zwht] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Bill[zwht].EncodeMsg(en)
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
	for zhct := range z.Flow {
		err = en.WriteInt32(z.Flow[zhct])
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
	for zcua := range z.Group {
		err = en.WriteInt64(z.Group[zcua])
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
	for zxhx := range z.Odds {
		err = en.WriteInt32(z.Odds[zxhx])
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
	for zlqf := range z.Rich {
		err = en.WriteInt32(z.Rich[zlqf])
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
	for zdaf := range z.UserBet {
		err = en.WriteInt64(z.UserBet[zdaf])
		if err != nil {
			return
		}
	}
	// write "k"
	err = en.Append(0xa1, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Bank)
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
	for zwht := range z.Bill {
		if z.Bill[zwht] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Bill[zwht].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "f"
	o = append(o, 0xa1, 0x66)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Flow)))
	for zhct := range z.Flow {
		o = msgp.AppendInt32(o, z.Flow[zhct])
	}
	// string "g"
	o = append(o, 0xa1, 0x67)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Group)))
	for zcua := range z.Group {
		o = msgp.AppendInt64(o, z.Group[zcua])
	}
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendBytes(o, z.Poker)
	// string "o"
	o = append(o, 0xa1, 0x6f)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Odds)))
	for zxhx := range z.Odds {
		o = msgp.AppendInt32(o, z.Odds[zxhx])
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
	for zlqf := range z.Rich {
		o = msgp.AppendInt32(o, z.Rich[zlqf])
	}
	// string "u"
	o = append(o, 0xa1, 0x75)
	o = msgp.AppendArrayHeader(o, uint32(len(z.UserBet)))
	for zdaf := range z.UserBet {
		o = msgp.AppendInt64(o, z.UserBet[zdaf])
	}
	// string "k"
	o = append(o, 0xa1, 0x6b)
	o = msgp.AppendInt32(o, z.Bank)
	// string "v"
	o = append(o, 0xa1, 0x76)
	o = msgp.AppendBool(o, z.Cheat)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameRound) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
			var zsnv uint32
			zsnv, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(zsnv) {
				z.Bill = (z.Bill)[:zsnv]
			} else {
				z.Bill = make([]*GameBill, zsnv)
			}
			for zwht := range z.Bill {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Bill[zwht] = nil
				} else {
					if z.Bill[zwht] == nil {
						z.Bill[zwht] = new(GameBill)
					}
					bts, err = z.Bill[zwht].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "f":
			var zkgt uint32
			zkgt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Flow) >= int(zkgt) {
				z.Flow = (z.Flow)[:zkgt]
			} else {
				z.Flow = make([]int32, zkgt)
			}
			for zhct := range z.Flow {
				z.Flow[zhct], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "g":
			var zema uint32
			zema, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Group) >= int(zema) {
				z.Group = (z.Group)[:zema]
			} else {
				z.Group = make([]int64, zema)
			}
			for zcua := range z.Group {
				z.Group[zcua], bts, err = msgp.ReadInt64Bytes(bts)
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
			var zpez uint32
			zpez, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Odds) >= int(zpez) {
				z.Odds = (z.Odds)[:zpez]
			} else {
				z.Odds = make([]int32, zpez)
			}
			for zxhx := range z.Odds {
				z.Odds[zxhx], bts, err = msgp.ReadInt32Bytes(bts)
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
			var zqke uint32
			zqke, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rich) >= int(zqke) {
				z.Rich = (z.Rich)[:zqke]
			} else {
				z.Rich = make([]int32, zqke)
			}
			for zlqf := range z.Rich {
				z.Rich[zlqf], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "u":
			var zqyh uint32
			zqyh, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.UserBet) >= int(zqyh) {
				z.UserBet = (z.UserBet)[:zqyh]
			} else {
				z.UserBet = make([]int64, zqyh)
			}
			for zdaf := range z.UserBet {
				z.UserBet[zdaf], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "k":
			z.Bank, bts, err = msgp.ReadInt32Bytes(bts)
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
	for zwht := range z.Bill {
		if z.Bill[zwht] == nil {
			s += msgp.NilSize
		} else {
			s += z.Bill[zwht].Msgsize()
		}
	}
	s += 2 + msgp.ArrayHeaderSize + (len(z.Flow) * (msgp.Int32Size)) + 2 + msgp.ArrayHeaderSize + (len(z.Group) * (msgp.Int64Size)) + 2 + msgp.BytesPrefixSize + len(z.Poker) + 2 + msgp.ArrayHeaderSize + (len(z.Odds) * (msgp.Int32Size)) + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.StringPrefixSize + len(z.Note) + 2 + msgp.ArrayHeaderSize + (len(z.Rich) * (msgp.Int32Size)) + 2 + msgp.ArrayHeaderSize + (len(z.UserBet) * (msgp.Int64Size)) + 2 + msgp.Int32Size + 2 + msgp.BoolSize
	return
}
