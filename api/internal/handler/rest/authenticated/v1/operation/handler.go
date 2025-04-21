package operation

import (
	"github.com/nhan1603/CryptographicAssignment/api/internal/controller/menus"
	"github.com/nhan1603/CryptographicAssignment/api/internal/controller/orders"
)

// Handler is the web handler for this pkg
type Handler struct {
	menuCtrl  menus.Controller
	orderCtrl orders.Controller
}

// New instantiates a new Handler and returns it
func New(menuCtrl menus.Controller, orderCtrl orders.Controller) Handler {
	return Handler{menuCtrl: menuCtrl, orderCtrl: orderCtrl}
}
