import random
import string
import hashlib


def random_string (length):
    # choose from all lowercase letter
    letters = string.hexdigits
    result_str = ''.join(random.choice(letters) for i in range(length))
    return result_str


def password_generator(password: str):
    result = hashlib.sha256(password.encode())
    return result.hexdigest()



