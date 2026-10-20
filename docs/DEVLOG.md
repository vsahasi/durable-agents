# Devlog

Dated entries, newest first. Each entry says what changed, what was learned, and what is next. Milestones link to their tag.

## 2026-10-20: day 0

- Chose the project: a durable runtime for AI agents, after ruling out an FPGA computer (too much hardware), an inference engine (too tight on 8 GB), and a plain agent (too thin to learn from).
- Chose Go over Rust. See ADR 0001.
- Wrote the roadmap with acceptance criteria that are tests, and the design questions that have to be answered before each stage.
- Named it durable-agents. The name should say what it is; "ballast" lasted an hour.
- Installed Go 1.27.3 for Apple silicon under ~/sdk. Homebrew on this Mac is the Intel build under Rosetta, so its Go was Intel too.
- Next: stage 1 weekend goal. An append-only event log, a replay loop, a three-step toy agent, `kill -9` after step two, resume without re-calling the model.
