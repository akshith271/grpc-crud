package connection

import (
	"github.com/akshith/grpc/models"
	"github.com/jinzhu/gorm"
)

func ConnectDB() {
	db, err := gorm.Open("postgres", "user=postgres dbname=grpc password=root  sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	// f := flag.String("dbname", "postgres", "added by akshith")
	defer db.Close()
	// fmt.Printf(*f)
	db.DropTableIfExists(&models.Emp{})
	db.SingularTable(true) //won't change the name to plural
	db.CreateTable(&models.Emp{})

	//for dept schema
	db.DropTableIfExists(&models.Dept{})
	db.SingularTable(true)
	db.CreateTable(&models.Dept{})
	dep := models.Dept{
		DName: "Maths",
		Employees: []models.Emp{
			{EName: "Akshith", Manager: "Sai", DeptID: 1},
			{EName: "Surya", Manager: "NG", DeptID: 2},
		},
	}
	db.Save(&dep)

}
