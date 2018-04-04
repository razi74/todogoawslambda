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

	log.Printf("Processing Lambda requestgettask %s\n", request.RequestContext)

	m := deleteTask(request)
	fmt.Println("Tables: after ")
	response, _ := json.Marshal(m)
	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
	}, nil

}

func deleteTask(r events.APIGatewayProxyRequest) map[string]string {

	var task Task
	//	defer r.Body.Close()
	task.ID = r.PathParameters["id"]
	fmt.Println(task)

	if err := dao.Delete(r.PathParameters["id"]); err != nil {
		//errorWithJSON(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err.Error())
		return map[string]string{"result": "error"}
	}
	return map[string]string{"result": "success"}
	//responseWithJSON(w, task, http.StatusCreated)

}

func main() {
	lambda.Start(Handler)
}
