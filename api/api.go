package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	config "github.com/siddhantprateek/Trakd/config"
	rbClient "github.com/siddhantprateek/Trakd/internals/broker"
	"github.com/siddhantprateek/Trakd/internals/consignment"
	"github.com/siddhantprateek/Trakd/internals/delivery"
	"github.com/siddhantprateek/Trakd/internals/track"
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

	brokerURL := config.GetEnv("BROKER_URL")
	if brokerURL == "" {
		brokerURL = os.Getenv("BROKER_URL")
	}

	rb, err := rbClient.NewRabbitMQClient(brokerURL)
	if err != nil {
		log.Panicln(err)
	}
	defer rb.Close()

	// go routines
	go func() {
		ctx := context.Background()
		time.Sleep(3 * time.Second)
		rb.Publish(ctx, consignment.Package{From: "Bordeaux", To: "Toulouse", VehicleID: "123"})
		time.Sleep(3 * time.Second)
		rb.Publish(ctx, consignment.Package{From: "Toulouse", To: "Monaco", VehicleID: "12"})
		time.Sleep(3 * time.Second)
		rb.Publish(ctx, consignment.Package{From: "Monaco", To: "Lyon", VehicleID: "123"})
		time.Sleep(3 * time.Second)
		rb.Publish(ctx, consignment.Package{From: "Lyon", To: "Paris", VehicleID: "123"})
		time.Sleep(3 * time.Second)
		rb.Publish(ctx, consignment.Package{From: "Paris", To: "Brussels", VehicleID: "123"})
		time.Sleep(3 * time.Second)
		rb.Publish(ctx, consignment.Package{From: "Brussels", To: "Rotterdam", VehicleID: "123"})
		time.Sleep(3 * time.Second)
		rb.Publish(ctx, consignment.Package{From: "Rotterdam", To: "Amsterdam", VehicleID: "123"})
	}()

	pub := track.NewConsignmentTrack(rb)
	delivery.NewConsignmentHandler(svr.fiber, pub)

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
