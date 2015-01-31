import math

num = 600851475143

stop = int(math.sqrt(num))

grt = 0

i = stop

while i >= 2:
    print i
    prime = True
    for j in range(2, i-1):
        if i % j == 0:
            prime = False
            break
    if prime and i > grt and num % i == 0:
            grt = i
            break
    i = i - 1
print grt
