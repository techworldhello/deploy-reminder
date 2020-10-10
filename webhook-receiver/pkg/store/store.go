package store

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/techworldhello/deploy-reminder/webhook-reciever/pkg/webhook_parser"
)

type Store struct {
	table *dynamodb.DynamoDB
}

func New() *Store{
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := dynamodb.New(sess)
	return &Store{client}
}

func (s Store) SaveData(data webhook_parser.JobFinished) error {
	data.Id = "abcd"
	d, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		return err
	}
	log.Printf("What is the d 8==D: %+v", d)
	input := &dynamodb.PutItemInput{
		Item:      d,
		TableName: aws.String("webhook-receiver-table"),
	}
	_, err = s.table.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}
