package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func StartDB() {
	db, err := gorm.Open("postgres", "user=postgres password=MohanNeelima@01 dbname=employee sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// db.DropTableIfExists(&Employee{})
	// db.CreateTable(&Employee{})
	// db.DropTableIfExists(&Department{})
	// db.CreateTable(&Department{})

	// db.Save(&Department{
	// 	DeptName: "Front-end",
	// 	Employee: []Employee{
	// 		{EmpName: "Dp", ManagerName: "Vamshee"},
	// 		{EmpName: "Sandeep", ManagerName: "Vamshee"},
	// 	},
	// })

	// db.Save(&Department{
	// 	DeptName: "Back-end",
	// 	Employee: []Employee{
	// 		{EmpName: "KHK", ManagerName: "Mani"},
	// 		{EmpName: "AB", ManagerName: "Mani"},
	// 		{EmpName: "GST", ManagerName: "Mani"},
	// 		{EmpName: "PST", ManagerName: "Mani"},
	// 	},
	// })
	// users := []Department{}

	// db.Preload("Employee").Find(&users)
	// for _, user := range users {
	// 	fmt.Printf("\n%v\n", user)
	// }

	// output := []Emp_dept{}
	// db.Debug().Model(&users).Select([]string{"emp_name", "manager_name", "dept_name"}).Scan(&output)
	// db.Debug().Model(&Department{}).Joins("inner join employees on employees.department_id = departments.id").
	// 	Select("employees.emp_name, employees.manager_name, departments.dept_name").Scan(&output)
	// for _, u := range output {
	// 	fmt.Printf("\n%v\n", u)
	// }

	// fmt.Println("Done")
}

type Department struct {
	gorm.Model
	DeptName string
	Employee []Employee
}

type Employee struct {
	gorm.Model
	EmpName      string
	DepartmentID uint
	ManagerName  string
}

type Emp_dept struct {
	EmpName     string
	ManagerName string
	DeptName    string
}
