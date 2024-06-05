// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package manager

import (
	"context"

	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/storage/pkg/pruner"
)

// DBManager is a manager for all pruners.
type DBManager[
	BeaconBlockT BeaconBlock,
	BlockEventT BlockEvent[BeaconBlockT],
	SubscriptionT Subscription,
] struct {
	pruners map[string]*pruner.Pruner[
		BeaconBlockT, BlockEventT, SubscriptionT,
	]
	logger log.Logger[any]
}

func NewDBManager[
	BeaconBlockT BeaconBlock,
	BlockEventT BlockEvent[BeaconBlockT],
	SubscriptionT Subscription,
](
	logger log.Logger[any],
	pruners ...*pruner.Pruner[BeaconBlockT, BlockEventT, SubscriptionT],
) (*DBManager[BeaconBlockT, BlockEventT, SubscriptionT], error) {
	m := &DBManager[
		BeaconBlockT, BlockEventT, SubscriptionT,
	]{
		logger: logger,
		pruners: make(map[string]*pruner.Pruner[
			BeaconBlockT, BlockEventT, SubscriptionT,
		]),
	}
	for _, p := range pruners {
		if _, ok := m.pruners[p.Name()]; ok {
			return nil, ErrDuplicatePruner
		}
		m.pruners[p.Name()] = p
	}
	return m, nil
}

// Name returns the name of the Basic Service.
func (m *DBManager[
	BeaconBlockT, BlockEventT, SubscriptionT,
]) Name() string {
	return "db-manager"
}

// TODO: fr implementation
func (m *DBManager[
	BeaconBlockT, BlockEventT, SubscriptionT,
]) Status() error {
	return nil
}

// TODO: fr implementation
func (m *DBManager[
	BeaconBlockT, BlockEventT, SubscriptionT,
]) WaitForHealthy(_ context.Context) {
}

// Start starts all pruners.
func (m *DBManager[
	BeaconBlockT, BlockEventT, SubscriptionT,
]) Start(ctx context.Context) error {
	for _, pruner := range m.pruners {
		pruner.Start(ctx)
	}
	return nil
}