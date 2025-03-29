package examples

import (
	"fmt"
)

// ANSI color codes
const (
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
	Blink     = "\033[5m"
	Reverse   = "\033[7m"
	Strikeout = "\033[9m"

	// Foreground colors
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	// Background colors
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

// RunAnsiExample demonstrates using ANSI escape codes for console styling
func RunAnsiExample() {
	// Print a styled title
	fmt.Println(BgBlue + Bold + White + Blink + " Task Tracker CLI " + Reset)
	fmt.Println()

	// Sample tasks
	tasks := []struct {
		ID        int
		Name      string
		Completed bool
	}{
		{1, "Implement task tracker CLI", false},
		{2, "Add beautiful console output", false},
		{3, "Write documentation", true},
	}

	// Print tasks with styling
	for _, task := range tasks {
		taskID := fmt.Sprintf("%d.", task.ID)

		if task.Completed {
			// Completed tasks are strikeout and gray
			fmt.Printf("%s %s%s%s ✓%s\n",
				Cyan+taskID+Reset,
				Strikeout,
				White,
				task.Name,
				Reset)
		} else {
			// Pending tasks are bold and white
			fmt.Printf("%s %s%s%s\n",
				Cyan+taskID+Reset,
				Bold,
				White,
				task.Name+Reset)
		}
	}
}
