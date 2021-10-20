package v1

import (
	"gin-template/core/util"

	"github.com/gin-gonic/gin"
)

// @Summary 测试用户登录
// @Produce  json
// @Param user_id query string false "user id"
// @Success 200 {string} json ""
// @Router /api/v1/login [post]
func UserLogin(c *gin.Context) {
	user_id := c.Query("user_id")
	token, err := util.CreateToken(user_id)
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