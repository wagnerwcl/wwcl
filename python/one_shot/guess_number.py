# Number Guessing Game Objectives:

# Include an ASCII art logo.
# Allow the player to submit a guess for a number between 1 and 100.
# Check user's guess against actual answer. Print "Too high." or "Too low." depending on the user's answer. 
# If they got the answer correct, show the actual answer to the player.
# Track the number of turns remaining.
# If they run out of turns, provide feedback to the player. 
# Include two different difficulty levels (e.g., 10 guesses in easy mode, only 5 guesses in hard mode).

from random import randint

logo = """ 
                            _                                        
 _ _ _ _       _      _____| |         _____ _   _     _   _         
| | | | |_ ___| |_   |     |_|_____   |_   _| |_|_|___| |_|_|___ ___ 
| | | |   | .'|  _|  |-   -| |     |    | | |   | |   | '_| |   | . |
|_____|_|_|__,|_|    |_____| |_|_|_|    |_| |_|_|_|_|_|_,_|_|_|_|_  |
                                                                |___|
"""
print(logo + "\n" + "Welcome to the Number Guessing Game!\n\n")



def dificulty(level):
    if level == 1:
        turns = 10
    elif level == 2:
        turns = 5
    return turns


def check_guess(user_guess, number):
    if user_guess < number:
        print("Too Low")
    elif user_guess > number:
        print("Too Hight")
    elif user_guess == number:
        print(f"You Win, the right number is {number}")
        is_game_over = True
        return is_game_over


def main():
    is_game_over = False
    number = randint(1,100)
    turns = 0
    
    level = int(input("Choose a level 1:Chicken and 2:Beast: "))
    turns = dificulty(level)

    while not is_game_over:
        print(f"You have {turns} turns left.\n")
        user_guess = int(input("Choose a number between 1 and 100: "))
        
        is_game_over = check_guess(user_guess, number)
        turns -= 1
        
        if turns == 0:
            print(f"Sorry, you run out of turns.\nThe right number is {number}")
            is_game_over = True


main()