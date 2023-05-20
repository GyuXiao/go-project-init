package app

import (
	"GyuBlog/global"
	"GyuBlog/pkg/errcode"
	"GyuBlog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// 根据客户端传入的 AppKey 和 AppSecret 以及在项目配置中所设置的签发者（Issuer）和过期时间（ExpiresAt），
// 根据指定的算法生成签名后的 Token

func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMd5(appKey),
		AppSecret: util.EncodeMd5(appSecret),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expireTime},
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	// 根据 Claims 结构体创建 Token 实例
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// JWT 令牌的解析和校验

func ParseToken(token string) (*Claims, error) {
	// ParseWithClaims 用于解析鉴权的声明，方法内部主要是具体的解码和校验
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		// Valid 验证基于时间的声明，如果没有任何声明在令牌中，仍然会被认为是有效的
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 不知道为什么，如果把下面的 JWT 单独放在 middleware/jwt.go 文件中，会报循环引用的错

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if res, exist := c.GetQuery("token"); exist {
			token = res
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := ParseToken(token)
			if err != nil {
				switch err {
				case jwt.ErrTokenExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}
		if ecode != errcode.Success {
			response := NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}
