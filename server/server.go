package main

import (
	"context"
	"log"
	"net"

	connection "github.com/akshith/grpc/connection"
	model "github.com/akshith/grpc/models"
	pb "github.com/akshith/grpc/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

// declaring the port number
const (
	port = "50051"
)

// connecting to the unimplemented methods via a struct
type empServer struct {
	pb.UnimplementedEmployeeDatabaseCrudServer
	db *gorm.DB //linking to db via this struct
}

// server side implementation of creating emp
func (s *empServer) CreateEmp(ctx context.Context, in *pb.NewEmp) (*pb.Emp, error) {
	log.Printf("createEmp method called from server side")
	newEmp := model.Emp{
		EName:   in.GetName(),
		Manager: in.GetManager(),
		DeptID:  int(in.GetDeptId()),
	}
	s.db.Save(&newEmp)
	return &pb.Emp{Name: in.GetName(), Manager: in.GetManager(), DeptId: in.GetDeptId()}, nil
}

// server side implementation of reading emp
func (s *empServer) ReadEmp(ctx context.Context, in *pb.VoidEmpRequest) (*pb.TotalEmp, error) {
	log.Printf("readEmp method called from server side")
	totalEmp := []model.Emp{}
	totalEmpData := []*pb.Emp{}
	s.db.Find(&totalEmp)
	for _, employee := range totalEmp {
		totalEmpData = append(totalEmpData, &pb.Emp{Name: employee.EName, Manager: employee.Manager, DeptId: int64(employee.DeptID)})
	}
	return &pb.TotalEmp{Emps: totalEmpData}, nil
}

// server side implementation of update emp
// updates the manager of an emp
func (s *empServer) UpdateEmp(ctx context.Context, in *pb.Emp) (*pb.Emp, error) {
	log.Printf("updateEmp method called from server side")
	s.db.Model(&model.Emp{}).Where("name=?", in.GetName()).Update("manager", in.GetManager())
	return &pb.Emp{Name: in.GetName(), Manager: in.GetManager()}, nil
}

// server side implementation of delete emp
func (s *empServer) DeleteEmp(ctx context.Context, in *pb.Emp) (*pb.VoidEmpResponse, error) {
	log.Printf("deleteEmp method called from server side")
	s.db.Where(&model.Emp{EName: in.GetName()}).Delete(&model.Emp{})
	return &pb.VoidEmpResponse{}, nil
}

func main() {

	//creating gorm instance
	connection.ConnectDB()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}

	//connecting gorm with grpc
	connection, err := gorm.Open("postgres", "user=postgres password=root dbname=grpc sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer connection.Close()

	//creating a new server
	s := grpc.NewServer()

	pb.RegisterEmployeeDatabaseCrudServer(s, &empServer{
		db: connection,
	})

	//if connection fails
	if err := s.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}

}
