from __future__ import annotations

from pathlib import Path
from typing import List, Set


def read_input(input: str) -> (List[List[bool]], List[(str, int)]):
    blank_index = 0
    lines = input.splitlines()
    for idx, line in enumerate(lines):
        if line == "":
            blank_index = idx
            break

    return read_grid(lines[:blank_index]), read_folds(lines[blank_index+1:])


def read_grid(lines: List[str]) -> List[List[bool]]:
    coords: List[(int, int)] = [tuple(map(int, line.split(","))) for line in lines]

    max_x = max(coords, key=lambda x: x[0])[0]
    max_y = max(coords, key=lambda y: y[1])[1]

    return [[(x, y) in coords for y in range(0, max_y+1)] for x in range(0, max_x+1)]


def read_folds(lines: List[str]) -> List[(str, int)]:
    return [(f[0], int(f[1])) for f in (line[len("fold along "):].split("=") for line in lines)]


def print_grid(grid: List[List[bool]]):
    print()
    for y in range(0, len(grid[0])):
        for x in range(0, len(grid)):
            print("X" if grid[x][y] else ".", end="")
        print()


def fold_grid(grid: List[List[bool]], folds: List[(str, int)]) -> List[List[bool]]:
    for fold in folds:
        if fold[0] == "y":
            # how many elements are there after the fold?
            bottom_len = len(grid[0]) - fold[1] - 1
            # how many elements are there before the fold?
            top_len = fold[1]
            # folds may not be centered, so range from before the fold, decrementing by the minimum number of elements
            # on either side of the fold
            for y in range(fold[1]-1, fold[1] - min(bottom_len, top_len) - 1, -1):
                for x in range(len(grid)):
                    mirror_y = fold[1]*2 - y
                    grid[x][y] = grid[x][y] or grid[x][mirror_y]

            # remove the bottom half of the fold
            grid = [col[0:fold[1]] for col in grid]
        elif fold[0] == "x":
            # similar logic as above but for vertical folds instead of horizontal
            right_len = len(grid) - fold[1] - 1
            left_len = fold[1]
            for y in range(len(grid[0])):
                for x in range(fold[1]-1, fold[1] - min(left_len, right_len) - 1, -1):
                    mirror_x = fold[1] * 2 - x
                    grid[x][y] = grid[x][y] or grid[mirror_x][y]

            # remove the right-hand side of the fold
            grid = grid[:fold[1]]

        # print_grid(grid)

    return grid


def part1(input: str) -> int:
    (grid, folds) = read_input(input)

    grid = fold_grid(grid, folds[:1])

    return sum(col.count(True) for col in grid)


def part2(input: str) -> int:
    (grid, folds) = read_input(input)

    grid = fold_grid(grid, folds)

    print_grid(grid)

    return 0


if __name__ == '__main__':
    example = Path('./example_input.txt').read_text()
    input = Path('./input.txt').read_text()

    for want_example, want_answer, fn in ((17, 693, part1), (0, 0, part2)):
        got_example = fn(example)
        assert got_example == want_example, f"example: got {got_example}, want {want_example}"

        got_answer = fn(input)
        assert got_answer == want_answer, f"answer: got {got_answer}, want {want_answer}"

        print("done!")
