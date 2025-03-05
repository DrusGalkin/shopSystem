package app

import (
	"StrongPakage/internal/entity/shop_system"
	"fmt"
)

func Run() {
	shop := shop_system.CreateShop("Магнит")
	err := shop.AddProduct(
		shop_system.NewProduct("Огурец", 100, 60, "Овощи"),
		shop_system.NewProduct("Картошка", 60, 100, "Овощи"),
		shop_system.NewProduct("Морковь", 50, 40, "Овощи"),
		shop_system.NewProduct("Бананы", 150, 90, "Фрукты"),
		shop_system.NewProduct("Арбуз", 80, 120, "Фрукты"),
	)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(shop_system.AllProducts())
	shop.GetWarehouse().GetStateWarehouse()

	h := shop_system.NewHuman("Andrew", 10_000)

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
