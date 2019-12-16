import sys

data = []
loop_cnt = 0
m = 1e+10
try:
    loop_cnt = int(sys.stdin.readline().split()[0])
except:
    sys.exit(0)

for _ in range(loop_cnt):
    items = sys.stdin.readline().split()
    if items[0] == '1':
        i = int(items[1])
        data.append(i)
        if m > i:
            m = i
    elif items[0] == '2':
        i = int(items[1])
        data.pop(data.index(i))
        if len(data) > 0:
            m = min(data)
        else:
            m = 1e+10
    else:
        print(m)
