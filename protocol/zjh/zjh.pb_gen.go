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
		case "type":
			{
				var zcmr int32
				zcmr, err = dc.ReadInt32()
				z.Type = ActionType(zcmr)
			}
			if err != nil {
				return
			}
		case "poker":
			z.Poker, err = dc.ReadBytes(z.Poker)
			if err != nil {
				return
			}
		case "uid":
			z.Uid, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "bet":
			z.Bet, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "play":
			var zajw uint32
			zajw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zajw) {
				z.Players = (z.Players)[:zajw]
			} else {
				z.Players = make([]int32, zajw)
			}
			for zxvk := range z.Players {
				z.Players[zxvk], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "win":
			var zwht uint32
			zwht, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Winners) >= int(zwht) {
				z.Winners = (z.Winners)[:zwht]
			} else {
				z.Winners = make([]int32, zwht)
			}
			for zbzg := range z.Winners {
				z.Winners[zbzg], err = dc.ReadInt32()
				if err != nil {
					return
				}
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
func (z *ActionAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "type"
	err = en.Append(0x87, 0xa4, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(int32(z.Type))
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
	// write "uid"
	err = en.Append(0xa3, 0x75, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Uid)
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
	// write "play"
	err = en.Append(0xa4, 0x70, 0x6c, 0x61, 0x79)
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
	// write "win"
	err = en.Append(0xa3, 0x77, 0x69, 0x6e)
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
func (z *ActionAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "type"
	o = append(o, 0x87, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendInt32(o, int32(z.Type))
	// string "poker"
	o = append(o, 0xa5, 0x70, 0x6f, 0x6b, 0x65, 0x72)
	o = msgp.AppendBytes(o, z.Poker)
	// string "uid"
	o = append(o, 0xa3, 0x75, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Uid)
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendInt32(o, z.Bet)
	// string "play"
	o = append(o, 0xa4, 0x70, 0x6c, 0x61, 0x79)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Players)))
	for zxvk := range z.Players {
		o = msgp.AppendInt32(o, z.Players[zxvk])
	}
	// string "win"
	o = append(o, 0xa3, 0x77, 0x69, 0x6e)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Winners)))
	for zbzg := range z.Winners {
		o = msgp.AppendInt32(o, z.Winners[zbzg])
	}
	// string "coin"
	o = append(o, 0xa4, 0x63, 0x6f, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.Coin)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ActionAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "type":
			{
				var zcua int32
				zcua, bts, err = msgp.ReadInt32Bytes(bts)
				z.Type = ActionType(zcua)
			}
			if err != nil {
				return
			}
		case "poker":
			z.Poker, bts, err = msgp.ReadBytesBytes(bts, z.Poker)
			if err != nil {
				return
			}
		case "uid":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "bet":
			z.Bet, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "play":
			var zxhx uint32
			zxhx, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zxhx) {
				z.Players = (z.Players)[:zxhx]
			} else {
				z.Players = make([]int32, zxhx)
			}
			for zxvk := range z.Players {
				z.Players[zxvk], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "win":
			var zlqf uint32
			zlqf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Winners) >= int(zlqf) {
				z.Winners = (z.Winners)[:zlqf]
			} else {
				z.Winners = make([]int32, zlqf)
			}
			for zbzg := range z.Winners {
				z.Winners[zbzg], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
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
func (z *ActionAck) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int32Size + 6 + msgp.BytesPrefixSize + len(z.Poker) + 4 + msgp.Int32Size + 4 + msgp.Int32Size + 5 + msgp.ArrayHeaderSize + (len(z.Players) * (msgp.Int32Size)) + 4 + msgp.ArrayHeaderSize + (len(z.Winners) * (msgp.Int32Size)) + 5 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ActionLog) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "s":
			z.Start, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "t":
			{
				var zcxo int32
				zcxo, err = dc.ReadInt32()
				z.Type = ActionType(zcxo)
			}
			if err != nil {
				return
			}
		case "u":
			z.Uid, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "c":
			z.Bet, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "p":
			var zeff uint32
			zeff, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zeff) {
				z.Players = (z.Players)[:zeff]
			} else {
				z.Players = make([]int32, zeff)
			}
			for zdaf := range z.Players {
				z.Players[zdaf], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "w":
			var zrsw uint32
			zrsw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Winners) >= int(zrsw) {
				z.Winners = (z.Winners)[:zrsw]
			} else {
				z.Winners = make([]int32, zrsw)
			}
			for zpks := range z.Winners {
				z.Winners[zpks], err = dc.ReadInt32()
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
	// write "s"
	err = en.Append(0x86, 0xa1, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Start)
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
	// write "u"
	err = en.Append(0xa1, 0x75)
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
	err = en.WriteInt32(z.Bet)
	if err != nil {
		return
	}
	// write "p"
	err = en.Append(0xa1, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Players)))
	if err != nil {
		return
	}
	for zdaf := range z.Players {
		err = en.WriteInt32(z.Players[zdaf])
		if err != nil {
			return
		}
	}
	// write "w"
	err = en.Append(0xa1, 0x77)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Winners)))
	if err != nil {
		return
	}
	for zpks := range z.Winners {
		err = en.WriteInt32(z.Winners[zpks])
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
	// string "s"
	o = append(o, 0x86, 0xa1, 0x73)
	o = msgp.AppendInt64(o, z.Start)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendInt32(o, int32(z.Type))
	// string "u"
	o = append(o, 0xa1, 0x75)
	o = msgp.AppendInt32(o, z.Uid)
	// string "c"
	o = append(o, 0xa1, 0x63)
	o = msgp.AppendInt32(o, z.Bet)
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Players)))
	for zdaf := range z.Players {
		o = msgp.AppendInt32(o, z.Players[zdaf])
	}
	// string "w"
	o = append(o, 0xa1, 0x77)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Winners)))
	for zpks := range z.Winners {
		o = msgp.AppendInt32(o, z.Winners[zpks])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ActionLog) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "s":
			z.Start, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "t":
			{
				var zdnj int32
				zdnj, bts, err = msgp.ReadInt32Bytes(bts)
				z.Type = ActionType(zdnj)
			}
			if err != nil {
				return
			}
		case "u":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "c":
			z.Bet, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "p":
			var zobc uint32
			zobc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zobc) {
				z.Players = (z.Players)[:zobc]
			} else {
				z.Players = make([]int32, zobc)
			}
			for zdaf := range z.Players {
				z.Players[zdaf], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "w":
			var zsnv uint32
			zsnv, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Winners) >= int(zsnv) {
				z.Winners = (z.Winners)[:zsnv]
			} else {
				z.Winners = make([]int32, zsnv)
			}
			for zpks := range z.Winners {
				z.Winners[zpks], bts, err = msgp.ReadInt32Bytes(bts)
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
	s = 1 + 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.ArrayHeaderSize + (len(z.Players) * (msgp.Int32Size)) + 2 + msgp.ArrayHeaderSize + (len(z.Winners) * (msgp.Int32Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ActionReq) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "type":
			{
				var zema int32
				zema, err = dc.ReadInt32()
				z.Type = ActionType(zema)
			}
			if err != nil {
				return
			}
		case "bet":
			z.Bet, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "opp":
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
	// write "type"
	err = en.Append(0x83, 0xa4, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(int32(z.Type))
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
	// write "opp"
	err = en.Append(0xa3, 0x6f, 0x70, 0x70)
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
	// string "type"
	o = append(o, 0x83, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendInt32(o, int32(z.Type))
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendInt32(o, z.Bet)
	// string "opp"
	o = append(o, 0xa3, 0x6f, 0x70, 0x70)
	o = msgp.AppendInt32(o, z.Opponent)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ActionReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "type":
			{
				var zqke int32
				zqke, bts, err = msgp.ReadInt32Bytes(bts)
				z.Type = ActionType(zqke)
			}
			if err != nil {
				return
			}
		case "bet":
			z.Bet, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "opp":
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
	s = 1 + 5 + msgp.Int32Size + 4 + msgp.Int32Size + 4 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ActionType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zqyh int32
		zqyh, err = dc.ReadInt32()
		(*z) = ActionType(zqyh)
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
		var zyzr int32
		zyzr, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = ActionType(zyzr)
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
		var zywj int32
		zywj, err = dc.ReadInt32()
		(*z) = Code(zywj)
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
		var zjpj int32
		zjpj, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = Code(zjpj)
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
	var zrfe uint32
	zrfe, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrfe > 0 {
		zrfe--
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
			z.OldCoin, err = dc.ReadInt64()
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
			var zgmo uint32
			zgmo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Pk) >= int(zgmo) {
				z.Pk = (z.Pk)[:zgmo]
			} else {
				z.Pk = make([]int32, zgmo)
			}
			for zzpf := range z.Pk {
				z.Pk[zzpf], err = dc.ReadInt32()
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
	err = en.WriteInt64(z.OldCoin)
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
	for zzpf := range z.Pk {
		err = en.WriteInt32(z.Pk[zzpf])
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
	o = msgp.AppendInt64(o, z.OldCoin)
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
	for zzpf := range z.Pk {
		o = msgp.AppendInt32(o, z.Pk[zzpf])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameBill) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "j":
			z.Job, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "c":
			z.OldCoin, bts, err = msgp.ReadInt64Bytes(bts)
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
			var zeth uint32
			zeth, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Pk) >= int(zeth) {
				z.Pk = (z.Pk)[:zeth]
			} else {
				z.Pk = make([]int32, zeth)
			}
			for zzpf := range z.Pk {
				z.Pk[zzpf], bts, err = msgp.ReadInt32Bytes(bts)
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
func (z *GameInitAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "play":
			var zawn uint32
			zawn, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zawn) {
				z.Players = (z.Players)[:zawn]
			} else {
				z.Players = make([]*Player, zawn)
			}
			for zsbz := range z.Players {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Players[zsbz] = nil
				} else {
					if z.Players[zsbz] == nil {
						z.Players[zsbz] = new(Player)
					}
					err = z.Players[zsbz].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "poker":
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
	// map header, size 7
	// write "table"
	err = en.Append(0x87, 0xa5, 0x74, 0x61, 0x62, 0x6c, 0x65)
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
	// write "play"
	err = en.Append(0xa4, 0x70, 0x6c, 0x61, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Players)))
	if err != nil {
		return
	}
	for zsbz := range z.Players {
		if z.Players[zsbz] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Players[zsbz].EncodeMsg(en)
			if err != nil {
				return
			}
		}
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
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameInitAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "table"
	o = append(o, 0x87, 0xa5, 0x74, 0x61, 0x62, 0x6c, 0x65)
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
	// string "play"
	o = append(o, 0xa4, 0x70, 0x6c, 0x61, 0x79)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Players)))
	for zsbz := range z.Players {
		if z.Players[zsbz] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Players[zsbz].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "poker"
	o = append(o, 0xa5, 0x70, 0x6f, 0x6b, 0x65, 0x72)
	o = msgp.AppendBytes(o, z.Poker)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameInitAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zwel uint32
	zwel, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zwel > 0 {
		zwel--
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
		case "play":
			var zrbe uint32
			zrbe, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zrbe) {
				z.Players = (z.Players)[:zrbe]
			} else {
				z.Players = make([]*Player, zrbe)
			}
			for zsbz := range z.Players {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Players[zsbz] = nil
				} else {
					if z.Players[zsbz] == nil {
						z.Players[zsbz] = new(Player)
					}
					bts, err = z.Players[zsbz].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "poker":
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
	s = 1 + 6 + msgp.Int32Size + 3 + msgp.Int32Size + 5 + msgp.Int64Size + 6 + msgp.Int32Size + 5 + msgp.Int32Size + 5 + msgp.ArrayHeaderSize
	for zsbz := range z.Players {
		if z.Players[zsbz] == nil {
			s += msgp.NilSize
		} else {
			s += z.Players[zsbz].Msgsize()
		}
	}
	s += 6 + msgp.BytesPrefixSize + len(z.Poker)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameResultAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zelx uint32
	zelx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zelx > 0 {
		zelx--
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
		case "winner":
			var zbal uint32
			zbal, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Winner) >= int(zbal) {
				z.Winner = (z.Winner)[:zbal]
			} else {
				z.Winner = make([]int32, zbal)
			}
			for zmfd := range z.Winner {
				z.Winner[zmfd], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "prize":
			var zjqz uint32
			zjqz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Prize) >= int(zjqz) {
				z.Prize = (z.Prize)[:zjqz]
			} else {
				z.Prize = make([]int64, zjqz)
			}
			for zzdc := range z.Prize {
				z.Prize[zzdc], err = dc.ReadInt64()
				if err != nil {
					return
				}
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
func (z *GameResultAck) EncodeMsg(en *msgp.Writer) (err error) {
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
	// write "winner"
	err = en.Append(0xa6, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Winner)))
	if err != nil {
		return
	}
	for zmfd := range z.Winner {
		err = en.WriteInt32(z.Winner[zmfd])
		if err != nil {
			return
		}
	}
	// write "prize"
	err = en.Append(0xa5, 0x70, 0x72, 0x69, 0x7a, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Prize)))
	if err != nil {
		return
	}
	for zzdc := range z.Prize {
		err = en.WriteInt64(z.Prize[zzdc])
		if err != nil {
			return
		}
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
func (z *GameResultAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "id"
	o = append(o, 0x86, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "winner"
	o = append(o, 0xa6, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Winner)))
	for zmfd := range z.Winner {
		o = msgp.AppendInt32(o, z.Winner[zmfd])
	}
	// string "prize"
	o = append(o, 0xa5, 0x70, 0x72, 0x69, 0x7a, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Prize)))
	for zzdc := range z.Prize {
		o = msgp.AppendInt64(o, z.Prize[zzdc])
	}
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
func (z *GameResultAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "winner":
			var ztmt uint32
			ztmt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Winner) >= int(ztmt) {
				z.Winner = (z.Winner)[:ztmt]
			} else {
				z.Winner = make([]int32, ztmt)
			}
			for zmfd := range z.Winner {
				z.Winner[zmfd], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "prize":
			var ztco uint32
			ztco, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Prize) >= int(ztco) {
				z.Prize = (z.Prize)[:ztco]
			} else {
				z.Prize = make([]int64, ztco)
			}
			for zzdc := range z.Prize {
				z.Prize[zzdc], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
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
func (z *GameResultAck) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 7 + msgp.ArrayHeaderSize + (len(z.Winner) * (msgp.Int32Size)) + 6 + msgp.ArrayHeaderSize + (len(z.Prize) * (msgp.Int64Size)) + 5 + msgp.Int64Size + 6 + msgp.BytesPrefixSize + len(z.Poker) + 6 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameRound) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zinl uint32
	zinl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zinl > 0 {
		zinl--
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
			var zare uint32
			zare, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(zare) {
				z.Bill = (z.Bill)[:zare]
			} else {
				z.Bill = make([]*GameBill, zare)
			}
			for zana := range z.Bill {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Bill[zana] = nil
				} else {
					if z.Bill[zana] == nil {
						z.Bill[zana] = new(GameBill)
					}
					err = z.Bill[zana].DecodeMsg(dc)
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
			var zljy uint32
			zljy, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Log) >= int(zljy) {
				z.Log = (z.Log)[:zljy]
			} else {
				z.Log = make([]*ActionLog, zljy)
			}
			for ztyy := range z.Log {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Log[ztyy] = nil
				} else {
					if z.Log[ztyy] == nil {
						z.Log[ztyy] = new(ActionLog)
					}
					err = z.Log[ztyy].DecodeMsg(dc)
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
	for zana := range z.Bill {
		if z.Bill[zana] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Bill[zana].EncodeMsg(en)
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
	for ztyy := range z.Log {
		if z.Log[ztyy] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Log[ztyy].EncodeMsg(en)
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
	for zana := range z.Bill {
		if z.Bill[zana] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Bill[zana].MarshalMsg(o)
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
	for ztyy := range z.Log {
		if z.Log[ztyy] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Log[ztyy].MarshalMsg(o)
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
			var zrsc uint32
			zrsc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bill) >= int(zrsc) {
				z.Bill = (z.Bill)[:zrsc]
			} else {
				z.Bill = make([]*GameBill, zrsc)
			}
			for zana := range z.Bill {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Bill[zana] = nil
				} else {
					if z.Bill[zana] == nil {
						z.Bill[zana] = new(GameBill)
					}
					bts, err = z.Bill[zana].UnmarshalMsg(bts)
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
			var zctn uint32
			zctn, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Log) >= int(zctn) {
				z.Log = (z.Log)[:zctn]
			} else {
				z.Log = make([]*ActionLog, zctn)
			}
			for ztyy := range z.Log {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Log[ztyy] = nil
				} else {
					if z.Log[ztyy] == nil {
						z.Log[ztyy] = new(ActionLog)
					}
					bts, err = z.Log[ztyy].UnmarshalMsg(bts)
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
	for zana := range z.Bill {
		if z.Bill[zana] == nil {
			s += msgp.NilSize
		} else {
			s += z.Bill[zana].Msgsize()
		}
	}
	s += 2 + msgp.Int64Size + 2 + msgp.Int32Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.ArrayHeaderSize
	for ztyy := range z.Log {
		if z.Log[ztyy] == nil {
			s += msgp.NilSize
		} else {
			s += z.Log[ztyy].Msgsize()
		}
	}
	s += 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.StringPrefixSize + len(z.Note) + 2 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameStartAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var znsg uint32
	znsg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for znsg > 0 {
		znsg--
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
			var zrus uint32
			zrus, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zrus) {
				z.Players = (z.Players)[:zrus]
			} else {
				z.Players = make([]*Player, zrus)
			}
			for zswy := range z.Players {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Players[zswy] = nil
				} else {
					if z.Players[zswy] == nil {
						z.Players[zswy] = new(Player)
					}
					err = z.Players[zswy].DecodeMsg(dc)
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
	err = en.WriteArrayHeader(uint32(len(z.Players)))
	if err != nil {
		return
	}
	for zswy := range z.Players {
		if z.Players[zswy] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Players[zswy].EncodeMsg(en)
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
	o = msgp.AppendArrayHeader(o, uint32(len(z.Players)))
	for zswy := range z.Players {
		if z.Players[zswy] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Players[zswy].MarshalMsg(o)
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
	var zsvm uint32
	zsvm, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsvm > 0 {
		zsvm--
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
			var zaoz uint32
			zaoz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zaoz) {
				z.Players = (z.Players)[:zaoz]
			} else {
				z.Players = make([]*Player, zaoz)
			}
			for zswy := range z.Players {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Players[zswy] = nil
				} else {
					if z.Players[zswy] == nil {
						z.Players[zswy] = new(Player)
					}
					bts, err = z.Players[zswy].UnmarshalMsg(bts)
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
	for zswy := range z.Players {
		if z.Players[zswy] == nil {
			s += msgp.NilSize
		} else {
			s += z.Players[zswy].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Player) DecodeMsg(dc *msgp.Reader) (err error) {
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
				var zsbo int32
				zsbo, err = dc.ReadInt32()
				z.State = Player_State(zsbo)
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
				var zqgz int32
				zqgz, bts, err = msgp.ReadInt32Bytes(bts)
				z.State = Player_State(zqgz)
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
		var zsnw int32
		zsnw, err = dc.ReadInt32()
		(*z) = Player_State(zsnw)
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
		var ztls int32
		ztls, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = Player_State(ztls)
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
