package mgm

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DefaultModel) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "create_time":
			z.CreatedAt, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "CreatedAt")
				return
			}
		case "last_modify_time":
			z.UpdatedAt, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "UpdatedAt")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z DefaultModel) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "id"
	err = en.Append(0x83, 0xa2, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.Id)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	// write "create_time"
	err = en.Append(0xab, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.CreatedAt)
	if err != nil {
		err = msgp.WrapError(err, "CreatedAt")
		return
	}
	// write "last_modify_time"
	err = en.Append(0xb0, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.UpdatedAt)
	if err != nil {
		err = msgp.WrapError(err, "UpdatedAt")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DefaultModel) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "id"
	o = append(o, 0x83, 0xa2, 0x69, 0x64)
	o = msgp.AppendString(o, z.Id)
	// string "create_time"
	o = append(o, 0xab, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt64(o, z.CreatedAt)
	// string "last_modify_time"
	o = append(o, 0xb0, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt64(o, z.UpdatedAt)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DefaultModel) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "create_time":
			z.CreatedAt, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedAt")
				return
			}
		case "last_modify_time":
			z.UpdatedAt, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "UpdatedAt")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DefaultModel) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.Id) + 12 + msgp.Int64Size + 17 + msgp.Int64Size
	return
}
