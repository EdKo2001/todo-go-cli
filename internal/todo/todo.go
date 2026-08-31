package todo

import (
	"fmt"
	"time"
)

type Todo struct {
	ID          int
	Title       string
	Completed   bool
	Due         string
	CreateAt    time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string, due string) Todo {
	todo := Todo{
		ID:       todos.nextID(),
		Title:    title,
		Due:      due,
		CreateAt: time.Now(),
	}

	*todos = append(*todos, todo)
	return todo
}

func (todos Todos) nextID() int {
	highestID := 0
	for _, todo := range todos {
		if todo.ID > highestID {
			highestID = todo.ID
		}
	}

	return highestID + 1
}

func (todos Todos) findIndex(id int) (int, error) {
	for index, todo := range todos {
		if todo.ID == id {
			return index, nil
		}
	}

	return -1, fmt.Errorf("todo %d not found", id)
}

func (todos *Todos) Complete(id int) (Todo, error) {
	index, err := todos.findIndex(id)
	if err != nil {
		return Todo{}, err
	}

	if !(*todos)[index].Completed {
		(*todos)[index].Completed = true
		now := time.Now()
		(*todos)[index].CompletedAt = &now
	}

	return (*todos)[index], nil
}

func (todos *Todos) Edit(id int, title string) (Todo, error) {
	index, err := todos.findIndex(id)
	if err != nil {
		return Todo{}, err
	}

	(*todos)[index].Title = title
	return (*todos)[index], nil
}

func (todos *Todos) Delete(id int) (Todo, error) {
	index, err := todos.findIndex(id)
	if err != nil {
		return Todo{}, err
	}

	deleted := (*todos)[index]
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return deleted, nil
}

func (todos Todos) Active() Todos {
	active := Todos{}
	for _, todo := range todos {
		if !todo.Completed {
			active = append(active, todo)
		}
	}
	return active
}

func (todos Todos) Completed() Todos {
	completed := Todos{}
	for _, todo := range todos {
		if todo.Completed {
			completed = append(completed, todo)
		}
	}
	return completed
}
