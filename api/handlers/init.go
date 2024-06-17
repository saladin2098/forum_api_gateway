package handlers

import "github.com/saladin2098/forum_api_gateway/clients"

type Handler struct {
	Clients *clients.Clients
}

func NewHandler() (*Handler,error) {
	clients,err := clients.NewClients() 
	if err!=nil {
        return nil,err
    }
	return &Handler{
        Clients: clients,
    },nil
}