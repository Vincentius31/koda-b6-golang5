# Go Interactive Authentication System

A terminal-based authentication flow built with **Go (Golang)**. This project demonstrates modular programming, user management, and basic security practices like password hashing without relying on external dependencies.

---

## 🚀 Features

- **User Registration**: Create an account with first name, last name, email, and password.
- **Secure Login**: Authentication system that verifies credentials against stored data.
- **Password Management**: "Forgot Password" functionality to reset credentials for existing users.
- **User Directory**: View a list of all registered users (accessible after login).
- **Security**: Password storage using **MD5 Hashing** with a custom Hexadecimal conversion implementation.
- **Stability**: Implementation of `recover` to prevent the system from crashing on unexpected errors.

---

## 🛠️ Technical Implementation

This project was built to practice core Go concepts, including:

* **Modules & Packages**: Logic is separated into the `auth` package for better maintainability.
* **Structs & Pointers**: Utilized to manage user data efficiently in memory.
* **Defer, Panic, & Recover**: Used for graceful error handling and resource cleanup.
* **Zero-Dependency Input**: Handling CLI input using `fmt.Scanln` and manual buffer management.
* **Manual Hex Conversion**: Converting MD5 byte arrays to Hex strings manually to understand low-level data manipulation without `encoding/hex`.

---

## 📂 Project Structure

```text
.
├── go.mod          # Go module definition
├── main.go         # Application entry point & menu navigation
└── auth/
    └── auth.go     # Core logic: Registration, Login, & Encryption

```

---

## ⚙️ How to Run

1. Open your terminal and navigate to the project directory.

2. Initialize the module (if you haven't already):

```text
go mod init koda-b6-golang5
```

3. Run the application
```
go run main.go
```

---

## 📝 Usage Flow

1. Register: Enter your details. The system will prompt for confirmation before saving. Your password is encrypted immediately using MD5.

2. Login: Enter your email and password. The system hashes your input and compares it with the stored record.

3. Forgot Password: If you forget your password, you can reset it by providing your registered email.

4. Home View: Once logged in, select "List All Users" to see the database, where you can verify that passwords are stored as encrypted hashes.





