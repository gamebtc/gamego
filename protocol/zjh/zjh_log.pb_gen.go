package zjh

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *ActionLog) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Start":
			z.Start, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Type":
			err = z.Type.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "Uid":
			z.Uid, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Bet":
			z.Bet, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Players":
			var zcmr uint32
			zcmr, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zcmr) {
				z.Players = (z.Players)[:zcmr]
			} else {
				z.Players = make([]int32, zcmr)
			}
			for zxvk := range z.Players {
				z.Players[zxvk], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "Winners":
			var zajw uint32
			zajw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Winners) >= int(zajw) {
				z.Winners = (z.Winners)[:zajw]
			} else {
				z.Winners = make([]int32, zajw)
			}
			for zbzg := range z.Winners {
				z.Winners[zbzg], err = dc.ReadInt32()
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
func (z *ActionLog) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "Start"
	err = en.Append(0x86, 0xa5, 0x53, 0x74, 0x61, 0x72, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Start)
	if err != nil {
		return
	}
	// write "Type"
	err = en.Append(0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = z.Type.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "Uid"
	err = en.Append(0xa3, 0x55, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Uid)
	if err != nil {
		return
	}
	// write "Bet"
	err = en.Append(0xa3, 0x42, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Bet)
	if err != nil {
		return
	}
	// write "Players"
	err = en.Append(0xa7, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Players)))
	if err != nil {
		return
	}
	for zxvk := range z.Players {
		err = en.WriteInt32(z.Players[zxvk])
		if err != nil {
			return
		}
	}
	// write "Winners"
	err = en.Append(0xa7, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Winners)))
	if err != nil {
		return
	}
	for zbzg := range z.Winners {
		err = en.WriteInt32(z.Winners[zbzg])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ActionLog) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Start"
	o = append(o, 0x86, 0xa5, 0x53, 0x74, 0x61, 0x72, 0x74)
	o = msgp.AppendInt64(o, z.Start)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o, err = z.Type.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "Uid"
	o = append(o, 0xa3, 0x55, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Uid)
	// string "Bet"
	o = append(o, 0xa3, 0x42, 0x65, 0x74)
	o = msgp.AppendInt32(o, z.Bet)
	// string "Players"
	o = append(o, 0xa7, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Players)))
	for zxvk := range z.Players {
		o = msgp.AppendInt32(o, z.Players[zxvk])
	}
	// string "Winners"
	o = append(o, 0xa7, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Winners)))
	for zbzg := range z.Winners {
		o = msgp.AppendInt32(o, z.Winners[zbzg])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ActionLog) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Start":
			z.Start, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Type":
			bts, err = z.Type.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "Uid":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Bet":
			z.Bet, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Players":
			var zhct uint32
			zhct, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zhct) {
				z.Players = (z.Players)[:zhct]
			} else {
				z.Players = make([]int32, zhct)
			}
			for zxvk := range z.Players {
				z.Players[zxvk], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "Winners":
			var zcua uint32
			zcua, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Winners) >= int(zcua) {
				z.Winners = (z.Winners)[:zcua]
			} else {
				z.Winners = make([]int32, zcua)
			}
			for zbzg := range z.Winners {
				z.Winners[zbzg], bts, err = msgp.ReadInt32Bytes(bts)
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
func (z *ActionLog) Msgsize() (s int) {
	s = 1 + 6 + msgp.Int64Size + 5 + z.Type.Msgsize() + 4 + msgp.Int32Size + 4 + msgp.Int32Size + 8 + msgp.ArrayHeaderSize + (len(z.Players) * (msgp.Int32Size)) + 8 + msgp.ArrayHeaderSize + (len(z.Winners) * (msgp.Int32Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameBill) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "u":
			z.Uid, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Job":
			z.Job, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "OldCoin":
			z.OldCoin, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Bet":
			z.Bet, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Poker":
			z.Poker, err = dc.ReadBytes(z.Poker)
			if err != nil {
				return
			}
		case "Weight":
			z.Weight, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Win":
			z.Win, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Tax":
			z.Tax, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Water":
			z.Water, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Lucky":
			z.Lucky, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Robot":
			z.Robot, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Pk":
			var zdaf uint32
			zdaf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Pk) >= int(zdaf) {
				z.Pk = (z.Pk)[:zdaf]
			} else {
				z.Pk = make([]int32, zdaf)
			}
			for zxhx := range z.Pk {
				z.Pk[zxhx], err = dc.ReadInt32()
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
	// map header, size 12
	// write "u"
	err = en.Append(0x8c, 0xa1, 0x75)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Uid)
	if err != nil {
		return
	}
	// write "Job"
	err = en.Append(0xa3, 0x4a, 0x6f, 0x62)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Job)
	if err != nil {
		return
	}
	// write "OldCoin"
	err = en.Append(0xa7, 0x4f, 0x6c, 0x64, 0x43, 0x6f, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.OldCoin)
	if err != nil {
		return
	}
	// write "Bet"
	err = en.Append(0xa3, 0x42, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Bet)
	if err != nil {
		return
	}
	// write "Poker"
	err = en.Append(0xa5, 0x50, 0x6f, 0x6b, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Poker)
	if err != nil {
		return
	}
	// write "Weight"
	err = en.Append(0xa6, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Weight)
	if err != nil {
		return
	}
	// write "Win"
	err = en.Append(0xa3, 0x57, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Win)
	if err != nil {
		return
	}
	// write "Tax"
	err = en.Append(0xa3, 0x54, 0x61, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Tax)
	if err != nil {
		return
	}
	// write "Water"
	err = en.Append(0xa5, 0x57, 0x61, 0x74, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Water)
	if err != nil {
		return
	}
	// write "Lucky"
	err = en.Append(0xa5, 0x4c, 0x75, 0x63, 0x6b, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Lucky)
	if err != nil {
		return
	}
	// write "Robot"
	err = en.Append(0xa5, 0x52, 0x6f, 0x62, 0x6f, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Robot)
	if err != nil {
		return
	}
	// write "Pk"
	err = en.Append(0xa2, 0x50, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Pk)))
	if err != nil {
		return
	}
	for zxhx := range z.Pk {
		err = en.WriteInt32(z.Pk[zxhx])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameBill) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 12
	// string "u"
	o = append(o, 0x8c, 0xa1, 0x75)
	o = msgp.AppendInt32(o, z.Uid)
	// string "Job"
	o = append(o, 0xa3, 0x4a, 0x6f, 0x62)
	o = msgp.AppendInt32(o, z.Job)
	// string "OldCoin"
	o = append(o, 0xa7, 0x4f, 0x6c, 0x64, 0x43, 0x6f, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.OldCoin)
	// string "Bet"
	o = append(o, 0xa3, 0x42, 0x65, 0x74)
	o = msgp.AppendInt64(o, z.Bet)
	// string "Poker"
	o = append(o, 0xa5, 0x50, 0x6f, 0x6b, 0x65, 0x72)
	o = msgp.AppendBytes(o, z.Poker)
	// string "Weight"
	o = append(o, 0xa6, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendInt32(o, z.Weight)
	// string "Win"
	o = append(o, 0xa3, 0x57, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.Win)
	// string "Tax"
	o = append(o, 0xa3, 0x54, 0x61, 0x78)
	o = msgp.AppendInt64(o, z.Tax)
	// string "Water"
	o = append(o, 0xa5, 0x57, 0x61, 0x74, 0x65, 0x72)
	o = msgp.AppendInt64(o, z.Water)
	// string "Lucky"
	o = append(o, 0xa5, 0x4c, 0x75, 0x63, 0x6b, 0x79)
	o = msgp.AppendInt64(o, z.Lucky)
	// string "Robot"
	o = append(o, 0xa5, 0x52, 0x6f, 0x62, 0x6f, 0x74)
	o = msgp.AppendInt64(o, z.Robot)
	// string "Pk"
	o = append(o, 0xa2, 0x50, 0x6b)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Pk)))
	for zxhx := range z.Pk {
		o = msgp.AppendInt32(o, z.Pk[zxhx])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameBill) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zpks uint32
	zpks, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zpks > 0 {
		zpks--
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
		case "Job":
			z.Job, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "OldCoin":
			z.OldCoin, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Bet":
			z.Bet, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Poker":
			z.Poker, bts, err = msgp.ReadBytesBytes(bts, z.Poker)
			if err != nil {
				return
			}
		case "Weight":
			z.Weight, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Win":
			z.Win, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Tax":
			z.Tax, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Water":
			z.Water, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Lucky":
			z.Lucky, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Robot":
			z.Robot, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Pk":
			var zjfb uint32
			zjfb, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Pk) >= int(zjfb) {
				z.Pk = (z.Pk)[:zjfb]
			} else {
				z.Pk = make([]int32, zjfb)
			}
			for zxhx := range z.Pk {
				z.Pk[zxhx], bts, err = msgp.ReadInt32Bytes(bts)
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
	s = 1 + 2 + msgp.Int32Size + 4 + msgp.Int32Size + 8 + msgp.Int64Size + 4 + msgp.Int64Size + 6 + msgp.BytesPrefixSize + len(z.Poker) + 7 + msgp.Int32Size + 4 + msgp.Int64Size + 4 + msgp.Int64Size + 6 + msgp.Int64Size + 6 + msgp.Int64Size + 6 + msgp.Int64Size + 3 + msgp.ArrayHeaderSize + (len(z.Pk) * (msgp.Int32Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameRound) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "End":
			z.End, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Room":
			z.Room, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Tab":
			z.Tab, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "b":
			var zxpk uint32
			zxpk, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(zxpk) {
				z.Bill = (z.Bill)[:zxpk]
			} else {
				z.Bill = make([]*GameBill, zxpk)
			}
			for zcxo := range z.Bill {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Bill[zcxo] = nil
				} else {
					if z.Bill[zcxo] == nil {
						z.Bill[zcxo] = new(GameBill)
					}
					err = z.Bill[zcxo].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "Ante":
			z.Ante, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Ring":
			z.Ring, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Sum":
			z.Sum, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Win":
			z.Win, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Tax":
			z.Tax, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Water":
			z.Water, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Log":
			var zdnj uint32
			zdnj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Log) >= int(zdnj) {
				z.Log = (z.Log)[:zdnj]
			} else {
				z.Log = make([]*ActionLog, zdnj)
			}
			for zeff := range z.Log {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Log[zeff] = nil
				} else {
					if z.Log[zeff] == nil {
						z.Log[zeff] = new(ActionLog)
					}
					err = z.Log[zeff].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "Pool":
			z.Pool, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Lucky":
			z.Lucky, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Note":
			z.Note, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Cheat":
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
	// write "End"
	err = en.Append(0xa3, 0x45, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.End)
	if err != nil {
		return
	}
	// write "Room"
	err = en.Append(0xa4, 0x52, 0x6f, 0x6f, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Room)
	if err != nil {
		return
	}
	// write "Tab"
	err = en.Append(0xa3, 0x54, 0x61, 0x62)
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
	for zcxo := range z.Bill {
		if z.Bill[zcxo] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Bill[zcxo].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "Ante"
	err = en.Append(0xa4, 0x41, 0x6e, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Ante)
	if err != nil {
		return
	}
	// write "Ring"
	err = en.Append(0xa4, 0x52, 0x69, 0x6e, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Ring)
	if err != nil {
		return
	}
	// write "Sum"
	err = en.Append(0xa3, 0x53, 0x75, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Sum)
	if err != nil {
		return
	}
	// write "Win"
	err = en.Append(0xa3, 0x57, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Win)
	if err != nil {
		return
	}
	// write "Tax"
	err = en.Append(0xa3, 0x54, 0x61, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Tax)
	if err != nil {
		return
	}
	// write "Water"
	err = en.Append(0xa5, 0x57, 0x61, 0x74, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Water)
	if err != nil {
		return
	}
	// write "Log"
	err = en.Append(0xa3, 0x4c, 0x6f, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Log)))
	if err != nil {
		return
	}
	for zeff := range z.Log {
		if z.Log[zeff] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Log[zeff].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "Pool"
	err = en.Append(0xa4, 0x50, 0x6f, 0x6f, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Pool)
	if err != nil {
		return
	}
	// write "Lucky"
	err = en.Append(0xa5, 0x4c, 0x75, 0x63, 0x6b, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Lucky)
	if err != nil {
		return
	}
	// write "Note"
	err = en.Append(0xa4, 0x4e, 0x6f, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Note)
	if err != nil {
		return
	}
	// write "Cheat"
	err = en.Append(0xa5, 0x43, 0x68, 0x65, 0x61, 0x74)
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
	// string "End"
	o = append(o, 0xa3, 0x45, 0x6e, 0x64)
	o = msgp.AppendInt64(o, z.End)
	// string "Room"
	o = append(o, 0xa4, 0x52, 0x6f, 0x6f, 0x6d)
	o = msgp.AppendInt32(o, z.Room)
	// string "Tab"
	o = append(o, 0xa3, 0x54, 0x61, 0x62)
	o = msgp.AppendInt32(o, z.Tab)
	// string "b"
	o = append(o, 0xa1, 0x62)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bill)))
	for zcxo := range z.Bill {
		if z.Bill[zcxo] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Bill[zcxo].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Ante"
	o = append(o, 0xa4, 0x41, 0x6e, 0x74, 0x65)
	o = msgp.AppendInt32(o, z.Ante)
	// string "Ring"
	o = append(o, 0xa4, 0x52, 0x69, 0x6e, 0x67)
	o = msgp.AppendInt32(o, z.Ring)
	// string "Sum"
	o = append(o, 0xa3, 0x53, 0x75, 0x6d)
	o = msgp.AppendInt64(o, z.Sum)
	// string "Win"
	o = append(o, 0xa3, 0x57, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.Win)
	// string "Tax"
	o = append(o, 0xa3, 0x54, 0x61, 0x78)
	o = msgp.AppendInt64(o, z.Tax)
	// string "Water"
	o = append(o, 0xa5, 0x57, 0x61, 0x74, 0x65, 0x72)
	o = msgp.AppendInt64(o, z.Water)
	// string "Log"
	o = append(o, 0xa3, 0x4c, 0x6f, 0x67)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Log)))
	for zeff := range z.Log {
		if z.Log[zeff] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Log[zeff].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Pool"
	o = append(o, 0xa4, 0x50, 0x6f, 0x6f, 0x6c)
	o = msgp.AppendInt64(o, z.Pool)
	// string "Lucky"
	o = append(o, 0xa5, 0x4c, 0x75, 0x63, 0x6b, 0x79)
	o = msgp.AppendInt64(o, z.Lucky)
	// string "Note"
	o = append(o, 0xa4, 0x4e, 0x6f, 0x74, 0x65)
	o = msgp.AppendString(o, z.Note)
	// string "Cheat"
	o = append(o, 0xa5, 0x43, 0x68, 0x65, 0x61, 0x74)
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
		case "End":
			z.End, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Room":
			z.Room, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Tab":
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
			for zcxo := range z.Bill {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Bill[zcxo] = nil
				} else {
					if z.Bill[zcxo] == nil {
						z.Bill[zcxo] = new(GameBill)
					}
					bts, err = z.Bill[zcxo].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Ante":
			z.Ante, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Ring":
			z.Ring, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Sum":
			z.Sum, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Win":
			z.Win, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Tax":
			z.Tax, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Water":
			z.Water, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Log":
			var zkgt uint32
			zkgt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Log) >= int(zkgt) {
				z.Log = (z.Log)[:zkgt]
			} else {
				z.Log = make([]*ActionLog, zkgt)
			}
			for zeff := range z.Log {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Log[zeff] = nil
				} else {
					if z.Log[zeff] == nil {
						z.Log[zeff] = new(ActionLog)
					}
					bts, err = z.Log[zeff].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Pool":
			z.Pool, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Lucky":
			z.Lucky, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Note":
			z.Note, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Cheat":
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
	s = 3 + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 4 + msgp.Int64Size + 5 + msgp.Int32Size + 4 + msgp.Int32Size + 2 + msgp.ArrayHeaderSize
	for zcxo := range z.Bill {
		if z.Bill[zcxo] == nil {
			s += msgp.NilSize
		} else {
			s += z.Bill[zcxo].Msgsize()
		}
	}
	s += 5 + msgp.Int32Size + 5 + msgp.Int32Size + 4 + msgp.Int64Size + 4 + msgp.Int64Size + 4 + msgp.Int64Size + 6 + msgp.Int64Size + 4 + msgp.ArrayHeaderSize
	for zeff := range z.Log {
		if z.Log[zeff] == nil {
			s += msgp.NilSize
		} else {
			s += z.Log[zeff].Msgsize()
		}
	}
	s += 5 + msgp.Int64Size + 6 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Note) + 6 + msgp.BoolSize
	return
}
