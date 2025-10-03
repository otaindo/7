package main

import (
    "errors"
    "fmt"
)

type Product struct {
    ID       int
    Name     string
    Price    float64
    Quantity int
}

type Inventory struct {
    products map[int]*Product
}

func (inv *Inventory) AddProduct(product Product) {
    inv.products[product.ID] = &product
    fmt.Printf("Продукт '%s' успешно добавлен на склад.\n", product.Name)
}

func (inv *Inventory) WriteOff(productID int, quantity int) error {
    product, exists := inv.products[productID]
    if !exists {
        return errors.New("продукт отсутствует на складе")
    }

    if quantity > product.Quantity {
        return errors.New("недостаточно товара на складе для списания")
    }

    product.Quantity -= quantity
    fmt.Printf("%d единиц продукта '%s' списано со склада.\n", quantity, product.Name)
    return nil
}

func (inv *Inventory) RemoveProduct(productID int) error {
    _, exists := inv.products[productID]
    if !exists {
        return errors.New("продукт отсутствует на складе")
    }
    delete(inv.products, productID)
    fmt.Printf("Продукт удалён со склада.\n")
    return nil
}

func (inv *Inventory) GetTotalValue() float64 {
    totalValue := 0.0
    for _, product := range inv.products {
        totalValue += product.Price * float64(product.Quantity)
    }
    return totalValue
}

func main() {
    inventory := Inventory{
        products: make(map[int]*Product),
    }

    inventory.AddProduct(Product{ID: 1, Name: "Хлеб", Price: 50, Quantity: 10})
    inventory.AddProduct(Product{ID: 2, Name: "Молоко", Price: 80, Quantity: 5})

    err := inventory.WriteOff(1, 3)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("Общая стоимость товаров на складе: %.2f руб.\n", inventory.GetTotalValue())

    err = inventory.RemoveProduct(2)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("Новая общая стоимость товаров на складе: %.2f руб.\n", inventory.GetTotalValue())
}
