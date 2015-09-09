// Code generated by protoc-gen-go.
// source: response.proto
// DO NOT EDIT!

package zrpc

import proto "github.com/golang/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// Response is an envelope for a service responding to a request from a client
type Response struct {
	UUID             []byte          `protobuf:"bytes,1,opt" json:"UUID,omitempty"`
	Path             *string         `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Payload          []byte          `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	StatusCode       *uint32         `protobuf:"varint,4,opt,name=status_code" json:"status_code,omitempty"`
	Error            *Response_Error `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetUUID() []byte {
	if m != nil {
		return m.UUID
	}
	return nil
}

func (m *Response) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *Response) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Response) GetStatusCode() uint32 {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return 0
}

func (m *Response) GetError() *Response_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Response_Error struct {
	Message          *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response_Error) Reset()         { *m = Response_Error{} }
func (m *Response_Error) String() string { return proto.CompactTextString(m) }
func (*Response_Error) ProtoMessage()    {}

func (m *Response_Error) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}
