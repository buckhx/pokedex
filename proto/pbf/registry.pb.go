// Code generated by protoc-gen-gogo.
// source: registry.proto
// DO NOT EDIT!

package pbf

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

import strconv "strconv"

import bytes "bytes"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import reflect "reflect"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// IRL Scopes would live in different go packages and not need to be prefixed
type RegistryScope int32

const (
	CIVILIAN   RegistryScope = 0
	TRAINER    RegistryScope = 1
	GYM_LEADER RegistryScope = 2
	ELITE_FOUR RegistryScope = 3
	PROFESSOR  RegistryScope = 4
)

var RegistryScope_name = map[int32]string{
	0: "CIVILIAN",
	1: "TRAINER",
	2: "GYM_LEADER",
	3: "ELITE_FOUR",
	4: "PROFESSOR",
}
var RegistryScope_value = map[string]int32{
	"CIVILIAN":   0,
	"TRAINER":    1,
	"GYM_LEADER": 2,
	"ELITE_FOUR": 3,
	"PROFESSOR":  4,
}

func (RegistryScope) EnumDescriptor() ([]byte, []int) { return fileDescriptorRegistry, []int{0} }

type Token_Type int32

const (
	ACCESS Token_Type = 0
)

var Token_Type_name = map[int32]string{
	0: "ACCESS",
}
var Token_Type_value = map[string]int32{
	"ACCESS": 0,
}

func (Token_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorRegistry, []int{0, 0} }

