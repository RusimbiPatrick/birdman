# Birdman

Birdman is a distributed **orchestrator** designed to schedule, execute, and monitor tasks across multiple worker nodes. It provides a robust framework for coordinating tasks, collecting statistics, and ensuring task health in a distributed environment.

## Features

- **Task Scheduling**: Efficiently schedule tasks to worker nodes using various scheduling algorithms.
- **Task Execution**: Execute tasks on worker nodes with support for Docker containers.
- **Task Monitoring**: Monitor the state of tasks and ensure their health.
- **Statistics Collection**: Collect and report statistics from worker nodes.
- **REST API**: Expose APIs for managing tasks, workers, and nodes.
- **Extensible Architecture**: Modular design for easy integration and customization.

## Project Structure

The project is organized as follows:

```
.
├── cmd/                # Command-line interface for interacting with the system
├── echo/               # Example echo service
├── manager/            # Manager component for task scheduling and coordination
├── node/               # Node abstraction for worker nodes
├── scheduler/          # Task scheduling algorithms
├── stats/              # Statistics collection and reporting
├── store/              # Data storage layer
├── task/               # Task definitions and state management
├── utils/              # Utility functions
├── worker/             # Worker component for task execution
├── add_task.json       # Example task definition
├── stop_task.json      # Example stop task definition
├── task*.json          # Example task files
├── go.mod              # Go module definition
├── go.sum              # Go module dependencies
└── README.md           # Project documentation
```

## Components

### Manager
The manager is responsible for:
- Scheduling tasks to worker nodes.
- Monitoring the health of tasks and nodes.
- Collecting statistics from worker nodes.

### Worker
The worker is responsible for:
- Executing tasks assigned by the manager.
- Reporting task status and statistics.
- Managing task lifecycle (start, stop, restart).

### Scheduler
The scheduler determines the best worker node for a task based on available resources and scheduling algorithms.

### Node
The node represents a worker in the system and provides APIs for collecting statistics and managing tasks.

### Task
The task package defines the structure and state of tasks, including transitions between states.

### Stats
The stats package collects and reports metrics such as memory usage, disk usage, and task counts.

### Store
The store package provides data storage for tasks and events, with support for in-memory and persistent storage.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/birdman.git
   cd birdman
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Build the project:
   ```sh
   go build ./...
   ```

## Usage

### Starting the Manager
Run the manager to coordinate tasks:
```sh
go run manager.go
```

### Starting a Worker
Run a worker to execute tasks:
```sh
go run worker.go --name worker1
```

### Adding a Task
Use the REST API to add a task:
```sh
curl -X POST -H "Content-Type: application/json" -d @add_task.json http://localhost:8080/tasks
```

### Checking Task Status
Use the `status` command to view task statuses:
```sh
go run status.go --manager localhost:8080
```

## Configuration

- **Database Type**: Configure the task database as `memory` or `persistent`.
- **Scheduler Type**: Choose from `greedy`, `roundrobin`, or `epvm` scheduling algorithms.

## API Endpoints

### Manager
- `POST /tasks`: Add a new task.
- `GET /tasks`: List all tasks.
- `DELETE /tasks/{taskID}`: Stop a task.

### Worker
- `GET /stats`: Retrieve worker statistics.
- `GET /tasks`: List tasks on the worker.
- `POST /tasks`: Start a task.

## Contributing

Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

## Acknowledgments

- [Go](https://golang.org/) for providing a powerful programming language.
- [Cobra](https://github.com/spf13/cobra) for CLI support.
- [Chi](https://github.com/go-chi/chi) for lightweight HTTP routing.
- **Tim Boring** - Author of *Build an Orchestrator from Scratch*, which inspired this project.

## Contact

For questions or feedback,.

## Contact

For questions or feedback, please contact [contactrusimb@gmail.com](mailto:yourname@example.com).
```