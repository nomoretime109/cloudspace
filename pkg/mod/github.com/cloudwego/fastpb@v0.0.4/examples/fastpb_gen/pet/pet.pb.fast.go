// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package pet

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	user "github.com/cloudwego/fastpb/examples/fastpb_gen/user"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Pet) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Pet[number], err)
}

func (x *Pet) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Pet) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Age, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Pet) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Owner, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Pet) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Group = user.Group(v)
	return offset, nil
}

func (x *Pet) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *Pet) fastWriteField1(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.Name)
	return offset
}

func (x *Pet) fastWriteField2(buf []byte) (offset int) {
	if x.Age == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.Age)
	return offset
}

func (x *Pet) fastWriteField3(buf []byte) (offset int) {
	if x.Owner == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.Owner)
	return offset
}

func (x *Pet) fastWriteField4(buf []byte) (offset int) {
	if x.Group == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, int32(x.Group))
	return offset
}

func (x *Pet) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *Pet) sizeField1() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(1, x.Name)
	return n
}

func (x *Pet) sizeField2() (n int) {
	if x.Age == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.Age)
	return n
}

func (x *Pet) sizeField3() (n int) {
	if x.Owner == "" {
		return n
	}
	n += fastpb.SizeString(3, x.Owner)
	return n
}

func (x *Pet) sizeField4() (n int) {
	if x.Group == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, int32(x.Group))
	return n
}

var fieldIDToName_Pet = map[int32]string{
	1: "Name",
	2: "Age",
	3: "Owner",
	4: "Group",
}

var _ = user.File_user_user_proto