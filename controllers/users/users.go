package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/KatherineEbel/bookstore_users-api/domain/users"
	"github.com/KatherineEbel/bookstore_users-api/services"
	"github.com/KatherineEbel/bookstore_users-api/utils/errors"
)

func Get(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	u, rErr := services.UsersService.Get(id)
	if rErr != nil {
		c.JSON(rErr.Code, rErr)
		return
	}
	c.JSON(http.StatusOK, u.Marshal(c.GetHeader("X-Public") == "true"))
}

func Insert(c *gin.Context) {
	var u users.User
	if err := c.ShouldBindJSON(&u); err != nil {
		e := errors.NewBadRequestError("invalid json data")
		c.JSON(e.Code, e)
		return
	}
	result, err := services.UsersService.Insert(&u)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func parseId(c *gin.Context) (int64, *errors.RestError) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("invalid request")
	}
	return userId, nil
}

func Update(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	var u users.User
	if err := c.ShouldBindJSON(&u); err != nil {
		rErr := errors.NewBadRequestError("invalid JSON")
		c.JSON(rErr.Code, rErr)
		return
	}
	u.Id = id
	isPartial := c.Request.Method == http.MethodPatch
	usr, updErr := services.UsersService.Update(isPartial, &u)
	if updErr != nil {
		c.JSON(updErr.Code, err)
		return
	}
	c.JSON(http.StatusOK, usr.Marshal(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	if err := services.UsersService.Delete(id); err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	byStatus, err := services.UsersService.FindByStatus(status)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, byStatus.Marshal(c.GetHeader("X-Public") == "true"))
}
