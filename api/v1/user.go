package v1

import (
	jwt_util "gin-template/core/util/jwt"

	"github.com/gin-gonic/gin"
)

// UserLogin
// @Summary 测试用户登录
// @Tags login
// @Produce  json
// @Param user_id query string false "user id"
// @Success 200 {string} json ""
// @Router /api/v1/login [post]
func UserLogin(c *gin.Context) {
	userId := c.Query("user_id")
	token, err := jwt_util.CreateToken(userId)
	if err != nil {
		println(err.Error())
		c.JSON(500, gin.H{
			"token": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}
