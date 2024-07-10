
option = True
auction = []
new_value = 0

def add_bidder(name, price):
    bid = {}
    bid["name"] = name
    bid["price"] = price
    auction.append(bid)
    

while option:
    name = input("What is your name? ")
    price = int(input("What is your bid? "))
    add_bidder(name, price)

    opt = input("Would you like to bid? Type 'yes' or 'no' ")
    if opt == "no":
        option = False

currently_value = 0
new_value = 0
for key in auction:
    currently_value = key["price"]
    if currently_value > new_value:
        new_value = currently_value
        winner = key["name"]
    else:
        print(f"The winner is {winner} with a bid of {new_value}")

