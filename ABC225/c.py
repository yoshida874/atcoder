N, M = map(int, input().split())
B = [list(map(int, input().split())) for _ in range(N)]

res = B[0][0]
for i in range(N):
    for j in range(M):
        if B[i][j] != i * 7 + j + res:
            print("No")
            exit()
        if j and B[i][j - 1] % 7 == 0:
            print("No")
            exit()

print("Yes")
