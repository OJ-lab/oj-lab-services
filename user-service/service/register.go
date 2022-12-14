package service

import (
	"github.com/OJ-lab/oj-lab-services/model"
	"github.com/OJ-lab/oj-lab-services/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	roleString := c.PostForm("role")
	role := model.String2Role(roleString)
	err := model.CreateUser(account, password, role)
	if err != nil {
		utils.ApplyError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}
