//go:build all || unit
// +build all unit

package hedera

/*-
 *
 * Hedera Go SDK
 *
 * Copyright (C) 2020 - 2024 Hedera Hashgraph, LLC
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
	"testing"

	"github.com/hashgraph/hedera-sdk-go/v2/proto/services"
	"github.com/stretchr/testify/assert"
)

func TestContractNonceInfoFromProtobuf(t *testing.T) {
	contractID := &ContractID{Shard: 0, Realm: 0, Contract: 123}
	nonce := int64(456)
	protobuf := &services.ContractNonceInfo{
		ContractId: contractID._ToProtobuf(),
		Nonce:      nonce,
	}

	result := _ContractNonceInfoFromProtobuf(protobuf)

	assert.Equal(t, contractID, result.ContractID)
	assert.Equal(t, nonce, result.Nonce)
}

func TestContractNonceInfoFromProtobuf_NilInput(t *testing.T) {
	result := _ContractNonceInfoFromProtobuf(nil)

	assert.Nil(t, result)
}
