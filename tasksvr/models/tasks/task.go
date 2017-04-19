package tasks

//NewTask represents a new task posted to the server
type NewTask struct {
	//TODO: fill out fields
}

//Task represents a task stored in the database
type Task struct {
	//TODO: fill out fields
}

//Validate will validate the NewTask
func (nt *NewTask) Validate() error {
	//TODO: implement
	return nil
}

//ToTask converts a NewTask to a Task
func (nt *NewTask) ToTask() *Task {
	//TODO: implement
	return nil
}
