# Design

The questions this project has to answer, with the decision recorded once it is made. A question with "Decision: open" is still mine to think through. The references are where to read before deciding, not a substitute for deciding.

## 1. Event model

- What is an event? One per step, or one per observable transition inside a step?
- What is the schema, and how is it versioned when the workflow code changes mid-run?
- How much of the model's response is recorded: the final text, or the full raw response including tool calls and usage?

Decision: open.

References: Temporal's event history documentation; Martin Kleppmann, *Designing Data-Intensive Applications*, chapter 11 on event sourcing.

## 2. Replay and determinism

- Which operations are deterministic and re-executed on replay, and which are recorded and replayed? Candidates for recording: model calls, time, random numbers, network reads.
- How is nondeterminism detected when code changes produce a different event sequence than the log?
- Is replay the only recovery mechanism, or is there also a materialized snapshot of run state?

Decision: open.

References: Temporal's workflow determinism constraints; the Durable Functions replay model.

## 3. Side effects

- How is the idempotency key derived? Run ID plus step index, or a content hash?
- Outbox pattern: the effect intent is committed with the decision, then performed by a separate loop. What marks it done, and what if "done" itself fails to commit?
- What is the policy for external systems that offer no idempotency (plain SMTP, some HTTP APIs)?

Decision: open.

References: the transactional outbox pattern; Stripe's idempotency key design; DDIA chapter 7 on transactions.

## 4. Storage

- Append-only file per run, or SQLite with WAL? What does fsync cost on this machine, measured?
- How are torn writes detected: length prefix plus checksum, or SQLite's own guarantees?
- Retention: when is a completed run's log compacted or archived?

Decision: open.

References: SQLite's atomic commit and WAL documentation; the "Files are hard" literature on fsync semantics.

## 5. Sandbox and budgets

- WebAssembly via wasmtime, or Apple container VMs, or both behind one interface?
- Capability model: what is the smallest set of grants that still lets useful tools run?
- Budgets: enforced at step boundaries only, or mid-step by cancellation? What state does a budget-exhausted run end in?

Decision: open.

References: wasmtime-go; the WASI capability model; the object-capability literature.

## 6. Scheduling

- Fairness across runs and tenants. Priorities. Starvation.
- Backpressure when a provider rate limit is hit: queue, slow, or shed?
- Cancellation and graceful shutdown semantics.

Decision: open.

References: token bucket algorithms; Go's context cancellation patterns.

## 7. Testing

- Where are the fault injection points, and how are they made exhaustive rather than sampled?
- Deterministic simulation: how are the clock, network, and model abstracted so the simulator can replace them without the production code knowing?

Decision: open.

References: Will Wilson, "Testing Distributed Systems with Deterministic Simulation" (Strange Loop 2014); Jepsen's approach to checking histories.
