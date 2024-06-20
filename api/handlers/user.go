package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/saladin2098/forum_api_gateway/genproto"
	t "github.com/saladin2098/forum_api_gateway/token"
)

// @Router 				/user/register [POST]
// @Summary 			Creates User
// @Description		 	This api create User
// @Tags 				User
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               user body pb.UserModel true "User Data"
// @Success 201 		{object} pb.User
// @Failure 400 		string Error
func (h *Handler) RegisterUser(c *gin.Context) {
	var model pb.UserModel
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while binding": err.Error()})
		return
	}
	user := pb.User{
		UserName: model.UserName,
		Email:    model.Email,
		Password: model.Password,
	}
	res, err := h.Clients.UserC.RegisterUser(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error creating user": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// @Router 				/admin/user/all [GET]
// @Summary 			Gets all User
// @Description		 	This api Get all Users
// @Tags 				User
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Success 200 		{object} pb.Users
// @Failure 400 		string Error
func (h *Handler) GetUsers(c *gin.Context) {
	res, err := h.Clients.UserC.GetUsers(c, &pb.Void{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Router 				/admin/user/{id} [PUT]
// @Summary 			Updates User
// @Description		 	This api Updates User
// @Tags 				User
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "user id to be updated"
// @Param               user body pb.UserModel true "User Data"
// @Success 200 		{object} pb.User
// @Failure 400 		string Error
func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var model pb.UserModel
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := pb.User{
		UserId:   id,
		UserName: model.UserName,
		Email:    model.Email,
		Password: model.Password,
	}
	res, err := h.Clients.UserC.UpdateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Router 				/admin/user/{id} [DELETE]
// @Summary 			Deletes User
// @Description		 	This api Deletes User
// @Tags 				User
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "user id to be deleted"
// @Success 200 		{object} string "user deleted successfully"
// @Failure 400 		string Error
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user := &pb.ById{
		Id: id,
	}
	_, err := h.Clients.UserC.DeleteUser(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

// @Router 				/user/login [POST]
// @Summary 			Logins User
// @Description		 	This api Logins User
// @Tags 				User
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               data body pb.LoginReq true "login credentials"
// @Success 201 		{object} pb.Token
// @Failure 400 		string Error
func (h *Handler) LoginUser(c *gin.Context) {
	var log_req pb.LoginReq
	if err := c.ShouldBindJSON(&log_req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Clients.UserC.LoginUser(c, &log_req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Router 				/user/info [GET]
// @Summary 			Gets User info
// @Description		 	This api Gets User info
// @Tags 				User
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Success 200 		{object} pb.User
// @Failure 400 		string Error
func (h *Handler) GetUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}
	claims, err := t.ExtractClaim(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	username, ok := claims["username"].(string)
	if !ok || username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in token"})
		return
	}
	user := &pb.ByUsername{
		Username: username,
	}
	res, err := h.Clients.UserC.GetUserInfo(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}


// @Router 				/user/posts/{id} [GET]
// @Summary 			Gets Posts of User
// @Description		 	This api Gets Posts of User
// @Tags 				User
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "User Id"
// @Success 200 		{object} pb.Posts
// @Failure 400 		string Error
func (h *Handler) GetUserPosts(c *gin.Context) {
	id := c.Param("id")
	filter := pb.PostFilter{
		UserId: id,
	}
	res, err := h.Clients.PostC.GetPosts(c, &filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
