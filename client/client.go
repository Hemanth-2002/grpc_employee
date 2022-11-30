package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc_employee/employeeproto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051" // port address for client
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	checkerror(err)
	defer conn.Close()

	client := pb.NewEmployeeDatabaseCrudClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//creating new employee
	new_employee, err := client.CreateEmployee(ctx, &pb.NewEmployee{EmpName: "sameer", DepartmentId: 1, ManagerName: "Vamshee"})
	checkerror(err)

	log.Printf("Employee Name: %v, Manager Name: %v, DepartmentId: %v", new_employee.GetEmpName(), new_employee.GetManagerName(), new_employee.GetDepartmentId())

	//get details of all employees
	AllEmployees, err := client.GetEmployees(ctx, &pb.EmptyEmployee{})
	if err != nil {
		log.Printf("error getting employees")
	}

	for _, emp := range AllEmployees.Employees {
		fmt.Println(emp.GetEmpName(), emp.GetManagerName(), emp.GetDepartmentId())
	}

	// updating manager of employee
	updated_manager, err := client.UpdateManager(ctx, &pb.Employee{EmpName: "KHK", ManagerName: "Ravi"})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(updated_manager)

	// deleting an employee
	delete, err := client.DeleteEmployee(ctx, &pb.Employee{EmpName: "sameer"})
	if err != nil {
		fmt.Println(delete)
		panic(err.Error())
	}

}

func checkerror(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
