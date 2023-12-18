package handler

import (
	

	
	"gofr.dev/pkg/gofr"
	
)
type Company struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Package  int    `json:"package"`
	Selected int    `json:"selected"`
	Ongoing  bool   `json:"ongoing"`
}
func Dashboard(ctx *gofr.Context) (interface{}, error) {

	return "Hello students! Welcome to the placement dashboard", nil
}
func PostName(ctx *gofr.Context) (interface{}, error) {
	name := ctx.PathParam("name")

	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO companies (name) VALUES (?)", name)

	return "successfully company name added, other values set 0 as default", err
}
func PostPackage(ctx *gofr.Context) (interface{}, error) {
	name := ctx.PathParam("name")
	packagee := ctx.PathParam("package")

	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO companies (name, package) VALUES (?, ?)", name, packagee)

	return "successfully added name and package", err
}
func SelectedStudent(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	selected := ctx.PathParam("selected")

	_, err := ctx.DB().ExecContext(ctx, "UPDATE companies SET selected = (?) WHERE id = (?);", selected, id)

	return "successfully updated selected count", err
}

func DeleteName(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM companies WHERE id = (?);", id)

	return "successfully deleted", err
}

func StopRecruitment(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	_, err := ctx.DB().ExecContext(ctx, "UPDATE companies SET ongoing = FALSE WHERE id = (?);", id)

	return "Drive's end stored", err
}

func StartRecruitment(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	_, err := ctx.DB().ExecContext(ctx, "UPDATE companies SET ongoing = TRUE WHERE id = (?);", id)

	return "drive continued", err
}

func CompanyDisplay(ctx *gofr.Context) (interface{}, error) {
	var companies []Company

	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM companies;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var company Company
		if err := rows.Scan(&company.ID, &company.Name, &company.Package, &company.Selected, &company.Ongoing); err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}