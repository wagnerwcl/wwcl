from random import randint
from data import *


def generate_questions():
    """Generate question using the dictionary data."""
    random = randint(0, len(questions) - 1)
    question = list(questions.keys())[random]

    return question


def main():
    questionA = generate_questions()
    questionB = generate_questions()
    score = 0
    
    while len(questions) > 1:

        while questionA == questionB:
            questionB = generate_questions()
        
        questionValueA = questions[questionA]
        questionValueB = questions[questionB]
        
        print("#"*40)
        
        print(questionA, ":", questionValueA)
        print(questionB, ":", questionValueB)

        if score > 0:
            print("You right, you score: ", score)

        print("Compare A: ", questionA)
        print(logo_vs)
        print("Compare B: ", questionB)
        print("-"*40)

        choice = input("Who have more points, type A or B: ")

        if choice == "A":
            if questionValueA > questionValueB:
                print("Right Answer")
                
                if questionB in questions:
                    questions.pop(questionB)
                
                questionB = generate_questions()
                score += 1
            else:
                print("Wrong Answer")
                break
        elif choice == "B":
            if questionValueB > questionValueA:
                print("Right Answer")
                
                if questionA in questions:
                    questions.pop(questionA)

                questionA = generate_questions()
                score += 1
            else:
                print("Wrong Answer")
                break
        else:
            print("Game Over")
        
        print(questions)
        print("#"*40+"\n\n")


if __name__ == "__main__":
    main()