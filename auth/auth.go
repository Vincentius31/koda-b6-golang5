package auth

import (
	"crypto/md5"
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type System struct {
	Users []User
}

// Manual Hex Conversion
func ToHexString(bytes [16]byte) string {
	const hexChars = "0123456789abcdef"
	res := make([]byte, 32)
	for i, b := range bytes {
		res[i*2] = hexChars[b>>4]
		res[i*2+1] = hexChars[b&0x0f]
	}
	return string(res)
}

func EncryptMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return ToHexString(hash)
}

func WaitEnter() {
	fmt.Scanln()
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func HandleSystemError() {
	if r := recover(); r != nil {
		fmt.Printf("\n[System Error]: %v. Memulihkan sistem...\n", r)
		WaitEnter()
	}
}

func (s *System) Register() {
	for {
		var fName, lName, email, pass, confirm, choice string

		ClearScreen()
		fmt.Println("--- Register ---")
		fmt.Print("What is your first name: ")
		fmt.Scanln(&fName)
		fmt.Print("What is your last name: ")
		fmt.Scanln(&lName)
		fmt.Print("What is your email: ")
		fmt.Scanln(&email)
		fmt.Print("Enter a strong password: ")
		fmt.Scanln(&pass)
		fmt.Print("Confirm your password: ")
		fmt.Scanln(&confirm)

		fmt.Println("\nIs it true?")
		fmt.Printf("First name: %s\n", fName)
		fmt.Printf("Last name: %s\n", lName)
		fmt.Printf("Email: %s\n", email)
		fmt.Print("Continue (y/n): ")
		fmt.Scanln(&choice)

		if choice == "y" || choice == "Y" {
			if pass != confirm {
				fmt.Println("Password does not match, press enter to restart..")
				WaitEnter()
				continue
			}
			newUser := User{
				FirstName: fName,
				LastName:  lName,
				Email:     email,
				Password:  EncryptMD5(pass),
			}
			s.Users = append(s.Users, newUser)
			fmt.Print("Register success, press enter to back..")
			WaitEnter()
			break
		}
	}
}

func (s *System) Login() *User {
	for {
		ClearScreen()
		var email, pass string
		fmt.Println("--- Login ---")
		fmt.Print("Enter your email: ")
		fmt.Scanln(&email)
		fmt.Print("Enter your password: ")
		fmt.Scanln(&pass)

		hashed := EncryptMD5(pass)
		for _, u := range s.Users {
			if u.Email == email && u.Password == hashed {
				fmt.Print("Login success, press enter to back..")
				WaitEnter()
				return &u
			}
		}
		fmt.Print("Wrong email or password, press enter to restart")
		WaitEnter()
	}
}

func (s *System) ForgotPassword() {
	for {
		ClearScreen()
		var email, pass, confirm string
		fmt.Println("--- Forgot Password ---")
		fmt.Print("Enter your email: ")
		fmt.Scanln(&email)
		fmt.Print("Enter a strong password: ")
		fmt.Scanln(&pass)
		fmt.Print("Confirm your password: ")
		fmt.Scanln(&confirm)

		found := -1
		for i, u := range s.Users {
			if u.Email == email {
				found = i
				break
			}
		}

		if found == -1 {
			fmt.Print("Email not found, press enter to restart")
			WaitEnter()
			continue
		}
		if pass != confirm {
			fmt.Print("Password mismatch, press enter to restart")
			WaitEnter()
			continue
		}
		s.Users[found].Password = EncryptMD5(pass)
		fmt.Print("Password changed, press enter to back")
		WaitEnter()
		break
	}
}
