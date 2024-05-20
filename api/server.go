package api

import (
	"fmt"

	"github.com/ekefan/bank_panda/token"
	"github.com/ekefan/bank_panda/utils"

	db "github.com/ekefan/bank_panda/db/sqlc" //exposes all the files in db sqlc to this server package
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// server will serve HTTP request for the Banking service
type Server struct {
	store  db.Store   //allows me interact with the database
	router *gin.Engine //help send each api  request to the correct handler for processing
	tokenMaker token.Maker
	config utils.Config
}

// NewServer create a new HTTP server and return a server instance
// and setup all api routes for that service on the server
func NewServer(store db.Store, config utils.Config) (*Server, error) {
	makeToken, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}
	server := &Server{
		store: store,
		tokenMaker: makeToken,
		config: config,
	}
	
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setUpRouter()
	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()
	//add routes to router
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)
	
	authRoutes.POST("/transfers", server.createTransfer)
	server.router = router
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
