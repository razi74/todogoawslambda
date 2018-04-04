package main

import (
	"bytes"
	"dao"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rs/xid"
	"log"
	. "model"
	"os"
	"time"
)

var (
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Processing Lambda requestgettask %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	svc := dynamodb.New(sess)

	result, err := svc.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Tables:")
	fmt.Println("")

	for _, n := range result.TableNames {
		fmt.Println(*n)
	}

	addTask(request)
	fmt.Println("Tables: after ")
	return events.APIGatewayProxyResponse{
		Body:       "Hellogettask tables " + request.Body,
		StatusCode: 200,
	}, nil

}

func addTask(r events.APIGatewayProxyRequest) {

	var task Task
	//	defer r.Body.Close()
	fmt.Println("Tables:" + r.Body)
	b := bytes.NewBufferString(r.Body)
	decoder := json.NewDecoder(b)
	err := decoder.Decode(&task)
	if err != nil {
		//errorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}
	task.Createddate = time.Now().String()

	task.ID = xid.New().String()
	fmt.Println(task)
	err = dao.Insert(task)

	//responseWithJSON(w, task, http.StatusCreated)

}

func main() {
	lambda.Start(Handler)
}
