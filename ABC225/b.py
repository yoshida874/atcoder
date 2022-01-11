N = int(input())
counter = [0] * (N + 1)

for _ in range(N - 1):
    a, b = map(int, input().split())
    counter[a] += 1
    counter[b] += 1

print("Yes" if N - 1 in counter else "No")