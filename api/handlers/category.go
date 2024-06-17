package handlers

import (

	"net/http"
	"github.com/gin-gonic/gin"
	pb "github.com/saladin2098/forum_api_gateway/genproto"
)

// @Router 				/category/create [POST]
// @Summary 			Creates Category
// @Description		 	This api create Category
// @Tags 				Category
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               name query string true "Category Name"
// @Success 201 		{object} pb.Category
// @Failure 400 		string Error
func (h *Handler) CreateCategory(c *gin.Context) {
	name := c.Query("name")
	category := &pb.Category{
		Name: name,
	}
	res, err := h.Clients.CategoryC.CreateCategory(c, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
// @Router 				/category/{name} [GET]
// @Summary 			Gets Category
// @Description		 	This api Gets Category by Name
// @Tags 				Category
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               name path string true "Category Name"
// @Success 200 		{object} pb.Category
// @Failure 400 		string Error
func (h *Handler) GetCategory(c *gin.Context) {
	name := c.Param("name")
    cat_name := &pb.ByName{
        Name: name,
    }
    res, err := h.Clients.CategoryC.GetCategory(c, cat_name)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}
// @Router 				/category/all [GET]
// @Summary 			Gets All Categories
// @Description		 	This api Gets all Categories
// @Tags 				Category
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Success 200 		{object} pb.Categories
// @Failure 400 		string Error
func (h *Handler) ListCategories(c *gin.Context) {
	res, err := h.Clients.CategoryC.ListCategories(c, &pb.Void{})
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}
// @Router 				/category/{name} [DELETE]
// @Summary 			Deletes Category
// @Description		 	This api Deletes Category by Name
// @Tags 				Category
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               name path string true "Category Name"
// @Success 200 	    {object} string "category deleted successfully!"
// @Failure 400 		string Error
func (h *Handler) DeleteCategory(c *gin.Context) {
	name := c.Param("name")
    cat_name := &pb.ByName{
        Name: name,
    }
    _, err := h.Clients.CategoryC.DeleteCategory(c, cat_name)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message":"category deleted successfully!"})
}

// @Router 				/category/{id} [PUT]
// @Summary 			Updates Category
// @Description		 	This api Updates Category by Name
// @Tags 				Category
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Category ID"
// @Param               name query string true "Category Name"
// @Success 200 		{object} pb.Category
// @Failure 400 		string Error
func (h *Handler) UpdateCategory(c *gin.Context) {
	name := c.Query("name")
	id := c.Param("id")
    category := pb.Category{
		CategoryId: id,
        Name: name,
	}
    res, err := h.Clients.CategoryC.UpdateCategory(c, &category)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}
// @Router 				/category/posts/{id} [GET]
// @Summary 			Gets Posts by Category
// @Description		 	This api Gets Posts by Category
// @Tags 				Category
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param               id path string true "Category ID"
// @Success 200 		{object} pb.Posts
// @Failure 400 		string Error
func (h *Handler) GetPostsByCategory(c *gin.Context) {
    id := c.Param("id")
    filter := pb.PostFilter{
        CategoryId: id,
    }
    res,err := h.Clients.PostC.GetPosts(c,&filter)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}