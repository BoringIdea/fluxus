// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayzevm

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

// AbortContext is an auto generated low-level Go binding around an user-defined struct.
type AbortContext struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}

// CallOptions is an auto generated low-level Go binding around an user-defined struct.
type CallOptions struct {
	GasLimit        *big.Int
	IsArbitraryCall bool
}

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

// GatewayZEVMOutboundTestMetaData contains all meta data concerning the GatewayZEVMOutboundTest contract.
var GatewayZEVMOutboundTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForDifferentZRC20WithdrawAndCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForDifferentZRC20Withdrawal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForZRC20WithdrawAndCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForZRC20Withdrawal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnProtocolFeesFailsWithInsufficientAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDeposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfSenderNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETA\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractWhenPaused\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfInsufficientProtocolBalance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsWhenPaused\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteAbortUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContractIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContractFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextData\",\"inputs\":[{\"name\":\"origin\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"msgSender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataAbort\",\"inputs\":[{\"name\":\"abortContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAbortContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outgoing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataRevert\",\"inputs\":[{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroGasPrice\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b506201953f806200003e6000396000f3fe608060405234801561001057600080fd5b50600436106104de5760003560e01c80637cec29b011610286578063c35cb5e41161016b578063eab7674e116100e3578063f1d98f1b11610097578063f589655d1161007c578063f589655d1461078c578063fa7626d414610794578063fb339a1c146107a157600080fd5b8063f1d98f1b1461077c578063f2a913581461078457600080fd5b8063ec294d9f116100c8578063ec294d9f14610764578063ee0f4ea11461076c578063ef2b53941461077457600080fd5b8063eab7674e14610754578063eb78bd7d1461075c57600080fd5b8063da8fd81d1161013a578063e20c9f711161011f578063e20c9f711461070f578063e36c0a2314610717578063e63ab1e91461071f57600080fd5b8063da8fd81d146106ff578063e09bc6591461070757600080fd5b8063c35cb5e4146106df578063c8814d2e146106e7578063ca26929c146106ef578063cf2c3d1d146106f757600080fd5b8063996b7675116101fe578063b5508aa9116101cd578063b84a3d2f116101b2578063b84a3d2f146106bf578063b936be8c14610645578063ba414fa6146106c757600080fd5b8063b5508aa9146106af578063b82ce921146106b757600080fd5b8063996b76751461068f5780639c9acd5d14610697578063ae4662ec1461069f578063b0464fdc146106a757600080fd5b8063884660a311610255578063916a17c61161023a578063916a17c61461067257806396d9d8761461068757806397f7661f1461064557600080fd5b8063884660a314610662578063890a2d671461066a57600080fd5b80637cec29b0146106355780637f924c4e1461063d578063828d267c1461064557806385226c811461064d57600080fd5b806333f7c81a116103c7578063514fcc7f1161033f57806361cfddb7116102f357806366d9a9a0116102d857806366d9a9a0146106105780636efa04b51461062557806375ca55871461062d57600080fd5b806361cfddb71461060057806362fe29151461060857600080fd5b80635b4c90e1116103245780635b4c90e11461055b5780635cec7db5146105f05780636163f8ef146105f857600080fd5b8063514fcc7f146105e057806358c9987f146105e857600080fd5b80633f7286f41161039657806348f4fd071161037b57806348f4fd07146105c857806350dd38ca146105d057806351336fb0146105d857600080fd5b80633f7286f4146105b8578063445bbf47146105c057600080fd5b806333f7c81a146105985780633626c616146105a05780633ab5b199146105a85780633e5e3c23146105b057600080fd5b806321d925911161045a5780632acb21b4116104295780632ed7cff61161040e5780632ed7cff6146105805780633177f38114610588578063339bd8281461059057600080fd5b80632acb21b4146105635780632ade38801461056b57600080fd5b806321d92591146105435780632468bc0f1461054b57806327820625146105535780632948df411461055b57600080fd5b80631832cb6e116104b15780631c785a14116104965780631c785a14146105155780631ed7831c1461051d5780631ff5bdd71461053b57600080fd5b80631832cb6e14610505578063195ac1431461050d57600080fd5b8063084fafab146104e35780630a9254e4146104ed5780630c4a14b4146104f557806314b7a6da146104fd575b600080fd5b6104eb6107a9565b005b6104eb61097f565b6104eb6117f6565b6104eb611a39565b6104eb611b70565b6104eb611d11565b6104eb611e86565b6105256127d4565b604051610532919061de5c565b60405180910390f35b6104eb612836565b6104eb61295f565b6104eb612acf565b6104eb612dd8565b6104eb612ea0565b6104eb613054565b610573613213565b604051610532919061def8565b6104eb613355565b6104eb613483565b6104eb613b6f565b6104eb613ca2565b6104eb613ec3565b6104eb614030565b6105256141c8565b610525614228565b6104eb614288565b6104eb6143db565b6104eb61450d565b6104eb614656565b6104eb6147a1565b6104eb6148e7565b6104eb614a9c565b6104eb614c80565b6104eb614dce565b6104eb6152f7565b6106186157c7565b604051610532919061e05e565b6104eb615949565b6104eb615d1f565b6104eb61655d565b6104eb616748565b6104eb61679a565b610655616930565b604051610532919061e0fc565b6104eb616a00565b6104eb616d11565b61067a616ecc565b604051610532919061e173565b6104eb616fc7565b6104eb61711a565b6104eb61726e565b6104eb6173a5565b61067a617470565b61065561756b565b6104eb61763b565b6104eb61777f565b6106cf617ed0565b6040519015158152602001610532565b6104eb617fa4565b6104eb61815f565b6104eb618288565b6104eb61853d565b6104eb6186fb565b6104eb6187c6565b61052561897f565b6104eb6189df565b6107467f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b604051908152602001610532565b6104eb618aaa565b6104eb618bf8565b6104eb618db6565b6104eb618f65565b6104eb61908e565b6104eb619226565b6104eb619374565b6104eb6196eb565b601f546106cf9060ff1681565b6104eb61995b565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561083857600080fd5b505af115801561084c573d6000803e3d6000fd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602d604051610880919061e33b565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b1580156108e257600080fd5b505af11580156108f6573d6000803e3d6000fd5b50506020546024546040517f184b07930000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063184b0793935061094b9290911690602d9060040161e34e565b600060405180830381600087803b15801561096557600080fd5b505af1158015610979573d6000803e3d6000fd5b50505050565b602680547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602780549091166112341790556040516109c59061dd5d565b604051809103906000f0801580156109e1573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600f81527f476174657761795a45564d2e736f6c000000000000000000000000000000000060208201526026549151602481019390935292166044820152610ac5919060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f485cc95500000000000000000000000000000000000000000000000000000000179052619af5565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b039384168102919091179182905560208054919092049092167fffffffffffffffffffffffff0000000000000000000000000000000000000000909216919091178155604051737cce3eb018bf23e1fe2a32692f2c77592d1103949160009190610b6090820161dd6b565b601f1982820381018352601f909101166040819052610b82919060200161e370565b60405160208183030381529060405290506000808251602084016000f590507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663b4d6c78284836001600160a01b0316803b806020016040519081016040528181526000908060200190933c6040518363ffffffff1660e01b8152600401610c1792919061e38c565b600060405180830381600087803b158015610c3157600080fd5b505af1158015610c45573d6000803e3d6000fd5b50506026546020546040517fc0c53b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820181905260248201529082166044820152908616925063c0c53b8b9150606401600060405180830381600087803b158015610cbb57600080fd5b505af1158015610ccf573d6000803e3d6000fd5b5050602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038781169182179092556020546040517fab7b49930000000000000000000000000000000000000000000000000000000081526004810192909252909116925063ab7b49939150602401600060405180830381600087803b158015610d6257600080fd5b505af1158015610d76573d6000803e3d6000fd5b505060208054604080517f2722feee00000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169450632722feee935060048082019392918290030181865afa158015610dda573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dfe919061e3ca565b602880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055604051610e429061dd79565b604051809103906000f080158015610e5e573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161781556028546040517f06447d5600000000000000000000000000000000000000000000000000000000815292166004830152737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d569101600060405180830381600087803b158015610efa57600080fd5b505af1158015610f0e573d6000803e3d6000fd5b505050506000806000604051610f239061dd87565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610f5f573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602054604051601293600193849360009391921690610fb59061dd95565b610fc49695949392919061e420565b604051809103906000f080158015610fe0573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179091556023546040517fee2815ba0000000000000000000000000000000000000000000000000000000081526001600482015260248101929092529091169063ee2815ba90604401600060405180830381600087803b15801561107757600080fd5b505af115801561108b573d6000803e3d6000fd5b50506023546040517fa7cb050700000000000000000000000000000000000000000000000000000000815260016004820181905260248201526001600160a01b03909116925063a7cb05079150604401600060405180830381600087803b1580156110f557600080fd5b505af1158015611109573d6000803e3d6000fd5b50506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152633b9aca006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b15801561118957600080fd5b505af115801561119d573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b1580156111f257600080fd5b505af1158015611206573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af115801561127a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061129e919061e4da565b506021546026546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a060248201529116906347e7ef24906044016020604051808303816000875af115801561130f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611333919061e4da565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561139257600080fd5b505af11580156113a6573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561141c57600080fd5b505af1158015611430573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a060248201529116925063095ea7b391506044016020604051808303816000875af11580156114a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114c8919061e4da565b50602260009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b15801561151a57600080fd5b505af115801561152e573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af11580156115a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115c6919061e4da565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561162557600080fd5b505af1158015611639573d6000803e3d6000fd5b5050604080516080810182526026546001600160a01b0390811682526000602080840182815260018587019081528651928301909652918152606084018190528351602d80549185167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161781559251602e8054919095169116179092559251602f559093509091506030906116d1908261e572565b50506040805160c0810190915260265460601b6bffffffffffffffffffffffff191660e082015290508060f4810160408051601f198184030181529181529082526000602083810182905260018484018190526060850183905260808501528251908101909252815260a09091015280516031908190611751908261e572565b5060208201516001820180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039092169190911790556040820151600282015560608201516003820180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790556080820151600482015560a082015160058201906117ee908261e572565b505050505050565b602854602054602480546040517f81bad6f300000000000000000000000000000000000000000000000000000000815260016004820181905292810183905260448101839052606481018390526001600160a01b03918216608482018190529294821631939190911631913190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561189c57600080fd5b505af11580156118b0573d6000803e3d6000fd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602d6040516118e4919061e33b565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561194557600080fd5b505af1158015611959573d6000803e3d6000fd5b50506020546024546040517fe90b9e5e0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063e90b9e5e935088926119ae921690602d9060040161e34e565b6000604051808303818588803b1580156119c757600080fd5b505af11580156119db573d6000803e3d6000fd5b50506028546001600160a01b0316319250611a0391506119fd9050868661e660565b82619b14565b6020546001600160a01b031631611a1a8482619b14565b6117ee611a27878561e673565b6024546001600160a01b031631619b14565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611a9257600080fd5b505af1158015611aa6573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015611afd57600080fd5b505af1158015611b11573d6000803e3d6000fd5b50506020546021546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260016024820152600060448201529116925063f45346dc915060640161094b565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611bff57600080fd5b505af1158015611c13573d6000803e3d6000fd5b505050507f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db76031604051611c47919061e6e4565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611ca857600080fd5b505af1158015611cbc573d6000803e3d6000fd5b50506020546024546040517f2095dedb0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450632095dedb935061094b929091169060319060040161e6f7565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611d6a57600080fd5b505af1158015611d7e573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015611dee57600080fd5b505af1158015611e02573d6000803e3d6000fd5b50506020546027546040517ff340fa010000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063f340fa0191506000906024015b6000604051808303818588803b158015611e6b57600080fd5b505af1158015611e7f573d6000803e3d6000fd5b5050505050565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611edf57600080fd5b505af1158015611ef3573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611fde919060040161e719565b600060405180830381600087803b158015611ff857600080fd5b505af115801561200c573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561206057600080fd5b505af1158015612074573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156120d157600080fd5b505af11580156120e5573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506121d0919060040161e719565b600060405180830381600087803b1580156121ea57600080fd5b505af11580156121fe573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561225257600080fd5b505af1158015612266573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156122c357600080fd5b505af11580156122d7573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561232b57600080fd5b505af115801561233f573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd93c066500000000000000000000000000000000000000000000000000000000600482015260019250737109709ecfa91a80626ff3989d68f67f5b1dd12d915063c31eb0e090602401600060405180830381600087803b1580156123b257600080fd5b505af11580156123c6573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561242357600080fd5b505af1158015612437573d6000803e3d6000fd5b50506020546021546027546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810187905290821660448201529116925063f45346dc9150606401600060405180830381600087803b1580156124af57600080fd5b505af11580156124c3573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561252057600080fd5b505af1158015612534573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561258857600080fd5b505af115801561259c573d6000803e3d6000fd5b50506021546027546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009450911691506370a08231906024015b602060405180830381865afa158015612609573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061262d919061e72c565b905061263a600082619b14565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561269357600080fd5b505af11580156126a7573d6000803e3d6000fd5b50506020546021546027546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810188905290821660448201529116925063f45346dc9150606401600060405180830381600087803b15801561271f57600080fd5b505af1158015612733573d6000803e3d6000fd5b50506021546027546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561279f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127c3919061e72c565b90506127cf8382619b14565b505050565b6060601680548060200260200160405190810160405280929190818152602001828054801561282c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161280e575b5050505050905090565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561288f57600080fd5b505af11580156128a3573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156128fa57600080fd5b505af115801561290e573d6000803e3d6000fd5b50506020546040517ff340fa01000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b03909116925063f340fa019150600190602401611e52565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156129ce57600080fd5b505af11580156129e2573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015612a3f57600080fd5b505af1158015612a53573d6000803e3d6000fd5b50506020546028546040517ff340fa010000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063f340fa01915083906024015b6000604051808303818588803b158015612abb57600080fd5b505af11580156117ee573d6000803e3d6000fd5b602854602080546024546040516001946001600160a01b039081163194938116319392163191600091612b02910161e745565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff191660808301529150600090806094810160408051808303601f190181529181529082526028546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015612bf057600080fd5b505af1158015612c04573d6000803e3d6000fd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e9450612c5f93506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602854602054612c8f936001600160a01b03928316928c92169061e782565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612cf057600080fd5b505af1158015612d04573d6000803e3d6000fd5b50506020546024546040517f30b103420000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506330b1034293508a92612d5b928792911690889060040161e833565b6000604051808303818588803b158015612d7457600080fd5b505af1158015612d88573d6000803e3d6000fd5b50506028546001600160a01b0316319250612daa91506119fd9050888861e660565b6020546001600160a01b031631612dc18682619b14565b612dce611a27898761e673565b5050505050505050565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612e4457600080fd5b505af1158015612e58573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024016108c8565b6000604051602001612eb19061e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015612f5f57600080fd5b505af1158015612f73573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015612fca57600080fd5b505af1158015612fde573d6000803e3d6000fd5b50506020546024546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063bcf7f32b935061303a92869260009260019290911690899060040161e867565b600060405180830381600087803b158015612abb57600080fd5b60006040516020016130659061e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b15801561312f57600080fd5b505af1158015613143573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015b600060405180830381600087803b1580156131a157600080fd5b505af11580156131b5573d6000803e3d6000fd5b50506020546021546024546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063bcf7f32b945061303a938793811692600192911690899060040161e867565b6060601e805480602002602001604051908101604052809291908181526020016000905b8282101561334c57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156133355783829060005260206000200180546132a89061e20a565b80601f01602080910402602001604051908101604052809291908181526020018280546132d49061e20a565b80156133215780601f106132f657610100808354040283529160200191613321565b820191906000526020600020905b81548152906001019060200180831161330457829003601f168201915b505050505081526020019060010190613289565b505050508152505081526020019060010190613237565b50505050905090565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156133b157600080fd5b505af11580156133c5573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561341c57600080fd5b505af1158015613430573d6000803e3d6000fd5b50506020546040517fe90b9e5e0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063e90b9e5e91508390612aa290600090602d9060040161e34e565b6027546040516001600160a01b03909116602482015260019060009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae7600000000000000000000000000000000000000000000000000000000017905260285490517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561356e57600080fd5b505af1158015613582573d6000803e3d6000fd5b50506021546040517ff687d12a000000000000000000000000000000000000000000000000000000008152620186a060048201526001600160a01b03909116925063f687d12a9150602401600060405180830381600087803b1580156135e757600080fd5b505af11580156135fb573d6000803e3d6000fd5b50506021546040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152620186a06004820152600093506001600160a01b03909116915063fc5fecd5906024016040805180830381865afa158015613665573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613689919061e8bc565b6021546026549193506001600160a01b0390811692506347e7ef2491166136b0848761e673565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015613713573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613737919061e4da565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561379657600080fd5b505af11580156137aa573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561382057600080fd5b505af1158015613834573d6000803e3d6000fd5b50506021546020546001600160a01b03918216935063095ea7b392501661385b848761e673565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af11580156138be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138e2919061e4da565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561394157600080fd5b505af1158015613955573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156139ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139d2919061e72c565b602080546027546040519394506001600160a01b0391821693637b15118b93613a189392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f198184030181526021548383018352620186a084526001602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b168152613a879391928a926001600160a01b03909116918a919060299060040161e945565b600060405180830381600087803b158015613aa157600080fd5b505af1158015613ab5573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613b0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b32919061e72c565b90506000613b40848761e673565b90506117ee613b4f828561e660565b836040518060600160405280603a8152602001620194d0603a9139619b94565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613bc857600080fd5b505af1158015613bdc573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015613c3357600080fd5b505af1158015613c47573d6000803e3d6000fd5b50506020546021546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450639d4ba465935061094b9290911690600190600090602d9060040161e9b3565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613cfe57600080fd5b505af1158015613d12573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b158015613d6657600080fd5b505af1158015613d7a573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd93c0665000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015613dea57600080fd5b505af1158015613dfe573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015b600060405180830381600087803b158015613e5c57600080fd5b505af1158015613e70573d6000803e3d6000fd5b50506020546027546040517ff340fa010000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063f340fa0191508390602401612aa2565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613f1f57600080fd5b505af1158015613f33573d6000803e3d6000fd5b5050604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015613fa357600080fd5b505af1158015613fb7573d6000803e3d6000fd5b50506020546021546028546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810187905290821660448201529116925063f45346dc91506064015b600060405180830381600087803b158015611e6b57600080fd5b60006040516020016140419061e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b1580156140ef57600080fd5b505af1158015614103573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561415a57600080fd5b505af115801561416e573d6000803e3d6000fd5b50506020546021546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063bcf7f32b935061303a9286921690600190600090899060040161e867565b6060601880548060200260200160405190810160405280929190818152602001828054801561282c576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161280e575050505050905090565b6060601780548060200260200160405190810160405280929190818152602001828054801561282c576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161280e575050505050905090565b60285460205460275460405163ca669fa760e01b81526001600160a01b039384166004820181905260019490319381163192163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156142f657600080fd5b505af115801561430a573d6000803e3d6000fd5b50506020546027546040517ff340fa010000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063f340fa01915086906024016000604051808303818588803b15801561437157600080fd5b505af1158015614385573d6000803e3d6000fd5b50506028546020546027546001600160a01b03928316319550908216319350163190506143bb6143b5888861e660565b84619b14565b6143c58583619b14565b6143d26119fd888661e673565b50505050505050565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561443457600080fd5b505af1158015614448573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561449f57600080fd5b505af11580156144b3573d6000803e3d6000fd5b50506020546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450639d4ba465935061094b92600092600192911690602d9060040161e9b3565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561457c57600080fd5b505af1158015614590573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156145ed57600080fd5b505af1158015614601573d6000803e3d6000fd5b50506020546028546040517fe90b9e5e0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063e90b9e5e93508592612aa2921690602d9060040161e34e565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156146c257600080fd5b505af11580156146d6573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561473357600080fd5b505af1158015614747573d6000803e3d6000fd5b50506020546021546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450639d4ba465935061094b92909116906001908590602d9060040161e9b3565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561481057600080fd5b505af1158015614824573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561488157600080fd5b505af1158015614895573d6000803e3d6000fd5b50506020546040517fe90b9e5e0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063e90b9e5e91508390612aa2908490602d9060040161e34e565b60006040516020016148f89061e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b1580156149a657600080fd5b505af11580156149ba573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015614a2a57600080fd5b505af1158015614a3e573d6000803e3d6000fd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca37945061303a938793811692600092911690899060040161e867565b604051600190600090614ab19060200161e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b158015614b7b57600080fd5b505af1158015614b8f573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015614bec57600080fd5b505af1158015614c00573d6000803e3d6000fd5b50506020546040517f30b103420000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506330b1034291508590614c539085908590889060040161e833565b6000604051808303818588803b158015614c6c57600080fd5b505af1158015612dce573d6000803e3d6000fd5b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015614cec57600080fd5b505af1158015614d00573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015614d5d57600080fd5b505af1158015614d71573d6000803e3d6000fd5b50506020546021546028546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061094b93928316926001921690602d9060040161e9b3565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015614e4357600080fd5b505af1158015614e57573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018690529116925063095ea7b391506044016020604051808303816000875af1158015614ec9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614eed919061e4da565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015614f4c57600080fd5b505af1158015614f60573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015614fbd57600080fd5b505af1158015614fd1573d6000803e3d6000fd5b50506021546040517ff687d12a00000000000000000000000000000000000000000000000000000000815262030d4060048201526001600160a01b03909116925063f687d12a9150602401600060405180830381600087803b15801561503657600080fd5b505af115801561504a573d6000803e3d6000fd5b50506021546040517ffc5fecd500000000000000000000000000000000000000000000000000000000815262030d406004820152600093508392506001600160a01b039091169063fc5fecd5906024016040805180830381865afa1580156150b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906150da919061e8bc565b60208054604080516001600160a01b03868116602483015290921660448301526064808301859052815180840390910181526084909201815291810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f667011120000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152929450909250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916151ad9160040161e719565b600060405180830381600087803b1580156151c757600080fd5b505af11580156151db573d6000803e3d6000fd5b5050602080546027546040805160609290921b6bffffffffffffffffffffffff19168285015280518083036014018152602154600460348501818152605886018552605490950180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f19ff1d21000000000000000000000000000000000000000000000000000000001790528351808501855262030d40815260019781019790975292517f06cb89830000000000000000000000000000000000000000000000000000000081526001600160a01b0395861698506306cb898397506152c996929590911693926029910161e9ea565b600060405180830381600087803b1580156152e357600080fd5b505af11580156143d2573d6000803e3d6000fd5b60265460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561535357600080fd5b505af1158015615367573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156153bb57600080fd5b505af11580156153cf573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd93c0665000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561543f57600080fd5b505af1158015615453573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156154b057600080fd5b505af11580156154c4573d6000803e3d6000fd5b50506020546024546040517fe90b9e5e0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063e90b9e5e93508592615519921690602d9060040161e34e565b6000604051808303818588803b15801561553257600080fd5b505af1158015615546573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063ca669fa792506024019050600060405180830381600087803b1580156155a557600080fd5b505af11580156155b9573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561560d57600080fd5b505af1158015615621573d6000803e3d6000fd5b5050602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156156b457600080fd5b505af11580156156c8573d6000803e3d6000fd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602d6040516156fc919061e33b565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b15801561575e57600080fd5b505af1158015615772573d6000803e3d6000fd5b50506020546024546040517fe90b9e5e0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063e90b9e5e93508592612aa2921690602d9060040161e34e565b6060601b805480602002602001604051908101604052809291908181526020016000905b8282101561334c578382906000526020600020906002020160405180604001604052908160008201805461581e9061e20a565b80601f016020809104026020016040519081016040528092919081815260200182805461584a9061e20a565b80156158975780601f1061586c57610100808354040283529160200191615897565b820191906000526020600020905b81548152906001019060200180831161587a57829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561593157602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116158de5790505b505050505081525050815260200190600101906157eb565b602154602480546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009391909116916370a082319101602060405180830381865afa1580156159b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906159d6919061e72c565b90506159e3600082619b14565b60006040516020016159f49061e745565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff191660808301529150600090806094810160408051808303601f190181529181529082526028546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015615ae257600080fd5b505af1158015615af6573d6000803e3d6000fd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e9450615b5193506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602854602054615b82936001600160a01b0392831692600192169061e782565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015615be357600080fd5b505af1158015615bf7573d6000803e3d6000fd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450615c55938793811692600192911690899060040161e867565b600060405180830381600087803b158015615c6f57600080fd5b505af1158015615c83573d6000803e3d6000fd5b5050602154602480546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009550921692506370a082319101602060405180830381865afa158015615cee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615d12919061e72c565b9050610979600182619b14565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015615d9157600080fd5b505af1158015615da5573d6000803e3d6000fd5b50506023546020546040516000945060129350600192600292620186a0926001600160a01b039283169290911690615ddc9061dd95565b615deb9695949392919061ea45565b604051809103906000f080158015615e07573d6000803e3d6000fd5b506026546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152606460248201529192508216906347e7ef24906044016020604051808303816000875af1158015615e76573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615e9a919061e4da565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015615ef957600080fd5b505af1158015615f0d573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015615f8357600080fd5b505af1158015615f97573d6000803e3d6000fd5b50506020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260646024820152908416925063095ea7b391506044016020604051808303816000875af1158015616007573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061602b919061e4da565b506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482015291169063095ea7b3906044016020604051808303816000875af115801561609c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906160c0919061e4da565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561611f57600080fd5b505af1158015616133573d6000803e3d6000fd5b50506027546040516001600160a01b039091166024820152600192506000915060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae76000000000000000000000000000000000000000000000000000000000179052517ffc5fecd5000000000000000000000000000000000000000000000000000000008152620186a0600482015290915060009081906001600160a01b0386169063fc5fecd5906024016040805180830381865afa158015616213573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616237919061e8bc565b60215491935091506162539083906001600160a01b0316619c14565b602154604080517f18160ddd00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916318160ddd9160048083019260209291908290030181865afa1580156162b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906162da919061e72c565b90506000866001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561631c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616340919061e72c565b602080546027546040519394506001600160a01b0391821693637b15118b936163869392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f198184030181528282018252620186a083526001602084015290517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526163e5928b918d918c9160299060040161e945565b600060405180830381600087803b1580156163ff57600080fd5b505af1158015616413573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561646c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616490919061e72c565b90506000886001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156164d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906164f6919061e72c565b9050616525616505868661e660565b836040518060600160405280602b815260200162019447602b9139619b94565b616552616532898561e660565b826040518060600160405280603881526020016201949860389139619b94565b505050505050505050565b600060405160200161656e9061e745565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff191660808301529150600090806094810160408051808303601f190181529181529082526028546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561665c57600080fd5b505af1158015616670573d6000803e3d6000fd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e94506166cb93506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526028546020546166fc936001600160a01b0392831692600192169061e782565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401613187565b6021546027546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260019260009216906370a08231906024016125ec565b6040516001906000906167af9060200161e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b15801561685d57600080fd5b505af1158015616871573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156168c857600080fd5b505af11580156168dc573d6000803e3d6000fd5b50506020546040517f30b103420000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506330b1034291508590614c53908590600090889060040161e833565b6060601a805480602002602001604051908101604052809291908181526020016000905b8282101561334c5783829060005260206000200180546169739061e20a565b80601f016020809104026020016040519081016040528092919081815260200182805461699f9061e20a565b80156169ec5780601f106169c1576101008083540402835291602001916169ec565b820191906000526020600020905b8154815290600101906020018083116169cf57829003601f168201915b505050505081526020019060010190616954565b602154602480546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009391909116916370a082319101602060405180830381865afa158015616a69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616a8d919061e72c565b9050616a9a600082619b14565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015616b2957600080fd5b505af1158015616b3d573d6000803e3d6000fd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602d604051616b71919061e33b565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015616bd257600080fd5b505af1158015616be6573d6000803e3d6000fd5b50506020546021546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba4659450616c4393928316926001921690602d9060040161e9b3565b600060405180830381600087803b158015616c5d57600080fd5b505af1158015616c71573d6000803e3d6000fd5b5050602154602480546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009550921692506370a082319101602060405180830381865afa158015616cdc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616d00919061e72c565b9050616d0d600182619b14565b5050565b604051600190600090616d269060200161e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b158015616df057600080fd5b505af1158015616e04573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015616e6157600080fd5b505af1158015616e75573d6000803e3d6000fd5b50506020546028546040517f30b103420000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506330b1034293508792614c53928792911690889060040161e833565b6060601d805480602002602001604051908101604052809291908181526020016000905b8282101561334c5760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015616faf57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411616f5c5790505b50505050508152505081526020019060010190616ef0565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561702357600080fd5b505af1158015617037573d6000803e3d6000fd5b5050604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156170a757600080fd5b505af11580156170bb573d6000803e3d6000fd5b50506020546021546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101869052911660448201819052925063f45346dc9150606401614016565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561717357600080fd5b505af1158015617187573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156171f757600080fd5b505af115801561720b573d6000803e3d6000fd5b50506020546021546027546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526000602482015290821660448201529116925063f45346dc915060640161094b565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156172c757600080fd5b505af11580156172db573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561733257600080fd5b505af1158015617346573d6000803e3d6000fd5b50506020546027546040517ff45346dc00000000000000000000000000000000000000000000000000000000815260006004820152600160248201526001600160a01b0391821660448201529116925063f45346dc915060640161094b565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561741457600080fd5b505af1158015617428573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401613e42565b6060601c805480602002602001604051908101604052809291908181526020016000905b8282101561334c5760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561755357602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116175005790505b50505050508152505081526020019060010190617494565b60606019805480602002602001604051908101604052809291908181526020016000905b8282101561334c5783829060005260206000200180546175ae9061e20a565b80601f01602080910402602001604051908101604052809291908181526020018280546175da9061e20a565b80156176275780601f106175fc57610100808354040283529160200191617627565b820191906000526020600020905b81548152906001019060200180831161760a57829003601f168201915b50505050508152602001906001019061758f565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156176aa57600080fd5b505af11580156176be573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561771b57600080fd5b505af115801561772f573d6000803e3d6000fd5b50506020546040517ff340fa010000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201819052925063f340fa0191508390602401612aa2565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156177f157600080fd5b505af1158015617805573d6000803e3d6000fd5b50506023546020546040516000945060129350600192600292620186a0926001600160a01b03928316929091169061783c9061dd95565b61784b9695949392919061ea45565b604051809103906000f080158015617867573d6000803e3d6000fd5b506026546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152606460248201529192508216906347e7ef24906044016020604051808303816000875af11580156178d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906178fa919061e4da565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561795957600080fd5b505af115801561796d573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b1580156179e357600080fd5b505af11580156179f7573d6000803e3d6000fd5b50506020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260646024820152908416925063095ea7b391506044016020604051808303816000875af1158015617a67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617a8b919061e4da565b506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482015291169063095ea7b3906044016020604051808303816000875af1158015617afc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617b20919061e4da565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015617b7f57600080fd5b505af1158015617b93573d6000803e3d6000fd5b50506040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152620186a06004820152600192506000915081906001600160a01b0385169063fc5fecd5906024016040805180830381865afa158015617bfe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617c22919061e8bc565b6021549193509150617c3e9083906001600160a01b0316619c14565b602154604080517f18160ddd00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916318160ddd9160048083019260209291908290030181865afa158015617ca1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617cc5919061e72c565b90506000856001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015617d07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617d2b919061e72c565b602080546027546040519394506001600160a01b0391821693637c0dcb5f93617d719392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b604051602081830303815290604052878960296040518563ffffffff1660e01b8152600401617da3949392919061ead5565b600060405180830381600087803b158015617dbd57600080fd5b505af1158015617dd1573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015617e2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617e4e919061e72c565b90506000876001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015617e90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617eb4919061e72c565b9050617ec3616505868661e660565b612dce616532888561e660565b60085460009060ff1615617ee8575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015617f79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617f9d919061e72c565b1415905090565b604051600190600090617fb99060200161e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b15801561808357600080fd5b505af1158015618097573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156180f457600080fd5b505af1158015618108573d6000803e3d6000fd5b50506020546024546040517f30b103420000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506330b1034293508792614c53928792911690889060040161e833565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156181b857600080fd5b505af11580156181cc573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561822357600080fd5b505af1158015618237573d6000803e3d6000fd5b50506020546040517f184b07930000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063184b0793915061094b90600090602d9060040161e34e565b6021546027546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa1580156182f2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618316919061e72c565b9050618323600082619b14565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561837c57600080fd5b505af1158015618390573d6000803e3d6000fd5b5050604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561840057600080fd5b505af1158015618414573d6000803e3d6000fd5b50506020546021546027546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810188905290821660448201529116925063f45346dc9150606401600060405180830381600087803b15801561848c57600080fd5b505af11580156184a0573d6000803e3d6000fd5b50506021546027546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561850c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618530919061e72c565b90506127cf600082619b14565b600060405160200161854e9061e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b15801561861857600080fd5b505af115801561862c573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561868957600080fd5b505af115801561869d573d6000803e3d6000fd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca37945061303a938793811692600192911690899060040161e867565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152600090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561875757600080fd5b505af115801561876b573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401615744565b60006040516020016187d79061e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b1580156188a157600080fd5b505af11580156188b5573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561891257600080fd5b505af1158015618926573d6000803e3d6000fd5b50506020546021546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca37935061303a92869216906001908690899060040161e867565b6060601580548060200260200160405190810160405280929190818152602001828054801561282c576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161280e575050505050905090565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015618a4e57600080fd5b505af1158015618a62573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401615744565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015618b0357600080fd5b505af1158015618b17573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015618b8757600080fd5b505af1158015618b9b573d6000803e3d6000fd5b50506020546021546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061094b93928316926000921690602d9060040161e9b3565b6000604051602001618c099061e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b158015618cd357600080fd5b505af1158015618ce7573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015618d4457600080fd5b505af1158015618d58573d6000803e3d6000fd5b50506020546021546028546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca37945061303a938793811692600192911690899060040161e867565b600080604051602001618dc89061e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015618e7657600080fd5b505af1158015618e8a573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015618efa57600080fd5b505af1158015618f0e573d6000803e3d6000fd5b50506020546021546040517f30b103420000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506330b1034293508792614c53928792911690889060040161e833565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015618fbe57600080fd5b505af1158015618fd2573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561902957600080fd5b505af115801561903d573d6000803e3d6000fd5b50506020546040517f2095dedb0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632095dedb915061094b9060009060319060040161e6f7565b600060405160200161909f9061e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b15801561914d57600080fd5b505af1158015619161573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156191b857600080fd5b505af11580156191cc573d6000803e3d6000fd5b50506020546021546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca37935061303a9286921690600190600090899060040161e867565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561929257600080fd5b505af11580156192a6573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561930357600080fd5b505af1158015619317573d6000803e3d6000fd5b50506020546021546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061094b93928316926001921690602d9060040161e9b3565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156193d057600080fd5b505af11580156193e4573d6000803e3d6000fd5b50506021546040517ff687d12a00000000000000000000000000000000000000000000000000000000815261c35060048201526001600160a01b03909116925063f687d12a9150602401600060405180830381600087803b15801561944857600080fd5b505af115801561945c573d6000803e3d6000fd5b50506021546040517ffc5fecd500000000000000000000000000000000000000000000000000000000815261c3506004820152600093506001600160a01b03909116915063fc5fecd5906024016040805180830381865afa1580156194c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906194e9919061e8bc565b9150506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015619541573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190619565919061e72c565b602080546027546040519394506001600160a01b0391821693637c0dcb5f936195ab9392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526196039288916001600160a01b03169060299060040161ead5565b600060405180830381600087803b15801561961d57600080fd5b505af1158015619631573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561968a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906196ae919061e72c565b905060006196bc848661e673565b9050611e7f6196cb828561e660565b836040518060600160405280602681526020016201947260269139619b94565b60285460405163ca669fa760e01b81526001600160a01b03909116600482018190523190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561974857600080fd5b505af115801561975c573d6000803e3d6000fd5b5061099992506108fc9150619774905060018461e660565b6040518115909202916000818181858888f1935050505015801561979c573d6000803e3d6000fd5b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156197fb57600080fd5b505af115801561980f573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561986c57600080fd5b505af1158015619880573d6000803e3d6000fd5b50506020546027546040516001600160a01b039182166024820152600094509116915060029060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff340fa01000000000000000000000000000000000000000000000000000000001790525161990c919061e370565b60006040518083038185875af1925050503d8060008114619949576040519150601f19603f3d011682016040523d82523d6000602084013e61994e565b606091505b50509050616d0d81619c75565b600060405160200161996c9061e745565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015619a1a57600080fd5b505af1158015619a2e573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015619a8557600080fd5b505af1158015619a99573d6000803e3d6000fd5b50506020546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca37935061303a92869260009260019290911690899060040161e867565b6000619aff61dda3565b619b0a848483619cee565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015b60006040518083038186803b158015619b8057600080fd5b505afa1580156117ee573d6000803e3d6000fd5b6040517f88b44c85000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d906388b44c8590619be89086908690869060040161eb0f565b60006040518083038186803b158015619c0057600080fd5b505afa1580156143d2573d6000803e3d6000fd5b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f690604401619b68565b6040517fa59828850000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a59828859060240160006040518083038186803b158015619cda57600080fd5b505afa158015611e7f573d6000803e3d6000fd5b600080619cfb8584619d69565b9050619d5e6040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f78790000008152508286604051602001619d4992919061e38c565b60405160208183030381529060405285619d75565b9150505b9392505050565b6000619d628383619da3565b60c08101515160009015619d9957619d9284848460c00151619dbe565b9050619d62565b619d928484619f64565b6000619daf838361a04f565b619d6283836020015184619d75565b600080619dc961a05b565b90506000619dd7868361a12e565b90506000619dee826060015183602001518561a5d4565b90506000619dfe8383898961a7e6565b90506000619e0b8261b663565b602081015181519192509060030b15619e7e57898260400151604051602001619e3592919061eb2e565b60408051601f19818403018152908290527f08c379a0000000000000000000000000000000000000000000000000000000008252619e759160040161e719565b60405180910390fd5b6000619ec16040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a20000000000000000000000081525083600161b832565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d90619f1490849060040161e719565b602060405180830381865afa158015619f31573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190619f55919061e3ca565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc92590619fb990879060040161e719565b600060405180830381865afa158015619fd6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619ffe919081019061ec68565b9050600061a02c828560405160200161a01892919061ec9d565b60405160208183030381529060405261ba32565b90506001600160a01b038116619b0a578484604051602001619e3592919061eccc565b616d0d8282600061ba45565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c9061a0e290849060040161ed77565b600060405180830381865afa15801561a0ff573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a127919081019061edbe565b9250505090565b61a1606040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d905061a1ab6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b61a1b48561bb48565b6020820152600061a1c48661bf2d565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa15801561a206573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a22e919081019061edbe565b8683856020015160405160200161a248949392919061ee07565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb119061a2a090859060040161e719565b600060405180830381865afa15801561a2bd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a2e5919081019061edbe565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f69061a32d90849060040161ef0b565b602060405180830381865afa15801561a34a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061a36e919061e4da565b61a3835781604051602001619e35919061ef5d565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061a3c890849060040161efef565b600060405180830381865afa15801561a3e5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a40d919081019061edbe565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f69061a45490849060040161f041565b602060405180830381865afa15801561a471573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061a495919061e4da565b1561a52a576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061a4df90849060040161f041565b600060405180830381865afa15801561a4fc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a524919081019061edbe565b60408501525b846001600160a01b03166349c4fac882866000015160405160200161a54f919061f093565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161a57b92919061f0ff565b600060405180830381865afa15801561a598573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a5c0919081019061edbe565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b606081526020019060019003908161a5f05790505090506040518060400160405280600481526020017f67726570000000000000000000000000000000000000000000000000000000008152508160008151811061a6505761a65061f124565b60200260200101819052506040518060400160405280600381526020017f2d726c00000000000000000000000000000000000000000000000000000000008152508160018151811061a6a45761a6a461f124565b60200260200101819052508460405160200161a6c0919061f153565b6040516020818303038152906040528160028151811061a6e25761a6e261f124565b60200260200101819052508260405160200161a6fe919061f1bf565b6040516020818303038152906040528160038151811061a7205761a72061f124565b6020026020010181905250600061a7368261b663565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184526000808252908601528251808401909352905182529281019290925291925061a7c7906040805180820182526000808252602091820152815180830190925284518252808501908201529061c1b0565b61a7dc5785604051602001619e35919061f200565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d901561a836565b511590565b61a9aa5782602001511561a8f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401619e75565b8260c001511561a9aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401619e75565b6040805160ff8082526120008201909252600091816020015b606081526020019060019003908161a9c357905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061aa1e9061f291565b935060ff168151811061aa335761aa3361f124565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e370000000000000000000000000000000000000081525060405160200161aa84919061f2b0565b60405160208183030381529060405282828061aa9f9061f291565b935060ff168151811061aab45761aab461f124565b60200260200101819052506040518060400160405280600681526020017f6465706c6f79000000000000000000000000000000000000000000000000000081525082828061ab019061f291565b935060ff168151811061ab165761ab1661f124565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d6500000000000000000000000000000000000081525082828061ab639061f291565b935060ff168151811061ab785761ab7861f124565b6020026020010181905250876020015182828061ab949061f291565b935060ff168151811061aba95761aba961f124565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163745061746800000000000000000000000000000000000081525082828061abf69061f291565b935060ff168151811061ac0b5761ac0b61f124565b60209081029190910101528751828261ac238161f291565b935060ff168151811061ac385761ac3861f124565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e4964000000000000000000000000000000000000000000000081525082828061ac859061f291565b935060ff168151811061ac9a5761ac9a61f124565b602002602001018190525061acae4661c211565b828261acb98161f291565b935060ff168151811061acce5761acce61f124565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c65000000000000000000000000000000000081525082828061ad1b9061f291565b935060ff168151811061ad305761ad3061f124565b60200260200101819052508682828061ad489061f291565b935060ff168151811061ad5d5761ad5d61f124565b602090810291909101015285511561ae845760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f646500000000000000000000006020820152828261adae8161f291565b935060ff168151811061adc35761adc361f124565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d9061ae1390899060040161e719565b600060405180830381865afa15801561ae30573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261ae58919081019061edbe565b828261ae638161f291565b935060ff168151811061ae785761ae7861f124565b60200260200101819052505b84602001511561af545760408051808201909152601281527f2d2d766572696679536f75726365436f646500000000000000000000000000006020820152828261aecd8161f291565b935060ff168151811061aee25761aee261f124565b60200260200101819052506040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525082828061af2f9061f291565b935060ff168151811061af445761af4461f124565b602002602001018190525061b11b565b61af8c61a8318660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b61b01f5760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261afcf8161f291565b935060ff168151811061afe45761afe461f124565b60200260200101819052508460a0015160405160200161b004919061f153565b60405160208183030381529060405282828061af2f9061f291565b8460c0015115801561b06257506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261b06090511590565b155b1561b11b5760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261b0a68161f291565b935060ff168151811061b0bb5761b0bb61f124565b602002602001018190525061b0cf8861c2b1565b60405160200161b0df919061f153565b60405160208183030381529060405282828061b0fa9061f291565b935060ff168151811061b10f5761b10f61f124565b60200260200101819052505b6040808601518151808301835260008082526020918201528251808401909352815183529081019082015261b14f90511590565b61b1e45760408051808201909152600b81527f2d2d72656c6179657249640000000000000000000000000000000000000000006020820152828261b1928161f291565b935060ff168151811061b1a75761b1a761f124565b6020026020010181905250846040015182828061b1c39061f291565b935060ff168151811061b1d85761b1d861f124565b60200260200101819052505b60608501511561b3055760408051808201909152600681527f2d2d73616c7400000000000000000000000000000000000000000000000000006020820152828261b22d8161f291565b935060ff168151811061b2425761b24261f124565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa15801561b2b1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261b2d9919081019061edbe565b828261b2e48161f291565b935060ff168151811061b2f95761b2f961f124565b60200260200101819052505b60e0850151511561b3ac5760408051808201909152600a81527f2d2d6761734c696d6974000000000000000000000000000000000000000000006020820152828261b34f8161f291565b935060ff168151811061b3645761b36461f124565b602002602001018190525061b3808560e001516000015161c211565b828261b38b8161f291565b935060ff168151811061b3a05761b3a061f124565b60200260200101819052505b60e0850151602001511561b4565760408051808201909152600a81527f2d2d6761735072696365000000000000000000000000000000000000000000006020820152828261b3f98161f291565b935060ff168151811061b40e5761b40e61f124565b602002602001018190525061b42a8560e001516020015161c211565b828261b4358161f291565b935060ff168151811061b44a5761b44a61f124565b60200260200101819052505b60e0850151604001511561b5005760408051808201909152600e81527f2d2d6d61784665655065724761730000000000000000000000000000000000006020820152828261b4a38161f291565b935060ff168151811061b4b85761b4b861f124565b602002602001018190525061b4d48560e001516040015161c211565b828261b4df8161f291565b935060ff168151811061b4f45761b4f461f124565b60200260200101819052505b60e0850151606001511561b5aa5760408051808201909152601681527f2d2d6d61785072696f72697479466565506572476173000000000000000000006020820152828261b54d8161f291565b935060ff168151811061b5625761b56261f124565b602002602001018190525061b57e8560e001516060015161c211565b828261b5898161f291565b935060ff168151811061b59e5761b59e61f124565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561b5c85761b5c861e4fc565b60405190808252806020026020018201604052801561b5fb57816020015b606081526020019060019003908161b5e65790505b50905060005b8260ff168160ff16101561b65457838160ff168151811061b6245761b62461f124565b6020026020010151828260ff168151811061b6415761b64161f124565b602090810291909101015260010161b601565b5093505050505b949350505050565b61b68a6040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c9161b7109186910161f31b565b600060405180830381865afa15801561b72d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261b755919081019061edbe565b9050600061b763868361cda0565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b815260040161b793919061e0fc565b6000604051808303816000875af115801561b7b2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261b7da919081019061f362565b805190915060030b1580159061b7f35750602081015151155b801561b8025750604081015151155b1561a7dc578160008151811061b81a5761b81a61f124565b6020026020010151604051602001619e35919061f418565b6060600061b8678560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252865182528087019082015290915061b89e9082905b9061cef5565b1561b9fb57600061b91b8261b9158461b90f61b8e18a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061cf1c565b9061cf7e565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061b97f90829061cef5565b1561b9e957604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b9e6905b829061d003565b90505b61b9f28161d029565b92505050619d62565b821561ba14578484604051602001619e3592919061f604565b5050604080516020810190915260008152619d62565b509392505050565b6000808251602084016000f09392505050565b8160a001511561ba5457505050565b600061ba6184848461d092565b9050600061ba6e8261b663565b602081015181519192509060030b15801561bb0a5750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bb0a9060408051808201825260008082526020918201528151808301909252845182528085019082015261b898565b1561bb1757505050505050565b6040820151511561bb37578160400151604051602001619e35919061f6ab565b80604051602001619e35919061f709565b6060600061bb7d8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061bbe2905b829061c1b0565b1561bc5157604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619d629061bc4c90839061d62d565b61d029565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bcb3905b829061d6b7565b60010361bd8057604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bd199061b9df565b50604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619d629061bc4c905b839061d003565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bddf9061bbdb565b1561bf1657604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061be4790839061d751565b90506000816001835161be5a919061e660565b8151811061be6a5761be6a61f124565b6020026020010151905061bf0d61bc4c61bee06040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925285518252808601908201529061d62d565b95945050505050565b82604051602001619e35919061f774565b50919050565b6060600061bf628360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061bfc49061bbdb565b1561bfd257619d628161d029565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c0319061bcac565b60010361c09b57604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619d629061bc4c9061bd79565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c0fa9061bbdb565b1561bf1657604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061c16290839061d751565b905060018151111561c19e57806002825161c17d919061e660565b8151811061c18d5761c18d61f124565b602002602001015192505050919050565b5082604051602001619e35919061f774565b80518251600091111561c1c557506000619b0e565b8151835160208501516000929161c1db9161e673565b61c1e5919061e660565b90508260200151810361c1fc576001915050619b0e565b82516020840151819020912014905092915050565b6060600061c21e8361d7f6565b600101905060008167ffffffffffffffff81111561c23e5761c23e61e4fc565b6040519080825280601f01601f19166020018201604052801561c268576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461c27257509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e534544000000000000000000000000000000000000000000008184019081528551808701875283815284019290925284518086019095525184529083015260609161c33d905b829061d8d8565b1561c37d57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c3dc9061c336565b1561c41c57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c47b9061c336565b1561c4bb57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c51a9061c336565b8061c57f5750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c57f9061c336565b1561c5bf57505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c61e9061c336565b8061c6835750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c6839061c336565b1561c6c357505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c7229061c336565b8061c7875750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c7879061c336565b1561c7c757505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c8269061c336565b8061c88b5750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c88b9061c336565b1561c8cb57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c92a9061c336565b1561c96a57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c9c99061c336565b1561ca0957505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ca689061c336565b1561caa857505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261cb079061c336565b1561cb4757505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261cba69061c336565b1561cbe657505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261cc459061c336565b8061ccaa5750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ccaa9061c336565b1561ccea57505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261cd499061c336565b1561cd8957505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151619e35929060200161f852565b60608060005b845181101561ce2b578185828151811061cdc25761cdc261f124565b602002602001015160405160200161cddb92919061ec9d565b60405160208183030381529060405291506001855161cdfa919061e660565b811461ce23578160405160200161ce11919061f9bb565b60405160208183030381529060405291505b60010161cda6565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161ce44579050509050838160008151811061ce6f5761ce6f61f124565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061cec35761cec361f124565b6020026020010181905250818160028151811061cee25761cee261f124565b6020908102919091010152949350505050565b602080830151835183519284015160009361cf13929184919061d8ec565b14159392505050565b6040805180820190915260008082526020820152600061cf4e846000015185602001518560000151866020015161d9fd565b905083602001518161cf60919061e660565b8451859061cf6f90839061e660565b90525060208401525090919050565b604080518082019091526000808252602082015281518351101561cfa3575081619b0e565b602080830151908401516001911461cfca5750815160208481015190840151829020919020145b801561cffb5782518451859061cfe190839061e660565b905250825160208501805161cff790839061e673565b9052505b509192915050565b604080518082019091526000808252602082015261d02283838361db1d565b5092915050565b60606000826000015167ffffffffffffffff81111561d04a5761d04a61e4fc565b6040519080825280601f01601f19166020018201604052801561d074576020820181803683370190505b509050600060208201905061d022818560200151866000015161dbc8565b6060600061d09e61a05b565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161d0bb57905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061d1169061f291565b935060ff168151811061d12b5761d12b61f124565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161d17c919061f9fc565b60405160208183030381529060405282828061d1979061f291565b935060ff168151811061d1ac5761d1ac61f124565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061d1f99061f291565b935060ff168151811061d20e5761d20e61f124565b60200260200101819052508260405160200161d22a919061f1bf565b60405160208183030381529060405282828061d2459061f291565b935060ff168151811061d25a5761d25a61f124565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061d2a79061f291565b935060ff168151811061d2bc5761d2bc61f124565b602002602001018190525061d2d1878461dc42565b828261d2dc8161f291565b935060ff168151811061d2f15761d2f161f124565b60209081029190910101528551511561d39d5760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261d3438161f291565b935060ff168151811061d3585761d35861f124565b602002602001018190525061d37186600001518461dc42565b828261d37c8161f291565b935060ff168151811061d3915761d39161f124565b60200260200101819052505b85608001511561d40b5760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261d3e68161f291565b935060ff168151811061d3fb5761d3fb61f124565b602002602001018190525061d471565b841561d4715760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261d4508161f291565b935060ff168151811061d4655761d46561f124565b60200260200101819052505b6040860151511561d50d5760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261d4bb8161f291565b935060ff168151811061d4d05761d4d061f124565b6020026020010181905250856040015182828061d4ec9061f291565b935060ff168151811061d5015761d50161f124565b60200260200101819052505b85606001511561d5775760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261d5568161f291565b935060ff168151811061d56b5761d56b61f124565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561d5955761d59561e4fc565b60405190808252806020026020018201604052801561d5c857816020015b606081526020019060019003908161d5b35790505b50905060005b8260ff168160ff16101561d62157838160ff168151811061d5f15761d5f161f124565b6020026020010151828260ff168151811061d60e5761d60e61f124565b602090810291909101015260010161d5ce565b50979650505050505050565b604080518082019091526000808252602082015281518351101561d652575081619b0e565b8151835160208501516000929161d6689161e673565b61d672919061e660565b6020840151909150600190821461d693575082516020840151819020908220145b801561d6ae5783518551869061d6aa90839061e660565b9052505b50929392505050565b600080826000015161d6db856000015186602001518660000151876020015161d9fd565b61d6e5919061e673565b90505b8351602085015161d6f9919061e673565b811161d022578161d7098161fa41565b925050826000015161d74085602001518361d724919061e660565b865161d730919061e660565b838660000151876020015161d9fd565b61d74a919061e673565b905061d6e8565b6060600061d75f848461d6b7565b61d76a90600161e673565b67ffffffffffffffff81111561d7825761d78261e4fc565b60405190808252806020026020018201604052801561d7b557816020015b606081526020019060019003908161d7a05790505b50905060005b815181101561ba2a5761d7d161bc4c868661d003565b82828151811061d7e35761d7e361f124565b602090810291909101015260010161d7bb565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061d83f577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061d86b576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061d88957662386f26fc10000830492506010015b6305f5e100831061d8a1576305f5e100830492506008015b612710831061d8b557612710830492506004015b6064831061d8c7576064830492506002015b600a8310619b0e5760010192915050565b600061d8e4838361dc82565b159392505050565b60008085841161d9f3576020841161d99f576000841561d93757600161d91386602061e660565b61d91e90600861fa5b565b61d92990600261fb59565b61d933919061e660565b1990505b835181168561d946898961e673565b61d950919061e660565b805190935082165b81811461d98a5787841161d972578794505050505061b65b565b8361d97c8161fb65565b94505082845116905061d958565b61d994878561e673565b94505050505061b65b565b83832061d9ac858861e660565b61d9b6908761e673565b91505b85821061d9f15784822080820361d9de5761d9d4868461e673565b935050505061b65b565b61d9e960018461e660565b92505061d9b9565b505b5092949350505050565b6000838186851161db08576020851161dab7576000851561da4957600161da2587602061e660565b61da3090600861fa5b565b61da3b90600261fb59565b61da45919061e660565b1990505b8451811660008761da5a8b8b61e673565b61da64919061e660565b855190915083165b82811461daa95781861061da915761da848b8b61e673565b965050505050505061b65b565b8561da9b8161fa41565b96505083865116905061da6c565b85965050505050505061b65b565b508383206000905b61dac9868961e660565b821161db065785832080820361dae5578394505050505061b65b565b61daf060018561e673565b935050818061dafe9061fa41565b92505061dabf565b505b61db12878761e673565b979650505050505050565b6040805180820190915260008082526020820152600061db4f856000015186602001518660000151876020015161d9fd565b60208087018051918601919091525190915061db6b908261e660565b83528451602086015161db7e919061e673565b810361db8d576000855261dbbf565b8351835161db9b919061e673565b8551869061dbaa90839061e660565b905250835161dbb9908261e673565b60208601525b50909392505050565b6020811061dc00578151835261dbdf60208461e673565b925061dbec60208361e673565b915061dbf960208261e660565b905061dbc8565b600019811561dc2f57600161dc1683602061e660565b61dc229061010061fb59565b61dc2c919061e660565b90505b9151835183169219169190911790915250565b6060600061dc50848461a12e565b805160208083015160405193945061dc6a9390910161fb7c565b60405160208183030381529060405291505092915050565b815181516000919081111561dc95575081515b6020808501519084015160005b8381101561dd4e578251825180821461dd1e57600019602087101561dcfd5760018461dccf89602061e660565b61dcd9919061e673565b61dce490600861fa5b565b61dcef90600261fb59565b61dcf9919061e660565b1990505b818116838216818103911461dd1b579750619b0e9650505050505050565b50505b61dd2960208661e673565b945061dd3660208561e673565b9350505060208161dd47919061e673565b905061dca2565b508451865161a7dc919061fbd4565b610b67806200fbf583390190565b6151c8806201075c83390190565b610a42806201592483390190565b61106f806201636683390190565b61207280620173d583390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161dde661ddeb565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161dde66040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561de9d5783516001600160a01b031683526020938401939092019160010161de76565b509095945050505050565b60005b8381101561dec357818101518382015260200161deab565b50506000910152565b6000815180845261dee481602086016020860161dea8565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561dff4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561dfda577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261dfc484865161decc565b602095860195909450929092019160010161df8a565b50919750505060209485019492909201915060010161df20565b50929695505050505050565b600081518084526020840193506020830160005b8281101561e0545781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161e014565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561dff4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261e0ca604088018261decc565b905060208201519150868103602088015261e0e5818361e000565b96505050602093840193919091019060010161e086565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561dff4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261e15e85835161decc565b9450602093840193919091019060010161e124565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561dff4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261e1f4604087018261e000565b955050602093840193919091019060010161e19b565b600181811c9082168061e21e57607f821691505b60208210810361bf27577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000815461e2648161e20a565b80855260018216801561e27e576001811461e2b85761e2ef565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083166020870152602082151560051b870101935061e2ef565b84600052602060002060005b8381101561e2e65781546020828a01015260018201915060208101905061e2c4565b87016020019450505b50505092915050565b6001600160a01b0381541682526001600160a01b03600182015416602083015260028101546040830152608060608301526000619d62608084016003840161e257565b602081526000619d62602083018461e2f8565b6001600160a01b038316815260406020820152600061b65b604083018461e2f8565b6000825161e38281846020870161dea8565b9190910192915050565b6001600160a01b038316815260406020820152600061b65b604083018461decc565b80516001600160a01b038116811461e3c557600080fd5b919050565b60006020828403121561e3dc57600080fd5b619d628261e3ae565b6003811061e41c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff8816604083015286606083015261e4b0608083018761e3e5565b8460a08301526001600160a01b03841660c083015261db1260e08301846001600160a01b03169052565b60006020828403121561e4ec57600080fd5b81518015158114619d6257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f8211156127cf57806000526020600020601f840160051c8101602085101561e5525750805b601f840160051c820191505b81811015611e7f576000815560010161e55e565b815167ffffffffffffffff81111561e58c5761e58c61e4fc565b61e5a08161e59a845461e20a565b8461e52b565b6020601f82116001811461e5d4576000831561e5bc5750848201515b600019600385901b1c1916600184901b178455611e7f565b600084815260208120601f198516915b8281101561e604578785015182556020948501946001909201910161e5e4565b508482101561e6225786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810381811115619b0e57619b0e61e631565b80820180821115619b0e57619b0e61e631565b60c08252600061e69960c084018361e257565b6001600160a01b0360018401541660208501526002830154604085015260ff600384015416151560608501526004830154608085015283810360a0850152619b0a816005850161e257565b602081526000619d62602083018461e686565b6001600160a01b038316815260406020820152600061b65b604083018461e686565b602081526000619d62602083018461decc565b60006020828403121561e73e57600080fd5b5051919050565b602081526000619b0e60208301600581527f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260400190565b60a08152600061e79560a083018761decc565b6001600160a01b03861660208401528460408401526001600160a01b0384166060840152828103608084015261db1281600581527f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260400190565b600081516060845261e80a606085018261decc565b90506001600160a01b036020840151166020850152604083015160408501528091505092915050565b60608152600061e846606083018661e7f5565b6001600160a01b0385166020840152828103604084015261a7dc818561decc565b60a08152600061e87a60a083018861e7f5565b6001600160a01b03871660208401528560408401526001600160a01b0385166060840152828103608084015261e8b0818561decc565b98975050505050505050565b6000806040838503121561e8cf57600080fd5b61e8d88361e3ae565b6020939093015192949293505050565b600081546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b03600183015416604084015260a0606084015261e93160a084016002840161e257565b600383015460808501528091505092915050565b60e08152600061e95860e083018961decc565b8760208401526001600160a01b0387166040840152828103606084015261e97f818761decc565b855160808501526020860151151560a0850152905082810360c084015261e9a6818561e8e8565b9998505050505050505050565b6001600160a01b03851681528360208201526001600160a01b038316604082015260806060820152600061a7dc608083018461e2f8565b60c08152600061e9fd60c083018861decc565b6001600160a01b0387166020840152828103604084015261ea1e818761decc565b85516060850152602086015115156080850152905082810360a084015261e8b0818561e8e8565b610100815260066101008201527f5345434f4e440000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f534543000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff8816604083015286606083015261e4b0608083018761e3e5565b60808152600061eae8608083018761decc565b8560208401526001600160a01b0385166040840152828103606084015261db12818561e8e8565b83815282602082015260606040820152600061bf0d606083018461decc565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161eb6681601a85016020880161dea8565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161eba381601c84016020880161dea8565b01601c01949350505050565b6040516060810167ffffffffffffffff8111828210171561ebd25761ebd261e4fc565b60405290565b60008067ffffffffffffffff84111561ebf35761ebf361e4fc565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561ec225761ec2261e4fc565b60405283815290508082840185101561ec3a57600080fd5b61ba2a84602083018561dea8565b600082601f83011261ec5957600080fd5b619d628383516020850161ebd8565b60006020828403121561ec7a57600080fd5b815167ffffffffffffffff81111561ec9157600080fd5b619b0a8482850161ec48565b6000835161ecaf81846020880161dea8565b83519083019061ecc381836020880161dea8565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161ed0481601a85016020880161dea8565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161ed4181603384016020880161dea8565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000619d62608083018461decc565b60006020828403121561edd057600080fd5b815167ffffffffffffffff81111561ede757600080fd5b8201601f8101841361edf857600080fd5b619b0a8482516020840161ebd8565b6000855161ee19818460208a0161dea8565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161ee53816001840160208a0161dea8565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161ee9181600284016020890161dea8565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161eed381600284016020880161dea8565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061ef1e604083018461decc565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161ef9581601f85016020870161dea8565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061f002604083018461decc565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061f054604083018461decc565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161f0cb81601485016020870161dea8565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061f112604083018561decc565b8281036020840152619d5e818561decc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161f18b81600185016020870161dea8565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161f1d181846020870161dea8565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161f28481604b85016020870161dea8565b91909101604b0192915050565b600060ff821660ff810361f2a75761f2a761e631565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161f30e81602985016020870161dea8565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000619d62608083018461decc565b60006020828403121561f37457600080fd5b815167ffffffffffffffff81111561f38b57600080fd5b82016060818503121561f39d57600080fd5b61f3a561ebaf565b81518060030b811461f3b657600080fd5b8152602082015167ffffffffffffffff81111561f3d257600080fd5b61f3de8682850161ec48565b602083015250604082015167ffffffffffffffff81111561f3fe57600080fd5b61f40a8682850161ec48565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161f47681602185016020870161dea8565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161f66281602185016020880161dea8565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161f69f81602e84016020880161dea8565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161f30e81602985016020870161dea8565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161f76781602285016020870161dea8565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161f7ac81600e85016020870161dea8565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161f88a81601885016020880161dea8565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161f8c781601c84016020880161dea8565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161f9cd81846020870161dea8565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161fa3481601c85016020870161dea8565b91909101601c0192915050565b6000600019820361fa545761fa5461e631565b5060010190565b8082028115828204841417619b0e57619b0e61e631565b6001815b600184111561faad5780850481111561fa915761fa9161e631565b600184161561fa9f57908102905b60019390931c92800261fa76565b935093915050565b60008261fac457506001619b0e565b8161fad157506000619b0e565b816001811461fae7576002811461faf15761fb0d565b6001915050619b0e565b60ff84111561fb025761fb0261e631565b50506001821b619b0e565b5060208310610133831016604e8410600b841016171561fb30575081810a619b0e565b61fb3d600019848461fa72565b806000190482111561fb515761fb5161e631565b029392505050565b6000619d62838361fab5565b60008161fb745761fb7461e631565b506000190190565b6000835161fb8e81846020880161dea8565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161fbc881600184016020880161dea8565b01600101949350505050565b818103600083128015838313168383128216171561d0225761d02261e63156fe60c0604052600d60809081526c2bb930b83832b21022ba3432b960991b60a05260009061002c9082610114565b506040805180820190915260048152630ae8aa8960e31b60208201526001906100559082610114565b506002805460ff1916601217905534801561006f57600080fd5b506101d2565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061009f57607f821691505b6020821081036100bf57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561010f57806000526020600020601f840160051c810160208510156100ec5750805b601f840160051c820191505b8181101561010c57600081556001016100f8565b50505b505050565b81516001600160401b0381111561012d5761012d610075565b6101418161013b845461008b565b846100c5565b6020601f821160018114610175576000831561015d5750848201515b600019600385901b1c1916600184901b17845561010c565b600084815260208120601f198516915b828110156101a55787850151825560209485019460019092019101610185565b50848210156101c35786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610986806101e16000396000f3fe6080604052600436106100c05760003560e01c8063313ce56711610074578063a9059cbb1161004e578063a9059cbb146101fa578063d0e30db01461021a578063dd62ed3e1461022257600080fd5b8063313ce5671461018c57806370a08231146101b857806395d89b41146101e557600080fd5b806318160ddd116100a557806318160ddd1461012f57806323b872dd1461014c5780632e1a7d4d1461016c57600080fd5b806306fdde03146100d4578063095ea7b3146100ff57600080fd5b366100cf576100cd61025a565b005b600080fd5b3480156100e057600080fd5b506100e96102b5565b6040516100f69190610745565b60405180910390f35b34801561010b57600080fd5b5061011f61011a3660046107da565b610343565b60405190151581526020016100f6565b34801561013b57600080fd5b50475b6040519081526020016100f6565b34801561015857600080fd5b5061011f610167366004610804565b6103bd565b34801561017857600080fd5b506100cd610187366004610841565b610647565b34801561019857600080fd5b506002546101a69060ff1681565b60405160ff90911681526020016100f6565b3480156101c457600080fd5b5061013e6101d336600461085a565b60036020526000908152604090205481565b3480156101f157600080fd5b506100e9610724565b34801561020657600080fd5b5061011f6102153660046107da565b610731565b6100cd61025a565b34801561022e57600080fd5b5061013e61023d366004610875565b600460209081526000928352604080842090915290825290205481565b33600090815260036020526040812080543492906102799084906108d7565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b600080546102c2906108ea565b80601f01602080910402602001604051908101604052809291908181526020018280546102ee906108ea565b801561033b5780601f106103105761010080835404028352916020019161033b565b820191906000526020600020905b81548152906001019060200180831161031e57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103ab9086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604081205482111561042b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841633148015906104a1575073ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156105605773ffffffffffffffffffffffffffffffffffffffff8416600090815260046020908152604080832033845290915290205482111561051a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b73ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091528120805484929061055a90849061093d565b90915550505b73ffffffffffffffffffffffffffffffffffffffff84166000908152600360205260408120805484929061059590849061093d565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040812080548492906105cf9084906108d7565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161063591815260200190565b60405180910390a35060019392505050565b3360009081526003602052604090205481111561069a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b33600090815260036020526040812080548392906106b990849061093d565b9091555050604051339082156108fc029083906000818181858888f193505050501580156106eb573d6000803e3d6000fd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b600180546102c2906108ea565b600061073e3384846103bd565b9392505050565b602081526000825180602084015260005b818110156107735760208186018101516040868401015201610756565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146107d557600080fd5b919050565b600080604083850312156107ed57600080fd5b6107f6836107b1565b946020939093013593505050565b60008060006060848603121561081957600080fd5b610822846107b1565b9250610830602085016107b1565b929592945050506040919091013590565b60006020828403121561085357600080fd5b5035919050565b60006020828403121561086c57600080fd5b61073e826107b1565b6000806040838503121561088857600080fd5b610891836107b1565b915061089f602084016107b1565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156103b7576103b76108a8565b600181811c908216806108fe57607f821691505b602082108103610937577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b818103818111156103b7576103b76108a856fea2646970667358221220bb5b9dcef0ba90bcdefcbd63f71b1df95b50e29550a7456c69c6b9ff9dcdd20e64736f6c634300081a003360a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516150cb6100fd60003960008181612898015281816128c10152612a7401526150cb6000f3fe6080604052600436106102a05760003560e01c80638f2839701161016e578063bd8fde1c116100cb578063d547741f1161007f578063e9d6c5ba11610064578063e9d6c5ba1461082d578063f354b31f1461085f578063f851a4401461087f57600080fd5b8063d547741f146107d9578063e63ab1e9146107f957600080fd5b8063c1bd469f116100b0578063c1bd469f14610777578063cc5ad8b614610799578063d3523ea2146107b957600080fd5b8063bd8fde1c14610723578063c0c53b8b1461075757600080fd5b8063a217fddf11610122578063a8f2cb9611610107578063a8f2cb961461069a578063aa808c06146106ba578063ad3cb1cc146106da57600080fd5b8063a217fddf1461066e578063a3ebd14c1461068357600080fd5b806391d148541161015357806391d14854146105c557806394cc86831461062a5780639ca220dd1461064c57600080fd5b80638f283970146105855780639060bda9146105a557600080fd5b80633f4ba83a1161021c578063631d62e4116101d05780637066b18d116101b55780637066b18d14610515578063804ea334146105425780638456cb591461057057600080fd5b8063631d62e4146104d55780636e9e2d3f146104f557600080fd5b806352d1902d1161020157806352d1902d1461045b5780635c975abb146104705780635cf92c9f146104a757600080fd5b80633f4ba83a146104335780634f1ef2861461044857600080fd5b80632259e9e5116102735780632f2ff15d116102585780632f2ff15d146103d35780633500c24b146103f357806336568abe1461041357600080fd5b80632259e9e514610356578063248a9ca31461037657600080fd5b806301ffc9a7146102a55780630c63109e146102da57806310d29b9e1461031257806318d3ce9614610334575b600080fd5b3480156102b157600080fd5b506102c56102c0366004613fd6565b61089f565b60405190151581526020015b60405180910390f35b3480156102e657600080fd5b506001546102fa906001600160a01b031681565b6040516001600160a01b0390911681526020016102d1565b34801561031e57600080fd5b5061033261032d36600461406f565b610938565b005b34801561034057600080fd5b506103496109f3565b6040516102d1919061411e565b34801561036257600080fd5b506103326103713660046141df565b610c95565b34801561038257600080fd5b506103c561039136600461425e565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016102d1565b3480156103df57600080fd5b506103326103ee36600461428c565b610d27565b3480156103ff57600080fd5b5061033261040e3660046142bc565b610d71565b34801561041f57600080fd5b5061033261042e36600461428c565b610f04565b34801561043f57600080fd5b50610332610f55565b610332610456366004614308565b610f6b565b34801561046757600080fd5b506103c5610f8a565b34801561047c57600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166102c5565b3480156104b357600080fd5b506104c76104c23660046143d5565b610fb9565b6040516102d1929190614421565b3480156104e157600080fd5b506103326104f03660046141df565b6110b4565b34801561050157600080fd5b50610332610510366004614444565b61115a565b34801561052157600080fd5b506105356105303660046143d5565b611219565b6040516102d1919061451f565b34801561054e57600080fd5b5061056261055d36600461425e565b6112e2565b6040516102d1929190614532565b34801561057c57600080fd5b5061033261139a565b34801561059157600080fd5b506103326105a03660046142bc565b6113cc565b3480156105b157600080fd5b506103326105c0366004614554565b611520565b3480156105d157600080fd5b506102c56105e036600461428c565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561063657600080fd5b5061063f6115ae565b6040516102d19190614582565b34801561065857600080fd5b50610661611606565b6040516102d191906145c5565b34801561067a57600080fd5b506103c5600081565b34801561068f57600080fd5b506103c56207a12081565b3480156106a657600080fd5b506103326106b5366004614672565b6117c5565b3480156106c657600080fd5b506102fa6106d53660046143d5565b611845565b3480156106e657600080fd5b506105356040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b34801561072f57600080fd5b506103c57ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa381565b34801561076357600080fd5b506103326107723660046146e6565b611885565b34801561078357600080fd5b5061078c611c32565b6040516102d19190614731565b3480156107a557600080fd5b50600b546102fa906001600160a01b031681565b3480156107c557600080fd5b506105356107d43660046141df565b611f39565b3480156107e557600080fd5b506103326107f436600461428c565b612021565b34801561080557600080fd5b506103c57f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b34801561083957600080fd5b5061084d6108483660046142bc565b612065565b6040516102d196959493929190614828565b34801561086b57600080fd5b5061033261087a366004614886565b6122b7565b34801561088b57600080fd5b506000546102fa906001600160a01b031681565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061093257507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361096281612353565b61096a61235d565b610976858585856123bb565b6109828585858561250f565b7f6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c6007600087815260200190815260200160002085856040516109c6929190614936565b90815260200160405180910390206001016040516109e49190614a1c565b60405180910390a15050505050565b6004546060908067ffffffffffffffff811115610a1257610a126142d9565b604051908082528060200260200182016040528015610a7057816020015b610a5d60405180608001604052806000151581526020016060815260200160608152602001600081525090565b815260200190600190039081610a305790505b50915060005b81811015610c9057600060048281548110610a9357610a93614a2f565b906000526020600020906002020160405180604001604052908160008201548152602001600182018054610ac690614946565b80601f0160208091040260200160405190810160405280929190818152602001828054610af290614946565b8015610b3f5780601f10610b1457610100808354040283529160200191610b3f565b820191906000526020600020905b815481529060010190602001808311610b2257829003601f168201915b505050505081525050905060008160000151905060008260200151905060405180608001604052806007600085815260200190815260200160002083604051610b889190614a5e565b90815260408051602092819003830190205460ff161515835260008681526007835281902090519290910191610bbf908590614a5e565b90815260200160405180910390206001018054610bdb90614946565b80601f0160208091040260200160405190810160405280929190818152602001828054610c0790614946565b8015610c545780601f10610c2957610100808354040283529160200191610c54565b820191906000526020600020905b815481529060010190602001808311610c3757829003601f168201915b5050505050815260200182815260200183815250868581518110610c7a57610c7a614a2f565b6020908102919091010152505050600101610a76565b505090565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3610cbf81612353565b610cc761235d565b610cd48686868686612592565b610ce18686868686612628565b857f40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee63486868686604051610d179493929190614aa5565b60405180910390a2505050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610d6181612353565b610d6b83836126a6565b50505050565b6000610d7c81612353565b6001600160a01b038216610dbc576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610de67ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3836126a6565b50610e117f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a836126a6565b50600154610e49907ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3906001600160a01b0316612775565b50600154610e81907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b0316612775565b50600154604080516001600160a01b03928316815291841660208301527f6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279910160405180910390a150600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6001600160a01b0381163314610f46576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f508282612775565b505050565b6000610f6081612353565b610f6861281b565b50565b610f7361288d565b610f7c8261295d565b610f868282612968565b5050565b6000610f94612a69565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b600083815260076020526040808220905160609190610fdb9086908690614936565b908152604080519182900360209081018320546000898152600790925291902060ff9091169350906110109086908690614936565b9081526020016040518091039020600101805461102c90614946565b80601f016020809104026020016040519081016040528092919081815260200182805461105890614946565b80156110a55780601f1061107a576101008083540402835291602001916110a5565b820191906000526020600020905b81548152906001019060200180831161108857829003601f168201915b50505050509050935093915050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36110de81612353565b6110e661235d565b6110f38686868686612acb565b6111008686868686612daf565b8484604051611110929190614936565b6040518091039020867f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581858560405161114a929190614ad7565b60405180910390a3505050505050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361118481612353565b61118c61235d565b61119d8a8a8a8a8a8a8a8a8a612e2d565b6111ae8a8a8a8a8a8a8a8a8a613163565b896001600160a01b031686866040516111c8929190614936565b60405180910390207fa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367848a8d8d6040516112059493929190614aeb565b60405180910390a350505050505050505050565b6060600660008581526020019081526020016000206004018383604051611241929190614936565b9081526020016040518091039020805461125a90614946565b80601f016020809104026020016040519081016040528092919081815260200182805461128690614946565b80156112d35780601f106112a8576101008083540402835291602001916112d3565b820191906000526020600020905b8154815290600101906020018083116112b657829003601f168201915b505050505090505b9392505050565b60008181526006602052604090206002810154600390910180546001600160a01b03909216916060919061131590614946565b80601f016020809104026020016040519081016040528092919081815260200182805461134190614946565b801561138e5780601f106113635761010080835404028352916020019161138e565b820191906000526020600020905b81548152906001019060200180831161137157829003601f168201915b50505050509050915091565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6113c481612353565b610f686131f5565b60006113d781612353565b6001600160a01b038216611417576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114226000836126a6565b5061144d7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a836126a6565b506000805461146591906001600160a01b0316612775565b5060005461149d907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b0316612775565b50600054604080516001600160a01b03928316815291841660208301527f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f910160405180910390a150600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361154a81612353565b61155261235d565b61155c8383613250565b6115668383613340565b604080516001600160a01b038516815283151560208201527f9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8910160405180910390a1505050565b606060028054806020026020016040519081016040528092919081815260200182805480156115fc57602002820191906000526020600020905b8154815260200190600101908083116115e8575b5050505050905090565b6003546060908067ffffffffffffffff811115611625576116256142d9565b60405190808252806020026020018201604052801561169457816020015b60408051608081018252600080825260208083018290529282015260608082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092019101816116435790505b50915060005b81811015610c90576000600382815481106116b7576116b7614a2f565b6000918252602080832090910154604080516080810182528285526006808552828620805460ff161515835282860185905260028101546001600160a01b031693830193909352948390529390925260039091018054919350606083019161171e90614946565b80601f016020809104026020016040519081016040528092919081815260200182805461174a90614946565b80156117975780601f1061176c57610100808354040283529160200191611797565b820191906000526020600020905b81548152906001019060200180831161177a57829003601f168201915b50505050508152508483815181106117b1576117b1614a2f565b60209081029190910101525060010161169a565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36117ef81612353565b6117f761235d565b61180486868686866133c3565b6118118686868686613579565b857fc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e83604051610d17911515815260200190565b6000838152600a602052604080822090516118639085908590614936565b908152604051908190036020019020546001600160a01b031690509392505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156118d05750825b905060008267ffffffffffffffff1660011480156118ed5750303b155b9050811580156118fb575080155b15611932576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156119935784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b03881615806119b057506001600160a01b038716155b806119c257506001600160a01b038616155b156119f9576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611a016135f7565b611a096135f7565b611a116135ff565b611a1c6000896126a6565b50611a477ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3886126a6565b50611a727f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a886126a6565b50611a9d7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a896126a6565b50600080546001600160a01b038a81167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316178355600180548b8316908416178155600b8054928b16929093169190911790915546808352600660208181526040808620805460ff1916909517855580513060601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016818401528151808203601401815260349091019091529290945290925260030190611b619082614b5f565b50600280546001818101909255467f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9091018190556003805492830181556000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b909101558315611c285784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b6005546060908067ffffffffffffffff811115611c5157611c516142d9565b604051908082528060200260200182016040528015611cd057816020015b611cbd6040518060e0016040528060001515815260200160006001600160a01b0316815260200160608152602001600081526020016060815260200160608152602001600060ff1681525090565b815260200190600190039081611c6f5790505b50915060005b81811015610c9057600060058281548110611cf357611cf3614a2f565b60009182526020808320909101546001600160a01b0390811680845260088352604093849020845160e081018652815460ff811615158252610100900490931693830193909352600183018054919550919384019190611d5290614946565b80601f0160208091040260200160405190810160405280929190818152602001828054611d7e90614946565b8015611dcb5780601f10611da057610100808354040283529160200191611dcb565b820191906000526020600020905b815481529060010190602001808311611dae57829003601f168201915b5050505050815260200160028201548152602001600382018054611dee90614946565b80601f0160208091040260200160405190810160405280929190818152602001828054611e1a90614946565b8015611e675780601f10611e3c57610100808354040283529160200191611e67565b820191906000526020600020905b815481529060010190602001808311611e4a57829003601f168201915b50505050508152602001600482018054611e8090614946565b80601f0160208091040260200160405190810160405280929190818152602001828054611eac90614946565b8015611ef95780601f10611ece57610100808354040283529160200191611ef9565b820191906000526020600020905b815481529060010190602001808311611edc57829003601f168201915b50505091835250506005919091015460ff166020909101528451859084908110611f2557611f25614a2f565b602090810291909101015250600101611cd6565b6060600760008781526020019081526020016000208585604051611f5e929190614936565b90815260200160405180910390206003018383604051611f7f929190614936565b90815260200160405180910390208054611f9890614946565b80601f0160208091040260200160405190810160405280929190818152602001828054611fc490614946565b80156120115780601f10611fe657610100808354040283529160200191612011565b820191906000526020600020905b815481529060010190602001808311611ff457829003601f168201915b5050505050905095945050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461205b81612353565b610d6b8383612775565b6001600160a01b038082166000908152600860209081526040808320815160e081018352815460ff811615158252610100900490951692850192909252600182018054939460609486948694859487948594909392840191906120c790614946565b80601f01602080910402602001604051908101604052809291908181526020018280546120f390614946565b80156121405780601f1061211557610100808354040283529160200191612140565b820191906000526020600020905b81548152906001019060200180831161212357829003601f168201915b505050505081526020016002820154815260200160038201805461216390614946565b80601f016020809104026020016040519081016040528092919081815260200182805461218f90614946565b80156121dc5780601f106121b1576101008083540402835291602001916121dc565b820191906000526020600020905b8154815290600101906020018083116121bf57829003601f168201915b505050505081526020016004820180546121f590614946565b80601f016020809104026020016040519081016040528092919081815260200182805461222190614946565b801561226e5780601f106122435761010080835404028352916020019161226e565b820191906000526020600020905b81548152906001019060200180831161225157829003601f168201915b50505091835250506005919091015460ff16602090910152805160808201516060830151604084015160a085015160c090950151939d929c50909a509850919650945092505050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36122e181612353565b6122e961235d565b6122f888888888888888613632565b6123078888888888888861378a565b877faea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c588888888888860405161234196959493929190614c5a565b60405180910390a25050505050505050565b610f68813361380c565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16156123b9576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60008481526006602052604090205460ff1661240b576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018590526024015b60405180910390fd5b60008290036124495782826040517ec10cfd000000000000000000000000000000000000000000000000000000008152600401612402929190614ad7565b6000848152600760205260409081902090516124689085908590614936565b9081526020016040518091039020600101805461248490614946565b90506000036124c5578383836040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161240293929190614ca3565b806007600086815260200190815260200160002084846040516124e9929190614936565b908152604051908190036020019020805491151560ff1990921691909117905550505050565b6000848484846040516024016125289493929190614cbd565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f10d29b9e00000000000000000000000000000000000000000000000000000000179052905061258b81613899565b5050505050565b60008581526006602052604090205460ff166125dd576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612402565b8181600660008881526020019081526020016000206004018686604051612605929190614936565b90815260200160405180910390209182612620929190614cea565b505050505050565b60008585858585604051602401612643959493929190614de6565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f2259e9e500000000000000000000000000000000000000000000000000000000179052905061262081613899565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1661276b576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556127213390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610932565b6000915050610932565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff161561276b576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610932565b612823613941565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061292657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661291a7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b156123b9576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610f8681612353565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156129c2575060408051601f3d908101601f191682019092526129bf91810190614e1f565b60015b612a03576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401612402565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612a5f576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401612402565b610f50838361399c565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146123b9576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008581526006602052604090205460ff16612b16576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612402565b6000839003612b545783836040517ec10cfd000000000000000000000000000000000000000000000000000000008152600401612402929190614ad7565b6000819003612b8f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000858152600760205260408082209051612bad9087908790614936565b90815260200160405180910390206001018054612bc990614946565b90501115612c0d5784848484846040517f2b192eab000000000000000000000000000000000000000000000000000000008152600401612402959493929190614de6565b6001600760008781526020019081526020016000208585604051612c32929190614936565b9081526040805160209281900383018120805460ff1916941515949094179093556000888152600790925290208391839190612c719088908890614936565b90815260200160405180910390206001019182612c8f929190614cea565b508383600760008881526020019081526020016000208686604051612cb5929190614936565b90815260200160405180910390206002019182612cd3929190614cea565b506004604051806040016040528087815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939094525050835460018181018655948252602091829020845160029092020190815590830151929390929083019150612d539082614b5f565b5050508383604051612d66929190614936565b6040518091039020857f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce25818484604051612da0929190614ad7565b60405180910390a35050505050565b60008585858585604051602401612dca959493929190614de6565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f631d62e400000000000000000000000000000000000000000000000000000000179052905061262081613899565b6001600160a01b038916612e6d576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000879003612ed7576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f53796d626f6c2063616e6e6f7420626520656d707479000000000000000000006044820152606401612402565b6001600160a01b0389811660009081526008602052604090205461010090041615612f39576040517f63f4ee1f0000000000000000000000000000000000000000000000000000000081526001600160a01b038a166004820152602401612402565b60006001600160a01b031660098989604051612f56929190614936565b908152604051908190036020019020546001600160a01b031614612faa5787876040517fb295cac9000000000000000000000000000000000000000000000000000000008152600401612402929190614ad7565b6001600160a01b0389166000818152600860205260409020805460017fffffffffffffffffffffff000000000000000000000000000000000000000000909116610100909302929092178217815501613004858783614cea565b506001600160a01b038916600090815260086020526040902060028101879055600301613032888a83614cea565b506001600160a01b038916600090815260086020526040902060058101805460ff191660ff841617905560040161306a838583614cea565b5088600a6000888152602001908152602001600020868660405161308f929190614936565b908152602001604051809103902060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555088600989896040516130d4929190614936565b90815260405190819003602001902080546001600160a01b039283167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617909155600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180549b9092169a16999099179098555050505050505050565b600089898989898989898960405160240161318699989796959493929190614e38565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6e9e2d3f0000000000000000000000000000000000000000000000000000000017905290506131e981613899565b50505050505050505050565b6131fd61235d565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361286f565b6001600160a01b038216613290576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260086020526040902054610100900416613315576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5a52433230206e6f7420726567697374657265640000000000000000000000006044820152606401612402565b6001600160a01b03919091166000908152600860205260409020805460ff1916911515919091179055565b6040516001600160a01b0383166024820152811515604482015260009060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9060bda9000000000000000000000000000000000000000000000000000000001790529050610f5081613899565b60008581526006602052604090205460ff1680156133de5750805b15613418576040517fa1452cb000000000000000000000000000000000000000000000000000000000815260048101869052602401612402565b60008581526006602052604090205460ff16158015613435575080155b1561346f576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612402565b6000858152600660205260409020600201546001600160a01b03161580156134975750468514155b156134d257600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018590555b6000858152600660205260409020805460ff19168215151781556002810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03871617905560030161352f838583614cea565b50801561357057600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0185905561258b565b61258b856139f2565b60008585858585604051602401613594959493929190614ea3565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa8f2cb9600000000000000000000000000000000000000000000000000000000179052905061262081613899565b6123b9613aa0565b613607613aa0565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b60008781526006602052604090205460ff1661367d576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101889052602401612402565b60008590036136bb5785856040517ec10cfd000000000000000000000000000000000000000000000000000000008152600401612402929190614ad7565b6000878152600760205260409081902090516136da9088908890614936565b9081526040519081900360200190205460ff16613729578686866040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161240293929190614ca3565b8181600760008a8152602001908152602001600020888860405161374e929190614936565b9081526020016040518091039020600301868660405161376f929190614936565b90815260200160405180910390209182611c28929190614cea565b6000878787878787876040516024016137a99796959493929190614ee0565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff354b31f000000000000000000000000000000000000000000000000000000001790529050611c2881613899565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610f86576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401612402565b60005b600254811015610f865746600282815481106138ba576138ba614a2f565b9060005260206000200154148061390e575060066000600283815481106138e3576138e3614a2f565b90600052602060002001548152602001908152602001600020600301805461390a90614946565b1590505b613939576139396002828154811061392857613928614a2f565b906000526020600020015483613b07565b60010161389c565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166123b9576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6139a582613dc9565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156139ea57610f508282613e71565b610f86613ee7565b60005b600254811015610f86578160028281548110613a1357613a13614a2f565b906000526020600020015403613a985760028054613a3390600190614f30565b81548110613a4357613a43614a2f565b906000526020600020015460028281548110613a6157613a61614a2f565b6000918252602090912001556002805480613a7e57613a7e614f6a565b600190038181906000526020600020016000905590555050565b6001016139f5565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff166123b9576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526207a120815260006020808301829052835160a0810185528281529081018290529283018190526060808401526080830152906000848152600660205260408082206002015490517ffc5fecd50000000000000000000000000000000000000000000000000000000081526207a12060048201526001600160a01b039091169190829063fc5fecd5906024016040805180830381865afa158015613bb8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613bdc9190614f99565b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018290529092506001600160a01b03841691506323b872dd906064016020604051808303816000875af1158015613c4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c709190614fc7565b613ca6576040517f90b8ec1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303816000875af1158015613d12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d369190614fc7565b50600b546000878152600660205260409081902090517f06cb89830000000000000000000000000000000000000000000000000000000081526001600160a01b03909216916306cb898391613d9b9160039091019086908a908a908a90600401614fe4565b600060405180830381600087803b158015613db557600080fd5b505af11580156131e9573d6000803e3d6000fd5b806001600160a01b03163b600003613e18576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401612402565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051613e8e9190614a5e565b600060405180830381855af49150503d8060008114613ec9576040519150601f19603f3d011682016040523d82523d6000602084013e613ece565b606091505b5091509150613ede858383613f1f565b95945050505050565b34156123b9576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082613f3457613f2f82613f94565b6112db565b8151158015613f4b57506001600160a01b0384163b155b15613f8d576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401612402565b50806112db565b805115613fa45780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215613fe857600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146112db57600080fd5b60008083601f84011261402a57600080fd5b50813567ffffffffffffffff81111561404257600080fd5b60208301915083602082850101111561405a57600080fd5b9250929050565b8015158114610f6857600080fd5b6000806000806060858703121561408557600080fd5b84359350602085013567ffffffffffffffff8111156140a357600080fd5b6140af87828801614018565b90945092505060408501356140c381614061565b939692955090935050565b60005b838110156140e95781810151838201526020016140d1565b50506000910152565b6000815180845261410a8160208601602086016140ce565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156141d3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180511515865260208101516080602088015261419660808801826140f2565b9050604082015187820360408901526141af82826140f2565b60609384015198909301979097525094506020938401939190910190600101614146565b50929695505050505050565b6000806000806000606086880312156141f757600080fd5b85359450602086013567ffffffffffffffff81111561421557600080fd5b61422188828901614018565b909550935050604086013567ffffffffffffffff81111561424157600080fd5b61424d88828901614018565b969995985093965092949392505050565b60006020828403121561427057600080fd5b5035919050565b6001600160a01b0381168114610f6857600080fd5b6000806040838503121561429f57600080fd5b8235915060208301356142b181614277565b809150509250929050565b6000602082840312156142ce57600080fd5b81356112db81614277565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806040838503121561431b57600080fd5b823561432681614277565b9150602083013567ffffffffffffffff81111561434257600080fd5b8301601f8101851361435357600080fd5b803567ffffffffffffffff81111561436d5761436d6142d9565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff8211171561439d5761439d6142d9565b6040528181528282016020018710156143b557600080fd5b816020840160208301376000602083830101528093505050509250929050565b6000806000604084860312156143ea57600080fd5b83359250602084013567ffffffffffffffff81111561440857600080fd5b61441486828701614018565b9497909650939450505050565b821515815260406020820152600061443c60408301846140f2565b949350505050565b600080600080600080600080600060c08a8c03121561446257600080fd5b893561446d81614277565b985060208a013567ffffffffffffffff81111561448957600080fd5b6144958c828d01614018565b90995097505060408a0135955060608a013567ffffffffffffffff8111156144bc57600080fd5b6144c88c828d01614018565b90965094505060808a013567ffffffffffffffff8111156144e857600080fd5b6144f48c828d01614018565b90945092505060a08a013560ff8116811461450e57600080fd5b809150509295985092959850929598565b6020815260006112db60208301846140f2565b6001600160a01b038316815260406020820152600061443c60408301846140f2565b6000806040838503121561456757600080fd5b823561457281614277565b915060208301356142b181614061565b602080825282518282018190526000918401906040840190835b818110156145ba57835183526020938401939092019160010161459c565b509095945050505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156141d3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160208701526001600160a01b036040820151166040870152606081015190506080606087015261465c60808701826140f2565b95505060209384019391909101906001016145ed565b60008060008060006080868803121561468a57600080fd5b85359450602086013561469c81614277565b9350604086013567ffffffffffffffff8111156146b857600080fd5b6146c488828901614018565b90945092505060608601356146d881614061565b809150509295509295909350565b6000806000606084860312156146fb57600080fd5b833561470681614277565b9250602084013561471681614277565b9150604084013561472681614277565b809150509250925092565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156141d3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051151586526001600160a01b036020820151166020870152604081015160e060408801526147bc60e08801826140f2565b905060608201516060880152608082015187820360808901526147df82826140f2565b91505060a082015187820360a08901526147f982826140f2565b91505060c0820151915061481260c088018360ff169052565b9550506020938401939190910190600101614759565b861515815260c06020820152600061484360c08301886140f2565b866040840152828103606084015261485b81876140f2565b9050828103608084015261486f81866140f2565b91505060ff831660a0830152979650505050505050565b60008060008060008060006080888a0312156148a157600080fd5b87359650602088013567ffffffffffffffff8111156148bf57600080fd5b6148cb8a828b01614018565b909750955050604088013567ffffffffffffffff8111156148eb57600080fd5b6148f78a828b01614018565b909550935050606088013567ffffffffffffffff81111561491757600080fd5b6149238a828b01614018565b989b979a50959850939692959293505050565b8183823760009101908152919050565b600181811c9082168061495a57607f821691505b602082108103614993577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b600081546149a681614946565b8085526001821680156149c057600181146149dc57614a13565b60ff1983166020870152602082151560051b8701019350614a13565b84600052602060002060005b83811015614a0a5781546020828a0101526001820191506020810190506149e8565b87016020019450505b50505092915050565b6020815260006112db6020830184614999565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008251614a708184602087016140ce565b9190910192915050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b604081526000614ab9604083018688614a7a565b8281036020840152614acc818587614a7a565b979650505050505050565b60208152600061443c602083018486614a7a565b60ff85168152836020820152606060408201526000614b0e606083018486614a7a565b9695505050505050565b601f821115610f5057806000526020600020601f840160051c81016020851015614b3f5750805b601f840160051c820191505b8181101561258b5760008155600101614b4b565b815167ffffffffffffffff811115614b7957614b796142d9565b614b8d81614b878454614946565b84614b18565b6020601f821160018114614bdf5760008315614ba95750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b17845561258b565b600084815260208120601f198516915b82811015614c0f5787850151825560209485019460019092019101614bef565b5084821015614c4b57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b606081526000614c6e60608301888a614a7a565b8281036020840152614c81818789614a7a565b90508281036040840152614c96818587614a7a565b9998505050505050505050565b838152604060208201526000613ede604083018486614a7a565b848152606060208201526000614cd7606083018587614a7a565b9050821515604083015295945050505050565b67ffffffffffffffff831115614d0257614d026142d9565b614d1683614d108354614946565b83614b18565b6000601f841160018114614d685760008515614d325750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b17835561258b565b600083815260209020601f19861690835b82811015614d995786850135825560209485019460019092019101614d79565b5086821015614dd4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b858152606060208201526000614e00606083018688614a7a565b8281036040840152614e13818587614a7a565b98975050505050505050565b600060208284031215614e3157600080fd5b5051919050565b6001600160a01b038a16815260c060208201526000614e5b60c083018a8c614a7a565b8860408401528281036060840152614e7481888a614a7a565b90508281036080840152614e89818688614a7a565b91505060ff831660a08301529a9950505050505050505050565b8581526001600160a01b0385166020820152608060408201526000614ecc608083018587614a7a565b905082151560608301529695505050505050565b878152608060208201526000614efa60808301888a614a7a565b8281036040840152614f0d818789614a7a565b90508281036060840152614f22818587614a7a565b9a9950505050505050505050565b81810381811115610932577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008060408385031215614fac57600080fd5b8251614fb781614277565b6020939093015192949293505050565b600060208284031215614fd957600080fd5b81516112db81614061565b60c081526000614ff760c0830188614999565b6001600160a01b0387166020840152828103604084015261501881876140f2565b90508451606084015260208501511515608084015282810360a08401526001600160a01b0384511681526020840151151560208201526001600160a01b036040850151166040820152606084015160a0606083015261507a60a08301826140f2565b9050608085015160808301528092505050969550505050505056fea2646970667358221220185320aaa2ff1b87c3db2fc5c5c150ce1b0fb71a2ace4fd7700966ba6b36ecda64736f6c634300081a003360a060405234801561001057600080fd5b50737cce3eb018bf23e1fe2a32692f2c77592d1103946001600160a01b031663cc5ad8b66040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610065573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610089919061009a565b6001600160a01b03166080526100ca565b6000602082840312156100ac57600080fd5b81516001600160a01b03811681146100c357600080fd5b9392505050565b60805161095e6100e46000396000607e015261095e6000f3fe6080604052600436106100635760003560e01c80637b103999116100405780637b10399914610109578063c9028a3614610131578063ebf9b2aa1461014457005b8063116191b61461006c5780632d4cfb7e146100c95780635bcfd616146100e957005b3661006a57005b005b34801561007857600080fd5b506100a07f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100d557600080fd5b5061006a6100e43660046102b4565b610157565b3480156100f557600080fd5b5061006a610104366004610380565b610191565b34801561011557600080fd5b506100a0737cce3eb018bf23e1fe2a32692f2c77592d11039481565b61006a61013f36600461040a565b61020c565b61006a610152366004610445565b61023b565b7f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7816040516101869190610560565b60405180910390a150565b606081156101a8576101a582840184610666565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e6101d3878061075c565b6101e360408a0160208b016107c1565b896040013533866040516101fc969594939291906107dc565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c481604051610186919061089e565b606081156102525761024f82840184610666565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e61027d858061075c565b61028d60408801602089016107c1565b876040013533866040516102a6969594939291906107dc565b60405180910390a150505050565b6000602082840312156102c657600080fd5b813567ffffffffffffffff8111156102dd57600080fd5b820160c081850312156102ef57600080fd5b9392505050565b60006060828403121561030857600080fd5b50919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461033257600080fd5b919050565b60008083601f84011261034957600080fd5b50813567ffffffffffffffff81111561036157600080fd5b60208301915083602082850101111561037957600080fd5b9250929050565b60008060008060006080868803121561039857600080fd5b853567ffffffffffffffff8111156103af57600080fd5b6103bb888289016102f6565b9550506103ca6020870161030e565b935060408601359250606086013567ffffffffffffffff8111156103ed57600080fd5b6103f988828901610337565b969995985093965092949392505050565b60006020828403121561041c57600080fd5b813567ffffffffffffffff81111561043357600080fd5b8201608081850312156102ef57600080fd5b60008060006040848603121561045a57600080fd5b833567ffffffffffffffff81111561047157600080fd5b61047d868287016102f6565b935050602084013567ffffffffffffffff81111561049a57600080fd5b6104a686828701610337565b9497909650939450505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126104e857600080fd5b830160208101925035905067ffffffffffffffff81111561050857600080fd5b80360382131561037957600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60208152600061057083846104b3565b60c0602085015261058560e085018284610517565b91505073ffffffffffffffffffffffffffffffffffffffff6105a96020860161030e565b1660408401526000604085013590508060608501525060608401358015158082146105d357600080fd5b80608086015250506000608085013590508060a0850152506105f860a08501856104b3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c086015261062d838284610517565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561067857600080fd5b813567ffffffffffffffff81111561068f57600080fd5b8201601f810184136106a057600080fd5b803567ffffffffffffffff8111156106ba576106ba610637565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561072657610726610637565b60405281815282820160200186101561073e57600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261079157600080fd5b83018035915067ffffffffffffffff8211156107ac57600080fd5b60200191503681900382131561037957600080fd5b6000602082840312156107d357600080fd5b6102ef8261030e565b60a0815260006107f060a08301888a610517565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b8181101561085a5760208187018101518483018201520161083e565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff6108c08361030e565b16602082015273ffffffffffffffffffffffffffffffffffffffff6108e76020840161030e565b166040820152600080604084013590508060608401525061090b60608401846104b3565b60808085015261091f60a085018284610517565b9594505050505056fea26469706673582212206f7565d8db8f7480b60517069080a5ccd0f8b46abb36bd8b344abd6930ec4d0664736f6c634300081a003360c060405234801561001057600080fd5b5060405161106f38038061106f83398101604081905261002f916100db565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006357604051632b2add3d60e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0385811691909117909155828116608052811660a0526040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac590600090a150505061011e565b80516001600160a01b03811681146100d657600080fd5b919050565b6000806000606084860312156100f057600080fd5b6100f9846100bf565b9250610107602085016100bf565b9150610115604085016100bf565b90509250925092565b60805160a051610f2561014a60003960006101e50152600081816102b9015261045b0152610f256000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806397770dff11610097578063c63585cc11610066578063c63585cc14610273578063d7fd7afb14610286578063d936a012146102b4578063ee2815ba146102db57600080fd5b806397770dff1461021a578063a7cb05071461022d578063c39aca3714610240578063c62178ac1461025357600080fd5b8063513a9c05116100d3578063513a9c051461018a578063569541b9146101c0578063842da36d146101e057806391dd645f1461020757600080fd5b80630be15547146100fa5780631f0e251b1461015a5780633ce4a5bc1461016f575b600080fd5b610130610108366004610bd1565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61016d610168366004610c13565b6102ee565b005b61013073735b14bb79463307aacbed86daf3322b1e6226ab81565b610130610198366004610bd1565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101309073ffffffffffffffffffffffffffffffffffffffff1681565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d610215366004610c35565b610402565b61016d610228366004610c13565b610526565b61016d61023b366004610c61565b610633565b61016d61024e366004610c83565b6106ce565b6004546101309073ffffffffffffffffffffffffffffffffffffffff1681565b610130610281366004610d53565b6108cd565b6102a6610294366004610bd1565b60006020819052908152604090205481565b604051908152602001610151565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d6102e9366004610c35565b610a02565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461033b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610388576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461044f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354600090610497907f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff16846108cd565b60008481526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251878152918201529192507f0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e910160405180910390a1505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610573576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166105c0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e906020016103f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610680576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461071b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831673735b14bb79463307aacbed86daf3322b1e6226ab1480610768575073ffffffffffffffffffffffffffffffffffffffff831630145b1561079f576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690528616906347e7ef24906044016020604051808303816000875af1158015610814573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108389190610d96565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063de43156e906108939089908990899088908890600401610e01565b600060405180830381600087803b1580156108ad57600080fd5b505af11580156108c1573d6000803e3d6000fd5b50505050505050505050565b60008060006108dc8585610ad3565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b16603482015291935091508690604801604051602081830303815290604052805190602001206040516020016109c29291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a4f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d91016106c2565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b3b576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610b75578284610b78565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610bca576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b600060208284031215610be357600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c0e57600080fd5b919050565b600060208284031215610c2557600080fd5b610c2e82610bea565b9392505050565b60008060408385031215610c4857600080fd5b82359150610c5860208401610bea565b90509250929050565b60008060408385031215610c7457600080fd5b50508035926020909101359150565b60008060008060008060a08789031215610c9c57600080fd5b863567ffffffffffffffff811115610cb357600080fd5b87016060818a031215610cc557600080fd5b9550610cd360208801610bea565b945060408701359350610ce860608801610bea565b9250608087013567ffffffffffffffff811115610d0457600080fd5b8701601f81018913610d1557600080fd5b803567ffffffffffffffff811115610d2c57600080fd5b896020828401011115610d3e57600080fd5b60208201935080925050509295509295509295565b600080600060608486031215610d6857600080fd5b610d7184610bea565b9250610d7f60208501610bea565b9150610d8d60408501610bea565b90509250925092565b600060208284031215610da857600080fd5b81518015158114610c2e57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60808152600086357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018112610e3957600080fd5b870160208101903567ffffffffffffffff811115610e5657600080fd5b803603821315610e6557600080fd5b60606080850152610e7a60e085018284610db8565b91505073ffffffffffffffffffffffffffffffffffffffff610e9e60208a01610bea565b1660a0840152604088013560c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401528281036060840152610ee3818587610db8565b9897505050505050505056fea26469706673582212204635fa482dcaad895b3b4ec621fd47a9c4c1d3aa3f8abdb3203f0116e91c4ab964736f6c634300081a003360c060405234801561001057600080fd5b5060405161207238038061207283398101604081905261002f916101f0565b6001600160a01b038216158061004c57506001600160a01b038116155b1561006a5760405163d92e233d60e01b815260040160405180910390fd5b60066100768982610342565b5060076100838882610342565b506008805460ff191660ff881617905560808590528360028111156100aa576100aa610400565b60a08160028111156100be576100be610400565b905250600192909255600080546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506104169350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261013357600080fd5b81516001600160401b0381111561014c5761014c61010c565b604051601f8201601f19908116603f011681016001600160401b038111828210171561017a5761017a61010c565b60405281815283820160200185101561019257600080fd5b60005b828110156101b157602081860181015183830182015201610195565b506000918101602001919091529392505050565b8051600381106101d457600080fd5b919050565b80516001600160a01b03811681146101d457600080fd5b600080600080600080600080610100898b03121561020d57600080fd5b88516001600160401b0381111561022357600080fd5b61022f8b828c01610122565b60208b015190995090506001600160401b0381111561024d57600080fd5b6102598b828c01610122565b975050604089015160ff8116811461027057600080fd5b60608a0151909650945061028660808a016101c5565b60a08a0151909450925061029c60c08a016101d9565b91506102aa60e08a016101d9565b90509295985092959890939650565b600181811c908216806102cd57607f821691505b6020821081036102ed57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561033d57806000526020600020601f840160051c8101602085101561031a5750805b601f840160051c820191505b8181101561033a5760008155600101610326565b50505b505050565b81516001600160401b0381111561035b5761035b61010c565b61036f8161036984546102b9565b846102f3565b6020601f8211600181146103a3576000831561038b5750848201515b600019600385901b1c1916600184901b17845561033a565b600084815260208120601f198516915b828110156103d357878501518255602094850194600190920191016103b3565b50848210156103f15786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a051611c1b61045760003960006103440152600081816102f001528181610bdc01528181610ce201528181610efe01526110040152611c1b6000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806395d89b41116100f9578063ccc7759911610097578063eddeb12311610071578063eddeb12314610461578063f2441b3214610474578063f687d12a14610494578063fc5fecd5146104a757600080fd5b8063ccc77599146103d4578063d9eeebed146103e7578063dd62ed3e1461041b57600080fd5b8063b84c8246116100d3578063b84c824614610386578063c47f00271461039b578063c7012626146103ae578063c835d7cc146103c157600080fd5b806395d89b4114610337578063a3413d031461033f578063a9059cbb1461037357600080fd5b80633ce4a5bc116101665780634d8943bb116101405780634d8943bb146102ac57806370a08231146102b557806385e1f4d0146102eb5780638b851b951461031257600080fd5b80633ce4a5bc1461024657806342966c681461028657806347e7ef241461029957600080fd5b806318160ddd1161019757806318160ddd1461021657806323b872dd1461021e578063313ce5671461023157600080fd5b806306fdde03146101be578063091d2788146101dc578063095ea7b3146101f3575b600080fd5b6101c66104ba565b6040516101d39190611648565b60405180910390f35b6101e560015481565b6040519081526020016101d3565b610206610201366004611687565b61054c565b60405190151581526020016101d3565b6005546101e5565b61020661022c3660046116b3565b610563565b60085460405160ff90911681526020016101d3565b61026173735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6102066102943660046116f4565b6105fa565b6102066102a7366004611687565b61060e565b6101e560025481565b6101e56102c336600461170d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101e57f000000000000000000000000000000000000000000000000000000000000000081565b60085461026190610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101c6610767565b6103667f000000000000000000000000000000000000000000000000000000000000000081565b6040516101d3919061172a565b610206610381366004611687565b610776565b610399610394366004611832565b610783565b005b6103996103a9366004611832565b6107e0565b6102066103bc366004611883565b610839565b6103996103cf36600461170d565b610988565b6103996103e236600461170d565b610a9c565b6103ef610bb0565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101d3565b6101e56104293660046118dc565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b61039961046f3660046116f4565b610dce565b6000546102619073ffffffffffffffffffffffffffffffffffffffff1681565b6103996104a23660046116f4565b610e50565b6103ef6104b53660046116f4565b610ed2565b6060600680546104c990611915565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590611915565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b5050505050905090565b60006105593384846110ee565b5060015b92915050565b60006105708484846111f7565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054828110156105db576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ef85336105ea8685611997565b6110ee565b506001949350505050565b600061060633836113b2565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab1480159061064c575060005473ffffffffffffffffffffffffffffffffffffffff163314155b80156106755750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b156106ac576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106b683836114f4565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107569186906119aa565b60405180910390a250600192915050565b6060600780546104c990611915565b60006105593384846111f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107d0576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107dc8282611a1b565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461082d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107dc8282611a1b565b6000806000610846610bb0565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156108d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fc9190611b34565b610932576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61093c33856113b2565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161097591899189918791611b56565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109d5576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a22576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610ae9576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b36576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c679190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610cb6576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d699190611ba2565b905080600003610da5576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610db89190611bbb565b610dc29190611bd2565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e1b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a91565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e9d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f899190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610fd8576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015611067573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108b9190611ba2565b9050806000036110c7576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000906110d78784611bbb565b6110e19190611bd2565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661113b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611188576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611244576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611291576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040902054818110156112f1576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112fb8282611997565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260036020526040808220939093559085168152908120805484929061133e908490611bd2565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113a491815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113ff576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260409020548181101561145f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114698282611997565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812091909155600580548492906114a4908490611997565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111ea565b73ffffffffffffffffffffffffffffffffffffffff8216611541576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008282546115539190611bd2565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805483929061158d908490611bd2565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561160a576020818501810151868301820152016115ee565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061165b60208301846115e4565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461168457600080fd5b50565b6000806040838503121561169a57600080fd5b82356116a581611662565b946020939093013593505050565b6000806000606084860312156116c857600080fd5b83356116d381611662565b925060208401356116e381611662565b929592945050506040919091013590565b60006020828403121561170657600080fd5b5035919050565b60006020828403121561171f57600080fd5b813561165b81611662565b6020810160038310611765577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff8411156117b5576117b561176b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156118025761180261176b565b60405283815290508082840185101561181a57600080fd5b83836020830137600060208583010152509392505050565b60006020828403121561184457600080fd5b813567ffffffffffffffff81111561185b57600080fd5b8201601f8101841361186c57600080fd5b61187b8482356020840161179a565b949350505050565b6000806040838503121561189657600080fd5b823567ffffffffffffffff8111156118ad57600080fd5b8301601f810185136118be57600080fd5b6118cd8582356020840161179a565b95602094909401359450505050565b600080604083850312156118ef57600080fd5b82356118fa81611662565b9150602083013561190a81611662565b809150509250929050565b600181811c9082168061192957607f821691505b602082108103611962577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561055d5761055d611968565b6040815260006119bd60408301856115e4565b90508260208301529392505050565b601f821115611a1657806000526020600020601f840160051c810160208510156119f35750805b601f840160051c820191505b81811015611a1357600081556001016119ff565b50505b505050565b815167ffffffffffffffff811115611a3557611a3561176b565b611a4981611a438454611915565b846119cc565b6020601f821160018114611a9b5760008315611a655750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611a13565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611ae95787850151825560209485019460019092019101611ac9565b5084821015611b2557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b600060208284031215611b4657600080fd5b8151801515811461165b57600080fd5b608081526000611b6960808301876115e4565b6020830195909552506040810192909252606090910152919050565b600060208284031215611b9757600080fd5b815161165b81611662565b600060208284031215611bb457600080fd5b5051919050565b808202811582820484141761055d5761055d611968565b8082018082111561055d5761055d61196856fea26469706673582212205c01f2682d8e2240fabc45c195d1ebe0c11d19758f25f18c2d70d3c6c00973bc64736f6c634300081a003347617320666565206e6f74206275726e656420636f72726563746c792066726f6d20676173205a524332305a5243323020746f6b656e732077657265206e6f74206275726e656420636f72726563746c795769746864726177616c20616d6f756e74206e6f74206275726e656420636f72726563746c792066726f6d207365636f6e64205a524332305a5243323020746f6b656e732077657265206e6f74206275726e656420636f72726563746c7920666f72207769746864726177416e6443616c6ca2646970667358221220cddc5dc4ae55e29547a121e3ec41c38bf2057e27e9efde56bee4fbd5ddb1b37064736f6c634300081a0033",
}

// GatewayZEVMOutboundTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayZEVMOutboundTestMetaData.ABI instead.
var GatewayZEVMOutboundTestABI = GatewayZEVMOutboundTestMetaData.ABI

// GatewayZEVMOutboundTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayZEVMOutboundTestMetaData.Bin instead.
var GatewayZEVMOutboundTestBin = GatewayZEVMOutboundTestMetaData.Bin

// DeployGatewayZEVMOutboundTest deploys a new Ethereum contract, binding an instance of GatewayZEVMOutboundTest to it.
func DeployGatewayZEVMOutboundTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayZEVMOutboundTest, error) {
	parsed, err := GatewayZEVMOutboundTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayZEVMOutboundTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayZEVMOutboundTest{GatewayZEVMOutboundTestCaller: GatewayZEVMOutboundTestCaller{contract: contract}, GatewayZEVMOutboundTestTransactor: GatewayZEVMOutboundTestTransactor{contract: contract}, GatewayZEVMOutboundTestFilterer: GatewayZEVMOutboundTestFilterer{contract: contract}}, nil
}

// GatewayZEVMOutboundTest is an auto generated Go binding around an Ethereum contract.
type GatewayZEVMOutboundTest struct {
	GatewayZEVMOutboundTestCaller     // Read-only binding to the contract
	GatewayZEVMOutboundTestTransactor // Write-only binding to the contract
	GatewayZEVMOutboundTestFilterer   // Log filterer for contract events
}

// GatewayZEVMOutboundTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMOutboundTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMOutboundTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayZEVMOutboundTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMOutboundTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayZEVMOutboundTestSession struct {
	Contract     *GatewayZEVMOutboundTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GatewayZEVMOutboundTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayZEVMOutboundTestCallerSession struct {
	Contract *GatewayZEVMOutboundTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// GatewayZEVMOutboundTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayZEVMOutboundTestTransactorSession struct {
	Contract     *GatewayZEVMOutboundTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// GatewayZEVMOutboundTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestRaw struct {
	Contract *GatewayZEVMOutboundTest // Generic contract binding to access the raw methods on
}

// GatewayZEVMOutboundTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestCallerRaw struct {
	Contract *GatewayZEVMOutboundTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayZEVMOutboundTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestTransactorRaw struct {
	Contract *GatewayZEVMOutboundTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayZEVMOutboundTest creates a new instance of GatewayZEVMOutboundTest, bound to a specific deployed contract.
func NewGatewayZEVMOutboundTest(address common.Address, backend bind.ContractBackend) (*GatewayZEVMOutboundTest, error) {
	contract, err := bindGatewayZEVMOutboundTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTest{GatewayZEVMOutboundTestCaller: GatewayZEVMOutboundTestCaller{contract: contract}, GatewayZEVMOutboundTestTransactor: GatewayZEVMOutboundTestTransactor{contract: contract}, GatewayZEVMOutboundTestFilterer: GatewayZEVMOutboundTestFilterer{contract: contract}}, nil
}

// NewGatewayZEVMOutboundTestCaller creates a new read-only instance of GatewayZEVMOutboundTest, bound to a specific deployed contract.
func NewGatewayZEVMOutboundTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayZEVMOutboundTestCaller, error) {
	contract, err := bindGatewayZEVMOutboundTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestCaller{contract: contract}, nil
}

