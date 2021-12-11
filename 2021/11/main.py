from __future__ import annotations

from pathlib import Path
from typing import List, Set


def increment_grid(g: List[List[int]]) -> List[List[int]]:
    for i in range(0, len(g)):
        for j in range(0, len(g[i])):
            g[i][j] += 1

    return g


def reset_grid(g: List[List[int]]) -> List[List[int]]:
    for i in range(0, len(g)):
        for j in range(0, len(g[i])):
            if g[i][j] > 9:
                g[i][j] = 0

    return g


def calc_to_flash(g: List[List[int]]) -> Set[(int, int)]:
    f: Set[(int, int)] = set()
    for i in range(0, len(g)):
        for j in range(0, len(g[i])):
            if g[i][j] > 9:
                f.add((i, j))

    return f


def calc_next_coords(g: List[List[int]], coords: (int, int), move: (int, int)) -> bool:
    (i, j) = (coords[0] + move[0], coords[1] + move[1])
    if 0 <= i < len(g) and 0 <= j < len(g[coords[0]]):
        return (i, j)

    return None

def part1(input: str) -> int:
    grid: List[List[int]] = [list(map(int, row)) for row in input.splitlines()]

    # up, down, left, right, up-left, up-right, down-left, down-right
    directions = [(-1, 0), (1, 0), (0, -1), (0, 1), (-1, -1), (-1, 1), (1, -1), (1, 1)]

    num_flashes = 0
    for _ in range(0, 100):
        grid = increment_grid(grid)

        to_flash_coords: Set[(int, int)] = calc_to_flash(grid)
        flashed: Set[(int, int)] = set()
        while to_flash_coords:
            next_to_flash: Set[(int, int)] = set()
            for coords in to_flash_coords:
                num_flashes += 1
                flashed.add(coords)

                for move in directions:
                    next_coords = calc_next_coords(grid, coords, move)
                    if next_coords is None:
                        continue

                    (i, j) = next_coords
                    grid[i][j] += 1

                    if grid[i][j] > 9:
                        next_to_flash.add((i, j))

            to_flash_coords = next_to_flash - flashed
        grid = reset_grid(grid)

    return num_flashes


def part2(input: str) -> int:
    grid: List[List[int]] = [list(map(int, row)) for row in input.splitlines()]

    # up, down, left, right, up-left, up-right, down-left, down-right
    directions = [(-1, 0), (1, 0), (0, -1), (0, 1), (-1, -1), (-1, 1), (1, -1), (1, 1)]

    step = 0
    while True:
        step += 1
        grid = increment_grid(grid)

        to_flash_coords: Set[(int, int)] = calc_to_flash(grid)
        flashed: Set[(int, int)] = set()
        while to_flash_coords:
            next_to_flash: Set[(int, int)] = set()
            for coords in to_flash_coords:
                flashed.add(coords)

                for move in directions:
                    next_coords = calc_next_coords(grid, coords, move)
                    if next_coords is None:
                        continue

                    (i, j) = next_coords
                    grid[i][j] += 1

                    if grid[i][j] > 9:
                        next_to_flash.add((i, j))

            to_flash_coords = next_to_flash - flashed

        if len(flashed) == len(grid) * len(grid[0]):
            return step

        grid = reset_grid(grid)

    return 0


if __name__ == '__main__':
    example = "\n".join([
        "5483143223",
        "2745854711",
        "5264556173",
        "6141336146",
        "6357385478",
        "4167524645",
        "2176841721",
        "6882881134",
        "4846848554",
        "5283751526",
    ])

    input = Path('./input.txt').read_text()

    for want_example, want_answer, fn in ((1656, 1755, part1), (195, 212, part2)):
        got_example = fn(example)
        assert got_example == want_example, f"example: got {got_example}, want {want_example}"

        got_answer = fn(input)
        print(got_answer)
        assert got_answer == want_answer, f"answer: got {got_answer}, want {want_answer}"
