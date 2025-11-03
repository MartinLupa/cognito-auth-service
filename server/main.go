package main

import (
	"fmt"
	"log"

	"github.com/MartinLupa/go-cognito-auth/aws"
	"github.com/MartinLupa/go-cognito-auth/config"
	"github.com/MartinLupa/go-cognito-auth/internal/handlers"
	"github.com/MartinLupa/go-cognito-auth/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load general config
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	router := gin.Default()
	cognitoClient := aws.NewCognitoClient(&cfg.AWS)
	authService := services.NewAuthService(cognitoClient)
	authHandlers := handlers.NewAuthHandlers(authService)

	fmt.Println("Cognito Service:", cognitoClient)

	// Define routes
	router.POST("/signup", authHandlers.Signup)
	router.POST("/confirm-email", authHandlers.ConfirmEmail)
	router.POST("/resend-confirmation-code", authHandlers.ResendConfirmationCode)
	router.POST("/signin", authHandlers.Signin)
	router.POST("/verify-session", authHandlers.VerifySession)
	router.POST("/signout", authHandlers.Signout)

	err = router.Run(cfg.Server.Port)
	if err != nil {
		panic("[Error] failed to start server due to: " + err.Error())
	}
}
