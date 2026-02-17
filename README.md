# Queens Puzzle Solver

A brute force-based solver for the Queens puzzle (inspired by the LinkedIn Queens game). Given an N x N grid divided into colored regions, the solver places exactly one queen per region such that no two queens share the same row, column, or are adjacent.  

---

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Dependencies](#dependencies)
- [Installation](#installation)
- [How to Run](#how-to-run)
  - [CLI Mode](#cli-mode)
  - [GUI Mode](#gui-mode)
- [Input Format](#input-format)
  - [Text File (.txt)](#text-file-txt)
  - [Image File (.png / .jpg)](#image-file-png--jpg)
- [Algorithms](#algorithms)
  - [Pure Bruteforce](#pure-bruteforce)
  - [Optimized Bruteforce (Smallest Region First)](#optimized-bruteforce-smallest-region-first)
- [Output](#output)
- [Test Files](#test-files)

---

## Features

- Two solver algorithms: Pure Bruteforce and Optimized Bruteforce (MRV heuristic)
- Two input modes: text file (.txt) and image file (.png/.jpg) with automatic grid detection
- Two interface modes: CLI (Command Line Interface) and GUI (Graphical User Interface)
- GUI step-by-step visualization with adjustable speed delay
- Auto cell-size detection for image inputs using border stretch analysis
- Save solutions to .txt file
- Supports grid sizes from 4x4 up to 26x26 (regions A-Z)

---

## Project Structure

```
Tucil1/
├── go.mod                          # Go module definition
├── go.sum                          # Dependency checksums
├── README.md
├── data/                           # Test input files
│   ├── test.txt
│   ├── 4x4.txt ... 26x26.txt
│   ├── test1.png
│   ├── test2.png
│   └── test3.png
└── src/                            # Source code
    ├── main.go                     # Entry point (CLI + GUI mode selection)
    └── packages/
        ├── bruteforce/             # Pure bruteforce solver
        │   ├── Solve.go
        │   ├── Validation.go
        │   └── Region.go
        ├── bruteforce-optimized/   # Optimized solver (Smallest Region First)
        │   └── SmallestRegion.go
        ├── gui/                    # Fyne-based GUI
        │   ├── app.go
        │   ├── menu.go
        │   ├── fileinput.go
        │   ├── gridview.go
        │   ├── solver.go
        │   └── savedialog.go
        ├── imageprocessor/         # Image-to-grid converter
        │   └── ImageReader.go
        ├── output/                 # File output (save solution)
        │   └── FileIO.go
        └── utils/                  # CLI menu utilities
            └── Menu.go
```

---

## Prerequisites

- **Go** version 1.24 or later — [https://go.dev/dl/](https://go.dev/dl/)
- **C compiler** (required by Fyne for GUI mode)

> **Note:** If you only intend to use CLI mode, you still need a C compiler installed because the Fyne dependency is compiled regardless. Ensure `gcc` is available in your system PATH.

---

## Dependencies

| Dependency | Version | Description |
|---|---|---|
| [fyne.io/fyne/v2](https://fyne.io/) | v2.7.2 | Cross-platform GUI toolkit for Go |

All dependencies are managed via Go modules and will be downloaded automatically during the build step.

---

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd Tucil1
```

2. Download dependencies:

```bash
go get .
go mod tidy
```

---

## How to Run

From the **project root directory** (`Tucil1/`), run:

```bash
go run ./src/
```

You will be prompted to choose between CLI Mode and GUI Mode.

### CLI Mode

1. Select `1` for CLI Mode.
2. Choose an input mode:
   - `1` — Text File (.txt)
   - `2` — Image File (.png/.jpg)
   - `3` — Exit Program
3. Enter the filename (e.g., `5x5.txt` or `test1.png`). The program automatically reads from the `data/` directory.
4. Choose a solver algorithm:
   - `1` — Pure Bruteforce (recommended for grids up to 8x8)
   - `2` — Optimized Bruteforce (recommended for larger grids)
   - `3` — Back to input selection
5. After a solution is found, you will be asked whether to save the result to a `.txt` file.

### GUI Mode

1. Select `2` for GUI Mode.
2. The application window will open with a main menu.
3. Choose **Text File** or **Image File**.
4. Enter the filename and click **Load File**.
5. The grid will be displayed with colored regions.
6. Choose a solver:
   - **Pure Bruteforce** — exhaustive search
   - **Optimized Bruteforce** — uses Smallest Region First heuristic
7. Use the **Delay slider** (0–500 ms) to control visualization speed.
8. Use **Stop** to halt the solver, **Reset** to clear the grid, or **Save Solution** to export to a `.txt` file.
9. Click **Back to Menu** to load a different puzzle.

---

## Input Format

### Text File (.txt)

A plain text file where each character represents a region label (uppercase letter A–Z). Each row of the grid is one line. No spaces between characters.

Example (`5x5.txt`):

```
AABBB
AACCC
DDCCE
DDDEE
DDDEE
```

- Each unique letter represents a distinct region.
- The grid must be rectangular (all rows have equal length).
- Maximum 26 regions (A through Z).

### Image File (.png / .jpg)

A screenshot or image of a Queens puzzle grid. The program automatically:

1. Detects cell boundaries by scanning for dark border lines.
2. Determines cell size using the median distance between border centers.
3. Samples the center pixel of each cell to identify region colors.
4. Maps similar colors to the same region letter.

For best results:
- Use a clean screenshot with clearly visible dark borders between cells.
- Avoid images with text overlays, shadows, or rounded corners that obscure the grid.

---

## Algorithms

### Pure Bruteforce

Generates all possible combinations of N positions on an N x N grid (where N is the number of regions) and checks each combination against the constraints:

- No two queens in the same row
- No two queens in the same column
- No two queens in the same region
- No two queens adjacent to each other (king's move — includes diagonals)

**Time complexity:** $O\binom{N^2}{N}$ — exponential. Practical for grids up to approximately 8x8.

### Optimized Bruteforce (Smallest Region First)

A backtracking algorithm enhanced with the **Minimum Remaining Values (MRV) heuristic**:

1. At each step, select the unsolved region with the fewest valid candidate cells.
2. Try placing a queen in each valid cell of that region.
3. If a dead end is detected (any unsolved region has zero valid cells), prune that branch immediately.
4. Backtrack if no valid placement leads to a solution.

This approach drastically reduces the search space by:
- Focusing on the most constrained region first (fail-fast strategy).
- Pruning entire subtrees when a region becomes unsolvable.

**Performance:** Solves grids up to 26x26 in under 1 second for most cases.

---

## Output

When saving a solution, the program writes a `.txt` file where:
- Region letters are preserved as-is for cells without a queen.
- Cells containing a queen are marked with `#`.

Example output:

```
A # B B B
A A C C #
D # C C E
D D D # E
# D D E E
```

---

## Test Files

The `data/` directory includes the following test cases:

| File | Grid Size | Type |
|---|---|---|
| test.txt | 4x4 | Text |
| 4x4.txt | 4x4 | Text |
| 5x5.txt | 5x5 | Text |
| 6x6.txt | 6x6 | Text |
| 7x7.txt | 7x7 | Text |
| 9x9.txt | 9x9 | Text |
| 12x12.txt | 12x12 | Text |
| 15x15.txt | 15x15 | Text |
| 20x20.txt | 20x20 | Text |
| 23x23.txt | 23x23 | Text |
| 26x26.txt | 26x26 | Text |
| test1.png | Varies | Image |
| test2.png | Varies | Image |
| test3.png | Varies | Image |

---

## Author
Athilla Zaidan Zidna Fann 13524068

IF2211 Algorithm Strategy — Tugas Kecil 1
Institut Teknologi Bandung
