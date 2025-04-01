package models

const (
	Pending    string = "pending    "
	InProgress string = "in progress"
	Completed  string = "completed  "
	Aborted    string = "aborted    "
	ToDo       string = "to do      "
)

func IsValidStatus(status string) bool {
	switch status {
	case Pending, InProgress, Completed, Aborted, ToDo:
		return true
	default:
		return false
	}
}
