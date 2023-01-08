package http

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Config struct {
	Server struct {
		Port    string `yaml:"port"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"server"`

	Context struct {
		Timeout int `yaml:"timeout"`
	} `yaml:"context"`
}

type Server struct {
	config            Config
	echo              *echo.Echo
	queryController   QueryController
	commandController CommandController
}

func NewServer(config Config, queryController QueryController, commandController CommandController) *Server {
	s := &Server{
		config:            config,
		echo:              echo.New(),
		queryController:   queryController,
		commandController: commandController,
	}

	s.useMiddleware()
	s.useRoutes()

	return s
}

func (s *Server) Start() error {
	port := s.config.Server.Port

	return s.echo.Start(port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *Server) Fatal(err error) {
	s.echo.Logger.Fatal(err)
}

func (s *Server) Config() Config {
	return s.config
}
