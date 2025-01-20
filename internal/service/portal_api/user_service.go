package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/cbglabs/nexusgate-pbmapi/internal/db"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types/request"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types/response"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

const (
	userPrefix = "user#"
	pkPrefix   = userPrefix + "cbg"
)

// UserService defines the interface for user operations
type UserService interface {
	CreateUser(newUserRequest request.AddUserRequest) (response.CommonResponse, error)
	ReadUser(uuid string) (response.CommonResponse, error)
	UpdateUser(updateUserRequest request.UpdateUserRequest, uuid string) (response.CommonResponse, error)
	DeleteUser(uuid string) (response.CommonResponse, error)
	ListUsers() ([]response.CommonResponse, error)
}

// userService implements the UserService interface
type userService struct {
	dbClient db.DynamoDBClient
}

// Verify userService implements UserService
var _ UserService = (*userService)(nil)

// NewUserService creates a new UserService instance
func NewUserService(dbClient db.DynamoDBClient) UserService {
	return &userService{
		dbClient: dbClient,
	}
}

func (s *userService) CreateUser(newUserRequest request.AddUserRequest) (response.CommonResponse, error) {
	newUserRequestJSON, err := json.Marshal(newUserRequest)
	if err != nil {
		return response.CommonResponse{}, fmt.Errorf("Error serializing user request: %w", err)
	}

	userSK := userPrefix + uuid.New().String()
	newUserPayload := map[string]string{
		"item_body":  string(newUserRequestJSON),
		"pk":         pkPrefix,
		"sk":         userSK,
		"created_at": time.Now().Format(time.RFC3339Nano),
	}

	av, err := dynamodbattribute.MarshalMap(newUserPayload)
	if err != nil {
		return response.CommonResponse{}, fmt.Errorf("Error marshalling payload: %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE_NAME")),
	}

	if _, err = s.dbClient.PutItem(input); err != nil {
		return response.CommonResponse{}, fmt.Errorf("Error putting item: %w", err)
	}

	// Get and return the newly created user
	user, err := s.ReadUser(userSK[len(userPrefix):])
	if err != nil {
		return response.CommonResponse{}, fmt.Errorf("Error retrieving created user: %w", err)
	}

	return user, nil
}

func (s *userService) ReadUser(uuid string) (response.CommonResponse, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE_NAME")),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(pkPrefix)},
			"sk": {S: aws.String(userPrefix + uuid)},
		},
	}

	result, err := s.dbClient.GetItem(input)
	if err != nil {
		return response.CommonResponse{}, fmt.Errorf("Error getting item: %w", err)
	}

	if result.Item == nil {
		return response.CommonResponse{}, fmt.Errorf("User not found with UUID: %s", uuid)
	}

	var user response.CommonResponse
	if err := dynamodbattribute.UnmarshalMap(result.Item, &user); err != nil {
		return response.CommonResponse{}, fmt.Errorf("Failed to unmarshal record: %w", err)
	}

	return user, nil
}

func (s *userService) UpdateUser(updateUserRequest request.UpdateUserRequest, uuid string) (response.CommonResponse, error) {
	updateUserData, err := s.ReadUser(uuid)
	if err != nil {
		return response.CommonResponse{}, fmt.Errorf("Error getting user: %w", err)
	}

	var itemBody request.AddUserRequest
	if err := json.Unmarshal([]byte(updateUserData.ItemBody), &itemBody); err != nil {
		return response.CommonResponse{}, fmt.Errorf("Failed to unmarshal item body: %w", err)
	}

	itemBody.FirstName = updateUserRequest.FirstName
	itemBody.LastName = updateUserRequest.LastName

	updatedItemBody, err := json.Marshal(itemBody)
	if err != nil {
		return response.CommonResponse{}, fmt.Errorf("Failed to marshal updated item body: %w", err)
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":item_body": {S: aws.String(string(updatedItemBody))},
		},
		UpdateExpression: aws.String("SET item_body = :item_body"),
		TableName:        aws.String(os.Getenv("DYNAMODB_TABLE_NAME")),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(pkPrefix)},
			"sk": {S: aws.String(userPrefix + uuid)},
		},
		ReturnValues: aws.String("UPDATED_NEW"),
	}

	user, err := s.dbClient.UpdateItem(input)
	if err != nil {
		return response.CommonResponse{}, fmt.Errorf("Error updating item: %w", err)
	}

	var updatedUser response.CommonResponse
	err = dynamodbattribute.UnmarshalMap(user.Attributes, &updatedUser)
	if err != nil {
		return response.CommonResponse{}, fmt.Errorf("Failed to unmarshal updated user: %w", err)
	}

	return updatedUser, nil
}

func (s *userService) DeleteUser(uuid string) (response.CommonResponse, error) {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(pkPrefix)},
			"sk": {S: aws.String(userPrefix + uuid)},
		},
		TableName:    aws.String(os.Getenv("DYNAMODB_TABLE_NAME")),
		ReturnValues: aws.String("ALL_OLD"),
	}

	result, err := s.dbClient.DeleteItem(input)
	if err != nil {
		return response.CommonResponse{}, fmt.Errorf("Error deleting item: %w", err)
	}

	if result.Attributes == nil {
		return response.CommonResponse{}, fmt.Errorf("User not found with UUID: %s", uuid)
	}

	var user response.CommonResponse
	if err := dynamodbattribute.UnmarshalMap(result.Attributes, &user); err != nil {
		return response.CommonResponse{}, fmt.Errorf("Failed to unmarshal deleted user: %w", err)
	}

	return user, nil
}

func (s *userService) ListUsers() ([]response.CommonResponse, error) {
	result, err := s.dbClient.GetItems(&dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE_NAME")),
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to scan table: %w", err)
	}

	var users []response.CommonResponse
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal items: %w", err)
	}

	return users, nil
}
