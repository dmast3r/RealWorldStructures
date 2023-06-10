from random import choice
from string import ascii_letters


def generate_random_string(length):
    letters = ascii_letters
    result_str = ''.join(choice(letters) for i in range(length))
    return result_str
