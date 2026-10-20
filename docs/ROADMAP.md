# Roadmap

Five stages. Each one is shippable on its own, has acceptance criteria that are tests, and names the evidence it publishes. The "claim" line is the sentence the stage earns: it goes on the resume and in the README only once the criteria pass.

## Stage 1: Event log and deterministic replay

**Goal.** A run is a sequence of events in an append-only log. Replaying the log reproduces the run's state. Model calls are recorded on first execution and replayed, never re-issued.

**Acceptance criteria.**
- A three-step toy agent runs to completion through the runtime.
- The process is killed with SIGKILL at every event boundary (one test run per boundary). On restart, the run finishes without re-calling the model, and the final state matches an uninterrupted run byte for byte.
- A partially written final record (torn write) is detected and discarded on open. The log is fsynced before an event is considered committed.
- Replaying a log against changed workflow code that would produce a different event sequence is detected and reported as a nondeterminism error, not silently accepted.

**Evidence.** The boundary-kill test and its count. A terminal recording of `kill -9` followed by resume.

**Claim.** Event-sourced agent runtime with deterministic replay; runs survive SIGKILL at any point and resume without repeating model calls.

## Stage 2: Exactly-once side effects

**Goal.** A tool call that changes the world happens exactly once from the world's point of view, no matter where the runtime crashes.

**Acceptance criteria.**
- Each effect carries an idempotency key derived from the run and step. Effects go through an outbox that is committed in the same write as the decision to perform them.
- A fake external world records every effect it receives. A fault injector crashes the runtime before and after every effect across at least 10,000 runs. The world shows zero duplicate and zero missing effects.
- Retries use exponential backoff with jitter. A step that fails permanently ends the run in a defined terminal state with the failure recorded.
- Effects against an external system that offers no idempotency are documented as at-least-once, with the compensating action named.

**Evidence.** Fault-injection results table in the README: runs, crash points, duplicates, losses.

**Claim.** Exactly-once tool effects proven by fault injection at every crash point across 10,000+ runs.

## Stage 3: Sandbox and budgets

**Goal.** Code the agent writes cannot touch anything it was not granted, and cannot spend more than it was given.

**Acceptance criteria.**
- Agent-written code runs in WebAssembly via wasmtime. Filesystem paths, network hosts, and clock access are capabilities granted per run, denied by default.
- CPU is metered with fuel and memory is capped. Exceeding either terminates the step with a recorded reason.
- Token, dollar, and wall-clock budgets are enforced by the runtime per run. A run that exhausts a budget stops at the next step boundary with a defined state.
- An escape test suite (read a path outside the grant, open a socket to an ungranted host, spin forever, allocate without bound) passes by failing every attempt.

**Evidence.** The escape suite. A short note comparing WebAssembly isolation with Apple's container VMs on macOS, with measurements.

**Claim.** Capability-based WebAssembly sandbox for agent-written code with fuel, memory, and spend budgets enforced by the runtime.

## Stage 4: Scheduler and observability

**Goal.** Many runs at once, fairly, without overrunning provider limits, and with every step visible.

**Acceptance criteria.**
- N concurrent runs with priorities. Provider rate limits are modeled as token buckets; the scheduler applies backpressure rather than failing calls.
- Every step is a trace span with parent run, timing, tokens, and cost. Traces are viewable locally and exportable as OTLP.
- Graceful shutdown drains in-flight steps to an event boundary before exit.

**Evidence.** Load test: runs per minute, p50 and p99 step latency, runtime overhead versus calling the model directly.

**Claim.** Scheduler with provider-aware backpressure and per-step tracing; measured overhead of X ms per step at N concurrent runs.

## Stage 5: Deterministic simulation and a real workload

**Goal.** Whole multi-agent runs reproduce from a seed, and the runtime carries production work.

**Acceptance criteria.**
- The clock, network, and model are replaced by simulated versions driven from a seed. A failing run reproduces exactly from its seed.
- Schedule fuzzing (randomized interleavings and injected faults) finds at least one real bug, which is fixed with a regression seed committed.
- The outbound email pipeline runs on the runtime nightly for at least two weeks, with reliability and cost numbers compared to the previous implementation.

**Evidence.** A write-up with the numbers and the bugs the simulator found.

**Claim.** Deterministic simulation testing in the style of FoundationDB; runtime in production for a nightly multi-step agent workload.
