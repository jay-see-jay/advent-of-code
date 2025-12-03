import sys
from one.one import handle11, handle12
from two import handle21, handle22

if __name__ == "__main__":
    day = int(sys.argv[1])
    part = int(sys.argv[2])

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
            else:
                print(f"Running day {day}, part two")
                handle22()

        case _:
            print(f"Day {day} is not yet implemented")

