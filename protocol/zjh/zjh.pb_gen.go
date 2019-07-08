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
		case "Type":
			{
				var zcmr int32
				zcmr, err = dc.ReadInt32()
				z.Type = ActionType(zcmr)
			}
			if err != nil {
				return
			}
		case "Poker":
			z.Poker, err = dc.ReadBytes(z.Poker)
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
		case "Winners":
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
		case "Coin":
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
	// write "Type"
	err = en.Append(0x87, 0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(int32(z.Type))
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
	// write "Coin"
	err = en.Append(0xa4, 0x43, 0x6f, 0x69, 0x6e)
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
	// string "Type"
	o = append(o, 0x87, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt32(o, int32(z.Type))
	// string "Poker"
	o = append(o, 0xa5, 0x50, 0x6f, 0x6b, 0x65, 0x72)
	o = msgp.AppendBytes(o, z.Poker)
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
	// string "Coin"
	o = append(o, 0xa4, 0x43, 0x6f, 0x69, 0x6e)
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
		case "Type":
			{
				var zcua int32
				zcua, bts, err = msgp.ReadInt32Bytes(bts)
				z.Type = ActionType(zcua)
			}
			if err != nil {
				return
			}
		case "Poker":
			z.Poker, bts, err = msgp.ReadBytesBytes(bts, z.Poker)
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
		case "Winners":
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
		case "Coin":
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
	s = 1 + 5 + msgp.Int32Size + 6 + msgp.BytesPrefixSize + len(z.Poker) + 4 + msgp.Int32Size + 4 + msgp.Int32Size + 8 + msgp.ArrayHeaderSize + (len(z.Players) * (msgp.Int32Size)) + 8 + msgp.ArrayHeaderSize + (len(z.Winners) * (msgp.Int32Size)) + 5 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ActionReq) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zdaf uint32
	zdaf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zdaf > 0 {
		zdaf--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Type":
			{
				var zpks int32
				zpks, err = dc.ReadInt32()
				z.Type = ActionType(zpks)
			}
			if err != nil {
				return
			}
		case "Bet":
			z.Bet, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Opponent":
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
	// write "Type"
	err = en.Append(0x83, 0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(int32(z.Type))
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
	// write "Opponent"
	err = en.Append(0xa8, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74)
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
	// string "Type"
	o = append(o, 0x83, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt32(o, int32(z.Type))
	// string "Bet"
	o = append(o, 0xa3, 0x42, 0x65, 0x74)
	o = msgp.AppendInt32(o, z.Bet)
	// string "Opponent"
	o = append(o, 0xa8, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74)
	o = msgp.AppendInt32(o, z.Opponent)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ActionReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjfb uint32
	zjfb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjfb > 0 {
		zjfb--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Type":
			{
				var zcxo int32
				zcxo, bts, err = msgp.ReadInt32Bytes(bts)
				z.Type = ActionType(zcxo)
			}
			if err != nil {
				return
			}
		case "Bet":
			z.Bet, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Opponent":
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
	s = 1 + 5 + msgp.Int32Size + 4 + msgp.Int32Size + 9 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ActionType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zeff int32
		zeff, err = dc.ReadInt32()
		(*z) = ActionType(zeff)
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
		var zrsw int32
		zrsw, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = ActionType(zrsw)
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
		var zxpk int32
		zxpk, err = dc.ReadInt32()
		(*z) = Code(zxpk)
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
		var zdnj int32
		zdnj, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = Code(zdnj)
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
func (z *GameInitAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsnv uint32
	zsnv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsnv > 0 {
		zsnv--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Table":
			z.Table, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Id":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Pool":
			z.Pool, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "State":
			z.State, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Ring":
			z.Ring, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Players":
			var zkgt uint32
			zkgt, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zkgt) {
				z.Players = (z.Players)[:zkgt]
			} else {
				z.Players = make([]*Player, zkgt)
			}
			for zobc := range z.Players {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Players[zobc] = nil
				} else {
					if z.Players[zobc] == nil {
						z.Players[zobc] = new(Player)
					}
					err = z.Players[zobc].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "Poker":
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
	// write "Table"
	err = en.Append(0x87, 0xa5, 0x54, 0x61, 0x62, 0x6c, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Table)
	if err != nil {
		return
	}
	// write "Id"
	err = en.Append(0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
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
	// write "State"
	err = en.Append(0xa5, 0x53, 0x74, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.State)
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
	// write "Players"
	err = en.Append(0xa7, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Players)))
	if err != nil {
		return
	}
	for zobc := range z.Players {
		if z.Players[zobc] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Players[zobc].EncodeMsg(en)
			if err != nil {
				return
			}
		}
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
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameInitAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Table"
	o = append(o, 0x87, 0xa5, 0x54, 0x61, 0x62, 0x6c, 0x65)
	o = msgp.AppendInt32(o, z.Table)
	// string "Id"
	o = append(o, 0xa2, 0x49, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "Pool"
	o = append(o, 0xa4, 0x50, 0x6f, 0x6f, 0x6c)
	o = msgp.AppendInt64(o, z.Pool)
	// string "State"
	o = append(o, 0xa5, 0x53, 0x74, 0x61, 0x74, 0x65)
	o = msgp.AppendInt32(o, z.State)
	// string "Ring"
	o = append(o, 0xa4, 0x52, 0x69, 0x6e, 0x67)
	o = msgp.AppendInt32(o, z.Ring)
	// string "Players"
	o = append(o, 0xa7, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Players)))
	for zobc := range z.Players {
		if z.Players[zobc] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Players[zobc].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Poker"
	o = append(o, 0xa5, 0x50, 0x6f, 0x6b, 0x65, 0x72)
	o = msgp.AppendBytes(o, z.Poker)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameInitAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zema uint32
	zema, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zema > 0 {
		zema--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Table":
			z.Table, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Pool":
			z.Pool, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "State":
			z.State, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Ring":
			z.Ring, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Players":
			var zpez uint32
			zpez, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zpez) {
				z.Players = (z.Players)[:zpez]
			} else {
				z.Players = make([]*Player, zpez)
			}
			for zobc := range z.Players {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Players[zobc] = nil
				} else {
					if z.Players[zobc] == nil {
						z.Players[zobc] = new(Player)
					}
					bts, err = z.Players[zobc].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Poker":
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
	s = 1 + 6 + msgp.Int32Size + 3 + msgp.Int32Size + 5 + msgp.Int64Size + 6 + msgp.Int32Size + 5 + msgp.Int32Size + 8 + msgp.ArrayHeaderSize
	for zobc := range z.Players {
		if z.Players[zobc] == nil {
			s += msgp.NilSize
		} else {
			s += z.Players[zobc].Msgsize()
		}
	}
	s += 6 + msgp.BytesPrefixSize + len(z.Poker)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameResultAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Id":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Winner":
			var zywj uint32
			zywj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Winner) >= int(zywj) {
				z.Winner = (z.Winner)[:zywj]
			} else {
				z.Winner = make([]int32, zywj)
			}
			for zqke := range z.Winner {
				z.Winner[zqke], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "Prize":
			var zjpj uint32
			zjpj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Prize) >= int(zjpj) {
				z.Prize = (z.Prize)[:zjpj]
			} else {
				z.Prize = make([]int64, zjpj)
			}
			for zqyh := range z.Prize {
				z.Prize[zqyh], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "Coin":
			z.Coin, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Poker":
			z.Poker, err = dc.ReadBytes(z.Poker)
			if err != nil {
				return
			}
		case "Lucky":
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
	// write "Id"
	err = en.Append(0x86, 0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
	}
	// write "Winner"
	err = en.Append(0xa6, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Winner)))
	if err != nil {
		return
	}
	for zqke := range z.Winner {
		err = en.WriteInt32(z.Winner[zqke])
		if err != nil {
			return
		}
	}
	// write "Prize"
	err = en.Append(0xa5, 0x50, 0x72, 0x69, 0x7a, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Prize)))
	if err != nil {
		return
	}
	for zqyh := range z.Prize {
		err = en.WriteInt64(z.Prize[zqyh])
		if err != nil {
			return
		}
	}
	// write "Coin"
	err = en.Append(0xa4, 0x43, 0x6f, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Coin)
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
	// write "Lucky"
	err = en.Append(0xa5, 0x4c, 0x75, 0x63, 0x6b, 0x79)
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
	// string "Id"
	o = append(o, 0x86, 0xa2, 0x49, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "Winner"
	o = append(o, 0xa6, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Winner)))
	for zqke := range z.Winner {
		o = msgp.AppendInt32(o, z.Winner[zqke])
	}
	// string "Prize"
	o = append(o, 0xa5, 0x50, 0x72, 0x69, 0x7a, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Prize)))
	for zqyh := range z.Prize {
		o = msgp.AppendInt64(o, z.Prize[zqyh])
	}
	// string "Coin"
	o = append(o, 0xa4, 0x43, 0x6f, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.Coin)
	// string "Poker"
	o = append(o, 0xa5, 0x50, 0x6f, 0x6b, 0x65, 0x72)
	o = msgp.AppendBytes(o, z.Poker)
	// string "Lucky"
	o = append(o, 0xa5, 0x4c, 0x75, 0x63, 0x6b, 0x79)
	o = msgp.AppendInt64(o, z.Lucky)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameResultAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zzpf uint32
	zzpf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zzpf > 0 {
		zzpf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Winner":
			var zrfe uint32
			zrfe, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Winner) >= int(zrfe) {
				z.Winner = (z.Winner)[:zrfe]
			} else {
				z.Winner = make([]int32, zrfe)
			}
			for zqke := range z.Winner {
				z.Winner[zqke], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "Prize":
			var zgmo uint32
			zgmo, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Prize) >= int(zgmo) {
				z.Prize = (z.Prize)[:zgmo]
			} else {
				z.Prize = make([]int64, zgmo)
			}
			for zqyh := range z.Prize {
				z.Prize[zqyh], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "Coin":
			z.Coin, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Poker":
			z.Poker, bts, err = msgp.ReadBytesBytes(bts, z.Poker)
			if err != nil {
				return
			}
		case "Lucky":
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
func (z *GameStartAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Id":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Pool":
			z.Pool, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Players":
			var zsbz uint32
			zsbz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zsbz) {
				z.Players = (z.Players)[:zsbz]
			} else {
				z.Players = make([]*Player, zsbz)
			}
			for ztaf := range z.Players {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Players[ztaf] = nil
				} else {
					if z.Players[ztaf] == nil {
						z.Players[ztaf] = new(Player)
					}
					err = z.Players[ztaf].DecodeMsg(dc)
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
	// write "Id"
	err = en.Append(0x83, 0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
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
	// write "Players"
	err = en.Append(0xa7, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Players)))
	if err != nil {
		return
	}
	for ztaf := range z.Players {
		if z.Players[ztaf] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Players[ztaf].EncodeMsg(en)
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
	// string "Id"
	o = append(o, 0x83, 0xa2, 0x49, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "Pool"
	o = append(o, 0xa4, 0x50, 0x6f, 0x6f, 0x6c)
	o = msgp.AppendInt64(o, z.Pool)
	// string "Players"
	o = append(o, 0xa7, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Players)))
	for ztaf := range z.Players {
		if z.Players[ztaf] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Players[ztaf].MarshalMsg(o)
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
	var zrjx uint32
	zrjx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrjx > 0 {
		zrjx--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Pool":
			z.Pool, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Players":
			var zawn uint32
			zawn, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zawn) {
				z.Players = (z.Players)[:zawn]
			} else {
				z.Players = make([]*Player, zawn)
			}
			for ztaf := range z.Players {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Players[ztaf] = nil
				} else {
					if z.Players[ztaf] == nil {
						z.Players[ztaf] = new(Player)
					}
					bts, err = z.Players[ztaf].UnmarshalMsg(bts)
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
	s = 1 + 3 + msgp.Int32Size + 5 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for ztaf := range z.Players {
		if z.Players[ztaf] == nil {
			s += msgp.NilSize
		} else {
			s += z.Players[ztaf].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Player) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Id":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Icon":
			z.Icon, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Vip":
			z.Vip, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Chair":
			z.Chair, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Coin":
			z.Coin, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Bet":
			z.Bet, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "State":
			{
				var zrbe int32
				zrbe, err = dc.ReadInt32()
				z.State = Player_State(zrbe)
			}
			if err != nil {
				return
			}
		case "Look":
			z.Look, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Down":
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
	// write "Id"
	err = en.Append(0x8a, 0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
	}
	// write "Icon"
	err = en.Append(0xa4, 0x49, 0x63, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Icon)
	if err != nil {
		return
	}
	// write "Vip"
	err = en.Append(0xa3, 0x56, 0x69, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Vip)
	if err != nil {
		return
	}
	// write "Chair"
	err = en.Append(0xa5, 0x43, 0x68, 0x61, 0x69, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Chair)
	if err != nil {
		return
	}
	// write "Coin"
	err = en.Append(0xa4, 0x43, 0x6f, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Coin)
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
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "State"
	err = en.Append(0xa5, 0x53, 0x74, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(int32(z.State))
	if err != nil {
		return
	}
	// write "Look"
	err = en.Append(0xa4, 0x4c, 0x6f, 0x6f, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Look)
	if err != nil {
		return
	}
	// write "Down"
	err = en.Append(0xa4, 0x44, 0x6f, 0x77, 0x6e)
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
	// string "Id"
	o = append(o, 0x8a, 0xa2, 0x49, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "Icon"
	o = append(o, 0xa4, 0x49, 0x63, 0x6f, 0x6e)
	o = msgp.AppendInt32(o, z.Icon)
	// string "Vip"
	o = append(o, 0xa3, 0x56, 0x69, 0x70)
	o = msgp.AppendInt32(o, z.Vip)
	// string "Chair"
	o = append(o, 0xa5, 0x43, 0x68, 0x61, 0x69, 0x72)
	o = msgp.AppendInt32(o, z.Chair)
	// string "Coin"
	o = append(o, 0xa4, 0x43, 0x6f, 0x69, 0x6e)
	o = msgp.AppendInt64(o, z.Coin)
	// string "Bet"
	o = append(o, 0xa3, 0x42, 0x65, 0x74)
	o = msgp.AppendInt64(o, z.Bet)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "State"
	o = append(o, 0xa5, 0x53, 0x74, 0x61, 0x74, 0x65)
	o = msgp.AppendInt32(o, int32(z.State))
	// string "Look"
	o = append(o, 0xa4, 0x4c, 0x6f, 0x6f, 0x6b)
	o = msgp.AppendBool(o, z.Look)
	// string "Down"
	o = append(o, 0xa4, 0x44, 0x6f, 0x77, 0x6e)
	o = msgp.AppendInt32(o, z.Down)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Player) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zmfd uint32
	zmfd, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zmfd > 0 {
		zmfd--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Icon":
			z.Icon, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Vip":
			z.Vip, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Chair":
			z.Chair, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Coin":
			z.Coin, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Bet":
			z.Bet, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "State":
			{
				var zzdc int32
				zzdc, bts, err = msgp.ReadInt32Bytes(bts)
				z.State = Player_State(zzdc)
			}
			if err != nil {
				return
			}
		case "Look":
			z.Look, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "Down":
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
		var zelx int32
		zelx, err = dc.ReadInt32()
		(*z) = Player_State(zelx)
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
		var zbal int32
		zbal, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = Player_State(zbal)
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
