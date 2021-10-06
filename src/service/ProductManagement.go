package service

import (
	"Ridwan/test_sagara/src/model"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
)

func (this *ServiceKanggo) UploadProcess(file *multipart.FileHeader, name string) {
	file_target := fmt.Sprintf("%s/%s", this.Properties.TargetPath, name)
	fmt.Println(file_target)
	src, err := file.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer src.Close()
	dst, err := os.Create(file_target)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Upload File Sukses")
}

func (this *ServiceKanggo) CreateProduct(context echo.Context, db *sql.DB) error {
	var req model.ReqCreateProduct
	var res = model.Response{Status: "failed"}
	file, err := context.FormFile("files")

	if err != nil {
		res.Error = err.Error()
		return context.JSON(http.StatusInternalServerError, res)
	}
	// fmt.Println(context.FormFile("files"))
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(form.File["files"])
	// }

	if err := context.Bind(&req); err != nil {
		res.Error = err.Error()
		return context.JSON(http.StatusInternalServerError, res)
	} else {

		token := context.Request().Header.Get("Authorization")
		if val, ok := this.Token[token]; ok {
			toname, err := file.Open()
			if err != nil {
				res.Error = err.Error()
				return context.JSON(http.StatusInternalServerError, res)
			}
			defer toname.Close()
			byteContainer, _ := ioutil.ReadAll(toname)
			format := strings.Split(file.Filename, ".")
			fmt.Println(file.Filename)
			fmt.Println(format)
			name := fmt.Sprintf("%x", sha256.Sum256(byteContainer))
			filename := fmt.Sprintf("%s.%s", name, format[len(format)-1:][0])
			// fmt.Println(file)
			go this.UploadProcess(file, filename)
			var status bool
			err = db.QueryRow("select true from tbl_user where email = $1 and status = $2", val, "admin").Scan(&status)
			fmt.Println(status)
			if err == nil && status {
				_, err := db.Query("insert into tbl_product (name, price, qty, path) values($1, $2, $3, $4)", req.Name,
					req.Price, req.Quantity, filename)
				if err != nil {
					res.Error = err.Error()
					return context.JSON(http.StatusInternalServerError, res)
				} else {
					res.Status = "success"
				}
			}
		}

	}
	return context.JSON(http.StatusOK, res)
}

func (this *ServiceKanggo) EditProduct(context echo.Context, db *sql.DB) error {
	var req model.ReqUpdateProduct
	var res = model.Response{Status: "failed"}
	if err := context.Bind(&req); err != nil {
		return context.JSON(http.StatusInternalServerError, res)
	} else {
		token := context.Request().Header.Get("Authorization")
		if val, ok := this.Token[token]; ok {
			_, err := db.Query("update tbl_product set name=$1, price=$2, qty=$3 where id = $4 and exists (select true from tbl_user where email = $5 and status = $6)",
				req.Name, req.Price, req.Quantity, req.Id, val, "admin")
			if err != nil {
				res.Error = err.Error()
				return context.JSON(http.StatusInternalServerError, res)
			} else {
				res.Status = "success"
			}
		}

	}
	return context.JSON(http.StatusOK, res)
}

func (this *ServiceKanggo) ListProduct(context echo.Context, db *sql.DB) error {
	var res = model.ResProductList{Response: model.Response{Status: "failed"}}
	var data model.DataProductList
	token := context.Request().Header.Get("Authorization")
	if val, ok := this.Token[token]; ok {
		result, err := db.Query("select * from tbl_product where exists (select true from tbl_user where email = $1)", val)
		if err != nil {
			res.Error = err.Error()
			return context.JSON(http.StatusInternalServerError, res)
		} else {
			for result.Next() {
				if err := result.Scan(
					&data.Id,
					&data.Name,
					&data.Price,
					&data.Quantity,
					&data.Path); err != nil {
					res.Error = err.Error()
					return context.JSON(http.StatusInternalServerError, res)
				} else {
					res.Data = append(res.Data, data)
				}
			}
			res.Status = "success"
		}
	}

	// }
	return context.JSON(http.StatusOK, res)
}

func (this *ServiceKanggo) DeleteProduct(context echo.Context, db *sql.DB) error {
	var req model.ReqDeleteProduct
	var res = model.Response{Status: "failed"}
	if err := context.Bind(&req); err != nil {
		return context.JSON(http.StatusInternalServerError, res)
	} else {
		token := context.Request().Header.Get("Authorization")
		if val, ok := this.Token[token]; ok {
			_, err := db.Query("delete from tbl_product where id = $1 and exists (select true from tbl_user where email = $2 and status = $3)",
				req.Id, val, "admin")
			if err != nil {
				res.Error = err.Error()
				return context.JSON(http.StatusInternalServerError, res)
			} else {
				res.Status = "success"
			}
		}

	}
	return context.JSON(http.StatusOK, res)
}
