package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	data "github.com/microsoft/commercial-marketplace-offer-deploy/internal/data"
	"github.com/microsoft/commercial-marketplace-offer-deploy/pkg/generated"
	"gorm.io/gorm"
)

// HTTP handler for creating deployments
func CreateEventSubscription(c echo.Context, db *gorm.DB) error {
	eventType := c.Param("eventType")

	var command *generated.CreateEventSubscription
	err := c.Bind(&command)

	if err != nil {
		return err
	}

	model := data.FromCreateEventSubscription(eventType, command)
	tx := db.Create(&model)

	log.Printf("Event Subscription [%s] created for event type %s.", model.Name, model.EventType)

	if tx.Error != nil {
		return err
	}

	id := model.ID.String()
	result := generated.EventSubscription{
		ID:        &id,
		Name:      &model.Name,
		Callback:  &model.Callback,
		EventType: &model.EventType,
	}

	return c.JSON(http.StatusOK, result)
}