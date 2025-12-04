from utils import Logger

# Puzzle input = joltage ratings for batteries
# The batteries are arranged into banks; each line of
# of the input corresponds to a single bank.

# Within each bank, **2** batteries must be turned on.

# The joltage produced by the bank is determined by
# concatanating the numbers for the two batteries that
# are on.

# e.g. "2" and "4" would mean a joltage of 24.

# The order cannot be changed, you must use the order
# as it is in the input.

# You'll need to find the **largest possible joltage**

# The total output is then the sum of the maximum voltage from each bank.


class Bank:
    def __init__(self, joltages: str, battery_count: int, logger: Logger):
        self.joltages = joltages
        self.battery_count = battery_count
        self.size = len(joltages)
        self.logger = logger
        self.candidates = []
        logger.print(f"Init with : {joltages}")

    def combine_candidates(self) -> int:
        self.logger.print("Combining candidates")
        sum = 0
        factor = 1
        for candidate in reversed(self.candidates):
            self.logger.print(f"Adding: {candidate} * {factor}")
            sum += candidate * factor
            factor *= 10

        self.logger.print(f"Sum: {sum}")

        return sum

    def find_max(self, index=0) -> int:
        self.logger.print(f"Starting from: {index}")
        if len(self.candidates) >= self.battery_count:
            self.logger.print(f"Reached battery count: {len(self.candidates)}")
            return self.combine_candidates()

        if index >= self.size:
            self.logger.print("Exceeded bank")
            return self.combine_candidates()

        max = int(self.joltages[index])
        max_i = index

        i = index + 1
        end = self.size - self.battery_count + len(self.candidates)
        self.logger.print(f"Ending at: {end}")

        while i < end:
            candidate = int(self.joltages[i])
            if max < candidate:
                max_i = i
                max = candidate
            i += 1

        self.candidates.append(max)
        self.logger.print(
            f"Candidates: {", ".join(str(x) for x in self.candidates)}")

        return self.find_max(max_i + 1)


def process_input(battery_count: int, logger: Logger):
    total_output = 0

    with open("./years/twenty_five/three/input.txt") as input:
        while line := input.readline():
            logger.print(f"Processing line: {
                         line}, length: {len(line)}")
            bank = Bank(line, battery_count, logger)
            bank_max = bank.find_max()
            total_output += bank_max

    # Iterate over each character to find the largest. The largest may appear multiple times, so need to pick the one that preceeds the second largest.

    # Iterate through and extract candidates, return the max

    print(f"Answer: {total_output}")


def handle31(logger: Logger):
    process_input(2, logger)

    # For part, need to find max with 12 batteries
    # Find largest number between 0 index and
    # length - 12 index
    # Find the next highest number and keep bisecting the
    # remaining batteries


def handle32(logger: Logger):
    process_input(12, logger)
