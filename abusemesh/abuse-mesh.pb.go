// Code generated by protoc-gen-go. DO NOT EDIT.
// source: abuse-mesh.proto

package abusemesh

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The type of update to a table
type TableUpdateType int32

const (
	// A new entity is added to the table
	TableUpdateType_TABLE_UPDATE_NEW TableUpdateType = 0
	// A existing entity was changed in the table
	TableUpdateType_TABLE_UPDATE_EDIT TableUpdateType = 1
	// A entity was deleted from the table
	TableUpdateType_TABLE_UPDATE_DELETE TableUpdateType = 2
)

var TableUpdateType_name = map[int32]string{
	0: "TABLE_UPDATE_NEW",
	1: "TABLE_UPDATE_EDIT",
	2: "TABLE_UPDATE_DELETE",
}
var TableUpdateType_value = map[string]int32{
	"TABLE_UPDATE_NEW":    0,
	"TABLE_UPDATE_EDIT":   1,
	"TABLE_UPDATE_DELETE": 2,
}

func (x TableUpdateType) String() string {
	return proto.EnumName(TableUpdateType_name, int32(x))
}
func (TableUpdateType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type GetNodeRequest struct {
}

func (m *GetNodeRequest) Reset()                    { *m = GetNodeRequest{} }
func (m *GetNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeRequest) ProtoMessage()               {}
func (*GetNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type GetNodeTableRequest struct {
}

func (m *GetNodeTableRequest) Reset()                    { *m = GetNodeTableRequest{} }
func (m *GetNodeTableRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeTableRequest) ProtoMessage()               {}
func (*GetNodeTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type GetNeighborTableRequest struct {
	// If true the server should return the global neighbor table
	// If false the server should return only its own neighbor table
	Global bool `protobuf:"varint,1,opt,name=global" json:"global,omitempty"`
}

func (m *GetNeighborTableRequest) Reset()                    { *m = GetNeighborTableRequest{} }
func (m *GetNeighborTableRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNeighborTableRequest) ProtoMessage()               {}
func (*GetNeighborTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetNeighborTableRequest) GetGlobal() bool {
	if m != nil {
		return m.Global
	}
	return false
}

type GetReportTableRequest struct {
}

func (m *GetReportTableRequest) Reset()                    { *m = GetReportTableRequest{} }
func (m *GetReportTableRequest) String() string            { return proto.CompactTextString(m) }
func (*GetReportTableRequest) ProtoMessage()               {}
func (*GetReportTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type GetDelistRequestTableRequest struct {
}

func (m *GetDelistRequestTableRequest) Reset()                    { *m = GetDelistRequestTableRequest{} }
func (m *GetDelistRequestTableRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDelistRequestTableRequest) ProtoMessage()               {}
func (*GetDelistRequestTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type GetDelistAcceptanceTableRequest struct {
}

func (m *GetDelistAcceptanceTableRequest) Reset()                    { *m = GetDelistAcceptanceTableRequest{} }
func (m *GetDelistAcceptanceTableRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDelistAcceptanceTableRequest) ProtoMessage()               {}
func (*GetDelistAcceptanceTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type TableUpdateStreamRequest struct {
}

func (m *TableUpdateStreamRequest) Reset()                    { *m = TableUpdateStreamRequest{} }
func (m *TableUpdateStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*TableUpdateStreamRequest) ProtoMessage()               {}
func (*TableUpdateStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type GetNodeTableResponse struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *GetNodeTableResponse) Reset()                    { *m = GetNodeTableResponse{} }
func (m *GetNodeTableResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNodeTableResponse) ProtoMessage()               {}
func (*GetNodeTableResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *GetNodeTableResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type GetNeighborTableResponse struct {
	Neighbors []*Neighbor `protobuf:"bytes,1,rep,name=neighbors" json:"neighbors,omitempty"`
}

func (m *GetNeighborTableResponse) Reset()                    { *m = GetNeighborTableResponse{} }
func (m *GetNeighborTableResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNeighborTableResponse) ProtoMessage()               {}
func (*GetNeighborTableResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *GetNeighborTableResponse) GetNeighbors() []*Neighbor {
	if m != nil {
		return m.Neighbors
	}
	return nil
}

type GetReportTableResponse struct {
	Reports []*Report `protobuf:"bytes,1,rep,name=reports" json:"reports,omitempty"`
}

func (m *GetReportTableResponse) Reset()                    { *m = GetReportTableResponse{} }
func (m *GetReportTableResponse) String() string            { return proto.CompactTextString(m) }
func (*GetReportTableResponse) ProtoMessage()               {}
func (*GetReportTableResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *GetReportTableResponse) GetReports() []*Report {
	if m != nil {
		return m.Reports
	}
	return nil
}

type GetDelistRequestTableResponse struct {
	DelistRequests []*DelistRequest `protobuf:"bytes,1,rep,name=delist_requests,json=delistRequests" json:"delist_requests,omitempty"`
}

func (m *GetDelistRequestTableResponse) Reset()                    { *m = GetDelistRequestTableResponse{} }
func (m *GetDelistRequestTableResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDelistRequestTableResponse) ProtoMessage()               {}
func (*GetDelistRequestTableResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *GetDelistRequestTableResponse) GetDelistRequests() []*DelistRequest {
	if m != nil {
		return m.DelistRequests
	}
	return nil
}

type GetDelistAcceptanceTableResponse struct {
	DelistAcceptance []*DelistAcceptance `protobuf:"bytes,1,rep,name=delist_acceptance,json=delistAcceptance" json:"delist_acceptance,omitempty"`
}

func (m *GetDelistAcceptanceTableResponse) Reset()         { *m = GetDelistAcceptanceTableResponse{} }
func (m *GetDelistAcceptanceTableResponse) String() string { return proto.CompactTextString(m) }
func (*GetDelistAcceptanceTableResponse) ProtoMessage()    {}
func (*GetDelistAcceptanceTableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{11}
}

func (m *GetDelistAcceptanceTableResponse) GetDelistAcceptance() []*DelistAcceptance {
	if m != nil {
		return m.DelistAcceptance
	}
	return nil
}

type TableUpdate struct {
	UpdateType TableUpdateType `protobuf:"varint,1,opt,name=update_type,json=updateType,enum=abusemesh.TableUpdateType" json:"update_type,omitempty"`
	// Types that are valid to be assigned to TableEntity:
	//	*TableUpdate_Node
	//	*TableUpdate_Neighbor
	//	*TableUpdate_Report
	//	*TableUpdate_DelistRequests
	//	*TableUpdate_DelistAcceptance
	TableEntity isTableUpdate_TableEntity `protobuf_oneof:"table_entity"`
}

func (m *TableUpdate) Reset()                    { *m = TableUpdate{} }
func (m *TableUpdate) String() string            { return proto.CompactTextString(m) }
func (*TableUpdate) ProtoMessage()               {}
func (*TableUpdate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

type isTableUpdate_TableEntity interface {
	isTableUpdate_TableEntity()
}

type TableUpdate_Node struct {
	Node *Node `protobuf:"bytes,2,opt,name=node,oneof"`
}
type TableUpdate_Neighbor struct {
	Neighbor *Neighbor `protobuf:"bytes,3,opt,name=neighbor,oneof"`
}
type TableUpdate_Report struct {
	Report *Report `protobuf:"bytes,4,opt,name=report,oneof"`
}
type TableUpdate_DelistRequests struct {
	DelistRequests *DelistRequest `protobuf:"bytes,5,opt,name=delist_requests,json=delistRequests,oneof"`
}
type TableUpdate_DelistAcceptance struct {
	DelistAcceptance *DelistAcceptance `protobuf:"bytes,6,opt,name=delist_acceptance,json=delistAcceptance,oneof"`
}

func (*TableUpdate_Node) isTableUpdate_TableEntity()             {}
func (*TableUpdate_Neighbor) isTableUpdate_TableEntity()         {}
func (*TableUpdate_Report) isTableUpdate_TableEntity()           {}
func (*TableUpdate_DelistRequests) isTableUpdate_TableEntity()   {}
func (*TableUpdate_DelistAcceptance) isTableUpdate_TableEntity() {}

func (m *TableUpdate) GetTableEntity() isTableUpdate_TableEntity {
	if m != nil {
		return m.TableEntity
	}
	return nil
}

func (m *TableUpdate) GetUpdateType() TableUpdateType {
	if m != nil {
		return m.UpdateType
	}
	return TableUpdateType_TABLE_UPDATE_NEW
}

func (m *TableUpdate) GetNode() *Node {
	if x, ok := m.GetTableEntity().(*TableUpdate_Node); ok {
		return x.Node
	}
	return nil
}

func (m *TableUpdate) GetNeighbor() *Neighbor {
	if x, ok := m.GetTableEntity().(*TableUpdate_Neighbor); ok {
		return x.Neighbor
	}
	return nil
}

func (m *TableUpdate) GetReport() *Report {
	if x, ok := m.GetTableEntity().(*TableUpdate_Report); ok {
		return x.Report
	}
	return nil
}

func (m *TableUpdate) GetDelistRequests() *DelistRequest {
	if x, ok := m.GetTableEntity().(*TableUpdate_DelistRequests); ok {
		return x.DelistRequests
	}
	return nil
}

func (m *TableUpdate) GetDelistAcceptance() *DelistAcceptance {
	if x, ok := m.GetTableEntity().(*TableUpdate_DelistAcceptance); ok {
		return x.DelistAcceptance
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TableUpdate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TableUpdate_OneofMarshaler, _TableUpdate_OneofUnmarshaler, _TableUpdate_OneofSizer, []interface{}{
		(*TableUpdate_Node)(nil),
		(*TableUpdate_Neighbor)(nil),
		(*TableUpdate_Report)(nil),
		(*TableUpdate_DelistRequests)(nil),
		(*TableUpdate_DelistAcceptance)(nil),
	}
}

func _TableUpdate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TableUpdate)
	// table_entity
	switch x := m.TableEntity.(type) {
	case *TableUpdate_Node:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Node); err != nil {
			return err
		}
	case *TableUpdate_Neighbor:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Neighbor); err != nil {
			return err
		}
	case *TableUpdate_Report:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Report); err != nil {
			return err
		}
	case *TableUpdate_DelistRequests:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DelistRequests); err != nil {
			return err
		}
	case *TableUpdate_DelistAcceptance:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DelistAcceptance); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TableUpdate.TableEntity has unexpected type %T", x)
	}
	return nil
}

func _TableUpdate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TableUpdate)
	switch tag {
	case 2: // table_entity.node
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Node)
		err := b.DecodeMessage(msg)
		m.TableEntity = &TableUpdate_Node{msg}
		return true, err
	case 3: // table_entity.neighbor
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Neighbor)
		err := b.DecodeMessage(msg)
		m.TableEntity = &TableUpdate_Neighbor{msg}
		return true, err
	case 4: // table_entity.report
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Report)
		err := b.DecodeMessage(msg)
		m.TableEntity = &TableUpdate_Report{msg}
		return true, err
	case 5: // table_entity.delist_requests
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DelistRequest)
		err := b.DecodeMessage(msg)
		m.TableEntity = &TableUpdate_DelistRequests{msg}
		return true, err
	case 6: // table_entity.delist_acceptance
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DelistAcceptance)
		err := b.DecodeMessage(msg)
		m.TableEntity = &TableUpdate_DelistAcceptance{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TableUpdate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TableUpdate)
	// table_entity
	switch x := m.TableEntity.(type) {
	case *TableUpdate_Node:
		s := proto.Size(x.Node)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TableUpdate_Neighbor:
		s := proto.Size(x.Neighbor)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TableUpdate_Report:
		s := proto.Size(x.Report)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TableUpdate_DelistRequests:
		s := proto.Size(x.DelistRequests)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TableUpdate_DelistAcceptance:
		s := proto.Size(x.DelistAcceptance)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*GetNodeRequest)(nil), "abusemesh.GetNodeRequest")
	proto.RegisterType((*GetNodeTableRequest)(nil), "abusemesh.GetNodeTableRequest")
	proto.RegisterType((*GetNeighborTableRequest)(nil), "abusemesh.GetNeighborTableRequest")
	proto.RegisterType((*GetReportTableRequest)(nil), "abusemesh.GetReportTableRequest")
	proto.RegisterType((*GetDelistRequestTableRequest)(nil), "abusemesh.GetDelistRequestTableRequest")
	proto.RegisterType((*GetDelistAcceptanceTableRequest)(nil), "abusemesh.GetDelistAcceptanceTableRequest")
	proto.RegisterType((*TableUpdateStreamRequest)(nil), "abusemesh.TableUpdateStreamRequest")
	proto.RegisterType((*GetNodeTableResponse)(nil), "abusemesh.GetNodeTableResponse")
	proto.RegisterType((*GetNeighborTableResponse)(nil), "abusemesh.GetNeighborTableResponse")
	proto.RegisterType((*GetReportTableResponse)(nil), "abusemesh.GetReportTableResponse")
	proto.RegisterType((*GetDelistRequestTableResponse)(nil), "abusemesh.GetDelistRequestTableResponse")
	proto.RegisterType((*GetDelistAcceptanceTableResponse)(nil), "abusemesh.GetDelistAcceptanceTableResponse")
	proto.RegisterType((*TableUpdate)(nil), "abusemesh.TableUpdate")
	proto.RegisterEnum("abusemesh.TableUpdateType", TableUpdateType_name, TableUpdateType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AbuseMesh service

type AbuseMeshClient interface {
	// Returns the Node data of the current node
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*Node, error)
	// Returns the contents of the node table
	GetNodeTable(ctx context.Context, in *GetNodeTableRequest, opts ...grpc.CallOption) (*GetNodeTableResponse, error)
	// Returns the contents of the neighbor table of the node
	GetNeighborTable(ctx context.Context, in *GetNeighborTableRequest, opts ...grpc.CallOption) (*GetNeighborTableResponse, error)
	// Returns the contents of the report table of the node
	GetReportTable(ctx context.Context, in *GetReportTableRequest, opts ...grpc.CallOption) (*GetReportTableResponse, error)
	// Returns the contents of the report delist requests table of the node
	GetDelistRequestTable(ctx context.Context, in *GetDelistRequestTableRequest, opts ...grpc.CallOption) (*GetDelistRequestTableResponse, error)
	// Returns the contents of the report delist acceptance table of the node
	GetDelistAcceptanceTable(ctx context.Context, in *GetDelistAcceptanceTableRequest, opts ...grpc.CallOption) (*GetDelistAcceptanceTableResponse, error)
	// Opens a stream on which all updates to the tables of a node are published
	TableUpdateStream(ctx context.Context, in *TableUpdateStreamRequest, opts ...grpc.CallOption) (AbuseMesh_TableUpdateStreamClient, error)
}

