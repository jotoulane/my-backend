package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"my-backend/global/code"
	"my-backend/global/response"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const TokenExpireDuration = time.Hour * 24

var CustomSecret = []byte("想见你只想见你未来过去")

type CustomClaims struct {
	Id       uint   `json:"id"`
	UserName string `json:"userName"`
	Phone    string `json:"phone"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(id uint, userName string, phoneNumber string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	// 创建一个我们自己声明的数据
	claims := CustomClaims{
		id,
		userName,
		phoneNumber,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			Issuer:    "my-backend",      // 签发人
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(CustomSecret)

	return token, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	var mc = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.ResponseErrorWithMsg(code.StatusUnauthorized, "请求未携带token，无权限访问"))
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, response.ResponseErrorWithMsg(code.StatusUnauthorized, "token格式不正确"))

			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.ResponseErrorWithMsg(code.StatusUnauthorized, err.Error()))
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("userID", mc.Id)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
