package service

import (
	"Ridwan/test_sagara/src/model"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func (this *ServiceKanggo) OrderTransaction(context echo.Context, db *sql.DB) error {
	var req model.ReqOrder
	var res = model.ResOrder{Response: model.Response{Status: "failed"}}
	if err := context.Bind(&req); err != nil {
		return context.JSON(http.StatusInternalServerError, res)
	} else {
		token := context.Request().Header.Get("Authorization")
		if val, ok := this.Token[token]; ok {
			timeNow := strconv.Itoa(int(time.Now().UnixNano() / int64(time.Millisecond)))
			orderId := fmt.Sprintf("%x", sha256.Sum256([]byte(timeNow+":"+val)))
			var status bool
			var price int
			err := db.QueryRow("select true, price*$1 from tbl_product where qty >= $2 and id = $3", req.Amount, req.Amount, req.ProductId).Scan(&status, &price)
			if err == nil && status {
				_, err := db.Query("insert into tbl_transaksi (order_id, user_id, product_id, amount, status) values($1, $2, $3, 44, $5) ", orderId,
					val, req.ProductId, req.Amount, "pending")
				if err != nil {
					return context.JSON(http.StatusInternalServerError, res)
				} else {
					res.Total = price
					res.OrderId = orderId
					res.Status = "success"
					this.TotalPrice[orderId] = Order{ProductId: req.ProductId, Qty: req.Amount, TotalPrice: price}
				}
			}
		}

	}
	return context.JSON(http.StatusOK, res)
}

func (this *ServiceKanggo) PaymentTransaction(context echo.Context, db *sql.DB) error {
	var req model.ReqPayment
	var res = model.Response{Status: "failed"}
	if err := context.Bind(&req); err != nil {
		return context.JSON(http.StatusInternalServerError, res)
	} else {
		token := context.Request().Header.Get("Authorization")
		if val, ok := this.Token[token]; ok {
			if val == req.Email {
				if val1, ok := this.TotalPrice[req.OrderId]; ok {
					if val1.TotalPrice == req.TotPrice {
						_, err := db.Query("update tbl_transaksi set status = $1 where order_id = $2", "paid", req.OrderId)
						if err != nil {
							return context.JSON(http.StatusInternalServerError, res)
						} else {
							_, err := db.Query("update tbl_product set qty = qty-$1 where id = $2", val1.Qty, val1.ProductId)
							if err != nil {
								return context.JSON(http.StatusInternalServerError, res)
							} else {
								res.Status = "success"
							}
						}
					}
				}
			}
		}
	}
	return context.JSON(http.StatusOK, res)
}
