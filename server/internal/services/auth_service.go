package services

import (
	"github.com/MartinLupa/go-cognito-auth/aws"
	"github.com/MartinLupa/go-cognito-auth/internal/models"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type AuthService interface {
	Signup(user *models.User) error
	ConfirmEmail(email, code string) error
	ResendConfirmationCode(email string) error
	Signin(email, password string) (string, error)
	Signout(accessToken string) error
	ListUsers() ([]types.UserType, error)
}

type authService struct {
	cognitoClient aws.CognitoClient
}

func NewAuthService(cognitoClient aws.CognitoClient) AuthService {
	return &authService{
		cognitoClient: cognitoClient,
	}
}

func (a *authService) Signup(user *models.User) error {
	err := a.cognitoClient.Signup(&models.User{
		GivenName:  user.GivenName,
		FamilyName: user.FamilyName,
		Email:      user.Email,
		Password:   user.Password,
	})

	if err != nil {
		return err
	}

	return nil
}

func (a *authService) ConfirmEmail(email, code string) error {
	err := a.cognitoClient.ConfirmEmail(email, code)

	if err != nil {
		return err
	}

	return nil
}

func (a *authService) ResendConfirmationCode(email string) error {
	err := a.cognitoClient.ResendConfirmationCode(email)

	if err != nil {
		return err
	}

	return nil
}

func (a *authService) Signin(email, password string) (string, error) {
	token, err := a.cognitoClient.Signin(email, password)

	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *authService) Signout(accessToken string) error {
	err := a.cognitoClient.Signout(accessToken)

	if err != nil {
		return err
	}

	return nil
}

func (a *authService) ListUsers() ([]types.UserType, error) {
	return a.cognitoClient.ListUsers()
}