// NewGatewayZEVMOutboundTestTransactor creates a new write-only instance of GatewayZEVMOutboundTest, bound to a specific deployed contract.
func NewGatewayZEVMOutboundTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayZEVMOutboundTestTransactor, error) {
	contract, err := bindGatewayZEVMOutboundTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestTransactor{contract: contract}, nil
}

// NewGatewayZEVMOutboundTestFilterer creates a new log filterer instance of GatewayZEVMOutboundTest, bound to a specific deployed contract.
func NewGatewayZEVMOutboundTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayZEVMOutboundTestFilterer, error) {
	contract, err := bindGatewayZEVMOutboundTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestFilterer{contract: contract}, nil
}

// bindGatewayZEVMOutboundTest binds a generic wrapper to an already deployed contract.
func bindGatewayZEVMOutboundTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayZEVMOutboundTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMOutboundTest.Contract.GatewayZEVMOutboundTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.GatewayZEVMOutboundTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.GatewayZEVMOutboundTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMOutboundTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ISTEST() (bool, error) {
	return _GatewayZEVMOutboundTest.Contract.ISTEST(&_GatewayZEVMOutboundTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ISTEST() (bool, error) {
	return _GatewayZEVMOutboundTest.Contract.ISTEST(&_GatewayZEVMOutboundTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVMOutboundTest.Contract.PAUSERROLE(&_GatewayZEVMOutboundTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVMOutboundTest.Contract.PAUSERROLE(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeArtifacts(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeArtifacts(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeContracts(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeContracts(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeSenders(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeSenders(&_GatewayZEVMOutboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) Failed() (bool, error) {
	return _GatewayZEVMOutboundTest.Contract.Failed(&_GatewayZEVMOutboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) Failed() (bool, error) {
	return _GatewayZEVMOutboundTest.Contract.Failed(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetArtifactSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetArtifactSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetArtifacts(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetArtifacts(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetContracts(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetContracts(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetInterfaces(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetInterfaces(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetSenders(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetSenders(&_GatewayZEVMOutboundTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.SetUp(&_GatewayZEVMOutboundTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.SetUp(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForDifferentZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x75ca5587.
//
// Solidity: function testBurnGasFeeForDifferentZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnGasFeeForDifferentZRC20WithdrawAndCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnGasFeeForDifferentZRC20WithdrawAndCall")
}

// TestBurnGasFeeForDifferentZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x75ca5587.
//
// Solidity: function testBurnGasFeeForDifferentZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnGasFeeForDifferentZRC20WithdrawAndCall() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForDifferentZRC20WithdrawAndCall(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForDifferentZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x75ca5587.
//
// Solidity: function testBurnGasFeeForDifferentZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnGasFeeForDifferentZRC20WithdrawAndCall() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForDifferentZRC20WithdrawAndCall(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForDifferentZRC20Withdrawal is a paid mutator transaction binding the contract method 0xb84a3d2f.
//
// Solidity: function testBurnGasFeeForDifferentZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnGasFeeForDifferentZRC20Withdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnGasFeeForDifferentZRC20Withdrawal")
}

// TestBurnGasFeeForDifferentZRC20Withdrawal is a paid mutator transaction binding the contract method 0xb84a3d2f.
//
// Solidity: function testBurnGasFeeForDifferentZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnGasFeeForDifferentZRC20Withdrawal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForDifferentZRC20Withdrawal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForDifferentZRC20Withdrawal is a paid mutator transaction binding the contract method 0xb84a3d2f.
//
// Solidity: function testBurnGasFeeForDifferentZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnGasFeeForDifferentZRC20Withdrawal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForDifferentZRC20Withdrawal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x3177f381.
//
// Solidity: function testBurnGasFeeForZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnGasFeeForZRC20WithdrawAndCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnGasFeeForZRC20WithdrawAndCall")
}

// TestBurnGasFeeForZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x3177f381.
//
// Solidity: function testBurnGasFeeForZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnGasFeeForZRC20WithdrawAndCall() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForZRC20WithdrawAndCall(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x3177f381.
//
// Solidity: function testBurnGasFeeForZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnGasFeeForZRC20WithdrawAndCall() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForZRC20WithdrawAndCall(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForZRC20Withdrawal is a paid mutator transaction binding the contract method 0xf2a91358.
//
// Solidity: function testBurnGasFeeForZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnGasFeeForZRC20Withdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnGasFeeForZRC20Withdrawal")
}

// TestBurnGasFeeForZRC20Withdrawal is a paid mutator transaction binding the contract method 0xf2a91358.
//
// Solidity: function testBurnGasFeeForZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnGasFeeForZRC20Withdrawal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForZRC20Withdrawal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForZRC20Withdrawal is a paid mutator transaction binding the contract method 0xf2a91358.
//
// Solidity: function testBurnGasFeeForZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnGasFeeForZRC20Withdrawal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForZRC20Withdrawal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnProtocolFeesFailsWithInsufficientAllowance is a paid mutator transaction binding the contract method 0x61cfddb7.
//
// Solidity: function testBurnProtocolFeesFailsWithInsufficientAllowance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnProtocolFeesFailsWithInsufficientAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnProtocolFeesFailsWithInsufficientAllowance")
}

// TestBurnProtocolFeesFailsWithInsufficientAllowance is a paid mutator transaction binding the contract method 0x61cfddb7.
//
// Solidity: function testBurnProtocolFeesFailsWithInsufficientAllowance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnProtocolFeesFailsWithInsufficientAllowance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnProtocolFeesFailsWithInsufficientAllowance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnProtocolFeesFailsWithInsufficientAllowance is a paid mutator transaction binding the contract method 0x61cfddb7.
//
// Solidity: function testBurnProtocolFeesFailsWithInsufficientAllowance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnProtocolFeesFailsWithInsufficientAllowance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnProtocolFeesFailsWithInsufficientAllowance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDeposit is a paid mutator transaction binding the contract method 0x7f924c4e.
//
// Solidity: function testDeposit() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDeposit")
}

// TestDeposit is a paid mutator transaction binding the contract method 0x7f924c4e.
//
// Solidity: function testDeposit() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDeposit() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDeposit(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDeposit is a paid mutator transaction binding the contract method 0x7f924c4e.
//
// Solidity: function testDeposit() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDeposit() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDeposit(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x884660a3.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContract")
}

// TestDepositAndRevertZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x884660a3.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x884660a3.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway is a paid mutator transaction binding the contract method 0x51336fb0.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway is a paid mutator transaction binding the contract method 0x51336fb0.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway is a paid mutator transaction binding the contract method 0x51336fb0.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol is a paid mutator transaction binding the contract method 0x6163f8ef.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol is a paid mutator transaction binding the contract method 0x6163f8ef.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol is a paid mutator transaction binding the contract method 0x6163f8ef.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xeab7674e.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xeab7674e.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xeab7674e.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xf1d98f1b.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xf1d98f1b.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xf1d98f1b.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x339bd828.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x339bd828.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x339bd828.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x48f4fd07.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x48f4fd07.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x48f4fd07.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x996b7675.
//
// Solidity: function testDepositFailsIfAmountIs0() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfAmountIs0")
}

// TestDepositFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x996b7675.
//
// Solidity: function testDepositFailsIfAmountIs0() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfAmountIs0(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x996b7675.
//
// Solidity: function testDepositFailsIfAmountIs0() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfAmountIs0(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfSenderNotProtocol is a paid mutator transaction binding the contract method 0xca26929c.
//
// Solidity: function testDepositFailsIfSenderNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfSenderNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfSenderNotProtocol")
}

// TestDepositFailsIfSenderNotProtocol is a paid mutator transaction binding the contract method 0xca26929c.
//
// Solidity: function testDepositFailsIfSenderNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfSenderNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfSenderNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfSenderNotProtocol is a paid mutator transaction binding the contract method 0xca26929c.
//
// Solidity: function testDepositFailsIfSenderNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfSenderNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfSenderNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x96d9d876.
//
// Solidity: function testDepositFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfTargetIsGateway")
}

// TestDepositFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x96d9d876.
//
// Solidity: function testDepositFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x96d9d876.
//
// Solidity: function testDepositFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x3626c616.
//
// Solidity: function testDepositFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfTargetIsProtocol")
}

// TestDepositFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x3626c616.
//
// Solidity: function testDepositFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x3626c616.
//
// Solidity: function testDepositFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x14b7a6da.
//
// Solidity: function testDepositFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfTargetIsZeroAddress")
}

// TestDepositFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x14b7a6da.
//
// Solidity: function testDepositFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x14b7a6da.
//
// Solidity: function testDepositFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x9c9acd5d.
//
// Solidity: function testDepositFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfZRC20IsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfZRC20IsZeroAddress")
}

// TestDepositFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x9c9acd5d.
//
// Solidity: function testDepositFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x9c9acd5d.
//
// Solidity: function testDepositFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositTogglePause is a paid mutator transaction binding the contract method 0x1c785a14.
//
// Solidity: function testDepositTogglePause() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositTogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositTogglePause")
}

// TestDepositTogglePause is a paid mutator transaction binding the contract method 0x1c785a14.
//
// Solidity: function testDepositTogglePause() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositTogglePause() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositTogglePause(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositTogglePause is a paid mutator transaction binding the contract method 0x1c785a14.
//
// Solidity: function testDepositTogglePause() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositTogglePause() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositTogglePause(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETA is a paid mutator transaction binding the contract method 0x445bbf47.
//
// Solidity: function testDepositZETA() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETA")
}

// TestDepositZETA is a paid mutator transaction binding the contract method 0x445bbf47.
//
// Solidity: function testDepositZETA() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETA() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETA(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETA is a paid mutator transaction binding the contract method 0x445bbf47.
//
// Solidity: function testDepositZETA() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETA() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETA(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversal is a paid mutator transaction binding the contract method 0x828d267c.
//
// Solidity: function testDepositZETAAndCallUniversal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversal")
}

// TestDepositZETAAndCallUniversal is a paid mutator transaction binding the contract method 0x828d267c.
//
// Solidity: function testDepositZETAAndCallUniversal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversal is a paid mutator transaction binding the contract method 0x828d267c.
//
// Solidity: function testDepositZETAAndCallUniversal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContract is a paid mutator transaction binding the contract method 0x2468bc0f.
//
// Solidity: function testDepositZETAAndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContract")
}

// TestDepositZETAAndCallUniversalContract is a paid mutator transaction binding the contract method 0x2468bc0f.
//
// Solidity: function testDepositZETAAndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContract is a paid mutator transaction binding the contract method 0x2468bc0f.
//
// Solidity: function testDepositZETAAndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xc35cb5e4.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol")
}

// TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xc35cb5e4.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xc35cb5e4.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero is a paid mutator transaction binding the contract method 0xec294d9f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero")
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero is a paid mutator transaction binding the contract method 0xec294d9f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero is a paid mutator transaction binding the contract method 0xec294d9f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x5cec7db5.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway")
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x5cec7db5.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x5cec7db5.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x890a2d67.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol")
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x890a2d67.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x890a2d67.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x97f7661f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress")
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x97f7661f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x97f7661f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0xb936be8c.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfZeroAddress")
}

// TestDepositZETAAndCallUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0xb936be8c.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0xb936be8c.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContract is a paid mutator transaction binding the contract method 0x0c4a14b4.
//
// Solidity: function testDepositZETAAndRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContract")
}

// TestDepositZETAAndRevertUniversalContract is a paid mutator transaction binding the contract method 0x0c4a14b4.
//
// Solidity: function testDepositZETAAndRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContract is a paid mutator transaction binding the contract method 0x0c4a14b4.
//
// Solidity: function testDepositZETAAndRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xda8fd81d.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero")
}

// TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xda8fd81d.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xda8fd81d.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xe36c0a23.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol")
}

// TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xe36c0a23.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xe36c0a23.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x514fcc7f.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway")
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x514fcc7f.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x514fcc7f.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x50dd38ca.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol")
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x50dd38ca.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x50dd38ca.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x2ed7cff6.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress")
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x2ed7cff6.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x2ed7cff6.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractWhenPaused is a paid mutator transaction binding the contract method 0x62fe2915.
//
// Solidity: function testDepositZETAAndRevertUniversalContractWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractWhenPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractWhenPaused")
}

// TestDepositZETAAndRevertUniversalContractWhenPaused is a paid mutator transaction binding the contract method 0x62fe2915.
//
// Solidity: function testDepositZETAAndRevertUniversalContractWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractWhenPaused() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractWhenPaused(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractWhenPaused is a paid mutator transaction binding the contract method 0x62fe2915.
//
// Solidity: function testDepositZETAAndRevertUniversalContractWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractWhenPaused() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractWhenPaused(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x195ac143.
//
// Solidity: function testDepositZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfAmountIsZero")
}

// TestDepositZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x195ac143.
//
// Solidity: function testDepositZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x195ac143.
//
// Solidity: function testDepositZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfInsufficientProtocolBalance is a paid mutator transaction binding the contract method 0xf589655d.
//
// Solidity: function testDepositZETAFailsIfInsufficientProtocolBalance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfInsufficientProtocolBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfInsufficientProtocolBalance")
}

// TestDepositZETAFailsIfInsufficientProtocolBalance is a paid mutator transaction binding the contract method 0xf589655d.
//
// Solidity: function testDepositZETAFailsIfInsufficientProtocolBalance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfInsufficientProtocolBalance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfInsufficientProtocolBalance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfInsufficientProtocolBalance is a paid mutator transaction binding the contract method 0xf589655d.
//
// Solidity: function testDepositZETAFailsIfInsufficientProtocolBalance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfInsufficientProtocolBalance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfInsufficientProtocolBalance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xae4662ec.
//
// Solidity: function testDepositZETAFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfSenderIsNotProtocol")
}

// TestDepositZETAFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xae4662ec.
//
// Solidity: function testDepositZETAFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xae4662ec.
//
// Solidity: function testDepositZETAFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0xb82ce921.
//
// Solidity: function testDepositZETAFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfTargetIsGateway")
}

// TestDepositZETAFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0xb82ce921.
//
// Solidity: function testDepositZETAFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0xb82ce921.
//
// Solidity: function testDepositZETAFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x21d92591.
//
// Solidity: function testDepositZETAFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfTargetIsProtocol")
}

// TestDepositZETAFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x21d92591.
//
// Solidity: function testDepositZETAFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x21d92591.
//
// Solidity: function testDepositZETAFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x1ff5bdd7.
//
// Solidity: function testDepositZETAFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfTargetIsZeroAddress")
}

// TestDepositZETAFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x1ff5bdd7.
//
// Solidity: function testDepositZETAFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x1ff5bdd7.
//
// Solidity: function testDepositZETAFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsWhenPaused is a paid mutator transaction binding the contract method 0x33f7c81a.
//
// Solidity: function testDepositZETAFailsWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsWhenPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsWhenPaused")
}

// TestDepositZETAFailsWhenPaused is a paid mutator transaction binding the contract method 0x33f7c81a.
//
// Solidity: function testDepositZETAFailsWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsWhenPaused() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsWhenPaused(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsWhenPaused is a paid mutator transaction binding the contract method 0x33f7c81a.
//
// Solidity: function testDepositZETAFailsWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsWhenPaused() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsWhenPaused(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x6efa04b5.
//
// Solidity: function testDepositZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContract")
}

// TestDepositZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x6efa04b5.
//
// Solidity: function testDepositZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x6efa04b5.
//
// Solidity: function testDepositZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x58c9987f.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero")
}

// TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x58c9987f.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x58c9987f.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xcf2c3d1d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol")
}

// TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xcf2c3d1d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xcf2c3d1d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xef2b5394.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress")
}

// TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xef2b5394.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xef2b5394.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0xfb339a1c.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress")
}

// TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0xfb339a1c.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0xfb339a1c.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsGateway is a paid mutator transaction binding the contract method 0xe09bc659.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractIfTargetIsGateway")
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsGateway is a paid mutator transaction binding the contract method 0xe09bc659.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsGateway is a paid mutator transaction binding the contract method 0xe09bc659.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol is a paid mutator transaction binding the contract method 0xeb78bd7d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractIfTargetIsProtocol")
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol is a paid mutator transaction binding the contract method 0xeb78bd7d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol is a paid mutator transaction binding the contract method 0xeb78bd7d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteAbortUniversalContract is a paid mutator transaction binding the contract method 0x1832cb6e.
//
// Solidity: function testExecuteAbortUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteAbortUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteAbortUniversalContract")
}

// TestExecuteAbortUniversalContract is a paid mutator transaction binding the contract method 0x1832cb6e.
//
// Solidity: function testExecuteAbortUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteAbortUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteAbortUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteAbortUniversalContract is a paid mutator transaction binding the contract method 0x1832cb6e.
//
// Solidity: function testExecuteAbortUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteAbortUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteAbortUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xee0f4ea1.
//
// Solidity: function testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress")
}

// TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xee0f4ea1.
//
// Solidity: function testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xee0f4ea1.
//
// Solidity: function testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x3ab5b199.
//
// Solidity: function testExecuteFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteFailsIfTargetIsZeroAddress")
}

// TestExecuteFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x3ab5b199.
//
// Solidity: function testExecuteFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x3ab5b199.
//
// Solidity: function testExecuteFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x2948df41.
//
// Solidity: function testExecuteFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteFailsIfZRC20IsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteFailsIfZRC20IsZeroAddress")
}

// TestExecuteFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x2948df41.
//
// Solidity: function testExecuteFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x2948df41.
//
// Solidity: function testExecuteFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContract is a paid mutator transaction binding the contract method 0x084fafab.
//
// Solidity: function testExecuteRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteRevertUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteRevertUniversalContract")
}

// TestExecuteRevertUniversalContract is a paid mutator transaction binding the contract method 0x084fafab.
//
// Solidity: function testExecuteRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteRevertUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContract is a paid mutator transaction binding the contract method 0x084fafab.
//
// Solidity: function testExecuteRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteRevertUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xc8814d2e.
//
// Solidity: function testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress")
}

// TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xc8814d2e.
//
// Solidity: function testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xc8814d2e.
//
// Solidity: function testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContractIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x27820625.
//
// Solidity: function testExecuteRevertUniversalContractIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteRevertUniversalContractIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteRevertUniversalContractIfSenderIsNotProtocol")
}

// TestExecuteRevertUniversalContractIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x27820625.
//
// Solidity: function testExecuteRevertUniversalContractIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteRevertUniversalContractIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContractIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContractIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x27820625.
//
// Solidity: function testExecuteRevertUniversalContractIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteRevertUniversalContractIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContractIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContract is a paid mutator transaction binding the contract method 0x7cec29b0.
//
// Solidity: function testExecuteUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteUniversalContract")
}

// TestExecuteUniversalContract is a paid mutator transaction binding the contract method 0x7cec29b0.
//
// Solidity: function testExecuteUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContract is a paid mutator transaction binding the contract method 0x7cec29b0.
//
// Solidity: function testExecuteUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x2acb21b4.
//
// Solidity: function testExecuteUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteUniversalContractFailsIfSenderIsNotProtocol")
}

// TestExecuteUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x2acb21b4.
//
// Solidity: function testExecuteUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x2acb21b4.
//
// Solidity: function testExecuteUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x5b4c90e1.
//
// Solidity: function testExecuteUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteUniversalContractFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteUniversalContractFailsIfZeroAddress")
}

// TestExecuteUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x5b4c90e1.
//
// Solidity: function testExecuteUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteUniversalContractFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContractFailsIfZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x5b4c90e1.
//
// Solidity: function testExecuteUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteUniversalContractFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContractFailsIfZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// GatewayZEVMOutboundTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestCalledIterator struct {
	Event *GatewayZEVMOutboundTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestCalled)
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
		it.Event = new(GatewayZEVMOutboundTestCalled)
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
func (it *GatewayZEVMOutboundTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestCalled represents a Called event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestCalled struct {
	Sender        common.Address
	Zrc20         common.Address
	Receiver      []byte
	Message       []byte
	CallOptions   CallOptions
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, zrc20 []common.Address) (*GatewayZEVMOutboundTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestCalledIterator{contract: _GatewayZEVMOutboundTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestCalled, sender []common.Address, zrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestCalled)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "Called", log); err != nil {
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

// ParseCalled is a log parse operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseCalled(log types.Log) (*GatewayZEVMOutboundTestCalled, error) {
	event := new(GatewayZEVMOutboundTestCalled)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestContextDataIterator is returned from FilterContextData and is used to iterate over the raw logs and unpacked data for ContextData events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataIterator struct {
	Event *GatewayZEVMOutboundTestContextData // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestContextDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestContextData)
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
		it.Event = new(GatewayZEVMOutboundTestContextData)
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
func (it *GatewayZEVMOutboundTestContextDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestContextDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestContextData represents a ContextData event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextData struct {
	Origin    []byte
	Sender    common.Address
	ChainID   *big.Int
	MsgSender common.Address
	Message   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContextData is a free log retrieval operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterContextData(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestContextDataIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestContextDataIterator{contract: _GatewayZEVMOutboundTest.contract, event: "ContextData", logs: logs, sub: sub}, nil
}

// WatchContextData is a free log subscription operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchContextData(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestContextData) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestContextData)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextData", log); err != nil {
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

// ParseContextData is a log parse operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseContextData(log types.Log) (*GatewayZEVMOutboundTestContextData, error) {
	event := new(GatewayZEVMOutboundTestContextData)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestContextDataAbortIterator is returned from FilterContextDataAbort and is used to iterate over the raw logs and unpacked data for ContextDataAbort events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataAbortIterator struct {
	Event *GatewayZEVMOutboundTestContextDataAbort // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestContextDataAbortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestContextDataAbort)
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
		it.Event = new(GatewayZEVMOutboundTestContextDataAbort)
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
func (it *GatewayZEVMOutboundTestContextDataAbortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestContextDataAbortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestContextDataAbort represents a ContextDataAbort event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataAbort struct {
	AbortContext AbortContext
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContextDataAbort is a free log retrieval operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterContextDataAbort(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestContextDataAbortIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "ContextDataAbort")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestContextDataAbortIterator{contract: _GatewayZEVMOutboundTest.contract, event: "ContextDataAbort", logs: logs, sub: sub}, nil
}

// WatchContextDataAbort is a free log subscription operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchContextDataAbort(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestContextDataAbort) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "ContextDataAbort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestContextDataAbort)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextDataAbort", log); err != nil {
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

// ParseContextDataAbort is a log parse operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseContextDataAbort(log types.Log) (*GatewayZEVMOutboundTestContextDataAbort, error) {
	event := new(GatewayZEVMOutboundTestContextDataAbort)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextDataAbort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestContextDataRevertIterator is returned from FilterContextDataRevert and is used to iterate over the raw logs and unpacked data for ContextDataRevert events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataRevertIterator struct {
	Event *GatewayZEVMOutboundTestContextDataRevert // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestContextDataRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestContextDataRevert)
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
		it.Event = new(GatewayZEVMOutboundTestContextDataRevert)
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
func (it *GatewayZEVMOutboundTestContextDataRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestContextDataRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestContextDataRevert represents a ContextDataRevert event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataRevert struct {
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContextDataRevert is a free log retrieval operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterContextDataRevert(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestContextDataRevertIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestContextDataRevertIterator{contract: _GatewayZEVMOutboundTest.contract, event: "ContextDataRevert", logs: logs, sub: sub}, nil
}

// WatchContextDataRevert is a free log subscription operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchContextDataRevert(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestContextDataRevert) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestContextDataRevert)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
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

// ParseContextDataRevert is a log parse operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseContextDataRevert(log types.Log) (*GatewayZEVMOutboundTestContextDataRevert, error) {
	event := new(GatewayZEVMOutboundTestContextDataRevert)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestWithdrawnIterator struct {
	Event *GatewayZEVMOutboundTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestWithdrawn)
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
		it.Event = new(GatewayZEVMOutboundTestWithdrawn)
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
func (it *GatewayZEVMOutboundTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestWithdrawn represents a Withdrawn event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestWithdrawn struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Zrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	CallOptions     CallOptions
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMOutboundTestWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestWithdrawnIterator{contract: _GatewayZEVMOutboundTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestWithdrawn)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseWithdrawn(log types.Log) (*GatewayZEVMOutboundTestWithdrawn, error) {
	event := new(GatewayZEVMOutboundTestWithdrawn)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestWithdrawnAndCalledIterator struct {
	Event *GatewayZEVMOutboundTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestWithdrawnAndCalled)
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
		it.Event = new(GatewayZEVMOutboundTestWithdrawnAndCalled)
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
func (it *GatewayZEVMOutboundTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestWithdrawnAndCalled struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Zrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	CallOptions     CallOptions
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMOutboundTestWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestWithdrawnAndCalledIterator{contract: _GatewayZEVMOutboundTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestWithdrawnAndCalled)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayZEVMOutboundTestWithdrawnAndCalled, error) {
	event := new(GatewayZEVMOutboundTestWithdrawnAndCalled)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogIterator struct {
	Event *GatewayZEVMOutboundTestLog // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLog)
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
		it.Event = new(GatewayZEVMOutboundTestLog)
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
func (it *GatewayZEVMOutboundTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLog represents a Log event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLog)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLog(log types.Log) (*GatewayZEVMOutboundTestLog, error) {
	event := new(GatewayZEVMOutboundTestLog)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogAddressIterator struct {
	Event *GatewayZEVMOutboundTestLogAddress // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogAddress)
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
		it.Event = new(GatewayZEVMOutboundTestLogAddress)
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
func (it *GatewayZEVMOutboundTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogAddress represents a LogAddress event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogAddressIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogAddressIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogAddress)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogAddress(log types.Log) (*GatewayZEVMOutboundTestLogAddress, error) {
	event := new(GatewayZEVMOutboundTestLogAddress)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArrayIterator struct {
	Event *GatewayZEVMOutboundTestLogArray // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogArray)
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
		it.Event = new(GatewayZEVMOutboundTestLogArray)
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
func (it *GatewayZEVMOutboundTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogArray represents a LogArray event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogArrayIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogArrayIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogArray)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogArray(log types.Log) (*GatewayZEVMOutboundTestLogArray, error) {
	event := new(GatewayZEVMOutboundTestLogArray)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray0Iterator struct {
	Event *GatewayZEVMOutboundTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogArray0)
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
		it.Event = new(GatewayZEVMOutboundTestLogArray0)
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
func (it *GatewayZEVMOutboundTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogArray0 represents a LogArray0 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogArray0Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogArray0)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogArray0(log types.Log) (*GatewayZEVMOutboundTestLogArray0, error) {
	event := new(GatewayZEVMOutboundTestLogArray0)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray1Iterator struct {
	Event *GatewayZEVMOutboundTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogArray1)
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
		it.Event = new(GatewayZEVMOutboundTestLogArray1)
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
func (it *GatewayZEVMOutboundTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogArray1 represents a LogArray1 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogArray1Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogArray1)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogArray1(log types.Log) (*GatewayZEVMOutboundTestLogArray1, error) {
	event := new(GatewayZEVMOutboundTestLogArray1)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogBytesIterator struct {
	Event *GatewayZEVMOutboundTestLogBytes // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogBytes)
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
		it.Event = new(GatewayZEVMOutboundTestLogBytes)
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
func (it *GatewayZEVMOutboundTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogBytes represents a LogBytes event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogBytesIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogBytesIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogBytes)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogBytes(log types.Log) (*GatewayZEVMOutboundTestLogBytes, error) {
	event := new(GatewayZEVMOutboundTestLogBytes)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogBytes32Iterator struct {
	Event *GatewayZEVMOutboundTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogBytes32)
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
		it.Event = new(GatewayZEVMOutboundTestLogBytes32)
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
func (it *GatewayZEVMOutboundTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogBytes32 represents a LogBytes32 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogBytes32Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogBytes32)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogBytes32(log types.Log) (*GatewayZEVMOutboundTestLogBytes32, error) {
	event := new(GatewayZEVMOutboundTestLogBytes32)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogIntIterator struct {
	Event *GatewayZEVMOutboundTestLogInt // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogInt)
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
		it.Event = new(GatewayZEVMOutboundTestLogInt)
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
func (it *GatewayZEVMOutboundTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogInt represents a LogInt event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogIntIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogIntIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogInt)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogInt(log types.Log) (*GatewayZEVMOutboundTestLogInt, error) {
	event := new(GatewayZEVMOutboundTestLogInt)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedAddressIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedAddress)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedAddress)
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
func (it *GatewayZEVMOutboundTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedAddressIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedAddress)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayZEVMOutboundTestLogNamedAddress, error) {
	event := new(GatewayZEVMOutboundTestLogNamedAddress)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArrayIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedArray)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedArray)
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
func (it *GatewayZEVMOutboundTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedArray represents a LogNamedArray event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedArrayIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedArray)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayZEVMOutboundTestLogNamedArray, error) {
	event := new(GatewayZEVMOutboundTestLogNamedArray)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray0Iterator struct {
	Event *GatewayZEVMOutboundTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedArray0)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedArray0)
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
func (it *GatewayZEVMOutboundTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedArray0Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedArray0)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayZEVMOutboundTestLogNamedArray0, error) {
	event := new(GatewayZEVMOutboundTestLogNamedArray0)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray1Iterator struct {
	Event *GatewayZEVMOutboundTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedArray1)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedArray1)
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
func (it *GatewayZEVMOutboundTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedArray1Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedArray1)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayZEVMOutboundTestLogNamedArray1, error) {
	event := new(GatewayZEVMOutboundTestLogNamedArray1)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedBytesIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedBytes)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedBytes)
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
func (it *GatewayZEVMOutboundTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedBytesIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedBytes)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayZEVMOutboundTestLogNamedBytes, error) {
	event := new(GatewayZEVMOutboundTestLogNamedBytes)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedBytes32Iterator struct {
	Event *GatewayZEVMOutboundTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedBytes32)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedBytes32)
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
func (it *GatewayZEVMOutboundTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedBytes32Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedBytes32)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayZEVMOutboundTestLogNamedBytes32, error) {
	event := new(GatewayZEVMOutboundTestLogNamedBytes32)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedDecimalIntIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedDecimalInt)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedDecimalInt)
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
func (it *GatewayZEVMOutboundTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedDecimalIntIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedDecimalInt)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayZEVMOutboundTestLogNamedDecimalInt, error) {
	event := new(GatewayZEVMOutboundTestLogNamedDecimalInt)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedDecimalUintIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedDecimalUint)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedDecimalUint)
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
func (it *GatewayZEVMOutboundTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedDecimalUintIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedDecimalUint)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayZEVMOutboundTestLogNamedDecimalUint, error) {
	event := new(GatewayZEVMOutboundTestLogNamedDecimalUint)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedIntIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedInt)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedInt)
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
func (it *GatewayZEVMOutboundTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedInt represents a LogNamedInt event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedIntIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedInt)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayZEVMOutboundTestLogNamedInt, error) {
	event := new(GatewayZEVMOutboundTestLogNamedInt)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedStringIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedString)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedString)
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
func (it *GatewayZEVMOutboundTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedString represents a LogNamedString event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedStringIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedString)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedString(log types.Log) (*GatewayZEVMOutboundTestLogNamedString, error) {
	event := new(GatewayZEVMOutboundTestLogNamedString)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedUintIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedUint)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedUint)
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
func (it *GatewayZEVMOutboundTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedUint represents a LogNamedUint event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedUintIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedUint)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayZEVMOutboundTestLogNamedUint, error) {
	event := new(GatewayZEVMOutboundTestLogNamedUint)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogStringIterator struct {
	Event *GatewayZEVMOutboundTestLogString // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogString)
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
		it.Event = new(GatewayZEVMOutboundTestLogString)
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
func (it *GatewayZEVMOutboundTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogString represents a LogString event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogStringIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogStringIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogString)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogString(log types.Log) (*GatewayZEVMOutboundTestLogString, error) {
	event := new(GatewayZEVMOutboundTestLogString)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogUintIterator struct {
	Event *GatewayZEVMOutboundTestLogUint // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogUint)
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
		it.Event = new(GatewayZEVMOutboundTestLogUint)
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
func (it *GatewayZEVMOutboundTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogUint represents a LogUint event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogUintIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogUintIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogUint)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogUint(log types.Log) (*GatewayZEVMOutboundTestLogUint, error) {
	event := new(GatewayZEVMOutboundTestLogUint)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogsIterator struct {
	Event *GatewayZEVMOutboundTestLogs // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogs)
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
		it.Event = new(GatewayZEVMOutboundTestLogs)
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
func (it *GatewayZEVMOutboundTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogs represents a Logs event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogsIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogsIterator{contract: _GatewayZEVMOutboundTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogs)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "logs", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogs(log types.Log) (*GatewayZEVMOutboundTestLogs, error) {
	event := new(GatewayZEVMOutboundTestLogs)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
