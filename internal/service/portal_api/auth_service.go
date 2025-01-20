package service

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cbglabs/nexusgate-pbmapi/internal/db"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types/request"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types/response"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"log"
)

// GetStates retrieves a list of states.
// It returns a slice of StatesListResponse and an error if any occurs.
func GetStates() ([]response.StatesListResponse, error) {
	return []response.StatesListResponse{
		{Id: 1, Name: "Alabama"},
		{Id: 2, Name: "Alaska"},
		{Id: 3, Name: "Arizona"},
	}, nil // Added nil to indicate no error
}

func Login(loginRequest request.LoginRequest) (string, error) {

	svc, databaseError := db.NewDynamoDBClient()
	if databaseError != nil {
		fmt.Println("Error connecting to the database.")
	}
	pk := "user#cbg"
	input := &dynamodb.QueryInput{
		TableName:              aws.String(os.Getenv("DYNAMODB_TABLE_NAME")),
		IndexName:              aws.String("email-index"), // Specifying GSI name if querying a secondary index
		KeyConditionExpression: aws.String("pk = :pk AND email = :email"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pk": {
				S: aws.String(pk),
			},
			":email": {
				S: aws.String(loginRequest.Email),
			},
		},
	}
	result, err := svc.QueryItem(input)
	// Initialize response
	response := response.UserLoginRecord{}
	// Unmarshal the result into the response
	if result.Count != nil && *result.Count > 0 {

		err = dynamodbattribute.UnmarshalMap(result.Items[0], &response)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal record: %v", err)
		}

		if response.Pk == "" {
			log.Fatalf("User not found: %v", err)
			return "User not found", nil
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(loginRequest.Password)); err != nil {
		fmt.Println("invalid password")
		return "Invalid Credentials.", nil
	}
	parts := strings.Split(response.Sk, "#")
	userId := ""
	// Check if there is a second part
	if len(parts) > 1 {
		userId = parts[1]
	} else {
		fmt.Println("No '#' found in the input string.")
	}
	claim := types.Claim{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // Token expiration time
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func Register(registerRequest request.RegisterRequest) string {
	svc, databaseError := db.NewDynamoDBClient()
	if databaseError != nil {
		fmt.Println("Error connecting to the database.")
	}

	var sk = fmt.Sprintf("user#%s", uuid.New().String())
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Got error while password hashing: %s", err)
	}

	// Initialize the payload map
	newUserPayload := map[string]string{
		"first_name": string(registerRequest.FirstName),
		"last_name":  string(registerRequest.LastName),
		"email":      string(registerRequest.Email),
		"password":   string(passwordHash),
		"pk":         "user#cbg",
		"sk":         sk,
		"created_at": time.Now().Format(time.RFC3339Nano),
	}

	av, err := dynamodbattribute.MarshalMap(newUserPayload)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	// Create item in table Movies
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
	return sk
}
