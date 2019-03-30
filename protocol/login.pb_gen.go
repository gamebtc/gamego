package protocol

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DeviceInfo) DecodeMsg(dc *msgp.Reader) (err error) {
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
			z.Id, err = dc.ReadString()
			if err != nil {
				return
			}
		case "vend":
			z.Vend, err = dc.ReadString()
			if err != nil {
				return
			}
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "mac":
			z.Mac, err = dc.ReadString()
			if err != nil {
				return
			}
		case "imei":
			z.Imei, err = dc.ReadString()
			if err != nil {
				return
			}
		case "emid":
			z.Emid, err = dc.ReadString()
			if err != nil {
				return
			}
		case "sn":
			z.Sn, err = dc.ReadString()
			if err != nil {
				return
			}
		case "osLang":
			z.OsLang, err = dc.ReadString()
			if err != nil {
				return
			}
		case "osVer":
			z.OsVer, err = dc.ReadString()
			if err != nil {
				return
			}
		case "other":
			z.Other, err = dc.ReadString()
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
func (z *DeviceInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "id"
	err = en.Append(0x8a, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Id)
	if err != nil {
		return
	}
	// write "vend"
	err = en.Append(0xa4, 0x76, 0x65, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Vend)
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
	// write "mac"
	err = en.Append(0xa3, 0x6d, 0x61, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Mac)
	if err != nil {
		return
	}
	// write "imei"
	err = en.Append(0xa4, 0x69, 0x6d, 0x65, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Imei)
	if err != nil {
		return
	}
	// write "emid"
	err = en.Append(0xa4, 0x65, 0x6d, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Emid)
	if err != nil {
		return
	}
	// write "sn"
	err = en.Append(0xa2, 0x73, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Sn)
	if err != nil {
		return
	}
	// write "osLang"
	err = en.Append(0xa6, 0x6f, 0x73, 0x4c, 0x61, 0x6e, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteString(z.OsLang)
	if err != nil {
		return
	}
	// write "osVer"
	err = en.Append(0xa5, 0x6f, 0x73, 0x56, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.OsVer)
	if err != nil {
		return
	}
	// write "other"
	err = en.Append(0xa5, 0x6f, 0x74, 0x68, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Other)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DeviceInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "id"
	o = append(o, 0x8a, 0xa2, 0x69, 0x64)
	o = msgp.AppendString(o, z.Id)
	// string "vend"
	o = append(o, 0xa4, 0x76, 0x65, 0x6e, 0x64)
	o = msgp.AppendString(o, z.Vend)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "mac"
	o = append(o, 0xa3, 0x6d, 0x61, 0x63)
	o = msgp.AppendString(o, z.Mac)
	// string "imei"
	o = append(o, 0xa4, 0x69, 0x6d, 0x65, 0x69)
	o = msgp.AppendString(o, z.Imei)
	// string "emid"
	o = append(o, 0xa4, 0x65, 0x6d, 0x69, 0x64)
	o = msgp.AppendString(o, z.Emid)
	// string "sn"
	o = append(o, 0xa2, 0x73, 0x6e)
	o = msgp.AppendString(o, z.Sn)
	// string "osLang"
	o = append(o, 0xa6, 0x6f, 0x73, 0x4c, 0x61, 0x6e, 0x67)
	o = msgp.AppendString(o, z.OsLang)
	// string "osVer"
	o = append(o, 0xa5, 0x6f, 0x73, 0x56, 0x65, 0x72)
	o = msgp.AppendString(o, z.OsVer)
	// string "other"
	o = append(o, 0xa5, 0x6f, 0x74, 0x68, 0x65, 0x72)
	o = msgp.AppendString(o, z.Other)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DeviceInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
			z.Id, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "vend":
			z.Vend, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "mac":
			z.Mac, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "imei":
			z.Imei, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "emid":
			z.Emid, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "sn":
			z.Sn, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "osLang":
			z.OsLang, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "osVer":
			z.OsVer, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "other":
			z.Other, bts, err = msgp.ReadStringBytes(bts)
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
func (z *DeviceInfo) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.Id) + 5 + msgp.StringPrefixSize + len(z.Vend) + 5 + msgp.StringPrefixSize + len(z.Name) + 4 + msgp.StringPrefixSize + len(z.Mac) + 5 + msgp.StringPrefixSize + len(z.Imei) + 5 + msgp.StringPrefixSize + len(z.Emid) + 3 + msgp.StringPrefixSize + len(z.Sn) + 7 + msgp.StringPrefixSize + len(z.OsLang) + 6 + msgp.StringPrefixSize + len(z.OsVer) + 6 + msgp.StringPrefixSize + len(z.Other)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Envirnment) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "id":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "pack":
			z.Pack, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "ver":
			z.Ver, err = dc.ReadString()
			if err != nil {
				return
			}
		case "chan":
			z.Chan, err = dc.ReadString()
			if err != nil {
				return
			}
		case "refer":
			z.Refer, err = dc.ReadString()
			if err != nil {
				return
			}
		case "other":
			z.Other, err = dc.ReadString()
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
func (z *Envirnment) EncodeMsg(en *msgp.Writer) (err error) {
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
	// write "pack"
	err = en.Append(0xa4, 0x70, 0x61, 0x63, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Pack)
	if err != nil {
		return
	}
	// write "ver"
	err = en.Append(0xa3, 0x76, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Ver)
	if err != nil {
		return
	}
	// write "chan"
	err = en.Append(0xa4, 0x63, 0x68, 0x61, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Chan)
	if err != nil {
		return
	}
	// write "refer"
	err = en.Append(0xa5, 0x72, 0x65, 0x66, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Refer)
	if err != nil {
		return
	}
	// write "other"
	err = en.Append(0xa5, 0x6f, 0x74, 0x68, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Other)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Envirnment) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "id"
	o = append(o, 0x86, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "pack"
	o = append(o, 0xa4, 0x70, 0x61, 0x63, 0x6b)
	o = msgp.AppendInt32(o, z.Pack)
	// string "ver"
	o = append(o, 0xa3, 0x76, 0x65, 0x72)
	o = msgp.AppendString(o, z.Ver)
	// string "chan"
	o = append(o, 0xa4, 0x63, 0x68, 0x61, 0x6e)
	o = msgp.AppendString(o, z.Chan)
	// string "refer"
	o = append(o, 0xa5, 0x72, 0x65, 0x66, 0x65, 0x72)
	o = msgp.AppendString(o, z.Refer)
	// string "other"
	o = append(o, 0xa5, 0x6f, 0x74, 0x68, 0x65, 0x72)
	o = msgp.AppendString(o, z.Other)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Envirnment) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "id":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "pack":
			z.Pack, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "ver":
			z.Ver, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "chan":
			z.Chan, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "refer":
			z.Refer, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "other":
			z.Other, bts, err = msgp.ReadStringBytes(bts)
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
func (z *Envirnment) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 5 + msgp.Int32Size + 4 + msgp.StringPrefixSize + len(z.Ver) + 5 + msgp.StringPrefixSize + len(z.Chan) + 6 + msgp.StringPrefixSize + len(z.Refer) + 6 + msgp.StringPrefixSize + len(z.Other)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ErrorInfo) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "reqId":
			z.ReqId, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "code":
			z.Code, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "msg":
			z.Msg, err = dc.ReadString()
			if err != nil {
				return
			}
		case "key":
			z.Key, err = dc.ReadString()
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
func (z *ErrorInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "reqId"
	err = en.Append(0x84, 0xa5, 0x72, 0x65, 0x71, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.ReqId)
	if err != nil {
		return
	}
	// write "code"
	err = en.Append(0xa4, 0x63, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Code)
	if err != nil {
		return
	}
	// write "msg"
	err = en.Append(0xa3, 0x6d, 0x73, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Msg)
	if err != nil {
		return
	}
	// write "key"
	err = en.Append(0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Key)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ErrorInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "reqId"
	o = append(o, 0x84, 0xa5, 0x72, 0x65, 0x71, 0x49, 0x64)
	o = msgp.AppendInt32(o, z.ReqId)
	// string "code"
	o = append(o, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt32(o, z.Code)
	// string "msg"
	o = append(o, 0xa3, 0x6d, 0x73, 0x67)
	o = msgp.AppendString(o, z.Msg)
	// string "key"
	o = append(o, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ErrorInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "reqId":
			z.ReqId, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "code":
			z.Code, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "msg":
			z.Msg, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
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
func (z *ErrorInfo) Msgsize() (s int) {
	s = 1 + 6 + msgp.Int32Size + 5 + msgp.Int32Size + 4 + msgp.StringPrefixSize + len(z.Msg) + 4 + msgp.StringPrefixSize + len(z.Key)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FatalInfo) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "reqId":
			z.ReqId, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "code":
			z.Code, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "msg":
			z.Msg, err = dc.ReadString()
			if err != nil {
				return
			}
		case "key":
			z.Key, err = dc.ReadString()
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
func (z *FatalInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "reqId"
	err = en.Append(0x84, 0xa5, 0x72, 0x65, 0x71, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.ReqId)
	if err != nil {
		return
	}
	// write "code"
	err = en.Append(0xa4, 0x63, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Code)
	if err != nil {
		return
	}
	// write "msg"
	err = en.Append(0xa3, 0x6d, 0x73, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Msg)
	if err != nil {
		return
	}
	// write "key"
	err = en.Append(0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Key)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FatalInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "reqId"
	o = append(o, 0x84, 0xa5, 0x72, 0x65, 0x71, 0x49, 0x64)
	o = msgp.AppendInt32(o, z.ReqId)
	// string "code"
	o = append(o, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt32(o, z.Code)
	// string "msg"
	o = append(o, 0xa3, 0x6d, 0x73, 0x67)
	o = msgp.AppendString(o, z.Msg)
	// string "key"
	o = append(o, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FatalInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "reqId":
			z.ReqId, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "code":
			z.Code, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "msg":
			z.Msg, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
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
func (z *FatalInfo) Msgsize() (s int) {
	s = 1 + 6 + msgp.Int32Size + 5 + msgp.Int32Size + 4 + msgp.StringPrefixSize + len(z.Msg) + 4 + msgp.StringPrefixSize + len(z.Key)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Handshake) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "code":
			z.Code, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "seed":
			z.Seed, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "msg":
			z.Msg, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ip":
			var zdaf uint32
			zdaf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Ip) >= int(zdaf) {
				z.Ip = (z.Ip)[:zdaf]
			} else {
				z.Ip = make([]string, zdaf)
			}
			for zxhx := range z.Ip {
				z.Ip[zxhx], err = dc.ReadString()
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
func (z *Handshake) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "code"
	err = en.Append(0x84, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Code)
	if err != nil {
		return
	}
	// write "seed"
	err = en.Append(0xa4, 0x73, 0x65, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Seed)
	if err != nil {
		return
	}
	// write "msg"
	err = en.Append(0xa3, 0x6d, 0x73, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Msg)
	if err != nil {
		return
	}
	// write "ip"
	err = en.Append(0xa2, 0x69, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Ip)))
	if err != nil {
		return
	}
	for zxhx := range z.Ip {
		err = en.WriteString(z.Ip[zxhx])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Handshake) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "code"
	o = append(o, 0x84, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt32(o, z.Code)
	// string "seed"
	o = append(o, 0xa4, 0x73, 0x65, 0x65, 0x64)
	o = msgp.AppendInt32(o, z.Seed)
	// string "msg"
	o = append(o, 0xa3, 0x6d, 0x73, 0x67)
	o = msgp.AppendString(o, z.Msg)
	// string "ip"
	o = append(o, 0xa2, 0x69, 0x70)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Ip)))
	for zxhx := range z.Ip {
		o = msgp.AppendString(o, z.Ip[zxhx])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Handshake) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "code":
			z.Code, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "seed":
			z.Seed, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "msg":
			z.Msg, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ip":
			var zjfb uint32
			zjfb, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Ip) >= int(zjfb) {
				z.Ip = (z.Ip)[:zjfb]
			} else {
				z.Ip = make([]string, zjfb)
			}
			for zxhx := range z.Ip {
				z.Ip[zxhx], bts, err = msgp.ReadStringBytes(bts)
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
func (z *Handshake) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int32Size + 5 + msgp.Int32Size + 4 + msgp.StringPrefixSize + len(z.Msg) + 3 + msgp.ArrayHeaderSize
	for zxhx := range z.Ip {
		s += msgp.StringPrefixSize + len(z.Ip[zxhx])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *HeartBeatAck) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcxo uint32
	zcxo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcxo > 0 {
		zcxo--
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
func (z HeartBeatAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "id"
	err = en.Append(0x81, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z HeartBeatAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "id"
	o = append(o, 0x81, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HeartBeatAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zeff uint32
	zeff, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zeff > 0 {
		zeff--
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
func (z HeartBeatAck) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *HeartBeatReq) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "id":
			z.Id, err = dc.ReadInt32()
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
func (z HeartBeatReq) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "id"
	err = en.Append(0x81, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z HeartBeatReq) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "id"
	o = append(o, 0x81, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HeartBeatReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
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
func (z HeartBeatReq) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LoginFailAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "code":
			z.Code, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "msg":
			z.Msg, err = dc.ReadString()
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
func (z LoginFailAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "code"
	err = en.Append(0x82, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Code)
	if err != nil {
		return
	}
	// write "msg"
	err = en.Append(0xa3, 0x6d, 0x73, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Msg)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z LoginFailAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "code"
	o = append(o, 0x82, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt32(o, z.Code)
	// string "msg"
	o = append(o, 0xa3, 0x6d, 0x73, 0x67)
	o = msgp.AppendString(o, z.Msg)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LoginFailAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "code":
			z.Code, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "msg":
			z.Msg, bts, err = msgp.ReadStringBytes(bts)
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
func (z LoginFailAck) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int32Size + 4 + msgp.StringPrefixSize + len(z.Msg)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LoginReq) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "type":
			z.Type, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "pwd":
			z.Pwd, err = dc.ReadString()
			if err != nil {
				return
			}
		case "udid":
			z.Udid, err = dc.ReadString()
			if err != nil {
				return
			}
		case "uid":
			z.Uid, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "time":
			z.Time, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "dev":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Dev = nil
			} else {
				if z.Dev == nil {
					z.Dev = new(DeviceInfo)
				}
				err = z.Dev.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "env":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Env = nil
			} else {
				if z.Env == nil {
					z.Env = new(Envirnment)
				}
				err = z.Env.DecodeMsg(dc)
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
func (z *LoginReq) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 8
	// write "type"
	err = en.Append(0x88, 0xa4, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Type)
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
	// write "pwd"
	err = en.Append(0xa3, 0x70, 0x77, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Pwd)
	if err != nil {
		return
	}
	// write "udid"
	err = en.Append(0xa4, 0x75, 0x64, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Udid)
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
	// write "time"
	err = en.Append(0xa4, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Time)
	if err != nil {
		return
	}
	// write "dev"
	err = en.Append(0xa3, 0x64, 0x65, 0x76)
	if err != nil {
		return err
	}
	if z.Dev == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Dev.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "env"
	err = en.Append(0xa3, 0x65, 0x6e, 0x76)
	if err != nil {
		return err
	}
	if z.Env == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Env.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *LoginReq) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "type"
	o = append(o, 0x88, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendInt32(o, z.Type)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "pwd"
	o = append(o, 0xa3, 0x70, 0x77, 0x64)
	o = msgp.AppendString(o, z.Pwd)
	// string "udid"
	o = append(o, 0xa4, 0x75, 0x64, 0x69, 0x64)
	o = msgp.AppendString(o, z.Udid)
	// string "uid"
	o = append(o, 0xa3, 0x75, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Uid)
	// string "time"
	o = append(o, 0xa4, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt32(o, z.Time)
	// string "dev"
	o = append(o, 0xa3, 0x64, 0x65, 0x76)
	if z.Dev == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Dev.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "env"
	o = append(o, 0xa3, 0x65, 0x6e, 0x76)
	if z.Env == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Env.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LoginReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zkgt uint32
	zkgt, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zkgt > 0 {
		zkgt--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "type":
			z.Type, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "pwd":
			z.Pwd, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "udid":
			z.Udid, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "uid":
			z.Uid, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "time":
			z.Time, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "dev":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Dev = nil
			} else {
				if z.Dev == nil {
					z.Dev = new(DeviceInfo)
				}
				bts, err = z.Dev.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "env":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Env = nil
			} else {
				if z.Env == nil {
					z.Env = new(Envirnment)
				}
				bts, err = z.Env.UnmarshalMsg(bts)
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
func (z *LoginReq) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int32Size + 5 + msgp.StringPrefixSize + len(z.Name) + 4 + msgp.StringPrefixSize + len(z.Pwd) + 5 + msgp.StringPrefixSize + len(z.Udid) + 4 + msgp.Int32Size + 5 + msgp.Int32Size + 4
	if z.Dev == nil {
		s += msgp.NilSize
	} else {
		s += z.Dev.Msgsize()
	}
	s += 4
	if z.Env == nil {
		s += msgp.NilSize
	} else {
		s += z.Env.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LoginSuccessAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
			z.Id, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "agent":
			z.Agent, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "icon":
			z.Icon, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "sex":
			z.Sex, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "vip":
			z.Vip, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "act":
			z.Act, err = dc.ReadString()
			if err != nil {
				return
			}
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "phone":
			z.Phone, err = dc.ReadString()
			if err != nil {
				return
			}
		case "bag":
			var zqyh uint32
			zqyh, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Bag == nil && zqyh > 0 {
				z.Bag = make(map[string]int64, zqyh)
			} else if len(z.Bag) > 0 {
				for key, _ := range z.Bag {
					delete(z.Bag, key)
				}
			}
			for zqyh > 0 {
				zqyh--
				var zema string
				var zpez int64
				zema, err = dc.ReadString()
				if err != nil {
					return
				}
				zpez, err = dc.ReadInt64()
				if err != nil {
					return
				}
				z.Bag[zema] = zpez
			}
		case "kind":
			z.Kind, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "level":
			z.Level, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "room":
			z.Room, err = dc.ReadInt32()
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
func (z *LoginSuccessAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 12
	// write "id"
	err = en.Append(0x8c, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		return
	}
	// write "agent"
	err = en.Append(0xa5, 0x61, 0x67, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Agent)
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
	// write "sex"
	err = en.Append(0xa3, 0x73, 0x65, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Sex)
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
	// write "act"
	err = en.Append(0xa3, 0x61, 0x63, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Act)
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
	// write "phone"
	err = en.Append(0xa5, 0x70, 0x68, 0x6f, 0x6e, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Phone)
	if err != nil {
		return
	}
	// write "bag"
	err = en.Append(0xa3, 0x62, 0x61, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Bag)))
	if err != nil {
		return
	}
	for zema, zpez := range z.Bag {
		err = en.WriteString(zema)
		if err != nil {
			return
		}
		err = en.WriteInt64(zpez)
		if err != nil {
			return
		}
	}
	// write "kind"
	err = en.Append(0xa4, 0x6b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Kind)
	if err != nil {
		return
	}
	// write "level"
	err = en.Append(0xa5, 0x6c, 0x65, 0x76, 0x65, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Level)
	if err != nil {
		return
	}
	// write "room"
	err = en.Append(0xa4, 0x72, 0x6f, 0x6f, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Room)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *LoginSuccessAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 12
	// string "id"
	o = append(o, 0x8c, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt32(o, z.Id)
	// string "agent"
	o = append(o, 0xa5, 0x61, 0x67, 0x65, 0x6e, 0x74)
	o = msgp.AppendInt64(o, z.Agent)
	// string "icon"
	o = append(o, 0xa4, 0x69, 0x63, 0x6f, 0x6e)
	o = msgp.AppendInt32(o, z.Icon)
	// string "sex"
	o = append(o, 0xa3, 0x73, 0x65, 0x78)
	o = msgp.AppendInt32(o, z.Sex)
	// string "vip"
	o = append(o, 0xa3, 0x76, 0x69, 0x70)
	o = msgp.AppendInt32(o, z.Vip)
	// string "act"
	o = append(o, 0xa3, 0x61, 0x63, 0x74)
	o = msgp.AppendString(o, z.Act)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "phone"
	o = append(o, 0xa5, 0x70, 0x68, 0x6f, 0x6e, 0x65)
	o = msgp.AppendString(o, z.Phone)
	// string "bag"
	o = append(o, 0xa3, 0x62, 0x61, 0x67)
	o = msgp.AppendMapHeader(o, uint32(len(z.Bag)))
	for zema, zpez := range z.Bag {
		o = msgp.AppendString(o, zema)
		o = msgp.AppendInt64(o, zpez)
	}
	// string "kind"
	o = append(o, 0xa4, 0x6b, 0x69, 0x6e, 0x64)
	o = msgp.AppendInt32(o, z.Kind)
	// string "level"
	o = append(o, 0xa5, 0x6c, 0x65, 0x76, 0x65, 0x6c)
	o = msgp.AppendInt32(o, z.Level)
	// string "room"
	o = append(o, 0xa4, 0x72, 0x6f, 0x6f, 0x6d)
	o = msgp.AppendInt32(o, z.Room)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LoginSuccessAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zyzr uint32
	zyzr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zyzr > 0 {
		zyzr--
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
		case "agent":
			z.Agent, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "icon":
			z.Icon, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "sex":
			z.Sex, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "vip":
			z.Vip, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "act":
			z.Act, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "phone":
			z.Phone, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "bag":
			var zywj uint32
			zywj, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Bag == nil && zywj > 0 {
				z.Bag = make(map[string]int64, zywj)
			} else if len(z.Bag) > 0 {
				for key, _ := range z.Bag {
					delete(z.Bag, key)
				}
			}
			for zywj > 0 {
				var zema string
				var zpez int64
				zywj--
				zema, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zpez, bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
				z.Bag[zema] = zpez
			}
		case "kind":
			z.Kind, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "level":
			z.Level, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "room":
			z.Room, bts, err = msgp.ReadInt32Bytes(bts)
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
func (z *LoginSuccessAck) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 6 + msgp.Int64Size + 5 + msgp.Int32Size + 4 + msgp.Int32Size + 4 + msgp.Int32Size + 4 + msgp.StringPrefixSize + len(z.Act) + 5 + msgp.StringPrefixSize + len(z.Name) + 6 + msgp.StringPrefixSize + len(z.Phone) + 4 + msgp.MapHeaderSize
	if z.Bag != nil {
		for zema, zpez := range z.Bag {
			_ = zpez
			s += msgp.StringPrefixSize + len(zema) + msgp.Int64Size
		}
	}
	s += 5 + msgp.Int32Size + 6 + msgp.Int32Size + 5 + msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VerCheckAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "code":
			z.Code, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "msg":
			z.Msg, err = dc.ReadString()
			if err != nil {
				return
			}
		case "canReg":
			z.CanReg, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "canLogin":
			z.CanLogin, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "url":
			z.Url, err = dc.ReadString()
			if err != nil {
				return
			}
		case "country":
			z.Country, err = dc.ReadString()
			if err != nil {
				return
			}
		case "region":
			z.Region, err = dc.ReadString()
			if err != nil {
				return
			}
		case "city":
			z.City, err = dc.ReadString()
			if err != nil {
				return
			}
		case "conf":
			var zgmo uint32
			zgmo, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Conf == nil && zgmo > 0 {
				z.Conf = make(map[string]string, zgmo)
			} else if len(z.Conf) > 0 {
				for key, _ := range z.Conf {
					delete(z.Conf, key)
				}
			}
			for zgmo > 0 {
				zgmo--
				var zjpj string
				var zzpf string
				zjpj, err = dc.ReadString()
				if err != nil {
					return
				}
				zzpf, err = dc.ReadString()
				if err != nil {
					return
				}
				z.Conf[zjpj] = zzpf
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
func (z *VerCheckAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "code"
	err = en.Append(0x89, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Code)
	if err != nil {
		return
	}
	// write "msg"
	err = en.Append(0xa3, 0x6d, 0x73, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Msg)
	if err != nil {
		return
	}
	// write "canReg"
	err = en.Append(0xa6, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.CanReg)
	if err != nil {
		return
	}
	// write "canLogin"
	err = en.Append(0xa8, 0x63, 0x61, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.CanLogin)
	if err != nil {
		return
	}
	// write "url"
	err = en.Append(0xa3, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Url)
	if err != nil {
		return
	}
	// write "country"
	err = en.Append(0xa7, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Country)
	if err != nil {
		return
	}
	// write "region"
	err = en.Append(0xa6, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Region)
	if err != nil {
		return
	}
	// write "city"
	err = en.Append(0xa4, 0x63, 0x69, 0x74, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.City)
	if err != nil {
		return
	}
	// write "conf"
	err = en.Append(0xa4, 0x63, 0x6f, 0x6e, 0x66)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Conf)))
	if err != nil {
		return
	}
	for zjpj, zzpf := range z.Conf {
		err = en.WriteString(zjpj)
		if err != nil {
			return
		}
		err = en.WriteString(zzpf)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VerCheckAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "code"
	o = append(o, 0x89, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt32(o, z.Code)
	// string "msg"
	o = append(o, 0xa3, 0x6d, 0x73, 0x67)
	o = msgp.AppendString(o, z.Msg)
	// string "canReg"
	o = append(o, 0xa6, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x67)
	o = msgp.AppendInt32(o, z.CanReg)
	// string "canLogin"
	o = append(o, 0xa8, 0x63, 0x61, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e)
	o = msgp.AppendInt32(o, z.CanLogin)
	// string "url"
	o = append(o, 0xa3, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.Url)
	// string "country"
	o = append(o, 0xa7, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79)
	o = msgp.AppendString(o, z.Country)
	// string "region"
	o = append(o, 0xa6, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Region)
	// string "city"
	o = append(o, 0xa4, 0x63, 0x69, 0x74, 0x79)
	o = msgp.AppendString(o, z.City)
	// string "conf"
	o = append(o, 0xa4, 0x63, 0x6f, 0x6e, 0x66)
	o = msgp.AppendMapHeader(o, uint32(len(z.Conf)))
	for zjpj, zzpf := range z.Conf {
		o = msgp.AppendString(o, zjpj)
		o = msgp.AppendString(o, zzpf)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VerCheckAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "code":
			z.Code, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "msg":
			z.Msg, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "canReg":
			z.CanReg, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "canLogin":
			z.CanLogin, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "url":
			z.Url, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "country":
			z.Country, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "region":
			z.Region, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "city":
			z.City, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "conf":
			var zeth uint32
			zeth, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Conf == nil && zeth > 0 {
				z.Conf = make(map[string]string, zeth)
			} else if len(z.Conf) > 0 {
				for key, _ := range z.Conf {
					delete(z.Conf, key)
				}
			}
			for zeth > 0 {
				var zjpj string
				var zzpf string
				zeth--
				zjpj, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zzpf, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Conf[zjpj] = zzpf
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
func (z *VerCheckAck) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int32Size + 4 + msgp.StringPrefixSize + len(z.Msg) + 7 + msgp.Int32Size + 9 + msgp.Int32Size + 4 + msgp.StringPrefixSize + len(z.Url) + 8 + msgp.StringPrefixSize + len(z.Country) + 7 + msgp.StringPrefixSize + len(z.Region) + 5 + msgp.StringPrefixSize + len(z.City) + 5 + msgp.MapHeaderSize
	if z.Conf != nil {
		for zjpj, zzpf := range z.Conf {
			_ = zzpf
			s += msgp.StringPrefixSize + len(zjpj) + msgp.StringPrefixSize + len(zzpf)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VerCheckReq) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsbz uint32
	zsbz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsbz > 0 {
		zsbz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "env":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Env = nil
			} else {
				if z.Env == nil {
					z.Env = new(Envirnment)
				}
				err = z.Env.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "time":
			z.Time, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "check":
			z.Check, err = dc.ReadInt32()
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
func (z *VerCheckReq) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "env"
	err = en.Append(0x83, 0xa3, 0x65, 0x6e, 0x76)
	if err != nil {
		return err
	}
	if z.Env == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Env.EncodeMsg(en)
		if err != nil {
			return
		}
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
	// write "check"
	err = en.Append(0xa5, 0x63, 0x68, 0x65, 0x63, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Check)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VerCheckReq) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "env"
	o = append(o, 0x83, 0xa3, 0x65, 0x6e, 0x76)
	if z.Env == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Env.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "time"
	o = append(o, 0xa4, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt32(o, z.Time)
	// string "check"
	o = append(o, 0xa5, 0x63, 0x68, 0x65, 0x63, 0x6b)
	o = msgp.AppendInt32(o, z.Check)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VerCheckReq) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "env":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Env = nil
			} else {
				if z.Env == nil {
					z.Env = new(Envirnment)
				}
				bts, err = z.Env.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "time":
			z.Time, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "check":
			z.Check, bts, err = msgp.ReadInt32Bytes(bts)
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
func (z *VerCheckReq) Msgsize() (s int) {
	s = 1 + 4
	if z.Env == nil {
		s += msgp.NilSize
	} else {
		s += z.Env.Msgsize()
	}
	s += 5 + msgp.Int32Size + 6 + msgp.Int32Size
	return
}
