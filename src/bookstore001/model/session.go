package model

// Session 结构体
type Session struct {
	SessionID string
	UserName  string
	UserID    int
	Cart      *Cart
	Order     *Order
	Orders    []*Order
}
