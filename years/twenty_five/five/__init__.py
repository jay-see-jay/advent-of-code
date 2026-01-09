from utils import Logger
from typing import Self

# Set up ranges of ingredient IDs
# Iterate through and check if any of the given
# ingredient IDs fit within the range.


class Range:
    def __init__(
            self,
            min: str,
            max: str,
            prev: Self | None = None,
            next: Self | None = None
    ):
        self.min = int(min)
        self.max = int(max)
        self.prev = prev
        self.next = next

    def is_in_range(self, id: int) -> bool:
        return self.min <= id and self.max >= id

    def __repr__(self):
        return "Range() {:,}-{:,}".format(self.min, self.max)


class Ranges:
    def __init__(self, logger: Logger):
        # Turn this into a linked list, that is sorted
        # and consolidated as ranges are added.
        self.list = None
        self.len = 0
        self.logger = logger

    def __repr__(self):
        ranges = []
        curr = self.list
        while curr is not None:
            ranges.append(curr.__repr__())
            curr = curr.next

        return f"Ranges() {"\n".join(ranges)}"

    def consolidate_ranges(self, curr: Range):
        next = curr.next
        while next is not None and curr.max >= next.min:
            curr.min = min(curr.min, next.min)
            curr.max = max(curr.max, next.max)

            curr.next = next.next
            self.len -= 1

            curr = next
            next = next.next

    def add_range(self, line: str):
        min, max = line.split("-")
        r = Range(min, max)
        self.logger.print(f"Processing range: {r}")
        if self.list is None:
            self.logger.print(f"Adding first range: {r}")
            self.list = r
            self.len += 1
            return

        if r.max < self.list.min:
            r.next = self.list
            self.list.prev = r
            self.list = r
            self.len += 1
            return

        prev = None
        curr = self.list
        self.logger.print(f"Setting curr to {self.list}")
        # 3 - 5
        # 10 - 14
        while curr is not None and curr.max < r.min:
            self.logger.print(f"Setting prev to {
                              curr} and curr to {curr.next}")
            prev = curr
            curr = curr.next

        self.logger.print(f"Curr is {curr}")

        if curr is None:
            self.logger.print(f"Appending {r} to list")
            r.prev = prev
            prev.next = r
            self.len += 1
            return

        # 10-15 >= 7-8... 15 >= 7
        if r.max < curr.min:
            self.logger.print(f"Prepending {r} to list")
            r.next = curr
            curr.prev.next = r
            r.prev = curr.prev
            curr.prev = r
            self.len += 1
            return

        if r.min < curr.min:
            curr.min = r.min

        if r.max > curr.max:
            curr.max = r.max

        self.consolidate_ranges(curr)

    def check_id(self, id_str: str) -> bool:
        id = int(id_str)
        for i in range(self.len):
            r = self.list[i]
            if r.is_in_range(id):
                self.logger.print(f"{id} in range: {r.min}-{r.max}")
                return True

        return False

    def get_fresh_ingredient_count(self):
        total = 0
        curr = self.list
        while curr is not None:
            total += (curr.max - curr.min) + 1
            curr = curr.next

        return total


def process_input(logger: Logger, is_test=False):
    input_file = "test" if is_test else "input"

    ranges = Ranges(logger)

    are_ranges_complete = False

    fresh_ingredients = 0

    with open(f"./years/twenty_five/five/{input_file}.txt") as input:
        while line := input.readline():
            line = line.strip()
            logger.print(f"Processing line: {line}, {str(len(line))}")
            if len(line) == 0:
                logger.print(f"Found gap: {line}")
                are_ranges_complete = True
                continue

            if not are_ranges_complete:
                ranges.add_range(line)
            else:
                if ranges.check_id(line):
                    fresh_ingredients += 1

    print(f"Answer: {fresh_ingredients}")


def handle51(logger: Logger, is_test=False):
    process_input(logger, is_test)


def process_input2(logger: Logger, is_test=False):
    input_file = "test" if is_test else "input"

    ranges = Ranges(logger)

    with open(f"./years/twenty_five/five/{input_file}.txt") as input:
        while line := input.readline():
            line = line.strip()
            logger.print(f"Processing line: {line}, {str(len(line))}")
            if len(line) == 0:
                logger.print(f"Found gap: {line}")
                break

            ranges.add_range(line)

    logger.print(ranges)

    print(f"Answer: {ranges.get_fresh_ingredient_count()}")


def handle52(logger: Logger, is_test=False):
    process_input2(logger, is_test)
