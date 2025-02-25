//go:build all || e2e
// +build all e2e

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

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestIntegrationTokenNftGetInfoByNftIDCanExecute(t *testing.T) { // nolint
	t.Parallel()
	env := NewIntegrationTestEnv(t)
	defer CloseIntegrationTestEnv(env, nil)

	newBalance := NewHbar(2)

	assert.Equal(t, 2*HbarUnits.Hbar._NumberOfTinybar(), newBalance.tinybar)

	tokenID, err := createNft(&env)
	require.NoError(t, err)

	metaData := []byte{50}

	mint, err := NewTokenMintTransaction().
		SetTokenID(tokenID).
		SetMetadata(metaData).
		Execute(env.Client)
	require.NoError(t, err)

	mintReceipt, err := mint.SetValidateStatus(true).GetReceipt(env.Client)
	require.NoError(t, err)

	nftID := tokenID.Nft(mintReceipt.SerialNumbers[0])

	info, err := NewTokenNftInfoQuery().
		SetNftID(nftID).
		Execute(env.Client)
	require.NoError(t, err)

	value := false
	for _, nftInfo := range info {
		if tokenID.String() == nftInfo.NftID.TokenID.String() {
			value = true
		}
	}
	assert.Truef(t, value, "token nft transfer transaction failed")
	assert.Equal(t, len(info), 1)
	assert.Equal(t, info[0].NftID, nftID)
	assert.Equal(t, info[0].Metadata[0], byte(50))
	parsedInfo, err := TokenNftInfoFromBytes(info[0].ToBytes())
	assert.NoError(t, err)
	assert.Equal(t, parsedInfo, info[0])
}
