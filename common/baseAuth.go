package common

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//MySigningKey := []byte("kh+khha%2b(p9(uw+^*#x%o51*hrsohxi&xtdklm2tfxx5mq_h")

type TokenStruct struct {
	Username string `json:"username"`
	Exp      int    `json:"exp"`
}

func GetJwtAuth(name string) (ss string, err error) {
	claims := &jwt.MapClaims{
		"username": name,
		"exp":      int64(time.Now().Unix() + 43200),
		//"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(beego.AppConfig.String("MySigningKey"))
	ss, err = token.SignedString([]byte(beego.AppConfig.String("MySigningKey")))

	return
}

func GetJwtName(ss string) (name string, err error) {
	t, err := jwt.Parse(ss, func(*jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("MySigningKey")), nil
	})

	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return
	}
	fmt.Println("还原后的token信息claims部分:", t.Claims)
	claims, _ := t.Claims.(jwt.MapClaims)
	name = claims["username"].(string)
	fmt.Println(name)
	return
}
func main() {
	mySigningKey := []byte("kh+khha%2b(p9(uw+^*#x%o51*hrsohxi&xtdklm2tfxx5mq_h")
	// Create the Claims
	claims := &jwt.MapClaims{
		"username": "lqm",
		"exp":      int64(time.Now().Unix() + 43200),
		//"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	}
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println("签名后的token信息:", ss)
	t, err := jwt.Parse(ss, func(*jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return
	}
	fmt.Println("还原后的token信息claims部分:", t.Claims)
	cc, erro := GetJwtAuth("lqm")
	fmt.Println(cc)
	GetJwtName(cc)
	fmt.Println(erro)
}
