package product

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Id       int
	Name     string
	Price    int
	Quantity int
}

type AddParam struct {
	Name     string
	Price    int
	Quantity int
}

func ReadProductData() ([]Product, error) {
	rawText, err := os.ReadFile("products.txt")
	if err != nil {
		return nil, err
	}
	// HANDLE EMPTY FILE
	if string(rawText) == "" {
		return []Product{}, nil
	}
	arrText := strings.Split(string(rawText), ";\n")
	datas := []Product{}
	// LAST ITEM IN datas GOING TO BE "" BECAUSE OF NEWLINE IN THE END OF FILE
	for i := 0; i < len(arrText)-1; i++ {
		data := strings.Split(arrText[i], ",")
		id, _ := strconv.Atoi(data[0])
		name := data[1]
		price, _ := strconv.Atoi(data[2])
		quantity, _ := strconv.Atoi(data[3])
		datas = append(datas, Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Quantity: quantity,
		})
	}
	return datas, nil
}

func FindProductById(id int) (Product, error) {
	data, err := ReadProductData()
	if err != nil {
		return Product{}, err
	}
	for _, item := range data {
		if item.Id == id {
			return item, nil
		}
	}
	return Product{}, errors.New("Product not found.")
}

func PrintAllProduct() error {
	data, err := ReadProductData()
	if err != nil {
		return err
	}
	if len(data) == 0 {
		fmt.Println("No Product Registered.")
		return nil
	}
	fmt.Println("| ID | Product Name | Price | Quantity |")
	for _, item := range data {
		fmt.Printf("| %d | %s | %d | %d |\n", item.Id, item.Name, item.Price, item.Quantity)
	}
	return nil
}

func AddProduct(param AddParam) error {
	if param.Name == "" {
		return errors.New("Name shouldn't be empty.")
	}
	if param.Price < 0 {
		return errors.New("Price should be minimal 0.")
	}
	if param.Quantity < 0 {
		return errors.New("Quantity should be minimal 0.")
	}
	data, err := ReadProductData()
	if err != nil {
		return err
	}
	nextId := 1
	if len(data) != 0 {
		nextId = data[len(data)-1].Id + 1
	}
	text := fmt.Sprintf("%d,%s,%d,%d;\n", nextId, param.Name, param.Price, param.Quantity)
	file, err := os.OpenFile("products.txt", os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, errWrite := file.WriteString(text)
	if errWrite != nil {
		return errWrite
	}
	return nil
}

func UpdateDataById(param Product) error {
	if param.Name == "" {
		return errors.New("Name shouldn't be empty.")
	}
	if param.Price < 0 {
		return errors.New("Price should be minimal 0.")
	}
	if param.Quantity < 0 {
		return errors.New("Quantity should be minimal 0.")
	}
	data, err := ReadProductData()
	if err != nil {
		return err
	}
	text := ""
	for _, item := range data {
		if param.Id == item.Id {
			text += fmt.Sprintf("%d,%s,%d,%d;\n", item.Id, param.Name, param.Price, param.Quantity)
		} else {
			text += fmt.Sprintf("%d,%s,%d,%d;\n", item.Id, item.Name, item.Price, item.Quantity)
		}
	}
	file, err := os.OpenFile("products.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, errWrite := file.WriteString(text)
	if errWrite != nil {
		return errWrite
	}
	return nil
}

func DeleteProduct(id int) error {
	data, err := ReadProductData()
	if err != nil {
		return err
	}
	text := ""
	for _, item := range data {
		if item.Id == id {
			continue
		}
		text += fmt.Sprintf("%d,%s,%d,%d;\n", item.Id, item.Name, item.Price, item.Quantity)
	}
	if text == "" {
		return errors.New("ID not registered.")
	}
	file, err := os.OpenFile("products.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, errWrite := file.WriteString(text)
	if errWrite != nil {
		return errWrite
	}
	return nil
}
