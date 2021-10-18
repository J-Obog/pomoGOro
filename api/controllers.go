package api

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/pomoGOro/gormdb"
	"github.com/gorilla/mux"
)


func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	//add new tasks to db
	db := gormdb.Connect()
	db.Model(&Task{}).Create(&task)

	w.WriteHeader(200)
	w.Write([]byte("Task successfully created"))
}	

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task

	//get all tasks from db
	db := gormdb.Connect()
	db.Find(&tasks)

	res, err := json.Marshal(tasks)
	
	if err != nil {
		w.WriteHeader(500)
		return 
	}

	w.WriteHeader(200)
	w.Write(res)
}	

func RemoveTask(w http.ResponseWriter, r *http.Request) {

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	
	if err != nil {
		w.WriteHeader(500)
		return 
	}

	var task Task

	db := gormdb.Connect()
	//update task with associated id
	db.First(&task,id)
	db.Model(&task).Updates(body)
		
	w.WriteHeader(200)	
	w.Write([]byte("Task successfully updated"))
}