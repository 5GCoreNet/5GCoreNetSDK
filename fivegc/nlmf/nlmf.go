package nlmf

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// Server represents a NLMF server.
type Server struct {
	ip        string
	port      string
	apiRoot   string
	location  Location
	broadcast Broadcast
	logger    *log.Logger
	router    *gin.Engine
	stop      chan bool
}

// NewServer creates a new Server NLMF server instance.
func NewServer(ip string, port string, apiRoot string, logger *log.Logger) *Server {
	return &Server{
		ip:      ip,
		port:    port,
		apiRoot: apiRoot,
		logger:  logger,
		stop:    make(chan bool),
	}
}

// AttachLocation attaches a Location client to the NLMF Server.
func (n *Server) AttachLocation(l Location) {
	n.location = l
}

// AttachBroadcast attaches a Broadcast client to the NLMF Server.
func (n *Server) AttachBroadcast(b Broadcast) {
	n.broadcast = b
}

// Start starts the NLMF Server.
func (n *Server) Start() {
	n.router = gin.Default()
	n.router.Use(gin.Recovery())
	n.router.Use(gin.LoggerWithWriter(n.logger.Writer()))
	root := n.router.Group(n.apiRoot)
	if n.location != nil {
		attachLocationHandler(root, n.location)
	}
	if n.broadcast != nil {
		attachBroadcastHandler(root, n.broadcast)
	}
	go n.router.Run(fmt.Sprintf("%s:%s", n.ip, n.port))
	<-n.stop
	return
}

// Stop stops the NLMF Server.
func (n *Server) Stop() {
	n.stop <- true
}
