// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package types

import "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

type commonConsensusData struct {
	// use to verify block builder
	proposerAddress []byte

	// used to build next block and validate current payload timestamp
	nextPayloadTimestamp math.U64

	// useful for logging
	consensusBlkHeight math.U64
}

// GetProposerAddress returns the address of the validator
// selected by consensus to propose the block.
func (c *commonConsensusData) GetProposerAddress() []byte {
	return c.proposerAddress
}

// GetNextPayloadTimestamp returns the timestamp proposed by consensus
// for the next payload to be proposed. It is also used to bound
// current payload upon validation.
func (c *commonConsensusData) GetNextPayloadTimestamp() math.U64 {
	return c.nextPayloadTimestamp
}

// GetConsensusBlockHeight returns the height of consensus block,
// which may be different from execution payload height
// in some networks. Currently only used for logging.
func (c *commonConsensusData) GetConsensusBlockHeight() math.U64 {
	return c.consensusBlkHeight
}