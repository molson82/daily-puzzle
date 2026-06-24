# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Purpose

This repo holds the user's solutions to daily Golang coding puzzles that Claude generates to keep their skills sharp. The intended daily workflow (driven by a Claude Routine):

1. **Grade yesterday's solution** — run its unit tests, verify they pass, audit the code quality, and assign a grade.
2. **Recalculate ELO** — a progressive ELO/rank system (starting around "bronze" / Codewars level-8 kata) that scales difficulty up as the user proves their skills. Wins climb gradually across many problems of a tier before ranking up; failures cause a slight drop. One win does not jump a tier.
3. **Generate today's puzzle** — a new problem inspired by LeetCode, Codewars katas, Advent of Code (with much shorter problem descriptions), and HackerRank, at the difficulty matching the current ELO.

The user solves problems using the Go standard library, optionally with popular quality-of-life utility libraries.

The routine is driven by the task prompt in `docs/routine.md`. Progress (current ELO, tier, and the history of graded puzzles) is tracked in `PROGRESS.md` at the repo root — read it before grading and update + commit it after. Tier bands and the starting baseline (ELO 800, Bronze) are defined there.

## Layout

Each puzzle lives in its own dated directory: `cmd/<YYYY>/<MM>/<DD>/`. A puzzle directory contains four files (Advent-of-Code style, with two inputs):

- `main.go` — the user writes their solution here.
- `main_test.go` — Claude-generated unit tests the solution must pass.
- `sample.txt` — short, simple sample input.
- `input.txt` — the full input: longer, more complex, with edge cases the solution must handle.

When generating a new puzzle, create the directory for today's date and populate these four files (leave `main.go` for the user to fill in, or with a starter signature only).

## Commands

Run from the repo root. Replace the date path with the puzzle you're working on (e.g. today's `cmd/2026/06/24`).

```bash
# Run a day's solution
go run ./cmd/2026/06/24/

# Test a day's solution
go test ./cmd/2026/06/24/

# Run a single test
go test -run TestName ./cmd/2026/06/24/ -v

# Test / build everything
go test ./...
go build ./...

# Format and vet
gofmt -l -w .
go vet ./...
```
