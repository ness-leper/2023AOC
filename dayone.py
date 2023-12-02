def CalibrationSum(line:str):
    first = "UNKNOWN"
    last = "UNKNOWN"
    if len(line) > 0:
        for char in line:
            if char.strip().isdigit():
                if first == "UNKNOWN":
                    first = char.strip()
                last = char.strip()

    return int(f"{first}{last}")

def find_all(a_str, sub):
    start = 0
    while True:
        start = a_str.find(sub, start)
        if start == -1: return
        yield start
        start += len(sub) # use start += 1 to find overlapping matches

def SolveDayOne_PartOne():
    f = open("./dayoneinput.txt", "r")
    sumCalibration = 0
    for line in f:
        sumCalibration += CalibrationSum(line)

    print(sumCalibration)
    f.close()

def D1P2ConvertWords(line: str):
    numbers = [
            "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"
            ]

    numberString = {} 
    
    for idx,number in enumerate(numbers):
        for occ in find_all(line,number):
            print(occ)
            numberString[occ] = str(idx + 1)
    
    for charIdx, char in enumerate(line): 
        if char.strip().isdigit():
            numberString[charIdx] = char.strip()

    output = ""
    for idx in sorted(numberString.keys()):
        output += str(numberString.get(idx))

    return output

def SolveDayOne_PartTwo(file: str): 
    f = open(file, "r")
    sumCal = 0
    for line in f:
        # Convert words into numbers
        # Convert overlapping words
        # build a new coordinate with the original numbers and the converted word numbers
        coordinate = D1P2ConvertWords(line)
        print(f"{line.strip()}: {coordinate} {CalibrationSum(coordinate)}")
        sumCal += CalibrationSum(coordinate)
    f.close()

    print (sumCal)

SolveDayOne_PartTwo("./dayoneinput.txt")
# SolveDayOne_PartTwo("./d1p2_test.txt")
