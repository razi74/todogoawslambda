package model

import ()

type Task struct {
	ID          string `json:"id"`
	Taskname    string `json:"taskname"`
	Createddate string `json:"createddate"`
}
