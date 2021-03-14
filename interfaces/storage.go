package main

import (
	"errors"
	"fmt"
)

type employee struct {
	id     int
	name   string
	age    string
	salary int
}

// --- Memory Storage Start ---

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]employee
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (s *memoryStorage) insert(e employee) error {
	s.data[e.id] = e
	return nil
}

func (s *memoryStorage) get(id int) (employee, error) {
	e, exist := s.data[id]

	if !exist {
		return employee{}, errors.New("employee with such id dosn't exist")
	}

	return e, nil
}

func (s *memoryStorage) delete(id int) (error) {
	delete(s.data, id)
	return nil
}
// --- Memory Storage End ---

// --- Dump Storage Start ---

type dumpStorage struct {}

func newDumpStorage() *dumpStorage {
	return &dumpStorage
}

func (s *dumpStorage) insert(e employee) error {
	fmt.Printf("Insert user with id: %d successful", e.id)
	return nil
}

// --- Dump Storage End ---

func main() {
	var s storage

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	s = newMemoryStorage()

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)
}
