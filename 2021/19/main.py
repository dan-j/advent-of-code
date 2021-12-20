from __future__ import annotations

from datetime import datetime
from typing import List, Set, Tuple, Optional

import numpy as np

Position = Tuple[int, int, int]

ROTATIONS = {
    'x': np.matrix([[1, 0, 0], [0, 0, -1], [0, 1, 0]], dtype=np.int16),
    'y': np.matrix([[0, 0, -1], [0, 1, 0], [1, 0, 0]], dtype=np.int16),
    'z': np.matrix([[0, -1, 0], [1, 0, 0], [0, 0, 1]], dtype=np.int16),
}

# These rotations provide initial 6 directions the scanner can face, you then
# rotate through 'x' 4 times to get all 24 permutations
OUTER_ROTATIONS = [
    np.identity(3, dtype=np.int16),
    ROTATIONS['z'],
    ROTATIONS['z'],
    ROTATIONS['z'],
    ROTATIONS['y'],
    np.matmul(ROTATIONS['y'], ROTATIONS['y']),
]


def read_input(path: str) -> List[Set[Position]]:
    scanners: List[Set[Position]] = []
    with open(path, 'r') as f:
        # read --- scanner N --- header
        line = f.readline()
        coords: Set[(int, int, int)] = set()
        while line:
            line = f.readline().rstrip()
            if line == "":
                line = f.readline()  # consume the --- scanner N --- header
                scanners.append(coords)
                coords = set()
                continue

            coords.add(tuple(map(int, line.split(","))))

    return scanners


def find_scanner_correction(base: Set[Position], beacons: Set[Position]) -> Optional[(Position, Set[Position])]:
    beacons_array = np.array(list(beacons), dtype=np.int16)
    base_array = np.array(list(base), dtype=np.int16)
    for outer_rotation in OUTER_ROTATIONS:
        beacons_array = np.matmul(beacons_array, outer_rotation)
        for _ in range(4):
            beacons_array = np.array(np.matmul(beacons_array, ROTATIONS['x']))
            for ref in base_array:
                deltas = np.subtract(ref, beacons_array)
                for delta in deltas:
                    s2_relative_to_base = set(tuple(b) for b in np.add(beacons_array, delta))

                    intersection = s2_relative_to_base.intersection(base)
                    if len(intersection) >= 12:
                        # print("\t\tFound it!")
                        return tuple(delta), s2_relative_to_base.union(base)
    return None


def part1(input: str) -> (int, List[Position]):
    scanners = read_input(input)

    positions: List[Optional[Position]] = [None] * len(scanners)
    positions[0] = (0, 0, 0)
    all_beacons = scanners[0]
    while any([p is None for p in positions]):
        print(positions)
        for next_idx, next_scanner in enumerate(scanners):
            if positions[next_idx] is not None:
                continue

            result = find_scanner_correction(all_beacons, next_scanner)
            if result is not None:
                print("Found scanner", next_idx, result[0])
                positions[next_idx], all_beacons = result

    return len(all_beacons), positions


def part2(scanners: List[Position]) -> int:
    return max(
        sum(d) for d in [
            list(map(abs, np.subtract(s1, s2))) for s1 in scanners for s2 in scanners
        ]
    )


def main():
    example = "./example_input.txt"
    input = "./input.txt"

    full_start, start = datetime.now(), datetime.now()
    got_example, example_scanners = part1(example)
    want_example1 = 79
    assert got_example == want_example1, f"example: got {got_example}, want {want_example1}"

    print(f"Part1(example): {got_example} {datetime.now() - start}")

    start = datetime.now()
    got_answer, scanners = part1(input)
    answer1 = 491
    assert got_answer == answer1, f"answer: got {got_answer}, want {answer1}"

    print(f"Part1(real): {got_answer} ({datetime.now() - start})")

    start = datetime.now()
    got_example = part2(example_scanners)
    want_example1 = 3621
    assert got_example == want_example1, f"example: got {got_example}, want {want_example1}"

    print(f"Part2(example): {got_example} {datetime.now() - start}")

    start = datetime.now()
    got_answer = part2(scanners)
    answer1 = 13374
    assert got_answer == answer1, f"answer: got {got_answer}, want {answer1}"

    print(f"Part2(real): {got_answer} ({datetime.now() - start})")
    print(f"done! ({datetime.now() - full_start})")


if __name__ == '__main__':
    main()
