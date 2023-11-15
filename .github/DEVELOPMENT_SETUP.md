# Development Setup

## Requirements and Initial Setup
To work with this repository, your system should be provisioned with the [latest version](https://go.dev/doc/install) 
of Golang and the [GNU Make](https://www.gnu.org/software/make/) tool. With those
requirements met and the repository cloned please be sure to run the following 
command in the repository's root to complete the initial setup.
```
make
```

## Makefile Targets
The `Makefile` in the repository's root provides targets for automating certain
mundane actions. A brief explanation for each such target is provided below.

### `setup`
This target sets up the git hooks provided in the `.git-hooks` directory and makes
all the shell scripts present in the `scripts` directory executable. Omitting `setup` 
and just running `make` as in "Requirements and Initial Setup" section above 👆,
results in the same behaviour.
```
make setup
```

### `init`
This target creates a [Go module](https://go.dev/blog/using-go-modules) for a single
puzzle's solution. Suppose we are attempting to solve the ["Calorie Counting"](https://adventofcode.com/2022/day/1)
which is the puzzle for day 1 of the year 2022. The following command will create
said Go module with some starter code, and include it in the `go.work` file.
```
make init YEAR=2022 DAY=1 NAME="Calorie Counting" 
```

### `run`
This target will run your solution against the input provided by Advent of Code for
the puzzle in question, for the "Calorie Counting" puzzle mentioned above 👆 this 
target will be invoked as follows.
```
make run YEAR=2022 DAY=1 INPUT="./path/to/adventofcode/puzzle/input.txt"
```

If the puzzle in question has multiple parts and you wish to execute the solution 
for a specific part then this can be done in the following manner. We assume that
the `SolutionPart` variable exposed by the `utilities` package (which corresponds
to the provide `PART` argument), is used in the solution code to account for the
different parts.
```
make run YEAR=2022 DAY=1 PART=2 INPUT="./path/to/adventofcode/puzzle/input.txt"
```

### `test`
This target will run all provided unit tests for all puzzle solutions and the Go 
modules in `packages` directory.
```
make test
```

### `format`
This target will format your source code using the `gofmt` tool, normally you would
not need to run this as a git hook is in place to format your code before committing.
```
make format
```

### `version`
This target is intended for setting the `Go` version for all packages and solution
modules to the latest version, this target is _not_ intended for use by contributors. 
The repository owner is reponsible for bumping all `Go` modules to the latest version.
```
make version GO=1.21.4
```

### `solution`
This target will commit your solution with a pre-defined commit message, this will
ensure that all commits carrying a solution follow a consistent pattern, keeping 
with the "Calorie Counting" example the following command will create the commit 
message `solution: year-2022/day-01`.
```
make solution YEAR=2022 DAY=1
```

### `improve`
This target will open your git editor with a pre-existing commit template for 
commiting improvements you have made to a **single** solution, keeping with the
"Calorie Counting" example the following command
```
make improve YEAR=2022 DAY=1
```

will open the configured editor with the following message, please provide each
improvement you make as a single bullet point.
```
improvement: year-2022/day-01

# please explain what improvements you made
- improvement-1
- improvement-2
- improvement-3
- etc.
```
