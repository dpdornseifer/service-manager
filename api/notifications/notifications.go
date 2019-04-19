package notifications

import (
	"github.com/Peripli/service-manager/pkg/web"
)

func (c *Controller) handleWS(req *web.Request) (*web.Response, error) {
	rw := req.HijackResponseWriter()
	_, err := c.wsUpgrader.Upgrade(rw, req.Request, nil)
	if err != nil {
		return nil, err
	}

	return &web.Response{}, nil
}
