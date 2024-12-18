package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"rummy-logic-v3/global"
	"rummy-logic-v3/pkg/response"
	"rummy-logic-v3/pkg/xerr"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Id uint `json:"id"`
	jwt.MapClaims
}

const TokenExpireDuration = time.Hour * 2

//const TokenExpireDuration = time.Minute / 3

var JWTSecret = []byte("82D9ED82-B0F0-389B-89EE-AF893E324780")

func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("无效的Token")
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			global.Logger.WithCaller(1).Error(context.Background(), "用户没有传递token", nil)
			response.Fail(c, xerr.UnauthorizedTokenError)
			c.Abort()
			return
		}
		// 按空格分割
		//parts := strings.SplitN(authHeader, " ", 2)
		//if !(len(parts) == 2 && parts[0] == "Bearer") {
		//	log.Println("传递的不是有效的token", parts)
		//	response.FailHttpStatus(c, xerr.TokenInvalid, "token格式错误", http.StatusOK)
		//	c.Abort()
		//	return
		//}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(authHeader)
		if err != nil {
			global.Logger.WithCaller(1).Error(context.Background(), "用户传递的token解析错误: "+err.Error(), nil)
			response.Fail(c, xerr.UnauthorizedTokenError)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("userId", mc.Id)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func GenderToken(id uint, exp time.Time) (string, error) {
	c := MyClaims{
		Id: id,
		MapClaims: jwt.MapClaims{
			"id":  id,
			"exp": jwt.NewNumericDate(exp),
			"iat": jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(JWTSecret)
}
