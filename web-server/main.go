package main

import (
	"log"
	"os"

	"os-container-project/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		ProxyHeader:             fiber.HeaderXForwardedFor,
	})

	healthHandler := &handler.HealthHandler{}
	argoHandler := &handler.ArgoHandler{}
	memberHandler := &handler.MemberHandler{}
	memberIdHandler := &handler.MemberIdHandler{}

	router.Get("/health", healthHandler.Handle)
	router.Get("/argo", argoHandler.Handle)
	router.Get("/member", memberHandler.Handle)
	router.Get("/member/:id", memberIdHandler.Handle)

	if err := router.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
