package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/saladin2098/forum_api_gateway/genproto"
)

// @Router 				/post-tag/create [POST]
// @Summary 			Creates PostTag
// @Description		 	This api Creates PostTag
// @Tags 				Post-Tag
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               data body pb.PostTagModel true "Post Tag Data"
// @Success 201 		{object} pb.PostTag
// @Failure 400 		string Error
func (h *Handler) CreatePostTag(c *gin.Context) {
	var model pb.PostTagModel
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postTag := pb.PostTag{
        PostId: model.PostId,
        TagId:  model.TagId,
    }
	res, err := h.Clients.PostTagC.CreatePostTag(c, &postTag)
	if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusCreated, res)
}
// @Router 				/post-tag/all [GET]
// @Summary 			Gets all PostTags
// @Description		 	This api Gets all PostTags
// @Tags 				Post-Tag
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               post_id query string false "Post ID"
// @Success 200 		{object} pb.PostTags
// @Failure 400 		string Error
func (h *Handler) GetPostTags(c *gin.Context) {
	post_id := c.Query("post_id")
	filter := pb.ByPost{
		PostId: post_id,
	}
	res, err := h.Clients.PostTagC.GetPostTags(c, &filter)
	if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK, res)
}
// @Router 				/post-tag/{id} [GET]
// @Summary 			Gets PostTag by id
// @Description		 	This api Gets PostTag by id
// @Tags 				Post-Tag
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string false "Post-Tag ID"
// @Success 200 		{object} pb.PostTag
// @Failure 400 		string Error
func (h *Handler) GetPostTag(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Clients.PostTagC.GetPostTag(c, &pb.ById{Id: id})
	if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK, res)
}
// @Router 				/post-tag/{id} [DELETE]
// @Summary 			Deletes PostTag
// @Description		 	This api Deletes PostTag
// @Tags 				Post-Tag
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Post-Tag ID"
// @Success 200 		{object} string "Post-Tag Deleted Successfully"
// @Failure 400 		string Error
func (h *Handler) DeletePostTagById(c *gin.Context) {
	id := c.Param("id")
	_, err := h.Clients.PostTagC.DeletePostTagById(c, &pb.ById{Id: id})
	if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK, gin.H{"message":"Post-Tag Deleted Successfully"})
}
// @Router 				/post-tag/{id} [PUT]
// @Summary 			Updates PostTag
// @Description		 	This api updates PostTag
// @Tags 				Post-Tag
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Post-Tag ID"
// @Param  data         body  pb.PostTagModel true "Post-Tag data"
// @Success 200 		{object} pb.PostTag
// @Failure 400 		string Error
func (h *Handler) UpdatePostTag(c *gin.Context) {
	id := c.Param("id")
    var model pb.PostTagModel
    if err := c.ShouldBindJSON(&model); err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    postTag := pb.PostTag{
		PostTagId: id,
        PostId: model.PostId,
        TagId:  model.TagId,
    }
    res, err := h.Clients.PostTagC.UpdatePostTag(c, &postTag)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}