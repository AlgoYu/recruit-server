package controllers

import (
	"anydevelop.cn/recruit-server/common"
	"anydevelop.cn/recruit-server/models"
	"anydevelop.cn/recruit-server/redis"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	SECRET = "QQ794763733"
	EXPIRE = time.Hour * 24
)

type LoginAdmin struct {
	beego.Controller
}

func (c *LoginAdmin) URLMapping() {
	c.Mapping("LoginAdmin", c.LoginAdmin)
	c.Mapping("LogoutAdmin", c.LogoutAdmin)
}

// @router /LoginAdmin [post]
func (c *LoginAdmin) LoginAdmin() {
	var input models.Admin
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)
	v, err := models.GetAdminByName(input.Name)
	if err == nil {
		if bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(input.Password)) == nil {
			tokenString := getAdminToken(v)
			if tokenString != "" {
				addAdminToRedis(v)
				c.Data["json"] = common.Success(tokenString)
			}
		}
	} else {
		c.Data["json"] = common.Fail(err.Error())
	}
	c.ServeJSON()
}

// @router /LogoutAdmin [delete]
func (c *LoginAdmin) LogoutAdmin() {
	tokenString := c.Ctx.Request.Header.Get("Admin-Token")
	if tokenString != "" {
		token, err := jwt.Parse(tokenString, parse)
		if err == nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				redis.Rdb.Del("admin" + claims["name"].(string))
			}
		} else {
			beego.Error(err.Error())
		}
	}
	c.Data["json"] = common.Ok()
	c.ServeJSON()
}

func addAdminToRedis(v *models.Admin) {
	redis.Rdb.HSet("admin"+string(v.Id), "id", v.Id)
	redis.Rdb.HSet("admin"+string(v.Id), "name", v.Name)
	redis.Rdb.HSet("admin"+string(v.Id), "picture", v.Picture)
	redis.Rdb.Expire("admin"+string(v.Id), EXPIRE)
}

func getAdminToken(v *models.Admin) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   v.Id,
		"name": v.Name,
	})
	tokenString, err := token.SignedString([]byte(SECRET))
	if err == nil {
		return tokenString
	} else {
		beego.Error(err.Error())
	}
	return ""
}

func parse(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return []byte(SECRET), nil
}
