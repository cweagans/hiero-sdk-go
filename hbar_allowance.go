package hedera

/*-
 *
 * Hedera Go SDK
 *
 * Copyright (C) 2020 - 2022 Hedera Hashgraph, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import (
	"fmt"

	"github.com/hashgraph/hedera-protobufs-go/services"
)

type HbarAllowance struct {
	OwnerAccountID   *AccountID
	SpenderAccountID *AccountID
	Amount           int64
}

func NewHbarAllowance(ownerAccountID AccountID, spenderAccountID AccountID, amount int64) HbarAllowance { //nolint
	return HbarAllowance{
		OwnerAccountID:   &ownerAccountID,
		SpenderAccountID: &spenderAccountID,
		Amount:           amount,
	}
}

func _HbarAllowanceFromProtobuf(pb *services.CryptoAllowance) HbarAllowance {
	body := HbarAllowance{
		Amount: pb.Amount,
	}

	if pb.Spender != nil {
		body.SpenderAccountID = _AccountIDFromProtobuf(pb.Spender)
	}

	if pb.Owner != nil {
		body.OwnerAccountID = _AccountIDFromProtobuf(pb.Owner)
	}

	return body
}

func _HbarAllowanceFromGrantedProtobuf(pb *services.GrantedCryptoAllowance) HbarAllowance {
	body := HbarAllowance{
		Amount: pb.Amount,
	}

	if pb.Spender != nil {
		body.SpenderAccountID = _AccountIDFromProtobuf(pb.Spender)
	}

	return body
}

func _HbarWipeAllowanceFromProtobuf(pb *services.CryptoWipeAllowance) HbarAllowance {
	body := HbarAllowance{}

	if pb.Owner != nil {
		body.OwnerAccountID = _AccountIDFromProtobuf(pb.Owner)
	}

	return body
}

func (approval *HbarAllowance) _ToProtobuf() *services.CryptoAllowance {
	body := &services.CryptoAllowance{
		Amount: approval.Amount,
	}

	if approval.SpenderAccountID != nil {
		body.Spender = approval.SpenderAccountID._ToProtobuf()
	}

	if approval.OwnerAccountID != nil {
		body.Owner = approval.OwnerAccountID._ToProtobuf()
	}

	return body
}

func (approval *HbarAllowance) _ToGrantedProtobuf() *services.GrantedCryptoAllowance {
	body := &services.GrantedCryptoAllowance{
		Amount: approval.Amount,
	}

	if approval.SpenderAccountID != nil {
		body.Spender = approval.SpenderAccountID._ToProtobuf()
	}

	return body
}

func (approval *HbarAllowance) _ToWipeProtobuf() *services.CryptoWipeAllowance {
	body := &services.CryptoWipeAllowance{}

	if approval.OwnerAccountID != nil {
		body.Owner = approval.OwnerAccountID._ToProtobuf()
	}

	return body
}

func (approval *HbarAllowance) String() string {
	if approval.OwnerAccountID != nil && approval.SpenderAccountID != nil { //nolint
		return fmt.Sprintf("OwnerAccountID: %s, SpenderAccountID: %s, Amount: %s", approval.OwnerAccountID.String(), approval.SpenderAccountID.String(), HbarFromTinybar(approval.Amount).String())
	} else if approval.OwnerAccountID != nil {
		return fmt.Sprintf("OwnerAccountID: %s, Amount: %s", approval.OwnerAccountID.String(), HbarFromTinybar(approval.Amount).String())
	} else if approval.SpenderAccountID != nil {
		return fmt.Sprintf("SpenderAccountID: %s, Amount: %s", approval.SpenderAccountID.String(), HbarFromTinybar(approval.Amount).String())
	}

	return ""
}
