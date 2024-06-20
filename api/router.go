package api

import (
	"github.com/gin-gonic/gin"
	"github.com/saladin2098/forum_api_gateway/api/handlers"
	_ "github.com/saladin2098/forum_api_gateway/docs"
	"github.com/saladin2098/forum_api_gateway/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Online Forum API Documentation
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// TODO kamroq middleware ishlating. Iloji boricha umumiy middleware ishlatishga harakat qiling
func NewGin(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	category := r.Group("/category")
	category.Use(middleware.Middleware())
	{
		category.POST("/create",h.CreateCategory)
		category.GET("/all",h.ListCategories)
		category.GET("/:name",h.GetCategory)
		category.DELETE("/:name",h.DeleteCategory)
		category.PUT("/:id",h.UpdateCategory)
		category.GET("/posts/:id",h.GetPostsByCategory)
	}

	userAdmin := r.Group("/admin/user")
	userAdmin.Use(middleware.MiddlewareAdmin())
	{
		userAdmin.GET("/all",h.GetUsers)
        userAdmin.DELETE("/:id",h.DeleteUser)
        userAdmin.PUT("/:id",h.UpdateUser)
	}

	user:= r.Group("/user")
	user.Use(middleware.Middleware())
	{
		user.POST("/register",h.RegisterUser)
        user.GET("/info",h.GetUserInfo)
        user.POST("/login",h.LoginUser)
		user.GET("/posts/:id",h.GetUserPosts)
	}

	postAdmin := r.Group("/admin/post")
	postAdmin.Use(middleware.MiddlewareAdmin())
	{
        postAdmin.DELETE("/:id",h.DeletePost)
        postAdmin.PUT("/:id",h.UpdatePost)
	}
	post := r.Group("/post")
	post.Use(middleware.Middleware())
	{
		post.POST("/create",h.CreatePost)
        post.GET("/all",h.GetPosts)
        post.GET("/:id",h.GetPost)
        post.GET("/search",h.SearchPosts)
		post.GET("/:id/comments",h.GetPostComments)
	}

	comment := r.Group("/comment")
	comment.Use(middleware.Middleware())
	{
		comment.POST("/create",h.CreateComment)
        comment.GET("/all",h.GetComments)
        comment.GET("/:id",h.GetComment)
        comment.DELETE("/:id",h.DeleteComment)
        comment.PUT("/:id",h.UpdateComment)
	}
	tag := r.Group("/tag")
	tag.Use(middleware.Middleware())
	{
		tag.POST("/create",h.CreateTag)
        tag.GET("/all",h.GetTags)
        tag.GET("/popular",h.GetPopularTags)
        tag.GET("/get/:name",h.GetTag)
        tag.DELETE("/:id",h.DeleteTag)
        tag.PUT("/:id",h.UpdateTag)
		tag.GET("/:id/posts",h.GetPostsBytag)
	}
	post_tag := r.Group("/post-tag")
	post_tag.Use(middleware.Middleware())
	{
		post_tag.POST("/create",h.CreatePostTag)
        post_tag.GET("/all",h.GetPostTags)
        post_tag.GET("/:id",h.GetPostTag)
        post_tag.DELETE("/:id",h.DeletePostTagById)
        post_tag.PUT("/:id",h.UpdatePostTag)
	}
	return r
}