package main

import (
	"gofr.dev/pkg/gofr"
	"vs-go/PlacementBuddy/handler"
)

type Company struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Package  int    `json:"package"`
	Selected int    `json:"selected"`
	Ongoing  bool   `json:"ongoing"`
}

func main() {

	app := gofr.New()

	app.GET("/dashboard", handler.Dashboard) //WELCOME DASHBOARD
	app.POST("/company/{name}", handler.PostName) // ADD NEW COMPANY NAME (default package set 0)
	app.POST("/company/{name}/{package}", handler.PostPackage) //ADD NEW COMPANY NAME AND ITS PACKAGE
	app.POST("/selected/{id}/{selected}", handler.SelectedStudent) //ADD NO OF STUDENTS SELECTED BY THE COMPANY's ID
	app.DELETE("/company/{id}", handler.DeleteName) //DELETE COMPANY USING ID
	app.POST("/stop/{id}", handler.StopRecruitment)//END COMPANY'S HIRING PROCESS
	app.POST("/start/{id}", handler.StartRecruitment) //RESUME COMPANY'S HIRING PROCESS
	app.GET("/company", handler.CompanyDisplay)	//DISPLAY COMPANY'S STATUS
	

	app.Start()
}
