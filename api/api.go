package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	config "github.com/siddhantprateek/Trakd/config"
)

type API interface {
	Start()
	Stop()
}

type server struct {
	fiber *fiber.App
	cfg   *config.APIConfiguration
}

func NewAPI(cfg *config.APIConfiguration) API {
	setDefaults(cfg)

	svr := server{
		fiber: fiber.New(
			fiber.Config{
				DisableStartupMessage:    true,
				DisableDefaultDate:       true,
				DisableHeaderNormalizing: true,
			},
		),
		cfg: cfg,
	}
	svr.fiber.Use(logger.New())
	svr.fiber.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":     "Trakd server.",
			"healthcheck": "ok",
		})
	})
	return &svr
}

func setDefaults(cfg *config.APIConfiguration) {
	if cfg.Port == 0 {
		cfg.Port = *intPointer(8090)
	}
}

func (s *server) Start() {
	log.Printf("server starting on port http://localhost:%d", s.cfg.Port)

	if err := s.fiber.Listen(fmt.Sprintf(":%d", s.cfg.Port)); err != nil {
		log.Fatalf("server cannot start on port %d", s.cfg.Port)
	}
}

func (s *server) Stop() {
	log.Println("server is closing")
	if err := s.fiber.Shutdown(); err != nil {
		log.Fatalf("server cannot be shutdown, err: %v", err)
	}
}

func intPointer(val int) *int {
	return &val
}
