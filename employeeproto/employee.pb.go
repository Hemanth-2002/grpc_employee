// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: employeeproto/employee.proto

package grpc_employee

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EmpName      string `protobuf:"bytes,2,opt,name=emp_name,json=empName,proto3" json:"emp_name,omitempty"`
	DepartmentId uint64 `protobuf:"varint,3,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	ManagerName  string `protobuf:"bytes,4,opt,name=manager_name,json=managerName,proto3" json:"manager_name,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employeeproto_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_employeeproto_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_employeeproto_employee_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Employee) GetEmpName() string {
	if x != nil {
		return x.EmpName
	}
	return ""
}

func (x *Employee) GetDepartmentId() uint64 {
	if x != nil {
		return x.DepartmentId
	}
	return 0
}

func (x *Employee) GetManagerName() string {
	if x != nil {
		return x.ManagerName
	}
	return ""
}

type NewEmployee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmpName      string `protobuf:"bytes,1,opt,name=emp_name,json=empName,proto3" json:"emp_name,omitempty"`
	DepartmentId uint64 `protobuf:"varint,2,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	ManagerName  string `protobuf:"bytes,3,opt,name=manager_name,json=managerName,proto3" json:"manager_name,omitempty"`
}

func (x *NewEmployee) Reset() {
	*x = NewEmployee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employeeproto_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewEmployee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewEmployee) ProtoMessage() {}

func (x *NewEmployee) ProtoReflect() protoreflect.Message {
	mi := &file_employeeproto_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewEmployee.ProtoReflect.Descriptor instead.
func (*NewEmployee) Descriptor() ([]byte, []int) {
	return file_employeeproto_employee_proto_rawDescGZIP(), []int{1}
}

func (x *NewEmployee) GetEmpName() string {
	if x != nil {
		return x.EmpName
	}
	return ""
}

func (x *NewEmployee) GetDepartmentId() uint64 {
	if x != nil {
		return x.DepartmentId
	}
	return 0
}

func (x *NewEmployee) GetManagerName() string {
	if x != nil {
		return x.ManagerName
	}
	return ""
}

type EmptyEmployee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyEmployee) Reset() {
	*x = EmptyEmployee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employeeproto_employee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyEmployee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyEmployee) ProtoMessage() {}

func (x *EmptyEmployee) ProtoReflect() protoreflect.Message {
	mi := &file_employeeproto_employee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyEmployee.ProtoReflect.Descriptor instead.
func (*EmptyEmployee) Descriptor() ([]byte, []int) {
	return file_employeeproto_employee_proto_rawDescGZIP(), []int{2}
}

type Employees struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employees []*Employee `protobuf:"bytes,1,rep,name=employees,proto3" json:"employees,omitempty"`
}

func (x *Employees) Reset() {
	*x = Employees{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employeeproto_employee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employees) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employees) ProtoMessage() {}

func (x *Employees) ProtoReflect() protoreflect.Message {
	mi := &file_employeeproto_employee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employees.ProtoReflect.Descriptor instead.
func (*Employees) Descriptor() ([]byte, []int) {
	return file_employeeproto_employee_proto_rawDescGZIP(), []int{3}
}

func (x *Employees) GetEmployees() []*Employee {
	if x != nil {
		return x.Employees
	}
	return nil
}

var File_employeeproto_employee_proto protoreflect.FileDescriptor

var file_employeeproto_employee_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x7d, 0x0a,
	0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6d, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6d, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x70, 0x0a, 0x0b,
	0x4e, 0x65, 0x77, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6d, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6d, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x0f,
	0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22,
	0x42, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x09,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x73, 0x32, 0xb9, 0x02, 0x0a, 0x14, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x72, 0x75, 0x64, 0x12, 0x47, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x1a,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x4e,
	0x65, 0x77, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x1a,
	0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x00, 0x42,
	0x1c, 0x5a, 0x1a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_employeeproto_employee_proto_rawDescOnce sync.Once
	file_employeeproto_employee_proto_rawDescData = file_employeeproto_employee_proto_rawDesc
)

func file_employeeproto_employee_proto_rawDescGZIP() []byte {
	file_employeeproto_employee_proto_rawDescOnce.Do(func() {
		file_employeeproto_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_employeeproto_employee_proto_rawDescData)
	})
	return file_employeeproto_employee_proto_rawDescData
}

var file_employeeproto_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_employeeproto_employee_proto_goTypes = []interface{}{
	(*Employee)(nil),      // 0: grpc_employee.Employee
	(*NewEmployee)(nil),   // 1: grpc_employee.NewEmployee
	(*EmptyEmployee)(nil), // 2: grpc_employee.EmptyEmployee
	(*Employees)(nil),     // 3: grpc_employee.Employees
}
var file_employeeproto_employee_proto_depIdxs = []int32{
	0, // 0: grpc_employee.Employees.employees:type_name -> grpc_employee.Employee
	1, // 1: grpc_employee.EmployeeDatabaseCrud.CreateEmployee:input_type -> grpc_employee.NewEmployee
	2, // 2: grpc_employee.EmployeeDatabaseCrud.GetEmployees:input_type -> grpc_employee.EmptyEmployee
	0, // 3: grpc_employee.EmployeeDatabaseCrud.UpdateManager:input_type -> grpc_employee.Employee
	0, // 4: grpc_employee.EmployeeDatabaseCrud.DeleteEmployee:input_type -> grpc_employee.Employee
	0, // 5: grpc_employee.EmployeeDatabaseCrud.CreateEmployee:output_type -> grpc_employee.Employee
	3, // 6: grpc_employee.EmployeeDatabaseCrud.GetEmployees:output_type -> grpc_employee.Employees
	0, // 7: grpc_employee.EmployeeDatabaseCrud.UpdateManager:output_type -> grpc_employee.Employee
	2, // 8: grpc_employee.EmployeeDatabaseCrud.DeleteEmployee:output_type -> grpc_employee.EmptyEmployee
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_employeeproto_employee_proto_init() }
func file_employeeproto_employee_proto_init() {
	if File_employeeproto_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_employeeproto_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_employeeproto_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewEmployee); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_employeeproto_employee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyEmployee); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_employeeproto_employee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employees); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_employeeproto_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_employeeproto_employee_proto_goTypes,
		DependencyIndexes: file_employeeproto_employee_proto_depIdxs,
		MessageInfos:      file_employeeproto_employee_proto_msgTypes,
	}.Build()
	File_employeeproto_employee_proto = out.File
	file_employeeproto_employee_proto_rawDesc = nil
	file_employeeproto_employee_proto_goTypes = nil
	file_employeeproto_employee_proto_depIdxs = nil
}
