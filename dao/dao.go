package dao

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	. "model"
	"os"
)

const (
	TABLE_TASKS = "TASKS"
)

func FindAll() ([]Task, error) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	svc := dynamodb.New(sess)

	params := &dynamodb.ScanInput{
		TableName: aws.String(TABLE_TASKS),
	}

	result, err := svc.Scan(params)

	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}
	fmt.Println(result.Items)
	tasks := []Task{}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &tasks)
	if err != nil {
		//exitWithError(fmt.Errorf("failed to unmarshal Query result items, %v", err))
	}

	return tasks, err
}

func Insert(task Task) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	svc := dynamodb.New(sess)
	fmt.Println(task)
	av, err := dynamodbattribute.MarshalMap(task)

	fmt.Println(av)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TABLE_TASKS),
	}
	fmt.Println("Insert1:")
	_, err = svc.PutItem(input)
	fmt.Println("Insert2:")
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully added 'The Big New Movie' (2015) to Movies table")

	return err
}

func Delete(id string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	svc := dynamodb.New(sess)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(TABLE_TASKS),
	}
	_, err = svc.DeleteItem(input)

	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		//return
	}
	return err
}

//
//func (m *TasksDAO) Update(task Task) error {
//	err := db.C(COLLECTION_TASKS).UpdateId(task.ID, &task)
//	return err
//}
