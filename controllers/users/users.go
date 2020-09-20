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
	id, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		e := errors.NewBadRequestError("bad parameter, expected a number")
		c.JSON(e.Code, e)
		return
	}
	u, rErr := services.Get(id)
	if rErr != nil {
		c.JSON(rErr.Code, rErr)
		return
	}
	c.JSON(http.StatusOK, u)
}

func Insert(c *gin.Context) {
	var u users.User
	if err := c.ShouldBindJSON(&u); err != nil {
		e := errors.NewBadRequestError("invalid json data")
		c.JSON(e.Code, e)
		return
	}
	result, err := services.Insert(&u)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
