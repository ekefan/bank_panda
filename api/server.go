package api

import (
	db "github.com/ekefan/bank_panda/db/sqlc" //exposes all the files in db sqlc to this server package
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// server will serve HTTP request for the Banking service
type Server struct {
	store  db.Store   //allows me interact with the database
	router *gin.Engine //help send each api  request to the correct handler for processing
}

// NewServer create a new HTTP server and return a server instance
// and setup all api routes for that service on the server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	//add routes to router
	router.POST("/users", server.createUser)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.POST("/transfers", server.createTransfer)
	server.router = router
	return server
}

// Start: runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
