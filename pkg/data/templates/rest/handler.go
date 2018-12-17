/*
 * Copyright (c) 2018. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

// Package rest exposes a simple API used by admins to query the whole tree directly without going through routers.
package rest

import (
	"github.com/emicklei/go-restful"

	"github.com/pmker/vash/pkg/common/proto/rest"
	"github.com/pmker/leven/data/templates"
)

type Handler struct {
	dao templates.DAO
}

// SwaggerTags list the names of the service tags declared in the swagger json implemented by this service
func (a *Handler) SwaggerTags() []string {
	return []string{"TemplatesService"}
}

// Filter returns a function to filter the swagger path
func (a *Handler) Filter() func(string) string {
	return func(s string) string {
		return s
	}
}

func (a *Handler) ListTemplates(req *restful.Request, rsp *restful.Response) {

	nodes := a.dao.List()
	response := &rest.ListTemplatesResponse{}
	for _, node := range nodes {
		response.Templates = append(response.Templates, node.AsTemplate())
	}

	rsp.WriteEntity(response)

}
