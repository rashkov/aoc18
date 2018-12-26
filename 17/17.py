import re

file = open("./input.txt", "r")

grid = {}
data = []
xs = []
ys = []

for line in file:
    x1, x2, y1, y2 = (None,)*4
    two_x = re.compile("x=(\d+)\.\.(\d+)").findall(line)
    one_x = re.compile("x=(\d+)").findall(line)
    if len(two_x) == 1:
        x1 = two_x[0][0]
        x2 = two_x[0][1]
    elif len(one_x) == 1:
        x1 = one_x[0]

    two_y = re.compile("y=(\d+)\.\.(\d+)").findall(line)
    one_y = re.compile("y=(\d+)").findall(line)
    if len(two_y) == 1:
        y1 = two_y[0][0]
        y2 = two_y[0][1]
    elif len(one_y) == 1:
        y1 = one_y[0]

    if x1 and x2:
        for i in range(int(x1), int(x2)):
            



#     [xs.append(i) for i in [x1, x2] if i != None]
#     [ys.append(i) for i in [y1, y2] if i != None]
#     data.append({'x1': x1, 'x2': x2, 'y1': y1, 'y2': y2})

# xs = sorted(xs)
# ys = sorted(ys)
# print(xs[0], xs[len(xs)-1])
# print(ys[0], ys[len(ys)-1])
# grid[]
