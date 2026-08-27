# Flag Practice Command

This is a separate, non-persistent command for practicing Go's standard
[`flag`](https://pkg.go.dev/flag) package. It prints the action it would perform
without changing the main application's `todos.json` file.

Run it from the repository root:

```bash
go run ./flag-practice -add "Learn flags"
go run ./flag-practice -add "Ship project" -due 2026-09-01
go run ./flag-practice -list
go run ./flag-practice -done 2
go run ./flag-practice -delete 2
go run ./flag-practice -clear
go run ./flag-practice -help
go run ./flag-practice -version
```

## What to Study

In `main.go`, look at:

1. The `options` struct that keeps all flag values together.
2. `StringVar`, `BoolVar`, and `IntVar` for different value types.
3. Default values such as `""`, `false`, and `-1`.
4. `flag.Parse()` for reading the flags from the command line.
5. `execute` and its `switch` for deciding which flag to act on.

## Practice Ideas

- Add a `-priority` string flag.
- Add a `-completed` filter for `-list`.
- Validate that `-due` uses `YYYY-MM-DD`.
- Add short aliases such as `-a` for `-add`.
- Replace the printed actions with a separate JSON storage file.
