// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"math/big"
	"testing"

	"github.com/ava-labs/avalanche-cli/internal/mocks"
	"github.com/ava-labs/avalanche-cli/pkg/application"
	"github.com/ava-labs/avalanche-cli/pkg/statemachine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
)

var testAirdropAddress = common.HexToAddress("0x098B69E43b1720Bd12378225519d74e5F3aD0eA5")

func TestGetAllocationCustomUnits(t *testing.T) {
	assert := setupTest(t)
	app := application.New()
	mockPrompt := &mocks.Prompter{}
	app.Prompt = mockPrompt

	airdropInputAmount := new(big.Int)
	airdropInputAmount.SetString("1000000", 10)

	expectedAmount := new(big.Int)
	expectedAmount.SetString(defaultEvmAirdropAmount, 10)

	mockPrompt.On("CaptureList", mock.Anything, mock.Anything).Return(customAirdrop, nil)
	mockPrompt.On("CaptureAddress", mock.Anything).Return(testAirdropAddress, nil)
	mockPrompt.On("CapturePositiveBigInt", mock.Anything).Return(airdropInputAmount, nil)
	mockPrompt.On("CaptureNoYes", mock.Anything).Return(false, nil)

	alloc, direction, err := getEVMAllocation(app)
	assert.NoError(err)
	assert.Equal(direction, statemachine.Forward)

	assert.Equal(alloc[testAirdropAddress].Balance, expectedAmount)
}
