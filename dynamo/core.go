package dynamo

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func InitializeDynamoDBClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return nil, err
	}
	return dynamodb.NewFromConfig(cfg), nil
}

func CreateTable(client *dynamodb.Client, name string) error {
	log.Printf("Table created with name %s\n", name)
	return nil
}

func GetTables(client *dynamodb.Client) ([]string, error) {
	resp, err := client.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
		return nil, err
	}

	return resp.TableNames, nil
}
