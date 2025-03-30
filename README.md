## Task-Tracker-Cli

This is a command-line interface (CLI) application designed to help you manage your to-do list efficiently. It is built using Go (requires version 1.23 or higher) and leverages SQLite as its database dependency.

### Features

- Add tasks with ease.
- Remove tasks you no longer need.
- Update task statuses (e.g., mark as completed).
- List all your tasks in one place.

### Usage

```bash
task-tracker-cli add 1 'Implement your own docs'
task-tracker-cli remove 2
task-tracker-cli update 1 -s Completed
task-tracker-cli list
```

### What's Next?

Future updates will include:

- A progress tracker to visualize your task completion.
- A Pomodoro timer to help you manage tasks within focused time intervals.
- Enhanced filtering and sorting options for better task organization.
- Cross-platform support for seamless usage.

Stay tuned for more features to enhance your productivity!
