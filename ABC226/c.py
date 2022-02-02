import sys
sys.setrecursionlimit(10**7)
 
n = int(input())
 
T = []
A = [[] for _ in range(n)]
for i in range(n):
    skill_row = list(map(int, input().split()))
    T.append(skill_row[0])
    A[i] = skill_row[2:]
 
flag = [False for _ in range(n)]
def dfs(now, time):
    now_time = time
    for node in A[now]:
        if flag[node-1] == False:
            now_time += dfs(node-1, time)
    flag[now] = True
    return now_time + T[now]
 
print(dfs(n-1, 0))