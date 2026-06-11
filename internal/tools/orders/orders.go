package orders

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/sxmmie/mcp-server-shop/internal/client"
	"github.com/sxmmie/mcp-server-shop/internal/mcp"
)

type OrderToolset struct {
	reg        *mcp.Registry
	logger     *logrus.Logger
	restClient *client.RestClient
}

func NewOrderToolset(reg *mcp.Registry, loggger *logrus.Logger, restClient *client.RestClient) *OrderToolset {
	ot := &OrderToolset{
		reg:        reg,
		logger:     loggger,
		restClient: restClient,
	}

	ot.registerOrderTools()

	return ot
}

func (r *OrderToolset) registerOrderTools() {
	r.reg.Register(mcp.Tool{
		Name:        "place_order",
		Description: "Place a new order with the items in the shopping cart (requires authentication)",
		InputSchema: mcp.InputSchema{
			Type:       "object",
			Properties: map[string]mcp.Property{},
			Required:   []string{},
		},
	}, r.handlePlaceOrder)
}

func (r *OrderToolset) handlePlaceOrder(_ context.Context, _ map[string]any) (mcp.CallToolResult, error) {
	res, err := r.restClient.WithToken().Post("/orders", nil)
	if err != nil {
		r.logger.WithError(err).Error("Failed to place order")
		return mcp.NewToolCallError("Failed to place order"), nil
	}

	var orderRes OrderResponse
	if err := json.Unmarshal(res, &orderRes); err != nil {
		r.logger.WithError(err).Error("Failed to unmarshal order response")
		return mcp.NewToolCallError("Failed to parse order response"), nil
	}

	return mcp.CallToolResult{
		Content: []mcp.Content{
			{
				Type: "text",
				Text: fmt.Sprintf("Order placed successfully! Order ID: %d, Total Amount: %%.2f", orderRes.Data.Id, orderRes.Data.Total),
			},
		},
		IsError: false,
	}, nil
}
