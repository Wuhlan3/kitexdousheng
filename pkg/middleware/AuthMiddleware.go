package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取token
		tokenString := ctx.Query("token")

		////validate token format
		//if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
		//	ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		//	ctx.Abort()
		//	return
		//}
		//
		//tokenString = tokenString[7:]

		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		////验证通过，获取claim中的userid
		//userId := claims.UserId
		//DB := common.GetDB()
		//var user model.User
		//DB.First(&user, userId)
		//
		////用户不存在
		//if user.ID == 0 {
		//	ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		//	ctx.Abort()
		//	return
		//}

		//用户存在，将user的信息写入context
		ctx.Set("uid", claims.UserId)
		ctx.Next()
	}
}
