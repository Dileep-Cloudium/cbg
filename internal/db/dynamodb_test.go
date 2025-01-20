package db

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDynamoDBClient is a mock implementation of dynamodbiface.DynamoDBAPI
type MockDynamoDBClient struct {
	mock.Mock
	dynamodbiface.DynamoDBAPI
}

// Test implementations for each DynamoDB operation
func (m *MockDynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.PutItemOutput), args.Error(1)
}

func (m *MockDynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.GetItemOutput), args.Error(1)
}

func (m *MockDynamoDBClient) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.DeleteItemOutput), args.Error(1)
}

func (m *MockDynamoDBClient) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.QueryOutput), args.Error(1)
}

// Example test cases
func TestPutItem(t *testing.T) {
	mockSvc := new(MockDynamoDBClient)
	client := &dynamoDBClient{svc: mockSvc}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("test-table"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String("123")},
		},
	}

	expectedOutput := &dynamodb.PutItemOutput{}
	mockSvc.On("PutItem", input).Return(expectedOutput, nil)

	result, err := client.PutItem(input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, result)
	mockSvc.AssertExpectations(t)
}

func TestGetItem(t *testing.T) {
	mockSvc := new(MockDynamoDBClient)
	client := &dynamoDBClient{svc: mockSvc}

	input := &dynamodb.GetItemInput{
		TableName: aws.String("test-table"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String("123")},
		},
	}

	expectedOutput := &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"id":   {S: aws.String("123")},
			"name": {S: aws.String("test")},
		},
	}

	mockSvc.On("GetItem", input).Return(expectedOutput, nil)

	result, err := client.GetItem(input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, result)
	mockSvc.AssertExpectations(t)
}

// Add error test case example
func TestGetItemError(t *testing.T) {
	mockSvc := new(MockDynamoDBClient)
	client := &dynamoDBClient{svc: mockSvc}

	input := &dynamodb.GetItemInput{
		TableName: aws.String("test-table"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String("123")},
		},
	}

	mockSvc.On("GetItem", input).Return(nil, assert.AnError)

	result, err := client.GetItem(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Failed to get item")
	mockSvc.AssertExpectations(t)
}
