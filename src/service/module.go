package service

import (
	"Ridwan/test_sagara/src/properties"
	"time"

	"github.com/jwt-go"
)

type Order struct {
	ProductId  int
	Qty        int
	TotalPrice int
}

type ServiceKanggo struct {
	Token      map[string]string
	TotalPrice map[string]Order
	Properties *properties.EndpointProperties
}

var JWT_SIGNATURE_KEY = []byte("testing")

func (this *ServiceKanggo) Init(Properties *properties.EndpointProperties) {
	this.Token = make(map[string]string)
	this.TotalPrice = make(map[string]Order)
	this.Properties = Properties
}

func (this *ServiceKanggo) CreateToken(userId string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}
	return token, nil
}
