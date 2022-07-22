// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"github.com/kukrer/savannahgo/ids"
	"github.com/kukrer/savannahgo/snow"
	"github.com/kukrer/savannahgo/vms"
)

var (
	// ID this VM should be referenced by
	ID = ids.ID{'e', 'v', 'm'}

	_ vms.Factory = &Factory{}
)

type Factory struct{}

func (f *Factory) New(*snow.Context) (interface{}, error) {
	return &VM{}, nil
}
