import math

def generateKyberParams(q, k, n):
    paramsPolyBytes = int(math.ceil(math.log2(q)) * n / 8)
    ctr = 0
    while True:
        if pow(q * ctr, 1, 2 ** 16) == 1:
            break
        ctr += 1
    print(f"q: {q}, k: {k}, n: {n}, paramsPolyBytes: {paramsPolyBytes}, paramsQInv: {ctr}")

generateKyberParams(3329, 2, 256)