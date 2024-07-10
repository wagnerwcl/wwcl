def add(number1, number2):
    """
    Sum two numbers.
    """
    return number1 + number2


def  sub(number1, number2):
    """
    Subtract two numbers.
    """
    return number1 - number2


def mult(number1, number2):
    """
    Multiply two numbers.
    """
    return number1 * number2


def div(number1, number2):
    """
    Divide two numbers.
    """
    return number1 / number2


operations = {
    "+": add,
    "-": sub,
    "x": mult,
    "/": div
}

function = operations["+"]
function(2,3)
print(function)