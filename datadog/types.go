package datadog

import "fmt"

type Event struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
}

type Error struct {
	Code int
	Body string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Datadog Error: %d %s", e.Code, e.Body)
}
