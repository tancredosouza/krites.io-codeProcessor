def fib(x):
    if x < 2:
        return x

    return fib(x - 1) + fib(x-2)


if __name__ == "__main__":
    N = int(input())
    while N > 0:
        x = int(input())
        print(fib(x))
        N = N - 1
