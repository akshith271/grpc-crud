package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/akshith/grpc/proto"
	"google.golang.org/grpc"
)

// declaring port for the client to run on
const (
	address = "localhost:50051"
)

func main() {

	connection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Connection Failed", err.Error())
	}
	defer connection.Close()

	c := pb.NewEmployeeDatabaseCrudClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//create emp from client
	new_emp, err := c.CreateEmp(ctx, &pb.NewEmp{Name: "Madhav", Manager: "Teja", DeptId: 1})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Employee Name: %v, Manager Name: %v", new_emp.GetName(), new_emp.GetManager())

	//total emp from client
	total_emp, err := c.ReadEmp(ctx, &pb.VoidEmpRequest{})
	if err != nil {
		log.Printf("error getting all employees")
	}
	//printing emp details from the total_emp
	for _, employee := range total_emp.Emps {
		fmt.Println(employee.GetName(), employee.GetManager())
	}

	//updating the manager of an emp
	updated_emp, err := c.UpdateEmp(ctx, &pb.Emp{Name: "SuryaMadhav", Manager: "Akshith"})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(updated_emp)

	//deleting an instance of an emp
	deleted_emp, err := c.DeleteEmp(ctx, &pb.Emp{Name: "Hemanth"})
	if err != nil {
		fmt.Println(deleted_emp)
		panic(err.Error())
	}
}
