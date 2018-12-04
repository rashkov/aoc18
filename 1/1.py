running_total = 0
total_from_first_pass = 0
frequency_dict = { 0: 1 }
passes = 0
should_loop = True

while should_loop:
    # Record answer for part 1
    if passes == 1:
        total_from_first_pass = running_total
    passes += 1

    file = open("./evilinput.txt", "r")
    for number in file:
        num = int(number)
        running_total += num
        times_seen = frequency_dict.get(running_total, 0)
        frequency_dict[running_total] = times_seen + 1
        if times_seen + 1 == 2:
            first_repeat = running_total
            should_loop = False
            break

print("Part 1 answer: ", total_from_first_pass) # Part A
print("Part 2 answer: ", first_repeat) # Part B
file.close()
