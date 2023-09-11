package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// loads ENVs
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	signedToken, err := SignToken(AdminClaims{
		ID:   1,
		Role: AdminRole{},
	})

	if err != nil {
		log.Fatal(err)
	}

	println(signedToken)

	parsedToken, err := VerifyToken(signedToken, Admin)

	if err != nil {
		log.Fatal(err)
	}

	println(parsedToken)
}
