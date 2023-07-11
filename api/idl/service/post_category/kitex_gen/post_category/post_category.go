// Code generated by thriftgo (0.2.9). DO NOT EDIT.

package post_category

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type PostCategoryInformation struct {
	PostCategoryId   int64  `thrift:"post_category_id,1" frugal:"1,default,i64" json:"post_category_id"`
	PostCategoryName string `thrift:"post_category_name,2" frugal:"2,default,string" json:"post_category_name"`
}

func NewPostCategoryInformation() *PostCategoryInformation {
	return &PostCategoryInformation{}
}

func (p *PostCategoryInformation) InitDefault() {
	*p = PostCategoryInformation{}
}

func (p *PostCategoryInformation) GetPostCategoryId() (v int64) {
	return p.PostCategoryId
}

func (p *PostCategoryInformation) GetPostCategoryName() (v string) {
	return p.PostCategoryName
}
func (p *PostCategoryInformation) SetPostCategoryId(val int64) {
	p.PostCategoryId = val
}
func (p *PostCategoryInformation) SetPostCategoryName(val string) {
	p.PostCategoryName = val
}

var fieldIDToName_PostCategoryInformation = map[int16]string{
	1: "post_category_id",
	2: "post_category_name",
}

func (p *PostCategoryInformation) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PostCategoryInformation[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PostCategoryInformation) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.PostCategoryId = v
	}
	return nil
}

func (p *PostCategoryInformation) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.PostCategoryName = v
	}
	return nil
}

func (p *PostCategoryInformation) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("PostCategoryInformation"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *PostCategoryInformation) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("post_category_id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.PostCategoryId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *PostCategoryInformation) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("post_category_name", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.PostCategoryName); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *PostCategoryInformation) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PostCategoryInformation(%+v)", *p)
}

func (p *PostCategoryInformation) DeepEqual(ano *PostCategoryInformation) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.PostCategoryId) {
		return false
	}
	if !p.Field2DeepEqual(ano.PostCategoryName) {
		return false
	}
	return true
}

func (p *PostCategoryInformation) Field1DeepEqual(src int64) bool {

	if p.PostCategoryId != src {
		return false
	}
	return true
}
func (p *PostCategoryInformation) Field2DeepEqual(src string) bool {

	if strings.Compare(p.PostCategoryName, src) != 0 {
		return false
	}
	return true
}

type PostCategoryAllInformation struct {
	PostCategoryId         int64  `thrift:"post_category_id,1" frugal:"1,default,i64" json:"post_category_id"`
	PostCategoryName       string `thrift:"post_category_name,2" frugal:"2,default,string" json:"post_category_name"`
	PostCategoryParentId   int64  `thrift:"post_category_parent_id,3" frugal:"3,default,i64" json:"post_category_parent_id"`
	PostCategoryParentName string `thrift:"post_category_parent_name,4" frugal:"4,default,string" json:"post_category_parent_name"`
}

func NewPostCategoryAllInformation() *PostCategoryAllInformation {
	return &PostCategoryAllInformation{}
}

func (p *PostCategoryAllInformation) InitDefault() {
	*p = PostCategoryAllInformation{}
}

func (p *PostCategoryAllInformation) GetPostCategoryId() (v int64) {
	return p.PostCategoryId
}

func (p *PostCategoryAllInformation) GetPostCategoryName() (v string) {
	return p.PostCategoryName
}

func (p *PostCategoryAllInformation) GetPostCategoryParentId() (v int64) {
	return p.PostCategoryParentId
}

func (p *PostCategoryAllInformation) GetPostCategoryParentName() (v string) {
	return p.PostCategoryParentName
}
func (p *PostCategoryAllInformation) SetPostCategoryId(val int64) {
	p.PostCategoryId = val
}
func (p *PostCategoryAllInformation) SetPostCategoryName(val string) {
	p.PostCategoryName = val
}
func (p *PostCategoryAllInformation) SetPostCategoryParentId(val int64) {
	p.PostCategoryParentId = val
}
func (p *PostCategoryAllInformation) SetPostCategoryParentName(val string) {
	p.PostCategoryParentName = val
}

var fieldIDToName_PostCategoryAllInformation = map[int16]string{
	1: "post_category_id",
	2: "post_category_name",
	3: "post_category_parent_id",
	4: "post_category_parent_name",
}

func (p *PostCategoryAllInformation) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PostCategoryAllInformation[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PostCategoryAllInformation) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.PostCategoryId = v
	}
	return nil
}

func (p *PostCategoryAllInformation) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.PostCategoryName = v
	}
	return nil
}

func (p *PostCategoryAllInformation) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.PostCategoryParentId = v
	}
	return nil
}

func (p *PostCategoryAllInformation) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.PostCategoryParentName = v
	}
	return nil
}

func (p *PostCategoryAllInformation) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("PostCategoryAllInformation"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *PostCategoryAllInformation) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("post_category_id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.PostCategoryId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *PostCategoryAllInformation) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("post_category_name", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.PostCategoryName); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *PostCategoryAllInformation) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("post_category_parent_id", thrift.I64, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.PostCategoryParentId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *PostCategoryAllInformation) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("post_category_parent_name", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.PostCategoryParentName); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *PostCategoryAllInformation) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PostCategoryAllInformation(%+v)", *p)
}

func (p *PostCategoryAllInformation) DeepEqual(ano *PostCategoryAllInformation) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.PostCategoryId) {
		return false
	}
	if !p.Field2DeepEqual(ano.PostCategoryName) {
		return false
	}
	if !p.Field3DeepEqual(ano.PostCategoryParentId) {
		return false
	}
	if !p.Field4DeepEqual(ano.PostCategoryParentName) {
		return false
	}
	return true
}

func (p *PostCategoryAllInformation) Field1DeepEqual(src int64) bool {

	if p.PostCategoryId != src {
		return false
	}
	return true
}
func (p *PostCategoryAllInformation) Field2DeepEqual(src string) bool {

	if strings.Compare(p.PostCategoryName, src) != 0 {
		return false
	}
	return true
}
func (p *PostCategoryAllInformation) Field3DeepEqual(src int64) bool {

	if p.PostCategoryParentId != src {
		return false
	}
	return true
}
func (p *PostCategoryAllInformation) Field4DeepEqual(src string) bool {

	if strings.Compare(p.PostCategoryParentName, src) != 0 {
		return false
	}
	return true
}
