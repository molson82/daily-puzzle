# Progress

Single source of truth for puzzle-solving progress. Updated by the daily routine
(see [docs/routine.md](docs/routine.md)).

## Current

- **ELO:** 830
- **Tier:** Bronze

## Tiers

Climb across many puzzles within a tier before ranking up. A single win never
jumps a tier; a loss drops ELO only slightly.

| Tier     | ELO range   |
| -------- | ----------- |
| Bronze   | 0 – 1199    |
| Silver   | 1200 – 1399 |
| Gold     | 1400 – 1599 |
| Platinum | 1600 – 1799 |
| Diamond  | 1800 – 1999 |
| Master   | 2000+       |

## ELO formula

Transparent and explainable. A win climbs gradually so a tier takes many
puzzles; a loss only nudges you down.

- **Win (tests pass):** `ΔELO = round(40 × gradeFactor)`
- **Loss (tests fail):** `ΔELO = -20`
- **Skipped (unsolved):** `ΔELO = 0`

| Grade | A+  | A   | A-   | B+   | B    | B-   | C   | D    |
| ----- | --- | --- | ---- | ---- | ---- | ---- | --- | ---- |
| factor| 1.0 | 0.9 | 0.85 | 0.75 | 0.65 | 0.55 | 0.4 | 0.25 |

## History

| Date | Puzzle | Difficulty / Tier | Result | Grade | ELO Δ | New ELO |
| ---- | ------ | ----------------- | ------ | ----- | ----- | ------- |
| 2026-06-24 | Digital Root Sum | Bronze (ELO 800) | GENERATED | — | 0 | 800 |
| 2026-06-25 | Digital Root Sum | Bronze (ELO 800) | WIN | B+ | +30 | 830 |
| 2026-06-25 | Balanced Brackets | Bronze (ELO 830) | GENERATED | — | 0 | 830 |
<<<<<<< HEAD
| 2026-06-26 | Balanced Brackets | Bronze (ELO 830) | SKIPPED | — | 0 | 830 |
| 2026-06-26 | Run-Length Encoding | Bronze (ELO 830) | GENERATED | — | 0 | 830 |
=======
| 2026-06-27 | Balanced Brackets | Bronze (ELO 830) | SKIPPED | — | 0 | 830 |
| 2026-06-27 | Warehouse Shelves | Bronze (ELO 830) | GENERATED | — | 0 | 830 |
>>>>>>> origin/claude/youthful-maxwell-k163yl

### 2026-06-25 — Digital Root Sum (graded), grade B+

Correct and passes all tests, including the edge cases (blank lines, embedded
whitespace, huge numbers). Clean, modern Go (`strings.SplitSeq`). Points to
improve, in priority order:

- **Don't parse each digit with `strconv.ParseInt`.** For a digit rune just use
  `r - '0'` — no allocation, no error to discard. The current code leans on
  `ParseInt(" ")` silently returning 0 to "handle" whitespace, which masks real
  errors rather than handling them.
- **Skip the string round-trip in the recursion.** `digitalRoot` converts the
  running sum back to a string via `strconv.Itoa` just to recurse. Keep summing
  ints instead.
- **There's an O(1) closed form.** For a non-negative `n`, the digital root is
  `0` if `n == 0`, else `1 + (n-1) % 9`. Worth knowing — it turns the whole
  per-number loop into one expression once the line is parsed.

Net: a solid, correct Bronze solution; the deductions are about idiom and
efficiency, not correctness. Hence B+ (factor 0.75 → +30 ELO).

<<<<<<< HEAD
### 2026-06-26 — Balanced Brackets, SKIPPED

`cmd/2026/06/25/main.go` was still the untouched stub (`// TODO: implement`,
`countBalanced` returning 0 and `isBalanced` returning false), so the puzzle was
unsolved. Per the routine this is neutral: no audit, no grade, ELO unchanged at
830 (Bronze). The Balanced Brackets puzzle stays on the board if you want to
come back to it — a classic stack problem: push openers, pop and match on
closers, reject on mismatch or leftover openers at the end.
=======
### 2026-06-27 — Balanced Brackets (skipped)

`cmd/2026/06/25/main.go` was left as the generated stub (`countBalanced`
returns 0, `isBalanced` returns false), so there was nothing to grade. Per the
routine, an unsolved puzzle is neutral: no ELO change. ELO stays 830, Bronze.

Side note: the "reference solution for june 25" commit added a learning
`reference.go` for the *already-graded* 06/24 Digital Root Sum puzzle (folding
in the prior feedback — `r - '0'`, the closed-form digital root, and the
mod-9 trick to stay overflow-proof). Nice follow-up, but it didn't touch the
06/25 puzzle, which remains open. If you want a grade for Balanced Brackets,
fill in its stub and it'll be picked up next run.

For reference, the idiomatic shape: scan runes, push the *expected closer* on a
stack for each opener (`( → )`, `[ → ]`, `{ → }`), and on a closer pop and
compare. Balanced ⇔ every closer matches the top and the stack is empty at the
end. Non-bracket runes are ignored. A `map[rune]rune` of opener→closer keeps it
tidy; a `[]rune` slice is a perfectly good stack.
>>>>>>> origin/claude/youthful-maxwell-k163yl
