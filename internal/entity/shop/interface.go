package shop

type BonusCardsSystem interface {
	AccrualBonus(sum float64)
	SetPercent(per float64)
	GetBalance() float64
}

type MagnetBonusCart struct {
	balance float64
	percent float64
}

func (m *MagnetBonusCart) AccrualBonus(sum float64) {
	m.balance = sum * m.percent
}

func (m *MagnetBonusCart) SetPercent(per float64) {
	m.percent = per
}

func (m *MagnetBonusCart) GetBalance() float64 {
	return m.balance
}
