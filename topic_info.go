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
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2/proto/services"
	protobuf "google.golang.org/protobuf/proto"
)

// TopicInfo is the information about a topic
type TopicInfo struct {
	TopicMemo          string
	RunningHash        []byte
	SequenceNumber     uint64
	ExpirationTime     time.Time
	AdminKey           Key
	SubmitKey          Key
	AutoRenewPeriod    time.Duration
	AutoRenewAccountID *AccountID
	LedgerID           LedgerID
}

func _TopicInfoFromProtobuf(topicInfo *services.ConsensusTopicInfo) (TopicInfo, error) {
	if topicInfo == nil {
		return TopicInfo{}, errParameterNull
	}
	var err error
	tempTopicInfo := TopicInfo{
		TopicMemo:      topicInfo.Memo,
		RunningHash:    topicInfo.RunningHash,
		SequenceNumber: topicInfo.SequenceNumber,
		LedgerID:       LedgerID{topicInfo.LedgerId},
	}

	if autoRenewPeriod := topicInfo.AutoRenewPeriod; autoRenewPeriod != nil {
		tempTopicInfo.AutoRenewPeriod = _DurationFromProtobuf(topicInfo.AutoRenewPeriod)
	}

	if expirationTime := topicInfo.ExpirationTime; expirationTime != nil {
		tempTopicInfo.ExpirationTime = _TimeFromProtobuf(expirationTime)
	}

	if adminKey := topicInfo.AdminKey; adminKey != nil {
		tempTopicInfo.AdminKey, err = _KeyFromProtobuf(adminKey)
	}

	if submitKey := topicInfo.SubmitKey; submitKey != nil {
		tempTopicInfo.SubmitKey, err = _KeyFromProtobuf(submitKey)
	}

	if autoRenewAccount := topicInfo.AutoRenewAccount; autoRenewAccount != nil {
		tempTopicInfo.AutoRenewAccountID = _AccountIDFromProtobuf(autoRenewAccount)
	}

	return tempTopicInfo, err
}

func (topicInfo *TopicInfo) _ToProtobuf() *services.ConsensusTopicInfo {
	return &services.ConsensusTopicInfo{
		Memo:           topicInfo.TopicMemo,
		RunningHash:    topicInfo.RunningHash,
		SequenceNumber: topicInfo.SequenceNumber,
		ExpirationTime: &services.Timestamp{
			Seconds: int64(topicInfo.ExpirationTime.Second()),
			Nanos:   int32(topicInfo.ExpirationTime.Nanosecond()),
		},
		AdminKey:         topicInfo.AdminKey._ToProtoKey(),
		SubmitKey:        topicInfo.SubmitKey._ToProtoKey(),
		AutoRenewPeriod:  _DurationToProtobuf(topicInfo.AutoRenewPeriod),
		AutoRenewAccount: topicInfo.AutoRenewAccountID._ToProtobuf(),
		LedgerId:         topicInfo.LedgerID.ToBytes(),
	}
}

// ToBytes returns a byte array representation of the TopicInfo object
func (topicInfo TopicInfo) ToBytes() []byte {
	data, err := protobuf.Marshal(topicInfo._ToProtobuf())
	if err != nil {
		return make([]byte, 0)
	}

	return data
}

// TopicInfoFromBytes returns a TopicInfo object from a byte array
func TopicInfoFromBytes(data []byte) (TopicInfo, error) {
	if data == nil {
		return TopicInfo{}, errByteArrayNull
	}
	pb := services.ConsensusTopicInfo{}
	err := protobuf.Unmarshal(data, &pb)
	if err != nil {
		return TopicInfo{}, err
	}

	info, err := _TopicInfoFromProtobuf(&pb)
	if err != nil {
		return TopicInfo{}, err
	}

	return info, nil
}
