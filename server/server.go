package server

import (
	"api-sample/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "3000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	route := routes.ConfigRoute(s.server)
	log.Print("Server is running at port:", s.port)
	log.Fatal(route.Run(":" + s.port))
}
