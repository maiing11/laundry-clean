package api

import (
	"fmt"
	"log"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/api/controllers"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/api/routes"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/config"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/manager"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/usecase"
)

type Server struct {
	ucManager manager.UsecaseManager
	domain    usecase.Domains
	handler   config.RequestHandler
	host      string
}
type Runner interface{}

func (s *Server) setupRoute() {
	customer := controllers.NewCustomerController(s.domain)
	service := controllers.NewServiceController(s.domain)
	bill := controllers.NewBillController(s.domain)

	routes.NewRoutes(
		routes.NewCustomerRoutes(s.handler, customer),
		routes.NewServiceRoutes(s.handler, service),
		routes.NewBillRoutes(s.handler, bill),
	)
}

func (s *Server) Run() {
	s.setupRoute()
	if err := s.handler.Gin.Run(s.host); err != nil {
		log.Fatalf("error ni bos :%v", err.Error())
	}
}

func NewServer() *Server {
	// Infra manager
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)
	host := fmt.Sprintf(":%s", cfg.APIPort)

	return &Server{
		ucManager: usecaseManager,
		host:      host,
	}
}