type Token struct {
	Token []byte     `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Type  Token_Type `protobuf:"varint,2,opt,name=type,proto3,enum=buckhx.safari.registry.Token_Type" json:"type,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptorRegistry, []int{0} }

type Token_Scopes struct {
	Registry RegistryScope `protobuf:"varint,1,opt,name=registry,proto3,enum=buckhx.safari.registry.RegistryScope" json:"registry,omitempty"`
	Pokedex  PokedexScope  `protobuf:"varint,2,opt,name=pokedex,proto3,enum=buckhx.safari.pokedex.PokedexScope" json:"pokedex,omitempty"`
	Zone     ZoneScope     `protobuf:"varint,3,opt,name=zone,proto3,enum=buckhx.safari.zone.ZoneScope" json:"zone,omitempty"`
}

func (m *Token_Scopes) Reset()                    { *m = Token_Scopes{} }
func (*Token_Scopes) ProtoMessage()               {}
func (*Token_Scopes) Descriptor() ([]byte, []int) { return fileDescriptorRegistry, []int{0, 0} }

func init() {
	proto.RegisterType((*Token)(nil), "buckhx.safari.registry.Token")
	proto.RegisterType((*Token_Scopes)(nil), "buckhx.safari.registry.Token.Scopes")
	proto.RegisterEnum("buckhx.safari.registry.RegistryScope", RegistryScope_name, RegistryScope_value)
	proto.RegisterEnum("buckhx.safari.registry.Token_Type", Token_Type_name, Token_Type_value)
}
func (x RegistryScope) String() string {
	s, ok := RegistryScope_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x Token_Type) String() string {
	s, ok := Token_Type_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Token) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Token)
	if !ok {
		that2, ok := that.(Token)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Token, that1.Token) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	return true
}
func (this *Token_Scopes) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Token_Scopes)
	if !ok {
		that2, ok := that.(Token_Scopes)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Registry != that1.Registry {
		return false
	}
	if this.Pokedex != that1.Pokedex {
		return false
	}
	if this.Zone != that1.Zone {
		return false
	}
	return true
}
func (this *Token) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pbf.Token{")
	s = append(s, "Token: "+fmt.Sprintf("%#v", this.Token)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Token_Scopes) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&pbf.Token_Scopes{")
	s = append(s, "Registry: "+fmt.Sprintf("%#v", this.Registry)+",\n")
	s = append(s, "Pokedex: "+fmt.Sprintf("%#v", this.Pokedex)+",\n")
	s = append(s, "Zone: "+fmt.Sprintf("%#v", this.Zone)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRegistry(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringRegistry(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Registry service

type RegistryClient interface {
	// Register makes a creates a new trainer in the safari
	//
	// Trainer name, password, age & gender are required.
	// Any other supplied fields will be ignored
	Register(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Response, error)
	// Get fetchs a trainer
	//
	// The populated fields will depend on the auth scope of the token
	Get(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Trainer, error)
	// Enter authenticates a user to retrieve a an access token to authorize requests for a safari
	//
	// HTTPS required w/ HTTP basic access authentication via a header
	// Authorization: Basic BASE64({user:pass})
	Enter(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Token, error)
}

type registryClient struct {
	cc *grpc.ClientConn
}

func NewRegistryClient(cc *grpc.ClientConn) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) Register(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/buckhx.safari.registry.Registry/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Get(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Trainer, error) {
	out := new(Trainer)
	err := grpc.Invoke(ctx, "/buckhx.safari.registry.Registry/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Enter(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := grpc.Invoke(ctx, "/buckhx.safari.registry.Registry/Enter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Registry service

type RegistryServer interface {
	// Register makes a creates a new trainer in the safari
	//
	// Trainer name, password, age & gender are required.
	// Any other supplied fields will be ignored
	Register(context.Context, *Trainer) (*Response, error)
	// Get fetchs a trainer
	//
	// The populated fields will depend on the auth scope of the token
	Get(context.Context, *Trainer) (*Trainer, error)
	// Enter authenticates a user to retrieve a an access token to authorize requests for a safari
	//
	// HTTPS required w/ HTTP basic access authentication via a header
	// Authorization: Basic BASE64({user:pass})
	Enter(context.Context, *Trainer) (*Token, error)
}

func RegisterRegistryServer(s *grpc.Server, srv RegistryServer) {
	s.RegisterService(&_Registry_serviceDesc, srv)
}

func _Registry_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trainer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buckhx.safari.registry.Registry/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Register(ctx, req.(*Trainer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trainer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buckhx.safari.registry.Registry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Get(ctx, req.(*Trainer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Enter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trainer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Enter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buckhx.safari.registry.Registry/Enter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Enter(ctx, req.(*Trainer))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buckhx.safari.registry.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Registry_Register_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Registry_Get_Handler,
		},
		{
			MethodName: "Enter",
			Handler:    _Registry_Enter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorRegistry,
}

func (m *Token) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Token) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintRegistry(data, i, uint64(len(m.Token)))
		i += copy(data[i:], m.Token)
	}
	if m.Type != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintRegistry(data, i, uint64(m.Type))
	}
	return i, nil
}

func (m *Token_Scopes) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Token_Scopes) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Registry != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintRegistry(data, i, uint64(m.Registry))
	}
	if m.Pokedex != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintRegistry(data, i, uint64(m.Pokedex))
	}
	if m.Zone != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintRegistry(data, i, uint64(m.Zone))
	}
	return i, nil
}

func encodeFixed64Registry(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Registry(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRegistry(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Token) Size() (n int) {
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovRegistry(uint64(m.Type))
	}
	return n
}

func (m *Token_Scopes) Size() (n int) {
	var l int
	_ = l
	if m.Registry != 0 {
		n += 1 + sovRegistry(uint64(m.Registry))
	}
	if m.Pokedex != 0 {
		n += 1 + sovRegistry(uint64(m.Pokedex))
	}
	if m.Zone != 0 {
		n += 1 + sovRegistry(uint64(m.Zone))
	}
	return n
}

func sovRegistry(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRegistry(x uint64) (n int) {
	return sovRegistry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Token) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Token{`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Token_Scopes) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Token_Scopes{`,
		`Registry:` + fmt.Sprintf("%v", this.Registry) + `,`,
		`Pokedex:` + fmt.Sprintf("%v", this.Pokedex) + `,`,
		`Zone:` + fmt.Sprintf("%v", this.Zone) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRegistry(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Token) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegistry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = append(m.Token[:0], data[iNdEx:postIndex]...)
			if m.Token == nil {
				m.Token = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (Token_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRegistry(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegistry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Token_Scopes) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegistry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Scopes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Scopes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			m.Registry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Registry |= (RegistryScope(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pokedex", wireType)
			}
			m.Pokedex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Pokedex |= (PokedexScope(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zone", wireType)
			}
			m.Zone = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Zone |= (ZoneScope(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRegistry(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegistry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRegistry(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRegistry
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRegistry
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRegistry
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRegistry(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRegistry = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRegistry   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorRegistry = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6b, 0xd4, 0x50,
	0x10, 0xc7, 0x37, 0xd9, 0x1f, 0xdd, 0x8e, 0xbb, 0x21, 0x4c, 0xb5, 0x2c, 0xa9, 0x5d, 0x24, 0x22,
	0x48, 0x0f, 0x09, 0x5d, 0xc1, 0x83, 0xe0, 0x21, 0xae, 0x69, 0x09, 0xac, 0xdd, 0xf2, 0x36, 0x0a,
	0xad, 0x87, 0x92, 0x6d, 0x5f, 0xb7, 0xa1, 0x9a, 0x17, 0x92, 0xb7, 0xd0, 0x2a, 0x82, 0xf8, 0x17,
	0x08, 0xfe, 0x13, 0xf5, 0x8f, 0xf0, 0xee, 0xb1, 0xe0, 0xc5, 0xa3, 0xad, 0x1e, 0x3c, 0xfa, 0x27,
	0xf8, 0xf2, 0x92, 0x08, 0x5d, 0xb4, 0xf6, 0x30, 0x79, 0x33, 0x99, 0xf9, 0x7e, 0x98, 0x79, 0x6f,
	0x40, 0x4b, 0xe8, 0x24, 0x4c, 0x79, 0x72, 0x6c, 0xc5, 0x09, 0xe3, 0x0c, 0x17, 0xc7, 0xd3, 0xdd,
	0xc3, 0x83, 0x23, 0x2b, 0x0d, 0xf6, 0x83, 0x24, 0xb4, 0xca, 0xac, 0x71, 0x73, 0xc2, 0xd8, 0xe4,
	0x05, 0xb5, 0x83, 0x38, 0xb4, 0x83, 0x28, 0x62, 0x3c, 0xe0, 0x21, 0x8b, 0xd2, 0x5c, 0x65, 0xb4,
	0x63, 0x76, 0x48, 0xf7, 0xe8, 0x51, 0x11, 0xb6, 0x68, 0xc4, 0x43, 0x5e, 0x20, 0x0d, 0x78, 0xc5,
	0x22, 0x5a, 0xf8, 0xf3, 0x2f, 0xd3, 0x49, 0xee, 0x9a, 0x1f, 0x55, 0xa8, 0xfb, 0x42, 0x16, 0xe1,
	0x75, 0xa8, 0xf3, 0xcc, 0xe9, 0x28, 0xb7, 0x94, 0xbb, 0x2d, 0x92, 0x07, 0x78, 0x1f, 0x6a, 0xfc,
	0x38, 0xa6, 0x1d, 0x55, 0xfc, 0xd4, 0x7a, 0xa6, 0xf5, 0xf7, 0xc6, 0x2c, 0x89, 0xb0, 0x7c, 0x51,
	0x49, 0x64, 0xbd, 0xf1, 0x49, 0x81, 0xc6, 0x68, 0x97, 0xc5, 0x34, 0x45, 0x07, 0x9a, 0x65, 0x9d,
	0x64, 0x6b, 0xbd, 0x3b, 0xff, 0xc2, 0x90, 0xc2, 0x91, 0x4a, 0xf2, 0x47, 0x86, 0x0f, 0x61, 0xae,
	0x98, 0xad, 0x68, 0xe4, 0xf6, 0x0c, 0xa1, 0x9c, 0x7c, 0x33, 0x3f, 0x73, 0x7d, 0xa9, 0xc1, 0x55,
	0xa8, 0x65, 0xd3, 0x77, 0xaa, 0x52, 0xbb, 0x3c, 0xa3, 0x95, 0x17, 0xb3, 0x2d, 0x3e, 0xb9, 0x4a,
	0x96, 0x9a, 0x08, 0xb5, 0x6c, 0x1a, 0x04, 0x68, 0x38, 0xfd, 0xbe, 0x3b, 0x1a, 0xe9, 0x95, 0x95,
	0x2d, 0x68, 0x5f, 0x68, 0x10, 0x5b, 0xd0, 0xec, 0x7b, 0xcf, 0xbc, 0x81, 0xe7, 0x6c, 0xe8, 0x15,
	0xbc, 0x06, 0x73, 0x3e, 0x71, 0xbc, 0x0d, 0x97, 0xe8, 0x0a, 0x6a, 0x00, 0xeb, 0x5b, 0x4f, 0x76,
	0x06, 0xae, 0xf3, 0x58, 0xc4, 0x6a, 0x16, 0xbb, 0x03, 0xcf, 0x77, 0x77, 0xd6, 0x86, 0x4f, 0x89,
	0x5e, 0xc5, 0x36, 0xcc, 0x6f, 0x92, 0xe1, 0x9a, 0x00, 0x0f, 0x89, 0x5e, 0xeb, 0x9d, 0xa8, 0xd0,
	0x2c, 0xd9, 0xf8, 0xbc, 0xf4, 0x69, 0x82, 0xb3, 0xcd, 0x16, 0x6f, 0xea, 0x27, 0x41, 0x18, 0xd1,
	0xc4, 0x58, 0x9a, 0x49, 0x67, 0x0f, 0x4b, 0x68, 0x1a, 0x8b, 0xa5, 0xa0, 0xe6, 0xc2, 0xbb, 0x2f,
	0x3f, 0x3e, 0xa8, 0x6d, 0xb3, 0x69, 0xf3, 0xbc, 0xfc, 0x81, 0xb2, 0x22, 0xe0, 0xd5, 0x75, 0xca,
	0xff, 0xc7, 0xbd, 0x3c, 0x6d, 0x2e, 0x4a, 0xb2, 0x8e, 0x5a, 0x49, 0xb6, 0x5f, 0x4f, 0xc3, 0xbd,
	0x37, 0x38, 0x86, 0xba, 0x1b, 0x5d, 0xa1, 0xed, 0xe5, 0x4b, 0xf7, 0xc8, 0x5c, 0x92, 0xf8, 0x1b,
	0xb8, 0x70, 0x11, 0x6f, 0x07, 0x53, 0x7e, 0xf0, 0x68, 0xf5, 0xf4, 0xac, 0x5b, 0xf9, 0x2a, 0xec,
	0xd7, 0x59, 0x57, 0x79, 0x7b, 0xde, 0x55, 0x4e, 0x84, 0x7d, 0x16, 0x76, 0x2a, 0xec, 0x9b, 0xb0,
	0x9f, 0xe7, 0x22, 0x27, 0xce, 0xf7, 0xdf, 0xbb, 0x95, 0xed, 0x6a, 0x3c, 0xde, 0x1f, 0x37, 0xe4,
	0xae, 0xdf, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xce, 0x4b, 0xc0, 0x77, 0x67, 0x03, 0x00, 0x00,
}
