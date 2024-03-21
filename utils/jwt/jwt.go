package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// TokenExpireDuration token 过期时间
const TokenExpireDuration = time.Hour * 24 * 30

// CustomSecret 用于加盐的字符串
var CustomSecret = []byte("你出现在我诗的每一页")

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               int64  `json:"user_id"`
	PhoneNumber          string `json:"phone_number"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(userId int64, phoneNumber string) (string, error) {
	// 创建一个我们自己声明的数据
	claims := CustomClaims{
		userId,
		phoneNumber,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "histbook", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
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
