package main

import "github.com/ozonva/ova-plan-api/internal/server"

func main() {
	srv := server.New()

	err := srv.Run(":8080")
	if err != nil {
		return
	}

}
