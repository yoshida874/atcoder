from bisect import bisect
n, k = map(int, input().split())
info = []
for i in range(n):
    score = sum(map(int, input().split()))
    info.append([i, score, 0])
 
info.sort(key=lambda x:x[1])
scores = [v[1] for v in info]
 
for data in info:
    rank = n-bisect(scores, data[1]+300)+1
    data[2] = rank
    
info.sort(key=lambda x:x[0])
 
for data in info:
    if(data[2]<=k):
        print('Yes')
    else:
        print('No')