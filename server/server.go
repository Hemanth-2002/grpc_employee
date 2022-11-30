package main

import (
	"context"
	"log"
	"net"

	pb "grpc_employee/employeeproto"
	model "grpc_employee/model"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	port = ":50051" // choosing port number
)

type employeeServer struct {
	pb.UnimplementedEmployeeDatabaseCrudServer
	Db *gorm.DB
}

// function to add new employee on server
func (s *employeeServer) CreateEmployee(ctx context.Context, in *pb.NewEmployee) (*pb.Employee, error) {
	log.Printf("creating new employee called")
	newEmp := model.Employee{
		EmpName:      in.GetEmpName(),
		DepartmentID: uint(in.GetDepartmentId()),
		ManagerName:  in.GetManagerName(),
	}
	s.Db.Save(&newEmp)
	return &pb.Employee{EmpName: in.GetEmpName(), DepartmentId: in.GetDepartmentId(), ManagerName: in.GetManagerName(), Id: uint64(newEmp.ID)}, nil
}

// function to read employee detail on server
func (s *employeeServer) GetEmployees(ctx context.Context, in *pb.EmptyEmployee) (*pb.Employees, error) {
	log.Printf("Getting employees called")
	Employees := []model.Employee{}
	FinalEmployees := []*pb.Employee{}
	s.Db.Find(&Employees)
	for _, emp := range Employees {
		FinalEmployees = append(FinalEmployees, &pb.Employee{EmpName: emp.EmpName, DepartmentId: uint64(emp.DepartmentID), ManagerName: emp.ManagerName, Id: uint64(emp.ID)})
	}
	return &pb.Employees{Employees: FinalEmployees}, nil
}

// function to update manager on server
func (s *employeeServer) UpdateManager(ctx context.Context, in *pb.Employee) (*pb.Employee, error) {
	log.Printf("update manager called")
	s.Db.Model(&model.Employee{}).Where("name=?", in.EmpName).Update("manager", in.ManagerName)
	return &pb.Employee{EmpName: in.GetEmpName(), ManagerName: in.GetManagerName()}, nil
}

// function to delete employee on server
func (s *employeeServer) DeleteEmployee(ctx context.Context, in *pb.Employee) (*pb.EmptyEmployee, error) {
	log.Printf("delete employee called")
	s.Db.Where(&model.Employee{EmpName: in.GetEmpName()}).Delete(&model.Employee{})
	return &pb.EmptyEmployee{}, nil
}

func main() {

	model.StartDB()
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}

	//db connection
	db, err := gorm.Open("postgres", "user=postgres password=MohanNeelima@01 dbname=employee sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//create new server
	new_server := grpc.NewServer()
	pb.RegisterEmployeeDatabaseCrudServer(new_server, &employeeServer{
		Db: db,
	})

	log.Printf("Using port no %v", listen.Addr())

	if err := new_server.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}
