package api

import (
	"github.com/OwLu0905/simplebank_owlu/db/sqlc"
	"github.com/gin-gonic/gin"
)

// NOTE: Server serves HTTP requests for our banking service.
type Server struct {
	store  sqlc.Store
	router *gin.Engine
}

// NOTE: NewServer creates a new HTTP server and setup routing
func NewServer(store sqlc.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// TODO: Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
