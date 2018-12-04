file = open("./input.txt", "r")
all = []
twos_count = 0
threes_count = 0
for str in file:
    all.append(str.strip())
    char_counts = {}
    for char in str.strip():
        count = char_counts.get(char, 0)
        char_counts[char] = count + 1

    has_two = False
    has_three = False
    for x,y in char_counts.items():
        if y == 3:
            has_three = True
        if y == 2:
            has_two = True
    if has_two:
        twos_count += 1
    if has_three:
        threes_count += 1
    #print(list(filter(lambda x: x == 3, char_counts.items())))
print("Part 1:", twos_count * threes_count)

scores = {}
for index, word in enumerate(all):
    for index2, word2 in enumerate(all):
        # if index == index2:
        #     continue
        score = 0
        for i in range(0, len(word)):
            if word[i] != word2[i]:
                score += 1
        a = scores.get(score, None)
        if a == None:
            a = { 'words': [] }
        a['words'].append(word)
        a['words'].append(word2)
        scores[score] = a
print(set(scores[1]['words']))
# part 2 answer: zihwtxagifpbsnwleydukjmqv
