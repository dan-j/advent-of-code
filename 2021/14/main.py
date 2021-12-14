from __future__ import annotations

import operator

from pathlib import Path
from typing import Dict, List
from datetime import datetime


def read_input(input: str) -> (str, Dict[str, str]):
    lines = input.splitlines()
    return lines[0], dict([rule.split(" -> ") for rule in lines[2:]])


def expand_template(template: str, rules: Dict[str, str], iterations: int) -> (Dict[str, int], str):
    for iteration in range(iterations):
        next_template = ""
        for i in range(len(template)-1):
            word = template[i:i+2]
            insert_char = rules[word]
            next_template += word[0] + insert_char

        template = next_template + template[-1]

    counts = dict([(c, template.count(c)) for c in set(template)])

    return counts, template[-1]


def part1(input: str) -> int:
    template, rules = read_input(input)

    # counts = expand_template(template, rules, 10)
    counts = {}
    for i in range(len(template)-1):
        curr_counts, last_char = expand_template(template[i:i+2], rules, 10)
        if i < len(template) - 2:
            curr_counts[last_char] -= 1

        for key in curr_counts:
            if key in counts:
                counts[key] += curr_counts[key]
            else:
                counts[key] = curr_counts[key]

    return max(counts.items(), key=operator.itemgetter(1))[1] - min(counts.items(), key=operator.itemgetter(1))[1]


def iterate_pair(pair: str, rules: Dict[str, str], iterations: int) -> Dict[str, int]:
    curr_pairs = {pair: 1}
    for i in range(iterations):
        next_pairs = dict(curr_pairs)
        remove_pairs = {}
        for curr_pair in curr_pairs:
            remove_pairs[curr_pair] = curr_pairs[curr_pair]
            left_pair = curr_pair[0] + rules[curr_pair]
            right_pair = rules[curr_pair] + curr_pair[1]

            if left_pair in next_pairs:
                next_pairs[left_pair] += curr_pairs[curr_pair]
            else:
                next_pairs[left_pair] = curr_pairs[curr_pair]

            if right_pair in next_pairs:
                next_pairs[right_pair] += curr_pairs[curr_pair]
            else:
                next_pairs[right_pair] = curr_pairs[curr_pair]

            pass

        for rm in remove_pairs:
            next_pairs[rm] -= remove_pairs[rm]

        curr_pairs = dict(filter(lambda e: e[1] > 0, next_pairs.items()))
        # pairs = dict([(p, pairs[p] - curr_pairs[p]) for p in curr_pairs])
        pass

    return curr_pairs


def part2(input: str) -> int:
    template, rules = read_input(input)

    pairs = []
    for i in range(len(template) - 1):
        pairs.append(template[i:i+2])

    pair_counts = {}
    for pair in pairs:
        res = iterate_pair(pair, rules, 40)
        for p in res:
            if p in pair_counts:
                pair_counts[p] += res[p]
            else:
                pair_counts[p] = res[p]

    counts = dict([(k, 0) for k in set(t[1] for t in rules.items())])
    for pair in pair_counts:
        counts[pair[0]] += pair_counts[pair]

    counts[template[-1]] += 1

    return max(counts.items(), key=operator.itemgetter(1))[1] - min(counts.items(), key=operator.itemgetter(1))[1]


if __name__ == '__main__':
    example = Path('./example_input.txt').read_text()
    input = Path('./input.txt').read_text()

    for idx, (want_example, want_answer, fn) in enumerate([(1588, 2509, part1), (2188189693529, 2827627697643, part2)]):
        start = datetime.now()
        got_example = fn(example)
        assert got_example == want_example, f"example: got {got_example}, want {want_example}"

        print(f"Part{idx+1} (example): {datetime.now() - start}")
        start = datetime.now()
        got_answer = fn(input)
        assert got_answer == want_answer, f"answer: got {got_answer}, want {want_answer}"

        print(f"Part{idx+1}: {datetime.now() - start}")
        print("done!")
