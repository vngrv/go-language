package main

import (
	"fmt"
)

type employee struct {
	id     int
	name   string
	age    string
	salary int
}

// --- Memory Storage Start ---

// type storage interface {
// 	insert(e employee) error
// 	get(id int) (employee, error)
// 	delete(id int) error
// }

// type memoryStorage struct {
// 	data map[int]employee
// }

// func newMemoryStorage() *memoryStorage {
// 	return &memoryStorage{
// 		data: make(map[int]employee),
// 	}
// }

// func (s *memoryStorage) insert(e employee) error {
// 	s.data[e.id] = e
// 	return nil
// }

// func (s *memoryStorage) get(id int) (employee, error) {
// 	e, exist := s.data[id]

// 	if !exist {
// 		return employee{}, errors.New("employee with such id dosn't exist")
// 	}

// 	return e, nil
// }

// func (s *memoryStorage) delete(id int) (error) {
// 	delete(s.data, id)
// 	return nil
// }
// --- Memory Storage End ---

// --- Dump Storage Start ---

type dumpStorage struct{}

func newDumpStorage() *dumpStorage {
	return &dumpStorage{}
}

func (s *dumpStorage) insert(e employee) error {
	fmt.Printf("Insert user with id: %d successful\n", e.id)
	return nil
}

func (s *dumpStorage) get(id int) (employee, error) {
	e := employee{
		id: id,
	}

	return e, nil
}

func (e *dumpStorage) delete(id int) error {
	fmt.Printf("Delete user with id: %d successful\n", id)
	return nil
}

// --- Dump Storage End ---

func spawnEmployees(s storage) {
	for i := 1; i <= 10; i++ {
		s.insert(employee{id: i})
	}
}

func main() {
	ms := newMemoryStorage()
	ds := newDumpStorage()

	spawnEmployees(ms)
	fmt.Println(ms.get(3))

	spawnEmployees(ds)
}
