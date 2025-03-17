# CLI Task Manager

A command-line task management application written in Go.

## Project Structure

```
.
├── cmd/
│   └── cli/            # Main application entry point
├── pkg/
│   ├── models/         # Data models
│   │   └── tests/      # Model-specific tests
│   ├── database/       # Database operations
│   │   └── tests/      # Database-specific tests
│   └── utils/          # Utility functions
│       └── tests/      # Utility-specific tests
├── internal/
│   └── handlers/       # Command handlers
│       └── tests/      # Handler-specific tests
└── tests/              # Top-level test directory
    ├── integration/    # Integration tests
    └── e2e/           # End-to-end tests
```

## Requirements

- Go 1.21 or higher

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

3. Build the application
```bash
go build -o task ./cmd/cli
```

## Testing

Run all tests:
```bash
go test ./...
```

Run specific test suite:
```bash
go test ./pkg/models/...  # Run model tests
go test ./tests/integration/...  # Run integration tests
go test ./tests/e2e/...  # Run end-to-end tests
```

## Usage

Documentation for usage will be added as the project develops.

## License

This project is licensed under the terms of the LICENSE file included in the repository.
