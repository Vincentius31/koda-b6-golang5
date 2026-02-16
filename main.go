package main

import (
	"koda-b6-golang5/auth"
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("\nApplication closed.")
	system := &auth.System{Users: []auth.User{}}

	for {
		auth.ClearScreen()
		fmt.Println("--- Welcome to system ---")
		fmt.Println("1. Register\n2. Login\n3. Forgot Password\n\n0. Exit")
		fmt.Print("\nChoose a menu: ")

		var menu int
		fmt.Scanln(&menu)

		switch menu {
		case 1:
			system.Register()
		case 2:
			if len(system.Users) == 0 {
				fmt.Print("No users registered yet, press enter..")
				auth.WaitEnter()
				continue
			}
			user := system.Login()
			if user != nil {
				loggedInMenu(system, user)
			}
		case 3:
			system.ForgotPassword()
		case 0:
			os.Exit(0)
		}
	}
}

func loggedInMenu(s *auth.System, user *auth.User) {
	defer auth.HandleSystemError()
	for {
		auth.ClearScreen()
		fmt.Printf("--- Welcome to system ---\n\nHello \"%s\"\n", user.Email)
		fmt.Println("1. List All Users\n2. Logout\n\n0. Exit")
		fmt.Print("\nChoose a menu: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			auth.ClearScreen()
			for i, u := range s.Users {
				fmt.Printf("%d.\nFull Name: %s %s\nEmail: %s\nPassword: %s\n\n", 
					i+1, u.FirstName, u.LastName, u.Email, u.Password)
			}
			fmt.Print("Press enter to back")
			auth.WaitEnter()
		case 2:
			fmt.Print("Logout success, press enter to back..")
			auth.WaitEnter()
			return
		case 0:
			os.Exit(0)
		}
	}
}