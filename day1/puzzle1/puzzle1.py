import os

DIGITS = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9']

def find_first_digit(line: str) -> str:
    for letter in line:
        if letter in DIGITS:
            return letter
    return None

def find_last_digit(line: str) -> str:
    for letter in line[::-1]:
        if letter in DIGITS:
            return letter
    return None

def find_calibration_value(line: str) -> int:
    first_digit = find_first_digit(line)
    last_digit = find_last_digit(line)
    if first_digit is None or last_digit is None:
        return None
    return int(f"{first_digit}{last_digit}")

def find_calibration_sum(file_path: str) -> int:
    if os.path.exists(file_path):
        with open(file_path) as f:
            lines = f.readlines()
        
        calibration_value_sum = 0
        for line in lines:
            calibration_value = find_calibration_value(line)
            calibration_value_sum += calibration_value
        
        return calibration_value_sum
    else:
        raise FileNotFoundError

def test_func():
    test_string = "five3onelxjninenine45"
    assert find_first_digit(test_string) == "3"
    assert find_last_digit(test_string) == "5"
    assert find_calibration_value(test_string) == 35
    assert find_calibration_sum("test_input.txt") == 57 


def main():
    print(find_calibration_sum("puzzle_input.txt"))

if __name__ == "__main__":
    main()