from machine import *
import re

REGISTERS = [1, 0, 0, 0, 0, 0]

IP_REG = None

instructions = []

file = open("./input.txt", "r")
for line in file:
    ip_RE = '#ip (\d)'
    instr_RE = '(\w+) (\d+) (\d+) (\d+)'
    if re.compile(ip_RE).match(line):
        IP_REG = int(re.compile(ip_RE).findall(line)[0])
    elif re.compile(instr_RE).match(line):
        instructions.append(re.compile(instr_RE).findall(line)[0])
    else:
        throw("AH")

fix1 = False
# fix2 = False
fixed = False
n = 0
while REGISTERS[IP_REG] < len(instructions):
    n+=1

    [op, a, b, c] = instructions[REGISTERS[IP_REG]]
    [a, b, c] = [int(a), int(b), int(c)]

    op_fn = op_dispatcher[op]

    print(REGISTERS)
    print(op_fn, a, b, c)
    if not fix1 and REGISTERS[2] < REGISTERS[3] and [a,b,c] == [5,2,4]:
        fix1 = True
        REGISTERS[2] = REGISTERS[3]
    #     # import pdb; pdb.set_trace()
    # elif REGISTERS[5] < REGISTERS[3] and [a,b,c] == [5,1,5]:
    #     REGISTERS[5] = REGISTERS[3]
    # elif REGISTERS[1] == 35:
    #     fixed = True

    if [a,b,c] == [5,0,0]:
        import pdb; pdb.set_trace()

    REGISTERS = op_fn(REGISTERS, a, b, c)
    REGISTERS[IP_REG] += 1

print(REGISTERS)
