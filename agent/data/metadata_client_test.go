//go:build unit
// +build unit

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testKey = "test-key"
	testVal = "test-val"
)

func TestManageMetadata(t *testing.T) {
	testClient, cleanup := newTestClient(t)
	defer cleanup()

	require.NoError(t, testClient.SaveMetadata(testKey, testVal))

	val, err := testClient.GetMetadata(testKey)
	require.NoError(t, err)
	assert.Equal(t, testVal, val)
}
