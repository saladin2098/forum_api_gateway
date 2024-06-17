package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/saladin2098/forum_api_gateway/genproto"
)
// @Router 				/comment/create [POST]
// @Summary 			Creates Comment
// @Description		 	This api create Commnet
// @Tags 				Comment
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               data body pb.CommentModel true "Comment Data"
// @Success 201 		{object} pb.Comment
// @Failure 400 		string Error
func (h *Handler) CreateComment(c *gin.Context) {
	var model pb.CommentModel
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment := pb.Comment{
		PostId: model.PostId,
		UserId: model.UserId,
		Body:   model.Body,
	}
	res, err := h.Clients.CommentC.CreateComment(c, &comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}
// @Router 				/comment/{id} [GET]
// @Summary 			Gets Comment
// @Description		 	This api Gets Commnet
// @Tags 				Comment
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Comment ID"
// @Success 200 		{object} pb.Comment
// @Failure 400 		string Error
func (h *Handler) GetComment(c *gin.Context) {
	id := c.Param("id")
    res, err := h.Clients.CommentC.GetComment(c, &pb.ById{Id: id})
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}
// @Router 				/comment/all [GET]
// @Summary 			Gets All Comment
// @Description		 	This api Gets All Commnet
// @Tags 				Comment
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               post_id query string false "Post ID"
// @Param               user_id query string false "User ID"
// @Param               body    query string false "Post Body"
// @Success 200 		{object} pb.Comment
// @Failure 400 		string Error
func (h *Handler) GetComments(c *gin.Context) {
	post_id := c.Query("post_id")
    user_id := c.Query("user_id")
    body := c.Query("body")
    filter := pb.CommentFilter{
        PostId: post_id,
        UserId: user_id,
        Body:   body,
    }
    res, err := h.Clients.CommentC.GetComments(c, &filter)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}
// @Router 				/comment/{id} [PUT]
// @Summary 			Updates Comment
// @Description		 	This api Updates Commnet
// @Tags 				Comment
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Comment ID"
// @Param        data   body pb.CommentModel true "Comment Data"
// @Success 200 		{object} pb.Comment
// @Failure 400 		string Error
func (h *Handler) UpdateComment(c *gin.Context) {
	id := c.Param("id")
    var model pb.CommentModel
    if err := c.ShouldBindJSON(&model); err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    comment := pb.Comment{
        CommentId: id,
        PostId:    model.PostId,
        UserId:    model.UserId,
        Body:      model.Body,
    }
    res, err := h.Clients.CommentC.UpdateComment(c, &comment)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}
// @Router 				/comment/{id} [DELETE]
// @Summary 			Deletes Comment
// @Description		 	This api Deletes Commnet
// @Tags 				Comment
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Comment ID"
// @Success 200 		{object} pb.Comment
// @Failure 400 		string Error
func (h *Handler) DeleteComment(c *gin.Context) {
	id := c.Param("id")
    comment := &pb.ById{
        Id: id,
    }
    _, err := h.Clients.CommentC.DeleteComment(c, comment)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message":"comment deleted successfully!"})
}
