package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/saladin2098/forum_api_gateway/genproto"
)
// @Router 				/tag/create [POST]
// @Summary 			Creates Tag
// @Description		 	This api create Tag
// @Tags 				Tag
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               name query string true "Tag Name"
// @Success 201 		{object} pb.Tag
// @Failure 400 		string Error
func (h *Handler) CreateTag(c *gin.Context) {
	name := c.Query("name")
	tag := &pb.Tag{
		Name: name,
	}
	res, err := h.Clients.TagC.CreateTag(c, tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}
// @Router 				/tag/{id} [DELETE]
// @Summary 			Deletes Tag
// @Description		 	This api Deletes Tag
// @Tags 				Tag
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Tag ID"
// @Success 200 		{object} string "tag deleted successfully!"
// @Failure 400 		string Error
func (h *Handler) DeleteTag(c *gin.Context) {
	id := c.Param("id")
    tag := &pb.ById{
        Id: id,
    }
    _, err := h.Clients.TagC.DeleteTag(c, tag)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message":"tag deleted successfully!"})
}
// @Router 				/tag/{id} [PUT]
// @Summary 			Updates Tag
// @Description		 	This api Updates Tag
// @Tags 				Tag
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               name query string true "Tag Name"
// @Param               id path string true "Tag ID"
// @Success 200 		{object} pb.Tag
// @Failure 400 		string Error
func (h *Handler) UpdateTag(c *gin.Context) {
	id := c.Param("id")
    name := c.Query("name")
    tag := &pb.Tag{
        TagId: id,
        Name:  name,
    }
    res, err := h.Clients.TagC.UpdateTag(c, tag)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}
// @Router 				/tag/get/{name} [GET]
// @Summary 			Gets Tag by Name
// @Description		 	This api Gets Tag by Name
// @Tags 				Tag
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               name path string true "Tag Name"
// @Success 200 		{object} pb.Tag
// @Failure 400 		string Error
func (h *Handler) GetTag(c *gin.Context) {
	name := c.Param("name")
	req := pb.ByName{
		Name: name,
	}
	res,err := h.Clients.TagC.GetTag(c, &req)
	if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK,res)
}
// @Router 				/tag/all [GET]
// @Summary 			Gets All Tags 
// @Description		 	This api Gets All Tags 
// @Tags 				Tag
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Success 200 		{object} pb.TagList
// @Failure 400 		string Error
func (h *Handler) GetTags(c *gin.Context) {
	res, err := h.Clients.TagC.GetTags(c, &pb.Void{})
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}
// @Router 				/tag/popular [GET]
// @Summary 			Gets Popular Tags 
// @Description		 	This api Gets Popular Tags 
// @Tags 				Tag
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Success 200 		{object} pb.TagList
// @Failure 400 		string Error
func (h *Handler) GetPopularTags(c *gin.Context) {
	res, err := h.Clients.TagC.GetPopularTags(c, &pb.Void{})
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}	
// @Router 				/tag/{id}/posts [GET]
// @Summary 			Gets Post By Tag
// @Description		 	This api Gets Post By Tag
// @Tags 				Tag
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Tag ID"
// @Success 200 		{object} pb.Tag
// @Failure 400 		string Error
func (h *Handler) GetPostsBytag(c *gin.Context) {
    id := c.Param("id")
    tag := pb.TagFilter{
        Tag: id,
    }
    res, err := h.Clients.PostC.GetPostsByTag(c, &tag)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}