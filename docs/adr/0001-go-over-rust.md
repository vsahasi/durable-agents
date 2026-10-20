# ADR 0001: Go over Rust

Date: 2026-10-20. Status: accepted.

## Context

The runtime needs concurrency, a sandbox host, durable storage, and a lot of test harness code. Rust was the strongest systems signal and has the native wasmtime API. Go is what Temporal is written in, has a simpler concurrency story, a built-in race detector, and a faster path from idea to running code.

## Decision

Go. The learning target is durable execution, exactly-once effects, sandboxing, and simulation testing. Those are design problems, and Go gets out of the way of them. wasmtime is available through wasmtime-go.

## Consequences

- wasmtime-go uses cgo, so the toolchain must be native arm64 on Apple silicon. The Intel build under Rosetta is not acceptable for this project.
- `go test -race` runs in CI on every push.
- If the sandbox or performance work later outgrows Go, the sandbox host is the piece to rewrite, behind the interface defined in stage 3.
