# Beans - Agentic Issue Tracker

This project uses **beans**, an agentic-first issue tracker. Issues are called "beans", and you can
use the "beans" CLI to manage them.

All commands support --json for machine-readable output. Use this flag to parse responses easily.

## Core Rules

- Track ALL work using beans (no TodoWrite tool, no markdown TODOs)
- Use `beans new` to create issues, not TodoWrite tool
- After compaction or clear, run `beans prompt` to re-sync

## Finding work

- `beans list --json` to list all beans
