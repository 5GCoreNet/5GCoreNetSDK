package nlmf

import (
	"github.com/gin-gonic/gin"
	"log"
)

// Server represents a NLMF server.
type Server struct {
	address   string // IP:PORT
	apiRoot   string
	location  Location
	broadcast Broadcast
	logger    *log.Logger
	router    *gin.Engine
	stop      chan bool
}

// NewServer creates a new Server NLMF server instance.
// The address is the IP:PORT of the NLMF server.
func NewServer(address string, apiRoot string, logger *log.Logger) *Server {
	return &Server{
		address: address,
		apiRoot: apiRoot,
		logger:  logger,
		stop:    make(chan bool),
	}
}

// AttachLocation attaches a Location handler to the NLMF Server.
func (n *Server) AttachLocation(l Location) {
	n.location = l
}

// AttachBroadcast attaches a Broadcast handler to the NLMF Server.
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
	go n.router.Run(n.address)
	<-n.stop
	return
}

// Stop stops the NLMF Server.
func (n *Server) Stop() {
	n.stop <- true
}
