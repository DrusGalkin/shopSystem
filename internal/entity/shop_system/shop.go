package shop_system

import (
	"errors"
	"fmt"
)

var NotFoundShopCart = errors.New("Тележки закончились")
var NotFoundProduct = errors.New("Товар не найден по данной категории")
var NotFountInventoryCart = errors.New("У вас нет тележки для покупок!")

type Shop[T any] struct {
	name          string
	open          bool
	warehouse     *Warehouse
	bonusCard     BonusCardsSystem
	shoppingCarts map[int]*ShoppingCarts[T]
	UserList      []T
	maxCarts      uint
	useCarts      uint
}

type ShoppingCarts[T any] struct {
	id       uint
	shop     *Shop[T]
	count    uint
	products []Product
}

func CreateShop(name string) *Shop[any] {
	s := &Shop[any]{
		name:          name,
		maxCarts:      30,
		shoppingCarts: make(map[int]*ShoppingCarts[any]),
		bonusCard: &MagnetBonusCart{
			balance: 0,
			percent: 0.2,
		},
	}
	s.warehouse = &Warehouse{
		assortment:      make(map[category][]Product),
		maxCountProduct: 10_000,
	}
	return s
}

func (s *Shop[T]) GetCart() (*ShoppingCarts[T], error) {

	if s.useCarts >= s.maxCarts {
		return nil, NotFoundShopCart
	}

	s.useCarts++
	newCart := &ShoppingCarts[T]{
		id:       s.useCarts,
		shop:     s,
		count:    0,
		products: []Product{},
	}

	s.shoppingCarts[int(s.useCarts)] = newCart
	return newCart, nil
}

func (s Shop[T]) FindProductByName(name string) (*Product, error) {
	for _, el := range s.warehouse.assortment {
		for _, item := range el {
			if item.name == name {
				s.warehouse.countProduct--
				item.count--
				return &item, nil
			}
		}
	}
	return nil, NotFoundProduct
}

func (s *Shop[any]) AddProduct(product ...*Product) error {
	for _, el := range product {
		if err := s.warehouse.AddAssortment(el); err != nil {
			err = fmt.Errorf("%w %d из %d", err,
				s.warehouse.countProduct-s.warehouse.maxCountProduct,
				s.warehouse.GetMaxCounterProduct())
			return err
		}
	}
	return nil
}

func (s *Shop[T]) CashRegister(h Human[T]) float64 {
	var cost float64
	cart := any(h.Inventory.Tile).(ShoppingCarts[T])
	for _, el := range cart.products {
		cost += float64(el.price)
	}

	if h.Bonus == nil {
		h.Bonus = s.bonusCard
	}
	return cost
}

func (s *Shop[T]) PushUsers(h T) {
	s.UserList = append(s.UserList, h)
}

func (s Shop[any]) GetName() string {
	return s.name
}

func (s *Shop[any]) AllProducts() *map[category][]Product {
	return &s.warehouse.assortment
}

func (s *Shop[any]) GetWarehouse() *Warehouse {
	return s.warehouse
}

func (s *Shop[any]) OpenShop() {
	s.open = true
}

func (s *Shop[any]) CloseShop() {
	s.open = false
}

func (c *ShoppingCarts[T]) String() string {
	return fmt.Sprintf("Магазин: %s\nID: %d\nЗаполненность: %d\nТовары: %s", c.shop.GetName(), c.id, len(c.products), c.products)
}

func (c *ShoppingCarts[T]) PutInBasket(product Product) {
	c.products = append(c.products, product)
}
