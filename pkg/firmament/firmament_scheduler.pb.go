// Code generated by protoc-gen-go.
// source: firmament_scheduler.proto
// DO NOT EDIT!

package firmament

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

type TaskReplyType int32

const (
	TaskReplyType_TASK_COMPLETED_OK      TaskReplyType = 0
	TaskReplyType_TASK_SUBMITTED_OK      TaskReplyType = 1
	TaskReplyType_TASK_REMOVED_OK        TaskReplyType = 2
	TaskReplyType_TASK_FAILED_OK         TaskReplyType = 3
	TaskReplyType_TASK_UPDATED_OK        TaskReplyType = 4
	TaskReplyType_TASK_NOT_FOUND         TaskReplyType = 5
	TaskReplyType_TASK_JOB_NOT_FOUND     TaskReplyType = 6
	TaskReplyType_TASK_ALREADY_SUBMITTED TaskReplyType = 7
	TaskReplyType_TASK_STATE_NOT_CREATED TaskReplyType = 8
)

var TaskReplyType_name = map[int32]string{
	0: "TASK_COMPLETED_OK",
	1: "TASK_SUBMITTED_OK",
	2: "TASK_REMOVED_OK",
	3: "TASK_FAILED_OK",
	4: "TASK_UPDATED_OK",
	5: "TASK_NOT_FOUND",
	6: "TASK_JOB_NOT_FOUND",
	7: "TASK_ALREADY_SUBMITTED",
	8: "TASK_STATE_NOT_CREATED",
}
var TaskReplyType_value = map[string]int32{
	"TASK_COMPLETED_OK":      0,
	"TASK_SUBMITTED_OK":      1,
	"TASK_REMOVED_OK":        2,
	"TASK_FAILED_OK":         3,
	"TASK_UPDATED_OK":        4,
	"TASK_NOT_FOUND":         5,
	"TASK_JOB_NOT_FOUND":     6,
	"TASK_ALREADY_SUBMITTED": 7,
	"TASK_STATE_NOT_CREATED": 8,
}