type abuseMeshClient struct {
	cc *grpc.ClientConn
}

func NewAbuseMeshClient(cc *grpc.ClientConn) AbuseMeshClient {
	return &abuseMeshClient{cc}
}

func (c *abuseMeshClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := grpc.Invoke(ctx, "/abusemesh.AbuseMesh/GetNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abuseMeshClient) GetNodeTable(ctx context.Context, in *GetNodeTableRequest, opts ...grpc.CallOption) (*GetNodeTableResponse, error) {
	out := new(GetNodeTableResponse)
	err := grpc.Invoke(ctx, "/abusemesh.AbuseMesh/GetNodeTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abuseMeshClient) GetNeighborTable(ctx context.Context, in *GetNeighborTableRequest, opts ...grpc.CallOption) (*GetNeighborTableResponse, error) {
	out := new(GetNeighborTableResponse)
	err := grpc.Invoke(ctx, "/abusemesh.AbuseMesh/GetNeighborTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abuseMeshClient) GetReportTable(ctx context.Context, in *GetReportTableRequest, opts ...grpc.CallOption) (*GetReportTableResponse, error) {
	out := new(GetReportTableResponse)
	err := grpc.Invoke(ctx, "/abusemesh.AbuseMesh/GetReportTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abuseMeshClient) GetDelistRequestTable(ctx context.Context, in *GetDelistRequestTableRequest, opts ...grpc.CallOption) (*GetDelistRequestTableResponse, error) {
	out := new(GetDelistRequestTableResponse)
	err := grpc.Invoke(ctx, "/abusemesh.AbuseMesh/GetDelistRequestTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abuseMeshClient) GetDelistAcceptanceTable(ctx context.Context, in *GetDelistAcceptanceTableRequest, opts ...grpc.CallOption) (*GetDelistAcceptanceTableResponse, error) {
	out := new(GetDelistAcceptanceTableResponse)
	err := grpc.Invoke(ctx, "/abusemesh.AbuseMesh/GetDelistAcceptanceTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abuseMeshClient) TableUpdateStream(ctx context.Context, in *TableUpdateStreamRequest, opts ...grpc.CallOption) (AbuseMesh_TableUpdateStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AbuseMesh_serviceDesc.Streams[0], c.cc, "/abusemesh.AbuseMesh/TableUpdateStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &abuseMeshTableUpdateStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AbuseMesh_TableUpdateStreamClient interface {
	Recv() (*TableUpdate, error)
	grpc.ClientStream
}

type abuseMeshTableUpdateStreamClient struct {
	grpc.ClientStream
}

func (x *abuseMeshTableUpdateStreamClient) Recv() (*TableUpdate, error) {
	m := new(TableUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AbuseMesh service

type AbuseMeshServer interface {
	// Returns the Node data of the current node
	GetNode(context.Context, *GetNodeRequest) (*Node, error)
	// Returns the contents of the node table
	GetNodeTable(context.Context, *GetNodeTableRequest) (*GetNodeTableResponse, error)
	// Returns the contents of the neighbor table of the node
	GetNeighborTable(context.Context, *GetNeighborTableRequest) (*GetNeighborTableResponse, error)
	// Returns the contents of the report table of the node
	GetReportTable(context.Context, *GetReportTableRequest) (*GetReportTableResponse, error)
	// Returns the contents of the report delist requests table of the node
	GetDelistRequestTable(context.Context, *GetDelistRequestTableRequest) (*GetDelistRequestTableResponse, error)
	// Returns the contents of the report delist acceptance table of the node
	GetDelistAcceptanceTable(context.Context, *GetDelistAcceptanceTableRequest) (*GetDelistAcceptanceTableResponse, error)
	// Opens a stream on which all updates to the tables of a node are published
	TableUpdateStream(*TableUpdateStreamRequest, AbuseMesh_TableUpdateStreamServer) error
}

func RegisterAbuseMeshServer(s *grpc.Server, srv AbuseMeshServer) {
	s.RegisterService(&_AbuseMesh_serviceDesc, srv)
}

func _AbuseMesh_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbuseMeshServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abusemesh.AbuseMesh/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbuseMeshServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AbuseMesh_GetNodeTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbuseMeshServer).GetNodeTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abusemesh.AbuseMesh/GetNodeTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbuseMeshServer).GetNodeTable(ctx, req.(*GetNodeTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AbuseMesh_GetNeighborTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNeighborTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbuseMeshServer).GetNeighborTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abusemesh.AbuseMesh/GetNeighborTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbuseMeshServer).GetNeighborTable(ctx, req.(*GetNeighborTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AbuseMesh_GetReportTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbuseMeshServer).GetReportTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abusemesh.AbuseMesh/GetReportTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbuseMeshServer).GetReportTable(ctx, req.(*GetReportTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AbuseMesh_GetDelistRequestTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDelistRequestTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbuseMeshServer).GetDelistRequestTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abusemesh.AbuseMesh/GetDelistRequestTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbuseMeshServer).GetDelistRequestTable(ctx, req.(*GetDelistRequestTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AbuseMesh_GetDelistAcceptanceTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDelistAcceptanceTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbuseMeshServer).GetDelistAcceptanceTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abusemesh.AbuseMesh/GetDelistAcceptanceTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbuseMeshServer).GetDelistAcceptanceTable(ctx, req.(*GetDelistAcceptanceTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AbuseMesh_TableUpdateStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TableUpdateStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AbuseMeshServer).TableUpdateStream(m, &abuseMeshTableUpdateStreamServer{stream})
}

type AbuseMesh_TableUpdateStreamServer interface {
	Send(*TableUpdate) error
	grpc.ServerStream
}

type abuseMeshTableUpdateStreamServer struct {
	grpc.ServerStream
}

func (x *abuseMeshTableUpdateStreamServer) Send(m *TableUpdate) error {
	return x.ServerStream.SendMsg(m)
}

var _AbuseMesh_serviceDesc = grpc.ServiceDesc{
	ServiceName: "abusemesh.AbuseMesh",
	HandlerType: (*AbuseMeshServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNode",
			Handler:    _AbuseMesh_GetNode_Handler,
		},
		{
			MethodName: "GetNodeTable",
			Handler:    _AbuseMesh_GetNodeTable_Handler,
		},
		{
			MethodName: "GetNeighborTable",
			Handler:    _AbuseMesh_GetNeighborTable_Handler,
		},
		{
			MethodName: "GetReportTable",
			Handler:    _AbuseMesh_GetReportTable_Handler,
		},
		{
			MethodName: "GetDelistRequestTable",
			Handler:    _AbuseMesh_GetDelistRequestTable_Handler,
		},
		{
			MethodName: "GetDelistAcceptanceTable",
			Handler:    _AbuseMesh_GetDelistAcceptanceTable_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TableUpdateStream",
			Handler:       _AbuseMesh_TableUpdateStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "abuse-mesh.proto",
}

func init() { proto.RegisterFile("abuse-mesh.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xb5, 0xf3, 0x41, 0x92, 0x49, 0x44, 0xcc, 0xe6, 0x03, 0xd7, 0x6d, 0x13, 0xe2, 0x28, 0x6a,
	0x14, 0x14, 0x54, 0xa8, 0x7a, 0xaa, 0x7a, 0x20, 0xc5, 0xc2, 0xad, 0x92, 0xb4, 0xda, 0x1a, 0xe5,
	0xd0, 0x03, 0xb2, 0xf1, 0x2a, 0x50, 0x81, 0xed, 0xda, 0xcb, 0x81, 0xbf, 0xdd, 0x73, 0x0f, 0x15,
	0xcb, 0xda, 0xac, 0x3f, 0xa0, 0xbd, 0x79, 0xe7, 0xbd, 0x79, 0xb3, 0x9e, 0xf1, 0x3c, 0x83, 0x62,
	0x3b, 0xd3, 0x88, 0xdc, 0x4e, 0x48, 0x34, 0x6c, 0x04, 0xa1, 0x4f, 0x7d, 0xb4, 0xc7, 0x22, 0xf3,
	0x80, 0x56, 0x5d, 0x82, 0xb7, 0x03, 0x7f, 0x32, 0xf1, 0xbd, 0x05, 0x47, 0x57, 0xa0, 0xdc, 0x25,
	0xf4, 0xd1, 0x77, 0x09, 0x26, 0xbf, 0xa6, 0x24, 0xa2, 0xfa, 0x09, 0x1c, 0xf1, 0x88, 0x65, 0x3b,
	0xe3, 0x24, 0xdc, 0x84, 0xea, 0x3c, 0x4c, 0x46, 0xcf, 0x43, 0xc7, 0x0f, 0x45, 0x08, 0x9d, 0x42,
	0xe9, 0x79, 0xec, 0x3b, 0xf6, 0x58, 0x95, 0x6b, 0xf2, 0xf5, 0x2e, 0xe6, 0x27, 0xbd, 0x0a, 0x27,
	0x5d, 0x42, 0x31, 0x09, 0xfc, 0x90, 0xa6, 0xb4, 0xce, 0xe0, 0x55, 0x97, 0xd0, 0x0e, 0x19, 0x8f,
	0x22, 0xca, 0x63, 0x29, 0xfc, 0x02, 0xce, 0x13, 0xbc, 0x3d, 0x18, 0x90, 0x80, 0xda, 0xde, 0x20,
	0x7d, 0x1d, 0x0d, 0x54, 0x76, 0xee, 0x05, 0xae, 0x4d, 0xc9, 0x77, 0x1a, 0x12, 0x7b, 0x12, 0x63,
	0x1f, 0xe1, 0x38, 0xfd, 0x06, 0x51, 0xe0, 0x7b, 0x11, 0x41, 0x57, 0xb0, 0xed, 0xf9, 0x2e, 0x89,
	0x54, 0xb9, 0xb6, 0x79, 0xbd, 0xdf, 0x3a, 0x6c, 0x24, 0xfd, 0x69, 0xb0, 0x06, 0x2c, 0x50, 0xfd,
	0x01, 0xd4, 0xfc, 0x9b, 0x72, 0x89, 0x26, 0xec, 0x79, 0x1c, 0x88, 0x65, 0x8e, 0x44, 0x19, 0x8e,
	0xe1, 0x25, 0x4b, 0x37, 0xe0, 0x34, 0xdb, 0x05, 0x2e, 0x56, 0x87, 0x9d, 0x90, 0x85, 0x63, 0xa9,
	0x8a, 0x20, 0xb5, 0x48, 0xc0, 0x31, 0x43, 0x77, 0xe0, 0xf5, 0x8a, 0x9e, 0x71, 0xb5, 0x36, 0x1c,
	0xba, 0x0c, 0xed, 0x87, 0x0b, 0x38, 0x56, 0x55, 0x05, 0xd5, 0x54, 0x3e, 0x2e, 0xbb, 0xe2, 0x31,
	0xd2, 0xc7, 0x50, 0x5b, 0xdd, 0x77, 0x5e, 0xc6, 0x84, 0x0a, 0x2f, 0x63, 0x27, 0x0c, 0x5e, 0xe8,
	0x65, 0xae, 0xd0, 0x52, 0x04, 0x2b, 0x6e, 0x26, 0xa2, 0xff, 0xde, 0x80, 0x7d, 0x61, 0x86, 0xe8,
	0x03, 0xec, 0x4f, 0xd9, 0x53, 0x9f, 0xce, 0x02, 0xc2, 0xbe, 0xa5, 0x72, 0x4b, 0x13, 0x34, 0x05,
	0xb2, 0x35, 0x0b, 0x08, 0x86, 0x69, 0xf2, 0x8c, 0xae, 0x60, 0x6b, 0x3e, 0x3d, 0x75, 0xa3, 0x26,
	0x17, 0x8c, 0xd6, 0x94, 0x30, 0x83, 0x51, 0x13, 0x76, 0xe3, 0xc9, 0xa8, 0x9b, 0x8c, 0x5a, 0x34,
	0x3e, 0x53, 0xc2, 0x09, 0x0d, 0xd5, 0xa1, 0xb4, 0x98, 0x81, 0xba, 0xc5, 0x12, 0xf2, 0x43, 0x32,
	0x25, 0xcc, 0x29, 0xe8, 0x53, 0x7e, 0x08, 0xdb, 0x2c, 0x6b, 0xe5, 0x10, 0x4c, 0x29, 0x3b, 0x06,
	0xf4, 0xa5, 0xa8, 0xc5, 0x25, 0x26, 0xb3, 0xae, 0xc5, 0xa6, 0x94, 0x6f, 0xf2, 0x5d, 0x19, 0x0e,
	0xe8, 0xbc, 0x6d, 0x7d, 0xe2, 0xd1, 0x11, 0x9d, 0xdd, 0x3c, 0xc1, 0x61, 0xa6, 0x8d, 0xe8, 0x18,
	0x14, 0xab, 0x7d, 0x77, 0x6f, 0xf4, 0x7b, 0xdf, 0x3a, 0x6d, 0xcb, 0xe8, 0x3f, 0x1a, 0x4f, 0x8a,
	0x84, 0x4e, 0xa0, 0x92, 0x8a, 0x1a, 0x9d, 0xcf, 0x96, 0x22, 0xa3, 0x2a, 0x1c, 0xa5, 0xc2, 0x1d,
	0xe3, 0xde, 0xb0, 0x0c, 0x65, 0xa3, 0xf5, 0x67, 0x0b, 0xf6, 0xda, 0xf3, 0xbb, 0x3d, 0x90, 0x68,
	0x88, 0xde, 0xc3, 0x0e, 0x5f, 0x41, 0xf4, 0x42, 0xb8, 0x72, 0xda, 0x6a, 0xb4, 0xec, 0x98, 0xd0,
	0x57, 0x38, 0x10, 0x37, 0x17, 0x9d, 0xe5, 0x73, 0x45, 0x17, 0xd0, 0xce, 0x57, 0xe2, 0xfc, 0x6b,
	0xfd, 0x01, 0x4a, 0x76, 0x97, 0x91, 0x9e, 0x49, 0x2a, 0xb0, 0x34, 0xed, 0x72, 0x2d, 0x87, 0x8b,
	0xf7, 0x98, 0x77, 0x0a, 0x9b, 0x8d, 0x6a, 0xe9, 0xb4, 0xbc, 0xf5, 0x69, 0x17, 0x6b, 0x18, 0x5c,
	0xf6, 0x27, 0xb3, 0xcd, 0xfc, 0xa6, 0xa3, 0x37, 0xe9, 0xdc, 0x95, 0xfe, 0xa9, 0x5d, 0xff, 0x9b,
	0xc8, 0x6b, 0x45, 0xcc, 0xeb, 0x0a, 0x37, 0x1e, 0xdd, 0x14, 0xa9, 0x14, 0xdb, 0xb1, 0x56, 0xff,
	0x2f, 0x2e, 0x2f, 0x8a, 0xa1, 0x92, 0xf3, 0x6e, 0x74, 0x59, 0xbc, 0xe8, 0x29, 0x67, 0xd7, 0x4e,
	0x8b, 0x49, 0x6f, 0x65, 0xa7, 0xc4, 0x7e, 0x67, 0xef, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x1b,
	0x6a, 0x92, 0x9c, 0x06, 0x07, 0x00, 0x00,
}
