def euc_mod(x, mod):
    rem = q % n
    if rem < 0:
        rem += q
    return rem

def find_root_of_unity(n, q):
    for x in range(2, q-1):
        if pow(x, n, q) == 1:
            return x
    return 0 # AKA. None found

def br(i, k):
        bin_i = bin(i & (2**k - 1))[2:].zfill(k)
        return int(bin_i[::-1], 2)

def find_zetas(root_of_unity, q, mont_r):
    return [(mont_r * pow(root_of_unity,  br(i,7), q)) % q for i in range(128)]

def find_inv_zetas(zetas, q):
    return [pow(x, -1, q) % q for x in zetas]

n = 256
q = 10753

# We require q mod n = 1
if euc_mod(q, n) != 1:
    print("Error, requirement on q not satisfied!")

# Brute force root of unity
root_of_unity = find_root_of_unity(n, q)
zetas = find_zetas(root_of_unity, q, pow(2, 16, q))
print(root_of_unity)
print(zetas)
print(find_inv_zetas(zetas, q))
