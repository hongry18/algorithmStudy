import sys

stack = []
loop_cnt = 0
m = 0
try:
    loop_cnt = int(sys.stdin.readline().split()[0])
except:
    sys.exit(0)

for _ in range(loop_cnt):
    items = sys.stdin.readline().split()
    if items[0] == '1':
        i = int(items[1])
        stack.append(i)
        if m < i:
            m = i
    elif items[0] == '2':
        stack.pop()
        if len(stack) > 0:
            m = max(stack)
        else:
            m = 0
    else:
        print(m)
