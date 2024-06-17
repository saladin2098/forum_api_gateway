package clients

import (
	"github.com/saladin2098/forum_api_gateway/config"
	pb "github.com/saladin2098/forum_api_gateway/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type Clients struct {
	CategoryC pb.CategoryServiceClient
	PostC pb.PostServiceClient
	CommentC pb.CommentServiceClient
	TagC pb.TagServiceClient
	PostTagC pb.PostTagServiceClient
	UserC pb.UserServiceClient
}
func NewClients() (*Clients, error) {
	cfg := config.Load()
	auth_apth := "localhost" + cfg.HTTPPortAuth
	forum_path := "localhost" + cfg.HTTPPortForum
	conn_auth,err := grpc.NewClient(auth_apth, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	conn_forum,err := grpc.NewClient(forum_path,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!= nil {
        return nil, err
    }

	catC := pb.NewCategoryServiceClient(conn_forum)
	postC := pb.NewPostServiceClient(conn_forum)
	commentC := pb.NewCommentServiceClient(conn_forum)
	tagC := pb.NewTagServiceClient(conn_forum)
	postTagC := pb.NewPostTagServiceClient(conn_forum)

	userC := pb.NewUserServiceClient(conn_auth)
	return &Clients{
		CategoryC: catC,
        PostC: postC,
        CommentC: commentC,
        TagC: tagC,
        PostTagC: postTagC,
        UserC: userC,
	},nil
}