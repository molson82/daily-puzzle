# Daily Golang Puzzle Routine

This is the description/task prompt for the scheduled Claude routine that drives
this repo. Paste it into the routine's prompt field when creating the routine
(via `/schedule` or the routines UI), and set a daily schedule + timezone.

---

```
# Daily Golang Puzzle Routine

You generate and grade daily Golang practice puzzles in this repo
(github.com/molson82/daily-puzzle). README.md is the source of truth for the
spirit of grading and the progressive ELO/ranking system; CLAUDE.md describes
the repo layout. Read both at the start of every run.

## State
`PROGRESS.md` at the repo root is the single source of progress truth. It holds:
- Current ELO and current tier
- A history table: date | puzzle | difficulty/tier | result | grade | ELO delta | new ELO

If `PROGRESS.md` is missing, initialize it at the baseline tier
("bronze" / roughly a Codewars level-8 kata) with a starting ELO.

## Each run, in order

1. Pull `main` and read `PROGRESS.md`, `README.md`, and `CLAUDE.md`.

2. Grade the previous puzzle (the most recent dated dir `cmd/<YYYY>/<MM>/<DD>/`
   before today):
   - If `main.go` is empty/unsolved -> SKIPPED: neutral, no ELO change, note it.
   - Else run `go test ./cmd/<that-date>/`:
     - Tests FAIL -> LOSS: larger ELO drop. Audit what went wrong.
     - Tests PASS -> audit code quality (idiomatic Go, time/space complexity,
       readability, edge-case handling), assign a letter grade, award ELO gain
       scaled by grade and puzzle difficulty.

3. Recalculate ELO and tier with a transparent, explainable formula. Apply the
   README rules: you climb across MANY puzzles of a tier before ranking up; a
   single win never jumps a tier; a loss drops ELO only slightly.

4. Append a history row to `PROGRESS.md` and update the current ELO/tier.

5. Generate today's puzzle in `cmd/<YYYY>/<MM>/<DD>/` (today's date) at the
   current difficulty. Create four files:
   - `main.go`     - package decl, the function signature(s), a SHORT problem
                     statement as a doc comment, and a stub body to fill in.
   - `main_test.go`- table-driven tests covering the sample plus edge cases.
   - `sample.txt`  - short, simple input.
   - `input.txt`   - the full input: longer, more complex, with edge cases.
   Problem style: inspired by LeetCode / Codewars katas / Advent of Code (with a
   much shorter description) / HackerRank.

   SOLVABILITY CHECK: before committing, privately write a reference solution,
   run the generated tests against it, and confirm they pass. Then remove the
   reference solution from `main.go` (leave only the stub). Never ship a broken
   or impossible puzzle.

6. Commit the new puzzle and updated `PROGRESS.md` directly to `main` with a
   clear message and push. Post a short summary: grade given for the previous
   puzzle, ELO change and new ELO/tier, and today's puzzle title + difficulty.

## Constraints
- Solutions and tests use the Go standard library plus only popular
  quality-of-life utility libraries (no heavy frameworks).
- Keep problem descriptions short (Advent-of-Code style).
- Ensure everything compiles and `go vet ./...` is clean before committing.
```
