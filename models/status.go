package models

import "strings"

const (
	Pending    string = "pending    "
	InProgress string = "in progress"
	Completed  string = "completed  "
	Aborted    string = "aborted    "
	ToDo       string = "to do      "
)

func IsValidStatus(status string) (string, bool) {
	normalizedStatus := strings.TrimSpace(status)
	switch normalizedStatus {
	case strings.TrimSpace(Pending):
		return Pending, true
	case strings.TrimSpace(InProgress):
		return InProgress, true
	case strings.TrimSpace(Completed):
		return Completed, true
	case strings.TrimSpace(Aborted):
		return Aborted, true
	case strings.TrimSpace(ToDo):
		return ToDo, true
	default:
		return "", false
	}
}
