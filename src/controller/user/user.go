package user

import (
	// "middlewares"

	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"gitlab.com/abhishek.k8/crud/src/model"
	res "gitlab.com/abhishek.k8/crud/src/response"
)

//GetRoutes -
func (uc *UserController) GetRoutes(appGroup *gin.RouterGroup) {

	appGroup.POST("/register", uc.createUser)
	appGroup.PUT("/update", uc.updateUser)
	appGroup.DELETE("/delete", uc.deleteUser)
	appGroup.GET("/getUser", uc.getUser)
	appGroup.GET("/getAllUser", uc.getAllUser)

}

func (uc *UserController) createUser(c *gin.Context) {
	var user model.Users
	c.Bind(&user)
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(http.StatusPreconditionFailed, res.PreConditionFailed(c, err))
		return
	}
	err = user.Register()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, res.BadDataRequest(c, err))
		return
	}
	c.JSON(http.StatusCreated, res.SuccessResponse(c, user))
	return

}

func (uc *UserController) updateUser(c *gin.Context) {
	var user model.Update
	c.Bind(&user)
	var u model.Users
	val, err := u.Updateuser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.InternalServerError(c, res.CustomError("Internal server error")))
		return
	}
	c.JSON(http.StatusCreated, res.SuccessResponse(c, val))
	return

}

func (uc *UserController) deleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	if userId == 0 {
		c.JSON(http.StatusBadRequest, res.CustomError("user id can not be zero."))
	}
	var user model.Users
	err := user.DeleteUser(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.InternalServerError(c, res.CustomError("Internal server error")))
		return
	}
	c.JSON(http.StatusCreated, res.SuccessResponse(c, user))
	return

}

func (uc *UserController) getUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	if userId == 0 {
		c.JSON(http.StatusBadRequest, res.CustomError("user id can not be zero."))
	}
	var user model.Users
	err := user.GetUser(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.InternalServerError(c, res.CustomError("Internal server error")))
		return
	}
	c.JSON(http.StatusCreated, res.SuccessResponse(c, user))
	return

}

func (uc *UserController) getAllUser(c *gin.Context) {
	var user model.Users
	err := user.GetAllUsers()
	if err == nil {
		c.JSON(http.StatusInternalServerError, res.InternalServerError(c, res.CustomError("Internal server error")))
		return
	}
	c.JSON(http.StatusCreated, res.SuccessResponse(c, err))
	return

}
