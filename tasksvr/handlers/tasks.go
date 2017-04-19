package handlers

import (
	"net/http"
)

//HandleTasks will handle requests for the /v1/tasks resource
func (ctx *Context) HandleTasks(w http.ResponseWriter, r *http.Request) {
}

//HandleSpecificTask will handle requests for the /v1/tasks/some-task-id resource
func (ctx *Context) HandleSpecificTask(w http.ResponseWriter, r *http.Request) {
}
