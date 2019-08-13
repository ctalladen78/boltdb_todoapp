package store

import "github.com/asdine/storm"

type Store struct {
	db *storm.DB
}

type Task struct {
	Key   int
	Value []byte
}

func (s Store) Init(path string) error {

	return nil
}

func (s Store) AllTasks() ([]*Task, error) {
	res := []*Task{}

	return res, nil
}

func (s Store) CreateTask(t string) (int, error) {

	return -1, nil
}

func (s Store) DeleteTask(key int) error {
	return nil
}
