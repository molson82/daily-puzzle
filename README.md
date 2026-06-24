# Golang Daily Puzzle
This repo holds my solutions to daily puzzles created by Claude to keep my skill sharp in Golang.

The main purpose of this is to give me a range of difficult puzzles that challenge and refresh my memory of how to write proper Golang code. Claude will generate:
- The problem for me to solve
- The unit tests for me to use and test my solution against
- A sample input for me to use that is short and simple
- The full input that is usually longer, more complex, and has edge cases for my code to consider and catch in it's solution.

The files are similar to how Advent of Code (AOC) is with giving two inputs. One a small sample of the actual input. 

# Types of problems
The problems presented to me should be inspired by:
- Leet Code
- Codewars Katas
- Advent of Code (but much much shorter problem descriptions)
- HackerRank

All of these are great websites and tools that I have used in the past to test my skills and solve coding problems of various difficulties.

# How does the difficulty work?
Claude will determine the difficulty to test my Golang skills and knowledge. It will be a progressive ELO system where the problems start off as "bronze" or like Codewars a low level 8 kata. Then, once I prove my skills and knowlege, Claude will calculate my ELO (rank or pts) and scale up the difficulty to higher ranks and allowing me to solve more and more difficult questions. 

Just like a regular ELO system, I can fail and fall down slightly. And, a win doesn't automatically slingshot me to the next rank. You have to climb to various tiers by solving many problems of that tier covering various ground before ranking up. 

# How will I get graded?
I will solve the code using the Golang standard package and maybe some popular dependencies. Only util libraries that give quality of life functions that speed up problem solving. 

I will be graded by Claude the next day before generating the next question. Claude will run the unit tests against my code, verify they pass, and then audit my solution to give me a grade and get a sense of my skills and expertise. Then, after seeing my quality of code and grade, Claude will recalculate my ELO and give me my new problem.

# More info
- [CLAUDE.md](CLAUDE.md) — repo layout, conventions, and commands for Claude/contributors.
- [docs/routine.md](docs/routine.md) — the description/task prompt for the daily Claude Routine that pulls in this repo and performs all the actions (grade, recalculate ELO, generate the next puzzle).
- [PROGRESS.md](PROGRESS.md) — current ELO, tier, and the running history of graded puzzles.
