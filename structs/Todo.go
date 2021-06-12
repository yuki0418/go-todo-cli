package structs

import (
	"time"
)

type Todo struct {
	id         int
	title      string
	desc       string
	date       time.Time
	isComplete bool
}
