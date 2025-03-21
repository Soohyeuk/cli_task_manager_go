# Simple CLI Task Manager

A minimal command-line task management application written in Go.

## Features

- Add new tasks
- List all tasks
- Mark tasks as complete
- Delete tasks
- Store tasks in a simple JSON file

## Project Structure

```
.
├── cmd/
│   └── cli/         # Main application code
└── pkg/
    └── models/      # Task data model
```

## Requirements

- Go 1.21 or higher

## Setup

1. Clone the repository
```bash
git clone https://github.com/soohyeuk/cli_task_manager_go.git
cd cli_task_manager_go
```

2. Build and run
```bash
go build -o task ./cmd/cli
./task
```

## Usage

```bash
./task add "Buy groceries"     # Add a new task
```

## License

This project is licensed under the terms of the LICENSE file included in the repository.
