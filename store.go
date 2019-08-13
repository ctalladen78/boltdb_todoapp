package main

import "github.com/boltdb/bolt"

// import asdine/armor

type Store struct {
	db *bolt.DB
}

type Task struct {
	Id    []byte
	Value []byte
}

func (s Store) AllTasks() ([]*Task, error) {
	res := []*Task{}

	return res, nil
}

func (s Store) CreateTask(t string) (int, error) {

	return -1, nil
}
