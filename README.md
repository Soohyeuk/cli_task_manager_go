# CLI Task Manager

A command-line task management application written in Go.

## Project Structure

```
.
├── go.mod
├── Makefile
├── cmd/
│   └── cli/            # Main application entry point
├── pkg/
│   ├── models/         # Data models
│   │   └── tests/      # Unit test
│   ├── database/       # Database operations
│   │   └── tests/      # Unit test
│   └── utils/          # Utility functions
│       └── tests/      # Unit test
├── internal/
│   └── handlers/       # Command handlers
│       └── tests/      # Unit test
└── tests/              # Top-level test directory
    ├── integration/    # Integration tests
    └── e2e/           # End-to-end tests
```

## Requirements

- Go 1.21 or higher
- Make (for using Makefile commands)

## Setup

1. Clone the repository
```bash
git clone https://github.com/soohyeuk/cli_task_manager_go.git
cd cli_task_manager_go
```

2. Install dependencies
```bash
go mod tidy
```

## Available Make Commands

You can use the following make commands for development:

```bash
make help           # Show all available commands
make build         # Build the application
make test          # Run all tests
make clean         # Clean build files
make run           # Run the application
make fmt           # Format all Go files
make vet           # Run Go vet
make lint          # Run linter
make test-coverage # Run tests with coverage report
```

## Development Workflow

1. Format your code:
```bash
make fmt
```

2. Run tests:
```bash
make test
```

3. Check code quality:
```bash
make vet
make lint
```

4. Build and run:
```bash
make build
make run
```

## Usage

Documentation for usage will be added as the project develops.

## License

This project is licensed under the terms of the LICENSE file included in the repository.
