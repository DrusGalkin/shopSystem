package shop_system

type BonusCard struct {
	cardID uint
	money  uint
}

func (b *BonusCard) GetBonusCardID() uint {
	return b.cardID
}
