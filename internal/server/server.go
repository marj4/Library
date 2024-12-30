package server

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func StartServer(Database *sql.DB) *Server {
	router := SetupRouter(Database)
	server := &Server{
		Router: router,
	}

	return server
}

func CreatePort(ServerPort string) string {
	return ":" + ServerPort
}
