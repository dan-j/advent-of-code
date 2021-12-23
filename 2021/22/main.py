from __future__ import annotations

from datetime import datetime
from pathlib import Path
from typing import List, Optional, Generator


class Range:
    def __init__(self, start: int, end: int):
        self.start = start
        self.end = end

    def get_length(self) -> int:
        return self.end - self.start + 1


class Cuboid:
    def __init__(self, x: Range, y: Range, z: Range, on: bool):
        self.on = on
        self.x, self.y, self.z = x, y, z

    def intersection(self, c: Cuboid) -> Optional[Cuboid]:
        xr = Range(max(self.x.start, c.x.start), min(self.x.end, c.x.end))
        yr = Range(max(self.y.start, c.y.start), min(self.y.end, c.y.end))
        zr = Range(max(self.z.start, c.z.start), min(self.z.end, c.z.end))

        if xr.start > xr.end or yr.start > yr.end or zr.start > zr.end:
            return None

        return Cuboid(xr, yr, zr, not self.on)

    def calc_volume(self) -> int:
        return self.x.get_length() * self.y.get_length() * self.z.get_length()


class Reactor:
    def __init__(self):
        self.cuboids: List[Cuboid] = []

    def apply_cuboid(self, c: Cuboid):
        # For each existing cuboid, find if this cuboid intersects it. If it does, we add it to the list as
        # a negation of what it intersects, and later if it's "on" we add the full cuboid to the list. This is
        # equivalent in set-theory as `A ∪ B = A + B - (A ∩ B)`.
        # If the cuboid is in the "off" position, it's intersection with any other cuboids will be negated,
        # then we just don't add the full cuboid at the end. Whether the cuboid is actually "on" or "off" doesn't
        # matter, because it's just negating an existing offset... I was surprised this last bit worked!
        for i in range(len(self.cuboids)):
            intersection = self.cuboids[i].intersection(c)
            if intersection is not None:
                self.cuboids.append(intersection)

        if c.on:
            self.cuboids.append(c)

    def calc_volume(self) -> int:
        return sum((1 if c.on else -1) * c.calc_volume() for c in self.cuboids)


def parse_input(input: str, limit: int = None) -> Generator[Cuboid]:
    for line in input.splitlines():
        switch, rest = line.split(" ")
        [[x1, x2], [y1, y2], [z1, z2]] = list(map(lambda n: map(int, n.split("=")[1].split("..")), rest.split(",")))
        if limit is not None and any(not -limit < a < limit for a in [x1, x2, y1, y2, z1, z2]):
            continue

        yield Cuboid(Range(x1, x2), Range(y1, y2), Range(z1, z2), switch == "on")


def part1(input: str) -> int:
    reactor = Reactor()
    for cuboid in parse_input(input, 50):
        reactor.apply_cuboid(cuboid)

    return reactor.calc_volume()


def part2(input: str) -> int:
    reactor = Reactor()
    for cuboid in parse_input(input):
        reactor.apply_cuboid(cuboid)

    return reactor.calc_volume()


if __name__ == '__main__':
    example = Path('./example_input.txt').read_text()
    input = Path('./input.txt').read_text()

    for idx, (want_example, want_answer, fn) in enumerate([(474140, 623748, part1), (2758514936282235, 1227345351869476, part2)]):
        start = datetime.now()
        got_example = fn(example)
        assert got_example == want_example, f"example: got {got_example}, want {want_example}"

        print(f"Part{idx + 1}(example): {got_example} ({datetime.now() - start})")
        start = datetime.now()
        got_answer = fn(input)
        assert got_answer == want_answer, f"answer: got {got_answer}, want {want_answer}"

        print(f"Part{idx + 1}(real): {got_answer} ({datetime.now() - start})")
        print("done!")
