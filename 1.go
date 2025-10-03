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

func (ba *BankAccount) Deposit(amount float64) {
    if amount <= 0 {
        fmt.Println("Сумма должна быть положительной.")
        return
    }
    ba.balance += amount
    fmt.Printf("Пополнено %.2f рублей.\nНовый баланс: %.2f\n", amount, ba.balance)
}

func (ba *BankAccount) Withdraw(amount float64) error {
    if amount > ba.balance {
        return errors.New("недостаточно средств")
    }
    if amount <= 0 {
        return errors.New("сумма вывода должна быть больше нуля")
    }
    ba.balance -= amount
    fmt.Printf("Снято %.2f рублей.\nНовый баланс: %.2f\n", amount, ba.balance)
    return nil
}

func (ba *BankAccount) GetBalance() float64 {
    return ba.balance
}

func main() {
    account := &BankAccount{
        accountNumber: "12345",
        holderName:    "Иван Иванов",
        balance:  
