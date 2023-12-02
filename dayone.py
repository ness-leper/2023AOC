def SolveDayOne_PartOne():
    f = open("./dayoneinput.txt", "r")
    sumCalibration = 0
    for line in f:
        if len(line) > 0:
            first = "UNKNOWN"
            last = "UNKNOWN"
            for char in line:
                if char.strip().isdigit():
                    if first == "UNKNOWN":
                        first = char.strip()
                    last = char.strip()
            print(f"{first}{last}")
            sumCalibration += int(f"{first}{last}")

    print(sumCalibration)
    f.close()
