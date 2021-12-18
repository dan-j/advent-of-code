from __future__ import annotations

import copy
from datetime import datetime
from pathlib import Path
from typing import List


def parse_input(input: str) -> (List[any]):
    return [eval(line) for line in input.splitlines()]


def traverse_expression(expr: List[any], path: List[int]) -> any:
    if len(path) == 0:
        return expr

    sub_expr = expr[path[0]]
    for path in path[1:]:
        sub_expr = sub_expr[path]

    return sub_expr


def explode_pair(expr: List[any], path: List[int]) -> List[any]:
    pair = traverse_expression(expr, path)

    left_path, right_path = list(path), list(path)
    done = False
    while not done:
        while left_path[-1] > 0:
            left_path[-1] -= 1
            element = traverse_expression(expr, left_path)

            if isinstance(element, list):
                left_path += [len(element)]
                continue

            if isinstance(element, int):
                left_parent = traverse_expression(expr, left_path[:-1])
                left_parent[left_path[-1]] = element + pair[0]
                done = True
                break

        if done or len(left_path) <= 1:
            break

        left_path = left_path[:-1]

    done = False
    while not done:
        while right_path[-1] < len(traverse_expression(expr, right_path[:-1])) - 1:
            right_path[-1] += 1
            element = traverse_expression(expr, right_path)

            if isinstance(element, list):
                right_path += [-1]
                continue

            if isinstance(element, int):
                right_parent = traverse_expression(expr, right_path[:-1])
                right_parent[right_path[-1]] = element + pair[1]
                done = True
                break

        if done or len(right_path) <= 1:
            break

        right_path = right_path[:-1]

    parent = traverse_expression(expr, path[:-1])
    parent[path[-1]] = 0

    return expr


def split_pair(expr: List[any], path: List[int]) -> List[any]:
    parent = traverse_expression(expr, path[:-1])
    number = parent[path[-1]]
    left = number // 2
    parent[path[-1]] = [left, number - left]

    return expr


def explode_expression(expr: List[any], depth: int, path: List[int]) -> bool:
    for i, item in enumerate(traverse_expression(expr, path)):
        next_path = path + [i]

        if depth >= 4 and isinstance(item, list):
            explode_pair(expr, next_path)
            return True

        if isinstance(item, list):
            if explode_expression(expr, depth + 1, next_path):
                return True

    return False


def split_expression(expr: List[any], path: List[int]) -> bool:
    for i, item in enumerate(traverse_expression(expr, path)):
        next_path = path + [i]

        if isinstance(item, int) and item >= 10:
            split_pair(expr, next_path)
            return True

        if isinstance(item, list):
            if split_expression(expr, next_path):
                return True

    return False


def reduce_expression(expr: List[any]) -> List[any]:
    expr = copy.deepcopy(expr)
    while explode_expression(expr, 1, []) or split_expression(expr, []):
        pass

    return expr


def add_expressions(left: List[any], right: List[any]) -> List[any]:
    return [left, right]


def calc_magnitude(expr: List[any]) -> int:
    left = calc_magnitude(expr[0]) if isinstance(expr[0], list) else expr[0]
    right = calc_magnitude(expr[1]) if isinstance(expr[1], list) else expr[1]

    return left * 3 + right * 2


def part1(input: str) -> int:
    expressions = parse_input(input)

    current_result: List[any] = expressions[0]
    for expr in expressions[1:]:
        current_result = reduce_expression(add_expressions(current_result, expr))

    return calc_magnitude(current_result)


def part2(input: str) -> int:
    expressions = parse_input(input)
    n = len(expressions)
    return max([
        calc_magnitude(reduce_expression(add_expressions(expressions[i], expressions[j]))) if i != j else 0
        for i in range(n) for j in range(n)
    ])


if __name__ == '__main__':
    example = Path('./example_input.txt').read_text()
    input = Path('./input.txt').read_text()

    for idx, (want_example, want_answer, fn) in enumerate([(4140, 3494, part1), (3993, 4712, part2)]):
        start = datetime.now()
        got_example = fn(example)
        assert got_example == want_example, f"example: got {got_example}, want {want_example}"

        print(f"Part{idx + 1} (example): {datetime.now() - start}")
        start = datetime.now()
        got_answer = fn(input)
        assert got_answer == want_answer, f"answer: got {got_answer}, want {want_answer}"

        print(f"Part{idx + 1}: {datetime.now() - start}")
        print("done!")
