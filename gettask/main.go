package main

import (
	"dao"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"log"
	. "model"
)

var (
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Processing Lambda requestgettask %s\n", request.RequestContext.RequestID)

	tasks := findTasks(request)
	fmt.Println("Tables: after ")
	response, _ := json.Marshal(tasks)
	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
	}, nil

}

func findTasks(r events.APIGatewayProxyRequest) []Task {

	var tasks []Task
	tasks, err := dao.FindAll()
	if err != nil {
		//errorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed get all tasks: ", err)
	}
	fmt.Println(tasks)
	//responseWithJSON(w, task, http.StatusCreated)
	return tasks

}

func main() {
	lambda.Start(Handler)
}
