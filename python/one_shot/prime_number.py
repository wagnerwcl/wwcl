
def prime_checker(number):
    is_prime = True
    for i in range(2, number):
        if number % i == 0:
            is_prime = False
            break
    if is_prime:
        print(f"{number} is a prime number.")
    else:
        print(f"{number} is not a prime number.")

n = int(input())
prime_checker(number=n)

for i in range(2, 10):
    value = n % i
    print(f"{n} % {i} = {value}")
