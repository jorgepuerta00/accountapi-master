package server

import (
	"fmt"
	"net/http"

	"github.com/jorgepuerta00/accountapi-master/pkg/service"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func NewHTTPServer(logger logrus.FieldLogger, host, port string, AccountService service.Service) *HTTPServer {
	return &HTTPServer{
		logger:  logger,
		host:    host,
		port:    port,
		service: AccountService,
		router:  mux.NewRouter().StrictSlash(true),
	}
}

type HTTPServer struct {
	logger  logrus.FieldLogger
	service service.Service
	router  *mux.Router
	host    string
	port    string
}

func (s *HTTPServer) RegisterResource(url string, handler func(http.ResponseWriter, *http.Request)) {
	s.router.HandleFunc(url, handler)
}

func (s *HTTPServer) Run() {
	s.logger.Info("HTTPServer.Run:", "starting service")
	s.logger.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%s", s.host, s.port),
			s.router,
		),
	)
	s.logger.Info("HTTPServer.Run:", "service ended")
}
