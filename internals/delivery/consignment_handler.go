package delivery

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	websocket "github.com/gorilla/websocket"
	consignment "github.com/siddhantprateek/Trakd/internals/consignment"
)

type PackageHandler struct {
	upgrader    websocket.Upgrader
	ConsignInit consignment.PackageInit
	w           http.ResponseWriter
	r           *http.Request
}

func NewConsignmentHandler(app *fiber.App, ci consignment.PackageInit) {
	handler := &PackageHandler{
		upgrader:    websocket.Upgrader{},
		ConsignInit: ci,
	}

	app.Get("/consignment/track/:vehicleId", handler.TrackByVehicleID)
}

func (p *PackageHandler) TrackByVehicleID(c *fiber.Ctx) error {
	wsConn, err := p.upgrader.Upgrade(p.w, p.r, nil)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		_, _, err := wsConn.ReadMessage()
		if err != nil {
			cancel()
		}
	}()

	for {
		select {
		case <-ctx.Done():
			wsConn.Close()
			return nil
		default:
			p, err := p.ConsignInit.TrackByVehicleID(ctx, c.Params("vehicleId"))
			if err != nil {
				log.Panicln(err)
				continue
			}

			err = wsConn.WriteJSON(p)
			if err != nil {
				log.Panicln(err)
			}
		}
	}
}
