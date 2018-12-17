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

// Package scheduler provides default implementation for basic scheduler tasks.
package scheduler

import "github.com/pmker/leven/scheduler/actions"

func init() {

	manager := actions.GetActionsManager()

	manager.Register(pruneJobsActionName, func() actions.ConcreteAction {
		return &PruneJobsAction{}
	})

	actions.GetActionsManager().Register(fakeActionName, func() actions.ConcreteAction {
		return &FakeAction{}
	})

	actions.GetActionsManager().Register(fakeUserCreationActionName, func() actions.ConcreteAction {
		return &FakeUsersAction{}
	})

}