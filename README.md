# Task Manager CLI

A simple command-line task management application built with Go.

Sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## Overview
Manage your tasks efficiently with basic CRUD operations and persistent JSON storage.

## Features
- ✨ Create new tasks
- 📝 Edit task descriptions
- 🔄 Update task status (todo/in-progress/done)
- 🗑️ Delete tasks
- 📋 List all tasks
- 💾 Automatic JSON persistence

## Installation

```bash
# Clone the repository
git clone https://github.com/mazyaryousefinia/cli-task-tracker.git

# Navigate to project directory
cd cli-task-tracker

# Build the application
go build ./main.go
```

## Usage


Start the application with initial command:
```bash
./main.exe --command=add
```

### Available Commands
- `add` - Create a new task
- `update` - Modify task description
- `delete` - Remove a task
- `update-status` - Change task status
- `list` - Show all tasks
- `exit` - Close application
