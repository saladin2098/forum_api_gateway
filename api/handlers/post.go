package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/saladin2098/forum_api_gateway/genproto"
)

// @Router 				/post/create [POST]
// @Summary 			Creates Post
// @Description		 	This api create Post
// @Tags 				Post
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               data body pb.PostModel true "Post Data"
// @Success 201 		{object} pb.Post
// @Failure 400 		string Error
func (h *Handler) CreatePost(c *gin.Context) {
	var model pb.PostModel
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := pb.Post{
		UserId: model.UserId,
		Title:  model.Title,
        Body:   model.Body,
		CategoryId: model.CategoryId,
	}
	res, err := h.Clients.PostC.CreatePost(c, &post)
	if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusCreated, res)
}
// @Router 				/post/{id} [GET]
// @Summary 			Gets Post
// @Description		 	This api Gets Post by id
// @Tags 				Post
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Post Id"
// @Success 200 		{object} pb.Post
// @Failure 400 		string Error
func (h *Handler) GetPost(c *gin.Context) {
	id := c.Param("id")
    post := &pb.ById{
        Id: id,
    }
    res, err := h.Clients.PostC.GetPost(c, post)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}
// @Router 				/post/all [GET]
// @Summary 			Gets all Post 
// @Description		 	This api gets all Post
// @Tags 				Post
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param              user_id query string false "User Id"
// @Param              title query string false "Post Title"
// @Param              category_id query string false "Category Id"
// @Param              body query string false "Post Body"
// @Success 200 		{object} pb.Posts
// @Failure 400 		string Error
func (h *Handler) GetPosts(c *gin.Context) {
	user_id := c.Query("user_id")
	title := c.Query("title")
	category_id := c.Query("category_id")
	body := c.Query("body")
	filter := pb.PostFilter{
		UserId: user_id,
        Title:  title,
        CategoryId: category_id,
        Body: body,
	}
	res, err := h.Clients.PostC.GetPosts(c, &filter)
	if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK, res)
}
// @Router 				/admin/post/{id} [PUT]
// @Summary 			Updates Post 
// @Description		 	This api Updates Post
// @Tags 				Post
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Post Id"
// @Param               data body pb.PostModel true "Post Data"
// @Success 200 		{object} pb.Post
// @Failure 400 		string Error
func (h *Handler) UpdatePost(c *gin.Context) {
	id := c.Param("id")
    var model pb.PostModel
    if err := c.ShouldBindJSON(&model); err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    post := pb.Post{
        PostId: id,
        UserId: model.UserId,
        Title:  model.Title,
        Body:   model.Body,
        CategoryId: model.CategoryId,
    }
    res, err := h.Clients.PostC.UpdatePost(c, &post)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}
// @Router 				/admin/post/{id} [DELETE]
// @Summary 			Deletes Post 
// @Description		 	This api Deletes Post
// @Tags 				Post
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Post Id"
// @Success 200 		{object} string "post deleted successfully!"
// @Failure 400 		string Error
func (h *Handler) DeletePost(c *gin.Context) {
	id := c.Param("id")
    post := &pb.ById{
        Id: id,
    }
    _, err := h.Clients.PostC.DeletePost(c, post)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message":"post deleted successfully!"})
}
// @Router 				/post/search [GET]
// @Summary 			Search Post 
// @Description		 	This api Searchs Post
// @Tags 				Post
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               title query string false "Post Title"
// @Param               body query string false "Post Body"
// @Success 200 		{object} pb.Posts
// @Failure 400 		string Error
func (h *Handler) SearchPosts(c *gin.Context) {
	title := c.Query("title")
	body := c.Query("body")
    filter := pb.PostFilter{Title: title, Body: body}
	res, err := h.Clients.PostC.GetPosts(c, &filter)
	if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK, res)
}
// @Router 				/post/{id}/comments [GET]
// @Summary 			Search Comments by Post
// @Description		 	This api Search Comments by Post
// @Tags 				Post
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Post ID"
// @Success 200 		{object} pb.Posts
// @Failure 400 		string Error
func (h *Handler) GetPostComments(c *gin.Context) {
    id := c.Param("id")
    filter := pb.CommentFilter{
        PostId: id,
    }
    res,err := h.Clients.CommentC.GetComments(c,&filter)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK,res)
}
