/*
 * Copyright 2018 The Service Manager Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package notifications

import (
	"net/http"

	"github.com/Peripli/service-manager/pkg/ws"

	"github.com/Peripli/service-manager/pkg/web"
)

// Controller implements api.Controller by providing service plans API logic
type Controller struct {
	wsUpgrader ws.Upgrader
}

// Routes returns the routes for notifications
func (c *Controller) Routes() []web.Route {
	return []web.Route{
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   web.NotificationsURL,
			},
			Handler: c.handleWS,
		},
	}
}

// TODO: create the actual websocket handling and disable CRUD and List operations
func NewController(wsUpgrader ws.Upgrader) *Controller {
	return &Controller{
		wsUpgrader: wsUpgrader,
	}
}
