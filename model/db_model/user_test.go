package db_model

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"my-backend/middleware"
	"testing"
	"time"
)

// 创建密码
func TestPasswordCreate(t *testing.T) {
	// 1、获取用户传来的pwd,使用字节切片转换 []byte("123456")

	// 2、调用 bcrypt.GenerateFromPassword 生成加密字符串
	hashPwd, _ := bcrypt.GenerateFromPassword([]byte("123456"), 10)

	// 3、此时 hashPwd 为字节切片，实际加密字符串需使用string转换
	//s := // 可以将此加密串保存到数据库，可作为密码匹配验证
	fmt.Printf("s:%v\n", string(hashPwd))
	//$2a$10$suBVtQytqdyRLSWgnPD2B.V4YDHoCam6EdMbu/2iyPgzUj74grmw.
	//$2a$10$8w0HZ6bNgUOSoPMWDenutOv6kKKQ.PZ.xGowZtik72mdxQvYxjwvS
}

// 校验密码
func TestCheck(t *testing.T) {
	// 1、将数据库中的加密串做字节切片转换
	// $2a$10$xKydBNAiBM7ImkE8QKm7lOLTpSk8ueUY6Da.OIAq8ee0Y0lZ7t1aO
	byteHashPwd := []byte("123456") // 实际的加密字符串

	// 2、调用 bcrypt.CompareHashAndPassword 证密码是否匹配
	// 第一个参数为通过字节切片转换的加密的哈希串、第二个参数为字节切片转换过的用户输入密码值
	err := bcrypt.CompareHashAndPassword(byteHashPwd, []byte("123456"))
	// 没有错误则密码匹配
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("密码匹配")
}

func TestGenToken(t *testing.T) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := middleware.CustomClaims{
		Id:       1,
		UserName: "admin",
		Phone:    "12345678910",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "my-backend", // 签发人
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Printf("tokenClaims:%v\n", tokenClaims)
	token, err := tokenClaims.SignedString(middleware.CustomSecret)
	fmt.Printf("token%v, err:%v\n", token, err)
}
