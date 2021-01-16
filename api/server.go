/*
 * @File: server.api.go
 * @Description: file impls gin server and http-routing for our banking service
 * @LastModifiedTime: 2021-01-15 18:14:45
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package api impls REST api
package api

import (
	db "github.com/Akshit8/go-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking services.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server{
	server := &Server{
		store: store,
		router: gin.Default(),
	}

	server.router.POST("/accounts", server.createAccount)
	server.router.GET("/accounts/:id", server.getAccount)
	server.router.GET("/accounts", server.listAccount)

	return server
}

// Start runs the HTTP server on specific address.
func (server *Server) Start(address  string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
