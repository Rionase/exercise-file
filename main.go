package main

import (
	product "exercise-file/Product"
	"exercise-file/helper"
	"fmt"
)

func main() {
	var input string
	var newProductName string
	var newProductPrice, newProductQuantity int
	var id int
	for {
		helper.ClearScreen()
		fmt.Println("PRODUCT DATA MENU\n=================")
		fmt.Println("1. See All Products\n2. Add New Product\n3. Edit Product Data\n4. Delete Product Data\n5. Exit")
		fmt.Print("Input : ")
		helper.InputString(&input)
		switch input {

		case "1":
			helper.ClearScreen()
			fmt.Println("All Product Data\n================")
			err := product.PrintAllProduct()
			if err != nil {
				fmt.Println("!!!", err)
			}
			helper.WaitEnter()

		case "2":
			helper.ClearScreen()
			fmt.Println("Add New Product\n===============")
			fmt.Print("Product Name \t : ")
			helper.InputString(&newProductName)
			fmt.Print("Product Price \t : ")
			helper.InputInt(&newProductPrice)
			fmt.Print("Product Quantity : ")
			helper.InputInt(&newProductQuantity)
			err := product.AddProduct(product.AddParam{
				Name:     newProductName,
				Price:    newProductPrice,
				Quantity: newProductQuantity,
			})
			if err != nil {
				fmt.Println("!!!", err)
				helper.WaitEnter()
				continue
			}
			fmt.Println("New Data Added Successfully.")
			helper.WaitEnter()
		case "3":
			helper.ClearScreen()
			fmt.Println("All Product Data\n================")
			err := product.PrintAllProduct()
			if err != nil {
				fmt.Println("!!!", err)
			}
			fmt.Print("Edit data with ID : ")
			helper.InputInt(&id)
			item, err := product.FindProductById(id)
			if err != nil {
				fmt.Println("!!!", err)
				helper.WaitEnter()
				continue
			}
			helper.ClearScreen()
			fmt.Println("Edit data\n=========")
			fmt.Printf("Old Name \t : %s\nOld Price \t : %d\nOld Quantity \t : %d\n\n", item.Name, item.Price, item.Quantity)
			fmt.Print("New Name \t : ")
			helper.InputString(&newProductName)
			fmt.Print("New Price \t : ")
			helper.InputInt(&newProductPrice)
			fmt.Print("New Quantity \t : ")
			helper.InputInt(&newProductQuantity)
			updateErr := product.UpdateDataById(product.Product{
				Id:       item.Id,
				Name:     newProductName,
				Price:    newProductPrice,
				Quantity: newProductQuantity,
			})
			if updateErr != nil {
				fmt.Println("!!!", updateErr)
				helper.WaitEnter()
				continue
			}
			fmt.Println("Data successfully updated.")
			helper.WaitEnter()
		case "4":
			helper.ClearScreen()
			fmt.Println("All Product Data\n================")
			err := product.PrintAllProduct()
			if err != nil {
				fmt.Println("!!!", err)
			}
			fmt.Print("Delete data with ID : ")
			helper.InputInt(&id)
			errDelete := product.DeleteProduct(id)
			if errDelete != nil {
				fmt.Println("!!!", err)
				helper.WaitEnter()
				continue
			}
			fmt.Println("Date deleted successfully.")
			helper.WaitEnter()
		case "5":
			return
		default:
			fmt.Println("!!! Invalid input.")
			helper.WaitEnter()
		}
	}
}
