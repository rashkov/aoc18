import re

file = open("./input.txt", "r")

runset = {
    'before': None,
    'instruction': None,
    'after': None
}
runsets = [runset]

for line in file:
    if runset['before'] != None and runset['instruction'] != None and runset['after'] != None:
        runsets.append(runset)
        runset = {
            'before': None,
            'instruction': None,
            'after': None
        }

    if len(line) == 1:
        continue

    if re.compile('Before').match(line):
        four_ints = re.compile('(\d+)[, ]{1,}(\d+)[, ]{1,}(\d+)[, ]{1,}(\d+)').findall(line)[0]
        runset['before'] = four_ints
    elif re.compile('After').match(line):
        four_ints = re.compile('(\d+)[, ]{1,}(\d+)[, ]{1,}(\d+)[, ]{1,}(\d+)').findall(line)[0]
        runset['after'] = four_ints
    elif re.compile('^(\d+)[, ]').match(line):
        four_ints = re.compile('(\d+)[, ]{1,}(\d+)[, ]{1,}(\d+)[, ]{1,}(\d+)').findall(line)[0]
        runset['instruction'] = four_ints

# Multiplication:
#  mulr (multiply register) stores into register C the result of multiplying register A and register B.
def mulr(registers, a, b, c):
    registers = list(registers)
    registers[c] = registers[a] * registers[b]
    return registers
assert mulr((4, 1, 2, 3), *(2, 3, 2)) == [4, 1, 6, 3]

#  muli (multiply immediate) stores into register C the result of multiplying register A and value B.
def muli(registers, a, b, c):
    registers = list(registers)
    registers[c] = registers[a] * b
    return registers
assert muli((4,1,2,3), *(2,2,2)) == [4,1,4,3]

# Bitwise AND:
#     banr (bitwise AND register) stores into register C the result of the bitwise AND of register A and register B.
def banr(registers, a, b, c):
    registers = list(registers)
    registers[c] = registers[a] & registers[b]
    return registers
assert banr((4,1,2,3), *(1,3,0)) == [1,1,2,3]

#     bani (bitwise AND immediate) stores into register C the result of the bitwise AND of register A and value B.
def bani(registers, a, b, c):
    registers = list(registers)
    registers[c] = registers[a] & b
    return registers
assert banr((4,1,2,3), *(2,3,0)) == [2,1,2,3]

# Bitwise OR:
#     borr (bitwise OR register) stores into register C the result of the bitwise OR of register A and register B.
def borr(registers, a, b, c):
    registers = list(registers)
    registers[c] = registers[a] | registers[b]
    return registers
assert borr((4,1,2,3), *(0,1,0)) == [5,1,2,3]

#     bori (bitwise OR immediate) stores into register C the result of the bitwise OR of register A and value B.
def bori(registers, a, b, c):
    registers = list(registers)
    registers[c] = registers[a] | b
    return registers
assert borr((4,1,2,3), *(0,1,0)) == [5,1,2,3]

# 
# Assignment:
# 
#     setr (set register) copies the contents of register A into register C. (Input B is ignored.)
def setr(registers, a, b, c):
    registers = list(registers)
    registers[c] = registers[a]
    return registers
assert setr((4,1,2,3), *(3, None, 0)) == [3,1,2,3]

#     seti (set immediate) stores value A into register C. (Input B is ignored.)
def seti(registers, a, b, c):
    registers = list(registers)
    registers[c] = a
    return registers
assert seti((4,1,2,3), *(9,None,0)) == [9,1,2,3]

#
# Greater-than testing:
#     gtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0.
def gtir(registers, a, b, c):
    registers = list(registers)
    if a > registers[b]:
        registers[c] = 1
    else:
        registers[c] = 0
    return registers
assert gtir((4,1,2,3), *(5,1,0)) == [1,1,2,3]
assert gtir((4,1,2,3), *(0,1,0)) == [0,1,2,3]

#     gtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0.
def gtri(registers, a, b, c):
    registers = list(registers)
    if registers[a] > b:
        registers[c] = 1
    else:
        registers[c] = 0
    return registers
assert gtri((4,1,2,3), *(0,1,0)) == [1,1,2,3]
assert gtri((4,1,2,3), *(1,2,0)) == [0,1,2,3]

#     gtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0.
def gtrr(registers, a, b, c):
    registers = list(registers)
    if registers[a] > registers[b]:
        registers[c] = 1
    else:
        registers[c] = 0
    return registers
assert gtrr((4,1,2,3), *(0,1,0)) == [1,1,2,3]
assert gtrr((4,1,2,3), *(1,2,0)) == [0,1,2,3]

#
# Equality testing:
#
#     eqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0.
def eqir(registers, a, b, c):
    registers = list(registers)
    if a == registers[b]:
        registers[c] = 1
    else:
        registers[c] = 0
    return registers
assert eqir((4,4,2,3), *(4,0,0)) == [1,4,2,3]
assert eqir((4,1,2,3), *(9,0,0)) == [0,1,2,3]

#     eqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0.
def eqri(registers, a, b, c):
    registers = list(registers)
    if registers[a] == b:
        registers[c] = 1
    else:
        registers[c] = 0
    return registers
assert eqri((4,4,2,3), *(0,4,0)) == [1,4,2,3]
assert eqri((4,1,2,3), *(0,5,0)) == [0,1,2,3]

#     eqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0.
def eqrr(registers, a, b, c):
    registers = list(registers)
    if registers[a] == registers[b]:
        registers[c] = 1
    else:
        registers[c] = 0
    return registers

assert eqrr((4,4,2,3), *(0,1,0)) == [1,4,2,3]
assert eqrr((4,1,2,3), *(0,2,0)) == [0,1,2,3]

stats = {}
ops = [mulr, muli]
for op in ops:
   instruction = runsets[0]['instruction'][1:]
   before = runsets[0]['before']
   print(op(before, *instruction))

