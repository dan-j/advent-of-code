from __future__ import annotations

import functools
import itertools
from datetime import datetime
from pathlib import Path
from typing import List, Iterator


def parse_input(input: str) -> List[int]:
    return [int(line.split(": ")[1]) for line in input.splitlines()]


def move_player(pos: int, move: int) -> int:
    return (pos + move - 1) % 10 + 1


def next_move(die: Iterator[int], num: int) -> int:
    return sum([((next(die) - 1) % 100 + 1) for _ in range(num)])


def play_deterministic_game(p1_pos: int, p2_pos: int) -> int:
    p1_score, p2_score = 0, 0
    limit = 1000
    die = itertools.count(1)

    while True:
        p1_pos = move_player(p1_pos, next_move(die, 3))
        p1_score += p1_pos
        if p1_score >= limit:
            break

        p2_pos = move_player(p2_pos, next_move(die, 3))
        p2_score += p2_pos
        if p2_score >= limit:
            break

    return min(p1_score, p2_score) * (next(die) - 1)


@functools.cache
def play_dirac_game(player1_pos: int, player1_score: int, player2_pos: int, player2_score: int) -> (
        int, int
):
    """
    Uses recursion and caching to play through all combinations of dice-rolls. The x3 rolls are combined into a total
    move count and the quantity of outcomes which produce that move... i.e. 3 and 9 can only happen once, by rolling
    1 or 3 3-times respectively. The other totals can occur 3, 6 or 7 times.
    """
    p1_wins, p2_wins = 0, 0
    limit = 21
    for p1_move, p1_qty in [(3, 1), (4, 3), (5, 6), (6, 7), (7, 6), (8, 3), (9, 1)]:
        next_player1_pos = move_player(player1_pos, p1_move)
        next_p1_score = player1_score + next_player1_pos

        if next_p1_score >= limit:
            p1_wins += p1_qty
            continue

        for p2_move, p2_qty in [(3, 1), (4, 3), (5, 6), (6, 7), (7, 6), (8, 3), (9, 1)]:
            next_player2_pos = move_player(player2_pos, p2_move)
            next_p2_score = player2_score + next_player2_pos

            if next_p2_score >= limit:
                p2_wins += p2_qty * p1_qty
                continue

            p1_win, p2_win = play_dirac_game(next_player1_pos, next_p1_score, next_player2_pos, next_p2_score)
            p1_wins += p1_win * p1_qty * p2_qty
            p2_wins += p2_win * p1_qty * p2_qty

    return p1_wins, p2_wins


def part1(input: str) -> int:
    [p1, p2] = parse_input(input)
    return play_deterministic_game(p1, p2)


def part2(input: str) -> (int, int):
    [p1, p2] = parse_input(input)
    return play_dirac_game(p1, 0, p2, 0)


if __name__ == '__main__':
    example = Path('./example_input.txt').read_text()
    input = Path('./input.txt').read_text()

    for idx, (want_example, want_answer, fn) in enumerate(
            [(739785, 929625, part1),
             ((444356092776315, 341960390180808), (153087536629019, 175731756652760), part2)]
    ):
        start = datetime.now()
        got_example = fn(example)
        assert got_example == want_example, f"example: got {got_example}, want {want_example}"

        print(f"Part{idx + 1}(example): {got_example} ({datetime.now() - start})")
        start = datetime.now()
        got_answer = fn(input)
        assert got_answer == want_answer, f"answer: got {got_answer}, want {want_answer}"

        print(f"Part{idx + 1}(real): {got_answer} ({datetime.now() - start})")
        print("done!")
