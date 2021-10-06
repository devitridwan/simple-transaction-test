package service

import (
	"Ridwan/test_sagara/src/model"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
)

func (this *ServiceKanggo) Registrasi(context echo.Context, db *sql.DB) error {
	var req model.ReqRegister
	var res = model.Response{Status: "failed"}
	if err := context.Bind(&req); err != nil {
		return context.JSON(http.StatusInternalServerError, res)
	} else {
		_, err := db.Query("insert into tbl_user (name, email, password, status) values ($1, $2, $3, $4)", req.Name, req.Email,
			fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password))), "user")
		if err != nil {
			res.Error = err.Error()
			return context.JSON(http.StatusInternalServerError, res)
		} else {
			res.Status = "success"
		}
	}
	return context.JSON(http.StatusOK, res)
}

func (this *ServiceKanggo) Login(context echo.Context, db *sql.DB) error {
	var req model.ReqRegister
	var res = model.ResLogin{Response: model.Response{Status: "failed"}}
	if err := context.Bind(&req); err != nil {
		return context.JSON(http.StatusInternalServerError, res)
	} else {
		err := db.QueryRow("select email from tbl_user where email = $1 and password = $2", req.Email,
			fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))).Scan(&res.Email)
		if err != nil {
			res.Error = "user or password wrong"
			return context.JSON(http.StatusInternalServerError, res)
		} else {
			token, err := this.CreateToken(req.Email)
			if err != nil {
				res.Error = err.Error()
				return context.JSON(http.StatusInternalServerError, res)
			} else {
				this.Token[token] = req.Email
				res.Status = "success"
				res.Token = token
			}
		}
	}
	return context.JSON(http.StatusOK, res)
}