func (x TaskReplyType) String() string {
	return proto.EnumName(TaskReplyType_name, int32(x))
}
func (TaskReplyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type NodeReplyType int32

const (
	NodeReplyType_NODE_ADDED_OK       NodeReplyType = 0
	NodeReplyType_NODE_FAILED_OK      NodeReplyType = 1
	NodeReplyType_NODE_REMOVED_OK     NodeReplyType = 2
	NodeReplyType_NODE_UPDATED_OK     NodeReplyType = 3
	NodeReplyType_NODE_NOT_FOUND      NodeReplyType = 4
	NodeReplyType_NODE_ALREADY_EXISTS NodeReplyType = 5
)

var NodeReplyType_name = map[int32]string{
	0: "NODE_ADDED_OK",
	1: "NODE_FAILED_OK",
	2: "NODE_REMOVED_OK",
	3: "NODE_UPDATED_OK",
	4: "NODE_NOT_FOUND",
	5: "NODE_ALREADY_EXISTS",
}
var NodeReplyType_value = map[string]int32{
	"NODE_ADDED_OK":       0,
	"NODE_FAILED_OK":      1,
	"NODE_REMOVED_OK":     2,
	"NODE_UPDATED_OK":     3,
	"NODE_NOT_FOUND":      4,
	"NODE_ALREADY_EXISTS": 5,
}

func (x NodeReplyType) String() string {
	return proto.EnumName(NodeReplyType_name, int32(x))
}
func (NodeReplyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type ScheduleRequest struct {
}

func (m *ScheduleRequest) Reset()                    { *m = ScheduleRequest{} }
func (m *ScheduleRequest) String() string            { return proto.CompactTextString(m) }
func (*ScheduleRequest) ProtoMessage()               {}
func (*ScheduleRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type SchedulingDeltas struct {
	Deltas []*SchedulingDelta `protobuf:"bytes,1,rep,name=deltas" json:"deltas,omitempty"`
}

func (m *SchedulingDeltas) Reset()                    { *m = SchedulingDeltas{} }
func (m *SchedulingDeltas) String() string            { return proto.CompactTextString(m) }
func (*SchedulingDeltas) ProtoMessage()               {}
func (*SchedulingDeltas) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SchedulingDeltas) GetDeltas() []*SchedulingDelta {
	if m != nil {
		return m.Deltas
	}
	return nil
}

type TaskCompletedResponse struct {
	Type TaskReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.TaskReplyType" json:"type,omitempty"`
}

func (m *TaskCompletedResponse) Reset()                    { *m = TaskCompletedResponse{} }
func (m *TaskCompletedResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskCompletedResponse) ProtoMessage()               {}
func (*TaskCompletedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *TaskCompletedResponse) GetType() TaskReplyType {
	if m != nil {
		return m.Type
	}
	return TaskReplyType_TASK_COMPLETED_OK
}

type TaskDescription struct {
	TaskDescriptor *TaskDescriptor `protobuf:"bytes,1,opt,name=task_descriptor,json=taskDescriptor" json:"task_descriptor,omitempty"`
	JobDescriptor  *JobDescriptor  `protobuf:"bytes,2,opt,name=job_descriptor,json=jobDescriptor" json:"job_descriptor,omitempty"`
}

func (m *TaskDescription) Reset()                    { *m = TaskDescription{} }
func (m *TaskDescription) String() string            { return proto.CompactTextString(m) }
func (*TaskDescription) ProtoMessage()               {}
func (*TaskDescription) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *TaskDescription) GetTaskDescriptor() *TaskDescriptor {
	if m != nil {
		return m.TaskDescriptor
	}
	return nil
}

func (m *TaskDescription) GetJobDescriptor() *JobDescriptor {
	if m != nil {
		return m.JobDescriptor
	}
	return nil
}

type TaskSubmittedResponse struct {
	Type TaskReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.TaskReplyType" json:"type,omitempty"`
}

func (m *TaskSubmittedResponse) Reset()                    { *m = TaskSubmittedResponse{} }
func (m *TaskSubmittedResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskSubmittedResponse) ProtoMessage()               {}
func (*TaskSubmittedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *TaskSubmittedResponse) GetType() TaskReplyType {
	if m != nil {
		return m.Type
	}
	return TaskReplyType_TASK_COMPLETED_OK
}

type TaskRemovedResponse struct {
	Type TaskReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.TaskReplyType" json:"type,omitempty"`
}

func (m *TaskRemovedResponse) Reset()                    { *m = TaskRemovedResponse{} }
func (m *TaskRemovedResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskRemovedResponse) ProtoMessage()               {}
func (*TaskRemovedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *TaskRemovedResponse) GetType() TaskReplyType {
	if m != nil {
		return m.Type
	}
	return TaskReplyType_TASK_COMPLETED_OK
}

type TaskFailedResponse struct {
	Type TaskReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.TaskReplyType" json:"type,omitempty"`
}

func (m *TaskFailedResponse) Reset()                    { *m = TaskFailedResponse{} }
func (m *TaskFailedResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskFailedResponse) ProtoMessage()               {}
func (*TaskFailedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *TaskFailedResponse) GetType() TaskReplyType {
	if m != nil {
		return m.Type
	}
	return TaskReplyType_TASK_COMPLETED_OK
}

type TaskUpdatedResponse struct {
	Type TaskReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.TaskReplyType" json:"type,omitempty"`
}

func (m *TaskUpdatedResponse) Reset()                    { *m = TaskUpdatedResponse{} }
func (m *TaskUpdatedResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskUpdatedResponse) ProtoMessage()               {}
func (*TaskUpdatedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *TaskUpdatedResponse) GetType() TaskReplyType {
	if m != nil {
		return m.Type
	}
	return TaskReplyType_TASK_COMPLETED_OK
}

type NodeAddedResponse struct {
	Type NodeReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.NodeReplyType" json:"type,omitempty"`
}

func (m *NodeAddedResponse) Reset()                    { *m = NodeAddedResponse{} }
func (m *NodeAddedResponse) String() string            { return proto.CompactTextString(m) }
func (*NodeAddedResponse) ProtoMessage()               {}
func (*NodeAddedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *NodeAddedResponse) GetType() NodeReplyType {
	if m != nil {
		return m.Type
	}
	return NodeReplyType_NODE_ADDED_OK
}

type NodeRemovedResponse struct {
	Type NodeReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.NodeReplyType" json:"type,omitempty"`
}

func (m *NodeRemovedResponse) Reset()                    { *m = NodeRemovedResponse{} }
func (m *NodeRemovedResponse) String() string            { return proto.CompactTextString(m) }
func (*NodeRemovedResponse) ProtoMessage()               {}
func (*NodeRemovedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *NodeRemovedResponse) GetType() NodeReplyType {
	if m != nil {
		return m.Type
	}
	return NodeReplyType_NODE_ADDED_OK
}

type NodeFailedResponse struct {
	Type NodeReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.NodeReplyType" json:"type,omitempty"`
}

func (m *NodeFailedResponse) Reset()                    { *m = NodeFailedResponse{} }
func (m *NodeFailedResponse) String() string            { return proto.CompactTextString(m) }
func (*NodeFailedResponse) ProtoMessage()               {}
func (*NodeFailedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *NodeFailedResponse) GetType() NodeReplyType {
	if m != nil {
		return m.Type
	}
	return NodeReplyType_NODE_ADDED_OK
}

type NodeUpdatedResponse struct {
	Type NodeReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.NodeReplyType" json:"type,omitempty"`
}

func (m *NodeUpdatedResponse) Reset()                    { *m = NodeUpdatedResponse{} }
func (m *NodeUpdatedResponse) String() string            { return proto.CompactTextString(m) }
func (*NodeUpdatedResponse) ProtoMessage()               {}
func (*NodeUpdatedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *NodeUpdatedResponse) GetType() NodeReplyType {
	if m != nil {
		return m.Type
	}
	return NodeReplyType_NODE_ADDED_OK
}

type TaskStatsResponse struct {
	Type TaskReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.TaskReplyType" json:"type,omitempty"`
}

func (m *TaskStatsResponse) Reset()                    { *m = TaskStatsResponse{} }
func (m *TaskStatsResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskStatsResponse) ProtoMessage()               {}
func (*TaskStatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *TaskStatsResponse) GetType() TaskReplyType {
	if m != nil {
		return m.Type
	}
	return TaskReplyType_TASK_COMPLETED_OK
}

type ResourceStatsResponse struct {
	Type NodeReplyType `protobuf:"varint,1,opt,name=type,enum=firmament.NodeReplyType" json:"type,omitempty"`
}

func (m *ResourceStatsResponse) Reset()                    { *m = ResourceStatsResponse{} }
func (m *ResourceStatsResponse) String() string            { return proto.CompactTextString(m) }
func (*ResourceStatsResponse) ProtoMessage()               {}
func (*ResourceStatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *ResourceStatsResponse) GetType() NodeReplyType {
	if m != nil {
		return m.Type
	}
	return NodeReplyType_NODE_ADDED_OK
}

type TaskUID struct {
	TaskUid uint64 `protobuf:"varint,1,opt,name=task_uid,json=taskUid" json:"task_uid,omitempty"`
}

func (m *TaskUID) Reset()                    { *m = TaskUID{} }
func (m *TaskUID) String() string            { return proto.CompactTextString(m) }
func (*TaskUID) ProtoMessage()               {}
func (*TaskUID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *TaskUID) GetTaskUid() uint64 {
	if m != nil {
		return m.TaskUid
	}
	return 0
}

type ResourceUID struct {
	ResourceUid string `protobuf:"bytes,1,opt,name=resource_uid,json=resourceUid" json:"resource_uid,omitempty"`
}

func (m *ResourceUID) Reset()                    { *m = ResourceUID{} }
func (m *ResourceUID) String() string            { return proto.CompactTextString(m) }
func (*ResourceUID) ProtoMessage()               {}
func (*ResourceUID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *ResourceUID) GetResourceUid() string {
	if m != nil {
		return m.ResourceUid
	}
	return ""
}

func init() {
	proto.RegisterType((*ScheduleRequest)(nil), "firmament.ScheduleRequest")
	proto.RegisterType((*SchedulingDeltas)(nil), "firmament.SchedulingDeltas")
	proto.RegisterType((*TaskCompletedResponse)(nil), "firmament.TaskCompletedResponse")
	proto.RegisterType((*TaskDescription)(nil), "firmament.TaskDescription")
	proto.RegisterType((*TaskSubmittedResponse)(nil), "firmament.TaskSubmittedResponse")
	proto.RegisterType((*TaskRemovedResponse)(nil), "firmament.TaskRemovedResponse")
	proto.RegisterType((*TaskFailedResponse)(nil), "firmament.TaskFailedResponse")
	proto.RegisterType((*TaskUpdatedResponse)(nil), "firmament.TaskUpdatedResponse")
	proto.RegisterType((*NodeAddedResponse)(nil), "firmament.NodeAddedResponse")
	proto.RegisterType((*NodeRemovedResponse)(nil), "firmament.NodeRemovedResponse")
	proto.RegisterType((*NodeFailedResponse)(nil), "firmament.NodeFailedResponse")
	proto.RegisterType((*NodeUpdatedResponse)(nil), "firmament.NodeUpdatedResponse")
	proto.RegisterType((*TaskStatsResponse)(nil), "firmament.TaskStatsResponse")
	proto.RegisterType((*ResourceStatsResponse)(nil), "firmament.ResourceStatsResponse")
	proto.RegisterType((*TaskUID)(nil), "firmament.TaskUID")
	proto.RegisterType((*ResourceUID)(nil), "firmament.ResourceUID")
	proto.RegisterEnum("firmament.TaskReplyType", TaskReplyType_name, TaskReplyType_value)
	proto.RegisterEnum("firmament.NodeReplyType", NodeReplyType_name, NodeReplyType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FirmamentScheduler service

type FirmamentSchedulerClient interface {
	Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*SchedulingDeltas, error)
	TaskCompleted(ctx context.Context, in *TaskUID, opts ...grpc.CallOption) (*TaskCompletedResponse, error)
	TaskFailed(ctx context.Context, in *TaskUID, opts ...grpc.CallOption) (*TaskFailedResponse, error)
	TaskRemoved(ctx context.Context, in *TaskUID, opts ...grpc.CallOption) (*TaskRemovedResponse, error)
	TaskSubmitted(ctx context.Context, in *TaskDescription, opts ...grpc.CallOption) (*TaskSubmittedResponse, error)
	TaskUpdated(ctx context.Context, in *TaskDescription, opts ...grpc.CallOption) (*TaskUpdatedResponse, error)
	NodeAdded(ctx context.Context, in *ResourceTopologyNodeDescriptor, opts ...grpc.CallOption) (*NodeAddedResponse, error)
	NodeFailed(ctx context.Context, in *ResourceUID, opts ...grpc.CallOption) (*NodeFailedResponse, error)
	NodeRemoved(ctx context.Context, in *ResourceUID, opts ...grpc.CallOption) (*NodeRemovedResponse, error)
	NodeUpdated(ctx context.Context, in *ResourceTopologyNodeDescriptor, opts ...grpc.CallOption) (*NodeUpdatedResponse, error)
	AddTaskStats(ctx context.Context, in *TaskStats, opts ...grpc.CallOption) (*TaskStatsResponse, error)
	AddNodeStats(ctx context.Context, in *ResourceStats, opts ...grpc.CallOption) (*ResourceStatsResponse, error)
}

type firmamentSchedulerClient struct {
	cc *grpc.ClientConn
}

func NewFirmamentSchedulerClient(cc *grpc.ClientConn) FirmamentSchedulerClient {
	return &firmamentSchedulerClient{cc}
}

func (c *firmamentSchedulerClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*SchedulingDeltas, error) {
	out := new(SchedulingDeltas)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/Schedule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) TaskCompleted(ctx context.Context, in *TaskUID, opts ...grpc.CallOption) (*TaskCompletedResponse, error) {
	out := new(TaskCompletedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/TaskCompleted", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) TaskFailed(ctx context.Context, in *TaskUID, opts ...grpc.CallOption) (*TaskFailedResponse, error) {
	out := new(TaskFailedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/TaskFailed", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) TaskRemoved(ctx context.Context, in *TaskUID, opts ...grpc.CallOption) (*TaskRemovedResponse, error) {
	out := new(TaskRemovedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/TaskRemoved", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) TaskSubmitted(ctx context.Context, in *TaskDescription, opts ...grpc.CallOption) (*TaskSubmittedResponse, error) {
	out := new(TaskSubmittedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/TaskSubmitted", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) TaskUpdated(ctx context.Context, in *TaskDescription, opts ...grpc.CallOption) (*TaskUpdatedResponse, error) {
	out := new(TaskUpdatedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/TaskUpdated", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) NodeAdded(ctx context.Context, in *ResourceTopologyNodeDescriptor, opts ...grpc.CallOption) (*NodeAddedResponse, error) {
	out := new(NodeAddedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/NodeAdded", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) NodeFailed(ctx context.Context, in *ResourceUID, opts ...grpc.CallOption) (*NodeFailedResponse, error) {
	out := new(NodeFailedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/NodeFailed", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) NodeRemoved(ctx context.Context, in *ResourceUID, opts ...grpc.CallOption) (*NodeRemovedResponse, error) {
	out := new(NodeRemovedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/NodeRemoved", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) NodeUpdated(ctx context.Context, in *ResourceTopologyNodeDescriptor, opts ...grpc.CallOption) (*NodeUpdatedResponse, error) {
	out := new(NodeUpdatedResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/NodeUpdated", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) AddTaskStats(ctx context.Context, in *TaskStats, opts ...grpc.CallOption) (*TaskStatsResponse, error) {
	out := new(TaskStatsResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/AddTaskStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmamentSchedulerClient) AddNodeStats(ctx context.Context, in *ResourceStats, opts ...grpc.CallOption) (*ResourceStatsResponse, error) {
	out := new(ResourceStatsResponse)
	err := grpc.Invoke(ctx, "/firmament.FirmamentScheduler/AddNodeStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FirmamentScheduler service

type FirmamentSchedulerServer interface {
	Schedule(context.Context, *ScheduleRequest) (*SchedulingDeltas, error)
	TaskCompleted(context.Context, *TaskUID) (*TaskCompletedResponse, error)
	TaskFailed(context.Context, *TaskUID) (*TaskFailedResponse, error)
	TaskRemoved(context.Context, *TaskUID) (*TaskRemovedResponse, error)
	TaskSubmitted(context.Context, *TaskDescription) (*TaskSubmittedResponse, error)
	TaskUpdated(context.Context, *TaskDescription) (*TaskUpdatedResponse, error)
	NodeAdded(context.Context, *ResourceTopologyNodeDescriptor) (*NodeAddedResponse, error)
	NodeFailed(context.Context, *ResourceUID) (*NodeFailedResponse, error)
	NodeRemoved(context.Context, *ResourceUID) (*NodeRemovedResponse, error)
	NodeUpdated(context.Context, *ResourceTopologyNodeDescriptor) (*NodeUpdatedResponse, error)
	AddTaskStats(context.Context, *TaskStats) (*TaskStatsResponse, error)
	AddNodeStats(context.Context, *ResourceStats) (*ResourceStatsResponse, error)
}

func RegisterFirmamentSchedulerServer(s *grpc.Server, srv FirmamentSchedulerServer) {
	s.RegisterService(&_FirmamentScheduler_serviceDesc, srv)
}

func _FirmamentScheduler_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).Schedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/Schedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).Schedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_TaskCompleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).TaskCompleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/TaskCompleted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).TaskCompleted(ctx, req.(*TaskUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_TaskFailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).TaskFailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/TaskFailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).TaskFailed(ctx, req.(*TaskUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_TaskRemoved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).TaskRemoved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/TaskRemoved",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).TaskRemoved(ctx, req.(*TaskUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_TaskSubmitted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDescription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).TaskSubmitted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/TaskSubmitted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).TaskSubmitted(ctx, req.(*TaskDescription))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_TaskUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDescription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).TaskUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/TaskUpdated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).TaskUpdated(ctx, req.(*TaskDescription))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_NodeAdded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceTopologyNodeDescriptor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).NodeAdded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/NodeAdded",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).NodeAdded(ctx, req.(*ResourceTopologyNodeDescriptor))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_NodeFailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).NodeFailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/NodeFailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).NodeFailed(ctx, req.(*ResourceUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_NodeRemoved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).NodeRemoved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/NodeRemoved",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).NodeRemoved(ctx, req.(*ResourceUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_NodeUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceTopologyNodeDescriptor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).NodeUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/NodeUpdated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).NodeUpdated(ctx, req.(*ResourceTopologyNodeDescriptor))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_AddTaskStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskStats)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).AddTaskStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/AddTaskStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).AddTaskStats(ctx, req.(*TaskStats))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmamentScheduler_AddNodeStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceStats)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmamentSchedulerServer).AddNodeStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmament.FirmamentScheduler/AddNodeStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmamentSchedulerServer).AddNodeStats(ctx, req.(*ResourceStats))
	}
	return interceptor(ctx, in, info, handler)
}

var _FirmamentScheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "firmament.FirmamentScheduler",
	HandlerType: (*FirmamentSchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Schedule",
			Handler:    _FirmamentScheduler_Schedule_Handler,
		},
		{
			MethodName: "TaskCompleted",
			Handler:    _FirmamentScheduler_TaskCompleted_Handler,
		},
		{
			MethodName: "TaskFailed",
			Handler:    _FirmamentScheduler_TaskFailed_Handler,
		},
		{
			MethodName: "TaskRemoved",
			Handler:    _FirmamentScheduler_TaskRemoved_Handler,
		},
		{
			MethodName: "TaskSubmitted",
			Handler:    _FirmamentScheduler_TaskSubmitted_Handler,
		},
		{
			MethodName: "TaskUpdated",
			Handler:    _FirmamentScheduler_TaskUpdated_Handler,
		},
		{
			MethodName: "NodeAdded",
			Handler:    _FirmamentScheduler_NodeAdded_Handler,
		},
		{
			MethodName: "NodeFailed",
			Handler:    _FirmamentScheduler_NodeFailed_Handler,
		},
		{
			MethodName: "NodeRemoved",
			Handler:    _FirmamentScheduler_NodeRemoved_Handler,
		},
		{
			MethodName: "NodeUpdated",
			Handler:    _FirmamentScheduler_NodeUpdated_Handler,
		},
		{
			MethodName: "AddTaskStats",
			Handler:    _FirmamentScheduler_AddTaskStats_Handler,
		},
		{
			MethodName: "AddNodeStats",
			Handler:    _FirmamentScheduler_AddNodeStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "firmament_scheduler.proto",
}

func init() { proto.RegisterFile("firmament_scheduler.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 819 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0xe2, 0x46,
	0x18, 0x35, 0x1b, 0x36, 0x24, 0x1f, 0xcb, 0xdf, 0x97, 0x0d, 0x25, 0xb4, 0x5d, 0x11, 0xab, 0x17,
	0x69, 0x54, 0x45, 0x15, 0x7d, 0x80, 0xca, 0x60, 0x53, 0x91, 0x10, 0x1c, 0xd9, 0x26, 0x6a, 0x7b,
	0x63, 0x01, 0x9e, 0xa6, 0x4e, 0x00, 0xbb, 0xb6, 0xa9, 0xc4, 0x43, 0xf4, 0xb6, 0x0f, 0xd0, 0x17,
	0xeb, 0xab, 0x54, 0x1e, 0xdb, 0xe3, 0xdf, 0x44, 0x94, 0xbd, 0x9c, 0x33, 0x67, 0xce, 0x9c, 0xef,
	0x1b, 0xcf, 0x19, 0xc3, 0xc5, 0x6f, 0xa6, 0xb3, 0x9e, 0xaf, 0xc9, 0xc6, 0xd3, 0xdd, 0xe5, 0xef,
	0xc4, 0xd8, 0xae, 0x88, 0x73, 0x63, 0x3b, 0x96, 0x67, 0xe1, 0x29, 0x9b, 0xea, 0xd6, 0x9f, 0xad,
	0x85, 0x6e, 0x10, 0x77, 0x19, 0x4c, 0x75, 0x3f, 0x3a, 0xc4, 0xb5, 0xb6, 0xce, 0x92, 0xe8, 0xae,
	0x37, 0xf7, 0xdc, 0x10, 0xbd, 0x64, 0xa8, 0x67, 0xd9, 0xd6, 0xca, 0x7a, 0xda, 0xe9, 0x1b, 0xcb,
	0x20, 0xc9, 0x85, 0xed, 0x70, 0x13, 0x73, 0xf3, 0xa4, 0x1b, 0x64, 0xe5, 0xcd, 0x43, 0xbc, 0xe1,
	0xcd, 0xdd, 0x97, 0x24, 0xb1, 0x49, 0x81, 0x84, 0x3a, 0xdf, 0x82, 0x86, 0x1a, 0x3a, 0x54, 0xc8,
	0x1f, 0x5b, 0xe2, 0x7a, 0xfc, 0x08, 0x9a, 0x2a, 0xd3, 0x13, 0x7d, 0x39, 0x17, 0xfb, 0x70, 0x4c,
	0x85, 0xdd, 0x4e, 0xa9, 0x77, 0x74, 0x55, 0xed, 0x77, 0x6f, 0x58, 0x19, 0x37, 0x19, 0xb2, 0x12,
	0x32, 0x79, 0x09, 0xce, 0xb5, 0xb9, 0xfb, 0x32, 0xb4, 0xd6, 0xf6, 0x8a, 0x78, 0xc4, 0x50, 0x88,
	0x6b, 0x5b, 0x1b, 0x97, 0xe0, 0x77, 0x50, 0xf6, 0x76, 0x36, 0xe9, 0x94, 0x7a, 0xa5, 0xab, 0x7a,
	0xbf, 0x93, 0x90, 0xf2, 0xf9, 0x0a, 0xb1, 0x57, 0x3b, 0x6d, 0x67, 0x13, 0x85, 0xb2, 0xf8, 0xbf,
	0x4b, 0xd0, 0xf0, 0x71, 0x91, 0xb8, 0x4b, 0xc7, 0xb4, 0x3d, 0xd3, 0xda, 0xe0, 0x00, 0xe2, 0xd2,
	0x7c, 0xcc, 0x72, 0xa8, 0x58, 0xb5, 0x7f, 0x91, 0x11, 0x13, 0x19, 0x41, 0xa9, 0x7b, 0xa9, 0x31,
	0xfe, 0x08, 0xac, 0xff, 0xa1, 0xc4, 0x3b, 0x2a, 0x91, 0xf4, 0x73, 0x6b, 0x2d, 0x12, 0x0a, 0xb5,
	0xe7, 0xe4, 0x30, 0xaa, 0x4f, 0xdd, 0x2e, 0xd6, 0xa6, 0x77, 0x78, 0x7d, 0x43, 0x38, 0x0b, 0xe0,
	0xb5, 0xf5, 0xe7, 0xc1, 0x22, 0x03, 0x40, 0x1f, 0x1e, 0xcd, 0xcd, 0xd5, 0xe7, 0x1a, 0x99, 0xd9,
	0xc6, 0xfc, 0xf0, 0x6a, 0x04, 0x68, 0x4d, 0x2d, 0x83, 0x08, 0x86, 0xb1, 0x97, 0x84, 0xcf, 0x2d,
	0xf0, 0x11, 0xc0, 0xfb, 0x36, 0xa4, 0x48, 0x64, 0x00, 0xe8, 0xc3, 0x7b, 0x37, 0xe4, 0x0d, 0x23,
	0xfb, 0x37, 0xa4, 0x48, 0x44, 0x80, 0x16, 0xfd, 0x4a, 0xfc, 0x3b, 0x77, 0x60, 0x4f, 0x25, 0x38,
	0x57, 0xc2, 0x0c, 0xd8, 0x57, 0xa6, 0xc8, 0xc9, 0x37, 0x50, 0xa1, 0xe7, 0x3b, 0x16, 0xf1, 0x02,
	0x4e, 0xe8, 0xfd, 0xd9, 0x9a, 0x06, 0x5d, 0x5c, 0x56, 0x2a, 0xfe, 0x78, 0x66, 0x1a, 0xfc, 0xf7,
	0x50, 0x8d, 0x36, 0xf3, 0x99, 0x97, 0xf0, 0x81, 0xe5, 0x4f, 0xc4, 0x3e, 0x55, 0xaa, 0x11, 0x36,
	0x33, 0x8d, 0xeb, 0x7f, 0x4b, 0x50, 0x4b, 0xd9, 0xc6, 0x73, 0x68, 0x69, 0x82, 0x7a, 0xa7, 0x0f,
	0xe5, 0xfb, 0x87, 0x89, 0xa4, 0x49, 0xa2, 0x2e, 0xdf, 0x35, 0x39, 0x06, 0xab, 0xb3, 0xc1, 0xfd,
	0x58, 0x0b, 0xe1, 0x12, 0x9e, 0x41, 0x83, 0xc2, 0x8a, 0x74, 0x2f, 0x3f, 0x06, 0xe0, 0x3b, 0x44,
	0xa8, 0x53, 0x70, 0x24, 0x8c, 0x27, 0x01, 0x76, 0xc4, 0x88, 0xb3, 0x07, 0x51, 0x08, 0x57, 0x97,
	0x19, 0x71, 0x2a, 0x6b, 0xfa, 0x48, 0x9e, 0x4d, 0xc5, 0xe6, 0x7b, 0x6c, 0x03, 0x52, 0xec, 0x56,
	0x1e, 0x24, 0xf0, 0x63, 0xec, 0x42, 0x9b, 0xe2, 0xc2, 0x44, 0x91, 0x04, 0xf1, 0x97, 0xd8, 0x48,
	0xb3, 0xc2, 0xe6, 0x54, 0x4d, 0xd0, 0x24, 0xba, 0x6a, 0xa8, 0x48, 0xfe, 0x36, 0xcd, 0x93, 0xeb,
	0xbf, 0x4a, 0x50, 0x4b, 0x75, 0x14, 0x5b, 0x50, 0x9b, 0xca, 0xa2, 0xa4, 0x0b, 0xa2, 0x18, 0x55,
	0x87, 0x50, 0xa7, 0x50, 0xec, 0x98, 0x96, 0x46, 0xb1, 0x54, 0x69, 0x11, 0x98, 0x28, 0xe3, 0x88,
	0xad, 0x8e, 0xed, 0x96, 0xf1, 0x0b, 0x38, 0x0b, 0x36, 0x09, 0xed, 0x4a, 0x3f, 0x8f, 0x55, 0x4d,
	0x6d, 0xbe, 0xef, 0xff, 0x53, 0x01, 0x1c, 0x45, 0x67, 0x1d, 0xc5, 0xb7, 0x83, 0x12, 0x9c, 0x44,
	0x03, 0x2c, 0x08, 0xe8, 0x28, 0xe0, 0xbb, 0x5f, 0xbe, 0x1e, 0xde, 0x2e, 0xcf, 0xe1, 0x4f, 0xc1,
	0x71, 0xb2, 0xdc, 0x46, 0xcc, 0x7c, 0x9f, 0xb3, 0xb1, 0xd8, 0xed, 0x65, 0xb0, 0x5c, 0xca, 0xf3,
	0x1c, 0x0a, 0x00, 0x71, 0x28, 0x15, 0xaa, 0x7c, 0x9d, 0xc1, 0xd2, 0xd7, 0x95, 0xe7, 0x70, 0x08,
	0xd5, 0x44, 0x38, 0x16, 0x6a, 0x7c, 0xca, 0xdd, 0x9e, 0x54, 0x6e, 0xf0, 0x1c, 0xca, 0x41, 0x41,
	0x2c, 0xa8, 0x53, 0xcd, 0xc9, 0x3c, 0x2d, 0xb9, 0xc2, 0x72, 0xf1, 0xce, 0x73, 0x78, 0x17, 0xb8,
	0x0a, 0x83, 0xe1, 0x4d, 0xb9, 0xac, 0xbb, 0x4c, 0x98, 0xf0, 0x1c, 0x3e, 0xc2, 0x29, 0x4b, 0x4c,
	0xfc, 0x36, 0x41, 0x8f, 0xae, 0xa1, 0x16, 0x3e, 0xfb, 0x3e, 0x2b, 0x7e, 0x7e, 0xba, 0x5f, 0x65,
	0xae, 0x7b, 0x2a, 0x72, 0x79, 0x0e, 0x25, 0x80, 0x38, 0x01, 0xb1, 0x5d, 0x20, 0x9c, 0x3d, 0x81,
	0x7c, 0x60, 0xd2, 0xaf, 0xa1, 0x9a, 0x48, 0xe3, 0x57, 0x75, 0x3e, 0xe5, 0xc2, 0x27, 0x7b, 0x0a,
	0xbf, 0x06, 0x42, 0x51, 0xd3, 0xfe, 0x47, 0xa5, 0x59, 0xed, 0x7c, 0x0f, 0x45, 0xf8, 0x20, 0x18,
	0x06, 0xcb, 0x59, 0xfc, 0x98, 0x3d, 0x44, 0x1f, 0x4d, 0x75, 0x2c, 0x97, 0xc9, 0x3c, 0x87, 0x13,
	0xaa, 0xe2, 0xef, 0x10, 0xa8, 0x74, 0x0a, 0x2c, 0x06, 0x4a, 0xbd, 0xd7, 0x66, 0x62, 0xb5, 0xc5,
	0x31, 0xfd, 0xc1, 0xfa, 0xe1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0x31, 0x42, 0x11, 0x0c,
	0x0a, 0x00, 0x00,
}
