Your `README.md` is mostly well-structured but has some formatting issues and missing Markdown syntax for better readability. Here's a corrected and polished version:

---

```markdown
# Payment Processing System

A **Go-based Payment Processing System** that supports multiple payment gateways like PayPal and Stripe. This project includes unit tests and follows clean code principles with a focus on extensibility and maintainability.

---

## Project Description

This project simulates a payment processing system with the following features:
- Validate payment details such as amount, currency, and method.
- Process payments via supported gateways:
  - **PayPal**
  - **Stripe**
- A modular and extensible design for adding new payment gateways.
- Comprehensive unit tests using the `testify` library.

---

## Setup Instructions

### Prerequisites
- **Go** installed on your system ([Download Go](https://golang.org/dl)).
- A terminal or IDE like Visual Studio Code for running the project.

### Clone the Repository
```bash
git clone https://github.com/mostafashekari/payment-processing.git
cd payment-processing
```

### Install Dependencies
Use `go mod` to install required dependencies:
```bash
go mod tidy
```

### Run the Application
To run the application:
```bash
go run main.go
```

### Run Tests
To execute unit tests:
```bash
go test ./payment -v
```

---

## Usage Examples

### Example 1: Valid PayPal Payment
```bash
Input:
Amount:   100.00
Currency: USD
Method:   paypal

Output:
Payment processed successfully!
```

### Example 2: Invalid Payment Method
```bash
Input:
Amount:   100.00
Currency: USD
Method:   cash

Output:
Error: unsupported payment method
```

---

## Future Enhancements
- Add more payment gateways (e.g., ApplePay, GooglePay).
- Integrate a database for persisting payment details.
- Add an API layer using gRPC or HTTP for external access.
- Implement logging for better observability.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
```

---

### **Changes Made**
1. Fixed **duplicated title**.
2. Corrected broken formatting in the "Setup Instructions" section:
   - Fixed indentation and added backticks for code blocks.
   - Added missing section headers.
3. Organized "Usage Examples" with proper code formatting.
4. Improved readability with consistent spacing.
