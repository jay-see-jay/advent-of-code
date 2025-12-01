import re

# Due to new security protocols, the password is locked
# in the safe below. Please see the attached document
# for the new combination.

# There is dial with values 0 through 99
# Given a series of rotations, e.g. L8 or R10, will
# need to use modulo to handle 0 L5 to get 95.

# The dial starts pointing at 50.

# The actual password is the number of times the dial
# points at 0 after any rotation instruction.

# Before finding first 0, add up the rotations until you get a multiple of +/- 50. Then you are at 0.
# From then on add up the instructions until you get a multiple of +/- 100 (or 0)
# Where L is negative and R is positive

class Dial:
    def __init__(self, start = 50):
        self.position = start
        self.count = 0
        self.click_count = 0

    def move(self, instruction):
        direction = -1 if instruction.startswith("L") else 1

        amount_digits = re.findall("\\d", instruction)
        amount = int("".join(amount_digits))

        while amount > 0:
            vector = 1 * direction
            new_position = self.position + vector
            self.position = new_position % 100
            if self.position == 0:
                self.click_count += 1
            amount -= 1

        if self.position == 0:
            self.count += 1

def handle11():
    d = Dial()
    with open("./one/input.txt") as input:
        while line := input.readline():
            d.move(line.rstrip())

    print(f"Answer: {d.count}")

# Use password method 0x434C4943B
# i.e. count the number of times **any click** leads to the dial pointing at 0.

def handle12():
    d = Dial()
    with open("./one/input.txt") as input:
        while line := input.readline():
            d.move(line.rstrip())

    print(f"Answer: {d.click_count}")

