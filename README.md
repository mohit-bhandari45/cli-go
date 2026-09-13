# gurl - Concurrent HTTP / API Tester CLI

`gurl` is a custom HTTP client and load testing CLI tool built in Go. It combines the single-request capabilities of `curl` with the load-testing and benchmarking capabilities of `hey` / `vegeta`.

---

## 🚀 Features

- **Custom HTTP Requests**: Supports `GET`, `POST`, `PUT`, `DELETE` methods with custom headers (`-H`) and JSON body payloads (`-d`).
- **Concurrent Benchmarking**: Runs load tests by spawning $N$ concurrent worker goroutines sending $M$ total requests.
- **Latency Distribution Statistics**: Calculates p50 (median), p95, and p99 latency percentiles, min, max, average, and throughput (Requests Per Second).
- **Status Code Breakdown**: Tracks success and failure rates along with HTTP status code distributions (2xx, 4xx, 5xx).
- **Rich Output Formatting**: Beautiful, colorized terminal summary or raw JSON output (`-o json`).

---

## 🛠️ Go Concepts Covered

- **`net/http`**: Connection pooling (`http.Transport`), `http.Client`, custom headers, and request bodies.
- **Goroutines & Channels**: Channel-based worker pool pattern for managing concurrency.
- **`sync.WaitGroup`**: Synchronizing concurrent worker goroutines.
- **High-Resolution Math**: Calculating time durations (`time.Since`) and percentile distributions.
- **CLI Ecosystem**: Using `cobra` for CLI flags/commands and `fatih/color` for terminal styling.

---

## 📁 Project Structure (Planned)

```text
cli/
├── README.md                   # Project documentation
├── go.mod                      # Go module definition
├── main.go                     # Entry point initializing CLI commands
├── cmd/
│   └── root.go                 # Cobra command flags and execution handling
└── internal/
    ├── client/                 # Configured net/http client factory
    ├── runner/                 # Concurrency worker pool engine & timing
    ├── stats/                  # Latency percentile & throughput calculator
    └── printer/                # Colorized terminal & JSON formatters
```
