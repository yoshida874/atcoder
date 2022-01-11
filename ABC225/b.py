N = int(input())
counter = [0] * (N + 1)

for _ in range(N - 1):
    a, b = map(int, input().split())
    counter[a] += 1
    counter[b] += 1

if N - 1 in counter:
    print("Yes")
else:
    print("No")