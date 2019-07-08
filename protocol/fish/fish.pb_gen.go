package fish

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Bullet) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "id":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "uid":
			z.Uid, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "client":
			z.Client, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "created":
			z.Created, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "bet":
			z.Bet, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "direction":
			z.Direction, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "x":
			z.X, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "y":
			z.Y, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "fish":
			z.Fish, err = dc.ReadInt32()
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
func (z *Bullet) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "id"
	err = en.Append(0x89, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
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
	// write "client"
	err = en.Append(0xa6, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Client)
	if err != nil {
		return
	}
	// write "created"
	err = en.Append(0xa7, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Created)
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
	// write "direction"
	err = en.Append(0xa9, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Direction)
	if err != nil {
		return
	}
	// write "x"
	err = en.Append(0xa1, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.X)
	if err != nil {
		return
	}
	// write "y"
	err = en.Append(0xa1, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Y)
	if err != nil {
		return
	}
	// write "fish"
	err = en.Append(0xa4, 0x66, 0x69, 0x73, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Fish)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Bullet) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "id"
	o = append(o, 0x89, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "uid"
	o = append(o, 0xa3, 0x75, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Uid)
	// string "client"
	o = append(o, 0xa6, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74)
	o = msgp.AppendInt32(o, z.Client)
	// string "created"
	o = append(o, 0xa7, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Created)
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendInt32(o, z.Bet)
	// string "direction"
	o = append(o, 0xa9, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendFloat64(o, z.Direction)
	// string "x"
	o = append(o, 0xa1, 0x78)
	o = msgp.AppendFloat64(o, z.X)
	// string "y"
	o = append(o, 0xa1, 0x79)
	o = msgp.AppendFloat64(o, z.Y)
	// string "fish"
	o = append(o, 0xa4, 0x66, 0x69, 0x73, 0x68)
	o = msgp.AppendInt32(o, z.Fish)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Bullet) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "uid":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "client":
			z.Client, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "created":
			z.Created, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "bet":
			z.Bet, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "direction":
			z.Direction, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "x":
			z.X, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "y":
			z.Y, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "fish":
			z.Fish, bts, err = msgp.ReadInt32Bytes(bts)
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
func (z *Bullet) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 4 + msgp.Int32Size + 7 + msgp.Int32Size + 8 + msgp.Int64Size + 4 + msgp.Int32Size + 10 + msgp.Float64Size + 2 + msgp.Float64Size + 2 + msgp.Float64Size + 5 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Code) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zbai int32
		zbai, err = dc.ReadInt32()
		(*z) = Code(zbai)
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
		var zcmr int32
		zcmr, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = Code(zcmr)
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
func (z *Fish) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "id":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "tmpId":
			z.TmpId, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "pathId":
			z.PathId, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "boxId":
			z.BoxId, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "speed":
			z.Speed, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "created":
			z.Created, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "direction":
			z.Direction, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "x":
			z.X, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "y":
			z.Y, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "troop":
			z.Troop, err = dc.ReadBool()
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
func (z *Fish) EncodeMsg(en *msgp.Writer) (err error) {
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
	// write "tmpId"
	err = en.Append(0xa5, 0x74, 0x6d, 0x70, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.TmpId)
	if err != nil {
		return
	}
	// write "pathId"
	err = en.Append(0xa6, 0x70, 0x61, 0x74, 0x68, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.PathId)
	if err != nil {
		return
	}
	// write "boxId"
	err = en.Append(0xa5, 0x62, 0x6f, 0x78, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.BoxId)
	if err != nil {
		return
	}
	// write "speed"
	err = en.Append(0xa5, 0x73, 0x70, 0x65, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Speed)
	if err != nil {
		return
	}
	// write "created"
	err = en.Append(0xa7, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Created)
	if err != nil {
		return
	}
	// write "direction"
	err = en.Append(0xa9, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Direction)
	if err != nil {
		return
	}
	// write "x"
	err = en.Append(0xa1, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.X)
	if err != nil {
		return
	}
	// write "y"
	err = en.Append(0xa1, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Y)
	if err != nil {
		return
	}
	// write "troop"
	err = en.Append(0xa5, 0x74, 0x72, 0x6f, 0x6f, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Troop)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Fish) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "id"
	o = append(o, 0x8a, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "tmpId"
	o = append(o, 0xa5, 0x74, 0x6d, 0x70, 0x49, 0x64)
	o = msgp.AppendInt32(o, z.TmpId)
	// string "pathId"
	o = append(o, 0xa6, 0x70, 0x61, 0x74, 0x68, 0x49, 0x64)
	o = msgp.AppendInt32(o, z.PathId)
	// string "boxId"
	o = append(o, 0xa5, 0x62, 0x6f, 0x78, 0x49, 0x64)
	o = msgp.AppendInt32(o, z.BoxId)
	// string "speed"
	o = append(o, 0xa5, 0x73, 0x70, 0x65, 0x65, 0x64)
	o = msgp.AppendFloat64(o, z.Speed)
	// string "created"
	o = append(o, 0xa7, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Created)
	// string "direction"
	o = append(o, 0xa9, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendFloat64(o, z.Direction)
	// string "x"
	o = append(o, 0xa1, 0x78)
	o = msgp.AppendFloat64(o, z.X)
	// string "y"
	o = append(o, 0xa1, 0x79)
	o = msgp.AppendFloat64(o, z.Y)
	// string "troop"
	o = append(o, 0xa5, 0x74, 0x72, 0x6f, 0x6f, 0x70)
	o = msgp.AppendBool(o, z.Troop)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Fish) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "tmpId":
			z.TmpId, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "pathId":
			z.PathId, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "boxId":
			z.BoxId, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "speed":
			z.Speed, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "created":
			z.Created, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "direction":
			z.Direction, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "x":
			z.X, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "y":
			z.Y, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "troop":
			z.Troop, bts, err = msgp.ReadBoolBytes(bts)
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
func (z *Fish) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 6 + msgp.Int32Size + 7 + msgp.Int32Size + 6 + msgp.Int32Size + 6 + msgp.Float64Size + 8 + msgp.Int64Size + 10 + msgp.Float64Size + 2 + msgp.Float64Size + 2 + msgp.Float64Size + 6 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FishSeed) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "id":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "speed":
			z.Speed, err = dc.ReadFloat64()
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
func (z FishSeed) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "id"
	err = en.Append(0x82, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
	}
	// write "speed"
	err = en.Append(0xa5, 0x73, 0x70, 0x65, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Speed)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z FishSeed) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "id"
	o = append(o, 0x82, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "speed"
	o = append(o, 0xa5, 0x73, 0x70, 0x65, 0x65, 0x64)
	o = msgp.AppendFloat64(o, z.Speed)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FishSeed) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "speed":
			z.Speed, bts, err = msgp.ReadFloat64Bytes(bts)
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
func (z FishSeed) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 6 + msgp.Float64Size
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
		case "tick":
			z.Tick, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "scene":
			z.Scene, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "play":
			var zjfb uint32
			zjfb, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zjfb) {
				z.Players = (z.Players)[:zjfb]
			} else {
				z.Players = make([]*Player, zjfb)
			}
			for zxhx := range z.Players {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Players[zxhx] = nil
				} else {
					if z.Players[zxhx] == nil {
						z.Players[zxhx] = new(Player)
					}
					err = z.Players[zxhx].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "fishes":
			var zcxo uint32
			zcxo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Fishes) >= int(zcxo) {
				z.Fishes = (z.Fishes)[:zcxo]
			} else {
				z.Fishes = make([]*Fish, zcxo)
			}
			for zlqf := range z.Fishes {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Fishes[zlqf] = nil
				} else {
					if z.Fishes[zlqf] == nil {
						z.Fishes[zlqf] = new(Fish)
					}
					err = z.Fishes[zlqf].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "bull":
			var zeff uint32
			zeff, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bullets) >= int(zeff) {
				z.Bullets = (z.Bullets)[:zeff]
			} else {
				z.Bullets = make([]*Bullet, zeff)
			}
			for zdaf := range z.Bullets {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Bullets[zdaf] = nil
				} else {
					if z.Bullets[zdaf] == nil {
						z.Bullets[zdaf] = new(Bullet)
					}
					err = z.Bullets[zdaf].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "maxBullet":
			z.MaxBullet, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "fireInterval":
			z.FireInterval, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "stopFire":
			z.StopFire, err = dc.ReadBool()
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
	// write "tick"
	err = en.Append(0xa4, 0x74, 0x69, 0x63, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Tick)
	if err != nil {
		return
	}
	// write "scene"
	err = en.Append(0xa5, 0x73, 0x63, 0x65, 0x6e, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Scene)
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
	for zxhx := range z.Players {
		if z.Players[zxhx] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Players[zxhx].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "fishes"
	err = en.Append(0xa6, 0x66, 0x69, 0x73, 0x68, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Fishes)))
	if err != nil {
		return
	}
	for zlqf := range z.Fishes {
		if z.Fishes[zlqf] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Fishes[zlqf].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "bull"
	err = en.Append(0xa4, 0x62, 0x75, 0x6c, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Bullets)))
	if err != nil {
		return
	}
	for zdaf := range z.Bullets {
		if z.Bullets[zdaf] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Bullets[zdaf].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "maxBullet"
	err = en.Append(0xa9, 0x6d, 0x61, 0x78, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.MaxBullet)
	if err != nil {
		return
	}
	// write "fireInterval"
	err = en.Append(0xac, 0x66, 0x69, 0x72, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.FireInterval)
	if err != nil {
		return
	}
	// write "stopFire"
	err = en.Append(0xa8, 0x73, 0x74, 0x6f, 0x70, 0x46, 0x69, 0x72, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.StopFire)
	if err != nil {
		return
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
	// string "tick"
	o = append(o, 0xa4, 0x74, 0x69, 0x63, 0x6b)
	o = msgp.AppendInt64(o, z.Tick)
	// string "scene"
	o = append(o, 0xa5, 0x73, 0x63, 0x65, 0x6e, 0x65)
	o = msgp.AppendInt32(o, z.Scene)
	// string "play"
	o = append(o, 0xa4, 0x70, 0x6c, 0x61, 0x79)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Players)))
	for zxhx := range z.Players {
		if z.Players[zxhx] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Players[zxhx].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "fishes"
	o = append(o, 0xa6, 0x66, 0x69, 0x73, 0x68, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Fishes)))
	for zlqf := range z.Fishes {
		if z.Fishes[zlqf] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Fishes[zlqf].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "bull"
	o = append(o, 0xa4, 0x62, 0x75, 0x6c, 0x6c)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bullets)))
	for zdaf := range z.Bullets {
		if z.Bullets[zdaf] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Bullets[zdaf].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "maxBullet"
	o = append(o, 0xa9, 0x6d, 0x61, 0x78, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74)
	o = msgp.AppendInt32(o, z.MaxBullet)
	// string "fireInterval"
	o = append(o, 0xac, 0x66, 0x69, 0x72, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c)
	o = msgp.AppendInt32(o, z.FireInterval)
	// string "stopFire"
	o = append(o, 0xa8, 0x73, 0x74, 0x6f, 0x70, 0x46, 0x69, 0x72, 0x65)
	o = msgp.AppendBool(o, z.StopFire)
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
		case "tick":
			z.Tick, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "scene":
			z.Scene, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "play":
			var zxpk uint32
			zxpk, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zxpk) {
				z.Players = (z.Players)[:zxpk]
			} else {
				z.Players = make([]*Player, zxpk)
			}
			for zxhx := range z.Players {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Players[zxhx] = nil
				} else {
					if z.Players[zxhx] == nil {
						z.Players[zxhx] = new(Player)
					}
					bts, err = z.Players[zxhx].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "fishes":
			var zdnj uint32
			zdnj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Fishes) >= int(zdnj) {
				z.Fishes = (z.Fishes)[:zdnj]
			} else {
				z.Fishes = make([]*Fish, zdnj)
			}
			for zlqf := range z.Fishes {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Fishes[zlqf] = nil
				} else {
					if z.Fishes[zlqf] == nil {
						z.Fishes[zlqf] = new(Fish)
					}
					bts, err = z.Fishes[zlqf].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "bull":
			var zobc uint32
			zobc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bullets) >= int(zobc) {
				z.Bullets = (z.Bullets)[:zobc]
			} else {
				z.Bullets = make([]*Bullet, zobc)
			}
			for zdaf := range z.Bullets {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Bullets[zdaf] = nil
				} else {
					if z.Bullets[zdaf] == nil {
						z.Bullets[zdaf] = new(Bullet)
					}
					bts, err = z.Bullets[zdaf].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "maxBullet":
			z.MaxBullet, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "fireInterval":
			z.FireInterval, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "stopFire":
			z.StopFire, bts, err = msgp.ReadBoolBytes(bts)
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
	s = 1 + 6 + msgp.Int32Size + 5 + msgp.Int64Size + 6 + msgp.Int32Size + 5 + msgp.ArrayHeaderSize
	for zxhx := range z.Players {
		if z.Players[zxhx] == nil {
			s += msgp.NilSize
		} else {
			s += z.Players[zxhx].Msgsize()
		}
	}
	s += 7 + msgp.ArrayHeaderSize
	for zlqf := range z.Fishes {
		if z.Fishes[zlqf] == nil {
			s += msgp.NilSize
		} else {
			s += z.Fishes[zlqf].Msgsize()
		}
	}
	s += 5 + msgp.ArrayHeaderSize
	for zdaf := range z.Bullets {
		if z.Bullets[zdaf] == nil {
			s += msgp.NilSize
		} else {
			s += z.Bullets[zdaf].Msgsize()
		}
	}
	s += 10 + msgp.Int32Size + 13 + msgp.Int32Size + 9 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameUpdateAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zzpf uint32
	zzpf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zzpf > 0 {
		zzpf--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "stopFire":
			z.StopFire, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "tick":
			z.Tick, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "switchScene":
			z.SwitchScene, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "play":
			var zrfe uint32
			zrfe, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zrfe) {
				z.Players = (z.Players)[:zrfe]
			} else {
				z.Players = make([]*Player, zrfe)
			}
			for zsnv := range z.Players {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Players[zsnv] = nil
				} else {
					if z.Players[zsnv] == nil {
						z.Players[zsnv] = new(Player)
					}
					err = z.Players[zsnv].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "fishes":
			var zgmo uint32
			zgmo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Fishes) >= int(zgmo) {
				z.Fishes = (z.Fishes)[:zgmo]
			} else {
				z.Fishes = make([]*Fish, zgmo)
			}
			for zkgt := range z.Fishes {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Fishes[zkgt] = nil
				} else {
					if z.Fishes[zkgt] == nil {
						z.Fishes[zkgt] = new(Fish)
					}
					err = z.Fishes[zkgt].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "bull":
			var ztaf uint32
			ztaf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Bullets) >= int(ztaf) {
				z.Bullets = (z.Bullets)[:ztaf]
			} else {
				z.Bullets = make([]*Bullet, ztaf)
			}
			for zema := range z.Bullets {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Bullets[zema] = nil
				} else {
					if z.Bullets[zema] == nil {
						z.Bullets[zema] = new(Bullet)
					}
					err = z.Bullets[zema].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		case "dieBullets":
			var zeth uint32
			zeth, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.DieBullets) >= int(zeth) {
				z.DieBullets = (z.DieBullets)[:zeth]
			} else {
				z.DieBullets = make([]int32, zeth)
			}
			for zpez := range z.DieBullets {
				z.DieBullets[zpez], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "dieFishes":
			var zsbz uint32
			zsbz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.DieFishes) >= int(zsbz) {
				z.DieFishes = (z.DieFishes)[:zsbz]
			} else {
				z.DieFishes = make([]int32, zsbz)
			}
			for zqke := range z.DieFishes {
				z.DieFishes[zqke], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "kills":
			var zrjx uint32
			zrjx, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Kills) >= int(zrjx) {
				z.Kills = (z.Kills)[:zrjx]
			} else {
				z.Kills = make([]*KillFish, zrjx)
			}
			for zqyh := range z.Kills {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Kills[zqyh] = nil
				} else {
					if z.Kills[zqyh] == nil {
						z.Kills[zqyh] = new(KillFish)
					}
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
							z.Kills[zqyh].Id, err = dc.ReadInt32()
							if err != nil {
								return
							}
						case "uid":
							z.Kills[zqyh].Uid, err = dc.ReadInt32()
							if err != nil {
								return
							}
						case "score":
							z.Kills[zqyh].Score, err = dc.ReadInt64()
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
		case "seed":
			var zwel uint32
			zwel, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Seed) >= int(zwel) {
				z.Seed = (z.Seed)[:zwel]
			} else {
				z.Seed = make([]*FishSeed, zwel)
			}
			for zyzr := range z.Seed {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Seed[zyzr] = nil
				} else {
					if z.Seed[zyzr] == nil {
						z.Seed[zyzr] = new(FishSeed)
					}
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
							z.Seed[zyzr].Id, err = dc.ReadInt32()
							if err != nil {
								return
							}
						case "speed":
							z.Seed[zyzr].Speed, err = dc.ReadFloat64()
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
		case "offline":
			var zmfd uint32
			zmfd, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Offline) >= int(zmfd) {
				z.Offline = (z.Offline)[:zmfd]
			} else {
				z.Offline = make([]int32, zmfd)
			}
			for zywj := range z.Offline {
				z.Offline[zywj], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "describe":
			var zzdc uint32
			zzdc, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Describe) >= int(zzdc) {
				z.Describe = (z.Describe)[:zzdc]
			} else {
				z.Describe = make([]string, zzdc)
			}
			for zjpj := range z.Describe {
				z.Describe[zjpj], err = dc.ReadString()
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
func (z *GameUpdateAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 12
	// write "stopFire"
	err = en.Append(0x8c, 0xa8, 0x73, 0x74, 0x6f, 0x70, 0x46, 0x69, 0x72, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.StopFire)
	if err != nil {
		return
	}
	// write "tick"
	err = en.Append(0xa4, 0x74, 0x69, 0x63, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Tick)
	if err != nil {
		return
	}
	// write "switchScene"
	err = en.Append(0xab, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x53, 0x63, 0x65, 0x6e, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.SwitchScene)
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
	for zsnv := range z.Players {
		if z.Players[zsnv] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Players[zsnv].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "fishes"
	err = en.Append(0xa6, 0x66, 0x69, 0x73, 0x68, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Fishes)))
	if err != nil {
		return
	}
	for zkgt := range z.Fishes {
		if z.Fishes[zkgt] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Fishes[zkgt].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "bull"
	err = en.Append(0xa4, 0x62, 0x75, 0x6c, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Bullets)))
	if err != nil {
		return
	}
	for zema := range z.Bullets {
		if z.Bullets[zema] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Bullets[zema].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "dieBullets"
	err = en.Append(0xaa, 0x64, 0x69, 0x65, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.DieBullets)))
	if err != nil {
		return
	}
	for zpez := range z.DieBullets {
		err = en.WriteInt32(z.DieBullets[zpez])
		if err != nil {
			return
		}
	}
	// write "dieFishes"
	err = en.Append(0xa9, 0x64, 0x69, 0x65, 0x46, 0x69, 0x73, 0x68, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.DieFishes)))
	if err != nil {
		return
	}
	for zqke := range z.DieFishes {
		err = en.WriteInt32(z.DieFishes[zqke])
		if err != nil {
			return
		}
	}
	// write "kills"
	err = en.Append(0xa5, 0x6b, 0x69, 0x6c, 0x6c, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Kills)))
	if err != nil {
		return
	}
	for zqyh := range z.Kills {
		if z.Kills[zqyh] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 3
			// write "id"
			err = en.Append(0x83, 0xa2, 0x69, 0x64)
			if err != nil {
				return err
			}
			err = en.WriteInt32(z.Kills[zqyh].Id)
			if err != nil {
				return
			}
			// write "uid"
			err = en.Append(0xa3, 0x75, 0x69, 0x64)
			if err != nil {
				return err
			}
			err = en.WriteInt32(z.Kills[zqyh].Uid)
			if err != nil {
				return
			}
			// write "score"
			err = en.Append(0xa5, 0x73, 0x63, 0x6f, 0x72, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Kills[zqyh].Score)
			if err != nil {
				return
			}
		}
	}
	// write "seed"
	err = en.Append(0xa4, 0x73, 0x65, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Seed)))
	if err != nil {
		return
	}
	for zyzr := range z.Seed {
		if z.Seed[zyzr] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "id"
			err = en.Append(0x82, 0xa2, 0x69, 0x64)
			if err != nil {
				return err
			}
			err = en.WriteInt32(z.Seed[zyzr].Id)
			if err != nil {
				return
			}
			// write "speed"
			err = en.Append(0xa5, 0x73, 0x70, 0x65, 0x65, 0x64)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Seed[zyzr].Speed)
			if err != nil {
				return
			}
		}
	}
	// write "offline"
	err = en.Append(0xa7, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Offline)))
	if err != nil {
		return
	}
	for zywj := range z.Offline {
		err = en.WriteInt32(z.Offline[zywj])
		if err != nil {
			return
		}
	}
	// write "describe"
	err = en.Append(0xa8, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Describe)))
	if err != nil {
		return
	}
	for zjpj := range z.Describe {
		err = en.WriteString(z.Describe[zjpj])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameUpdateAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 12
	// string "stopFire"
	o = append(o, 0x8c, 0xa8, 0x73, 0x74, 0x6f, 0x70, 0x46, 0x69, 0x72, 0x65)
	o = msgp.AppendBool(o, z.StopFire)
	// string "tick"
	o = append(o, 0xa4, 0x74, 0x69, 0x63, 0x6b)
	o = msgp.AppendInt64(o, z.Tick)
	// string "switchScene"
	o = append(o, 0xab, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x53, 0x63, 0x65, 0x6e, 0x65)
	o = msgp.AppendInt32(o, z.SwitchScene)
	// string "play"
	o = append(o, 0xa4, 0x70, 0x6c, 0x61, 0x79)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Players)))
	for zsnv := range z.Players {
		if z.Players[zsnv] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Players[zsnv].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "fishes"
	o = append(o, 0xa6, 0x66, 0x69, 0x73, 0x68, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Fishes)))
	for zkgt := range z.Fishes {
		if z.Fishes[zkgt] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Fishes[zkgt].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "bull"
	o = append(o, 0xa4, 0x62, 0x75, 0x6c, 0x6c)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Bullets)))
	for zema := range z.Bullets {
		if z.Bullets[zema] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Bullets[zema].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "dieBullets"
	o = append(o, 0xaa, 0x64, 0x69, 0x65, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.DieBullets)))
	for zpez := range z.DieBullets {
		o = msgp.AppendInt32(o, z.DieBullets[zpez])
	}
	// string "dieFishes"
	o = append(o, 0xa9, 0x64, 0x69, 0x65, 0x46, 0x69, 0x73, 0x68, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.DieFishes)))
	for zqke := range z.DieFishes {
		o = msgp.AppendInt32(o, z.DieFishes[zqke])
	}
	// string "kills"
	o = append(o, 0xa5, 0x6b, 0x69, 0x6c, 0x6c, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Kills)))
	for zqyh := range z.Kills {
		if z.Kills[zqyh] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 3
			// string "id"
			o = append(o, 0x83, 0xa2, 0x69, 0x64)
			o = msgp.AppendInt32(o, z.Kills[zqyh].Id)
			// string "uid"
			o = append(o, 0xa3, 0x75, 0x69, 0x64)
			o = msgp.AppendInt32(o, z.Kills[zqyh].Uid)
			// string "score"
			o = append(o, 0xa5, 0x73, 0x63, 0x6f, 0x72, 0x65)
			o = msgp.AppendInt64(o, z.Kills[zqyh].Score)
		}
	}
	// string "seed"
	o = append(o, 0xa4, 0x73, 0x65, 0x65, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Seed)))
	for zyzr := range z.Seed {
		if z.Seed[zyzr] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "id"
			o = append(o, 0x82, 0xa2, 0x69, 0x64)
			o = msgp.AppendInt32(o, z.Seed[zyzr].Id)
			// string "speed"
			o = append(o, 0xa5, 0x73, 0x70, 0x65, 0x65, 0x64)
			o = msgp.AppendFloat64(o, z.Seed[zyzr].Speed)
		}
	}
	// string "offline"
	o = append(o, 0xa7, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Offline)))
	for zywj := range z.Offline {
		o = msgp.AppendInt32(o, z.Offline[zywj])
	}
	// string "describe"
	o = append(o, 0xa8, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Describe)))
	for zjpj := range z.Describe {
		o = msgp.AppendString(o, z.Describe[zjpj])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameUpdateAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "stopFire":
			z.StopFire, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "tick":
			z.Tick, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "switchScene":
			z.SwitchScene, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "play":
			var zbal uint32
			zbal, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Players) >= int(zbal) {
				z.Players = (z.Players)[:zbal]
			} else {
				z.Players = make([]*Player, zbal)
			}
			for zsnv := range z.Players {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Players[zsnv] = nil
				} else {
					if z.Players[zsnv] == nil {
						z.Players[zsnv] = new(Player)
					}
					bts, err = z.Players[zsnv].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "fishes":
			var zjqz uint32
			zjqz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Fishes) >= int(zjqz) {
				z.Fishes = (z.Fishes)[:zjqz]
			} else {
				z.Fishes = make([]*Fish, zjqz)
			}
			for zkgt := range z.Fishes {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Fishes[zkgt] = nil
				} else {
					if z.Fishes[zkgt] == nil {
						z.Fishes[zkgt] = new(Fish)
					}
					bts, err = z.Fishes[zkgt].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "bull":
			var zkct uint32
			zkct, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Bullets) >= int(zkct) {
				z.Bullets = (z.Bullets)[:zkct]
			} else {
				z.Bullets = make([]*Bullet, zkct)
			}
			for zema := range z.Bullets {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Bullets[zema] = nil
				} else {
					if z.Bullets[zema] == nil {
						z.Bullets[zema] = new(Bullet)
					}
					bts, err = z.Bullets[zema].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "dieBullets":
			var ztmt uint32
			ztmt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.DieBullets) >= int(ztmt) {
				z.DieBullets = (z.DieBullets)[:ztmt]
			} else {
				z.DieBullets = make([]int32, ztmt)
			}
			for zpez := range z.DieBullets {
				z.DieBullets[zpez], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "dieFishes":
			var ztco uint32
			ztco, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.DieFishes) >= int(ztco) {
				z.DieFishes = (z.DieFishes)[:ztco]
			} else {
				z.DieFishes = make([]int32, ztco)
			}
			for zqke := range z.DieFishes {
				z.DieFishes[zqke], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "kills":
			var zana uint32
			zana, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Kills) >= int(zana) {
				z.Kills = (z.Kills)[:zana]
			} else {
				z.Kills = make([]*KillFish, zana)
			}
			for zqyh := range z.Kills {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Kills[zqyh] = nil
				} else {
					if z.Kills[zqyh] == nil {
						z.Kills[zqyh] = new(KillFish)
					}
					var ztyy uint32
					ztyy, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for ztyy > 0 {
						ztyy--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "id":
							z.Kills[zqyh].Id, bts, err = msgp.ReadInt32Bytes(bts)
							if err != nil {
								return
							}
						case "uid":
							z.Kills[zqyh].Uid, bts, err = msgp.ReadInt32Bytes(bts)
							if err != nil {
								return
							}
						case "score":
							z.Kills[zqyh].Score, bts, err = msgp.ReadInt64Bytes(bts)
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
		case "seed":
			var zinl uint32
			zinl, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Seed) >= int(zinl) {
				z.Seed = (z.Seed)[:zinl]
			} else {
				z.Seed = make([]*FishSeed, zinl)
			}
			for zyzr := range z.Seed {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Seed[zyzr] = nil
				} else {
					if z.Seed[zyzr] == nil {
						z.Seed[zyzr] = new(FishSeed)
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
						case "id":
							z.Seed[zyzr].Id, bts, err = msgp.ReadInt32Bytes(bts)
							if err != nil {
								return
							}
						case "speed":
							z.Seed[zyzr].Speed, bts, err = msgp.ReadFloat64Bytes(bts)
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
		case "offline":
			var zljy uint32
			zljy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Offline) >= int(zljy) {
				z.Offline = (z.Offline)[:zljy]
			} else {
				z.Offline = make([]int32, zljy)
			}
			for zywj := range z.Offline {
				z.Offline[zywj], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "describe":
			var zixj uint32
			zixj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Describe) >= int(zixj) {
				z.Describe = (z.Describe)[:zixj]
			} else {
				z.Describe = make([]string, zixj)
			}
			for zjpj := range z.Describe {
				z.Describe[zjpj], bts, err = msgp.ReadStringBytes(bts)
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
func (z *GameUpdateAck) Msgsize() (s int) {
	s = 1 + 9 + msgp.BoolSize + 5 + msgp.Int64Size + 12 + msgp.Int32Size + 5 + msgp.ArrayHeaderSize
	for zsnv := range z.Players {
		if z.Players[zsnv] == nil {
			s += msgp.NilSize
		} else {
			s += z.Players[zsnv].Msgsize()
		}
	}
	s += 7 + msgp.ArrayHeaderSize
	for zkgt := range z.Fishes {
		if z.Fishes[zkgt] == nil {
			s += msgp.NilSize
		} else {
			s += z.Fishes[zkgt].Msgsize()
		}
	}
	s += 5 + msgp.ArrayHeaderSize
	for zema := range z.Bullets {
		if z.Bullets[zema] == nil {
			s += msgp.NilSize
		} else {
			s += z.Bullets[zema].Msgsize()
		}
	}
	s += 11 + msgp.ArrayHeaderSize + (len(z.DieBullets) * (msgp.Int32Size)) + 10 + msgp.ArrayHeaderSize + (len(z.DieFishes) * (msgp.Int32Size)) + 6 + msgp.ArrayHeaderSize
	for zqyh := range z.Kills {
		if z.Kills[zqyh] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 3 + msgp.Int32Size + 4 + msgp.Int32Size + 6 + msgp.Int64Size
		}
	}
	s += 5 + msgp.ArrayHeaderSize
	for zyzr := range z.Seed {
		if z.Seed[zyzr] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 3 + msgp.Int32Size + 6 + msgp.Float64Size
		}
	}
	s += 8 + msgp.ArrayHeaderSize + (len(z.Offline) * (msgp.Int32Size)) + 9 + msgp.ArrayHeaderSize
	for zjpj := range z.Describe {
		s += msgp.StringPrefixSize + len(z.Describe[zjpj])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *HitReq) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrsc uint32
	zrsc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrsc > 0 {
		zrsc--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "bulletId":
			z.BulletId, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "fishId":
			z.FishId, err = dc.ReadInt32()
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
func (z HitReq) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "bulletId"
	err = en.Append(0x82, 0xa8, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.BulletId)
	if err != nil {
		return
	}
	// write "fishId"
	err = en.Append(0xa6, 0x66, 0x69, 0x73, 0x68, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.FishId)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z HitReq) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "bulletId"
	o = append(o, 0x82, 0xa8, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64)
	o = msgp.AppendInt32(o, z.BulletId)
	// string "fishId"
	o = append(o, 0xa6, 0x66, 0x69, 0x73, 0x68, 0x49, 0x64)
	o = msgp.AppendInt32(o, z.FishId)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HitReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "bulletId":
			z.BulletId, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "fishId":
			z.FishId, bts, err = msgp.ReadInt32Bytes(bts)
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
func (z HitReq) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int32Size + 7 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *KillFish) DecodeMsg(dc *msgp.Reader) (err error) {
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
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "uid":
			z.Uid, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "score":
			z.Score, err = dc.ReadInt64()
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
func (z KillFish) EncodeMsg(en *msgp.Writer) (err error) {
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
	// write "uid"
	err = en.Append(0xa3, 0x75, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Uid)
	if err != nil {
		return
	}
	// write "score"
	err = en.Append(0xa5, 0x73, 0x63, 0x6f, 0x72, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Score)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z KillFish) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "id"
	o = append(o, 0x83, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "uid"
	o = append(o, 0xa3, 0x75, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Uid)
	// string "score"
	o = append(o, 0xa5, 0x73, 0x63, 0x6f, 0x72, 0x65)
	o = msgp.AppendInt64(o, z.Score)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *KillFish) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var znsg uint32
	znsg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for znsg > 0 {
		znsg--
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
		case "uid":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "score":
			z.Score, bts, err = msgp.ReadInt64Bytes(bts)
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
func (z KillFish) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 4 + msgp.Int32Size + 6 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Player) DecodeMsg(dc *msgp.Reader) (err error) {
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
	// map header, size 6
	// string "id"
	o = append(o, 0x86, 0xa2, 0x69, 0x64)
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
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Player) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
	s = 1 + 3 + msgp.Int32Size + 5 + msgp.Int32Size + 4 + msgp.Int32Size + 6 + msgp.Int32Size + 5 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ShootReq) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "fish":
			z.Fish, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "client":
			z.Client, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "bet":
			z.Bet, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "direction":
			z.Direction, err = dc.ReadFloat64()
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
func (z *ShootReq) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "fish"
	err = en.Append(0x84, 0xa4, 0x66, 0x69, 0x73, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Fish)
	if err != nil {
		return
	}
	// write "client"
	err = en.Append(0xa6, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Client)
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
	// write "direction"
	err = en.Append(0xa9, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Direction)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ShootReq) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "fish"
	o = append(o, 0x84, 0xa4, 0x66, 0x69, 0x73, 0x68)
	o = msgp.AppendInt32(o, z.Fish)
	// string "client"
	o = append(o, 0xa6, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74)
	o = msgp.AppendInt32(o, z.Client)
	// string "bet"
	o = append(o, 0xa3, 0x62, 0x65, 0x74)
	o = msgp.AppendInt32(o, z.Bet)
	// string "direction"
	o = append(o, 0xa9, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendFloat64(o, z.Direction)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ShootReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "fish":
			z.Fish, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "client":
			z.Client, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "bet":
			z.Bet, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "direction":
			z.Direction, bts, err = msgp.ReadFloat64Bytes(bts)
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
func (z *ShootReq) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int32Size + 7 + msgp.Int32Size + 4 + msgp.Int32Size + 10 + msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TimeSyncAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsbo uint32
	zsbo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsbo > 0 {
		zsbo--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Client":
			z.Client, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Server":
			z.Server, err = dc.ReadInt64()
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
func (z TimeSyncAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Client"
	err = en.Append(0x82, 0xa6, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Client)
	if err != nil {
		return
	}
	// write "Server"
	err = en.Append(0xa6, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Server)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TimeSyncAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Client"
	o = append(o, 0x82, 0xa6, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74)
	o = msgp.AppendInt64(o, z.Client)
	// string "Server"
	o = append(o, 0xa6, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72)
	o = msgp.AppendInt64(o, z.Server)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TimeSyncAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Client":
			z.Client, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Server":
			z.Server, bts, err = msgp.ReadInt64Bytes(bts)
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
func (z TimeSyncAck) Msgsize() (s int) {
	s = 1 + 7 + msgp.Int64Size + 7 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TimeSyncReq) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zqgz uint32
	zqgz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zqgz > 0 {
		zqgz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Tick":
			z.Tick, err = dc.ReadInt64()
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
func (z TimeSyncReq) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Tick"
	err = en.Append(0x81, 0xa4, 0x54, 0x69, 0x63, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Tick)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TimeSyncReq) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Tick"
	o = append(o, 0x81, 0xa4, 0x54, 0x69, 0x63, 0x6b)
	o = msgp.AppendInt64(o, z.Tick)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TimeSyncReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Tick":
			z.Tick, bts, err = msgp.ReadInt64Bytes(bts)
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
func (z TimeSyncReq) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int64Size
	return
}
