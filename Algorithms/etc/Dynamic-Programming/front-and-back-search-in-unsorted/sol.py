import random

def search(arr,x):

    n = len(arr)
    front = 0
    back = n - 1

    while front <= back:
        if arr[front] == x or arr[back] == x:
            return x
        front += 1
        back -= 1

if __name__ == '__main__':
    arr = [i**2 for i in range(50)]
    random.shuffle(arr)

    x = 3
    print(search(arr, x))
