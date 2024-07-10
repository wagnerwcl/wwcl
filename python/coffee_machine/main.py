# TODO: Step 4

from data import *

money = 0

def report():
    resources["money"] = money
    
    for item in resources:
        if item == "water" or item == "milk":
            print(f"{(item.capitalize())}: {resources[item]}ml")
        elif item == "coffee":
            print(f"{(item.capitalize())}: {resources[item]}g")
        elif item == "money":
            print(f"{(item.capitalize())}: {resources[item]}$")


userOption = "latte"
def drink(userOption):
    print(MENU[userOption]["ingredients"]["water"])




def main():
    machine_state = True

    while machine_state == True:
        option = input("What you want? (espresso, latte, cappuccino): ")
        
        if option == "report":
            report()
        elif option == "espresso":
            print("espresso")
        elif option == "latte":
            print("latte")
        elif option == "cappuccino":
            print("cappuccino")
        elif option == "off":
            machine_state = False
        else:
            print("Wrong Option")




drink(userOption)
    