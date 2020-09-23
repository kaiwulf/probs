def gcd(m, n):
    r = m % n
    if r == 0:
        return n
    else:
        return gcd(n, r)

def main():
    g = gcd(7, 14)
    print("gcd is: {}" + g)