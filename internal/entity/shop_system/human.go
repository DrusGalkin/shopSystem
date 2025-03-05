package shop_system

import (
	"errors"
	"fmt"
)

var NotFoundShopException = errors.New("Данное место не является магазином!")
var NotFoundPlaceException = errors.New("Вы не находитесь в месте, для которого предоставленна данная функция")
var NotEnoughMoney = errors.New("Пополните баланс! Недостаточно средств")

type Human[T any] struct {
	Name      string
	Money     uint
	Places    *Places[T]
	Inventory Inventory[T]
	Bonus     BonusCardsSystem
}

type Places[T any] struct {
	place T
}

type Inventory[T any] struct {
	Tile T
}

func NewHuman(name string, money uint) *Human[any] {
	return &Human[any]{
		Name:  name,
		Money: money,
	}
}

func (h *Human[T]) GoToShop(s T) error {
	h.Places = &Places[T]{place: s}
	shop, ok := any(s).(*Shop[T])
	if !ok {
		return NotFoundPlaceException
	}
	fmt.Printf("%s пошел в магазин %s", h.Name, shop.GetName())
	shop.PushUsers(any(h).(T))

	return nil
}

func (h *Human[T]) TakeCart() error {
	shop, ok := any(h.Places.place).(*Shop[T])
	if !ok {
		return NotFoundShopException
	}

	cart, err := shop.GetCart()
	if err != nil {
		fmt.Println(err)
	}
	h.Inventory.Tile = any(cart).(T)
	fmt.Println(any(cart.String()).(T))

	return nil
}

func (h *Human[T]) TakeProduct(nameProduct string, quantity uint) error {
	shop, ok := any(h.Places.place).(*Shop[T])
	if !ok {
		return NotFoundShopException
	}

	if any(h.Inventory.Tile).(*ShoppingCarts[any]) == nil {
		return NotFountInventoryCart
	}

	cart, ok := any(h.Inventory.Tile).(*ShoppingCarts[any])
	if !ok {
		return NotFountInventoryCart
	}

	for i := 0; i < int(quantity); i++ {
		product, err := shop.FindProductByName(nameProduct)
		if err != nil {
			fmt.Println(err)
		}
		cart.PutInBasket(*product)
	}
	return nil
}

func (h *Human[T]) BuyProducts(useBonus bool) error {
	shop, ok := any(h.Places.place).(Shop[T])
	if !ok {
		return NotFoundPlaceException
	}

	cost := shop.CashRegister(*h)

	if useBonus {
		cost -= h.Bonus.GetBalance()
	}

	finalSum := float64(h.Money) - cost
	if finalSum < 0 {
		return NotEnoughMoney
	}

	return nil
}
