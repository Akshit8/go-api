/*
 * @File: server.api.go
 * @Description: file impls gin server and http-routing for our banking service
 * @LastModifiedTime: 2021-01-15 18:14:45
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package api impls REST api
package api

import (
	"fmt"

	db "github.com/Akshit8/go-api/db/sqlc"
	"github.com/Akshit8/go-api/token"
	"github.com/Akshit8/go-api/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking services.
type Server struct {
	config util.Config
	store  db.Store
	token  token.Maker
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannnot create token maker: %w", err)
	}
	server := &Server{
		config: config,
		store:  store,
		token:  tokenMaker,
		router: gin.Default(),
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.router.POST("/accounts", server.createAccount)
	server.router.GET("/accounts/:id", server.getAccount)
	server.router.GET("/accounts", server.listAccount)

	server.router.POST("/transfers", server.createTransfer)

	server.router.POST("/users", server.createUser)

	return server, nil
}

// Start runs the HTTP server on specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
