import unittest

def find_min_max(numbers):
    """
    Finds the smallest and largest numbers in a list.

    Args:
        numbers (list): A list of numeric values.

    Returns:
        tuple: A tuple containing the smallest and largest numbers.
    """
    if not numbers:
        raise ValueError("The list is empty. Please provide a non-empty list.")
    
    smallest = min(numbers)
    largest = max(numbers)
    
    return smallest, largest

numbers_list = [10, 5, 23, -4, 0, 18]
smallest, largest = find_min_max(numbers_list)
print(f"Smallest number: {smallest}, Largest number: {largest}")
