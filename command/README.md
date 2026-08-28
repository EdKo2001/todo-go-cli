# Manual CLI Implementation

This folder contains the original todo CLI built without Cobra. It is kept as
a learning example for manually parsing commands, arguments, and flags.

The newer Cobra implementation lives in `../cmd` and starts from `../main.go`.

## Run

From the repository root:

```bash
go run ./command help
go run ./command add "Learn Go"
go run ./command add "Learn flags" --due 2026-09-01
go run ./command list
go run ./command done 1
go run ./command edit 1 "Learn Go CLI"
go run ./command delete 1
go run ./command clear
```

## Build

```bash
go build -o todo-manual ./command
```

On Windows:

```powershell
go build -o todo-manual.exe ./command
```

## How It Works

`main.go` loads `todos.json`, passes `os.Args[1:]` to `Execute`, and saves the
list when a command changes it.

`command.go` separates the first argument from the remaining arguments:

```go
command := args[0]
commandArgs := args[1:]
```

It then selects a command with a `switch`:

```go
switch command {
case "add":
    return runAdd(commandArgs, todos)
case "list":
    todos.Print()
}
```

The `add` command manually scans its arguments for `--due`. This demonstrates
the parsing work that a CLI library such as Cobra normally handles.

## Files

```text
command/
├── README.md
├── command.go   Command parsing and handlers
├── main.go      Program entry point
├── storage.go   JSON loading and saving
└── todo.go      Todo model and methods
```

Both the manual and Cobra implementations use `todos.json` in the directory
from which the program is run.
