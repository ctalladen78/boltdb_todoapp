package main

import (
	"golang-projects/boltdb_todoapp/cmd"
	"golang-projects/boltdb_todoapp/stormdb"
	"log"
	"os"
	"path/filepath"
	"strconv"
	//homedir "github.com/mitchellh/go-homedir"
)

var db store.Store

const (
	GET = "get"
	PUT = "new"
	DEL = "delete"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	//home, _ := homedir.Dir()
	//dbPath := filepath.Join(home, "store.db")
	//log.Println(dbPath)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	storePath := filepath.Join(dir, "/store.db")
	//log.Println(dir)
	log.Println(storePath)

	err = db.Init(storePath)
	if err != nil {
		log.Fatal(err)
	}

	// TODO run echo webapp
	todo := os.Args[1]
	if todo == "" {
		log.Fatal("pls enter a command")
	}
	log.Print("command: ", string(todo))

	switch todo {
	case GET:
		getTodos()
	case PUT:
		putTodo(string(todo))
	case DEL:
		{
			// get os.Args[2]
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
			if err := db.DeleteTask(id); err != nil {
				log.Print(err)
			}
		}
	default:
		// get all todos
		db.AllTasks()
	}

}

func getTodos() {

	tasks, err := db.AllTasks()
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
	id, err := db.CreateTask(t)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Dope! added %s to your task list with id %d \n", t, id)

}
