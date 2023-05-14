package domain

type Payment struct {
	Amount  uint32
	OrderID string
	Sign    string
}
