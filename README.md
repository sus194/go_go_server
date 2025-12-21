# go_go_server

A small Go server project organized around **typed request/response models** and **pluggable schedulers**.

This repo is structured with:
- `main.go` as the entrypoint
- `data_types/` for shared structs / models
- `schedulars/` for scheduler logic
- `go.mod` for module/dependency management :contentReference[oaicite:1]{index=1}

> If you’re using this as a learning repo: the main goal is to keep the codebase clean and modular—types in one place, scheduling logic in another, and the server wiring in `main.go`.

---

## Features
- Go HTTP server entrypoint (`main.go`)
- Centralized data models under `data_types/`
- Scheduler implementations under `schedulars/` (designed to be extended)
- Go modules support via `go.mod` :contentReference[oaicite:2]{index=2}

---

## Project Structure

```text
.
├── main.go
├── go.mod
├── data_types/
└── schedulars/
