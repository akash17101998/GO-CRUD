package main

import "gorm.io/gorm"

type Employee struct{
	gorm.Model
	EmpId int `json:"<-"`
	EmpName string     `json:"empname"`
	EmpSalary float32   `json:"salary"`
	EmpEmail string         `json:"email-id"`
	// EmpContact []string      `json:"contactNo."`
}