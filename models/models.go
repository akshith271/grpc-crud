package models

import "github.com/jinzhu/gorm"

// defining schema  for dept table
type Dept struct {
	gorm.Model
	DName     string
	Employees []Emp
}

// defining schema for emp table
type Emp struct {
	gorm.Model
	EName   string
	Manager string
	DeptID  int
}
