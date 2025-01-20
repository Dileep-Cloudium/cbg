package db

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// DynamoDBClient defines the interface for DynamoDB operations
type DynamoDBClient interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	GetItems(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
	UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
	DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
	QueryItem(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
}

// dynamoDBClient implements DynamoDBClient interface
type dynamoDBClient struct {
	svc dynamodbiface.DynamoDBAPI
}

// Verify dynamoDBClient implements DynamoDBClient
var _ DynamoDBClient = (*dynamoDBClient)(nil)

// NewDynamoDBClient creates a new DynamoDB client
func NewDynamoDBClient() (DynamoDBClient, error) {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-west-2" // Fallback default
	}

	var sess *session.Session
	var err error

	if os.Getenv("AWS_ACCESS_KEY_ID") != "" && os.Getenv("AWS_SECRET_ACCESS_KEY") != "" {
		// Use static credentials for local development
		sess, err = session.NewSession(&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AWS_ACCESS_KEY_ID"),
				os.Getenv("AWS_SECRET_ACCESS_KEY"),
				os.Getenv("AWS_SESSION_TOKEN"), // Optional
			),
		})
	} else {
		// Use default credentials chain (works for Lambda, EC2, ECS, etc.)
		sess, err = session.NewSessionWithOptions(session.Options{
			Config: aws.Config{
				Region: aws.String(region),
			},
			SharedConfigState: session.SharedConfigEnable,
		})
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to create AWS session: %w", err)
	}

	return &dynamoDBClient{
		svc: dynamodb.New(sess),
	}, nil
}

// PutItem adds an item to a DynamoDB table
func (db *dynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	result, err := db.svc.PutItem(input)
	if err != nil {
		return nil, fmt.Errorf("Failed to put item: %w", err)
	}
	return result, nil
}

// GetItem retrieves an item to a DynamoDB table
func (db *dynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	result, err := db.svc.GetItem(input)
	if err != nil {
		return nil, fmt.Errorf("Failed to get item: %w", err)
	}
	return result, nil
}

// GetItems scans all items from the DynamoDB table with pagination
func (db *dynamoDBClient) GetItems(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	result := &dynamodb.ScanOutput{
		Items:        make([]map[string]*dynamodb.AttributeValue, 0),
		Count:        aws.Int64(0),
		ScannedCount: aws.Int64(0),
	}

	err := db.svc.ScanPages(input, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		result.Items = append(result.Items, page.Items...)

		// Copy other relevant fields from the last page
		result.ConsumedCapacity = page.ConsumedCapacity
		result.Count = aws.Int64(*result.Count + *page.Count)
		result.ScannedCount = aws.Int64(*result.ScannedCount + *page.ScannedCount)
		return true // Continue scanning
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to scan table: %w", err)
	}

	return result, nil
}

// UpdateItem updates an item in a DynamoDB table and returns the complete updated record
func (db *dynamoDBClient) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	// Ensure we get all attributes after update
	input.ReturnValues = aws.String("ALL_NEW")

	result, err := db.svc.UpdateItem(input)
	if err != nil {
		return nil, fmt.Errorf("Failed to update item: %w", err)
	}
	return result, nil
}

// DeleteItem deletes an item from a DynamoDB table
func (db *dynamoDBClient) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	result, err := db.svc.DeleteItem(input)
	if err != nil {
		return nil, fmt.Errorf("Failed to delete item: %w", err)
	}
	return result, nil
}

// QueryItem queries items from a DynamoDB table using the provided input
func (db *dynamoDBClient) QueryItem(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	result, err := db.svc.Query(input)
	if err != nil {
		return nil, fmt.Errorf("Failed to query items: %w", err)
	}
	return result, nil
}
