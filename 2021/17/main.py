from __future__ import annotations

import math
from cmath import sqrt
from datetime import datetime
from pathlib import Path
from typing import List, Set


def parse_input(input: str) -> (int, int, int, int):
    split = input.strip().split(" ")
    [x1, x2] = split[2][:-1].split("=")[1].split("..")
    [y1, y2] = split[3].split("=")[1].split("..")

    return int(x1), int(x2), int(y1), int(y2)


def print_trench(steps: List[(int, int)], x1: int, x2: int, y1: int, y2: int):
    min_y = min(min(steps, key=lambda step: step[1])[1], y1)
    max_y = max(max(steps, key=lambda step: step[1])[1], y2)
    max_x = max(max(steps, key=lambda step: step[0])[0], x2)

    for y in range(max_y, min_y - 1, -1):
        for x in range(max_x + 1):
            char = "."

            if x1 <= x <= x2 and y1 <= y <= y2:
                char = "T"

            if (x, y) in steps:
                char = "S" if (x, y) == steps[0] else "#"

            print(char, end="")
        print()


def calc_step(velocity: (int, int), position: (int, int)) -> ((int, int), (int, int)):
    x = position[0] + velocity[0]
    y = position[1] + velocity[1]
    vx = velocity[0] + (-1 if velocity[0] > 0 else 1 if velocity[0] < 0 else 0)
    vy = velocity[1] - 1

    return (x, y), (vx, vy)


def calc_n_for_sum(sum: int) -> int:
    # calculate the minimum number of steps to reach a target "sum". this can be calculated by using the equation
    # n + ... + 4 + 3 + 2 + 1 = n * (n + 1) / 2
    #
    # use the equation for solving quadratic equations
    # solve for x = n*(n+1)/2
    # -->       0 = n**2 + n - 2x"
    d = 1 - (4 * (-(sum * 2)))

    sol1 = (-1 - sqrt(d).real) / 2
    sol2 = (-1 + sqrt(d).real) / 2
    return math.ceil(max(sol1, sol2))


def launch_probe(velocity: (int, int), x1: int, x2: int, y1: int, y2: int) -> (List[(int, int)], bool):
    x, y = (0, 0)
    steps = [(x, y)]
    hits_target = False
    while x <= x2 and y >= y1:
        (x, y), velocity = calc_step(velocity, steps[-1])
        steps.append((x, y))
        # print_trench(steps, x1, x2, y1, y2)

        if x1 <= x <= x2 and y1 <= y <= y2:
            hits_target = True
            break

    return steps, hits_target


def part1(input: str) -> int:
    (x1, x2, y1, y2) = parse_input(input)

    max_ys: List[int] = []
    start_vx = calc_n_for_sum(x1)
    for vx in range(start_vx, x2 + 1):
        for vy in range(1, y1 * -1):
            steps, hits_target = launch_probe((vx, vy), x1, x2, y1, y2)

            if hits_target:
                max_ys.append(max(steps, key=lambda step: step[1])[1])

    return max(max_ys)


def part2(input: str) -> int:
    (x1, x2, y1, y2) = parse_input(input)

    valid_velocities: Set[(int, int)] = set()
    start_vx = calc_n_for_sum(x1)

    for vx in range(start_vx, x2 + 1):
        for vy in range(y1, y1 * -1):
            steps, hits_target = launch_probe((vx, vy), x1, x2, y1, y2)

            if hits_target:
                valid_velocities.add((vx, vy))

    return len(valid_velocities)


if __name__ == '__main__':
    example = Path('./example_input.txt').read_text()
    input = Path('./input.txt').read_text()

    for idx, (want_example, want_answer, fn) in enumerate([(45, 6555, part1), (112, 4973, part2)]):
        start = datetime.now()
        got_example = fn(example)
        assert got_example == want_example, f"example: got {got_example}, want {want_example}"

        print(f"Part{idx + 1} (example): {datetime.now() - start}")
        start = datetime.now()
        got_answer = fn(input)
        assert got_answer == want_answer, f"answer: got {got_answer}, want {want_answer}"

        print(f"Part{idx + 1}: {datetime.now() - start}")
        print("done!")
