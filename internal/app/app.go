package app

import (
	"StrongPakage/internal/entity/shop"
	"StrongPakage/internal/entity/shop/common"
	"fmt"
)

func Run() {
	shop := shop.CreateShop("Магнит")
	err := shop.AddProduct(
		shop.NewProduct("Огурец", 100, 60, "Овощи"),
		shop.NewProduct("Картошка", 60, 100, "Овощи"),
		shop.NewProduct("Морковь", 50, 40, "Овощи"),
		shop.NewProduct("Бананы", 150, 90, "Фрукты"),
		shop.NewProduct("Арбуз", 80, 120, "Фрукты"),
	)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(shop.AllProducts())
	shop.GetWarehouse().GetStateWarehouse()

	h := common.NewHuman("Andrew", 10_000)

	if err2 := h.GoToShop(shop); err2 != nil {
		fmt.Println(err2)
	}
	err3 := h.TakeCart()
	if err3 != nil {
		fmt.Println(err)
	}

	h.TakeProduct("Огурец", 5)
	fmt.Println(h.Inventory.Tile)

	shop.GetWarehouse().GetStateWarehouse()
}
