package main

import (
	"fmt"
	"time"

	api "gitlab.com/iman_api/api"
	_ "gitlab.com/iman_api/api/docs"
	"gitlab.com/iman_api/pkg/logger"
	"gitlab.com/iman_api/pkg/token"
)

func main() {
	logger.Init()
	log := logger.GetLogger()
	log.Info("logger initialized")

	apiServer := api.New(api.RoutetOptions{
		Log: log,
	})

	token, err := token.CreateToken(&token.TokenParams{
		Duration: time.Hour * 24,
	})
	if err != nil {
		log.Fatalf("failed to create token")
	}
	fmt.Println("Token: ", token)

	if err := apiServer.Run(fmt.Sprintf(":%s", "8080")); err != nil {
		log.Fatalf("failed to run server: %s", err)
	}
}
