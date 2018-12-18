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

package main

import (
	"github.com/pmker/vash/cmd"

	//// Making sure they are initialised first
	//_ "github.com/pmker/vash/discovery/consul"
	//_ "github.com/pmker/vash/discovery/nats"
	//
	//_ "github.com/pmker/vash/discovery/config/grpc"
	//_ "github.com/pmker/vash/discovery/config/rest"
	//_ "github.com/pmker/vash/discovery/install/rest"
	//_ "github.com/pmker/vash/discovery/update/grpc"
	//_ "github.com/pmker/vash/discovery/update/rest"
	//
	//_ "github.com/pmker/vash/broker/activity/grpc"
	//_ "github.com/pmker/vash/broker/activity/rest"
	//_ "github.com/pmker/vash/broker/chat/grpc"
	//_ "github.com/pmker/vash/broker/log/grpc"
	//_ "github.com/pmker/vash/broker/log/rest"
	//_ "github.com/pmker/vash/broker/mailer/grpc"
	//_ "github.com/pmker/vash/broker/mailer/rest"
	//_ "github.com/pmker/vash/frontend/front-srv/rest"
	//_ "github.com/pmker/vash/frontend/front-srv/web"
	//
	//_ "github.com/pmker/vash/data/changes/grpc"
	//_ "github.com/pmker/vash/data/changes/rest"
	//_ "github.com/pmker/vash/data/docstore/grpc"
	//_ "github.com/pmker/vash/data/key/grpc"
	//_ "github.com/pmker/vash/data/meta/grpc"
	//_ "github.com/pmker/vash/data/meta/rest"
	//_ "github.com/pmker/vash/data/source/index/grpc"
	//_ "github.com/pmker/vash/data/source/objects/grpc"
	//_ "github.com/pmker/vash/data/source/sync/grpc"
	//_ "github.com/pmker/vash/data/source/test"
	//_ "github.com/pmker/vash/data/templates/rest"
	//_ "github.com/pmker/vash/data/tree/grpc"
	//_ "github.com/pmker/vash/data/tree/rest"
	//_ "github.com/pmker/vash/data/versions/grpc"
	//
	//_ "github.com/pmker/vash/discovery/config/grpc"
	//_ "github.com/pmker/vash/discovery/config/rest"
	//
	//_ "github.com/pmker/vash/gateway/data"
	//_ "github.com/pmker/vash/gateway/dav"
	//_ "github.com/pmker/vash/gateway/micro"
	//_ "github.com/pmker/vash/gateway/proxy"
	//_ "github.com/pmker/vash/gateway/websocket/api"
	//_ "github.com/pmker/vash/gateway/wopi"
	//
	//_ "github.com/pmker/vash/data/search/grpc"
	//_ "github.com/pmker/vash/data/search/rest"
	//_ "github.com/pmker/vash/idm/acl/grpc"
	//_ "github.com/pmker/vash/idm/acl/rest"
	//_ "github.com/pmker/vash/idm/auth/grpc"
	//_ "github.com/pmker/vash/idm/auth/rest"
	//_ "github.com/pmker/vash/idm/auth/web"
	//_ "github.com/pmker/vash/idm/graph/rest"
	//_ "github.com/pmker/vash/idm/key/grpc"
	//_ "github.com/pmker/vash/idm/meta/grpc"
	//_ "github.com/pmker/vash/idm/meta/rest"
	//_ "github.com/pmker/vash/idm/policy/grpc"
	//_ "github.com/pmker/vash/idm/policy/rest"
	//_ "github.com/pmker/vash/idm/role/grpc"
	//_ "github.com/pmker/vash/idm/role/rest"
	//_ "github.com/pmker/vash/idm/share/rest"
	//_ "github.com/pmker/vash/idm/user/grpc"
	//_ "github.com/pmker/vash/idm/user/rest"
	//_ "github.com/pmker/vash/idm/workspace/grpc"
	//_ "github.com/pmker/vash/idm/workspace/rest"
	//_ "github.com/pmker/vash/scheduler/jobs/grpc"
	//_ "github.com/pmker/vash/scheduler/jobs/rest"
	//_ "github.com/pmker/vash/scheduler/tasks/grpc"
	//_ "github.com/pmker/vash/scheduler/timer/grpc"
	//
	//// All Actions for scheduler
	//_ "github.com/pmker/vash/broker/activity/actions"
	//_ "github.com/pmker/vash/scheduler/actions/archive"
	//_ "github.com/pmker/vash/scheduler/actions/changes"
	//_ "github.com/pmker/vash/scheduler/actions/cmd"
	//_ "github.com/pmker/vash/scheduler/actions/images"
	//_ "github.com/pmker/vash/scheduler/actions/scheduler"
	//_ "github.com/pmker/vash/scheduler/actions/tree"

	"github.com/pmker/vash/common"
)

func main() {
	common.PackageType = "Vash Dev"
	common.PackageLabel = "Vash Dev Home Edition"
	cmd.Execute()
}
