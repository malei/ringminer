// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package orderbook

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/Loopring/ringminer/types"
)

var _ = (*cutoffIndexMarshaling)(nil)

func (c CutoffIndex) MarshalJSON() ([]byte, error) {
	type CutoffIndex struct {
		Address     types.Address `json:"address" gencodec:"required"`
		Timestamp   *types.Big    `json:"timestamp" gencodec:"required"`
		Cutoff      *types.Big    `json:"cutoff" gencodec:"required"`
		BlockNumber *types.Big    `json:"blocknumber" gencodec:"required"`
	}
	var enc CutoffIndex
	enc.Address = c.Address
	enc.Timestamp = (*types.Big)(c.Timestamp)
	enc.Cutoff = (*types.Big)(c.Cutoff)
	enc.BlockNumber = (*types.Big)(c.BlockNumber)
	return json.Marshal(&enc)
}

func (c *CutoffIndex) UnmarshalJSON(input []byte) error {
	type CutoffIndex struct {
		Address     *types.Address `json:"address" gencodec:"required"`
		Timestamp   *types.Big     `json:"timestamp" gencodec:"required"`
		Cutoff      *types.Big     `json:"cutoff" gencodec:"required"`
		BlockNumber *types.Big     `json:"blocknumber" gencodec:"required"`
	}
	var dec CutoffIndex
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Address == nil {
		return errors.New("missing required field 'address' for CutoffIndex")
	}
	c.Address = *dec.Address
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for CutoffIndex")
	}
	c.Timestamp = (*big.Int)(dec.Timestamp)
	if dec.Cutoff == nil {
		return errors.New("missing required field 'cutoff' for CutoffIndex")
	}
	c.Cutoff = (*big.Int)(dec.Cutoff)
	if dec.BlockNumber == nil {
		return errors.New("missing required field 'blocknumber' for CutoffIndex")
	}
	c.BlockNumber = (*big.Int)(dec.BlockNumber)
	return nil
}
