// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Amazon Software License (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/asl/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package model contains message struct for MDS/SSM messages.
package model

import (
	"github.com/aws/amazon-ssm-agent/agent/contracts"
)

// CancelPayload represents the json structure of a cancel command MDS message payload.
type CancelPayload struct {
	CancelMessageID string `json:"CancelMessageId"`
}

// SendCommandPayload parallels the structure of a send command MDS message payload.
type SendCommandPayload struct {
	Parameters         map[string]interface{}    `json:"Parameters"`
	DocumentContent    contracts.DocumentContent `json:"DocumentContent"`
	CommandID          string                    `json:"CommandId"`
	DocumentName       string                    `json:"DocumentName"`
	OutputS3KeyPrefix  string                    `json:"OutputS3KeyPrefix"`
	OutputS3BucketName string                    `json:"OutputS3BucketName"`
}

// SendReplyPayload represents the json structure of a reply sent to MDS.
type SendReplyPayload struct {
	AdditionalInfo      contracts.AdditionalInfo                  `json:"additionalInfo"`
	DocumentStatus      contracts.ResultStatus                    `json:"documentStatus"`
	DocumentTraceOutput string                                    `json:"documentTraceOutput"`
	RuntimeStatus       map[string]*contracts.PluginRuntimeStatus `json:"runtimeStatus"`
}

// PluginState represents information stored as interim state for any plugin
// This has both the configuration with which a plugin gets executed and a
// corresponding plugin result.
type PluginState struct {
	Configuration contracts.Configuration
	Result        contracts.PluginResult
	HasExecuted   bool
}

// DocumentInfo represents information stored as interim state for a document
type DocumentInfo struct {
	AdditionalInfo      contracts.AdditionalInfo
	CommandID           string
	Destination         string
	MessageID           string
	RunID               string
	CreatedDate         string
	DocumentName        string
	IsCommand           bool
	DocumentStatus      contracts.ResultStatus
	DocumentTraceOutput string
	RuntimeStatus       map[string]*contracts.PluginRuntimeStatus
	RunCount            int
	//ParsedDocumentContent string
	//RuntimeStatus
}

// CommandState represents information relevant to a command that gets executed by agent
type CommandState struct {
	DocumentInformation DocumentInfo
	PluginsInformation  map[string]PluginState
}

// CancelCommandState represents information relevant to a cancel-command that agent receives
// TODO  This might be revisited when Agent-cli is written to list previously executed commands
type CancelCommandState struct {
	Destination     string
	MessageID       string
	RunID           string
	CreatedDate     string
	Status          contracts.ResultStatus
	CancelMessageID string
	CancelCommandID string
	Payload         string
	DebugInfo       string
}
