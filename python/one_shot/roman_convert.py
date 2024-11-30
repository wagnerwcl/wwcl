
def convert_to_number(number):
    ROMAN_LETTERS = {
        "I": 1,
        "V": 5 ,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000
    }

    total = 0
    previous = 0

    for char in reversed(number):
        current = ROMAN_LETTERS[char]

        if current < previous:
            total -= current
        else:
            total += current

    previous = current

    return total