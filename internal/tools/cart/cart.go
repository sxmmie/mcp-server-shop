package cart

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/sxmmie/mcp-server-shop/internal/client"
	"github.com/sxmmie/mcp-server-shop/internal/mcp"
)

type CartToolset struct {
	reg        *mcp.Registry
	logger     *logrus.Logger
	restClient *client.RestClient
}

func NewCartToolset(reg *mcp.Registry, logger *logrus.Logger, restClient *client.RestClient) *CartToolset {
	ct := &CartToolset{
		reg:        reg,
		logger:     logger,
		restClient: restClient,
	}

	ct.registerCartTools()

	return ct
}

func (c *CartToolset) registerCartTools() {
	c.reg.Register(mcp.Tool{
		Name:        "add_to_cart",
		Description: "Add a product to the shopping cart (requires authentication)",
		InputSchema: mcp.InputSchema{
			Type: "object",
			Properties: map[string]mcp.Property{
				"product_id": {
					Type:        "number",
					Description: "ID of the product to add",
				},
				"quantity": {
					Type:        "number",
					Description: "Quantity to add (default: 1)",
				},
			},
			Required: []string{"product_id"},
		},
	}, c.handleAddToCart)

	c.reg.Register(mcp.Tool{
		Name:        "view_cart",
		Description: "ACiew current shopping cart contents (requires authentication)",
		InputSchema: mcp.InputSchema{
			Type:       "objetc",
			Properties: map[string]mcp.Property{},
			Required:   []string{},
		},
	}, c.handleVIewCart)
}

func (c *CartToolset) handleAddToCart(ctx context.Context, args map[string]any) (mcp.CallToolResult, error) {
	c.logger.WithField("args", args).Info("Adding to cart")

	productIDFloat, ok := args["product_id"].(float64)
	if !ok {
		return mcp.CallToolResult{}, fmt.Errorf("invalid prodict_id; %+v", args["product_id"])
	}

	productID := uint(productIDFloat)
	var quantity uint = 1
	if q, ok := args["quantity"].(uint); ok {
		if q == 0 {
			q = 1
		}
		quantity = q
	}

	body := map[string]interface{}{
		"product_id": productID,
		"quantity":   quantity,
	}

	response, err := c.restClient.WithToken().Post("/cart/items", body)
	if err != nil {
		return mcp.CallToolResult{}, fmt.Errorf("failed to add to cart: %w", err)
	}

	c.logger.WithField("response", string(response)).Info("Added to cart")

	var cartItem AddToCartResponse
	if err := json.Unmarshal(response, &cartItem); err != nil {
		return mcp.CallToolResult{}, fmt.Errorf("failed to parse response: %w", err)
	}

	if cartItem.Error != "" {
		return mcp.CallToolResult{
			Content: []mcp.Content{
				{
					Type: "text",
					Text: fmt.Sprintf("Error: %s", cartItem.Error),
				},
			},
			IsError: true,
		}, nil
	}

	return mcp.CallToolResult{
		Content: []mcp.Content{
			{
				Type: "text",
				Text: fmt.Sprintf("Successfully added product %d (quantity: %d) to cart", productID, quantity),
			},
		},
	}, nil
}

func (c *CartToolset) handleVIewCart() {}
