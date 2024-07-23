from menu import Menu, MenuItem
from cofee_maker import CoffeeMaker
from money_machine import MoneyMachine

cm = CoffeeMaker()
menu = Menu()
#menu_item = MenuItem()
mm = MoneyMachine()


def main():
    machine_state = True

    while machine_state == True:
        options = menu.get_items()
        choice = input(f"What you want? {options} : ")

        match choice:
            case "report":
                cm.report()
                mm.report()

            case "off":
                machine_state = False

            case _:
                drink = menu.find_drink(choice)
                if cm.is_resource_sufficient(drink):
                    if mm.make_payment(drink.cost):
                        cm.make_coffee(drink)


main()