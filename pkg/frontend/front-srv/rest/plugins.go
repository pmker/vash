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

// Package rest is a service for serving specific requests directly to frontend
package rest

import (
	"github.com/gobuffalo/packr"
	"github.com/pmker/vash/pkg/common/plugins"

	"encoding/gob"

	"github.com/pmker/vash/pkg/common"
	"github.com/pmker/vash/pkg/common/service"
	"github.com/pmker/vash/pkg/common/service/frontend"
	"github.com/pmker/leven/frontend/front-srv/rest/modifiers"
)

var BasePluginsBox = frontend.PluginBox{
	Box: packr.NewBox("../../../frontend/front-srv/assets"),
	Exposes: []string{
		"access.gateway",
		"access.homepage",
		"access.settings",
		"action.share",
		"auth.pydio",
		"authfront.session_login",
		"conf.pydio",
		"core.auth",
		"core.authfront",
		"core.conf",
		"core.mailer",
		"core.pydio",
		"core.uploader",
		"editor.libreoffice",
		"gui.ajax",
		"gui.mobile",
		"meta.user",
		"action.avatar",
		"action.compression",
		"action.demo_counter",
		"action.user",
		"core.activitystreams",
		"editor.browser",
		"editor.ckeditor",
		"editor.codemirror",
		"editor.diaporama",
		"editor.exif",
		"editor.infopanel",
		"editor.libreoffice",
		"editor.openlayer",
		"editor.pdfjs",
		"editor.soundmanager",
		"editor.text",
		"editor.video",
		"meta.comments",
		"meta.exif",
		"meta.simple_lock",
		"meta.versions",
		"uploader.html",
		"uploader.http",
	},
}

func init() {

	plugins.Register(func() {
		gob.Register(map[string]string{})

		frontend.RegisterRegModifier(modifiers.MetaUserRegModifier)
		frontend.RegisterPluginModifier(modifiers.MobileRegModifier)
		frontend.WrapAuthMiddleware(modifiers.LoginPasswordAuth)

		s := service.NewService(
			service.Name(common.SERVICE_REST_NAMESPACE_+common.SERVICE_FRONTEND),
			service.Tag(common.SERVICE_TAG_FRONTEND),
			service.Description("REST service for serving specific requests directly to frontend"),
			service.PluginBoxes(BasePluginsBox),
			service.WithWeb(func() service.WebHandler {
				return NewFrontendHandler()
			}),
		)
		// Make sure to have the WebSession wrapper happen before the policies
		// Exclude POST binaries for using Cookies as it's the only one subject to possible CSRF
		s.Init(service.WithWebSession("POST:/frontend/binaries"))
	})

}