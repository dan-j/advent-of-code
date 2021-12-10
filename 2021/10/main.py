from __future__ import annotations

from pathlib import Path
from typing import List


def part1(input: str) -> int:
    table = dict({')': 3, ']': 57, '}': 1197, '>': 25137})
    opener = ['(', '[', '{', '<']
    closer = [')', ']', '}', '>']

    def find_illegal(line: str) -> str:
        stack = []
        for i in range(0, len(line)):
            char = line[i]
            if char in opener:
                stack.append(char)
                continue

            if char not in closer:
                return char

            top = stack[-1]
            if char != closer[opener.index(top)]:
                return char

            stack.pop()

        return ''

    score = 0
    for line in input.splitlines():
        illegal = find_illegal(line)
        if illegal != '':
            score += table[illegal]

    return score


def part2(input: str) -> int:
    table = dict({')': 1, ']': 2, '}': 3, '>': 4})
    opener = ['(', '[', '{', '<']
    closer = [')', ']', '}', '>']

    def calc_score(line: str) -> int:
        stack = []
        for i in range(0, len(line)):
            char = line[i]
            if char in opener:
                stack.append(char)
                continue

            if char not in closer:
                return 0

            top = stack[-1]
            if char != closer[opener.index(top)]:
                return 0

            stack.pop()

        if len(stack) == 0:
            return 0

        score = 0
        for item in reversed(stack):
            close_char = closer[opener.index(item)]
            score = score * 5 + table[close_char]

        return score

    scores: List[int] = []
    for line in input.splitlines():
        score = calc_score(line)
        if score != 0:
            scores.append(score)

    scores.sort()

    return scores[int(len(scores)/2)]


if __name__ == '__main__':
    example = '\n'.join([
        "[({(<(())[]>[[{[]{<()<>>",
        "[(()[<>])]({[<{<<[]>>(",
        "{([(<{}[<>[]}>{[]{[(<()>",
        "(((({<>}<{<{<>}{[]{[]{}",
        "[[<[([]))<([[{}[[()]]]",
        "[{[{({}]{}}([{[{{{}}([]",
        "{<[[]]>}<{[{[{[]{()[[[]",
        "[<(<(<(<{}))><([]([]()",
        "<{([([[(<>()){}]>(<<{{",
        "<{([{{}}[<[[[<>{}]]]>[]]",
    ])

    input = Path('./input.txt').read_text()

    for want_example, want_answer, fn in ((26397, 364389, part1), (288957, 2870201088, part2)):
        got_example = fn(example)
        assert got_example == want_example, f"example: got {got_example}, want {want_example}"

        got_answer = fn(input)
        print(got_answer)
        assert got_answer == want_answer, f"answer: got {got_answer}, want {want_answer}"
