package api

import (
	db "github.com/SmoothWay/simpleBank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for our banking system
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	server.router = router
	return server
}

// Run the HTTP server on a specified address
func (s *Server) Start(address string) error {
	return s.router.Run()
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
