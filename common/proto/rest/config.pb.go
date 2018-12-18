// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package rest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tree "github.com/pmker/vash/common/proto/tree"
import object "github.com/pmker/vash/common/proto/object"
import ctl "github.com/pmker/vash/common/proto/ctl"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Configuration message. Data is an Json representation of any value
type Configuration struct {
	FullPath string `protobuf:"bytes,1,opt,name=FullPath" json:"FullPath,omitempty"`
	Data     string `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
}

func (m *Configuration) Reset()                    { *m = Configuration{} }
func (m *Configuration) String() string            { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()               {}
func (*Configuration) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Configuration) GetFullPath() string {
	if m != nil {
		return m.FullPath
	}
	return ""
}

func (m *Configuration) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ListDataSourceRequest struct {
}

func (m *ListDataSourceRequest) Reset()                    { *m = ListDataSourceRequest{} }
func (m *ListDataSourceRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDataSourceRequest) ProtoMessage()               {}
func (*ListDataSourceRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// Collection of datasources
type DataSourceCollection struct {
	DataSources []*object.DataSource `protobuf:"bytes,1,rep,name=DataSources" json:"DataSources,omitempty"`
	Total       int32                `protobuf:"varint,2,opt,name=Total" json:"Total,omitempty"`
}

func (m *DataSourceCollection) Reset()                    { *m = DataSourceCollection{} }
func (m *DataSourceCollection) String() string            { return proto.CompactTextString(m) }
func (*DataSourceCollection) ProtoMessage()               {}
func (*DataSourceCollection) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *DataSourceCollection) GetDataSources() []*object.DataSource {
	if m != nil {
		return m.DataSources
	}
	return nil
}

func (m *DataSourceCollection) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type DeleteDataSourceResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *DeleteDataSourceResponse) Reset()                    { *m = DeleteDataSourceResponse{} }
func (m *DeleteDataSourceResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteDataSourceResponse) ProtoMessage()               {}
func (*DeleteDataSourceResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *DeleteDataSourceResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListPeersAddressesRequest struct {
}

func (m *ListPeersAddressesRequest) Reset()                    { *m = ListPeersAddressesRequest{} }
func (m *ListPeersAddressesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPeersAddressesRequest) ProtoMessage()               {}
func (*ListPeersAddressesRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type ListPeersAddressesResponse struct {
	PeerAddresses []string `protobuf:"bytes,1,rep,name=PeerAddresses" json:"PeerAddresses,omitempty"`
}

func (m *ListPeersAddressesResponse) Reset()                    { *m = ListPeersAddressesResponse{} }
func (m *ListPeersAddressesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPeersAddressesResponse) ProtoMessage()               {}
func (*ListPeersAddressesResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *ListPeersAddressesResponse) GetPeerAddresses() []string {
	if m != nil {
		return m.PeerAddresses
	}
	return nil
}

type ListPeerFoldersRequest struct {
	PeerAddress string `protobuf:"bytes,1,opt,name=PeerAddress" json:"PeerAddress,omitempty"`
	Path        string `protobuf:"bytes,2,opt,name=Path" json:"Path,omitempty"`
}

func (m *ListPeerFoldersRequest) Reset()                    { *m = ListPeerFoldersRequest{} }
func (m *ListPeerFoldersRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPeerFoldersRequest) ProtoMessage()               {}
func (*ListPeerFoldersRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *ListPeerFoldersRequest) GetPeerAddress() string {
	if m != nil {
		return m.PeerAddress
	}
	return ""
}

func (m *ListPeerFoldersRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ListVersioningPolicyRequest struct {
}

func (m *ListVersioningPolicyRequest) Reset()                    { *m = ListVersioningPolicyRequest{} }
func (m *ListVersioningPolicyRequest) String() string            { return proto.CompactTextString(m) }
func (*ListVersioningPolicyRequest) ProtoMessage()               {}
func (*ListVersioningPolicyRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

type VersioningPolicyCollection struct {
	Policies []*tree.VersioningPolicy `protobuf:"bytes,1,rep,name=Policies" json:"Policies,omitempty"`
}

func (m *VersioningPolicyCollection) Reset()                    { *m = VersioningPolicyCollection{} }
func (m *VersioningPolicyCollection) String() string            { return proto.CompactTextString(m) }
func (*VersioningPolicyCollection) ProtoMessage()               {}
func (*VersioningPolicyCollection) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *VersioningPolicyCollection) GetPolicies() []*tree.VersioningPolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type ListVirtualNodesRequest struct {
}

func (m *ListVirtualNodesRequest) Reset()                    { *m = ListVirtualNodesRequest{} }
func (m *ListVirtualNodesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListVirtualNodesRequest) ProtoMessage()               {}
func (*ListVirtualNodesRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

type ListServiceRequest struct {
	StatusFilter ctl.ServiceStatus `protobuf:"varint,1,opt,name=StatusFilter,enum=ctl.ServiceStatus" json:"StatusFilter,omitempty"`
}

func (m *ListServiceRequest) Reset()                    { *m = ListServiceRequest{} }
func (m *ListServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*ListServiceRequest) ProtoMessage()               {}
func (*ListServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *ListServiceRequest) GetStatusFilter() ctl.ServiceStatus {
	if m != nil {
		return m.StatusFilter
	}
	return ctl.ServiceStatus_ANY
}

type ServiceCollection struct {
	Services []*ctl.Service `protobuf:"bytes,1,rep,name=Services" json:"Services,omitempty"`
	Total    int32          `protobuf:"varint,2,opt,name=Total" json:"Total,omitempty"`
}

func (m *ServiceCollection) Reset()                    { *m = ServiceCollection{} }
func (m *ServiceCollection) String() string            { return proto.CompactTextString(m) }
func (*ServiceCollection) ProtoMessage()               {}
func (*ServiceCollection) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *ServiceCollection) GetServices() []*ctl.Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ServiceCollection) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type ControlServiceRequest struct {
	ServiceName string             `protobuf:"bytes,1,opt,name=ServiceName" json:"ServiceName,omitempty"`
	NodeName    string             `protobuf:"bytes,2,opt,name=NodeName" json:"NodeName,omitempty"`
	Command     ctl.ServiceCommand `protobuf:"varint,3,opt,name=Command,enum=ctl.ServiceCommand" json:"Command,omitempty"`
}

func (m *ControlServiceRequest) Reset()                    { *m = ControlServiceRequest{} }
func (m *ControlServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*ControlServiceRequest) ProtoMessage()               {}
func (*ControlServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *ControlServiceRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ControlServiceRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *ControlServiceRequest) GetCommand() ctl.ServiceCommand {
	if m != nil {
		return m.Command
	}
	return ctl.ServiceCommand_START
}

type DiscoveryRequest struct {
	EndpointType string `protobuf:"bytes,1,opt,name=EndpointType" json:"EndpointType,omitempty"`
}

func (m *DiscoveryRequest) Reset()                    { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string            { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()               {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func (m *DiscoveryRequest) GetEndpointType() string {
	if m != nil {
		return m.EndpointType
	}
	return ""
}

type DiscoveryResponse struct {
	PackageType   string            `protobuf:"bytes,1,opt,name=PackageType" json:"PackageType,omitempty"`
	PackageLabel  string            `protobuf:"bytes,2,opt,name=PackageLabel" json:"PackageLabel,omitempty"`
	Version       string            `protobuf:"bytes,3,opt,name=Version" json:"Version,omitempty"`
	BuildStamp    int32             `protobuf:"varint,4,opt,name=BuildStamp" json:"BuildStamp,omitempty"`
	BuildRevision string            `protobuf:"bytes,5,opt,name=BuildRevision" json:"BuildRevision,omitempty"`
	Endpoints     map[string]string `protobuf:"bytes,6,rep,name=Endpoints" json:"Endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DiscoveryResponse) Reset()                    { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string            { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()               {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

func (m *DiscoveryResponse) GetPackageType() string {
	if m != nil {
		return m.PackageType
	}
	return ""
}

func (m *DiscoveryResponse) GetPackageLabel() string {
	if m != nil {
		return m.PackageLabel
	}
	return ""
}

func (m *DiscoveryResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DiscoveryResponse) GetBuildStamp() int32 {
	if m != nil {
		return m.BuildStamp
	}
	return 0
}

func (m *DiscoveryResponse) GetBuildRevision() string {
	if m != nil {
		return m.BuildRevision
	}
	return ""
}

func (m *DiscoveryResponse) GetEndpoints() map[string]string {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type ConfigFormRequest struct {
	ServiceName string `protobuf:"bytes,1,opt,name=ServiceName" json:"ServiceName,omitempty"`
}

func (m *ConfigFormRequest) Reset()                    { *m = ConfigFormRequest{} }
func (m *ConfigFormRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigFormRequest) ProtoMessage()               {}
func (*ConfigFormRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{15} }

func (m *ConfigFormRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type OpenApiResponse struct {
}

func (m *OpenApiResponse) Reset()                    { *m = OpenApiResponse{} }
func (m *OpenApiResponse) String() string            { return proto.CompactTextString(m) }
func (*OpenApiResponse) ProtoMessage()               {}
func (*OpenApiResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{16} }

func init() {
	proto.RegisterType((*Configuration)(nil), "rest.Configuration")
	proto.RegisterType((*ListDataSourceRequest)(nil), "rest.ListDataSourceRequest")
	proto.RegisterType((*DataSourceCollection)(nil), "rest.DataSourceCollection")
	proto.RegisterType((*DeleteDataSourceResponse)(nil), "rest.DeleteDataSourceResponse")
	proto.RegisterType((*ListPeersAddressesRequest)(nil), "rest.ListPeersAddressesRequest")
	proto.RegisterType((*ListPeersAddressesResponse)(nil), "rest.ListPeersAddressesResponse")
	proto.RegisterType((*ListPeerFoldersRequest)(nil), "rest.ListPeerFoldersRequest")
	proto.RegisterType((*ListVersioningPolicyRequest)(nil), "rest.ListVersioningPolicyRequest")
	proto.RegisterType((*VersioningPolicyCollection)(nil), "rest.VersioningPolicyCollection")
	proto.RegisterType((*ListVirtualNodesRequest)(nil), "rest.ListVirtualNodesRequest")
	proto.RegisterType((*ListServiceRequest)(nil), "rest.ListServiceRequest")
	proto.RegisterType((*ServiceCollection)(nil), "rest.ServiceCollection")
	proto.RegisterType((*ControlServiceRequest)(nil), "rest.ControlServiceRequest")
	proto.RegisterType((*DiscoveryRequest)(nil), "rest.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "rest.DiscoveryResponse")
	proto.RegisterType((*ConfigFormRequest)(nil), "rest.ConfigFormRequest")
	proto.RegisterType((*OpenApiResponse)(nil), "rest.OpenApiResponse")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5b, 0x6f, 0x1a, 0x3b,
	0x10, 0x16, 0xb9, 0xc2, 0x40, 0x72, 0x82, 0x4f, 0x2e, 0x1b, 0xa2, 0x73, 0x84, 0x56, 0x55, 0x95,
	0x97, 0x2e, 0x6a, 0x92, 0xa6, 0x55, 0x55, 0xa9, 0x4a, 0x20, 0x3c, 0x45, 0x29, 0x5a, 0xa2, 0xbe,
	0x2f, 0xbb, 0x53, 0xe2, 0xc6, 0xbb, 0xde, 0xda, 0x5e, 0x24, 0xde, 0xfb, 0xb7, 0xfa, 0xdf, 0x2a,
	0x1b, 0xef, 0x62, 0x92, 0x56, 0xca, 0x03, 0xe0, 0xf9, 0xbe, 0x99, 0xf1, 0x37, 0x17, 0x0c, 0xad,
	0x98, 0x67, 0xdf, 0xe8, 0x34, 0xc8, 0x05, 0x57, 0x9c, 0x6c, 0x08, 0x94, 0xaa, 0x73, 0x3e, 0xa5,
	0xea, 0xa1, 0x98, 0x04, 0x31, 0x4f, 0x7b, 0xf9, 0x3c, 0xa1, 0xbc, 0x17, 0x23, 0x63, 0xb2, 0x17,
	0xf3, 0x34, 0xe5, 0x59, 0xcf, 0xb8, 0xf6, 0x94, 0x40, 0x34, 0x5f, 0x8b, 0xd0, 0xce, 0xfb, 0x97,
	0x04, 0xf1, 0xc9, 0x77, 0x8c, 0x95, 0xfd, 0xb1, 0x81, 0x6f, 0x5f, 0x12, 0x18, 0x2b, 0xa6, 0x3f,
	0x8b, 0x10, 0xff, 0x33, 0xec, 0xf4, 0x8d, 0xec, 0x42, 0x44, 0x8a, 0xf2, 0x8c, 0x74, 0xa0, 0x3e,
	0x2c, 0x18, 0x1b, 0x45, 0xea, 0xc1, 0xab, 0x75, 0x6b, 0xa7, 0x8d, 0xb0, 0xb2, 0x09, 0x81, 0x8d,
	0x41, 0xa4, 0x22, 0x6f, 0xcd, 0xe0, 0xe6, 0xec, 0x1f, 0xc1, 0xc1, 0x2d, 0x95, 0x4a, 0x9f, 0xc7,
	0xbc, 0x10, 0x31, 0x86, 0xf8, 0xa3, 0x40, 0xa9, 0xfc, 0x09, 0xec, 0x2f, 0xc1, 0x3e, 0x67, 0x0c,
	0x63, 0x73, 0xc1, 0x05, 0x34, 0x97, 0xb8, 0xf4, 0x6a, 0xdd, 0xf5, 0xd3, 0xe6, 0x19, 0x09, 0x6c,
	0x21, 0x4e, 0x1e, 0xd7, 0x8d, 0xec, 0xc3, 0xe6, 0x3d, 0x57, 0x11, 0x33, 0x77, 0x6f, 0x86, 0x0b,
	0xc3, 0xbf, 0x00, 0x6f, 0x80, 0x0c, 0x15, 0xba, 0xd7, 0xcb, 0x9c, 0x67, 0x12, 0x89, 0x07, 0xdb,
	0xe3, 0x22, 0x8e, 0x51, 0x4a, 0x53, 0x47, 0x3d, 0x2c, 0x4d, 0xff, 0x04, 0x8e, 0xb5, 0xe4, 0x11,
	0xa2, 0x90, 0x57, 0x49, 0x22, 0x50, 0x4a, 0x94, 0xa5, 0xec, 0x6b, 0xe8, 0xfc, 0x89, 0xb4, 0x49,
	0x5f, 0xc1, 0x8e, 0x66, 0x2a, 0xc2, 0xc8, 0x6f, 0x84, 0xab, 0xa0, 0x7f, 0x07, 0x87, 0x65, 0x8e,
	0x21, 0x67, 0x09, 0x8a, 0x32, 0x3b, 0xe9, 0x42, 0xd3, 0x71, 0xb5, 0x0d, 0x76, 0x21, 0xdd, 0x63,
	0xd3, 0x7b, 0xdb, 0x63, 0x7d, 0xf6, 0xff, 0x83, 0x13, 0x9d, 0xef, 0x2b, 0x0a, 0x49, 0x79, 0x46,
	0xb3, 0xe9, 0x88, 0x33, 0x1a, 0xcf, 0x4b, 0xc9, 0x23, 0xe8, 0x3c, 0xa5, 0x9c, 0x7e, 0x9f, 0x41,
	0xdd, 0x60, 0xb4, 0x6a, 0xf6, 0x61, 0x60, 0x96, 0xed, 0x59, 0xba, 0xca, 0xcf, 0x3f, 0x86, 0x23,
	0x73, 0x21, 0x15, 0xaa, 0x88, 0xd8, 0x1d, 0x4f, 0x96, 0xfd, 0xb9, 0x05, 0xa2, 0xa9, 0x31, 0x8a,
	0x19, 0xad, 0x86, 0x4d, 0x2e, 0xa1, 0x35, 0x56, 0x91, 0x2a, 0xe4, 0x90, 0x32, 0x85, 0xc2, 0x14,
	0xb6, 0x7b, 0x46, 0x02, 0xbd, 0x68, 0xd6, 0x75, 0xc1, 0x87, 0x2b, 0x7e, 0xfe, 0x18, 0xda, 0x96,
	0x76, 0x14, 0x9f, 0x42, 0xdd, 0x82, 0xa5, 0xe2, 0x96, 0x9b, 0x28, 0xac, 0xd8, 0xbf, 0x6c, 0xc5,
	0xcf, 0x1a, 0x1c, 0xf4, 0x79, 0xa6, 0x04, 0x67, 0x4f, 0x64, 0x76, 0xa1, 0x69, 0x91, 0xbb, 0x28,
	0xc5, 0xb2, 0xfd, 0x0e, 0xa4, 0xd7, 0x5f, 0x97, 0x6b, 0xe8, 0xc5, 0x08, 0x2a, 0x9b, 0xbc, 0x81,
	0xed, 0x3e, 0x4f, 0xd3, 0x28, 0x4b, 0xbc, 0x75, 0x53, 0xdf, 0xbf, 0xae, 0x2c, 0x4b, 0x85, 0xa5,
	0x8f, 0x7f, 0x09, 0x7b, 0x03, 0x2a, 0x63, 0x3e, 0x43, 0x51, 0x8e, 0x8a, 0xf8, 0xd0, 0xba, 0xc9,
	0x92, 0x9c, 0xd3, 0x4c, 0xdd, 0xcf, 0xf3, 0x52, 0xc1, 0x0a, 0xe6, 0xff, 0x5a, 0x83, 0xb6, 0x13,
	0x68, 0x37, 0x4f, 0x6f, 0x4e, 0x14, 0x3f, 0x46, 0x53, 0x74, 0x02, 0x5d, 0x48, 0xe7, 0xb6, 0xe6,
	0x6d, 0x34, 0x41, 0x66, 0xe5, 0xaf, 0x60, 0xfa, 0x4f, 0x61, 0xc7, 0x6e, 0x4a, 0x68, 0x84, 0xa5,
	0x49, 0xfe, 0x07, 0xb8, 0x2e, 0x28, 0x4b, 0xc6, 0x2a, 0x4a, 0x73, 0x6f, 0xc3, 0xf4, 0xd3, 0x41,
	0xf4, 0xe6, 0x1b, 0x2b, 0xc4, 0x19, 0x35, 0xf1, 0x9b, 0x26, 0x7e, 0x15, 0x24, 0x03, 0x68, 0x94,
	0xb5, 0x48, 0x6f, 0xcb, 0xcc, 0xee, 0x75, 0xa0, 0x5f, 0xc2, 0xe0, 0x59, 0x45, 0x41, 0xe5, 0x78,
	0x93, 0x29, 0x31, 0x0f, 0x97, 0x81, 0x9d, 0x4f, 0xb0, 0xbb, 0x4a, 0x92, 0x3d, 0x58, 0x7f, 0xc4,
	0xb9, 0xad, 0x5a, 0x1f, 0xf5, 0xe8, 0x67, 0x11, 0x2b, 0xca, 0x29, 0x2d, 0x8c, 0x8f, 0x6b, 0x1f,
	0x6a, 0xfe, 0x3b, 0x68, 0x2f, 0x9e, 0xb4, 0x21, 0x17, 0xe9, 0x8b, 0x27, 0xef, 0xb7, 0xe1, 0x9f,
	0x2f, 0x39, 0x66, 0x57, 0x39, 0x2d, 0x15, 0x4e, 0xb6, 0xcc, 0x1b, 0x79, 0xfe, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0xad, 0x87, 0xcc, 0x3f, 0xda, 0x05, 0x00, 0x00,
}
