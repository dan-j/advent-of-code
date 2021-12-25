from __future__ import annotations

import sys
from datetime import datetime
from typing import Tuple, Callable, Iterator

"""
Proper pleased with this implementation. For different inputs, just take the 5th (div), 6th (add) and 16th (add) 
instruction parameter for each iteration of `inp`, and update `PARAMS` as required.

We reverse engineer the ALU to form the `calc_z` function. We then work backwards from the 14th digit to find valid 
ranges from an input z to produce `z=0` at the end. These ranges are stored in MIN_MAXES for later use.

We then brute-force all combinations of numbers, exiting early when a z value outside the allowed range is 
calculated. It's hardly brute-force though, both parts run in under 1s :D
"""

RangeFn = Callable[[], Iterator[int]]

PARAMS = [
    (1, 11, 1),
    (1, 11, 11),
    (1, 14, 1),
    (1, 11, 11),
    (26, -8, 2),
    (26, -5, 9),
    (1, 11, 7),
    (26, -13, 11),
    (1, 12, 6),
    (26, -1, 15),
    (1, 14, 7),
    (26, -5, 1),
    (26, -4, 8),
    (26, -8, 6),
]


def calc_z(w: int, z: int, p: Tuple[int]) -> int:
    """
    The ALU executes this function per digit from the model number. The tuple `p` is extracted from the input
    manually and are defined in the PARAMS global.
    """
    if (z % 26) + p[1] == w:
        return z // p[0]

    return z // p[0] * 26 + w + p[2]


def find_monad(model_number: str, range_fn: RangeFn, z: int) -> (int, int):
    n = len(model_number)
    for w in range_fn():
        next_z = calc_z(w, z, PARAMS[n])
        if n == 13:
            if next_z == 0:
                return int(model_number + str(w)), 0
            continue
        elif not (MIN_MAXES[n+1][0] <= next_z <= MIN_MAXES[n+1][1]):
            # it's impossible to generate a valid number from here since z is outside the allowed range
            continue

        result, next_z = find_monad(model_number + str(w), range_fn, next_z)
        if next_z == 0:
            return result, 0

    return int(model_number), -1


def init_min_maxes():
    """
    This function initialises a MIN_MAX global with the minimum and maximum values for z on each digit. This uses the
    knowledge of what `calc_z()` performs based on the PARAMS at that digit. It works backwards based on the knowledge
    that z must end in zero, hence z for the 2nd-to-last digit must be in a range
    """
    global MIN_MAXES
    min_z, max_z = 26, 0
    # brute force allowed z input for final digit
    for z in range(25):
        for w in range(9, 0, -1):
            if (calc_z(w, z, PARAMS[-1])) == 0:
                if z < min_z:
                    min_z = z
                elif z > max_z:
                    max_z = z

    MIN_MAXES = [(min_z, max_z)]
    for n in range(12, -1, -1):
        # if PARAMS[n][1] <= 9 then calc_z will always divide by 26, so candidate z's for the previous digit must be
        # within a range of min * 26 and max * 26. otherwise, calc_z will always at least multiply by 26, so we divide
        # by 26.
        next_min = (min_z * PARAMS[n][0]) if PARAMS[n][1] <= 9 else min_z // 26
        next_max = (max_z * PARAMS[n][0] + 1) if PARAMS[n][1] <= 9 else (max_z // 26 + 1)
        prev_min, prev_max = min_z, max_z
        min_z, max_z = sys.maxsize, 0
        for z in range(next_min, next_max):
            for w in range(9, 0, -1):
                if prev_min <= calc_z(w, z, PARAMS[n]) <= prev_max:
                    if z < min_z:
                        min_z = z
                    elif z > max_z:
                        max_z = z
        MIN_MAXES.insert(0, (min_z, max_z))


def part1() -> int:
    model_number, _ = find_monad("", lambda: range(9, 0, -1), 0)
    return model_number


def part2():
    model_number, _ = find_monad("", lambda: range(1, 10), 0)
    return model_number


if __name__ == '__main__':
    start = datetime.now()
    # YOU MUST INIT MIN_MAXES before running the parts
    init_min_maxes()

    for idx, (want_answer, fn) in enumerate([(92969593497992, part1), (81514171161381, part2)]):
        start_part = datetime.now()
        got_answer = fn()
        assert got_answer == want_answer, f"answer: got {got_answer}, want {want_answer}"

        print(f"Part{idx + 1}(real): {got_answer} ({datetime.now() - start_part})")

    print(f"done! ({datetime.now() - start})")
