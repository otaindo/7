package main

import (
    "errors"
    "fmt"
)

type BankAccount struct {
    accountNumber string
    holderName    string
    balance       float64
}

func (account *BankAccount) Deposit(amount float64) {
    account.balance += amount
}

func (account *BankAccount) Withdraw(amount float64) error {
    if amount > account.balance {
        return errors.New("Недостаточно средств на счете")
    }
    account.balance -= amount
    return nil
}

func (account *BankAccount) GetBalance() float64 {
    return account.balance
}

func main() {
    account := BankAccount{
        accountNumber: "123456789",
        holderName:    "Иван Иванов",
        balance:       1000.0,
    }

    fmt.Printf("Баланс счета: %.2f руб.\n", account.GetBalance())

    account.Deposit(500.0)
    fmt.Printf("Баланс после депозита: %.2f руб.\n", account.GetBalance())

    err := account.Withdraw(1200.0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Баланс после снятия: %.2f руб.\n", account.GetBalance())
    }


}
