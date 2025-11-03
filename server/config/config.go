package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	AWS    AWSConfig
}

type ServerConfig struct {
	Port string
}

type AWSConfig struct {
	AccessKey    string
	SecretKey    string
	Region       string
	AppID        string
	UserPoolID   string
	ClientSecret string
	SDKConfig    aws.Config
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// AWS SDK Configuration
	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return nil, nil
	}

	cfg := &Config{
		Server: ServerConfig{
			Port: os.Getenv("PORT"),
		},
		AWS: AWSConfig{
			AccessKey:    os.Getenv("AWS_ACCESS_KEY"),
			SecretKey:    os.Getenv("AWS_SECRET_KEY"),
			Region:       os.Getenv("AWS_REGION"),
			AppID:        os.Getenv("APP_CLIENT_ID"),
			UserPoolID:   os.Getenv("USER_POOL_ID"),
			ClientSecret: os.Getenv("SECRET_HASH"),
			SDKConfig:    sdkConfig,
		},
	}
	return cfg, nil
}
