package aws

import (
	"context"
	"fmt"

	"github.com/MartinLupa/go-cognito-auth/aws/utils"
	"github.com/MartinLupa/go-cognito-auth/config"
	"github.com/MartinLupa/go-cognito-auth/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/google/uuid"
)

type User struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CognitoClient interface {
	Signup(user *models.User) error
	ConfirmEmail(email, code string) error
	ResendConfirmationCode(email string) error
	Signin(email, password string) (string, error)
	VerifySession(accessToken string) error
	Signout(accessToken string) error
}

type cognitoService struct {
	AWSConfig     *config.AWSConfig
	cognitoClient *cognitoidentityprovider.Client
}

func NewCognitoClient(sdkConfig *config.AWSConfig) CognitoClient {
	cognitoClient := cognitoidentityprovider.NewFromConfig(sdkConfig.SDKConfig)

	return &cognitoService{
		cognitoClient: cognitoClient,
		AWSConfig:     sdkConfig,
	}
}

func (c *cognitoService) Signup(user *models.User) error {
	cognitoUser := &cognitoidentityprovider.SignUpInput{
		ClientId:   aws.String(c.AWSConfig.AppID),
		SecretHash: aws.String(utils.ComputeSecretHash(user.Email, c.AWSConfig.AppID, c.AWSConfig.ClientSecret)),
		Username:   aws.String(user.Email),
		Password:   aws.String(user.Password),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("custom:custom_id"),
				Value: aws.String(uuid.NewString()),
			},
			{
				Name:  aws.String("given_name"),
				Value: aws.String(user.GivenName),
			},
			{Name: aws.String("family_name"),
				Value: aws.String(user.FamilyName),
			},
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			},
		},
	}

	_, err := c.cognitoClient.SignUp(context.TODO(), cognitoUser)

	if err != nil {
		fmt.Println("Error during Cognito Signup: ", err)
		return err
	}

	return nil
}

func (c *cognitoService) ConfirmEmail(email, code string) error {
	confirmEmailInput := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(c.AWSConfig.AppID),
		SecretHash:       aws.String(utils.ComputeSecretHash(email, c.AWSConfig.AppID, c.AWSConfig.ClientSecret)),
		Username:         aws.String(email),
		ConfirmationCode: aws.String(code),
	}

	_, err := c.cognitoClient.ConfirmSignUp(context.TODO(), confirmEmailInput)
	if err != nil {
		return err
	}

	return nil
}

func (c *cognitoService) ResendConfirmationCode(email string) error {
	resendCodeInput := &cognitoidentityprovider.ResendConfirmationCodeInput{
		ClientId:   aws.String(c.AWSConfig.AppID),
		SecretHash: aws.String(utils.ComputeSecretHash(email, c.AWSConfig.AppID, c.AWSConfig.ClientSecret)),
		Username:   aws.String(email),
	}

	_, err := c.cognitoClient.ResendConfirmationCode(context.TODO(), resendCodeInput)

	if err != nil {
		return err
	}

	return nil
}

func (c *cognitoService) Signin(email, password string) (string, error) {
	initiateAuthInput := &cognitoidentityprovider.InitiateAuthInput{
		ClientId: aws.String(c.AWSConfig.AppID),
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		AuthParameters: map[string]string{
			"USERNAME":    email,
			"PASSWORD":    password,
			"SECRET_HASH": utils.ComputeSecretHash(email, c.AWSConfig.AppID, c.AWSConfig.ClientSecret),
		},
	}

	resp, err := c.cognitoClient.InitiateAuth(context.TODO(), initiateAuthInput)
	if err != nil {
		return "", err
	}

	fmt.Println("Cognito Signin Response: ", resp)

	return *resp.AuthenticationResult.AccessToken, nil
}

func (c *cognitoService) VerifySession(accessToken string) error {
	fmt.Println("Verifying session for access token: ", accessToken)
	getUserInput := &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(accessToken),
	}

	_, err := c.cognitoClient.GetUser(context.TODO(), getUserInput)
	if err != nil {
		return err
	}

	return nil
}

func (c *cognitoService) Signout(accessToken string) error {
	globalSignOutInput := &cognitoidentityprovider.GlobalSignOutInput{
		AccessToken: aws.String(accessToken),
	}

	resp, err := c.cognitoClient.GlobalSignOut(context.TODO(), globalSignOutInput)
	if err != nil {
		return err
	}

	fmt.Println("Cognito Signout Response: ", resp)

	return nil
}
