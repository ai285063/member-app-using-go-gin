package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int64
	Account  string
	Email    string
	Password string
}

func GetUserList(c *gin.Context) {
	var users []User
	if err := MysqlDB.Table("users").Find(&users).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "success",
			"Users":   users,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
}

func Register(c *gin.Context) {
	var user User
	user.Account = c.Request.FormValue("account")
	user.Password = c.Request.FormValue("password")
	user.Email = c.Request.FormValue("email")

	if err := MysqlDB.Table("users").Create(&user).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "Successfully registered.",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
}

func PutUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := findUser(int(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "User doesn't exist",
		})
		return
	}

	account := c.Request.FormValue("account")
	password := c.Request.FormValue("password")
	email := c.Request.FormValue("email")

	if err := MysqlDB.Table("users").
		Where("ID = ?", id).
		Updates(User{
			Account:  account,
			Email:    email,
			Password: password,
		}).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "User data updated.",
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := findUser(int(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "User doesn't exist",
		})
		return
	}

	if err := MysqlDB.Table("users").Delete(User{}, id).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "User deleted.",
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func GetViewCount(c *gin.Context) {
	val, err := RedisDB.Get(Ctx, ViewCount).Result()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"View count": val,
	})
}

func AddViewCount() {
	RedisDB.Incr(Ctx, ViewCount)
}

func findUser(id int) error {
	var user User
	err := MysqlDB.Table("users").Where("ID = ?", id).First(&user).Error
	return err
}
