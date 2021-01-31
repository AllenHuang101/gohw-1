package main

import (
	"encoding/json"
	"gohw-1/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/role", Get)

	router.GET("/role/:id", GetOne)

	router.POST("/role", Post)

	router.PUT("/role/:id", Put)

	router.DELETE("/role/:id", Delete)

	router.Run(":8080")
}

// 取得全部資料
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, model.Roles)
	return
}

// 取得單一筆資料
func GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "param id error!")
		return
	}

	for _, role := range model.Roles {
		if role.ID == uint(id) {
			c.JSON(http.StatusOK, role)
			return
		}
	}
	c.JSON(http.StatusNotFound, "No Find Data!")
	return
}

// 新增資料
func Post(c *gin.Context) {
	var role model.Role
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&role)

	if err != nil {
		c.JSON(http.StatusBadRequest, "param body error!")
		return
	}

	role.ID = uint(len(model.Roles)) + 1
	model.Roles = append(model.Roles, role)
	c.JSON(http.StatusCreated, role)
	return
}

func Put(c *gin.Context) {
	var newRole model.Role
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "param id error!")
		return
	}

	decoder := json.NewDecoder(c.Request.Body)
	err = decoder.Decode(&newRole)

	if err != nil {
		c.JSON(http.StatusBadRequest, "param body error!")
		return
	}

	for i := 0; i < len(model.Roles); i++ {
		if model.Roles[i].ID == uint(id) {
			model.Roles[i].Name = newRole.Name
			model.Roles[i].Summary = newRole.Summary
			c.JSON(http.StatusOK, model.Roles[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, "No Find Modify Data!")
	return
}

// 刪除資料
func Delete(c *gin.Context) {
	strNewid := c.Param("id")

	newid, err := strconv.Atoi(strNewid)

	if err != nil {
		c.JSON(http.StatusBadRequest, "param id error!")
		return
	}

	for i := 0; i < len(model.Roles); i++ {
		if model.Roles[i].ID == uint(newid) {
			// if i == len(model.Roles)-1 {
			// 	model.Roles = model.Roles[0:i]
			// } else {
			// 	model.Roles = append(model.Roles[0:i], model.Roles[i+1:]...)
			// }
			model.Roles = append(model.Roles[0:i], model.Roles[i+1:]...)
			c.JSON(http.StatusOK, "Delete ok")
			return
		}
	}
	c.JSON(http.StatusNotFound, "No Find Delete Data!")
	return

}

type RoleVM struct {
	ID      uint   `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}
