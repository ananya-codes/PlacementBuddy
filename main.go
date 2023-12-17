package main

import "gofr.dev/pkg/gofr"

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
    
    app := gofr.New()

    
    app.GET("/dashboard", func(ctx *gofr.Context) (interface{}, error) {

			return "Hello students! Welcome to the placement dashboard", nil
		})

		

    // Starts the server, it will listen on the default port 8000.
    // it can be over-ridden through configs
    app.Start()
}