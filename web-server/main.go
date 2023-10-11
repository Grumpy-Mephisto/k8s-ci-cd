package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

type APIHandler interface {
	Handle(c *fiber.Ctx) error
}

type HealthHandler struct{}

func (h *HealthHandler) Handle(c *fiber.Ctx) error {
	if err := c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "API is healthy",
	}); err != nil {
		return err
	}

	return nil
}

type ArgoHandler struct{}

func (h *ArgoHandler) Handle(c *fiber.Ctx) error {
	if err := c.Status(http.StatusOK).JSON(fiber.Map{
		"message":   "CI/CD pipeline is working!",
		"success":   true,
		"timestamp": time.Now().Unix(),
	}); err != nil {
		return err
	}

	return nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		ProxyHeader:             fiber.HeaderXForwardedFor,
	})

	healthHandler := &HealthHandler{}
	argoHandler := &ArgoHandler{}

	router.Get("/health", healthHandler.Handle)
	router.Get("/argo", argoHandler.Handle)

	log.Printf("Server is running on port %s", port)
	if err := router.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
