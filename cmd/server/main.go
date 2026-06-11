package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/sxmmie/mcp-server-shop/configs"
	"github.com/sxmmie/mcp-server-shop/internal/client"
	"github.com/sxmmie/mcp-server-shop/internal/mcp"
	"github.com/sxmmie/mcp-server-shop/internal/tools/cart"
	"github.com/sxmmie/mcp-server-shop/internal/tools/products"
)

func main() {
	logger := logrus.New()
	logger.SetOutput(os.Stderr) // Use Stderr for logs (stdout is for JSON-RPC)
	logger.SetFormatter(&logrus.JSONFormatter{})

	config, err := configs.Load()
	if err != nil {
		logger.WithError(err).Fatal("Failed to load configuration")
	}

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logger.WithError(err).Warn("Invalid log level, using info")
		level = logrus.InfoLevel
	}

	logger.SetLevel(level)

	// log to the console
	logger.WithFields(logrus.Fields{
		"api_url":               config.APIURL,
		"log_level":             config.LogLevel,
		"auth_token_configured": config.AuthToken != "",
	}).Info("Starting MCP E-Commerce Server")

	// create HTTP client
	restClient := client.NewRestyClient(config.APIURL, config.AuthToken, logger)

	toolRegistry := mcp.NewRegistry(logger)

	products.NewProductToolset(nil, logger, restClient)
	cart.NewCartToolset(toolRegistry, logger, restClient)

	logger.WithField("tool_count", len(toolRegistry.ListTools())).Info("Registered tools")

	mcpServer := mcp.NewServer(toolRegistry, logger)

	logger.Info("Starting server on stdio transport")

	if err := mcpServer.Start(); err != nil {
		logger.WithError(err).Fatal("Serevr error")
	}
}
