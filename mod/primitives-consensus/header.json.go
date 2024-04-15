// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package consensusprimitives

import (
	"encoding/json"

	"github.com/berachain/beacon-kit/mod/primitives"
)

// MarshalJSON marshals as JSON.
func (b BeaconBlockHeader) MarshalJSON() ([]byte, error) {
	type BeaconBlockHeader struct {
		Slot          primitives.Slot    `json:"slot"`
		ProposerIndex primitives.U64     `json:"proposerIndex"`
		ParentRoot    primitives.Bytes32 `json:"parentRoot"    ssz-size:"32"`
		StateRoot     primitives.Bytes32 `json:"stateRoot"     ssz-size:"32"`
		BodyRoot      primitives.Bytes32 `json:"bodyRoot"      ssz-size:"32"`
	}
	var enc BeaconBlockHeader
	enc.Slot = b.Slot
	enc.ProposerIndex = b.ProposerIndex
	enc.ParentRoot = b.ParentRoot
	enc.StateRoot = b.StateRoot
	enc.BodyRoot = b.BodyRoot
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (b *BeaconBlockHeader) UnmarshalJSON(input []byte) error {
	type BeaconBlockHeader struct {
		Slot          *primitives.Slot    `json:"slot"`
		ProposerIndex *primitives.U64     `json:"proposerIndex"`
		ParentRoot    *primitives.Bytes32 `json:"parentRoot"    ssz-size:"32"`
		StateRoot     *primitives.Bytes32 `json:"stateRoot"     ssz-size:"32"`
		BodyRoot      *primitives.Bytes32 `json:"bodyRoot"      ssz-size:"32"`
	}
	var dec BeaconBlockHeader
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Slot != nil {
		b.Slot = *dec.Slot
	}
	if dec.ProposerIndex != nil {
		b.ProposerIndex = *dec.ProposerIndex
	}
	if dec.ParentRoot != nil {
		b.ParentRoot = *dec.ParentRoot
	}
	if dec.StateRoot != nil {
		b.StateRoot = *dec.StateRoot
	}
	if dec.BodyRoot != nil {
		b.BodyRoot = *dec.BodyRoot
	}
	return nil
}
