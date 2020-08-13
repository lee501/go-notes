package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

//gin-jwt结构体
type GinJWTMiddleware struct {
	SingingAlgorithm string //算法
	Key 			[]byte	//密钥
	Timeout 		time.Duration   //有效期
	MaxRefresh		time.Duration	//刷新token有效期
	Authenticator   func(userId string, password string, c *gin.Context)   //认证回调函数
	Authorizator    func(userID string, c *gin.Context) bool	//认证后处理
	Unauthorized 	func(*gin.Context, int, string)   //认证失败函数
}

/*
	LoginHandler和RefreshHandler可以直接拿来使用
	在LoginHandler中会回调Authenticator 验证用户名和密码
*/

//Authenticator


type AuthBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var authMiddleware *jwt.GinJWTMiddleware
func init() {
	// the jwt middleware
	authMiddleware = &jwt.GinJWTMiddleware{
		Realm:         "Realmname",
		Key:           []byte("Secretkey"),
		Timeout:       time.Hour * 12,
		MaxRefresh:    time.Hour * 24,
		Authenticator: jwtAuthFunc,
		Unauthorized:  jwtUnAuthFunc,
		// 其他默认
	}
}

func jwtAuthFunc(c *gin.Context) (interface{}, error) {
	var body AuthBody
	if err := c.ShouldBind(&body); err != nil {
		return nil, err
	}
	//u := FindAdmin(body.Username)
	//u.checkPassword(body.ppassword)
	return "用户", nil
}




//Unauthorized

func jwtUnAuthFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

//在路由中使用中间件
func main() {
	route := gin.Default()
	//登录使用LoginHandler
	route.POST("/login", authMiddleware.LoginHandler)

	auth := route.Group("/")
	auth.Use(authMiddleware.MiddlewareFunc())
	auth.POST("/refresh_token", authMiddleware.RefreshHandler)
}
