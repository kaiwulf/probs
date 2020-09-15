# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(S, K):
    num_to_day = {0: "Mon", 1: "Tues", 2: "Wed", 3: "Thu", 4: "Fri", 5: "Sat", 6: "Sun"}
    day_to_num = {"Mon":0, "Tues":1, "Wed":2, "Thurs":3, "Fri":4, "Sat":5, "Sun":6}
    day = (K+day_to_num[S]) % 7
    return num_to_day[day]

def solution(X, Y, A):
    N = len(A)
    result = -1
    nX = 0
    nY = 0
    for i in range(N):
        if A[i] == X:
            nX += 1
        if A[i] == Y:
            nY += 1
        if nX == nY:
            result = i
    return result

    