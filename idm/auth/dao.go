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

// Package auth provides all authentication related features
package auth

import (
	"github.com/pmker/vash/common/dao"
	"github.com/pmker/vash/common/proto/auth"
	"github.com/pmker/vash/common/sql"
)

type DAO interface {
	PutToken(t *auth.Token) error
	GetInfo(value string) (string, error)
	DeleteToken(t string) error
	ListTokens(offset int, count int) (chan *auth.Token, error)
}

type DexDAO interface {
	DexPruneOfflineSessions(c Config) (pruned int64, e error)
	DexDeleteOfflineSessions(c Config, userUuid string, sessionUuid string) error
}

func NewDAO(o dao.DAO) dao.DAO {
	switch v := o.(type) {
	case sql.DAO:
		return &dexSql{DAO: v}
	}
	return nil
}
