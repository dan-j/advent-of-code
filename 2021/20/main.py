from __future__ import annotations

from datetime import datetime
from pathlib import Path
from typing import Set

Algorithm = str
Pixel = (int, int)
Bounds = ((int, int), (int, int))


def parse_input(input: str) -> (Algorithm, Set[Pixel], Bounds):
    lines = input.splitlines()

    image = set()
    for i, line in enumerate(lines[2:]):
        for j, c in enumerate(line):
            if c == "#":
                image.add((i, j))

    return lines[0], image, ((0, len(lines[2:])), (0, len(lines[-1])))


def expand_bounds(bounds: Bounds, iterations: int) -> Bounds:
    min_padding = iterations * 3
    max_padding = iterations * 3 - 1
    (min_i, max_i), (min_j, max_j) = bounds

    return (min_i - min_padding, max_i + max_padding), (min_j - min_padding, max_j + max_padding)


def calc_window_value(centre: (int, int), image: Set[Pixel]) -> int:
    value = 0
    ci, cj = centre
    for i in range(ci - 1, ci + 2):
        for j in range(cj - 1, cj + 2):
            value = value << 1 | (1 if (i, j) in image else 0)

    return value


def print_image(image: Set[Pixel], bounds: Bounds):
    print()
    (min_i, max_i), (min_j, max_j) = bounds
    for i in range(min_i, max_i + 1):
        for j in range(min_j, max_j + 1):
            print("#" if (i, j) in image else ".", end="")
        print()


def enhance_image(algo: Algorithm, image: Set[Pixel], bounds: Bounds) -> (Set[Pixel], Bounds):
    next_image: Set[Pixel] = set()
    (min_i, max_i), (min_j, max_j) = bounds
    next_bounds = ((min_i + 2, max_i - 2), (min_j + 2, max_j - 2))
    (min_i, max_i), (min_j, max_j) = next_bounds

    for i in range(min_i, max_i + 1):
        for j in range(min_j, max_j + 1):
            value = calc_window_value((i, j), image)
            if algo[value] == '#':
                next_image.add((i, j))

    return next_image, next_bounds


def iterate_enhancement(algo: Algorithm, image: Set[Pixel], bounds: Bounds, iterations: int) -> (Set[Pixel], Bounds):
    bounds = expand_bounds(bounds, iterations)
    for i in range(iterations):
        image, bounds = enhance_image(algo, image, bounds)

    return image, bounds


def part1(input: str) -> int:
    algo, image, bounds = parse_input(input)
    image, bounds = iterate_enhancement(algo, image, bounds, 2)

    print_image(image, bounds)

    return len(image)


def part2(input: str) -> int:
    algo, image, bounds = parse_input(input)
    image, bounds = iterate_enhancement(algo, image, bounds, 50)

    print_image(image, bounds)
    return len(image)


if __name__ == '__main__':
    example = Path('./example_input.txt').read_text()
    input = Path('./input.txt').read_text()

    for idx, (want_example, want_answer, fn) in enumerate([(35, 5597, part1), (3351, 18723, part2)]):
        start = datetime.now()
        got_example = fn(example)
        assert got_example == want_example, f"example: got {got_example}, want {want_example}"

        print(f"Part{idx + 1}(example): {got_example} ({datetime.now() - start})")
        start = datetime.now()
        got_answer = fn(input)
        assert got_answer == want_answer, f"answer: got {got_answer}, want {want_answer}"

        print(f"Part{idx + 1}(real): {got_answer} ({datetime.now() - start})")
        print("done!")
