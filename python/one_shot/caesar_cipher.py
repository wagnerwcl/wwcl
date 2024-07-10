alphabet = "abcdefghijklmnopqrstuvwxyz"
alphabet_list = list(alphabet)
alphabet_size = len(alphabet)

text = input("Enter the text to be encrypted: ")
shift = int(input("Enter the shift value: "))

def encrypt(text, shift):
    encrypt_text = []

    for letter in text:
        if letter in alphabet_list and alphabet_list.index(letter) + shift < alphabet_size:
            new_letter = alphabet_list[alphabet_list.index(letter) + shift]
            encrypt_text.append(new_letter)
        elif letter in alphabet_list and alphabet_list.index(letter) + shift >= alphabet_size:
            new_letter = alphabet_list[(alphabet_list.index(letter) + shift) - alphabet_size]
            encrypt_text.append(new_letter)
        else:
            print("The letter", letter, "is not in the alphabet")
    
    encrypt_text = "".join(encrypt_text)
    return encrypt_text


def decrypt(text, shift):
    decrypt_text = []

    for letter in text:
        if letter in alphabet_list and alphabet_list.index(letter) - shift >= 0:
            new_letter = alphabet_list[alphabet_list.index(letter) - shift]
            decrypt_text.append(new_letter)
        elif letter in alphabet_list and alphabet_list.index(letter) - shift < 0:
            new_letter = alphabet_list[(alphabet_list.index(letter) - shift) + alphabet_size]
            decrypt_text.append(new_letter)
        else:
            print("The letter", letter, "is not in the alphabet")
    
    decrypt_text = "".join(decrypt_text)
    return decrypt_text

def main():
    option = ''
    while option != 'q':
        option = input("Options \n'e' to encrypt, \n'd' to decrypt \n'r' to reset \n'q' to quit: ")
        if option == 'e':
            encrypted_text = encrypt(text, shift)
            print("The encrypted text is:", encrypted_text)
        elif option == 'd':
            decrypted_text = decrypt(encrypted_text, shift)
            print("The decrypted text is:", decrypted_text)
        # Reset and start again
        elif option == 'r':
            main()
        elif option == 'q':
            print("Goodbye")
            SystemExit
        else:
            print("Invalid option")

if __name__ == "__main__":
    main()

