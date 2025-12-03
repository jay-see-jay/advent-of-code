# Identify invalid product IDs

# Input consists of input ranges, separated by commas
# [start]-[end]

# Invalid IDs are those with repeating patterns e.g.
# 55, 6464 or 123123

# Need to find all the invalid IDs in the given ranges
# then sum them.

def get_ranges(input):
    ranges = input.read()
    return ranges.split(",")

def is_id_invalid(id, predicate):
    seq_length = len(id) // 2

    while seq_length > 0:
        test_seq = id[:seq_length]
        start = seq_length 
        matches = 1
        while start < len(id):
            end = start + seq_length 
            curr_seq = id[start:end]
            if test_seq == curr_seq:
                matches += 1

            start += seq_length

        if len(id) / seq_length == matches and predicate(matches):
            return True 

        seq_length -= 1

    return False


def handle_range(start, end, predicate):
    curr = start

    invalid_sum = 0

    while curr <= end:
        if is_id_invalid(str(curr), predicate):
            invalid_sum += curr

        curr += 1

    return invalid_sum
        
def process_input(predicate):
    invalid_sum = 0
    
    with open("./years/twenty_five/two/input.txt") as input:
        ranges = get_ranges(input)
        for i in range(len(ranges)):
            start, end = ranges[i].split("-")
            invalid_sum += handle_range(int(start), int(end), predicate)

    print(f"Answer: {invalid_sum}")


def handle21():

    def predicate(matches):
        return matches == 2

    process_input(predicate)

def handle22():

    def predicate(matches):
        return matches >= 2

    process_input(predicate)
