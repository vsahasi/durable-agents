# durable-agents

A runtime for AI agents, in Go.

An agent run is a sequence of steps: model calls, tool calls, and the code between them. The runtime records each step in an append-only log and reconstructs a run from its log. A process that dies at any point resumes the run from the last committed step without repeating a model call or a side effect. Code the agent writes runs in a sandbox with explicit capabilities, and every run has token, dollar, and time budgets enforced by the runtime.

Status: early. Nothing runs yet.

## Parts

**Event log.** One append-only log per run. Each record is length-prefixed and checksummed, and a record is committed only after fsync. A torn final record is detected and discarded on open. The log is the only source of truth; all other run state is derived from it.

**Replay.** Workflow code is deterministic and is re-executed against the log to rebuild state. Nondeterministic inputs are recorded the first time they happen and served from the log on replay: model responses, tool results, time, random numbers. If re-execution produces an event sequence that diverges from the log, the run halts with a nondeterminism error instead of continuing on wrong state.

**Effects.** A tool call that changes the world is an effect. The intent to perform it is committed to the log in the same write as the decision that produced it, then performed by a separate loop and marked done. Each effect carries an idempotency key derived from run and step, so a crash between intent and done produces a retry with the same key, not a duplicate. Retries use exponential backoff with jitter. A permanently failing step ends the run in a defined terminal state.

**Sandbox.** Agent-written code runs in WebAssembly through wasmtime. Filesystem paths, network hosts, and the clock are capabilities granted per run and denied by default. CPU is metered as fuel and memory is capped; exceeding either terminates the step with a recorded reason.

**Budgets.** Tokens, dollars, and wall-clock time per run, checked by the runtime at step boundaries. A run that exhausts a budget stops in a defined state.

**Scheduler.** Many runs execute concurrently with priorities. Provider rate limits are modeled as token buckets and the scheduler applies backpressure rather than letting calls fail. Shutdown drains in-flight steps to an event boundary.

**Tracing.** Every step is a span with parent run, timing, token counts, and cost, exportable as OTLP.

**Testing.** A fault injector crashes the runtime before and after every effect and checks the external world for duplicates and losses. A simulator replaces the clock, network, and model with seeded versions so a complete multi-agent run reproduces from a seed and interleavings can be fuzzed.

## Build

Go 1.27 or later.

    make test    # go vet + go test -race ./...

## License

MIT.
