package main

import(
  "log"
  "os"
  "path/filepath"
  "golang-projects/boltdb_todoapp/cmd"
  //homedir "github.com/mitchellh/go-homedir"
) 

var store Store

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


  err = store.Init(storePath)
  if err != nil {
    log.Fatal(err)
    }

  // TODO run echo webapp
  todo := os.Args[1]
  log.Print(string(todo))
  putTodo(string(todo))
}

func getTodos() {
  
  tasks, err := store.AllTasks()
  // error handling
  if err != nil { return }
  for i, task := range tasks {
    log.Printf("%d : %s \n", i+1, task.Value)
  }
}

func putTodo(t string) {

  // stdin argument (./boltdb_todoapp "first todo")
  //task := strings.Join(t, " ")
  id, err := store.CreateTask(t)
  if err != nil {log.Fatal(err)}
  log.Printf("Dope! added %s to your task list with id %d \n", t, id)
  
}

