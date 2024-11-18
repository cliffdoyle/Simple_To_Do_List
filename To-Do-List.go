package main

import "fmt"

type Task struct {
	ID          int
	Priority    int
	Description string
}

type Tasks struct {
	te []Task
}

func (to *Tasks) AddTask(priority int, des string) {
	task := Task{}
	task.ID = len(to.te) + 1
	task.Description = des
	if priority > 0 && priority <= 3 {
		task.Priority = priority
		to.te = append(to.te, task)
	}else{
	 	fmt.Println("Sorry mate, can't add that!!")
	}

}

func (to *Tasks) ViewTask() {
	for _, task := range to.te {
		fmt.Println(task.ID, task.Priority, task.Description)
	}
}

func (to *Tasks) DeleteTask(id int) {
	for i, task := range to.te {
		if task.ID == id {
			to.te = append(to.te[:i], to.te[i+1:]...)
			fmt.Printf("Task with ID %d deleted.\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
}

func (to *Tasks) UpdateTask(id, prio int, des string) {
	for i, task := range to.te {
		if task.ID == id {
			to.te[i].Priority = prio
			to.te[i].Description = des
			fmt.Printf("Task with ID %d updated.\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
}

func main() {
	t := &Tasks{}
	// Add tasks
	t.AddTask(2, "Hello World")
	t.AddTask(3, "Task Two")
	t.AddTask(1, "High Priority Task")
	t.AddTask(8, "High Priority Task wehsbcx")

	// View tasks
	t.ViewTask()

	// Update a task
	t.UpdateTask(2, 5, "Updated Task Two")
	t.ViewTask()

	// Delete a task
	t.DeleteTask(1)
	t.ViewTask()

	// Try to delete a non-existent task
	t.DeleteTask(10)
}
