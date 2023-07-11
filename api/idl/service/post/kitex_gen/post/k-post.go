// Code generated by Kitex v0.5.2. DO NOT EDIT.

package post

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"
	"resume-resolving/api/idl/service/post/kitex_gen/department"
	"resume-resolving/api/idl/service/post/kitex_gen/post_category"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
	_ = department.KitexUnusedProtection
	_ = post_category.KitexUnusedProtection
)

func (p *PostInfo) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.BYTE {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.BYTE {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField7(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 8:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField8(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 9:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField9(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PostInfo[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PostInfo) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PostId = v

	}
	return offset, nil
}

func (p *PostInfo) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PostBrief = v

	}
	return offset, nil
}

func (p *PostInfo) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PostDescription = v

	}
	return offset, nil
}

func (p *PostInfo) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PostRequire = v

	}
	return offset, nil
}

func (p *PostInfo) FastReadField5(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadByte(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.IsSchoolRecruitment = v

	}
	return offset, nil
}

func (p *PostInfo) FastReadField6(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadByte(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.IsInternship = v

	}
	return offset, nil
}

func (p *PostInfo) FastReadField7(buf []byte) (int, error) {
	offset := 0

	tmp := post_category.NewPostCategoryAllInformation()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.PostCategoryInformation = tmp
	return offset, nil
}

func (p *PostInfo) FastReadField8(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.CityInformation = make([]*department.CityInformation, 0, size)
	for i := 0; i < size; i++ {
		_elem := department.NewCityInformation()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		p.CityInformation = append(p.CityInformation, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

func (p *PostInfo) FastReadField9(buf []byte) (int, error) {
	offset := 0

	tmp := department.NewDepartmentInformation()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.DepartmentInformation = tmp
	return offset, nil
}

// for compatibility
func (p *PostInfo) FastWrite(buf []byte) int {
	return 0
}

func (p *PostInfo) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "PostInfo")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField7(buf[offset:], binaryWriter)
		offset += p.fastWriteField8(buf[offset:], binaryWriter)
		offset += p.fastWriteField9(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *PostInfo) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("PostInfo")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
		l += p.field8Length()
		l += p.field9Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *PostInfo) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "post_id", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.PostId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PostInfo) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "post_brief", thrift.STRING, 2)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.PostBrief)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PostInfo) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "post_description", thrift.STRING, 3)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.PostDescription)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PostInfo) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "post_require", thrift.STRING, 4)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.PostRequire)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PostInfo) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "is_school_recruitment", thrift.BYTE, 5)
	offset += bthrift.Binary.WriteByte(buf[offset:], p.IsSchoolRecruitment)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PostInfo) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "is_internship", thrift.BYTE, 6)
	offset += bthrift.Binary.WriteByte(buf[offset:], p.IsInternship)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PostInfo) fastWriteField7(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "post_category_information", thrift.STRUCT, 7)
	offset += p.PostCategoryInformation.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PostInfo) fastWriteField8(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "city_information", thrift.LIST, 8)
	listBeginOffset := offset
	offset += bthrift.Binary.ListBeginLength(thrift.STRUCT, 0)
	var length int
	for _, v := range p.CityInformation {
		length++
		offset += v.FastWriteNocopy(buf[offset:], binaryWriter)
	}
	bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
	offset += bthrift.Binary.WriteListEnd(buf[offset:])
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PostInfo) fastWriteField9(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "department_information", thrift.STRUCT, 9)
	offset += p.DepartmentInformation.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PostInfo) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("post_id", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.PostId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PostInfo) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("post_brief", thrift.STRING, 2)
	l += bthrift.Binary.StringLengthNocopy(p.PostBrief)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PostInfo) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("post_description", thrift.STRING, 3)
	l += bthrift.Binary.StringLengthNocopy(p.PostDescription)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PostInfo) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("post_require", thrift.STRING, 4)
	l += bthrift.Binary.StringLengthNocopy(p.PostRequire)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PostInfo) field5Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("is_school_recruitment", thrift.BYTE, 5)
	l += bthrift.Binary.ByteLength(p.IsSchoolRecruitment)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PostInfo) field6Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("is_internship", thrift.BYTE, 6)
	l += bthrift.Binary.ByteLength(p.IsInternship)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PostInfo) field7Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("post_category_information", thrift.STRUCT, 7)
	l += p.PostCategoryInformation.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PostInfo) field8Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("city_information", thrift.LIST, 8)
	l += bthrift.Binary.ListBeginLength(thrift.STRUCT, len(p.CityInformation))
	for _, v := range p.CityInformation {
		l += v.BLength()
	}
	l += bthrift.Binary.ListEndLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PostInfo) field9Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("department_information", thrift.STRUCT, 9)
	l += p.DepartmentInformation.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}
