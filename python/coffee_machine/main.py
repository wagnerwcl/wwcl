from data import *

money = 0
resources["money"] = money

quarters = 0.25
dimes = 0.10
nickles = 0.05
pennies = 0.01


def report():
    """Prints the current state of the resources."""
    for item in resources:
        if item == "water" or item == "milk":
            print(f"{(item.capitalize())}: {resources[item]}ml")
        elif item == "coffee":
            print(f"{(item.capitalize())}: {resources[item]}g")
        elif item == "money":
            print(f"{(item.capitalize())}: {resources[item]}$")


def checkResource(userOption):
    """Returns True when the resources are sufficient, or False if resources are insufficient."""
    for ingredient in MENU[userOption]["ingredients"]:
        if MENU[userOption]["ingredients"][ingredient] <= resources[ingredient]:
            print(f"Enough {ingredient}")
        else:
            print(f"Sorry there is not enough {ingredient}\n")
            return False
    print("")


def insertCoins():
    """Returns the total calculated from coins inserted."""
    print("Please insert coins.")
    quarters = int(input("How many quarters (0.25)?: "))
    dimes = int(input("How many dimes (0.10)?: "))
    nickles = int(input("How many nickles (0.05)?: "))
    pennies = int(input("How many pennies? (0.01): "))
    total = (quarters * 0.25) + (dimes * 0.10) + (nickles * 0.05) + (pennies * 0.01)
    print("Total: ", total)
    print("")
    return total


def checkTransaction(total, userOption):
    """Returns True when the payment is accepted, or False if money is insufficient."""
    if total >= MENU[userOption]["cost"]:
        change = total - MENU[userOption]["cost"]
        roundedNumber = round(change, 2)
        print(f"Here is ${roundedNumber} in change.")
        print(f"Here is your {userOption} ☕️. Enjoy!\n")
        resources["money"] += MENU[userOption]["cost"]
        
        for ingredient in MENU[userOption]["ingredients"]:
            resources[ingredient] -= MENU[userOption]["ingredients"][ingredient]
        return True
    else:
        print("Sorry that's not enough money. Money refunded.\n")
        return False


def main():
    machine_state = True

    while machine_state == True:
        option = input("What you want? (espresso, latte, cappuccino): ")
        
        if option == "report":
            report()
        elif option == "espresso":
            print("Espresso\n")
            if checkResource(option) == False:
                continue
            else:
                total = insertCoins()
                checkTransaction(total, option)
        elif option == "latte":
            print("Latte\n")
            if checkResource(option) == False:
                continue
            else:
                total = insertCoins()
                checkTransaction(total, option)
        elif option == "cappuccino":
            print("Cappuccino\n")
            if checkResource(option) == False:
                continue
            else:
                total = insertCoins()
                checkTransaction(total, option)
        elif option == "off":
            machine_state = False
        else:
            print("Wrong Option")


main()



    