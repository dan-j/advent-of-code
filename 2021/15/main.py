from __future__ import annotations

import heapq

from pathlib import Path
from typing import List
from sys import maxsize
from datetime import datetime


def read_input(input: str) -> List[List[int]]:
    lines = input.splitlines()
    grid = [[int(c) for c in line] for line in lines]
    return grid


def calc_neighbours(max_i: int, max_j: int, current: (int, int)) -> (int, int):
    directions = [(0, -1), (0, 1), (-1, 0), (1, 0)]
    neighbours: List[(int, int)] = []
    (ci, cj) = current
    for di, dj in directions:
        i, j = ci + di, cj + dj

        if 0 <= i < max_i and 0 <= j < max_j:
            neighbours.append((i, j))

    return neighbours


def dijkstra(grid: List[List[int]]) -> int:
    # queue = [(i, j) for j in range(len(grid)) for i in range(len(grid))]
    current = (0, 0)
    target = (len(grid)-1, len(grid[0])-1)

    distances = [[maxsize] * len(grid[i]) for i in range(len(grid))]
    distances[current[0]][current[1]] = 0

    queue: List[(int, (int, int))] = [(0, current)]

    visited = set()
    in_visited = lambda n: n not in visited

    while len(queue) > 0:
        current = heapq.heappop(queue)[1]
        current_distance = distances[current[0]][current[1]]

        neighbours = filter(in_visited, calc_neighbours(len(grid), len(grid[0]), current))
        for n in neighbours:
            if distances[n[0]][n[1]] > current_distance + grid[n[0]][n[1]]:
                distances[n[0]][n[1]] = current_distance + grid[n[0]][n[1]]
                heapq.heappush(queue, (distances[n[0]][n[1]], (n[0], n[1])))

        visited.add(current)

    return distances[-1][-1]


def part1(input: str) -> int:
    return dijkstra(read_input(input))


def part2(input: str) -> int:
    initial_grid = read_input(input)

    grid: List[List[int]] = [[]] * len(initial_grid)
    for x in range(5):
        for i, row in enumerate(initial_grid):
            grid[i] = grid[i] + list(map(lambda n: (n + x - 1) % 9 + 1, initial_grid[i]))

    initial_grid = list(grid)
    for x in range(1, 5):
        for row in initial_grid:
            grid += [[(n + x - 1) % 9 + 1 for n in row]]

    return dijkstra(grid)


if __name__ == '__main__':
    example = Path('./example_input.txt').read_text()
    input = Path('./input.txt').read_text()

    for idx, (want_example, want_answer, fn) in enumerate([(40, 604, part1), (315, 2907, part2)]):
        start = datetime.now()
        got_example = fn(example)
        assert got_example == want_example, f"example: got {got_example}, want {want_example}"

        print(f"Part{idx + 1} (example): {datetime.now() - start}")
        start = datetime.now()
        got_answer = fn(input)
        assert got_answer == want_answer, f"answer: got {got_answer}, want {want_answer}"

        print(f"Part{idx + 1}: {datetime.now() - start}")
        print("done!")
