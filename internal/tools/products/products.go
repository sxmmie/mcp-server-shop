package products

import (
	"github.com/sirupsen/logrus"
	"github.com/sxmmie/mcp-server-shop/internal/client"
	"github.com/sxmmie/mcp-server-shop/internal/mcp"
)

type ProductToolset struct {
	reg        *mcp.Registry
	logger     *logrus.Logger
	restClient *client.RestClient
}

func NewProductToolset(reg *mcp.Registry, logger *logrus.Logger, restClient *client.RestClient) *ProductToolset {
	pt := &ProductToolset{
		reg:        reg,
		logger:     logger,
		restClient: restClient,
	}

	pt.registerProductTools()

	return pt
}

func (p *ProductToolset) registerProductTools() {}
