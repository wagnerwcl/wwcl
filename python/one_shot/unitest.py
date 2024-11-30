import unittest

import roman_convert as roman

class Test(unittest.TestCase):

    def test_convert_to_number_I(self):
        roman_number = "I"
        number = roman.convert_to_number(roman_number)
        self.assertEqual(number, 1)

    def test_convert_to_number_V(self):
        roman_number = "V"
        number = roman.convert_to_number(roman_number)
        self.assertEqual(number, 5)
    
    def test_convert_to_number_X(self):
        roman_number = "X"
        number = roman.convert_to_number(roman_number)
        self.assertEqual(number, 10)

    def test_convert_to_number_L(self):
        roman_number = "L"
        number = roman.convert_to_number(roman_number)
        self.assertEqual(number, 50)

    def test_convert_to_number_C(self):
        roman_number = "C"
        number = roman.convert_to_number(roman_number)
        self.assertEqual(number, 100)

    def test_convert_to_number_D(self):
        roman_number = "D"
        number = roman.convert_to_number(roman_number)
        self.assertEqual(number, 500)

    def test_convert_to_number_M(self):
        roman_number = "M"
        number = roman.convert_to_number(roman_number)
        self.assertEqual(number, 1000)


if __name__ == '__main__':
    unittest.main()