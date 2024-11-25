
# Bank Teller API

[![Go](https://img.shields.io/badge/Go-1.23.3-brightgreen)](https://golang.org/dl/)  


## Overview
**Bank Teller API** is a Go-based application designed to manage bank transactions, including login, logout, and payment processing.

---

## Features
- **Feature 1**: User login and logout functionality
- **Feature 2**: Create new payment records
- **Feature 3**: Update payment status (Paid/Unpaid)
- **Feature 4**: Authentication middleware to secure payment routes

---

## Table of Contents
1. [Getting Started](#getting-started)
2. [Technologies Used](#technologies-used)
3. [Setup and Installation](#setup-and-installation)
4. [Usage](#usage)
5. [Endpoints](#endpoints)
6. [Contact](#contact)

---

## Getting Started
To get started with this project, ensure the following tools are installed on your system:
- [Go 1.23.3+](https://golang.org/dl/)

---

## Technologies Used
- **Framework**: Go with Gorilla Mux
- **Build Tool**: Go modules



---

To handle the `SECRET_KEY` environment variable for your application, you can update the `README` to include instructions for setting the environment variable before running the application. Here’s how to modify the section:

---

## Setup and Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Rafirvan/bank-teller-api
   cd bank-teller-api
   ```

2. Set the `SECRET_KEY` environment variable:
    - On **Linux/macOS**:
      ```bash
      export SECRET_KEY=your_secret_key_value
      ```

    - On **Windows** (PowerShell):
      ```powershell
      $env:SECRET_KEY="your_secret_key_value"
      ```

3. Build the project:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

---



## Usage
1. Access the application at:
   ```
   http://localhost:9090
   ```

2. Default login credentials:
    - `username: admin` / `password: password`

---

## Endpoints
Here’s a list of key API endpoints:

### Public Endpoints

| Method | Endpoint               | Description              |
|--------|------------------------|--------------------------|
| POST   | `/login`                | User login.              |

### Authenticated Endpoints

| Method | Endpoint                           | Description                                     |
|--------|------------------------------------|-------------------------------------------------|
| POST   | `/logout`                          | User logout.                                    |
| POST   | `/payment`                         | Create a new payment record.                   |
| PATCH  | `/payment/{id}`                    | Update payment status by ID (mark as Paid/Unpaid). |

---



## Contact
For inquiries or support, please reach out:
- **Name**: Rafirvan
- **Email**: rafirvan@gmail.com

---
