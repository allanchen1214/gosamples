package main

import (
	"net/http"

	"github.com/allanchen1214/bingo/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var orders = map[uint64]Order{
	1: Order{OrderID: 1, ProductName: "phone"},
	2: Order{OrderID: 2, ProductName: "car"},
	3: Order{OrderID: 3, ProductName: "food"},
}

type OrderHandler struct {
}

type Order struct {
	OrderID     uint64 `json:"orderid"`
	ProductName string `json:"productName"`
}

func newOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func RegisterOrderRouter(r *gin.Engine) {
	h := newOrderHandler()
	g := r.Group("/order")
	g.GET("/add", h.List)
}

func (h *OrderHandler) List(c *gin.Context) {
	log.Info("list all orders", zap.Any("data", orders))
	c.JSON(http.StatusOK, orders)
}
