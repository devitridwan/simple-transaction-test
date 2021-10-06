package main

import (
	"Ridwan/test_sagara/src/info"
	"Ridwan/test_sagara/src/properties"
	"Ridwan/test_sagara/src/service"
	"database/sql"
	"fmt"
	"strconv"

	_ "database/sql/driver"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Ekstration(target ...interface{}) error {
	configPath := "."
	configName := "sagara-endpoint-properties"

	// config file path
	viper.AddConfigPath(configPath)
	// config file name
	viper.SetConfigName(configName)

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	for _, element := range target {
		err = viper.Unmarshal(&element)
		if err != nil {
			return err
		}
	}

	viper.OnConfigChange(func(in fsnotify.Event) {
		for _, element := range target {
			err = viper.Unmarshal(&element)
		}
	})

	return nil
}

func main() {
	info.PrintHeader()
	properties := &properties.EndpointProperties{}
	fmt.Println("trace")
	err := Ekstration(&properties)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		Start(properties)
	}
}

func Start(properties *properties.EndpointProperties) {
	prop := properties
	dbCOnf := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s", prop.Database.Username, prop.Database.Password, prop.Database.Name, prop.Database.Host, prop.Database.Port, prop.Database.SllMode)
	// dbCOnf := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", prop.Database.Username, prop.Database.Password, prop.Database.Host, prop.Database.Port, prop.Database.Name)
	db, err := sql.Open("postgres", dbCOnf)
	if err != nil {
		panic(err.Error())
	} else {
	}
	service := service.ServiceKanggo{}
	service.Init(prop)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE}}))

	e.POST(prop.Topic.GetChannelRest("registrasi", "post"), func(context echo.Context) (err error) { return service.Registrasi(context, db) })
	e.POST(prop.Topic.GetChannelRest("login", "post"), func(context echo.Context) (err error) { return service.Login(context, db) })
	// e.POST(prop.Topic.GetChannelRest("create-product", "post"), func(context echo.Context) (err error) { return service.CreateProduct(context, db) })
	e.POST(prop.Topic.GetChannelRest("create-product", "post"), func(context echo.Context) (err error) {
		return service.CreateProduct(context, db)
	})
	e.POST(prop.Topic.GetChannelRest("edit-product", "post"), func(context echo.Context) (err error) { return service.EditProduct(context, db) })
	e.POST(prop.Topic.GetChannelRest("delete-product", "post"), func(context echo.Context) (err error) { return service.DeleteProduct(context, db) })
	e.GET(prop.Topic.GetChannelRest("list-product", "get"), func(context echo.Context) (err error) { return service.ListProduct(context, db) })
	e.POST(prop.Topic.GetChannelRest("order-transaction", "post"), func(context echo.Context) (err error) { return service.OrderTransaction(context, db) })
	e.POST(prop.Topic.GetChannelRest("payment-transaction", "post"), func(context echo.Context) (err error) { return service.PaymentTransaction(context, db) })
	// e.POST(this.properties.Topic.GetChannelRest("restore", "post"), func(context echo.Context) (err error) {
	// 	form, err := context.MultipartForm()
	// 	go this.processUploadRestoreRequest(form.File["files"], context.FormValue("id_user"))
	// 	return context.JSON(http.StatusOK, `{"status":"success"}`)
	// })
	// this.logger.Error(e.StartTLS(":"+strconv.Itoa(this.properties.Port), this.properties.CertDir+"cert.pem", this.properties.CertDir+"key.pem").Error())
	logger.Error(e.Start(":" + strconv.Itoa(prop.Port)).Error())

}
