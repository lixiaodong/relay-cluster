// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

// MarshalJSON marshals as JSON.
func (o OrderTxRecord) MarshalJSON() ([]byte, error) {
	type OrderTxRecord struct {
		Owner       common.Address `json:"owner"`
		TxHash      common.Hash    `json:"tx_hash"`
		OrderHash   common.Hash    `json:"order_hash"`
		OrderStatus OrderStatus    `json:"order_status"`
		Nonce       int64          `json:"nonce"`
	}
	var enc OrderTxRecord
	enc.Owner = o.Owner
	enc.TxHash = o.TxHash
	enc.OrderHash = o.OrderHash
	enc.OrderStatus = o.OrderStatus
	enc.Nonce = o.Nonce
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (o *OrderTxRecord) UnmarshalJSON(input []byte) error {
	type OrderTxRecord struct {
		Owner       *common.Address `json:"owner"`
		TxHash      *common.Hash    `json:"tx_hash"`
		OrderHash   *common.Hash    `json:"order_hash"`
		OrderStatus *OrderStatus    `json:"order_status"`
		Nonce       *int64          `json:"nonce"`
	}
	var dec OrderTxRecord
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Owner != nil {
		o.Owner = *dec.Owner
	}
	if dec.TxHash != nil {
		o.TxHash = *dec.TxHash
	}
	if dec.OrderHash != nil {
		o.OrderHash = *dec.OrderHash
	}
	if dec.OrderStatus != nil {
		o.OrderStatus = *dec.OrderStatus
	}
	if dec.Nonce != nil {
		o.Nonce = *dec.Nonce
	}
	return nil
}