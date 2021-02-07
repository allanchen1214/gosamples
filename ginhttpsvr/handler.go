package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/allanchen1214/bingo/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var orders []Order = []Order{
	{OrderID: 1, ProductName: "phone"},
	{OrderID: 2, ProductName: "car"},
	{OrderID: 3, ProductName: "food"},
}

type OrderHandler struct {
}

type Order struct {
	OrderID     uint64 `json:"orderID"`
	ProductName string `json:"productName"`
}

func newOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func RegisterOrderRouter(r *gin.Engine) {
	h := newOrderHandler()
	g := r.Group("/order")
	g.GET("/list", h.List)
	g.GET("/item/:id", h.GetByID)
}

func (h *OrderHandler) List(c *gin.Context) {
	log.Info("list all orders", zap.Any("data", orders))
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetByID(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)
	id64 := uint64(id)
	log.Info("get by id req", zap.Uint64("id", id64))
	for _, item := range orders {
		if item.OrderID == id64 {
			log.Info("get by id resp", zap.Any("data", item))
			c.JSON(http.StatusOK, item)
			return
		}
	}
	log.Info("get by id resp", zap.Error(errors.New("NotFound")))
	c.JSON(http.StatusBadRequest, "NotFound")
}
