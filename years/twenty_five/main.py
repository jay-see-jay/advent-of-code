import sys
from one.one import handle11, handle12
from two import handle21, handle22
from three import handle31, handle32
from utils import Logger

if __name__ == "__main__":
    day = int(sys.argv[1])
    part = int(sys.argv[2])
    is_debug = len(sys.argv) > 3 and sys.argv[3] == "debug"

    if is_debug:
        print("Running in debug")

    logger = Logger(is_debug)

    match day:
        case 1:
            if part == 1:
                print(f"Running day {day}, part one")
                handle11()
            elif part == 2:
                print(f"Running day {day}, part two")
                handle12()
            else:
                print("There are no other parts!")
        case 2:
            if part == 1:
                print(f"Running day {day}, part one")
                handle21()
            elif part == 2:
                print(f"Running day {day}, part two")
                handle22()
            else:
                print("There are no other parts!")
        case 3:
            if part == 1:
                print(f"Running day {day}, part one")
                handle31(logger)
            elif part == 2:
                print(f"Running day {day}, part two")
                handle32(logger)
            else:
                print("There are no other parts!")

        case _:
            print(f"Day {day} is not yet implemented")
