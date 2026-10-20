# durable-agents

A durable runtime for AI agents, written in Go.

It runs long-horizon agents so that a crash, a redeploy, or a `kill -9` at any instant loses nothing. The run resumes exactly where it was, no model call is repeated, and no side effect happens twice. Code the agent writes runs in a sandbox it cannot escape, inside token, dollar, and wall-clock budgets the agent cannot override.

**Status: day 0 (2026-10-20). Nothing works yet.** The roadmap below is the plan and the [devlog](docs/DEVLOG.md) is the record. Every stage has acceptance criteria that are tests, not adjectives.

## The promise

- **Durable.** Every step of a run is an event in an append-only log. Replaying the log reconstructs the run. Kill the process anywhere and it continues from the last event.
- **Exactly once.** Side effects (send the email, place the order, write the file) carry idempotency keys and go through an outbox. A crash between "decided to send" and "sent" cannot produce two emails or zero.
- **Sandboxed.** Agent-written code runs in WebAssembly with explicit capabilities. No filesystem, no network, no clock unless granted. CPU fuel and memory are capped.
- **Budgeted.** Tokens, dollars, and wall-clock time are enforced by the runtime, per run and per tenant.
- **Testable.** A fault injector crashes the runtime before and after every effect. A deterministic simulator replaces the clock, the network, and the model so an entire multi-agent run reproduces from a seed.

## Why

The loop around a model call is the easy part of an agent. The hard part is everything that makes it safe to leave running: durability, exactly-once effects, isolation, scheduling, and a way to test failure modes you cannot reproduce by hand. That is what this project is for, and it is built to run real workloads, starting with my own.

## Roadmap

| Stage | Milestone | Status |
|---|---|---|
| 1 | Event log and deterministic replay. Survives `kill -9` at any point. | not started |
| 2 | Exactly-once side effects, proven by fault injection at every boundary. | not started |
| 3 | WebAssembly sandbox with capabilities and budgets. | not started |
| 4 | Scheduler for many concurrent runs, with tracing. | not started |
| 5 | Deterministic simulation testing, then a real production workload. | not started |

Acceptance criteria for each stage are in [docs/ROADMAP.md](docs/ROADMAP.md).

## Design principles

1. The event log is the only source of truth. Everything else is a cache of it.
2. Anything nondeterministic (model output, time, randomness, network) is recorded once and replayed forever.
3. Every side effect is idempotent or it does not exist.
4. The runtime, not the agent, enforces budgets and permissions.
5. If a failure mode cannot be injected in a test, it is not handled.

## Non-goals

This is not a prompt framework, not an agent, and not a hosted service. It is the layer agents run on.

## Development

```sh
make test   # go vet + go test -race ./...
```

Requires Go 1.27 or later, native arm64 on Apple silicon.

## Documents

- [docs/ROADMAP.md](docs/ROADMAP.md): stages, acceptance criteria, and the evidence each stage publishes
- [docs/DESIGN.md](docs/DESIGN.md): the design questions, with decisions recorded as they are made
- [docs/adr/](docs/adr/): architecture decision records
- [docs/DEVLOG.md](docs/DEVLOG.md): dated progress log

## License

MIT. Built by [Veer Sahasi](https://github.com/vsahasi).
