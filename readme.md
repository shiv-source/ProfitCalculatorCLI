# Profit Calculator CLI

A simple command-line interface (CLI) tool to calculate profit after deducting expenses and taxes. This tool takes user inputs for revenue, expenses, and tax rate to compute earnings before tax, profit after tax, and the profit percentage.

---

## Features

- Accepts user inputs for:
  - Revenue
  - Expenses
  - Tax Rate (percentage)
- Calculates the following:
  - **Earnings Before Tax (EBT)**: `Revenue - Expenses`
  - **Profit After Tax**: `EBT - Taxes`
  - **Profit Percentage**: `(Profit After Tax / Revenue) × 100`

---

## Requirements

- **Go** (if running directly with `go run`): Ensure Go is installed on your system.

---

## How to Run

1. Clone the repository:

   ```bash
   git clone https://github.com/shiv-source/ProfitCalculatorCLI.git
   cd ProfitCalculatorCLI
   ```

2. Use the `make` command to run the application without requiring Go installation:

   ```bash
   make run
   ```

   Alternatively, if Go is installed, you can run the application directly:

   ```bash
   make start
   ```

3. Follow the prompts to input the required data:

   ```
   Enter Revenue: <value>
   Enter Expenses: <value>
   Enter Tax Rate (in %): <value>
   ```

4. The tool will calculate and display:
   ```
   Earnings Before Tax (EBT): <calculated-value>
   Profit After Tax: <calculated-value>
   Profit Percentage: <calculated-value>%
   ```

---

## License

This project is licensed under the MIT License. See the [MIT License](LICENSE) file for details.

---
