package shop

import (
	"errors"
	"fmt"
)

var OverflowException = errors.New("Переполнение склада! Больше на")

type Warehouse struct {
	assortment      map[category][]Product
	maxCountProduct uint
	countProduct    uint
}

func (w *Warehouse) AddAssortment(product *Product) error {
	cat := product.Category
	if _, empty := w.assortment[cat]; !empty {
		w.assortment[cat] = []Product{}
	}

	w.assortment[cat] = append(w.assortment[cat], *product)
	w.countProduct += product.count
	if w.countProduct > w.maxCountProduct {
		return OverflowException
	}
	return nil
}

func (w *Warehouse) GetMaxCounterProduct() uint {
	return w.maxCountProduct
}

func (w *Warehouse) GetCounterProduct() uint {
	return w.countProduct
}

func (w *Warehouse) GetStateWarehouse() {
	fmt.Printf("Заполнен на %d из %d\n", w.countProduct, w.maxCountProduct)
}
