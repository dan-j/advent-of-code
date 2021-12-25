from __future__ import annotations

from datetime import datetime
from pathlib import Path
from typing import Set, Tuple

Dimensions = Tuple[int, int]
Position = Tuple[int, int]
State = (Set[Position], Set[Position])


def parse_input(input: str) -> (Dimensions, State):
    east: Set[Position] = set()
    south: Set[Position] = set()
    lines = input.splitlines()
    for i, line in enumerate(lines):
        for j, c in enumerate(line):
            if c == ">":
                east.add((i, j))
            elif c == "v":
                south.add((i, j))

    return (len(lines) - 1, len(lines[0]) - 1), (east, south)


def move_cucumbers(dimensions: Dimensions, state: State) -> (State, bool):
    max_i, max_j = dimensions
    east, south = state
    next_east: Set[Position] = set()
    next_south: Set[Position] = set()
    all_cucumbers = east.union(south)
    moved = False
    for i, j in east:
        next_j = (j + 1) % (max_j + 1)
        if (i, next_j) not in all_cucumbers:
            moved |= True
            next_east.add((i, next_j))
        else:
            next_east.add((i, j))

    all_cucumbers = next_east.union(south)
    for i, j in south:
        next_i = (i + 1) % (max_i + 1)
        if (next_i, j) not in all_cucumbers:
            moved |= True
            next_south.add((next_i, j))
        else:
            next_south.add((i, j))

    return (next_east, next_south), moved


def print_cucumbers(dimensions: Dimensions, state: State):
    max_i, max_j = dimensions
    east, south = state
    for i in range(max_i + 1):
        for j in range(max_j + 1):
            if (i, j) in east:
                print(">", end="")
            elif (i, j) in south:
                print("v", end="")
            else:
                print(".", end="")
        print("")


def part1(input: str) -> int:
    dimensions, state = parse_input(input)
    n = 0
    moved = True
    while moved:
        state, moved = move_cucumbers(dimensions, state)
        n += 1
        # print("step: ", n)
        # print_cucumbers(dimensions, state)

    print_cucumbers(dimensions, state)

    return n


if __name__ == '__main__':
    example = Path('./example_input.txt').read_text()
    input = Path('./input.txt').read_text()
    start = datetime.now()
    want_example = 58
    got_example = part1(example)
    assert got_example == want_example, f"example: got {got_example}, want {want_example}"

    print(f"Part1(example): {got_example} ({datetime.now() - start})")
    start = datetime.now()
    want_answer = 523
    got_answer = part1(input)
    assert got_answer == want_answer, f"answer: got {got_answer}, want {want_answer}"

    print(f"Part1(real): {got_answer} ({datetime.now() - start})")
    print("done!")
