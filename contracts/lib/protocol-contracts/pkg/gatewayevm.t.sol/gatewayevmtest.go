// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayevm

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// RevertOptions is an auto generated low-level Go binding around an user-defined struct.
type RevertOptions struct {
	RevertAddress    common.Address
	CallOnRevert     bool
	AbortAddress     common.Address
	RevertMessage    []byte
	OnRevertGasLimit *big.Int
}

// StdInvariantFuzzArtifactSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzArtifactSelector struct {
	Artifact  string
	Selectors [][4]byte
}

// StdInvariantFuzzInterface is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzInterface struct {
	Addr      common.Address
	Artifacts []string
}

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}

// GatewayEVMTestMetaData contains all meta data concerning the GatewayEVMTest contract.
var GatewayEVMTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ASSET_HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfDestinationIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevert\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertFailsIfSenderIsNotTSS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteWithERC20FailsIfNotCustodyOrConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteWithERC20Token\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteWithMsgContextFailsIfDestinationIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteWithNonReturnApprovalToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteWithTokenRevertingOnZeroApproval\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNoParams\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNoParamsTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNonPayable\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveOnCallFails\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveOnCallUsingAuthCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveOnRevertFails\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceivePayable\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertWithERC20FailsIfNotCustodyOrConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetConnectorFailsIfSenderIsNotAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetConnectorFailsIfSet\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetConnectorFailsIfZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetCustodyFailsIfSenderIsNotAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetCustodyFailsIfSet\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetCustodyFailsIfZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgrade\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfSenderIsNotTSSUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpgradeAndForwardCallToReceivePayable\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositedAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedV2\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedCustodyTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x600c8054600160ff199182168117909255601f8054909116909117905560a06040526000608052603180546001600160a01b0319169055348015604157600080fd5b506201191f80620000536000396000f3fe608060405234801561001057600080fd5b50600436106103155760003560e01c806385226c81116101a7578063ce5871e1116100ee578063e63ab1e911610097578063fa18c09b11610071578063fa18c09b14610560578063fa7626d414610568578063fe7bdbb21461057557600080fd5b8063e63ab1e914610529578063e6afc79014610550578063f68bd1c01461055857600080fd5b8063dd51e82f116100c8578063dd51e82f14610511578063e20c9f7114610519578063e54182191461052157600080fd5b8063ce5871e1146104f9578063cebad2a614610501578063d38b66cd1461050957600080fd5b8063a783c78911610150578063b5508aa91161012a578063b5508aa9146104d1578063ba414fa6146104d9578063ccf20616146104f157600080fd5b8063a783c7891461049a578063b0464fdc146104c1578063b124dbf5146104c957600080fd5b8063a217fddf11610181578063a217fddf14610482578063a397ffd21461048a578063a56f7a4b1461049257600080fd5b806385226c811461043157806385f438c114610446578063916a17c61461046d57600080fd5b806343fd8c7d1161026b5780635d62c860116102145780637a380ebf116101ee5780637a380ebf146104195780637d7f772a146104215780637ebf744f1461042957600080fd5b80635d62c860146103c757806366d9a9a0146103fc5780636bdd212b1461041157600080fd5b806351010e491161024557806351010e49146103af578063520d989d146103b757806352ff5939146103bf57600080fd5b806343fd8c7d1461039757806344671b941461039f5780634df42da1146103a757600080fd5b80631855c337116102cd5780633801e58f116102a75780633801e58f1461037f5780633e5e3c23146103875780633f7286f41461038f57600080fd5b80631855c337146103445780631ed7831c1461034c5780632ade38801461036a57600080fd5b80630a9254e4116102fe5780630a9254e41461032c5780631226c65514610334578063130daf591461033c57600080fd5b806304b974f91461031a578063070f2ad014610324575b600080fd5b61032261057d565b005b61032261077c565b610322610954565b61032261136a565b6103226114c6565b61032261167e565b6103546117df565b604051610361919061b606565b60405180910390f35b610372611841565b604051610361919061b6a2565b610322611983565b610354611d85565b610354611de5565b610322611e45565b610322612191565b610322612454565b6103226125b0565b610322612773565b610322612eb0565b6103ee7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b981565b604051908152602001610361565b6104046133ff565b604051610361919061b808565b61032261356c565b61032261372b565b610322613d4f565b610322614040565b6104396141a1565b604051610361919061b902565b6103ee7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b610475614271565b604051610361919061b915565b6103ee600081565b610322614357565b610322614664565b6103ee7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b610475614788565b61032261486e565b610439614aea565b6104e1614bba565b6040519015158152602001610361565b610322614c8e565b610322614db2565b610322614f0e565b610322615122565b610322615268565b610354615a24565b610322615a84565b6103ee7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b610322615f67565b6103226161b2565b61032261656a565b601f546104e19060ff1681565b6103226168af565b6040805160048082526024820183526020820180516001600160e01b03167f6ed7016900000000000000000000000000000000000000000000000000000000179052602754925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa79261060f926001600160a01b031691016001600160a01b0391909116815260200190565b600060405180830381600087803b15801561062957600080fd5b505af115801561063d573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156106c657600080fd5b505af11580156106da573d6000803e3d6000fd5b5050601f546040517f38e225270000000000000000000000000000000000000000000000000000000081526101009091046001600160a01b031692506338e22527915061073190603190600090869060040161b9ac565b6000604051808303816000875af1158015610750573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610778919081019061bac6565b5050565b6027546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156107d557600080fd5b505af11580156107e9573d6000803e3d6000fd5b5050602754604080516001600160a01b039092166024830152600060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506108a0919060040161bafb565b600060405180830381600087803b1580156108ba57600080fd5b505af11580156108ce573d6000803e3d6000fd5b5050601f546025546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015261010090920416925063950837aa91506024015b600060405180830381600087803b15801561093a57600080fd5b505af115801561094e573d6000803e3d6000fd5b50505050565b602580547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602680548216611234179055602780548216615678179055602880549091166198761790556040516109b29061b507565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f080158015610a37573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556027546040519116908190610a809061b515565b6001600160a01b03928316815291166020820152604001604051809103906000f080158015610ab3573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081178255604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260275460255492519086169481019490945260448401929092529092166064820152600091610b92916084015b60408051601f198184030181529190526020810180516001600160e01b03167fc0c53b8b00000000000000000000000000000000000000000000000000000000179052616cf2565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0384811682029290921792839055604080518082018252601081527f4552433230437573746f64792e736f6c0000000000000000000000000000000060208201526027546025549251939095048416602484015293831660448301529091166064820152919250610c3591608401610b4a565b602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691909117909155604080518082018252601a81527f5a657461436f6e6e6563746f724e6f6e4e61746976652e736f6c0000000000006020820152601f546024805460275460255495516101009094048716928401929092528516604483015284166064820152919092166084820152919250610d279160a40160408051601f198184030181529190526020810180516001600160e01b03167ff8c8765e00000000000000000000000000000000000000000000000000000000179052616cf2565b602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038381169190911790915560275460405163ca669fa760e01b815291166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015610db257600080fd5b505af1158015610dc6573d6000803e3d6000fd5b5050602480546027546022546040517f15d57fd40000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152908216938101939093521692506315d57fd49150604401600060405180830381600087803b158015610e3957600080fd5b505af1158015610e4d573d6000803e3d6000fd5b50505050604051610e5d9061b523565b604051809103906000f080158015610e79573d6000803e3d6000fd5b50602080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556027546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b158015610f2557600080fd5b505af1158015610f39573d6000803e3d6000fd5b50506025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015610f9657600080fd5b505af1158015610faa573d6000803e3d6000fd5b5050601f546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015261010090920416925063ae7a3a6f9150602401600060405180830381600087803b15801561101557600080fd5b505af1158015611029573d6000803e3d6000fd5b5050601f546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526101009092041692506310188aef9150602401600060405180830381600087803b15801561109457600080fd5b505af11580156110a8573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561110a57600080fd5b505af115801561111e573d6000803e3d6000fd5b50506023546025546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42406024820152911692506340c10f199150604401600060405180830381600087803b15801561118d57600080fd5b505af11580156111a1573d6000803e3d6000fd5b50506023546021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526207a12060248201529116925063a9059cbb91506044016020604051808303816000875af1158015611215573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611239919061bb0e565b506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b1580156112ba57600080fd5b505af11580156112ce573d6000803e3d6000fd5b5050604080516080810182526025546001600160a01b039081168252602354811660208084019182526001848601908152855191820190955260008152606084018190528351602d80549185167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161781559251602e8054919095169116179092559251602f5590935090915060309061094e908261bbc4565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156113c357600080fd5b505af11580156113d7573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561146057600080fd5b505af1158015611474573d6000803e3d6000fd5b5050601f546040517fae7a3a6f000000000000000000000000000000000000000000000000000000008152600060048201526101009091046001600160a01b0316925063ae7a3a6f9150602401610920565b6040805160048082526024820183526020820180516001600160e01b03167fc9028a3600000000000000000000000000000000000000000000000000000000179052602754925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa792611558926001600160a01b031691016001600160a01b0391909116815260200190565b600060405180830381600087803b15801561157257600080fd5b505af1158015611586573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527ff3459a96000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b600060405180830381600087803b15801561161057600080fd5b505af1158015611624573d6000803e3d6000fd5b5050601f546020546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b03610100909304831694506338e225279350610731926031921690869060040161b9ac565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156116d757600080fd5b505af11580156116eb573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fb337f378000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b600060405180830381600087803b15801561177557600080fd5b505af1158015611789573d6000803e3d6000fd5b5050601f546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015261010090920416925063ae7a3a6f9150602401610920565b6060601680548060200260200160405190810160405280929190818152602001828054801561183757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611819575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020016000905b8282101561197a57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156119635783829060005260206000200180546118d69061bb30565b80601f01602080910402602001604051908101604052809291908181526020018280546119029061bb30565b801561194f5780601f106119245761010080835404028352916020019161194f565b820191906000526020600020905b81548152906001019060200180831161193257829003601f168201915b5050505050815260200190600101906118b7565b505050508152505081526020019060010190611865565b50505050905090565b60006040516119919061b531565b61199a9061bc83565b604051809103906000f0801580156119b6573d6000803e3d6000fd5b506025546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424060248201529192508216906340c10f1990604401600060405180830381600087803b158015611a2257600080fd5b505af1158015611a36573d6000803e3d6000fd5b50506025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015611a9357600080fd5b505af1158015611aa7573d6000803e3d6000fd5b50506021546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301529091169250639b19251a9150602401600060405180830381600087803b158015611b0b57600080fd5b505af1158015611b1f573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611b8157600080fd5b505af1158015611b95573d6000803e3d6000fd5b50506021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482018190529350908416915063a9059cbb906044016020604051808303816000875af1158015611c0a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c2e919061bb0e565b506040805160048082526024820183526020820180516001600160e01b03167f6ed7016900000000000000000000000000000000000000000000000000000000179052602154925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa792611cc1926001600160a01b031691016001600160a01b0391909116815260200190565b600060405180830381600087803b158015611cdb57600080fd5b505af1158015611cef573d6000803e3d6000fd5b5050601f546020546040517f7bbe9afa0000000000000000000000000000000000000000000000000000000081526001600160a01b0361010090930483169450637bbe9afa9350611d4e9260319289929116908890889060040161bcfc565b600060405180830381600087803b158015611d6857600080fd5b505af1158015611d7c573d6000803e3d6000fd5b50505050505050565b60606018805480602002602001604051908101604052809291908181526020018280548015611837576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611819575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015611837576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611819575050505050905090565b604080516001808252818301909252600091816020015b6060815260200190600190039081611e5c5790505090506040518060400160405280600f81526020017f48656c6c6f2c20466f756e64727921000000000000000000000000000000000081525081600081518110611ebc57611ebc61bd43565b6020908102919091010152604080516001808252818301909252600091816020016020820280368337019050509050602a81600081518110611f0057611f0061bd43565b6020908102919091010152604051600190600090611f269085908590859060240161bda4565b60408051601f198184030181529181526020820180516001600160e01b03167ff05b6abf00000000000000000000000000000000000000000000000000000000179052602554905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611fc457600080fd5b505af1158015611fd8573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506120ae919060040161bafb565b600060405180830381600087803b1580156120c857600080fd5b505af11580156120dc573d6000803e3d6000fd5b5050601f54604080516020808201835261012382525491517f38e225270000000000000000000000000000000000000000000000000000000081526101009093046001600160a01b0390811695506338e2252794506121439391921690869060040161bddc565b6000604051808303816000875af1158015612162573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261218a919081019061bac6565b5050505050565b604080516004808252602482018352602080830180516001600160e01b03167f6ed70169000000000000000000000000000000000000000000000000000000001790525492517ff30c7ba30000000000000000000000000000000000000000000000000000000081529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba392612231926001600160a01b03169160009187910161be0e565b600060405180830381600087803b15801561224b57600080fd5b505af115801561225f573d6000803e3d6000fd5b50506020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156122d857600080fd5b505af11580156122ec573d6000803e3d6000fd5b5050601f546040516101009091046001600160a01b031681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09250602001905060405180910390a1601f546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156123ae57600080fd5b505af11580156123c2573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f915061240890600090859061be36565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024016115f6565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156124ad57600080fd5b505af11580156124c1573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561254a57600080fd5b505af115801561255e573d6000803e3d6000fd5b5050601f546040517f950837aa000000000000000000000000000000000000000000000000000000008152600060048201526101009091046001600160a01b0316925063950837aa9150602401610920565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152670de0b6b3a76400009060009060250160408051808303601f190181529082905260275463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561265057600080fd5b505af1158015612664573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156126ed57600080fd5b505af1158015612701573d6000803e3d6000fd5b5050601f546040517fcb7ba8e50000000000000000000000000000000000000000000000000000000081526101009091046001600160a01b0316925063cb7ba8e59150849061275a906000908690602d9060040161bf2c565b6000604051808303818588803b158015611d6857600080fd5b602354602654604051620186a0602482018190526001600160a01b03938416604483015292909116606482015260009060840160408051601f198184030181529181526020820180516001600160e01b03167f357fc5a20000000000000000000000000000000000000000000000000000000017905260255490516303223eab60e11b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561284457600080fd5b505af1158015612858573d6000803e3d6000fd5b5050602354601f546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0361010090920482166004820152602481018790529116925063a9059cbb91506044016020604051808303816000875af11580156128cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128f3919061bb0e565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561295257600080fd5b505af1158015612966573d6000803e3d6000fd5b5050602354601f546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b036101009092048216600482015260009450911691506370a0823190602401602060405180830381865afa1580156129d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129fb919061bf60565b6023546026546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa158015612a66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a8a919061bf60565b9050612a968285616d11565b612aa1816000616d11565b6020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015612b1657600080fd5b505af1158015612b2a573d6000803e3d6000fd5b5050601f54602354602654604080516101009094046001600160a01b039081168552602085018b9052928316908401521660608201527f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609250608001905060405180910390a1601f546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015612c0957600080fd5b505af1158015612c1d573d6000803e3d6000fd5b50506020546023546040516001600160a01b039283169450911691507f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b738290612c68908890889061be36565b60405180910390a360215460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612cc957600080fd5b505af1158015612cdd573d6000803e3d6000fd5b5050601f546023546020546040517f7bbe9afa0000000000000000000000000000000000000000000000000000000081526001600160a01b0361010090940484169550637bbe9afa9450612d3f9360319381169216908a908a9060040161bcfc565b600060405180830381600087803b158015612d5957600080fd5b505af1158015612d6d573d6000803e3d6000fd5b5050602354601f546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b036101009092048216600482015260009450911691506370a0823190602401602060405180830381865afa158015612dde573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e02919061bf60565b6023546026546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa158015612e6d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e91919061bf60565b9050612e9e826000616d11565b612ea88187616d11565b505050505050565b601f546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb600482015261432160248201819052916000916101009091046001600160a01b0316906391d1485490604401602060405180830381865afa158015612f43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f67919061bb0e565b9050612f7281616d91565b601f546027546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b0391821660248201526000926101009004909116906391d1485490604401602060405180830381865afa158015613006573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061302a919061bb0e565b905061303581616e0b565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561308e57600080fd5b505af11580156130a2573d6000803e3d6000fd5b5050601f546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561311f57600080fd5b505af1158015613133573d6000803e3d6000fd5b5050602754604080516001600160a01b03928316815291871660208301527f3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579935001905060405180910390a1601f546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301526101009092049091169063950837aa90602401600060405180830381600087803b1580156131e557600080fd5b505af11580156131f9573d6000803e3d6000fd5b5050505061327d83601f60019054906101000a90046001600160a01b03166001600160a01b0316635b1125916040518163ffffffff1660e01b8152600401602060405180830381865afa158015613254573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613278919061bf79565b616e5d565b601f546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b038581166024830152610100909204909116906391d1485490604401602060405180830381865afa15801561330c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613330919061bb0e565b915061333b82616e0b565b601f546027546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b03918216602482015261010090920416906391d1485490604401602060405180830381865afa1580156133cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133ef919061bb0e565b90506133fa81616d91565b505050565b6060601b805480602002602001604051908101604052809291908181526020016000905b8282101561197a57838290600052602060002090600202016040518060400160405290816000820180546134569061bb30565b80601f01602080910402602001604051908101604052809291908181526020018280546134829061bb30565b80156134cf5780601f106134a4576101008083540402835291602001916134cf565b820191906000526020600020905b8154815290600101906020018083116134b257829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561355457602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116135165790505b50505050508152505081526020019060010190613423565b6040805160048082526024820183526020820180516001600160e01b03167f6ed7016900000000000000000000000000000000000000000000000000000000179052602754925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa7926135fe926001600160a01b031691016001600160a01b0391909116815260200190565b600060405180830381600087803b15801561361857600080fd5b505af115801561362c573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156136b557600080fd5b505af11580156136c9573d6000803e3d6000fd5b5050601f5460408051602081018252610123815290517f38e225270000000000000000000000000000000000000000000000000000000081526101009092046001600160a01b031693506338e22527925061073191600090869060040161bddc565b601f54604080518082018252601981527f4761746577617945564d55706772616465546573742e736f6c00000000000000602080830191909152825190810190925260008252602554613790936001600160a01b036101009091048116939116616ebe565b601f54604080517fdda79b7500000000000000000000000000000000000000000000000000000000815290516101009092046001600160a01b031691600091839163dda79b75916004808201926020929091908290030181865afa1580156137fc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613820919061bf79565b90506000601f60019054906101000a90046001600160a01b03166001600160a01b0316635b1125916040518163ffffffff1660e01b8152600401602060405180830381865afa158015613877573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061389b919061bf79565b604080518082018252600f81527f48656c6c6f2c20466f756e6472792100000000000000000000000000000000006020820152905191925090602a90600190670de0b6b3a7640000906000906138f99086908690869060240161bfa2565b60408051601f19818403018152918152602080830180516001600160e01b03167fe04d4f97000000000000000000000000000000000000000000000000000000001790525490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161399f916001600160a01b0391909116908690869060040161be0e565b600060405180830381600087803b1580156139b957600080fd5b505af11580156139cd573d6000803e3d6000fd5b50506020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015613a4657600080fd5b505af1158015613a5a573d6000803e3d6000fd5b505050507f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa601f60019054906101000a90046001600160a01b031683878787604051613aaa95949392919061bfcc565b60405180910390a1601f546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015613b2b57600080fd5b505af1158015613b3f573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e85469150613b84908590859061be36565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613be557600080fd5b505af1158015613bf9573d6000803e3d6000fd5b50506020546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b038c811694506338e2252793508692613c4f9260319290911690879060040161b9ac565b60006040518083038185885af1158015613c6d573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052613c96919081019061bac6565b50613cee87601f60019054906101000a90046001600160a01b03166001600160a01b031663dda79b756040518163ffffffff1660e01b8152600401602060405180830381865afa158015613254573d6000803e3d6000fd5b613d4586601f60019054906101000a90046001600160a01b03166001600160a01b0316635b1125916040518163ffffffff1660e01b8152600401602060405180830381865afa158015613254573d6000803e3d6000fd5b5050505050505050565b604080516001808252818301909252600091816020015b6060815260200190600190039081613d665790505090506040518060400160405280600f81526020017f48656c6c6f2c20466f756e64727921000000000000000000000000000000000081525081600081518110613dc657613dc661bd43565b6020908102919091010152604080516001808252818301909252600091816020016020820280368337019050509050602a81600081518110613e0a57613e0a61bd43565b6020908102919091010152604051600190600090613e309085908590859060240161bda4565b60408051601f198184030181529181526020820180516001600160e01b03167ff05b6abf00000000000000000000000000000000000000000000000000000000179052602554905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613ece57600080fd5b505af1158015613ee2573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250613fb8919060040161bafb565b600060405180830381600087803b158015613fd257600080fd5b505af1158015613fe6573d6000803e3d6000fd5b5050601f546020546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b03610100909304831694506338e225279350612143926031921690869060040161b9ac565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561409957600080fd5b505af11580156140ad573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527f0c8dc016000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b600060405180830381600087803b15801561413757600080fd5b505af115801561414b573d6000803e3d6000fd5b5050601f546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526101009092041692506310188aef9150602401610920565b6060601a805480602002602001604051908101604052809291908181526020016000905b8282101561197a5783829060005260206000200180546141e49061bb30565b80601f01602080910402602001604051908101604052809291908181526020018280546142109061bb30565b801561425d5780601f106142325761010080835404028352916020019161425d565b820191906000526020600020905b81548152906001019060200180831161424057829003601f168201915b5050505050815260200190600101906141c5565b6060601d805480602002602001604051908101604052809291908181526020016000905b8282101561197a5760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561433f57602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116143015790505b50505050508152505081526020019060010190614295565b6020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156143cc57600080fd5b505af11580156143e0573d6000803e3d6000fd5b5050604080518082018252600181527f31000000000000000000000000000000000000000000000000000000000000006020820152905161012393507fd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501925061444a91849161c00d565b60405180910390a1601f546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156144cb57600080fd5b505af11580156144df573d6000803e3d6000fd5b505060208054604080518082018252600181527f310000000000000000000000000000000000000000000000000000000000000093810193909352516001600160a01b0390911693507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f92506145579160009161be36565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156145b857600080fd5b505af11580156145cc573d6000803e3d6000fd5b5050601f5460408051602080820183526001600160a01b038781168352815484518086018652600181527f31000000000000000000000000000000000000000000000000000000000000009381019390935293517f38e22527000000000000000000000000000000000000000000000000000000008152610100909504811696506338e225279550610731949293169160040161bddc565b6027546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156146bd57600080fd5b505af11580156146d1573d6000803e3d6000fd5b5050602754604080516001600160a01b039092166024830152600060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061175b919060040161bafb565b6060601c805480602002602001604051908101604052809291908181526020016000905b8282101561197a5760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561485657602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116148185790505b505050505081525050815260200190600101906147ac565b602354602654604051620186a0602482018190526001600160a01b03938416604483015292909116606482015260009060840160408051601f198184030181529181526020820180516001600160e01b03167f357fc5a200000000000000000000000000000000000000000000000000000000179052602554905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561493f57600080fd5b505af1158015614953573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b960448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614a29919060040161bafb565b600060405180830381600087803b158015614a4357600080fd5b505af1158015614a57573d6000803e3d6000fd5b5050601f546023546026546040517faa0c0fc10000000000000000000000000000000000000000000000000000000081526101009093046001600160a01b03908116955063aa0c0fc19450614abc939281169291169087908790602d9060040161c02f565b600060405180830381600087803b158015614ad657600080fd5b505af1158015612ea8573d6000803e3d6000fd5b60606019805480602002602001604051908101604052809291908181526020016000905b8282101561197a578382906000526020600020018054614b2d9061bb30565b80601f0160208091040260200160405190810160405280929190818152602001828054614b599061bb30565b8015614ba65780601f10614b7b57610100808354040283529160200191614ba6565b820191906000526020600020905b815481529060010190602001808311614b8957829003601f168201915b505050505081526020019060010190614b0e565b60085460009060ff1615614bd2575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015614c63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614c87919061bf60565b1415905090565b6027546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015614ce757600080fd5b505af1158015614cfb573d6000803e3d6000fd5b5050602754604080516001600160a01b039092166024830152600060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061411d919060040161bafb565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015614e0b57600080fd5b505af1158015614e1f573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015614ea857600080fd5b505af1158015614ebc573d6000803e3d6000fd5b5050601f546040517f10188aef000000000000000000000000000000000000000000000000000000008152600060048201526101009091046001600160a01b031692506310188aef9150602401610920565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152670de0b6b3a76400009060009060250160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614fae57600080fd5b505af1158015614fc2573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250615098919060040161bafb565b600060405180830381600087803b1580156150b257600080fd5b505af11580156150c6573d6000803e3d6000fd5b5050601f546020546040517fcb7ba8e50000000000000000000000000000000000000000000000000000000081526001600160a01b036101009093048316945063cb7ba8e59350869261275a9216908690602d9060040161bf2c565b604080516020810182526000808252915161514291607b9160240161c00d565b60408051601f198184030181529181526020820180516001600160e01b03167f676cc05400000000000000000000000000000000000000000000000000000000179052602754905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156151e057600080fd5b505af11580156151f4573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fed699775000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024016115f6565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156152c157600080fd5b505af11580156152d5573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506153ab919060040161bafb565b600060405180830381600087803b1580156153c557600080fd5b505af11580156153d9573d6000803e3d6000fd5b50505050601f60019054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561542d57600080fd5b505af1158015615441573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561549e57600080fd5b505af11580156154b2573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250615588919060040161bafb565b600060405180830381600087803b1580156155a257600080fd5b505af11580156155b6573d6000803e3d6000fd5b50505050601f60019054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561560a57600080fd5b505af115801561561e573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561567b57600080fd5b505af115801561568f573d6000803e3d6000fd5b50505050601f60019054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156156e357600080fd5b505af11580156156f7573d6000803e3d6000fd5b5050604080516004808252602480830184526020830180516001600160e01b03167f6ed701690000000000000000000000000000000000000000000000000000000017905292517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd93c06650000000000000000000000000000000000000000000000000000000091810191909152909350737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09101600060405180830381600087803b1580156157c457600080fd5b505af11580156157d8573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561583557600080fd5b505af1158015615849573d6000803e3d6000fd5b5050601f546020546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b03610100909304831694506338e2252793506158a3926031921690869060040161b9ac565b6000604051808303816000875af11580156158c2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526158ea919081019061bac6565b5060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561594457600080fd5b505af1158015615958573d6000803e3d6000fd5b50505050601f60019054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156159ac57600080fd5b505af11580156159c0573d6000803e3d6000fd5b50506020546040517ff30c7ba3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f30c7ba39250612231916001600160a01b031690600090869060040161be0e565b60606015805480602002602001604051908101604052809291908181526020018280548015611837576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611819575050505050905090565b6000604051615a929061b53f565b615a9b9061c084565b604051809103906000f080158015615ab7573d6000803e3d6000fd5b506025546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424060248201529192508216906340c10f1990604401600060405180830381600087803b158015615b2357600080fd5b505af1158015615b37573d6000803e3d6000fd5b50506025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015615b9457600080fd5b505af1158015615ba8573d6000803e3d6000fd5b50506021546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301529091169250639b19251a9150602401600060405180830381600087803b158015615c0c57600080fd5b505af1158015615c20573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015615c8257600080fd5b505af1158015615c96573d6000803e3d6000fd5b50506021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482018190529350908416915063a9059cbb906044016020604051808303816000875af1158015615d0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615d2f919061bb0e565b5060408051600480825260248083018452602080840180516001600160e01b03167f6ed70169000000000000000000000000000000000000000000000000000000001790525493516381bad6f360e01b815260019281018390529081018290526044810182905260648101919091526001600160a01b03909216608483015290737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015615de857600080fd5b505af1158015615dfc573d6000803e3d6000fd5b5050601f546040516101009091046001600160a01b031681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09250602001905060405180910390a1601f546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015615ebe57600080fd5b505af1158015615ed2573d6000803e3d6000fd5b50506020546040516001600160a01b03918216935090861691507f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b738290615f1b908690869061be36565b60405180910390a360215460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401611cc1565b602354602654604051620186a0602482018190526001600160a01b03938416604483015292909116606482015260009060840160408051601f198184030181529181526020820180516001600160e01b03167f357fc5a200000000000000000000000000000000000000000000000000000000179052602554905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561603857600080fd5b505af115801561604c573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b960448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250616122919060040161bafb565b600060405180830381600087803b15801561613c57600080fd5b505af1158015616150573d6000803e3d6000fd5b5050601f546023546026546040517f7bbe9afa0000000000000000000000000000000000000000000000000000000081526001600160a01b0361010090940484169550637bbe9afa9450614abc9360319381169216908890889060040161bcfc565b604080516001808252818301909252600091816020015b60608152602001906001900390816161c95790505090506040518060400160405280600f81526020017f48656c6c6f2c20466f756e647279210000000000000000000000000000000000815250816000815181106162295761622961bd43565b6020908102919091010152604080516001808252818301909252600091816020016020820280368337019050509050602a8160008151811061626d5761626d61bd43565b60209081029190910101526040516001906000906162939085908590859060240161bda4565b60408051601f19818403018152918152602080830180516001600160e01b03167ff05b6abf000000000000000000000000000000000000000000000000000000001790525490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161633a916001600160a01b039190911690600090869060040161be0e565b600060405180830381600087803b15801561635457600080fd5b505af1158015616368573d6000803e3d6000fd5b50506020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156163e157600080fd5b505af11580156163f5573d6000803e3d6000fd5b505050507f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146601f60019054906101000a90046001600160a01b0316858585604051616443949392919061c0fd565b60405180910390a1601f546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156164c457600080fd5b505af11580156164d8573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f915061651e90600090859061be36565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401613fb8565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152670de0b6b3a76400009060009060250160408051601f198184030181529190526020549091506001600160a01b0316316165ce816000616d11565b6020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561664357600080fd5b505af1158015616657573d6000803e3d6000fd5b505050507f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b601f60019054906101000a90046001600160a01b0316602d6040516166a292919061c145565b60405180910390a1601f546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561672357600080fd5b505af1158015616737573d6000803e3d6000fd5b5050602054604051600093506001600160a01b0390911691507fde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e0359061678a90670de0b6b3a7640000908790602d9061c167565b60405180910390a360275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156167eb57600080fd5b505af11580156167ff573d6000803e3d6000fd5b5050601f546020546040517fcb7ba8e50000000000000000000000000000000000000000000000000000000081526001600160a01b036101009093048316945063cb7ba8e59350879261685b9216908790602d9060040161bf2c565b6000604051808303818588803b15801561687457600080fd5b505af1158015616888573d6000803e3d6000fd5b50506020546001600160a01b031631925061094e9150829050670de0b6b3a7640000616d11565b60408051808201909152600f81527f48656c6c6f2c20466f756e64727921000000000000000000000000000000000060208083019190915254602a90600190670de0b6b3a76400009061690e906000906001600160a01b031631616d11565b60008484846040516024016169259392919061bfa2565b60408051601f19818403018152918152602080830180516001600160e01b03167fe04d4f97000000000000000000000000000000000000000000000000000000001790525490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba3916169d3916001600160a01b039190911690670de0b6b3a764000090869060040161be0e565b600060405180830381600087803b1580156169ed57600080fd5b505af1158015616a01573d6000803e3d6000fd5b50506020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015616a7a57600080fd5b505af1158015616a8e573d6000803e3d6000fd5b505050507f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa601f60019054906101000a90046001600160a01b031683878787604051616ade95949392919061bfcc565b60405180910390a1601f546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015616b5f57600080fd5b505af1158015616b73573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f9150616bc090670de0b6b3a764000090859061be36565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015616c2157600080fd5b505af1158015616c35573d6000803e3d6000fd5b5050601f546020546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b03610100909304831694506338e2252793508692616c9292603192911690879060040161b9ac565b60006040518083038185885af1158015616cb0573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052616cd9919081019061bac6565b5060205461218a9083906001600160a01b031631616d11565b6000616cfc61b54d565b616d07848483616ed3565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015b60006040518083038186803b158015616d7d57600080fd5b505afa158015612ea8573d6000803e3d6000fd5b6040517fa59828850000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a5982885906024015b60006040518083038186803b158015616df757600080fd5b505afa15801561218a573d6000803e3d6000fd5b6040517f0c9fd5810000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90630c9fd58190602401616ddf565b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f690604401616d65565b616ec661b54d565b61218a8585858486616f4e565b600080616ee08584617035565b9050616f436040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f78790000008152508286604051602001616f2e92919061c00d565b60405160208183030381529060405285617041565b9150505b9392505050565b6040516303223eab60e11b81526001600160a01b03821660048201528190737109709ecfa91a80626ff3989d68f67f5b1dd12d9081906306447d5690602401600060405180830381600087803b158015616fa757600080fd5b505af1925050508015616fb8575060015b616fcd57616fc88787878761706f565b611d7c565b616fd98787878761706f565b806001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561701457600080fd5b505af1158015617028573d6000803e3d6000fd5b5050505050505050505050565b6000616f478383617088565b60c081015151600090156170655761705e84848460c001516170a3565b9050616f47565b61705e8484617249565b600061707b8483617334565b905061218a858285617340565b6000617094838361770a565b616f4783836020015184617041565b6000806170ae617716565b905060006170bc86836177e9565b905060006170d38260600151836020015185617c8f565b905060006170e383838989617ea1565b905060006170f082618d1e565b602081015181519192509060030b156171635789826040015160405160200161711a92919061c180565b60408051601f19818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261715a9160040161bafb565b60405180910390fd5b60006171a66040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001618eed565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906171f990849060040161bafb565b602060405180830381865afa158015617216573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061723a919061bf79565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc9259061729e90879060040161bafb565b600060405180830381865afa1580156172bb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526172e3919081019061bac6565b9050600061731182856040516020016172fd92919061c201565b6040516020818303038152906040526190ed565b90506001600160a01b038116616d0757848460405160200161711a92919061c230565b60006170948383619100565b6040517f667f9d700000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90600090829063667f9d7090604401602060405180830381865afa1580156173dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617400919061bf60565b9050806175a75760006174128661910c565b604080518082018252600581527f352e302e300000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061749d905b604080518082018252600080825260209182015281518083019092528451825280850190820152906191ef565b806174a9575060008451115b1561752c576040517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b03871690634f1ef286906174f5908890889060040161c00d565b600060405180830381600087803b15801561750f57600080fd5b505af1158015617523573d6000803e3d6000fd5b505050506175a1565b6040517f3659cfe60000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152871690633659cfe690602401600060405180830381600087803b15801561758857600080fd5b505af115801561759c573d6000803e3d6000fd5b505050505b5061218a565b8060006175b38261910c565b604080518082018252600581527f352e302e300000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061761590617470565b80617621575060008551115b156176a6576040517f9623609d0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690639623609d9061766f908a908a908a9060040161c2db565b600060405180830381600087803b15801561768957600080fd5b505af115801561769d573d6000803e3d6000fd5b50505050611d7c565b6040517f99a88ec40000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015287811660248301528316906399a88ec490604401600060405180830381600087803b15801561701457600080fd5b61077882826000619203565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c9061779d90849060040161c30c565b600060405180830381865afa1580156177ba573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526177e2919081019061c353565b9250505090565b61781b6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d90506178666040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b61786f85619306565b6020820152600061787f866196eb565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa1580156178c1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526178e9919081019061c353565b86838560200151604051602001617903949392919061c39c565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb119061795b90859060040161bafb565b600060405180830381865afa158015617978573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526179a0919081019061c353565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906179e890849060040161c4a0565b602060405180830381865afa158015617a05573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617a29919061bb0e565b617a3e578160405160200161711a919061c4f2565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890617a8390849060040161c584565b600060405180830381865afa158015617aa0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617ac8919081019061c353565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f690617b0f90849060040161c5d6565b602060405180830381865afa158015617b2c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617b50919061bb0e565b15617be5576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890617b9a90849060040161c5d6565b600060405180830381865afa158015617bb7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617bdf919081019061c353565b60408501525b846001600160a01b03166349c4fac8828660000151604051602001617c0a919061c628565b6040516020818303038152906040526040518363ffffffff1660e01b8152600401617c3692919061c694565b600060405180830381865afa158015617c53573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617c7b919081019061c353565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b6060815260200190600190039081617cab5790505090506040518060400160405280600481526020017f677265700000000000000000000000000000000000000000000000000000000081525081600081518110617d0b57617d0b61bd43565b60200260200101819052506040518060400160405280600381526020017f2d726c000000000000000000000000000000000000000000000000000000000081525081600181518110617d5f57617d5f61bd43565b602002602001018190525084604051602001617d7b919061c6b9565b60405160208183030381529060405281600281518110617d9d57617d9d61bd43565b602002602001018190525082604051602001617db9919061c725565b60405160208183030381529060405281600381518110617ddb57617ddb61bd43565b60200260200101819052506000617df182618d1e565b602080820151604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000008185019081528251808401845260008082529086015282518084019093529051825292810192909252919250617e82906040805180820182526000808252602091820152815180830190925284518252808501908201529061996e565b617e97578560405160200161711a919061c766565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015617ef1565b511590565b61806557826020015115617fad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a40161715a565b8260c0015115618065576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a40161715a565b6040805160ff8082526120008201909252600091816020015b606081526020019060019003908161807e57905050905060006040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806180d99061c826565b935060ff16815181106180ee576180ee61bd43565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e370000000000000000000000000000000000000081525060405160200161813f919061c845565b60405160208183030381529060405282828061815a9061c826565b935060ff168151811061816f5761816f61bd43565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806181bc9061c826565b935060ff16815181106181d1576181d161bd43565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d6500000000000000000000000000000000000081525082828061821e9061c826565b935060ff16815181106182335761823361bd43565b6020026020010181905250876020015182828061824f9061c826565b935060ff16815181106182645761826461bd43565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e7472616374506174680000000000000000000000000000000000008152508282806182b19061c826565b935060ff16815181106182c6576182c661bd43565b6020908102919091010152875182826182de8161c826565b935060ff16815181106182f3576182f361bd43565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e496400000000000000000000000000000000000000000000008152508282806183409061c826565b935060ff16815181106183555761835561bd43565b6020026020010181905250618369466199cf565b82826183748161c826565b935060ff16815181106183895761838961bd43565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806183d69061c826565b935060ff16815181106183eb576183eb61bd43565b6020026020010181905250868282806184039061c826565b935060ff16815181106184185761841861bd43565b602090810291909101015285511561853f5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826184698161c826565b935060ff168151811061847e5761847e61bd43565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906184ce90899060040161bafb565b600060405180830381865afa1580156184eb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618513919081019061c353565b828261851e8161c826565b935060ff16815181106185335761853361bd43565b60200260200101819052505b84602001511561860f5760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826185888161c826565b935060ff168151811061859d5761859d61bd43565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806185ea9061c826565b935060ff16815181106185ff576185ff61bd43565b60200260200101819052506187d6565b618647617eec8660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6186da5760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261868a8161c826565b935060ff168151811061869f5761869f61bd43565b60200260200101819052508460a001516040516020016186bf919061c6b9565b6040516020818303038152906040528282806185ea9061c826565b8460c0015115801561871d57506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261871b90511590565b155b156187d65760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826187618161c826565b935060ff16815181106187765761877661bd43565b602002602001018190525061878a88619a6f565b60405160200161879a919061c6b9565b6040516020818303038152906040528282806187b59061c826565b935060ff16815181106187ca576187ca61bd43565b60200260200101819052505b6040808601518151808301835260008082526020918201528251808401909352815183529081019082015261880a90511590565b61889f5760408051808201909152600b81527f2d2d72656c6179657249640000000000000000000000000000000000000000006020820152828261884d8161c826565b935060ff16815181106188625761886261bd43565b6020026020010181905250846040015182828061887e9061c826565b935060ff16815181106188935761889361bd43565b60200260200101819052505b6060850151156189c05760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826188e88161c826565b935060ff16815181106188fd576188fd61bd43565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa15801561896c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618994919081019061c353565b828261899f8161c826565b935060ff16815181106189b4576189b461bd43565b60200260200101819052505b60e08501515115618a675760408051808201909152600a81527f2d2d6761734c696d69740000000000000000000000000000000000000000000060208201528282618a0a8161c826565b935060ff1681518110618a1f57618a1f61bd43565b6020026020010181905250618a3b8560e00151600001516199cf565b8282618a468161c826565b935060ff1681518110618a5b57618a5b61bd43565b60200260200101819052505b60e08501516020015115618b115760408051808201909152600a81527f2d2d67617350726963650000000000000000000000000000000000000000000060208201528282618ab48161c826565b935060ff1681518110618ac957618ac961bd43565b6020026020010181905250618ae58560e00151602001516199cf565b8282618af08161c826565b935060ff1681518110618b0557618b0561bd43565b60200260200101819052505b60e08501516040015115618bbb5760408051808201909152600e81527f2d2d6d617846656550657247617300000000000000000000000000000000000060208201528282618b5e8161c826565b935060ff1681518110618b7357618b7361bd43565b6020026020010181905250618b8f8560e00151604001516199cf565b8282618b9a8161c826565b935060ff1681518110618baf57618baf61bd43565b60200260200101819052505b60e08501516060015115618c655760408051808201909152601681527f2d2d6d61785072696f726974794665655065724761730000000000000000000060208201528282618c088161c826565b935060ff1681518110618c1d57618c1d61bd43565b6020026020010181905250618c398560e00151606001516199cf565b8282618c448161c826565b935060ff1681518110618c5957618c5961bd43565b60200260200101819052505b60008160ff1667ffffffffffffffff811115618c8357618c8361b9de565b604051908082528060200260200182016040528015618cb657816020015b6060815260200190600190039081618ca15790505b50905060005b8260ff168160ff161015618d0f57838160ff1681518110618cdf57618cdf61bd43565b6020026020010151828260ff1681518110618cfc57618cfc61bd43565b6020908102919091010152600101618cbc565b5093505050505b949350505050565b618d456040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c91618dcb9186910161c8b0565b600060405180830381865afa158015618de8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618e10919081019061c353565b90506000618e1e868361a55e565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b8152600401618e4e919061b902565b6000604051808303816000875af1158015618e6d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618e95919081019061c8f7565b805190915060030b15801590618eae5750602081015151155b8015618ebd5750604081015151155b15617e975781600081518110618ed557618ed561bd43565b602002602001015160405160200161711a919061c9ad565b60606000618f228560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528651825280870190820152909150618f599082905b9061a6b3565b156190b6576000618fd682618fd084618fca618f9c8a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061a6da565b9061a73c565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061903a90829061a6b3565b156190a457604080518082018252600181527f0a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526190a1905b829061a7c1565b90505b6190ad8161a7e7565b92505050616f47565b82156190cf57848460405160200161711a92919061cb99565b5050604080516020810190915260008152616f47565b509392505050565b6000808251602084016000f09392505050565b61077882826001619203565b60408051600481526024810182526020810180516001600160e01b03167fad3cb1cc00000000000000000000000000000000000000000000000000000000179052905160609160009182916001600160a01b0386169161916c919061cc40565b600060405180830381855afa9150503d80600081146191a7576040519150601f19603f3d011682016040523d82523d6000602084013e6191ac565b606091505b50915091508180156191bf575060208151115b156191d85780806020019051810190618d16919061c353565b505060408051602081019091526000815292915050565b60006191fb838361a850565b159392505050565b8160a001511561921257505050565b600061921f84848461a92b565b9050600061922c82618d1e565b602081015181519192509060030b1580156192c85750604080518082018252600781527f5355434345535300000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526192c890604080518082018252600080825260209182015281518083019092528451825280850190820152618f53565b156192d557505050505050565b604082015151156192f557816040015160405160200161711a919061cc5c565b8060405160200161711a919061ccba565b6060600061933b8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506193a0905b829061996e565b1561940f57604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616f479061940a90839061aec6565b61a7e7565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619471905b829061af50565b60010361953e57604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526194d79061909a565b50604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616f479061940a905b839061a7c1565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261959d90619399565b156196d457604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061960590839061afea565b905060008160018351619618919061cd25565b815181106196285761962861bd43565b602002602001015190506196cb61940a61969e6040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925285518252808601908201529061aec6565b95945050505050565b8260405160200161711a919061cd38565b50919050565b606060006197208360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061978290619399565b1561979057616f478161a7e7565b604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526197ef9061946a565b60010361985957604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616f479061940a90619537565b604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526198b890619399565b156196d457604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061992090839061afea565b905060018151111561995c57806002825161993b919061cd25565b8151811061994b5761994b61bd43565b602002602001015192505050919050565b508260405160200161711a919061cd38565b80518251600091111561998357506000616d0b565b815183516020850151600092916199999161ce16565b6199a3919061cd25565b9050826020015181036199ba576001915050616d0b565b82516020840151819020912014905092915050565b606060006199dc8361b08f565b600101905060008167ffffffffffffffff8111156199fc576199fc61b9de565b6040519080825280601f01601f191660200182016040528015619a26576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084619a3057509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e5345440000000000000000000000000000000000000000000081840190815285518087018752838152840192909252845180860190955251845290830152606091619afb905b82906191ef565b15619b3b57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e7365000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619b9a90619af4565b15619bda57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d4954000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619c3990619af4565b15619c7957505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c79000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619cd890619af4565b80619d3d5750604080518082018252601081527f47504c2d322e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619d3d90619af4565b15619d7d57505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c79000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619ddc90619af4565b80619e415750604080518082018252601081527f47504c2d332e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619e4190619af4565b15619e8157505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619ee090619af4565b80619f455750604080518082018252601181527f4c47504c2d322e312d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619f4590619af4565b15619f8557505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619fe490619af4565b8061a0495750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a04990619af4565b1561a08957505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a0e890619af4565b1561a12857505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a18790619af4565b1561a1c757505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a22690619af4565b1561a26657505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a2c590619af4565b1561a30557505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a36490619af4565b1561a3a457505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a40390619af4565b8061a4685750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a46890619af4565b1561a4a857505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a50790619af4565b1561a54757505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b6040808401518451915161711a929060200161ce29565b60608060005b845181101561a5e9578185828151811061a5805761a58061bd43565b602002602001015160405160200161a59992919061c201565b60405160208183030381529060405291506001855161a5b8919061cd25565b811461a5e1578160405160200161a5cf919061cf92565b60405160208183030381529060405291505b60010161a564565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161a602579050509050838160008151811061a62d5761a62d61bd43565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061a6815761a68161bd43565b6020026020010181905250818160028151811061a6a05761a6a061bd43565b6020908102919091010152949350505050565b602080830151835183519284015160009361a6d1929184919061b171565b14159392505050565b6040805180820190915260008082526020820152600061a70c846000015185602001518560000151866020015161b282565b905083602001518161a71e919061cd25565b8451859061a72d90839061cd25565b90525060208401525090919050565b604080518082019091526000808252602082015281518351101561a761575081616d0b565b602080830151908401516001911461a7885750815160208481015190840151829020919020145b801561a7b95782518451859061a79f90839061cd25565b905250825160208501805161a7b590839061ce16565b9052505b509192915050565b604080518082019091526000808252602082015261a7e083838361b3a2565b5092915050565b60606000826000015167ffffffffffffffff81111561a8085761a80861b9de565b6040519080825280601f01601f19166020018201604052801561a832576020820181803683370190505b509050600060208201905061a7e0818560200151866000015161b44d565b815181516000919081111561a863575081515b6020808501519084015160005b8381101561a91c578251825180821461a8ec57600019602087101561a8cb5760018461a89d89602061cd25565b61a8a7919061ce16565b61a8b290600861cfd3565b61a8bd90600261d0d1565b61a8c7919061cd25565b1990505b818116838216818103911461a8e9579750616d0b9650505050505050565b50505b61a8f760208661ce16565b945061a90460208561ce16565b9350505060208161a915919061ce16565b905061a870565b5084518651617e97919061d0dd565b6060600061a937617716565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161a95457905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061a9af9061c826565b935060ff168151811061a9c45761a9c461bd43565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161aa15919061d0fd565b60405160208183030381529060405282828061aa309061c826565b935060ff168151811061aa455761aa4561bd43565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061aa929061c826565b935060ff168151811061aaa75761aaa761bd43565b60200260200101819052508260405160200161aac3919061c725565b60405160208183030381529060405282828061aade9061c826565b935060ff168151811061aaf35761aaf361bd43565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061ab409061c826565b935060ff168151811061ab555761ab5561bd43565b602002602001018190525061ab6a878461b4c7565b828261ab758161c826565b935060ff168151811061ab8a5761ab8a61bd43565b60209081029190910101528551511561ac365760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261abdc8161c826565b935060ff168151811061abf15761abf161bd43565b602002602001018190525061ac0a86600001518461b4c7565b828261ac158161c826565b935060ff168151811061ac2a5761ac2a61bd43565b60200260200101819052505b85608001511561aca45760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261ac7f8161c826565b935060ff168151811061ac945761ac9461bd43565b602002602001018190525061ad0a565b841561ad0a5760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261ace98161c826565b935060ff168151811061acfe5761acfe61bd43565b60200260200101819052505b6040860151511561ada65760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261ad548161c826565b935060ff168151811061ad695761ad6961bd43565b6020026020010181905250856040015182828061ad859061c826565b935060ff168151811061ad9a5761ad9a61bd43565b60200260200101819052505b85606001511561ae105760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261adef8161c826565b935060ff168151811061ae045761ae0461bd43565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561ae2e5761ae2e61b9de565b60405190808252806020026020018201604052801561ae6157816020015b606081526020019060019003908161ae4c5790505b50905060005b8260ff168160ff16101561aeba57838160ff168151811061ae8a5761ae8a61bd43565b6020026020010151828260ff168151811061aea75761aea761bd43565b602090810291909101015260010161ae67565b50979650505050505050565b604080518082019091526000808252602082015281518351101561aeeb575081616d0b565b8151835160208501516000929161af019161ce16565b61af0b919061cd25565b6020840151909150600190821461af2c575082516020840151819020908220145b801561af475783518551869061af4390839061cd25565b9052505b50929392505050565b600080826000015161af74856000015186602001518660000151876020015161b282565b61af7e919061ce16565b90505b8351602085015161af92919061ce16565b811161a7e0578161afa28161d142565b925050826000015161afd985602001518361afbd919061cd25565b865161afc9919061cd25565b838660000151876020015161b282565b61afe3919061ce16565b905061af81565b6060600061aff8848461af50565b61b00390600161ce16565b67ffffffffffffffff81111561b01b5761b01b61b9de565b60405190808252806020026020018201604052801561b04e57816020015b606081526020019060019003908161b0395790505b50905060005b81518110156190e55761b06a61940a868661a7c1565b82828151811061b07c5761b07c61bd43565b602090810291909101015260010161b054565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061b0d8577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061b104576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061b12257662386f26fc10000830492506010015b6305f5e100831061b13a576305f5e100830492506008015b612710831061b14e57612710830492506004015b6064831061b160576064830492506002015b600a8310616d0b5760010192915050565b60008085841161b278576020841161b224576000841561b1bc57600161b19886602061cd25565b61b1a390600861cfd3565b61b1ae90600261d0d1565b61b1b8919061cd25565b1990505b835181168561b1cb898961ce16565b61b1d5919061cd25565b805190935082165b81811461b20f5787841161b1f75787945050505050618d16565b8361b2018161d15c565b94505082845116905061b1dd565b61b219878561ce16565b945050505050618d16565b83832061b231858861cd25565b61b23b908761ce16565b91505b85821061b2765784822080820361b2635761b259868461ce16565b9350505050618d16565b61b26e60018461cd25565b92505061b23e565b505b5092949350505050565b6000838186851161b38d576020851161b33c576000851561b2ce57600161b2aa87602061cd25565b61b2b590600861cfd3565b61b2c090600261d0d1565b61b2ca919061cd25565b1990505b8451811660008761b2df8b8b61ce16565b61b2e9919061cd25565b855190915083165b82811461b32e5781861061b3165761b3098b8b61ce16565b9650505050505050618d16565b8561b3208161d142565b96505083865116905061b2f1565b859650505050505050618d16565b508383206000905b61b34e868961cd25565b821161b38b5785832080820361b36a5783945050505050618d16565b61b37560018561ce16565b935050818061b3839061d142565b92505061b344565b505b61b397878761ce16565b979650505050505050565b6040805180820190915260008082526020820152600061b3d4856000015186602001518660000151876020015161b282565b60208087018051918601919091525190915061b3f0908261cd25565b83528451602086015161b403919061ce16565b810361b412576000855261b444565b8351835161b420919061ce16565b8551869061b42f90839061cd25565b905250835161b43e908261ce16565b60208601525b50909392505050565b6020811061b485578151835261b46460208461ce16565b925061b47160208361ce16565b915061b47e60208261cd25565b905061b44d565b600019811561b4b457600161b49b83602061cd25565b61b4a79061010061d0d1565b61b4b1919061cd25565b90505b9151835183169219169190911790915250565b6060600061b4d584846177e9565b805160208083015160405193945061b4ef9390910161d173565b60405160208183030381529060405291505092915050565b610c9f806200d1cc83390190565b6112a6806200de6b83390190565b610f2a806200f11183390190565b610bf8806201003b83390190565b610cb78062010c3383390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161b59061b595565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161b5906040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561b6475783516001600160a01b031683526020938401939092019160010161b620565b509095945050505050565b60005b8381101561b66d57818101518382015260200161b655565b50506000910152565b6000815180845261b68e81602086016020860161b652565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b79e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561b784577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261b76e84865161b676565b602095860195909450929092019160010161b734565b50919750505060209485019492909201915060010161b6ca565b50929695505050505050565b600081518084526020840193506020830160005b8281101561b7fe5781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161b7be565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b79e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261b874604088018261b676565b905060208201519150868103602088015261b88f818361b7aa565b96505050602093840193919091019060010161b830565b600082825180855260208501945060208160051b8301016020850160005b8381101561b8f657601f1985840301885261b8e083835161b676565b602098890198909350919091019060010161b8c4565b50909695505050505050565b602081526000616f47602083018461b8a6565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b79e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261b996604087018261b7aa565b955050602093840193919091019060010161b93d565b6001600160a01b0384541681526001600160a01b03831660208201526060604082015260006196cb606083018461b676565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561ba305761ba3061b9de565b60405290565b60008067ffffffffffffffff84111561ba515761ba5161b9de565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561ba805761ba8061b9de565b60405283815290508082840185101561ba9857600080fd5b6190e584602083018561b652565b600082601f83011261bab757600080fd5b616f478383516020850161ba36565b60006020828403121561bad857600080fd5b815167ffffffffffffffff81111561baef57600080fd5b616d078482850161baa6565b602081526000616f47602083018461b676565b60006020828403121561bb2057600080fd5b81518015158114616f4757600080fd5b600181811c9082168061bb4457607f821691505b6020821081036196e5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f8211156133fa57806000526020600020601f840160051c8101602085101561bba45750805b601f840160051c820191505b8181101561218a576000815560010161bbb0565b815167ffffffffffffffff81111561bbde5761bbde61b9de565b61bbf28161bbec845461bb30565b8461bb7d565b6020601f82116001811461bc26576000831561bc0e5750848201515b600019600385901b1c1916600184901b17845561218a565b600084815260208120601f198516915b8281101561bc56578785015182556020948501946001909201910161bc36565b508482101561bc745786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60408152600061bcc060408301600481527f5553445400000000000000000000000000000000000000000000000000000000602082015260400190565b8281036020840152616f4781600481527f5553445400000000000000000000000000000000000000000000000000000000602082015260400190565b6001600160a01b0386541681526001600160a01b03851660208201526001600160a01b038416604082015282606082015260a06080820152600061b39760a083018461b676565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081518084526020840193506020830160005b8281101561b7fe57815186526020958601959091019060010161bd86565b60608152600061bdb7606083018661b8a6565b828103602084015261bdc9818661bd72565b9150508215156040830152949350505050565b6001600160a01b0384511681526001600160a01b03831660208201526060604082015260006196cb606083018461b676565b6001600160a01b03841681528260208201526060604082015260006196cb606083018461b676565b828152604060208201526000618d16604083018461b676565b6001600160a01b0381541682526001600160a01b03600182015416602083015260028101546040830152600060038201608060608501526000815461be938161bb30565b806080880152600182166000811461beb2576001811461beec5761bf20565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660a089015260a082151560051b890101935061bf20565b84600052602060002060005b8381101561bf175781548a820160a0015260019091019060200161bef8565b890160a0019450505b50919695505050505050565b6001600160a01b038416815260606020820152600061bf4e606083018561b676565b8281036040840152617e97818561be4f565b60006020828403121561bf7257600080fd5b5051919050565b60006020828403121561bf8b57600080fd5b81516001600160a01b0381168114616f4757600080fd5b60608152600061bfb5606083018661b676565b602083019490945250901515604090910152919050565b6001600160a01b038616815284602082015260a06040820152600061bff460a083018661b676565b6060830194909452509015156080909101529392505050565b6001600160a01b0383168152604060208201526000618d16604083018461b676565b6001600160a01b03861681526001600160a01b038516602082015283604082015260a06060820152600061c06660a083018561b676565b828103608084015261c078818561be4f565b98975050505050505050565b60408152600061c0c160408301600381527f424e420000000000000000000000000000000000000000000000000000000000602082015260400190565b8281036020840152616f4781600381527f424e420000000000000000000000000000000000000000000000000000000000602082015260400190565b6001600160a01b038516815260806020820152600061c11f608083018661b8a6565b828103604084015261c131818661bd72565b915050821515606083015295945050505050565b6001600160a01b0383168152604060208201526000618d16604083018461be4f565b83815260606020820152600061bf4e606083018561b676565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161c1b881601a85016020880161b652565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161c1f581601c84016020880161b652565b01601c01949350505050565b6000835161c21381846020880161b652565b83519083019061c22781836020880161b652565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161c26881601a85016020880161b652565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161c2a581603384016020880161b652565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b6001600160a01b03841681526001600160a01b03831660208201526060604082015260006196cb606083018461b676565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000616f47608083018461b676565b60006020828403121561c36557600080fd5b815167ffffffffffffffff81111561c37c57600080fd5b8201601f8101841361c38d57600080fd5b616d078482516020840161ba36565b6000855161c3ae818460208a0161b652565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161c3e8816001840160208a0161b652565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161c42681600284016020890161b652565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161c46881600284016020880161b652565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061c4b3604083018461b676565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161c52a81601f85016020870161b652565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061c597604083018461b676565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061c5e9604083018461b676565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161c66081601485016020870161b652565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061c6a7604083018561b676565b8281036020840152616f43818561b676565b7f220000000000000000000000000000000000000000000000000000000000000081526000825161c6f181600185016020870161b652565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161c73781846020870161b652565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161c7ea81604b85016020870161b652565b91909101604b0192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600060ff821660ff810361c83c5761c83c61c7f7565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161c8a381602985016020870161b652565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000616f47608083018461b676565b60006020828403121561c90957600080fd5b815167ffffffffffffffff81111561c92057600080fd5b82016060818503121561c93257600080fd5b61c93a61ba0d565b81518060030b811461c94b57600080fd5b8152602082015167ffffffffffffffff81111561c96757600080fd5b61c9738682850161baa6565b602083015250604082015167ffffffffffffffff81111561c99357600080fd5b61c99f8682850161baa6565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161ca0b81602185016020870161b652565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161cbf781602185016020880161b652565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161cc3481602e84016020880161b652565b01602e01949350505050565b6000825161cc5281846020870161b652565b9190910192915050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161c8a381602985016020870161b652565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161cd1881602285016020870161b652565b9190910160220192915050565b81810381811115616d0b57616d0b61c7f7565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161cd7081600e85016020870161b652565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b80820180821115616d0b57616d0b61c7f7565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161ce6181601885016020880161b652565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161ce9e81601c84016020880161b652565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161cfa481846020870161b652565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b8082028115828204841417616d0b57616d0b61c7f7565b6001815b600184111561d0255780850481111561d0095761d00961c7f7565b600184161561d01757908102905b60019390931c92800261cfee565b935093915050565b60008261d03c57506001616d0b565b8161d04957506000616d0b565b816001811461d05f576002811461d0695761d085565b6001915050616d0b565b60ff84111561d07a5761d07a61c7f7565b50506001821b616d0b565b5060208310610133831016604e8410600b841016171561d0a8575081810a616d0b565b61d0b5600019848461cfea565b806000190482111561d0c95761d0c961c7f7565b029392505050565b6000616f47838361d02d565b818103600083128015838313168383128216171561a7e05761a7e061c7f7565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161d13581601c85016020870161b652565b91909101601c0192915050565b6000600019820361d1555761d15561c7f7565b5060010190565b60008161d16b5761d16b61c7f7565b506000190190565b6000835161d18581846020880161b652565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161d1bf81600184016020880161b652565b0160010194935050505056fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220837c7d9916de10b20cdb18567d8c7679613426bbd7b0b72548d8000a412f307b64736f6c634300081a0033608060405234801561001057600080fd5b506040516112a63803806112a683398101604081905261002f91610110565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007891906101e2565b50600461008582826101e2565b5050506001600160a01b03821615806100a557506001600160a01b038116155b156100c35760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b031991821617909155600780549290931691161790556102a0565b80516001600160a01b038116811461010b57600080fd5b919050565b6000806040838503121561012357600080fd5b61012c836100f4565b915061013a602084016100f4565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061016d57607f821691505b60208210810361018d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101dd57806000526020600020601f840160051c810160208510156101ba5750805b601f840160051c820191505b818110156101da57600081556001016101c6565b50505b505050565b81516001600160401b038111156101fb576101fb610143565b61020f816102098454610159565b84610193565b6020601f821160018114610243576000831561022b5750848201515b600019600385901b1c1916600184901b1784556101da565b600084815260208120601f198516915b828110156102735787850151825560209485019460019092019101610253565b50848210156102915786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610ff7806102af6000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806342966c68116100b257806379cc679011610081578063a9059cbb11610066578063a9059cbb1461028e578063bff9662a146102a1578063dd62ed3e146102c157600080fd5b806379cc67901461027357806395d89b411461028657600080fd5b806342966c68146102025780635b1125911461021557806370a0823114610235578063779e3b631461026b57600080fd5b80631e458bee116100ee5780631e458bee1461018857806323b872dd1461019b578063313ce567146101ae578063328a01d0146101bd57600080fd5b806306fdde0314610120578063095ea7b31461013e57806315d57fd41461016157806318160ddd14610176575b600080fd5b610128610307565b6040516101359190610d97565b60405180910390f35b61015161014c366004610e2c565b610399565b6040519015158152602001610135565b61017461016f366004610e56565b6103b3565b005b6002545b604051908152602001610135565b610174610196366004610e89565b61057e565b6101516101a9366004610ebc565b610631565b60405160128152602001610135565b6007546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610135565b610174610210366004610ef9565b610655565b6006546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a610243366004610f12565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b610174610662565b610174610281366004610e2c565b610786565b610128610837565b61015161029c366004610e2c565b610846565b6005546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a6102cf366004610e56565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60606003805461031690610f34565b80601f016020809104026020016040519081016040528092919081815260200182805461034290610f34565b801561038f5780601f106103645761010080835404028352916020019161038f565b820191906000526020600020905b81548152906001019060200180831161037257829003601f168201915b5050505050905090565b6000336103a7818585610854565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103f3575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610431576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82161580610468575073ffffffffffffffffffffffffffffffffffffffff8116155b1561049f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105d1576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6105db8383610866565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161062491815260200190565b60405180910390a3505050565b60003361063f8582856108c6565b61064a858585610995565b506001949350505050565b61065f3382610a40565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106b5576040517fe700765e000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b60065473ffffffffffffffffffffffffffffffffffffffff16610704576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107d9576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6107e38282610a9c565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161082b91815260200190565b60405180910390a25050565b60606004805461031690610f34565b6000336103a7818585610995565b6108618383836001610ab1565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108b6576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c260008383610bf9565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461098f5781811015610980576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610428565b61098f84848484036000610ab1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109e5576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8216610a35576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b610861838383610bf9565b73ffffffffffffffffffffffffffffffffffffffff8216610a90576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c282600083610bf9565b610aa78233836108c6565b6108c28282610a40565b73ffffffffffffffffffffffffffffffffffffffff8416610b01576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8316610b51576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600160209081526040808320938716835292905220829055801561098f578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610beb91815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c31578060026000828254610c269190610f87565b90915550610ce39050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610cb7576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610428565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610d0c57600280548290039055610d38565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161062491815260200190565b602081526000825180602084015260005b81811015610dc55760208186018101516040868401015201610da8565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610e2757600080fd5b919050565b60008060408385031215610e3f57600080fd5b610e4883610e03565b946020939093013593505050565b60008060408385031215610e6957600080fd5b610e7283610e03565b9150610e8060208401610e03565b90509250929050565b600080600060608486031215610e9e57600080fd5b610ea784610e03565b95602085013595506040909401359392505050565b600080600060608486031215610ed157600080fd5b610eda84610e03565b9250610ee860208501610e03565b929592945050506040919091013590565b600060208284031215610f0b57600080fd5b5035919050565b600060208284031215610f2457600080fd5b610f2d82610e03565b9392505050565b600181811c90821680610f4857607f821691505b602082108103610f81577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b808201808211156103ad577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220fec2a89bc2f37ba9f63e9d7e47fa9a0d2b02402309f8905f94cc44fb9e5fbba064736f6c634300081a00336080604052348015600f57600080fd5b506001600055610f06806100246000396000f3fe60806040526004361061006e5760003560e01c8063c51316911161004b578063c5131691146100d5578063c9028a36146100f5578063e04d4f9714610115578063f05b6abf1461012857005b8063357fc5a214610077578063676cc054146100975780636ed70169146100c057005b3661007557005b005b34801561008357600080fd5b50610075610092366004610709565b610148565b6100aa6100a5366004610745565b6101de565b6040516100b79190610840565b60405180910390f35b3480156100cc57600080fd5b5061007561023f565b3480156100e157600080fd5b506100756100f0366004610709565b610274565b34801561010157600080fd5b50610075610110366004610853565b61034f565b6100756101233660046109b3565b61038b565b34801561013457600080fd5b50610075610143366004610a9f565b6103cf565b610150610404565b61017273ffffffffffffffffffffffffffffffffffffffff8316338386610447565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101d96001600055565b505050565b60607fd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a35250161020e6020860186610b89565b848460405161021f93929190610bed565b60405180910390a1506040805160208101909152600081525b9392505050565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61027c610404565b6000610289600285610c26565b9050806000036102c5576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102e773ffffffffffffffffffffffffffffffffffffffff8416338484610447565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101d96001600055565b7f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b3382604051610380929190610c61565b60405180910390a150565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa33348585856040516103c2959493929190610d53565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146338484846040516103c29493929190610ddd565b600260005403610440576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526104dc9085906104e2565b50505050565b600061050473ffffffffffffffffffffffffffffffffffffffff84168361057d565b905080516000141580156105295750808060200190518101906105279190610e97565b155b156101d9576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b606061023883836000846000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105b09190610eb4565b60006040518083038185875af1925050503d80600081146105ed576040519150601f19603f3d011682016040523d82523d6000602084013e6105f2565b606091505b509150915061060286838361060c565b9695505050505050565b6060826106215761061c8261069b565b610238565b8151158015610645575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610694576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610574565b5080610238565b8051156106ab5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff8116811461070457600080fd5b919050565b60008060006060848603121561071e57600080fd5b8335925061072e602085016106e0565b915061073c604085016106e0565b90509250925092565b6000806000838503604081121561075b57600080fd5b602081121561076957600080fd5b50839250602084013567ffffffffffffffff81111561078757600080fd5b8401601f8101861361079857600080fd5b803567ffffffffffffffff8111156107af57600080fd5b8660208284010111156107c157600080fd5b939660209190910195509293505050565b60005b838110156107ed5781810151838201526020016107d5565b50506000910152565b6000815180845261080e8160208601602086016107d2565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061023860208301846107f6565b60006020828403121561086557600080fd5b813567ffffffffffffffff81111561087c57600080fd5b82016080818503121561023857600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156109045761090461088e565b604052919050565b600082601f83011261091d57600080fd5b813567ffffffffffffffff8111156109375761093761088e565b61096860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016108bd565b81815284602083860101111561097d57600080fd5b816020850160208301376000918101602001919091529392505050565b80151581146106dd57600080fd5b80356107048161099a565b6000806000606084860312156109c857600080fd5b833567ffffffffffffffff8111156109df57600080fd5b6109eb8682870161090c565b935050602084013591506040840135610a038161099a565b809150509250925092565b600067ffffffffffffffff821115610a2857610a2861088e565b5060051b60200190565b600082601f830112610a4357600080fd5b8135610a56610a5182610a0e565b6108bd565b8082825260208201915060208360051b860101925085831115610a7857600080fd5b602085015b83811015610a95578035835260209283019201610a7d565b5095945050505050565b600080600060608486031215610ab457600080fd5b833567ffffffffffffffff811115610acb57600080fd5b8401601f81018613610adc57600080fd5b8035610aea610a5182610a0e565b8082825260208201915060208360051b850101925088831115610b0c57600080fd5b602084015b83811015610b4e57803567ffffffffffffffff811115610b3057600080fd5b610b3f8b60208389010161090c565b84525060209283019201610b11565b509550505050602084013567ffffffffffffffff811115610b6e57600080fd5b610b7a86828701610a32565b92505061073c604085016109a8565b600060208284031215610b9b57600080fd5b610238826106e0565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff84168152604060208201526000610c1d604083018486610ba4565b95945050505050565b600082610c5c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610c9f836106e0565b16604082015273ffffffffffffffffffffffffffffffffffffffff610cc6602084016106e0565b166060820152600080604084013590508060808401525060608301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610d1257600080fd5b830160208101903567ffffffffffffffff811115610d2f57600080fd5b803603821315610d3e57600080fd5b608060a085015261060260c085018284610ba4565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610d8860a08301866107f6565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610dd3578151865260209586019590910190600101610db5565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610e70577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610e5b8583516107f6565b94506020938401939190910190600101610e21565b505050508281036040840152610e868186610da1565b915050610c1d606083018415159052565b600060208284031215610ea957600080fd5b81516102388161099a565b60008251610ec68184602087016107d2565b919091019291505056fea2646970667358221220cecc8e08efbffa03aeef8fe05cb9598999ded8d265bc5994494dac5aca93443464736f6c634300081a003360806040526002805460ff1916601217905534801561001d57600080fd5b50604051610bf8380380610bf883398101604081905261003c91610116565b60006100488382610208565b5060016100558282610208565b5050506102c6565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261008457600080fd5b81516001600160401b0381111561009d5761009d61005d565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100cb576100cb61005d565b6040528181528382016020018510156100e357600080fd5b60005b82811015610102576020818601810151838301820152016100e6565b506000918101602001919091529392505050565b6000806040838503121561012957600080fd5b82516001600160401b0381111561013f57600080fd5b61014b85828601610073565b602085015190935090506001600160401b0381111561016957600080fd5b61017585828601610073565b9150509250929050565b600181811c9082168061019357607f821691505b6020821081036101b357634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561020357806000526020600020601f840160051c810160208510156101e05750805b601f840160051c820191505b8181101561020057600081556001016101ec565b50505b505050565b81516001600160401b038111156102215761022161005d565b6102358161022f845461017f565b846101b9565b6020601f82116001811461026957600083156102515750848201515b600019600385901b1c1916600184901b178455610200565b600084815260208120601f198516915b828110156102995787850151825560209485019460019092019101610279565b50848210156102b75786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610923806102d56000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610182578063a9059cbb1461018a578063dd62ed3e1461019d57600080fd5b806340c10f191461014f57806370a082311461016257600080fd5b806318160ddd116100a757806318160ddd146100f657806323b872dd1461010d578063313ce5671461013057600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101c8565b6040516100d891906106f4565b60405180910390f35b6100f46100ef366004610789565b610256565b005b6100ff60035481565b6040519081526020016100d8565b61012061011b3660046107b3565b6102c3565b60405190151581526020016100d8565b60025461013d9060ff1681565b60405160ff90911681526020016100d8565b6100f461015d366004610789565b61051d565b6100ff6101703660046107f0565b60046020526000908152604090205481565b6100cb6105b9565b610120610198366004610789565b6105c6565b6100ff6101ab366004610812565b600560209081526000928352604080842090915290825290205481565b600080546101d590610845565b80601f016020809104026020016040519081016040528092919081815260200182805461020190610845565b801561024e5780601f106102235761010080835404028352916020019161024e565b820191906000526020600020905b81548152906001019060200180831161023157829003601f168201915b505050505081565b33600081815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff871680855290835292819020859055518481529192917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260046020526040812054821115610357576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e636500000000000000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841660009081526005602090815260408083203384529091529020548211156103f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e73756666696369656e7420616c6c6f77616e636500000000000000000000604482015260640161034e565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260046020526040812080548492906104269084906108c7565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600090815260046020526040812080548492906104609084906108da565b909155505073ffffffffffffffffffffffffffffffffffffffff84166000908152600560209081526040808320338452909152812080548492906104a59084906108c7565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161050b91815260200190565b60405180910390a35060019392505050565b806003600082825461052f91906108da565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260046020526040812080548392906105699084906108da565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016102b7565b600180546101d590610845565b3360009081526004602052604081205482111561063f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e6365000000000000000000000000604482015260640161034e565b336000908152600460205260408120805484929061065e9084906108c7565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600090815260046020526040812080548492906106989084906108da565b909155505060405182815273ffffffffffffffffffffffffffffffffffffffff84169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35060015b92915050565b602081526000825180602084015260005b818110156107225760208186018101516040868401015201610705565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461078457600080fd5b919050565b6000806040838503121561079c57600080fd5b6107a583610760565b946020939093013593505050565b6000806000606084860312156107c857600080fd5b6107d184610760565b92506107df60208501610760565b929592945050506040919091013590565b60006020828403121561080257600080fd5b61080b82610760565b9392505050565b6000806040838503121561082557600080fd5b61082e83610760565b915061083c60208401610760565b90509250929050565b600181811c9082168061085957607f821691505b602082108103610892577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156106ee576106ee610898565b808201808211156106ee576106ee61089856fea264697066735822122041013599089452b2c9cdcab1a4eb2194e2bcab74511cb3cee684a048cea2529f64736f6c634300081a0033608060405234801561001057600080fd5b50604051610cb7380380610cb783398101604081905261002f91610111565b81818181600361003f8382610203565b50600461004c8282610203565b505050505050506102c1565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007f57600080fd5b81516001600160401b0381111561009857610098610058565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c6576100c6610058565b6040528181528382016020018510156100de57600080fd5b60005b828110156100fd576020818601810151838301820152016100e1565b506000918101602001919091529392505050565b6000806040838503121561012457600080fd5b82516001600160401b0381111561013a57600080fd5b6101468582860161006e565b602085015190935090506001600160401b0381111561016457600080fd5b6101708582860161006e565b9150509250929050565b600181811c9082168061018e57607f821691505b6020821081036101ae57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fe57806000526020600020601f840160051c810160208510156101db5750805b601f840160051c820191505b818110156101fb57600081556001016101e7565b50505b505050565b81516001600160401b0381111561021c5761021c610058565b6102308161022a845461017a565b846101b4565b6020601f821160018114610264576000831561024c5750848201515b600019600385901b1c1916600184901b1784556101fb565b600084815260208120601f198516915b828110156102945787850151825560209485019460019092019101610274565b50848210156102b25786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109e7806102d06000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107da565b60405180910390f35b6100f46100ef36600461086f565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f4610124366004610899565b610297565b604051601281526020016100d8565b61014b61014636600461086f565b6102bb565b005b61010861015b3660046108d6565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c9565b6100f461019936600461086f565b6102d8565b6101086101ac3660046108f1565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610924565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610924565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b600080821161028457600080fd5b61028e83836102f0565b90505b92915050565b6000336102a58582856102fe565b6102b08585856103d2565b506001949350505050565b6102c58282610482565b5050565b6060600480546101f390610924565b6000336102e68185856103d2565b5060019392505050565b6000336102e68185856104de565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103cc57818110156103bd576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103cc848484840360006104e7565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610422576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103b4565b73ffffffffffffffffffffffffffffffffffffffff8216610472576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103b4565b61047d83838361062f565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166104d2576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103b4565b6102c56000838361062f565b61047d83838360015b73ffffffffffffffffffffffffffffffffffffffff8416610537576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103b4565b73ffffffffffffffffffffffffffffffffffffffff8316610587576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103b4565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103cc578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161062191815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661066757806002600082825461065c9190610977565b909155506107199050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106ed576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103b4565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff82166107425760028054829003905561076e565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107cd91815260200190565b60405180910390a3505050565b602081526000825180602084015260005b8181101561080857602081860181015160408684010152016107eb565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461086a57600080fd5b919050565b6000806040838503121561088257600080fd5b61088b83610846565b946020939093013593505050565b6000806000606084860312156108ae57600080fd5b6108b784610846565b92506108c560208501610846565b929592945050506040919091013590565b6000602082840312156108e857600080fd5b61028e82610846565b6000806040838503121561090457600080fd5b61090d83610846565b915061091b60208401610846565b90509250929050565b600181811c9082168061093857607f821691505b602082108103610971577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b80820180821115610291577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea264697066735822122073db0e30c05200fc444905446d879bf6bfcd0ec4eb5c23d25c82318bbdc3416b64736f6c634300081a0033a2646970667358221220d7bc68cf82da86f3fd98efda2eaf0e8e3720b1d5897d25fa036fd3707f97550364736f6c634300081a0033",
}

// GatewayEVMTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMTestMetaData.ABI instead.
var GatewayEVMTestABI = GatewayEVMTestMetaData.ABI

// GatewayEVMTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMTestMetaData.Bin instead.
var GatewayEVMTestBin = GatewayEVMTestMetaData.Bin

// DeployGatewayEVMTest deploys a new Ethereum contract, binding an instance of GatewayEVMTest to it.
func DeployGatewayEVMTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMTest, error) {
	parsed, err := GatewayEVMTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMTest{GatewayEVMTestCaller: GatewayEVMTestCaller{contract: contract}, GatewayEVMTestTransactor: GatewayEVMTestTransactor{contract: contract}, GatewayEVMTestFilterer: GatewayEVMTestFilterer{contract: contract}}, nil
}

// GatewayEVMTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMTest struct {
	GatewayEVMTestCaller     // Read-only binding to the contract
	GatewayEVMTestTransactor // Write-only binding to the contract
	GatewayEVMTestFilterer   // Log filterer for contract events
}

// GatewayEVMTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMTestSession struct {
	Contract     *GatewayEVMTest   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayEVMTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMTestCallerSession struct {
	Contract *GatewayEVMTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GatewayEVMTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMTestTransactorSession struct {
	Contract     *GatewayEVMTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GatewayEVMTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMTestRaw struct {
	Contract *GatewayEVMTest // Generic contract binding to access the raw methods on
}

// GatewayEVMTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMTestCallerRaw struct {
	Contract *GatewayEVMTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMTestTransactorRaw struct {
	Contract *GatewayEVMTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMTest creates a new instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMTest, error) {
	contract, err := bindGatewayEVMTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTest{GatewayEVMTestCaller: GatewayEVMTestCaller{contract: contract}, GatewayEVMTestTransactor: GatewayEVMTestTransactor{contract: contract}, GatewayEVMTestFilterer: GatewayEVMTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMTestCaller creates a new read-only instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMTestCaller, error) {
	contract, err := bindGatewayEVMTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestCaller{contract: contract}, nil
}

// NewGatewayEVMTestTransactor creates a new write-only instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMTestTransactor, error) {
	contract, err := bindGatewayEVMTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestTransactor{contract: contract}, nil
}

// NewGatewayEVMTestFilterer creates a new log filterer instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMTestFilterer, error) {
	contract, err := bindGatewayEVMTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestFilterer{contract: contract}, nil
}

// bindGatewayEVMTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMTest *GatewayEVMTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMTest.Contract.GatewayEVMTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMTest *GatewayEVMTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.GatewayEVMTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMTest *GatewayEVMTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.GatewayEVMTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMTest *GatewayEVMTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMTest *GatewayEVMTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMTest *GatewayEVMTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.contract.Transact(opts, method, params...)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) ASSETHANDLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "ASSET_HANDLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.ASSETHANDLERROLE(&_GatewayEVMTest.CallOpts)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.ASSETHANDLERROLE(&_GatewayEVMTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.DEFAULTADMINROLE(&_GatewayEVMTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.DEFAULTADMINROLE(&_GatewayEVMTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestSession) ISTEST() (bool, error) {
	return _GatewayEVMTest.Contract.ISTEST(&_GatewayEVMTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ISTEST() (bool, error) {
	return _GatewayEVMTest.Contract.ISTEST(&_GatewayEVMTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.PAUSERROLE(&_GatewayEVMTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.PAUSERROLE(&_GatewayEVMTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.TSSROLE(&_GatewayEVMTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.TSSROLE(&_GatewayEVMTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) WITHDRAWERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.WITHDRAWERROLE(&_GatewayEVMTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.WITHDRAWERROLE(&_GatewayEVMTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.ExcludeArtifacts(&_GatewayEVMTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.ExcludeArtifacts(&_GatewayEVMTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeContracts(&_GatewayEVMTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeContracts(&_GatewayEVMTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.ExcludeSelectors(&_GatewayEVMTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.ExcludeSelectors(&_GatewayEVMTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeSenders(&_GatewayEVMTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeSenders(&_GatewayEVMTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestSession) Failed() (bool, error) {
	return _GatewayEVMTest.Contract.Failed(&_GatewayEVMTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) Failed() (bool, error) {
	return _GatewayEVMTest.Contract.Failed(&_GatewayEVMTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMTest.Contract.TargetArtifactSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMTest.Contract.TargetArtifactSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.TargetArtifacts(&_GatewayEVMTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.TargetArtifacts(&_GatewayEVMTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetContracts(&_GatewayEVMTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetContracts(&_GatewayEVMTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMTest.Contract.TargetInterfaces(&_GatewayEVMTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMTest.Contract.TargetInterfaces(&_GatewayEVMTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.TargetSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.TargetSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetSenders(&_GatewayEVMTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetSenders(&_GatewayEVMTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.SetUp(&_GatewayEVMTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.SetUp(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x04b974f9.
//
// Solidity: function testExecuteFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteFailsIfDestinationIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteFailsIfDestinationIsZeroAddress")
}

// TestExecuteFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x04b974f9.
//
// Solidity: function testExecuteFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x04b974f9.
//
// Solidity: function testExecuteFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevert is a paid mutator transaction binding the contract method 0xfa18c09b.
//
// Solidity: function testExecuteRevert() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteRevert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteRevert")
}

// TestExecuteRevert is a paid mutator transaction binding the contract method 0xfa18c09b.
//
// Solidity: function testExecuteRevert() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteRevert() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevert(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevert is a paid mutator transaction binding the contract method 0xfa18c09b.
//
// Solidity: function testExecuteRevert() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteRevert() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevert(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51010e49.
//
// Solidity: function testExecuteRevertFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteRevertFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteRevertFailsIfReceiverIsZeroAddress")
}

// TestExecuteRevertFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51010e49.
//
// Solidity: function testExecuteRevertFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteRevertFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfReceiverIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51010e49.
//
// Solidity: function testExecuteRevertFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteRevertFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfReceiverIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0xcebad2a6.
//
// Solidity: function testExecuteRevertFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteRevertFailsIfSenderIsNotTSS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteRevertFailsIfSenderIsNotTSS")
}

// TestExecuteRevertFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0xcebad2a6.
//
// Solidity: function testExecuteRevertFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteRevertFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0xcebad2a6.
//
// Solidity: function testExecuteRevertFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteRevertFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xe6afc790.
//
// Solidity: function testExecuteWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteWithERC20FailsIfNotCustodyOrConnector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteWithERC20FailsIfNotCustodyOrConnector")
}

// TestExecuteWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xe6afc790.
//
// Solidity: function testExecuteWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xe6afc790.
//
// Solidity: function testExecuteWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithERC20Token is a paid mutator transaction binding the contract method 0x520d989d.
//
// Solidity: function testExecuteWithERC20Token() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteWithERC20Token(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteWithERC20Token")
}

// TestExecuteWithERC20Token is a paid mutator transaction binding the contract method 0x520d989d.
//
// Solidity: function testExecuteWithERC20Token() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteWithERC20Token() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithERC20Token(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithERC20Token is a paid mutator transaction binding the contract method 0x520d989d.
//
// Solidity: function testExecuteWithERC20Token() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteWithERC20Token() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithERC20Token(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x6bdd212b.
//
// Solidity: function testExecuteWithMsgContextFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteWithMsgContextFailsIfDestinationIsZeroAddress")
}

// TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x6bdd212b.
//
// Solidity: function testExecuteWithMsgContextFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x6bdd212b.
//
// Solidity: function testExecuteWithMsgContextFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithNonReturnApprovalToken is a paid mutator transaction binding the contract method 0x3801e58f.
//
// Solidity: function testExecuteWithNonReturnApprovalToken() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteWithNonReturnApprovalToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteWithNonReturnApprovalToken")
}

// TestExecuteWithNonReturnApprovalToken is a paid mutator transaction binding the contract method 0x3801e58f.
//
// Solidity: function testExecuteWithNonReturnApprovalToken() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteWithNonReturnApprovalToken() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithNonReturnApprovalToken(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithNonReturnApprovalToken is a paid mutator transaction binding the contract method 0x3801e58f.
//
// Solidity: function testExecuteWithNonReturnApprovalToken() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteWithNonReturnApprovalToken() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithNonReturnApprovalToken(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithTokenRevertingOnZeroApproval is a paid mutator transaction binding the contract method 0xe5418219.
//
// Solidity: function testExecuteWithTokenRevertingOnZeroApproval() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteWithTokenRevertingOnZeroApproval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteWithTokenRevertingOnZeroApproval")
}

// TestExecuteWithTokenRevertingOnZeroApproval is a paid mutator transaction binding the contract method 0xe5418219.
//
// Solidity: function testExecuteWithTokenRevertingOnZeroApproval() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteWithTokenRevertingOnZeroApproval() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithTokenRevertingOnZeroApproval(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithTokenRevertingOnZeroApproval is a paid mutator transaction binding the contract method 0xe5418219.
//
// Solidity: function testExecuteWithTokenRevertingOnZeroApproval() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteWithTokenRevertingOnZeroApproval() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithTokenRevertingOnZeroApproval(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParams is a paid mutator transaction binding the contract method 0x44671b94.
//
// Solidity: function testForwardCallToReceiveNoParams() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNoParams(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNoParams")
}

// TestForwardCallToReceiveNoParams is a paid mutator transaction binding the contract method 0x44671b94.
//
// Solidity: function testForwardCallToReceiveNoParams() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNoParams() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParams(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParams is a paid mutator transaction binding the contract method 0x44671b94.
//
// Solidity: function testForwardCallToReceiveNoParams() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNoParams() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParams(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParamsTogglePause is a paid mutator transaction binding the contract method 0xdd51e82f.
//
// Solidity: function testForwardCallToReceiveNoParamsTogglePause() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNoParamsTogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNoParamsTogglePause")
}

// TestForwardCallToReceiveNoParamsTogglePause is a paid mutator transaction binding the contract method 0xdd51e82f.
//
// Solidity: function testForwardCallToReceiveNoParamsTogglePause() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNoParamsTogglePause() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParamsTogglePause(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParamsTogglePause is a paid mutator transaction binding the contract method 0xdd51e82f.
//
// Solidity: function testForwardCallToReceiveNoParamsTogglePause() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNoParamsTogglePause() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParamsTogglePause(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayable is a paid mutator transaction binding the contract method 0xf68bd1c0.
//
// Solidity: function testForwardCallToReceiveNonPayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNonPayable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNonPayable")
}

// TestForwardCallToReceiveNonPayable is a paid mutator transaction binding the contract method 0xf68bd1c0.
//
// Solidity: function testForwardCallToReceiveNonPayable() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNonPayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayable(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayable is a paid mutator transaction binding the contract method 0xf68bd1c0.
//
// Solidity: function testForwardCallToReceiveNonPayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNonPayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayable(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x7d7f772a.
//
// Solidity: function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS")
}

// TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x7d7f772a.
//
// Solidity: function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x7d7f772a.
//
// Solidity: function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x43fd8c7d.
//
// Solidity: function testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS")
}

// TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x43fd8c7d.
//
// Solidity: function testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x43fd8c7d.
//
// Solidity: function testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnCallFails is a paid mutator transaction binding the contract method 0xd38b66cd.
//
// Solidity: function testForwardCallToReceiveOnCallFails() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveOnCallFails(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveOnCallFails")
}

// TestForwardCallToReceiveOnCallFails is a paid mutator transaction binding the contract method 0xd38b66cd.
//
// Solidity: function testForwardCallToReceiveOnCallFails() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveOnCallFails() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnCallFails(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnCallFails is a paid mutator transaction binding the contract method 0xd38b66cd.
//
// Solidity: function testForwardCallToReceiveOnCallFails() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveOnCallFails() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnCallFails(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnCallUsingAuthCall is a paid mutator transaction binding the contract method 0xa397ffd2.
//
// Solidity: function testForwardCallToReceiveOnCallUsingAuthCall() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveOnCallUsingAuthCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveOnCallUsingAuthCall")
}

// TestForwardCallToReceiveOnCallUsingAuthCall is a paid mutator transaction binding the contract method 0xa397ffd2.
//
// Solidity: function testForwardCallToReceiveOnCallUsingAuthCall() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveOnCallUsingAuthCall() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnCallUsingAuthCall(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnCallUsingAuthCall is a paid mutator transaction binding the contract method 0xa397ffd2.
//
// Solidity: function testForwardCallToReceiveOnCallUsingAuthCall() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveOnCallUsingAuthCall() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnCallUsingAuthCall(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnRevertFails is a paid mutator transaction binding the contract method 0x130daf59.
//
// Solidity: function testForwardCallToReceiveOnRevertFails() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveOnRevertFails(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveOnRevertFails")
}

// TestForwardCallToReceiveOnRevertFails is a paid mutator transaction binding the contract method 0x130daf59.
//
// Solidity: function testForwardCallToReceiveOnRevertFails() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveOnRevertFails() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnRevertFails(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnRevertFails is a paid mutator transaction binding the contract method 0x130daf59.
//
// Solidity: function testForwardCallToReceiveOnRevertFails() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveOnRevertFails() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnRevertFails(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0xfe7bdbb2.
//
// Solidity: function testForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceivePayable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceivePayable")
}

// TestForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0xfe7bdbb2.
//
// Solidity: function testForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceivePayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceivePayable(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0xfe7bdbb2.
//
// Solidity: function testForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceivePayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceivePayable(&_GatewayEVMTest.TransactOpts)
}

// TestRevertWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xb124dbf5.
//
// Solidity: function testRevertWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestRevertWithERC20FailsIfNotCustodyOrConnector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testRevertWithERC20FailsIfNotCustodyOrConnector")
}

// TestRevertWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xb124dbf5.
//
// Solidity: function testRevertWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestRevertWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestRevertWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestRevertWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xb124dbf5.
//
// Solidity: function testRevertWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestRevertWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestRevertWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xccf20616.
//
// Solidity: function testSetConnectorFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetConnectorFailsIfSenderIsNotAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetConnectorFailsIfSenderIsNotAdmin")
}

// TestSetConnectorFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xccf20616.
//
// Solidity: function testSetConnectorFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetConnectorFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xccf20616.
//
// Solidity: function testSetConnectorFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetConnectorFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSet is a paid mutator transaction binding the contract method 0x7ebf744f.
//
// Solidity: function testSetConnectorFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetConnectorFailsIfSet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetConnectorFailsIfSet")
}

// TestSetConnectorFailsIfSet is a paid mutator transaction binding the contract method 0x7ebf744f.
//
// Solidity: function testSetConnectorFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetConnectorFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSet is a paid mutator transaction binding the contract method 0x7ebf744f.
//
// Solidity: function testSetConnectorFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetConnectorFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfZero is a paid mutator transaction binding the contract method 0xce5871e1.
//
// Solidity: function testSetConnectorFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetConnectorFailsIfZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetConnectorFailsIfZero")
}

// TestSetConnectorFailsIfZero is a paid mutator transaction binding the contract method 0xce5871e1.
//
// Solidity: function testSetConnectorFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetConnectorFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfZero is a paid mutator transaction binding the contract method 0xce5871e1.
//
// Solidity: function testSetConnectorFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetConnectorFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xa56f7a4b.
//
// Solidity: function testSetCustodyFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetCustodyFailsIfSenderIsNotAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetCustodyFailsIfSenderIsNotAdmin")
}

// TestSetCustodyFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xa56f7a4b.
//
// Solidity: function testSetCustodyFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetCustodyFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xa56f7a4b.
//
// Solidity: function testSetCustodyFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetCustodyFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSet is a paid mutator transaction binding the contract method 0x1855c337.
//
// Solidity: function testSetCustodyFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetCustodyFailsIfSet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetCustodyFailsIfSet")
}

// TestSetCustodyFailsIfSet is a paid mutator transaction binding the contract method 0x1855c337.
//
// Solidity: function testSetCustodyFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetCustodyFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSet is a paid mutator transaction binding the contract method 0x1855c337.
//
// Solidity: function testSetCustodyFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetCustodyFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfZero is a paid mutator transaction binding the contract method 0x1226c655.
//
// Solidity: function testSetCustodyFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetCustodyFailsIfZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetCustodyFailsIfZero")
}

// TestSetCustodyFailsIfZero is a paid mutator transaction binding the contract method 0x1226c655.
//
// Solidity: function testSetCustodyFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetCustodyFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfZero is a paid mutator transaction binding the contract method 0x1226c655.
//
// Solidity: function testSetCustodyFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetCustodyFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestTSSUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testTSSUpgrade")
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgrade(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgrade(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testTSSUpgradeFailsIfSenderIsNotTSSUpdater")
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestTSSUpgradeFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testTSSUpgradeFailsIfZeroAddress")
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestUpgradeAndForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0x7a380ebf.
//
// Solidity: function testUpgradeAndForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestUpgradeAndForwardCallToReceivePayable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testUpgradeAndForwardCallToReceivePayable")
}

// TestUpgradeAndForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0x7a380ebf.
//
// Solidity: function testUpgradeAndForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestUpgradeAndForwardCallToReceivePayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestUpgradeAndForwardCallToReceivePayable(&_GatewayEVMTest.TransactOpts)
}

// TestUpgradeAndForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0x7a380ebf.
//
// Solidity: function testUpgradeAndForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestUpgradeAndForwardCallToReceivePayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestUpgradeAndForwardCallToReceivePayable(&_GatewayEVMTest.TransactOpts)
}

// GatewayEVMTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayEVMTest contract.
type GatewayEVMTestCalledIterator struct {
	Event *GatewayEVMTestCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestCalled represents a Called event raised by the GatewayEVMTest contract.
type GatewayEVMTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestCalledIterator{contract: _GatewayEVMTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestCalled)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Called", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCalled is a log parse operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseCalled(log types.Log) (*GatewayEVMTestCalled, error) {
	event := new(GatewayEVMTestCalled)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayEVMTest contract.
type GatewayEVMTestDepositedIterator struct {
	Event *GatewayEVMTestDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestDeposited represents a Deposited event raised by the GatewayEVMTest contract.
type GatewayEVMTestDeposited struct {
	Recipient []byte
	Asset     common.Address
	Amount    *big.Int
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterDeposited(opts *bind.FilterOpts, asset []common.Address) (*GatewayEVMTestDepositedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestDepositedIterator{contract: _GatewayEVMTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestDeposited, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestDeposited)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseDeposited(log types.Log) (*GatewayEVMTestDeposited, error) {
	event := new(GatewayEVMTestDeposited)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestDeposited0Iterator is returned from FilterDeposited0 and is used to iterate over the raw logs and unpacked data for Deposited0 events raised by the GatewayEVMTest contract.
type GatewayEVMTestDeposited0Iterator struct {
	Event *GatewayEVMTestDeposited0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestDeposited0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestDeposited0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestDeposited0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestDeposited0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestDeposited0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestDeposited0 represents a Deposited0 event raised by the GatewayEVMTest contract.
type GatewayEVMTestDeposited0 struct {
	Sender        common.Address
	Receiver      common.Address
	Amount        *big.Int
	Asset         common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposited0 is a free log retrieval operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterDeposited0(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMTestDeposited0Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Deposited0", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestDeposited0Iterator{contract: _GatewayEVMTest.contract, event: "Deposited0", logs: logs, sub: sub}, nil
}

// WatchDeposited0 is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchDeposited0(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestDeposited0, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Deposited0", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestDeposited0)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited0 is a log parse operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseDeposited0(log types.Log) (*GatewayEVMTestDeposited0, error) {
	event := new(GatewayEVMTestDeposited0)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestDepositedAndCalledIterator is returned from FilterDepositedAndCalled and is used to iterate over the raw logs and unpacked data for DepositedAndCalled events raised by the GatewayEVMTest contract.
type GatewayEVMTestDepositedAndCalledIterator struct {
	Event *GatewayEVMTestDepositedAndCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestDepositedAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestDepositedAndCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestDepositedAndCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestDepositedAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestDepositedAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestDepositedAndCalled represents a DepositedAndCalled event raised by the GatewayEVMTest contract.
type GatewayEVMTestDepositedAndCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Amount        *big.Int
	Asset         common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositedAndCalled is a free log retrieval operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterDepositedAndCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMTestDepositedAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestDepositedAndCalledIterator{contract: _GatewayEVMTest.contract, event: "DepositedAndCalled", logs: logs, sub: sub}, nil
}

// WatchDepositedAndCalled is a free log subscription operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchDepositedAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestDepositedAndCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestDepositedAndCalled)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositedAndCalled is a log parse operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseDepositedAndCalled(log types.Log) (*GatewayEVMTestDepositedAndCalled, error) {
	event := new(GatewayEVMTestDepositedAndCalled)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedIterator struct {
	Event *GatewayEVMTestExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestExecuted represents a Executed event raised by the GatewayEVMTest contract.
type GatewayEVMTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestExecutedIterator{contract: _GatewayEVMTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestExecuted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecuted is a log parse operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseExecuted(log types.Log) (*GatewayEVMTestExecuted, error) {
	event := new(GatewayEVMTestExecuted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestExecutedV2Iterator is returned from FilterExecutedV2 and is used to iterate over the raw logs and unpacked data for ExecutedV2 events raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedV2Iterator struct {
	Event *GatewayEVMTestExecutedV2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestExecutedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestExecutedV2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestExecutedV2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestExecutedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestExecutedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestExecutedV2 represents a ExecutedV2 event raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedV2 struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutedV2 is a free log retrieval operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterExecutedV2(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMTestExecutedV2Iterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ExecutedV2", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestExecutedV2Iterator{contract: _GatewayEVMTest.contract, event: "ExecutedV2", logs: logs, sub: sub}, nil
}

// WatchExecutedV2 is a free log subscription operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchExecutedV2(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestExecutedV2, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ExecutedV2", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestExecutedV2)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ExecutedV2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutedV2 is a log parse operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseExecutedV2(log types.Log) (*GatewayEVMTestExecutedV2, error) {
	event := new(GatewayEVMTestExecutedV2)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ExecutedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMTestExecutedWithERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestExecutedWithERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestExecutedWithERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestExecutedWithERC20Iterator{contract: _GatewayEVMTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestExecutedWithERC20)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutedWithERC20 is a log parse operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMTestExecutedWithERC20, error) {
	event := new(GatewayEVMTestExecutedWithERC20)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedERC20Iterator struct {
	Event *GatewayEVMTestReceivedERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReceivedERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedERC20 represents a ReceivedERC20 event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*GatewayEVMTestReceivedERC20Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedERC20Iterator{contract: _GatewayEVMTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedERC20)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceivedERC20 is a log parse operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedERC20(log types.Log) (*GatewayEVMTestReceivedERC20, error) {
	event := new(GatewayEVMTestReceivedERC20)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNoParamsIterator struct {
	Event *GatewayEVMTestReceivedNoParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedNoParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReceivedNoParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedNoParams represents a ReceivedNoParams event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*GatewayEVMTestReceivedNoParamsIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedNoParamsIterator{contract: _GatewayEVMTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedNoParams)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceivedNoParams is a log parse operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedNoParams(log types.Log) (*GatewayEVMTestReceivedNoParams, error) {
	event := new(GatewayEVMTestReceivedNoParams)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNonPayableIterator struct {
	Event *GatewayEVMTestReceivedNonPayable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedNonPayable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReceivedNonPayable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedNonPayable represents a ReceivedNonPayable event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*GatewayEVMTestReceivedNonPayableIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedNonPayableIterator{contract: _GatewayEVMTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedNonPayable)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceivedNonPayable is a log parse operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedNonPayable(log types.Log) (*GatewayEVMTestReceivedNonPayable, error) {
	event := new(GatewayEVMTestReceivedNonPayable)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedOnCallIterator is returned from FilterReceivedOnCall and is used to iterate over the raw logs and unpacked data for ReceivedOnCall events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedOnCallIterator struct {
	Event *GatewayEVMTestReceivedOnCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestReceivedOnCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedOnCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReceivedOnCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestReceivedOnCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedOnCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedOnCall represents a ReceivedOnCall event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedOnCall struct {
	Sender  common.Address
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCall is a free log retrieval operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedOnCall(opts *bind.FilterOpts) (*GatewayEVMTestReceivedOnCallIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedOnCallIterator{contract: _GatewayEVMTest.contract, event: "ReceivedOnCall", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCall is a free log subscription operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedOnCall(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedOnCall) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedOnCall)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceivedOnCall is a log parse operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedOnCall(log types.Log) (*GatewayEVMTestReceivedOnCall, error) {
	event := new(GatewayEVMTestReceivedOnCall)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedPayableIterator struct {
	Event *GatewayEVMTestReceivedPayable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedPayable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReceivedPayable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedPayable represents a ReceivedPayable event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedPayable struct {
	Sender common.Address
	Value  *big.Int
	Str    string
	Num    *big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedPayable is a free log retrieval operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*GatewayEVMTestReceivedPayableIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedPayableIterator{contract: _GatewayEVMTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedPayable)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceivedPayable is a log parse operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedPayable(log types.Log) (*GatewayEVMTestReceivedPayable, error) {
	event := new(GatewayEVMTestReceivedPayable)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedRevertIterator struct {
	Event *GatewayEVMTestReceivedRevert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedRevert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReceivedRevert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedRevert represents a ReceivedRevert event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*GatewayEVMTestReceivedRevertIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedRevertIterator{contract: _GatewayEVMTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedRevert)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceivedRevert is a log parse operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedRevert(log types.Log) (*GatewayEVMTestReceivedRevert, error) {
	event := new(GatewayEVMTestReceivedRevert)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVMTest contract.
type GatewayEVMTestRevertedIterator struct {
	Event *GatewayEVMTestReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReverted represents a Reverted event raised by the GatewayEVMTest contract.
type GatewayEVMTestReverted struct {
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestRevertedIterator{contract: _GatewayEVMTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReverted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReverted is a log parse operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReverted(log types.Log) (*GatewayEVMTestReverted, error) {
	event := new(GatewayEVMTestReverted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUnwhitelistedIterator is returned from FilterUnwhitelisted and is used to iterate over the raw logs and unpacked data for Unwhitelisted events raised by the GatewayEVMTest contract.
type GatewayEVMTestUnwhitelistedIterator struct {
	Event *GatewayEVMTestUnwhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUnwhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestUnwhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUnwhitelisted represents a Unwhitelisted event raised by the GatewayEVMTest contract.
type GatewayEVMTestUnwhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnwhitelisted is a free log retrieval operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUnwhitelisted(opts *bind.FilterOpts, token []common.Address) (*GatewayEVMTestUnwhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUnwhitelistedIterator{contract: _GatewayEVMTest.contract, event: "Unwhitelisted", logs: logs, sub: sub}, nil
}

// WatchUnwhitelisted is a free log subscription operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUnwhitelisted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUnwhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUnwhitelisted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnwhitelisted is a log parse operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUnwhitelisted(log types.Log) (*GatewayEVMTestUnwhitelisted, error) {
	event := new(GatewayEVMTestUnwhitelisted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUpdatedCustodyTSSAddressIterator is returned from FilterUpdatedCustodyTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedCustodyTSSAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedCustodyTSSAddressIterator struct {
	Event *GatewayEVMTestUpdatedCustodyTSSAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestUpdatedCustodyTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUpdatedCustodyTSSAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestUpdatedCustodyTSSAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestUpdatedCustodyTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUpdatedCustodyTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUpdatedCustodyTSSAddress represents a UpdatedCustodyTSSAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedCustodyTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCustodyTSSAddress is a free log retrieval operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUpdatedCustodyTSSAddress(opts *bind.FilterOpts) (*GatewayEVMTestUpdatedCustodyTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUpdatedCustodyTSSAddressIterator{contract: _GatewayEVMTest.contract, event: "UpdatedCustodyTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedCustodyTSSAddress is a free log subscription operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUpdatedCustodyTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUpdatedCustodyTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUpdatedCustodyTSSAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedCustodyTSSAddress is a log parse operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUpdatedCustodyTSSAddress(log types.Log) (*GatewayEVMTestUpdatedCustodyTSSAddress, error) {
	event := new(GatewayEVMTestUpdatedCustodyTSSAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedGatewayTSSAddressIterator struct {
	Event *GatewayEVMTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUpdatedGatewayTSSAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestUpdatedGatewayTSSAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedGatewayTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*GatewayEVMTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUpdatedGatewayTSSAddressIterator{contract: _GatewayEVMTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUpdatedGatewayTSSAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedGatewayTSSAddress is a log parse operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*GatewayEVMTestUpdatedGatewayTSSAddress, error) {
	event := new(GatewayEVMTestUpdatedGatewayTSSAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the GatewayEVMTest contract.
type GatewayEVMTestWhitelistedIterator struct {
	Event *GatewayEVMTestWhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestWhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWhitelisted represents a Whitelisted event raised by the GatewayEVMTest contract.
type GatewayEVMTestWhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWhitelisted(opts *bind.FilterOpts, token []common.Address) (*GatewayEVMTestWhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWhitelistedIterator{contract: _GatewayEVMTest.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWhitelisted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Whitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelisted is a log parse operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWhitelisted(log types.Log) (*GatewayEVMTestWhitelisted, error) {
	event := new(GatewayEVMTestWhitelisted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnIterator struct {
	Event *GatewayEVMTestWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWithdrawn represents a Withdrawn event raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawn struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWithdrawnIterator{contract: _GatewayEVMTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWithdrawn, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWithdrawn)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWithdrawn(log types.Log) (*GatewayEVMTestWithdrawn, error) {
	event := new(GatewayEVMTestWithdrawn)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndCalledIterator struct {
	Event *GatewayEVMTestWithdrawnAndCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWithdrawnAndCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestWithdrawnAndCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndCalled struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWithdrawnAndCalledIterator{contract: _GatewayEVMTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWithdrawnAndCalled, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWithdrawnAndCalled)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayEVMTestWithdrawnAndCalled, error) {
	event := new(GatewayEVMTestWithdrawnAndCalled)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndRevertedIterator struct {
	Event *GatewayEVMTestWithdrawnAndReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWithdrawnAndReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestWithdrawnAndReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndReverted struct {
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWithdrawnAndRevertedIterator{contract: _GatewayEVMTest.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWithdrawnAndReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWithdrawnAndReverted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*GatewayEVMTestWithdrawnAndReverted, error) {
	event := new(GatewayEVMTestWithdrawnAndReverted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogIterator struct {
	Event *GatewayEVMTestLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLog represents a Log event raised by the GatewayEVMTest contract.
type GatewayEVMTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayEVMTestLogIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogIterator{contract: _GatewayEVMTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLog)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLog is a log parse operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLog(log types.Log) (*GatewayEVMTestLog, error) {
	event := new(GatewayEVMTestLog)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogAddressIterator struct {
	Event *GatewayEVMTestLogAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogAddress represents a LogAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayEVMTestLogAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogAddressIterator{contract: _GatewayEVMTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogAddress is a log parse operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogAddress(log types.Log) (*GatewayEVMTestLogAddress, error) {
	event := new(GatewayEVMTestLogAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArrayIterator struct {
	Event *GatewayEVMTestLogArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogArray represents a LogArray event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayEVMTestLogArrayIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogArrayIterator{contract: _GatewayEVMTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogArray)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogArray is a log parse operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogArray(log types.Log) (*GatewayEVMTestLogArray, error) {
	event := new(GatewayEVMTestLogArray)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray0Iterator struct {
	Event *GatewayEVMTestLogArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogArray0 represents a LogArray0 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayEVMTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogArray0Iterator{contract: _GatewayEVMTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogArray0)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogArray0 is a log parse operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogArray0(log types.Log) (*GatewayEVMTestLogArray0, error) {
	event := new(GatewayEVMTestLogArray0)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray1Iterator struct {
	Event *GatewayEVMTestLogArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogArray1 represents a LogArray1 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayEVMTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogArray1Iterator{contract: _GatewayEVMTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogArray1)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogArray1 is a log parse operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogArray1(log types.Log) (*GatewayEVMTestLogArray1, error) {
	event := new(GatewayEVMTestLogArray1)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytesIterator struct {
	Event *GatewayEVMTestLogBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogBytes represents a LogBytes event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayEVMTestLogBytesIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogBytesIterator{contract: _GatewayEVMTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogBytes)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogBytes is a log parse operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogBytes(log types.Log) (*GatewayEVMTestLogBytes, error) {
	event := new(GatewayEVMTestLogBytes)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytes32Iterator struct {
	Event *GatewayEVMTestLogBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogBytes32 represents a LogBytes32 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayEVMTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogBytes32Iterator{contract: _GatewayEVMTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogBytes32)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogBytes32 is a log parse operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogBytes32(log types.Log) (*GatewayEVMTestLogBytes32, error) {
	event := new(GatewayEVMTestLogBytes32)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogIntIterator struct {
	Event *GatewayEVMTestLogInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogInt represents a LogInt event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayEVMTestLogIntIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogIntIterator{contract: _GatewayEVMTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogInt)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogInt is a log parse operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogInt(log types.Log) (*GatewayEVMTestLogInt, error) {
	event := new(GatewayEVMTestLogInt)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedAddressIterator struct {
	Event *GatewayEVMTestLogNamedAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedAddressIterator{contract: _GatewayEVMTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedAddress is a log parse operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayEVMTestLogNamedAddress, error) {
	event := new(GatewayEVMTestLogNamedAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArrayIterator struct {
	Event *GatewayEVMTestLogNamedArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedArray represents a LogNamedArray event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedArrayIterator{contract: _GatewayEVMTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedArray)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedArray is a log parse operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayEVMTestLogNamedArray, error) {
	event := new(GatewayEVMTestLogNamedArray)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray0Iterator struct {
	Event *GatewayEVMTestLogNamedArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedArray0Iterator{contract: _GatewayEVMTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedArray0)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedArray0 is a log parse operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayEVMTestLogNamedArray0, error) {
	event := new(GatewayEVMTestLogNamedArray0)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray1Iterator struct {
	Event *GatewayEVMTestLogNamedArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedArray1Iterator{contract: _GatewayEVMTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedArray1)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedArray1 is a log parse operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayEVMTestLogNamedArray1, error) {
	event := new(GatewayEVMTestLogNamedArray1)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytesIterator struct {
	Event *GatewayEVMTestLogNamedBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedBytesIterator{contract: _GatewayEVMTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedBytes)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedBytes is a log parse operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayEVMTestLogNamedBytes, error) {
	event := new(GatewayEVMTestLogNamedBytes)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytes32Iterator struct {
	Event *GatewayEVMTestLogNamedBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedBytes32Iterator{contract: _GatewayEVMTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedBytes32)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedBytes32 is a log parse operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayEVMTestLogNamedBytes32, error) {
	event := new(GatewayEVMTestLogNamedBytes32)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalIntIterator struct {
	Event *GatewayEVMTestLogNamedDecimalInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedDecimalInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedDecimalInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedDecimalIntIterator{contract: _GatewayEVMTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedDecimalInt)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedDecimalInt is a log parse operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayEVMTestLogNamedDecimalInt, error) {
	event := new(GatewayEVMTestLogNamedDecimalInt)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalUintIterator struct {
	Event *GatewayEVMTestLogNamedDecimalUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedDecimalUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedDecimalUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedDecimalUintIterator{contract: _GatewayEVMTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedDecimalUint)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedDecimalUint is a log parse operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayEVMTestLogNamedDecimalUint, error) {
	event := new(GatewayEVMTestLogNamedDecimalUint)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedIntIterator struct {
	Event *GatewayEVMTestLogNamedInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedInt represents a LogNamedInt event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedIntIterator{contract: _GatewayEVMTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedInt)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedInt is a log parse operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayEVMTestLogNamedInt, error) {
	event := new(GatewayEVMTestLogNamedInt)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedStringIterator struct {
	Event *GatewayEVMTestLogNamedString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedString represents a LogNamedString event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedStringIterator{contract: _GatewayEVMTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedString)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedString is a log parse operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedString(log types.Log) (*GatewayEVMTestLogNamedString, error) {
	event := new(GatewayEVMTestLogNamedString)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedUintIterator struct {
	Event *GatewayEVMTestLogNamedUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedUint represents a LogNamedUint event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedUintIterator{contract: _GatewayEVMTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedUint)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedUint is a log parse operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayEVMTestLogNamedUint, error) {
	event := new(GatewayEVMTestLogNamedUint)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogStringIterator struct {
	Event *GatewayEVMTestLogString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogString represents a LogString event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayEVMTestLogStringIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogStringIterator{contract: _GatewayEVMTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogString)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogString is a log parse operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogString(log types.Log) (*GatewayEVMTestLogString, error) {
	event := new(GatewayEVMTestLogString)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogUintIterator struct {
	Event *GatewayEVMTestLogUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogUint represents a LogUint event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayEVMTestLogUintIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogUintIterator{contract: _GatewayEVMTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogUint)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogUint is a log parse operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogUint(log types.Log) (*GatewayEVMTestLogUint, error) {
	event := new(GatewayEVMTestLogUint)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogsIterator struct {
	Event *GatewayEVMTestLogs // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogs)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogs)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogs represents a Logs event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayEVMTestLogsIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogsIterator{contract: _GatewayEVMTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogs)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "logs", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogs is a log parse operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogs(log types.Log) (*GatewayEVMTestLogs, error) {
	event := new(GatewayEVMTestLogs)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
