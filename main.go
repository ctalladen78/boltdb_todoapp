package main

import (
	"golang-projects/boltdb_todoapp/boltdb"
	"golang-projects/boltdb_todoapp/cmd"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var db store.Store
var storePath string

const (
	GET  = "get"
	PUT  = "new"
	DEL  = "delete"
	FIND = "find"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	storePath = filepath.Join(dir, "/bolt.db")
	log.Println(storePath)

	err = db.Init(storePath)
	if err != nil {
		log.Fatal(err)
	}

	// TODO run echo webapp
	input := os.Args[1]
	if input == "" {
		log.Fatal("pls enter a command")
	}
	log.Print("command: ", string(input))

	switch input {
	case GET:
		getTodos()
	case PUT:
		todo := os.Args[2]
		if input == "" {
			log.Fatal("pls enter a command")
		}
		putTodo(string(todo))
	case DEL:
		{
			// get os.Args[2]
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
			if err := db.DeleteTask(storePath, id); err != nil {
				log.Print(err)
			}
		}
	case FIND:
		{
			// query by key
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
			db.Select(id)

		}
	default:
		// get all todos
		db.AllTasks(storePath)
	}

}

func getTodos() {

	tasks, err := db.AllTasks(storePath)
	// error handling
	if err != nil {
		return
	}
	for _, task := range tasks {
		log.Printf("%d : %s \n", task.Key, task.Value)
	}
}

func putTodo(t string) {

	// stdin argument (./boltdb_todoapp "first todo")
	//task := strings.Join(t, " ")
	id, err := db.CreateTask(storePath, t)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Dope! added %s to your task list with id %d \n", t, id)

}
