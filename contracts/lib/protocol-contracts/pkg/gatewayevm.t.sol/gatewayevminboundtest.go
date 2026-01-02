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

// GatewayEVMInboundTestMetaData contains all meta data concerning the GatewayEVMInboundTest contract.
var GatewayEVMInboundTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testCallFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfCallOnRevertIsTrue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfDestinationIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTss\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZetaToConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositERC20ToCustodyIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositERC20ToCustodyIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositERC20ToCustodyWithPayloadIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssWithPayloadIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositedAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallOnRevertNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055620f4240602c55348015603357600080fd5b5061d147806100436000396000f3fe608060405234801561001057600080fd5b50600436106102d35760003560e01c806366d9a9a011610186578063b1526934116100e3578063c0fab86d11610097578063e306a97811610071578063e306a97814610464578063e85c5a071461046c578063fa7626d41461047457600080fd5b8063c0fab86d1461044c578063cdef966f14610454578063e20c9f711461045c57600080fd5b8063b5508aa9116100c8578063b5508aa914610424578063ba414fa61461042c578063ba46ba231461044457600080fd5b8063b152693414610414578063b28490631461041c57600080fd5b80639d96310a1161013a578063aa030c1c1161011f578063aa030c1c146103fc578063b0464fdc14610404578063b1409f711461040c57600080fd5b80639d96310a146103ec5780639fd1e597146103f457600080fd5b80638f5a42241161016b5780638f5a4224146103c7578063916a17c6146103cf57806395cd0445146103e457600080fd5b806366d9a9a01461039d57806385226c81146103b257600080fd5b806330f7c04f11610234578063466f332e116101e857806351da903d116101cd57806351da903d146103855780635c4013d01461038d5780636459542a1461039557600080fd5b8063466f332e146103755780634ce25c0a1461037d57600080fd5b80633f7286f4116102195780633f7286f41461035d57806341b8fec614610365578063458085f51461036d57600080fd5b806330f7c04f1461034d5780633e5e3c231461035557600080fd5b806318e141381161028b5780631f4f542f116102705780631f4f542f146103285780632a5dcf36146103305780632ade38801461033857600080fd5b806318e14138146103025780631ed7831c1461030a57600080fd5b806309b21ccb116102bc57806309b21ccb146102ea5780630a9254e4146102f25780630fc37f5e146102fa57600080fd5b806301a74aee146102d85780630724d8e3146102e2575b600080fd5b6102e0610481565b005b6102e061067b565b6102e061082f565b6102e061096d565b6102e061148e565b6102e061175c565b61031261189d565b60405161031f9190619716565b60405180910390f35b6102e06118ff565b6102e0611b18565b610340611c51565b60405161031f91906197b2565b6102e0611d93565b6103126121d8565b610312612238565b6102e0612298565b6102e06123bb565b6102e0612635565b6102e0612a0a565b6102e0612d42565b6102e0612e75565b6102e0612f7d565b6103a5613376565b60405161031f9190619918565b6103ba6134e3565b60405161031f91906199b6565b6102e06135b3565b6103d76136a4565b60405161031f9190619a2d565b6102e061378a565b6102e0613bf9565b6102e0613d31565b6102e0613f25565b6103d76140b4565b6102e061419a565b6102e061452e565b6102e06146ac565b6103ba614980565b610434614a50565b604051901515815260200161031f565b6102e0614b24565b6102e0614c9a565b6102e0614e7c565b610312614f84565b6102e0614fe4565b6102e06152a7565b601f546104349060ff1681565b6026546040516001600160a01b03909116602482015260009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a08101835261032180825260008284018190528285019190915283519283019093528282526060810191909152919250906080810161050e621e84806001619af3565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916105c191600401619b06565b600060405180830381600087803b1580156105db57600080fd5b505af11580156105ef573d6000803e3d6000fd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb49350610645929091169086908690600401619b74565b600060405180830381600087803b15801561065f57600080fd5b505af1158015610673573d6000803e3d6000fd5b505050505050565b6027546020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039182166084820152620186a092919091163190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561071757600080fd5b505af115801561072b573d6000803e3d6000fd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c9061077a908690600090602890619ce4565b60405180910390a36020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263726ac97c9286926107d59290911690602890600401619d1a565b6000604051808303818588803b1580156107ee57600080fd5b505af1158015610802573d6000803e3d6000fd5b50506027546001600160a01b031631925061082a915061082490508484619af3565b82615536565b505050565b6026546040516001600160a01b03909116602482015260019060009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156108fe57600080fd5b505af1158015610912573d6000803e3d6000fd5b50506020546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063d09e3b7893506106459260009288929116908790602890600401619d3c565b602580547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602680548216611234179055602780549091166156781790556040516109bf90619643565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f080158015610a44573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556027546040519116908190610a8d90619650565b6001600160a01b03928316815291166020820152604001604051809103906000f080158015610ac0573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081178255604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260275460255492519086169481019490945260448401929092529092166064820152610b9d91906084015b60408051601f198184030181529190526020810180516001600160e01b03167fc0c53b8b000000000000000000000000000000000000000000000000000000001790526155b5565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff000000000000000000000000000000000000000090911681178255604080518082018252601081527f4552433230437573746f64792e736f6c000000000000000000000000000000009381019390935260275460255491516024810193909352841660448301529092166064830152610c6d91608401610b55565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602180549190920483167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055604080518082018252601a81527f5a657461436f6e6e6563746f724e6f6e4e61746976652e736f6c00000000000060208083019190915254602480546027546025549551938716928401929092528516604483015284166064820152919092166084820152610d91919060a40160408051601f198184030181529190526020810180516001600160e01b03167ff8c8765e000000000000000000000000000000000000000000000000000000001790526155b5565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602280549190920483167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790556027546040517fca669fa700000000000000000000000000000000000000000000000000000000815291166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015610e6957600080fd5b505af1158015610e7d573d6000803e3d6000fd5b5050602480546027546022546040517f15d57fd40000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152908216938101939093521692506315d57fd49150604401600060405180830381600087803b158015610ef057600080fd5b505af1158015610f04573d6000803e3d6000fd5b50506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b158015610f8857600080fd5b505af1158015610f9c573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561101257600080fd5b505af1158015611026573d6000803e3d6000fd5b50506020546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f9150602401600060405180830381600087803b15801561108c57600080fd5b505af11580156110a0573d6000803e3d6000fd5b50506020546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef9150602401600060405180830381600087803b15801561110657600080fd5b505af115801561111a573d6000803e3d6000fd5b50506021546023546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639b19251a9150602401600060405180830381600087803b15801561118057600080fd5b505af1158015611194573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156111f657600080fd5b505af115801561120a573d6000803e3d6000fd5b5050602354602554602c546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810191909152911692506340c10f199150604401600060405180830381600087803b15801561127b57600080fd5b505af115801561128f573d6000803e3d6000fd5b50506027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561130557600080fd5b505af1158015611319573d6000803e3d6000fd5b5050602254602554602c546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810191909152600060448201529116925063106e62909150606401600060405180830381600087803b15801561139157600080fd5b505af11580156113a5573d6000803e3d6000fd5b50506040805160a0810182526103218082526000602080840182815284860193845285519182019095528181526060840181905260808401919091528251602880549551151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009096166001600160a01b0392831617959095178555915160298054919093167fffffffffffffffffffffffff00000000000000000000000000000000000000009190911617909155909350909150602a9061147e9082619e07565b5060808201518160030155905050565b60208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa1580156114ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115109190619ec6565b61151b906001619af3565b67ffffffffffffffff81111561153357611533619d91565b6040519080825280601f01601f19166020018201604052801561155d576020820181803683370190505b50602a9061156b9082619e07565b5060006028600201805461157e90619ba8565b905090506000602060009054906101000a90046001600160a01b03166001600160a01b031663a2ba19346040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115fb9190619ec6565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916116a291600401619b06565b600060405180830381600087803b1580156116bc57600080fd5b505af11580156116d0573d6000803e3d6000fd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c9350600192611726921690602890600401619d1a565b6000604051808303818588803b15801561173f57600080fd5b505af1158015611753573d6000803e3d6000fd5b50505050505050565b6026546040516001600160a01b039091166024820152600090819060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052517fc31eb0e00000000000000000000000000000000000000000000000000000000081527f951e19ed000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b600060405180830381600087803b15801561182b57600080fd5b505af115801561183f573d6000803e3d6000fd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b7894506106459392831692889216908790602890600401619d3c565b606060168054806020026020016040519081016040528092919081815260200182805480156118f557602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116118d7575b5050505050905090565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303816000875af1158015611974573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119989190619edf565b506040805160a081018252610321808252600060208084018290528385019290925283519182019093528281526060820152608081016119dc621e84806001619af3565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391611a8f91600401619b06565b600060405180830381600087803b158015611aa957600080fd5b505af1158015611abd573d6000803e3d6000fd5b50506020546026546023546040517f102614b00000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063102614b094506106459392831692889216908790600401619f01565b6026546040516001600160a01b039091166024820152600090819060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052517fc31eb0e00000000000000000000000000000000000000000000000000000000081527f7671265e000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015611be657600080fd5b505af1158015611bfa573d6000803e3d6000fd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b935086926117269216908690602890600401619f38565b6060601e805480602002602001604051908101604052809291908181526020016000905b82821015611d8a57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015611d73578382906000526020600020018054611ce690619ba8565b80601f0160208091040260200160405190810160405280929190818152602001828054611d1290619ba8565b8015611d5f5780601f10611d3457610100808354040283529160200191611d5f565b820191906000526020600020905b815481529060010190602001808311611d4257829003601f168201915b505050505081526020019060010190611cc7565b505050508152505081526020019060010190611c75565b50505050905090565b6023546021546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a09260009216906370a0823190602401602060405180830381865afa158015611dff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e239190619ec6565b9050611e30600082615536565b6026546040516001600160a01b03909116602482015260009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052602354905491517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101879052929350169063095ea7b3906044016020604051808303816000875af1158015611ee5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f099190619edf565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611f9857600080fd5b505af1158015611fac573d6000803e3d6000fd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f92612002928992909116908790602890619f6c565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789361206693908216928992909116908790602890600401619d3c565b600060405180830381600087803b15801561208057600080fd5b505af1158015612094573d6000803e3d6000fd5b50506023546021546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015612100573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121249190619ec6565b90506121308482615536565b6023546025546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa15801561219a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121be9190619ec6565b90506121d185602c546108249190619fa6565b5050505050565b606060188054806020026020016040519081016040528092919081815260200182805480156118f5576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116118d7575050505050905090565b606060178054806020026020016040519081016040528092919081815260200182805480156118f5576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116118d7575050505050905090565b6040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561232057600080fd5b505af1158015612334573d6000803e3d6000fd5b50506020546023546040517f102614b00000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063102614b0935061238d926000928792911690602890600401619fb9565b600060405180830381600087803b1580156123a757600080fd5b505af11580156121d1573d6000803e3d6000fd5b6026546040516001600160a01b039091166024820152620186a09060009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052602354905491517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101869052929350169063095ea7b3906044016020604051808303816000875af1158015612475573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124999190619edf565b506040805160a081018252610321808252600060208084018290528385019290925283519182019093528281526060820152608081016124dd621e84806001619af3565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39161259091600401619b06565b600060405180830381600087803b1580156125aa57600080fd5b505af11580156125be573d6000803e3d6000fd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b78945061261b93928316928992169088908890600401619ff0565b600060405180830381600087803b15801561173f57600080fd5b60208054604080517fa2ba19340000000000000000000000000000000000000000000000000000000081529051620186a0936000936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa1580156126a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126c59190619ec6565b6126cf919061a039565b67ffffffffffffffff8111156126e7576126e7619d91565b6040519080825280601f01601f191660200182016040528015612711576020820181803683370190505b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa158015612776573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061279a9190619ec6565b6127a4919061a039565b6127af906001619af3565b67ffffffffffffffff8111156127c7576127c7619d91565b6040519080825280601f01601f1916602001820160405280156127f1576020820181803683370190505b50602a906127ff9082619e07565b5060006028600201805461281290619ba8565b835161281e9250619af3565b60208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506000936001600160a01b039092169263a2ba1934926004808401938290030181865afa158015612882573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128a69190619ec6565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39161294d91600401619b06565b600060405180830381600087803b15801561296757600080fd5b505af115801561297b573d6000803e3d6000fd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b935088926129d29216908890602890600401619f38565b6000604051808303818588803b1580156129eb57600080fd5b505af11580156129ff573d6000803e3d6000fd5b505050505050505050565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303816000875af1158015612a7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612aa39190619edf565b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa158015612b02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b269190619ec6565b612b31906001619af3565b67ffffffffffffffff811115612b4957612b49619d91565b6040519080825280601f01601f191660200182016040528015612b73576020820181803683370190505b50602a90612b819082619e07565b50600060286002018054612b9490619ba8565b905090506000602060009054906101000a90046001600160a01b03166001600160a01b031663a2ba19346040518163ffffffff1660e01b8152600401602060405180830381865afa158015612bed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c119190619ec6565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391612cb891600401619b06565b600060405180830381600087803b158015612cd257600080fd5b505af1158015612ce6573d6000803e3d6000fd5b50506020546026546023546040517f102614b00000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063102614b0945061261b939283169289921690602890600401619fb9565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612e0e57600080fd5b505af1158015612e22573d6000803e3d6000fd5b50506020546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250631becceb4915061238d906000908590602890600401619f38565b6040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612efd57600080fd5b505af1158015612f11573d6000803e3d6000fd5b50506020546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063726ac97c91508390612f6490600090602890600401619d1a565b6000604051808303818588803b15801561065f57600080fd5b6023546021546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a09260009216906370a0823190602401602060405180830381865afa158015612fe9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061300d9190619ec6565b905061301a600082615536565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526024810185905291169063095ea7b3906044016020604051808303816000875af1158015613088573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130ac9190619edf565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561313b57600080fd5b505af115801561314f573d6000803e3d6000fd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c926131a392889290911690602890619ce4565b60405180910390a36020546026546023546040517f102614b00000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363102614b0936132059390821692889290911690602890600401619fb9565b600060405180830381600087803b15801561321f57600080fd5b505af1158015613233573d6000803e3d6000fd5b50506023546021546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561329f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132c39190619ec6565b90506132cf8382615536565b6023546025546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015613339573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061335d9190619ec6565b905061337084602c546108249190619fa6565b50505050565b6060601b805480602002602001604051908101604052809291908181526020016000905b82821015611d8a57838290600052602060002090600202016040518060400160405290816000820180546133cd90619ba8565b80601f01602080910402602001604051908101604052809291908181526020018280546133f990619ba8565b80156134465780601f1061341b57610100808354040283529160200191613446565b820191906000526020600020905b81548152906001019060200180831161342957829003601f168201915b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156134cb57602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841161348d5790505b5050505050815250508152602001906001019061339a565b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015611d8a57838290600052602060002001805461352690619ba8565b80601f016020809104026020016040519081016040528092919081815260200182805461355290619ba8565b801561359f5780601f106135745761010080835404028352916020019161359f565b820191906000526020600020905b81548152906001019060200180831161358257829003601f168201915b505050505081526020019060010190613507565b6040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527f7671265e000000000000000000000000000000000000000000000000000000006004820152600090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561363b57600080fd5b505af115801561364f573d6000803e3d6000fd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c93508592612f64921690602890600401619d1a565b6060601d805480602002602001604051908101604052809291908181526020016000905b82821015611d8a5760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561377257602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116137345790505b505050505081525050815260200190600101906136c8565b60208054604080517fa2ba19340000000000000000000000000000000000000000000000000000000081529051620186a0936000936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa1580156137f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061381a9190619ec6565b613824919061a039565b67ffffffffffffffff81111561383c5761383c619d91565b6040519080825280601f01601f191660200182016040528015613866576020820181803683370190505b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa1580156138cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138ef9190619ec6565b6138f9919061a039565b613904906001619af3565b67ffffffffffffffff81111561391c5761391c619d91565b6040519080825280601f01601f191660200182016040528015613946576020820181803683370190505b50602a906139549082619e07565b506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526024810185905291169063095ea7b3906044016020604051808303816000875af11580156139c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139e79190619edf565b506000602860020180546139fa90619ba8565b8351613a069250619af3565b60208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506000936001600160a01b039092169263a2ba1934926004808401938290030181865afa158015613a6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a8e9190619ec6565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391613b3591600401619b06565b600060405180830381600087803b158015613b4f57600080fd5b505af1158015613b63573d6000803e3d6000fd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b789450613bc193928316928a9216908990602890600401619d3c565b600060405180830381600087803b158015613bdb57600080fd5b505af1158015613bef573d6000803e3d6000fd5b5050505050505050565b6026546040516001600160a01b03909116602482015260019060009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015613cc857600080fd5b505af1158015613cdc573d6000803e3d6000fd5b50506020546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063744b9b8b91508490611726906000908690602890600401619f38565b6027546026546040516001600160a01b039182166024820152620186a09291909116319060009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015613e1457600080fd5b505af1158015613e28573d6000803e3d6000fd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f90613e799087906000908790602890619f6c565b60405180910390a36020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263744b9b8b928792613ed692909116908690602890600401619f38565b6000604051808303818588803b158015613eef57600080fd5b505af1158015613f03573d6000803e3d6000fd5b50506027546001600160a01b0316319250613370915061082490508585619af3565b6026546040516001600160a01b03909116602482015260009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015613ffa57600080fd5b505af115801561400e573d6000803e3d6000fd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d9749061405a90859060289061a074565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb49261238d929116908590602890600401619f38565b6060601c805480602002602001604051908101604052809291908181526020016000905b82821015611d8a5760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561418257602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116141445790505b505050505081525050815260200190600101906140d8565b60208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290516000936002936001600160a01b03169263a2ba193492600480830193928290030181865afa1580156141fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906142219190619ec6565b61422b919061a039565b67ffffffffffffffff81111561424357614243619d91565b6040519080825280601f01601f19166020018201604052801561426d576020820181803683370190505b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa1580156142d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906142f69190619ec6565b614300919061a039565b61430b906001619af3565b67ffffffffffffffff81111561432357614323619d91565b6040519080825280601f01601f19166020018201604052801561434d576020820181803683370190505b50602a9061435b9082619e07565b5060006028600201805461436e90619ba8565b835161437a9250619af3565b60208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506000936001600160a01b039092169263a2ba1934926004808401938290030181865afa1580156143de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906144029190619ec6565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916144a991600401619b06565b600060405180830381600087803b1580156144c357600080fd5b505af11580156144d7573d6000803e3d6000fd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb4935061261b92909116908790602890600401619f38565b6040805160a081018252610321808252600060208084018290528385019290925283519182019093528281526060820152620186a0919060808101614577621e84806001619af3565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39161462a91600401619b06565b600060405180830381600087803b15801561464457600080fd5b505af1158015614658573d6000803e3d6000fd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c93508692611726921690869060040161a099565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303816000875af1158015614721573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906147459190619edf565b506025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156147b857600080fd5b505af11580156147cc573d6000803e3d6000fd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a5904279150602401600060405180830381600087803b15801561483257600080fd5b505af1158015614846573d6000803e3d6000fd5b5050602354604080516001600160a01b0390921660248084019190915281518084039091018152604490920181526020820180516001600160e01b03167f1387a34900000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506148f69190600401619b06565b600060405180830381600087803b15801561491057600080fd5b505af1158015614924573d6000803e3d6000fd5b50506020546026546023546040517f102614b00000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063102614b0945061238d939283169287921690602890600401619fb9565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015611d8a5783829060005260206000200180546149c390619ba8565b80601f01602080910402602001604051908101604052809291908181526020018280546149ef90619ba8565b8015614a3c5780601f10614a1157610100808354040283529160200191614a3c565b820191906000526020600020905b815481529060010190602001808311614a1f57829003601f168201915b5050505050815260200190600101906149a4565b60085460009060ff1615614a68575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015614af9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614b1d9190619ec6565b1415905090565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052602880547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000179055517fc31eb0e00000000000000000000000000000000000000000000000000000000081527f19b4bff2000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015614c2f57600080fd5b505af1158015614c43573d6000803e3d6000fd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb4935061238d92909116908590602890600401619f38565b6026546040516001600160a01b039091166024820152620186a09060009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a081018352610321808252600082840181905282850191909152835192830190935282825260608101919091529192509060808101614d2c621e84806001619af3565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391614ddf91600401619b06565b600060405180830381600087803b158015614df957600080fd5b505af1158015614e0d573d6000803e3d6000fd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b93508792614e6392169087908790600401619b74565b6000604051808303818588803b158015613bdb57600080fd5b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260006024820181905292919091169063095ea7b3906044016020604051808303816000875af1158015614eef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614f139190619edf565b506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527f951e19ed000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024016148f6565b606060158054806020026020016040519081016040528092919081815260200182805480156118f5576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116118d7575050505050905090565b602480546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a093810184905291169063095ea7b3906044016020604051808303816000875af1158015615056573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061507a9190619edf565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561510957600080fd5b505af115801561511d573d6000803e3d6000fd5b50506026546025546024546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c9261517192879290911690602890619ce4565b60405180910390a36020546026546024546040517f102614b00000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363102614b0936151d39390821692879290911690602890600401619fb9565b600060405180830381600087803b1580156151ed57600080fd5b505af1158015615201573d6000803e3d6000fd5b5050602480546025546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa15801561526c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906152909190619ec6565b90506152a382602c546108249190619fa6565b5050565b6026546040516001600160a01b039091166024820152620186a09060009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052602354905491517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101869052929350169063095ea7b3906044016020604051808303816000875af1158015615361573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906153859190619edf565b506025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156153f857600080fd5b505af115801561540c573d6000803e3d6000fd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a5904279150602401600060405180830381600087803b15801561547257600080fd5b505af1158015615486573d6000803e3d6000fd5b5050602354604080516001600160a01b0390921660248084019190915281518084039091018152604490920181526020820180516001600160e01b03167f1387a34900000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506118119190600401619b06565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c549060440160006040518083038186803b1580156155a157600080fd5b505afa158015610673573d6000803e3d6000fd5b60006155bf61965d565b6155ca8484836155d4565b9150505b92915050565b6000806155e1858461564f565b90506156446040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f7879000000815250828660405160200161562f92919061a0bb565b6040516020818303038152906040528561565b565b9150505b9392505050565b60006156488383615689565b60c0810151516000901561567f5761567884848460c001516156a4565b9050615648565b615678848461584a565b60006156958383615935565b6156488383602001518461565b565b6000806156af615941565b905060006156bd8683615a14565b905060006156d48260600151836020015185615eba565b905060006156e4838389896160cc565b905060006156f182616f49565b602081015181519192509060030b156157645789826040015160405160200161571b92919061a0dd565b60408051601f19818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261575b91600401619b06565b60405180910390fd5b60006157a76040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001617118565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906157fa908490600401619b06565b602060405180830381865afa158015615817573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061583b919061a15e565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc9259061589f908790600401619b06565b600060405180830381865afa1580156158bc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526158e4919081019061a240565b9050600061591282856040516020016158fe92919061a275565b604051602081830303815290604052617318565b90506001600160a01b0381166155ca57848460405160200161571b92919061a2a4565b6152a38282600061732b565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c906159c890849060040161a34f565b600060405180830381865afa1580156159e5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615a0d919081019061a396565b9250505090565b615a466040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d9050615a916040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b615a9a8561742e565b60208201526000615aaa86617813565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa158015615aec573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615b14919081019061a396565b86838560200151604051602001615b2e949392919061a3df565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb1190615b86908590600401619b06565b600060405180830381865afa158015615ba3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615bcb919081019061a396565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f690615c1390849060040161a4e3565b602060405180830381865afa158015615c30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615c549190619edf565b615c69578160405160200161571b919061a535565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890615cae90849060040161a5c7565b600060405180830381865afa158015615ccb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615cf3919081019061a396565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f690615d3a90849060040161a619565b602060405180830381865afa158015615d57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615d7b9190619edf565b15615e10576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890615dc590849060040161a619565b600060405180830381865afa158015615de2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615e0a919081019061a396565b60408501525b846001600160a01b03166349c4fac8828660000151604051602001615e35919061a66b565b6040516020818303038152906040526040518363ffffffff1660e01b8152600401615e6192919061a6d7565b600060405180830381865afa158015615e7e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615ea6919081019061a396565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b6060815260200190600190039081615ed65790505090506040518060400160405280600481526020017f677265700000000000000000000000000000000000000000000000000000000081525081600081518110615f3657615f3661a6fc565b60200260200101819052506040518060400160405280600381526020017f2d726c000000000000000000000000000000000000000000000000000000000081525081600181518110615f8a57615f8a61a6fc565b602002602001018190525084604051602001615fa6919061a72b565b60405160208183030381529060405281600281518110615fc857615fc861a6fc565b602002602001018190525082604051602001615fe4919061a797565b604051602081830303815290604052816003815181106160065761600661a6fc565b6020026020010181905250600061601c82616f49565b602080820151604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000081850190815282518084018452600080825290860152825180840190935290518252928101929092529192506160ad9060408051808201825260008082526020918201528151808301909252845182528085019082015290617a96565b6160c2578560405160200161571b919061a7d8565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d901561611c565b511590565b616290578260200151156161d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a40161575b565b8260c0015115616290576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a40161575b565b6040805160ff8082526120008201909252600091816020015b60608152602001906001900390816162a957905050905060006040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806163049061a869565b935060ff16815181106163195761631961a6fc565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e370000000000000000000000000000000000000081525060405160200161636a919061a888565b6040516020818303038152906040528282806163859061a869565b935060ff168151811061639a5761639a61a6fc565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806163e79061a869565b935060ff16815181106163fc576163fc61a6fc565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806164499061a869565b935060ff168151811061645e5761645e61a6fc565b6020026020010181905250876020015182828061647a9061a869565b935060ff168151811061648f5761648f61a6fc565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e7472616374506174680000000000000000000000000000000000008152508282806164dc9061a869565b935060ff16815181106164f1576164f161a6fc565b6020908102919091010152875182826165098161a869565b935060ff168151811061651e5761651e61a6fc565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e4964000000000000000000000000000000000000000000000081525082828061656b9061a869565b935060ff16815181106165805761658061a6fc565b602002602001018190525061659446617af7565b828261659f8161a869565b935060ff16815181106165b4576165b461a6fc565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806166019061a869565b935060ff16815181106166165761661661a6fc565b60200260200101819052508682828061662e9061a869565b935060ff16815181106166435761664361a6fc565b602090810291909101015285511561676a5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826166948161a869565b935060ff16815181106166a9576166a961a6fc565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906166f9908990600401619b06565b600060405180830381865afa158015616716573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261673e919081019061a396565b82826167498161a869565b935060ff168151811061675e5761675e61a6fc565b60200260200101819052505b84602001511561683a5760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826167b38161a869565b935060ff16815181106167c8576167c861a6fc565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806168159061a869565b935060ff168151811061682a5761682a61a6fc565b6020026020010181905250616a01565b6168726161178660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6169055760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826168b58161a869565b935060ff16815181106168ca576168ca61a6fc565b60200260200101819052508460a001516040516020016168ea919061a72b565b6040516020818303038152906040528282806168159061a869565b8460c0015115801561694857506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261694690511590565b155b15616a015760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261698c8161a869565b935060ff16815181106169a1576169a161a6fc565b60200260200101819052506169b588617b97565b6040516020016169c5919061a72b565b6040516020818303038152906040528282806169e09061a869565b935060ff16815181106169f5576169f561a6fc565b60200260200101819052505b60408086015181518083018352600080825260209182015282518084019093528151835290810190820152616a3590511590565b616aca5760408051808201909152600b81527f2d2d72656c61796572496400000000000000000000000000000000000000000060208201528282616a788161a869565b935060ff1681518110616a8d57616a8d61a6fc565b60200260200101819052508460400151828280616aa99061a869565b935060ff1681518110616abe57616abe61a6fc565b60200260200101819052505b606085015115616beb5760408051808201909152600681527f2d2d73616c74000000000000000000000000000000000000000000000000000060208201528282616b138161a869565b935060ff1681518110616b2857616b2861a6fc565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015616b97573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616bbf919081019061a396565b8282616bca8161a869565b935060ff1681518110616bdf57616bdf61a6fc565b60200260200101819052505b60e08501515115616c925760408051808201909152600a81527f2d2d6761734c696d69740000000000000000000000000000000000000000000060208201528282616c358161a869565b935060ff1681518110616c4a57616c4a61a6fc565b6020026020010181905250616c668560e0015160000151617af7565b8282616c718161a869565b935060ff1681518110616c8657616c8661a6fc565b60200260200101819052505b60e08501516020015115616d3c5760408051808201909152600a81527f2d2d67617350726963650000000000000000000000000000000000000000000060208201528282616cdf8161a869565b935060ff1681518110616cf457616cf461a6fc565b6020026020010181905250616d108560e0015160200151617af7565b8282616d1b8161a869565b935060ff1681518110616d3057616d3061a6fc565b60200260200101819052505b60e08501516040015115616de65760408051808201909152600e81527f2d2d6d617846656550657247617300000000000000000000000000000000000060208201528282616d898161a869565b935060ff1681518110616d9e57616d9e61a6fc565b6020026020010181905250616dba8560e0015160400151617af7565b8282616dc58161a869565b935060ff1681518110616dda57616dda61a6fc565b60200260200101819052505b60e08501516060015115616e905760408051808201909152601681527f2d2d6d61785072696f726974794665655065724761730000000000000000000060208201528282616e338161a869565b935060ff1681518110616e4857616e4861a6fc565b6020026020010181905250616e648560e0015160600151617af7565b8282616e6f8161a869565b935060ff1681518110616e8457616e8461a6fc565b60200260200101819052505b60008160ff1667ffffffffffffffff811115616eae57616eae619d91565b604051908082528060200260200182016040528015616ee157816020015b6060815260200190600190039081616ecc5790505b50905060005b8260ff168160ff161015616f3a57838160ff1681518110616f0a57616f0a61a6fc565b6020026020010151828260ff1681518110616f2757616f2761a6fc565b6020908102919091010152600101616ee7565b5093505050505b949350505050565b616f706040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c91616ff69186910161a8f3565b600060405180830381865afa158015617013573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261703b919081019061a396565b905060006170498683618686565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b815260040161707991906199b6565b6000604051808303816000875af1158015617098573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526170c0919081019061a93a565b805190915060030b158015906170d95750602081015151155b80156170e85750604081015151155b156160c257816000815181106171005761710061a6fc565b602002602001015160405160200161571b919061a9f0565b6060600061714d8560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925286518252808701908201529091506171849082905b906187db565b156172e1576000617201826171fb846171f56171c78a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b90618802565b90618864565b604080518082018252600181527f0a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506172659082906187db565b156172cf57604080518082018252600181527f0a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526172cc905b82906188e9565b90505b6172d88161890f565b92505050615648565b82156172fa57848460405160200161571b92919061abdc565b5050604080516020810190915260008152615648565b509392505050565b6000808251602084016000f09392505050565b8160a001511561733a57505050565b6000617347848484618978565b9050600061735482616f49565b602081015181519192509060030b1580156173f05750604080518082018252600781527f5355434345535300000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526173f09060408051808201825260008082526020918201528151808301909252845182528085019082015261717e565b156173fd57505050505050565b6040820151511561741d57816040015160405160200161571b919061ac83565b8060405160200161571b919061ace1565b606060006174638360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506174c8905b8290617a96565b1561753757604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261564890617532908390618f13565b61890f565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617599905b8290618f9d565b60010361766657604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526175ff906172c5565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261564890617532905b83906188e9565b604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526176c5906174c1565b156177fc57604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061772d908390619037565b9050600081600183516177409190619fa6565b815181106177505761775061a6fc565b602002602001015190506177f36175326177c66040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252855182528086019082015290618f13565b95945050505050565b8260405160200161571b919061ad4c565b50919050565b606060006178488360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506178aa906174c1565b156178b8576156488161890f565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261791790617592565b60010361798157604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615648906175329061765f565b604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526179e0906174c1565b156177fc57604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290617a48908390619037565b9050600181511115617a84578060028251617a639190619fa6565b81518110617a7357617a7361a6fc565b602002602001015192505050919050565b508260405160200161571b919061ad4c565b805182516000911115617aab575060006155ce565b81518351602085015160009291617ac191619af3565b617acb9190619fa6565b905082602001518103617ae25760019150506155ce565b82516020840151819020912014905092915050565b60606000617b04836190dc565b600101905060008167ffffffffffffffff811115617b2457617b24619d91565b6040519080825280601f01601f191660200182016040528015617b4e576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084617b5857509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e5345440000000000000000000000000000000000000000000081840190815285518087018752838152840192909252845180860190955251845290830152606091617c23905b82906191be565b15617c6357505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e7365000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617cc290617c1c565b15617d0257505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d4954000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617d6190617c1c565b15617da157505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c79000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617e0090617c1c565b80617e655750604080518082018252601081527f47504c2d322e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617e6590617c1c565b15617ea557505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c79000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617f0490617c1c565b80617f695750604080518082018252601081527f47504c2d332e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617f6990617c1c565b15617fa957505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261800890617c1c565b8061806d5750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261806d90617c1c565b156180ad57505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261810c90617c1c565b806181715750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261817190617c1c565b156181b157505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261821090617c1c565b1561825057505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c617573650000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526182af90617c1c565b156182ef57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261834e90617c1c565b1561838e57505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e3000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526183ed90617c1c565b1561842d57505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261848c90617c1c565b156184cc57505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261852b90617c1c565b806185905750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261859090617c1c565b156185d057505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261862f90617c1c565b1561866f57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b6040808401518451915161571b929060200161ae2a565b60608060005b845181101561871157818582815181106186a8576186a861a6fc565b60200260200101516040516020016186c192919061a275565b6040516020818303038152906040529150600185516186e09190619fa6565b811461870957816040516020016186f7919061af93565b60405160208183030381529060405291505b60010161868c565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161872a57905050905083816000815181106187555761875561a6fc565b60200260200101819052506040518060400160405280600281526020017f2d63000000000000000000000000000000000000000000000000000000000000815250816001815181106187a9576187a961a6fc565b602002602001018190525081816002815181106187c8576187c861a6fc565b6020908102919091010152949350505050565b60208083015183518351928401516000936187f992918491906191d2565b14159392505050565b6040805180820190915260008082526020820152600061883484600001518560200151856000015186602001516192e3565b90508360200151816188469190619fa6565b84518590618855908390619fa6565b90525060208401525090919050565b60408051808201909152600080825260208201528151835110156188895750816155ce565b60208083015190840151600191146188b05750815160208481015190840151829020919020145b80156188e1578251845185906188c7908390619fa6565b90525082516020850180516188dd908390619af3565b9052505b509192915050565b6040805180820190915260008082526020820152618908838383619403565b5092915050565b60606000826000015167ffffffffffffffff81111561893057618930619d91565b6040519080825280601f01601f19166020018201604052801561895a576020820181803683370190505b509050600060208201905061890881856020015186600001516194ae565b60606000618984615941565b6040805160ff808252612000820190925291925060009190816020015b60608152602001906001900390816189a157905050905060006040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806189fc9061a869565b935060ff1681518110618a1157618a1161a6fc565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e3300000000000000000000000000000000000000000000000000815250604051602001618a62919061afd4565b604051602081830303815290604052828280618a7d9061a869565b935060ff1681518110618a9257618a9261a6fc565b60200260200101819052506040518060400160405280600881526020017f76616c6964617465000000000000000000000000000000000000000000000000815250828280618adf9061a869565b935060ff1681518110618af457618af461a6fc565b602002602001018190525082604051602001618b10919061a797565b604051602081830303815290604052828280618b2b9061a869565b935060ff1681518110618b4057618b4061a6fc565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e747261637400000000000000000000000000000000000000000000815250828280618b8d9061a869565b935060ff1681518110618ba257618ba261a6fc565b6020026020010181905250618bb78784619528565b8282618bc28161a869565b935060ff1681518110618bd757618bd761a6fc565b602090810291909101015285515115618c835760408051808201909152600b81527f2d2d7265666572656e636500000000000000000000000000000000000000000060208201528282618c298161a869565b935060ff1681518110618c3e57618c3e61a6fc565b6020026020010181905250618c57866000015184619528565b8282618c628161a869565b935060ff1681518110618c7757618c7761a6fc565b60200260200101819052505b856080015115618cf15760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b000000000000000060208201528282618ccc8161a869565b935060ff1681518110618ce157618ce161a6fc565b6020026020010181905250618d57565b8415618d575760408051808201909152601281527f2d2d726571756972655265666572656e6365000000000000000000000000000060208201528282618d368161a869565b935060ff1681518110618d4b57618d4b61a6fc565b60200260200101819052505b60408601515115618df35760408051808201909152600d81527f2d2d756e73616665416c6c6f770000000000000000000000000000000000000060208201528282618da18161a869565b935060ff1681518110618db657618db661a6fc565b60200260200101819052508560400151828280618dd29061a869565b935060ff1681518110618de757618de761a6fc565b60200260200101819052505b856060015115618e5d5760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d657300000000000000000000000060208201528282618e3c8161a869565b935060ff1681518110618e5157618e5161a6fc565b60200260200101819052505b60008160ff1667ffffffffffffffff811115618e7b57618e7b619d91565b604051908082528060200260200182016040528015618eae57816020015b6060815260200190600190039081618e995790505b50905060005b8260ff168160ff161015618f0757838160ff1681518110618ed757618ed761a6fc565b6020026020010151828260ff1681518110618ef457618ef461a6fc565b6020908102919091010152600101618eb4565b50979650505050505050565b6040805180820190915260008082526020820152815183511015618f385750816155ce565b81518351602085015160009291618f4e91619af3565b618f589190619fa6565b60208401519091506001908214618f79575082516020840151819020908220145b8015618f9457835185518690618f90908390619fa6565b9052505b50929392505050565b6000808260000151618fc185600001518660200151866000015187602001516192e3565b618fcb9190619af3565b90505b83516020850151618fdf9190619af3565b81116189085781618fef8161b019565b925050826000015161902685602001518361900a9190619fa6565b86516190169190619fa6565b83866000015187602001516192e3565b6190309190619af3565b9050618fce565b606060006190458484618f9d565b619050906001619af3565b67ffffffffffffffff81111561906857619068619d91565b60405190808252806020026020018201604052801561909b57816020015b60608152602001906001900390816190865790505b50905060005b8151811015617310576190b761753286866188e9565b8282815181106190c9576190c961a6fc565b60209081029190910101526001016190a1565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310619125577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310619151576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061916f57662386f26fc10000830492506010015b6305f5e1008310619187576305f5e100830492506008015b612710831061919b57612710830492506004015b606483106191ad576064830492506002015b600a83106155ce5760010192915050565b60006191ca8383619568565b159392505050565b6000808584116192d95760208411619285576000841561921d5760016191f9866020619fa6565b61920490600861b033565b61920f90600261b131565b6192199190619fa6565b1990505b835181168561922c8989619af3565b6192369190619fa6565b805190935082165b818114619270578784116192585787945050505050616f41565b836192628161b13d565b94505082845116905061923e565b61927a8785619af3565b945050505050616f41565b8383206192928588619fa6565b61929c9087619af3565b91505b8582106192d7578482208082036192c4576192ba8684619af3565b9350505050616f41565b6192cf600184619fa6565b92505061929f565b505b5092949350505050565b600083818685116193ee576020851161939d576000851561932f57600161930b876020619fa6565b61931690600861b033565b61932190600261b131565b61932b9190619fa6565b1990505b845181166000876193408b8b619af3565b61934a9190619fa6565b855190915083165b82811461938f578186106193775761936a8b8b619af3565b9650505050505050616f41565b856193818161b019565b965050838651169050619352565b859650505050505050616f41565b508383206000905b6193af8689619fa6565b82116193ec578583208082036193cb5783945050505050616f41565b6193d6600185619af3565b93505081806193e49061b019565b9250506193a5565b505b6193f88787619af3565b979650505050505050565b6040805180820190915260008082526020820152600061943585600001518660200151866000015187602001516192e3565b6020808701805191860191909152519091506194519082619fa6565b8352845160208601516194649190619af3565b810361947357600085526194a5565b835183516194819190619af3565b85518690619490908390619fa6565b905250835161949f9082619af3565b60208601525b50909392505050565b602081106194e657815183526194c5602084619af3565b92506194d2602083619af3565b91506194df602082619fa6565b90506194ae565b60001981156195155760016194fc836020619fa6565b6195089061010061b131565b6195129190619fa6565b90505b9151835183169219169190911790915250565b606060006195368484615a14565b80516020808301516040519394506195509390910161b154565b60405160208183030381529060405291505092915050565b815181516000919081111561957b575081515b6020808501519084015160005b8381101561963457825182518082146196045760001960208710156195e3576001846195b5896020619fa6565b6195bf9190619af3565b6195ca90600861b033565b6195d590600261b131565b6195df9190619fa6565b1990505b81811683821681810391146196015797506155ce9650505050505050565b50505b61960f602086619af3565b945061961c602085619af3565b9350505060208161962d9190619af3565b9050619588565b50845186516160c2919061b1ac565b610c9f8061b1cd83390190565b6112a68061be6c83390190565b6040518060e001604052806060815260200160608152602001606081526020016000151581526020016000151581526020016000151581526020016196a06196a5565b905290565b604051806101000160405280600015158152602001600015158152602001606081526020016000801916815260200160608152602001606081526020016000151581526020016196a06040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b818110156197575783516001600160a01b0316835260209384019390920191600101619730565b509095945050505050565b60005b8381101561977d578181015183820152602001619765565b50506000910152565b6000815180845261979e816020860160208601619762565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156198ae577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b81811015619894577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261987e848651619786565b6020958601959094509290920191600101619844565b5091975050506020948501949290920191506001016197da565b50929695505050505050565b600081518084526020840193506020830160005b8281101561990e5781517fffffffff00000000000000000000000000000000000000000000000000000000168652602095860195909101906001016198ce565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156198ae577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051604087526199846040880182619786565b905060208201519150868103602088015261999f81836198ba565b965050506020938401939190910190600101619940565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156198ae577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452619a18858351619786565b945060209384019391909101906001016199de565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156198ae577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b0381511686526020810151905060406020870152619aae60408701826198ba565b9550506020938401939190910190600101619a55565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156155ce576155ce619ac4565b6020815260006156486020830184619786565b6001600160a01b0381511682526020810151151560208301526001600160a01b0360408201511660408301526000606082015160a06060850152619b6060a0850182619786565b608093840151949093019390935250919050565b6001600160a01b0384168152606060208201526000619b966060830185619786565b82810360408401526160c28185619b19565b600181811c90821680619bbc57607f821691505b60208210810361780d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600081546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160a0606085015260008154619c3f81619ba8565b8060a08801526001821660008114619c5e5760018114619c9857619ccc565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660c089015260c082151560051b8901019350619ccc565b84600052602060002060005b83811015619cc35781548a820160c00152600190910190602001619ca4565b890160c0019450505b50505060038401546080860152809250505092915050565b8381526001600160a01b0383166020820152608060408201526000608082015260a0606082015260006177f360a0830184619bf5565b6001600160a01b0383168152604060208201526000616f416040830184619bf5565b6001600160a01b03861681528460208201526001600160a01b038416604082015260a060608201526000619d7360a0830185619786565b8281036080840152619d858185619bf5565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561082a57806000526020600020601f840160051c81016020851015619de75750805b601f840160051c820191505b818110156121d15760008155600101619df3565b815167ffffffffffffffff811115619e2157619e21619d91565b619e3581619e2f8454619ba8565b84619dc0565b6020601f821160018114619e695760008315619e515750848201515b600019600385901b1c1916600184901b1784556121d1565b600084815260208120601f198516915b82811015619e995787850151825560209485019460019092019101619e79565b5084821015619eb75786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b600060208284031215619ed857600080fd5b5051919050565b600060208284031215619ef157600080fd5b8151801515811461564857600080fd5b6001600160a01b03851681528360208201526001600160a01b03831660408201526080606082015260006160c26080830184619b19565b6001600160a01b0384168152606060208201526000619f5a6060830185619786565b82810360408401526160c28185619bf5565b8481526001600160a01b0384166020820152608060408201526000619f946080830185619786565b82810360608401526193f88185619bf5565b818103818111156155ce576155ce619ac4565b6001600160a01b03851681528360208201526001600160a01b03831660408201526080606082015260006160c26080830184619bf5565b6001600160a01b03861681528460208201526001600160a01b038416604082015260a06060820152600061a02760a0830185619786565b8281036080840152619d858185619b19565b60008261a06f577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60408152600061a0876040830185619786565b82810360208401526156448185619bf5565b6001600160a01b0383168152604060208201526000616f416040830184619b19565b6001600160a01b0383168152604060208201526000616f416040830184619786565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161a11581601a850160208801619762565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161a15281601c840160208801619762565b01601c01949350505050565b60006020828403121561a17057600080fd5b81516001600160a01b038116811461564857600080fd5b6040516060810167ffffffffffffffff8111828210171561a1aa5761a1aa619d91565b60405290565b60008067ffffffffffffffff84111561a1cb5761a1cb619d91565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561a1fa5761a1fa619d91565b60405283815290508082840185101561a21257600080fd5b617310846020830185619762565b600082601f83011261a23157600080fd5b6156488383516020850161a1b0565b60006020828403121561a25257600080fd5b815167ffffffffffffffff81111561a26957600080fd5b6155ca8482850161a220565b6000835161a287818460208801619762565b83519083019061a29b818360208801619762565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161a2dc81601a850160208801619762565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161a319816033840160208801619762565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f555400000000000000000000000000000000000000000060608201526080602082015260006156486080830184619786565b60006020828403121561a3a857600080fd5b815167ffffffffffffffff81111561a3bf57600080fd5b8201601f8101841361a3d057600080fd5b6155ca8482516020840161a1b0565b6000855161a3f1818460208a01619762565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161a42b816001840160208a01619762565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161a469816002840160208901619762565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161a4ab816002840160208801619762565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061a4f66040830184619786565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161a56d81601f850160208701619762565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061a5da6040830184619786565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061a62c6040830184619786565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161a6a3816014850160208701619762565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061a6ea6040830185619786565b82810360208401526156448185619786565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161a763816001850160208701619762565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161a7a9818460208701619762565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161a85c81604b850160208701619762565b91909101604b0192915050565b600060ff821660ff810361a87f5761a87f619ac4565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161a8e6816029850160208701619762565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f504154480000000000000000000060608201526080602082015260006156486080830184619786565b60006020828403121561a94c57600080fd5b815167ffffffffffffffff81111561a96357600080fd5b82016060818503121561a97557600080fd5b61a97d61a187565b81518060030b811461a98e57600080fd5b8152602082015167ffffffffffffffff81111561a9aa57600080fd5b61a9b68682850161a220565b602083015250604082015167ffffffffffffffff81111561a9d657600080fd5b61a9e28682850161a220565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161aa4e816021850160208701619762565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161ac3a816021850160208801619762565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161ac7781602e840160208801619762565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161a8e6816029850160208701619762565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161ad3f816022850160208701619762565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161ad8481600e850160208701619762565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161ae62816018850160208801619762565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161ae9f81601c840160208801619762565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161afa5818460208701619762565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161b00c81601c850160208701619762565b91909101601c0192915050565b6000600019820361b02c5761b02c619ac4565b5060010190565b80820281158282048414176155ce576155ce619ac4565b6001815b600184111561b0855780850481111561b0695761b069619ac4565b600184161561b07757908102905b60019390931c92800261b04e565b935093915050565b60008261b09c575060016155ce565b8161b0a9575060006155ce565b816001811461b0bf576002811461b0c95761b0e5565b60019150506155ce565b60ff84111561b0da5761b0da619ac4565b50506001821b6155ce565b5060208310610133831016604e8410600b841016171561b108575081810a6155ce565b61b115600019848461b04a565b806000190482111561b1295761b129619ac4565b029392505050565b6000615648838361b08d565b60008161b14c5761b14c619ac4565b506000190190565b6000835161b166818460208801619762565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161b1a0816001840160208801619762565b01600101949350505050565b818103600083128015838313168383128216171561890857618908619ac456fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220837c7d9916de10b20cdb18567d8c7679613426bbd7b0b72548d8000a412f307b64736f6c634300081a0033608060405234801561001057600080fd5b506040516112a63803806112a683398101604081905261002f91610110565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007891906101e2565b50600461008582826101e2565b5050506001600160a01b03821615806100a557506001600160a01b038116155b156100c35760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b031991821617909155600780549290931691161790556102a0565b80516001600160a01b038116811461010b57600080fd5b919050565b6000806040838503121561012357600080fd5b61012c836100f4565b915061013a602084016100f4565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061016d57607f821691505b60208210810361018d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101dd57806000526020600020601f840160051c810160208510156101ba5750805b601f840160051c820191505b818110156101da57600081556001016101c6565b50505b505050565b81516001600160401b038111156101fb576101fb610143565b61020f816102098454610159565b84610193565b6020601f821160018114610243576000831561022b5750848201515b600019600385901b1c1916600184901b1784556101da565b600084815260208120601f198516915b828110156102735787850151825560209485019460019092019101610253565b50848210156102915786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610ff7806102af6000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806342966c68116100b257806379cc679011610081578063a9059cbb11610066578063a9059cbb1461028e578063bff9662a146102a1578063dd62ed3e146102c157600080fd5b806379cc67901461027357806395d89b411461028657600080fd5b806342966c68146102025780635b1125911461021557806370a0823114610235578063779e3b631461026b57600080fd5b80631e458bee116100ee5780631e458bee1461018857806323b872dd1461019b578063313ce567146101ae578063328a01d0146101bd57600080fd5b806306fdde0314610120578063095ea7b31461013e57806315d57fd41461016157806318160ddd14610176575b600080fd5b610128610307565b6040516101359190610d97565b60405180910390f35b61015161014c366004610e2c565b610399565b6040519015158152602001610135565b61017461016f366004610e56565b6103b3565b005b6002545b604051908152602001610135565b610174610196366004610e89565b61057e565b6101516101a9366004610ebc565b610631565b60405160128152602001610135565b6007546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610135565b610174610210366004610ef9565b610655565b6006546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a610243366004610f12565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b610174610662565b610174610281366004610e2c565b610786565b610128610837565b61015161029c366004610e2c565b610846565b6005546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a6102cf366004610e56565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60606003805461031690610f34565b80601f016020809104026020016040519081016040528092919081815260200182805461034290610f34565b801561038f5780601f106103645761010080835404028352916020019161038f565b820191906000526020600020905b81548152906001019060200180831161037257829003601f168201915b5050505050905090565b6000336103a7818585610854565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103f3575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610431576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82161580610468575073ffffffffffffffffffffffffffffffffffffffff8116155b1561049f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105d1576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6105db8383610866565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161062491815260200190565b60405180910390a3505050565b60003361063f8582856108c6565b61064a858585610995565b506001949350505050565b61065f3382610a40565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106b5576040517fe700765e000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b60065473ffffffffffffffffffffffffffffffffffffffff16610704576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107d9576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6107e38282610a9c565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161082b91815260200190565b60405180910390a25050565b60606004805461031690610f34565b6000336103a7818585610995565b6108618383836001610ab1565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108b6576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c260008383610bf9565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461098f5781811015610980576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610428565b61098f84848484036000610ab1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109e5576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8216610a35576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b610861838383610bf9565b73ffffffffffffffffffffffffffffffffffffffff8216610a90576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c282600083610bf9565b610aa78233836108c6565b6108c28282610a40565b73ffffffffffffffffffffffffffffffffffffffff8416610b01576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8316610b51576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600160209081526040808320938716835292905220829055801561098f578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610beb91815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c31578060026000828254610c269190610f87565b90915550610ce39050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610cb7576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610428565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610d0c57600280548290039055610d38565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161062491815260200190565b602081526000825180602084015260005b81811015610dc55760208186018101516040868401015201610da8565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610e2757600080fd5b919050565b60008060408385031215610e3f57600080fd5b610e4883610e03565b946020939093013593505050565b60008060408385031215610e6957600080fd5b610e7283610e03565b9150610e8060208401610e03565b90509250929050565b600080600060608486031215610e9e57600080fd5b610ea784610e03565b95602085013595506040909401359392505050565b600080600060608486031215610ed157600080fd5b610eda84610e03565b9250610ee860208501610e03565b929592945050506040919091013590565b600060208284031215610f0b57600080fd5b5035919050565b600060208284031215610f2457600080fd5b610f2d82610e03565b9392505050565b600181811c90821680610f4857607f821691505b602082108103610f81577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b808201808211156103ad577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220fec2a89bc2f37ba9f63e9d7e47fa9a0d2b02402309f8905f94cc44fb9e5fbba064736f6c634300081a0033a264697066735822122099911d716e5475beeb8a146016592cf6a156b58d00b97bca2e662667a814057c64736f6c634300081a0033",
}

// GatewayEVMInboundTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMInboundTestMetaData.ABI instead.
var GatewayEVMInboundTestABI = GatewayEVMInboundTestMetaData.ABI

// GatewayEVMInboundTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMInboundTestMetaData.Bin instead.
var GatewayEVMInboundTestBin = GatewayEVMInboundTestMetaData.Bin

// DeployGatewayEVMInboundTest deploys a new Ethereum contract, binding an instance of GatewayEVMInboundTest to it.
func DeployGatewayEVMInboundTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMInboundTest, error) {
	parsed, err := GatewayEVMInboundTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMInboundTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMInboundTest{GatewayEVMInboundTestCaller: GatewayEVMInboundTestCaller{contract: contract}, GatewayEVMInboundTestTransactor: GatewayEVMInboundTestTransactor{contract: contract}, GatewayEVMInboundTestFilterer: GatewayEVMInboundTestFilterer{contract: contract}}, nil
}

// GatewayEVMInboundTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMInboundTest struct {
	GatewayEVMInboundTestCaller     // Read-only binding to the contract
	GatewayEVMInboundTestTransactor // Write-only binding to the contract
	GatewayEVMInboundTestFilterer   // Log filterer for contract events
}

// GatewayEVMInboundTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMInboundTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMInboundTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMInboundTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMInboundTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMInboundTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMInboundTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMInboundTestSession struct {
	Contract     *GatewayEVMInboundTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayEVMInboundTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMInboundTestCallerSession struct {
	Contract *GatewayEVMInboundTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// GatewayEVMInboundTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMInboundTestTransactorSession struct {
	Contract     *GatewayEVMInboundTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// GatewayEVMInboundTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMInboundTestRaw struct {
	Contract *GatewayEVMInboundTest // Generic contract binding to access the raw methods on
}

// GatewayEVMInboundTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMInboundTestCallerRaw struct {
	Contract *GatewayEVMInboundTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMInboundTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMInboundTestTransactorRaw struct {
	Contract *GatewayEVMInboundTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMInboundTest creates a new instance of GatewayEVMInboundTest, bound to a specific deployed contract.
func NewGatewayEVMInboundTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMInboundTest, error) {
	contract, err := bindGatewayEVMInboundTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTest{GatewayEVMInboundTestCaller: GatewayEVMInboundTestCaller{contract: contract}, GatewayEVMInboundTestTransactor: GatewayEVMInboundTestTransactor{contract: contract}, GatewayEVMInboundTestFilterer: GatewayEVMInboundTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMInboundTestCaller creates a new read-only instance of GatewayEVMInboundTest, bound to a specific deployed contract.
func NewGatewayEVMInboundTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMInboundTestCaller, error) {
	contract, err := bindGatewayEVMInboundTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestCaller{contract: contract}, nil
}

// NewGatewayEVMInboundTestTransactor creates a new write-only instance of GatewayEVMInboundTest, bound to a specific deployed contract.
func NewGatewayEVMInboundTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMInboundTestTransactor, error) {
	contract, err := bindGatewayEVMInboundTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestTransactor{contract: contract}, nil
}

// NewGatewayEVMInboundTestFilterer creates a new log filterer instance of GatewayEVMInboundTest, bound to a specific deployed contract.
func NewGatewayEVMInboundTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMInboundTestFilterer, error) {
	contract, err := bindGatewayEVMInboundTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestFilterer{contract: contract}, nil
}

// bindGatewayEVMInboundTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMInboundTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMInboundTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMInboundTest.Contract.GatewayEVMInboundTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.GatewayEVMInboundTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.GatewayEVMInboundTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMInboundTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ISTEST() (bool, error) {
	return _GatewayEVMInboundTest.Contract.ISTEST(&_GatewayEVMInboundTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ISTEST() (bool, error) {
	return _GatewayEVMInboundTest.Contract.ISTEST(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeArtifacts(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeArtifacts(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeContracts(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeContracts(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeSenders(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeSenders(&_GatewayEVMInboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) Failed() (bool, error) {
	return _GatewayEVMInboundTest.Contract.Failed(&_GatewayEVMInboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) Failed() (bool, error) {
	return _GatewayEVMInboundTest.Contract.Failed(&_GatewayEVMInboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMInboundTest.Contract.TargetArtifactSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMInboundTest.Contract.TargetArtifactSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMInboundTest.Contract.TargetArtifacts(&_GatewayEVMInboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMInboundTest.Contract.TargetArtifacts(&_GatewayEVMInboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.TargetContracts(&_GatewayEVMInboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.TargetContracts(&_GatewayEVMInboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMInboundTest.Contract.TargetInterfaces(&_GatewayEVMInboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMInboundTest.Contract.TargetInterfaces(&_GatewayEVMInboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMInboundTest.Contract.TargetSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMInboundTest.Contract.TargetSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.TargetSenders(&_GatewayEVMInboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.TargetSenders(&_GatewayEVMInboundTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.SetUp(&_GatewayEVMInboundTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.SetUp(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallFailsIfRevertGasLimitExceeded")
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayload is a paid mutator transaction binding the contract method 0xaa030c1c.
//
// Solidity: function testCallWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallWithPayload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallWithPayload")
}

// TestCallWithPayload is a paid mutator transaction binding the contract method 0xaa030c1c.
//
// Solidity: function testCallWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayload is a paid mutator transaction binding the contract method 0xaa030c1c.
//
// Solidity: function testCallWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfCallOnRevertIsTrue is a paid mutator transaction binding the contract method 0xba46ba23.
//
// Solidity: function testCallWithPayloadFailsIfCallOnRevertIsTrue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallWithPayloadFailsIfCallOnRevertIsTrue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallWithPayloadFailsIfCallOnRevertIsTrue")
}

// TestCallWithPayloadFailsIfCallOnRevertIsTrue is a paid mutator transaction binding the contract method 0xba46ba23.
//
// Solidity: function testCallWithPayloadFailsIfCallOnRevertIsTrue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallWithPayloadFailsIfCallOnRevertIsTrue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfCallOnRevertIsTrue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfCallOnRevertIsTrue is a paid mutator transaction binding the contract method 0xba46ba23.
//
// Solidity: function testCallWithPayloadFailsIfCallOnRevertIsTrue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallWithPayloadFailsIfCallOnRevertIsTrue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfCallOnRevertIsTrue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x51da903d.
//
// Solidity: function testCallWithPayloadFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallWithPayloadFailsIfDestinationIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallWithPayloadFailsIfDestinationIsZeroAddress")
}

// TestCallWithPayloadFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x51da903d.
//
// Solidity: function testCallWithPayloadFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallWithPayloadFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfDestinationIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x51da903d.
//
// Solidity: function testCallWithPayloadFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallWithPayloadFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfDestinationIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xb1409f71.
//
// Solidity: function testCallWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallWithPayloadFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallWithPayloadFailsIfPayloadSizeExceeded")
}

// TestCallWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xb1409f71.
//
// Solidity: function testCallWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xb1409f71.
//
// Solidity: function testCallWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x458085f5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded")
}

// TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x458085f5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x458085f5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc0fab86d.
//
// Solidity: function testDepositAndCallEthFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthFailsIfRevertGasLimitExceeded")
}

// TestDepositAndCallEthFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc0fab86d.
//
// Solidity: function testDepositAndCallEthFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc0fab86d.
//
// Solidity: function testDepositAndCallEthFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustody is a paid mutator transaction binding the contract method 0x6459542a.
//
// Solidity: function testDepositERC20ToCustody() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustody(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustody")
}

// TestDepositERC20ToCustody is a paid mutator transaction binding the contract method 0x6459542a.
//
// Solidity: function testDepositERC20ToCustody() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustody() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustody(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustody is a paid mutator transaction binding the contract method 0x6459542a.
//
// Solidity: function testDepositERC20ToCustody() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustody() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustody(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x4ce25c0a.
//
// Solidity: function testDepositERC20ToCustodyFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfPayloadSizeExceeded")
}

// TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x4ce25c0a.
//
// Solidity: function testDepositERC20ToCustodyFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x4ce25c0a.
//
// Solidity: function testDepositERC20ToCustodyFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x1f4f542f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded")
}

// TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x1f4f542f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x1f4f542f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xb2849063.
//
// Solidity: function testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted")
}

// TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xb2849063.
//
// Solidity: function testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xb2849063.
//
// Solidity: function testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayload is a paid mutator transaction binding the contract method 0x30f7c04f.
//
// Solidity: function testDepositERC20ToCustodyWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayload")
}

// TestDepositERC20ToCustodyWithPayload is a paid mutator transaction binding the contract method 0x30f7c04f.
//
// Solidity: function testDepositERC20ToCustodyWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayload is a paid mutator transaction binding the contract method 0x30f7c04f.
//
// Solidity: function testDepositERC20ToCustodyWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x95cd0445.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded")
}

// TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x95cd0445.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x95cd0445.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xe85c5a07.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted")
}

// TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xe85c5a07.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xe85c5a07.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTss is a paid mutator transaction binding the contract method 0x0724d8e3.
//
// Solidity: function testDepositEthToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTss(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTss")
}

// TestDepositEthToTss is a paid mutator transaction binding the contract method 0x0724d8e3.
//
// Solidity: function testDepositEthToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTss() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTss(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTss is a paid mutator transaction binding the contract method 0x0724d8e3.
//
// Solidity: function testDepositEthToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTss() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTss(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xb1526934.
//
// Solidity: function testDepositEthToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssFailsIfRevertGasLimitExceeded")
}

// TestDepositEthToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xb1526934.
//
// Solidity: function testDepositEthToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xb1526934.
//
// Solidity: function testDepositEthToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayload is a paid mutator transaction binding the contract method 0x9fd1e597.
//
// Solidity: function testDepositEthToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssWithPayload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssWithPayload")
}

// TestDepositEthToTssWithPayload is a paid mutator transaction binding the contract method 0x9fd1e597.
//
// Solidity: function testDepositEthToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayload is a paid mutator transaction binding the contract method 0x9fd1e597.
//
// Solidity: function testDepositEthToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x466f332e.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded")
}

// TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x466f332e.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x466f332e.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositZetaToConnector is a paid mutator transaction binding the contract method 0xe306a978.
//
// Solidity: function testDepositZetaToConnector() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositZetaToConnector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositZetaToConnector")
}

// TestDepositZetaToConnector is a paid mutator transaction binding the contract method 0xe306a978.
//
// Solidity: function testDepositZetaToConnector() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositZetaToConnector() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositZetaToConnector(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositZetaToConnector is a paid mutator transaction binding the contract method 0xe306a978.
//
// Solidity: function testDepositZetaToConnector() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositZetaToConnector() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositZetaToConnector(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyIfAmountIs0 is a paid mutator transaction binding the contract method 0xcdef966f.
//
// Solidity: function testRevertDepositERC20ToCustodyIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositERC20ToCustodyIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositERC20ToCustodyIfAmountIs0")
}

// TestRevertDepositERC20ToCustodyIfAmountIs0 is a paid mutator transaction binding the contract method 0xcdef966f.
//
// Solidity: function testRevertDepositERC20ToCustodyIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositERC20ToCustodyIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyIfAmountIs0 is a paid mutator transaction binding the contract method 0xcdef966f.
//
// Solidity: function testRevertDepositERC20ToCustodyIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositERC20ToCustodyIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x41b8fec6.
//
// Solidity: function testRevertDepositERC20ToCustodyIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositERC20ToCustodyIfReceiverIsZeroAddress")
}

// TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x41b8fec6.
//
// Solidity: function testRevertDepositERC20ToCustodyIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x41b8fec6.
//
// Solidity: function testRevertDepositERC20ToCustodyIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0 is a paid mutator transaction binding the contract method 0x18e14138.
//
// Solidity: function testRevertDepositERC20ToCustodyWithPayloadIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositERC20ToCustodyWithPayloadIfAmountIs0")
}

// TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0 is a paid mutator transaction binding the contract method 0x18e14138.
//
// Solidity: function testRevertDepositERC20ToCustodyWithPayloadIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0 is a paid mutator transaction binding the contract method 0x18e14138.
//
// Solidity: function testRevertDepositERC20ToCustodyWithPayloadIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x09b21ccb.
//
// Solidity: function testRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress")
}

// TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x09b21ccb.
//
// Solidity: function testRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x09b21ccb.
//
// Solidity: function testRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfAmountIs0 is a paid mutator transaction binding the contract method 0x8f5a4224.
//
// Solidity: function testRevertDepositEthToTssIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositEthToTssIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositEthToTssIfAmountIs0")
}

// TestRevertDepositEthToTssIfAmountIs0 is a paid mutator transaction binding the contract method 0x8f5a4224.
//
// Solidity: function testRevertDepositEthToTssIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositEthToTssIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfAmountIs0 is a paid mutator transaction binding the contract method 0x8f5a4224.
//
// Solidity: function testRevertDepositEthToTssIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositEthToTssIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x0fc37f5e.
//
// Solidity: function testRevertDepositEthToTssIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositEthToTssIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositEthToTssIfPayloadSizeExceeded")
}

// TestRevertDepositEthToTssIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x0fc37f5e.
//
// Solidity: function testRevertDepositEthToTssIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositEthToTssIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x0fc37f5e.
//
// Solidity: function testRevertDepositEthToTssIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositEthToTssIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x5c4013d0.
//
// Solidity: function testRevertDepositEthToTssIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositEthToTssIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositEthToTssIfReceiverIsZeroAddress")
}

// TestRevertDepositEthToTssIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x5c4013d0.
//
// Solidity: function testRevertDepositEthToTssIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositEthToTssIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x5c4013d0.
//
// Solidity: function testRevertDepositEthToTssIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositEthToTssIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssWithPayloadIfAmountIs0 is a paid mutator transaction binding the contract method 0x2a5dcf36.
//
// Solidity: function testRevertDepositEthToTssWithPayloadIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositEthToTssWithPayloadIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositEthToTssWithPayloadIfAmountIs0")
}

// TestRevertDepositEthToTssWithPayloadIfAmountIs0 is a paid mutator transaction binding the contract method 0x2a5dcf36.
//
// Solidity: function testRevertDepositEthToTssWithPayloadIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositEthToTssWithPayloadIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssWithPayloadIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssWithPayloadIfAmountIs0 is a paid mutator transaction binding the contract method 0x2a5dcf36.
//
// Solidity: function testRevertDepositEthToTssWithPayloadIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositEthToTssWithPayloadIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssWithPayloadIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9d96310a.
//
// Solidity: function testRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress")
}

// TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9d96310a.
//
// Solidity: function testRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9d96310a.
//
// Solidity: function testRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// GatewayEVMInboundTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestCalledIterator struct {
	Event *GatewayEVMInboundTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestCalled)
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
		it.Event = new(GatewayEVMInboundTestCalled)
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
func (it *GatewayEVMInboundTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestCalled represents a Called event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMInboundTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestCalledIterator{contract: _GatewayEVMInboundTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestCalled)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseCalled(log types.Log) (*GatewayEVMInboundTestCalled, error) {
	event := new(GatewayEVMInboundTestCalled)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestDepositedIterator struct {
	Event *GatewayEVMInboundTestDeposited // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestDeposited)
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
		it.Event = new(GatewayEVMInboundTestDeposited)
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
func (it *GatewayEVMInboundTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestDeposited represents a Deposited event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestDeposited struct {
	Sender        common.Address
	Receiver      common.Address
	Amount        *big.Int
	Asset         common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterDeposited(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMInboundTestDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestDepositedIterator{contract: _GatewayEVMInboundTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestDeposited, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestDeposited)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseDeposited(log types.Log) (*GatewayEVMInboundTestDeposited, error) {
	event := new(GatewayEVMInboundTestDeposited)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestDepositedAndCalledIterator is returned from FilterDepositedAndCalled and is used to iterate over the raw logs and unpacked data for DepositedAndCalled events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestDepositedAndCalledIterator struct {
	Event *GatewayEVMInboundTestDepositedAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestDepositedAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestDepositedAndCalled)
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
		it.Event = new(GatewayEVMInboundTestDepositedAndCalled)
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
func (it *GatewayEVMInboundTestDepositedAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestDepositedAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestDepositedAndCalled represents a DepositedAndCalled event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestDepositedAndCalled struct {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterDepositedAndCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMInboundTestDepositedAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestDepositedAndCalledIterator{contract: _GatewayEVMInboundTest.contract, event: "DepositedAndCalled", logs: logs, sub: sub}, nil
}

// WatchDepositedAndCalled is a free log subscription operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchDepositedAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestDepositedAndCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestDepositedAndCalled)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseDepositedAndCalled(log types.Log) (*GatewayEVMInboundTestDepositedAndCalled, error) {
	event := new(GatewayEVMInboundTestDepositedAndCalled)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestExecutedIterator struct {
	Event *GatewayEVMInboundTestExecuted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestExecuted)
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
		it.Event = new(GatewayEVMInboundTestExecuted)
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
func (it *GatewayEVMInboundTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestExecuted represents a Executed event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMInboundTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestExecutedIterator{contract: _GatewayEVMInboundTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestExecuted)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Executed", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseExecuted(log types.Log) (*GatewayEVMInboundTestExecuted, error) {
	event := new(GatewayEVMInboundTestExecuted)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMInboundTestExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestExecutedWithERC20)
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
		it.Event = new(GatewayEVMInboundTestExecutedWithERC20)
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
func (it *GatewayEVMInboundTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMInboundTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestExecutedWithERC20Iterator{contract: _GatewayEVMInboundTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestExecutedWithERC20)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMInboundTestExecutedWithERC20, error) {
	event := new(GatewayEVMInboundTestExecutedWithERC20)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedERC20Iterator struct {
	Event *GatewayEVMInboundTestReceivedERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedERC20)
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
		it.Event = new(GatewayEVMInboundTestReceivedERC20)
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
func (it *GatewayEVMInboundTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedERC20 represents a ReceivedERC20 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedERC20Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedERC20Iterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedERC20)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedERC20(log types.Log) (*GatewayEVMInboundTestReceivedERC20, error) {
	event := new(GatewayEVMInboundTestReceivedERC20)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedNoParamsIterator struct {
	Event *GatewayEVMInboundTestReceivedNoParams // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedNoParams)
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
		it.Event = new(GatewayEVMInboundTestReceivedNoParams)
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
func (it *GatewayEVMInboundTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedNoParams represents a ReceivedNoParams event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedNoParamsIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedNoParamsIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedNoParams)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedNoParams(log types.Log) (*GatewayEVMInboundTestReceivedNoParams, error) {
	event := new(GatewayEVMInboundTestReceivedNoParams)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedNonPayableIterator struct {
	Event *GatewayEVMInboundTestReceivedNonPayable // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedNonPayable)
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
		it.Event = new(GatewayEVMInboundTestReceivedNonPayable)
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
func (it *GatewayEVMInboundTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedNonPayable represents a ReceivedNonPayable event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedNonPayableIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedNonPayableIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedNonPayable)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedNonPayable(log types.Log) (*GatewayEVMInboundTestReceivedNonPayable, error) {
	event := new(GatewayEVMInboundTestReceivedNonPayable)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedOnCallIterator is returned from FilterReceivedOnCall and is used to iterate over the raw logs and unpacked data for ReceivedOnCall events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedOnCallIterator struct {
	Event *GatewayEVMInboundTestReceivedOnCall // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestReceivedOnCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedOnCall)
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
		it.Event = new(GatewayEVMInboundTestReceivedOnCall)
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
func (it *GatewayEVMInboundTestReceivedOnCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedOnCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedOnCall represents a ReceivedOnCall event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedOnCall struct {
	Sender  common.Address
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCall is a free log retrieval operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedOnCall(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedOnCallIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedOnCallIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedOnCall", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCall is a free log subscription operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedOnCall(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedOnCall) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedOnCall)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedOnCall(log types.Log) (*GatewayEVMInboundTestReceivedOnCall, error) {
	event := new(GatewayEVMInboundTestReceivedOnCall)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedPayableIterator struct {
	Event *GatewayEVMInboundTestReceivedPayable // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedPayable)
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
		it.Event = new(GatewayEVMInboundTestReceivedPayable)
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
func (it *GatewayEVMInboundTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedPayable represents a ReceivedPayable event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedPayable struct {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedPayableIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedPayableIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedPayable)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedPayable(log types.Log) (*GatewayEVMInboundTestReceivedPayable, error) {
	event := new(GatewayEVMInboundTestReceivedPayable)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedRevertIterator struct {
	Event *GatewayEVMInboundTestReceivedRevert // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedRevert)
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
		it.Event = new(GatewayEVMInboundTestReceivedRevert)
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
func (it *GatewayEVMInboundTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedRevert represents a ReceivedRevert event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedRevertIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedRevertIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedRevert)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedRevert(log types.Log) (*GatewayEVMInboundTestReceivedRevert, error) {
	event := new(GatewayEVMInboundTestReceivedRevert)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestRevertedIterator struct {
	Event *GatewayEVMInboundTestReverted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReverted)
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
		it.Event = new(GatewayEVMInboundTestReverted)
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
func (it *GatewayEVMInboundTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReverted represents a Reverted event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReverted struct {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMInboundTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestRevertedIterator{contract: _GatewayEVMInboundTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReverted)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Reverted", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReverted(log types.Log) (*GatewayEVMInboundTestReverted, error) {
	event := new(GatewayEVMInboundTestReverted)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator struct {
	Event *GatewayEVMInboundTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestUpdatedGatewayTSSAddress)
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
		it.Event = new(GatewayEVMInboundTestUpdatedGatewayTSSAddress)
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
func (it *GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedGatewayTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator{contract: _GatewayEVMInboundTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestUpdatedGatewayTSSAddress)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*GatewayEVMInboundTestUpdatedGatewayTSSAddress, error) {
	event := new(GatewayEVMInboundTestUpdatedGatewayTSSAddress)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogIterator struct {
	Event *GatewayEVMInboundTestLog // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLog)
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
		it.Event = new(GatewayEVMInboundTestLog)
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
func (it *GatewayEVMInboundTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLog represents a Log event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogIterator{contract: _GatewayEVMInboundTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLog)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLog(log types.Log) (*GatewayEVMInboundTestLog, error) {
	event := new(GatewayEVMInboundTestLog)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogAddressIterator struct {
	Event *GatewayEVMInboundTestLogAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogAddress)
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
		it.Event = new(GatewayEVMInboundTestLogAddress)
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
func (it *GatewayEVMInboundTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogAddress represents a LogAddress event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogAddressIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogAddressIterator{contract: _GatewayEVMInboundTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogAddress)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogAddress(log types.Log) (*GatewayEVMInboundTestLogAddress, error) {
	event := new(GatewayEVMInboundTestLogAddress)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArrayIterator struct {
	Event *GatewayEVMInboundTestLogArray // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogArray)
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
		it.Event = new(GatewayEVMInboundTestLogArray)
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
func (it *GatewayEVMInboundTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogArray represents a LogArray event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogArrayIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogArrayIterator{contract: _GatewayEVMInboundTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogArray)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogArray(log types.Log) (*GatewayEVMInboundTestLogArray, error) {
	event := new(GatewayEVMInboundTestLogArray)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray0Iterator struct {
	Event *GatewayEVMInboundTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogArray0)
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
		it.Event = new(GatewayEVMInboundTestLogArray0)
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
func (it *GatewayEVMInboundTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogArray0 represents a LogArray0 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogArray0Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogArray0)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogArray0(log types.Log) (*GatewayEVMInboundTestLogArray0, error) {
	event := new(GatewayEVMInboundTestLogArray0)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray1Iterator struct {
	Event *GatewayEVMInboundTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogArray1)
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
		it.Event = new(GatewayEVMInboundTestLogArray1)
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
func (it *GatewayEVMInboundTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogArray1 represents a LogArray1 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogArray1Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogArray1)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogArray1(log types.Log) (*GatewayEVMInboundTestLogArray1, error) {
	event := new(GatewayEVMInboundTestLogArray1)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogBytesIterator struct {
	Event *GatewayEVMInboundTestLogBytes // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogBytes)
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
		it.Event = new(GatewayEVMInboundTestLogBytes)
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
func (it *GatewayEVMInboundTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogBytes represents a LogBytes event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogBytesIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogBytesIterator{contract: _GatewayEVMInboundTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogBytes)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogBytes(log types.Log) (*GatewayEVMInboundTestLogBytes, error) {
	event := new(GatewayEVMInboundTestLogBytes)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogBytes32Iterator struct {
	Event *GatewayEVMInboundTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogBytes32)
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
		it.Event = new(GatewayEVMInboundTestLogBytes32)
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
func (it *GatewayEVMInboundTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogBytes32 represents a LogBytes32 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogBytes32Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogBytes32)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogBytes32(log types.Log) (*GatewayEVMInboundTestLogBytes32, error) {
	event := new(GatewayEVMInboundTestLogBytes32)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogIntIterator struct {
	Event *GatewayEVMInboundTestLogInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogInt)
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
		it.Event = new(GatewayEVMInboundTestLogInt)
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
func (it *GatewayEVMInboundTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogInt represents a LogInt event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogIntIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogIntIterator{contract: _GatewayEVMInboundTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogInt)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogInt(log types.Log) (*GatewayEVMInboundTestLogInt, error) {
	event := new(GatewayEVMInboundTestLogInt)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedAddressIterator struct {
	Event *GatewayEVMInboundTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedAddress)
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
		it.Event = new(GatewayEVMInboundTestLogNamedAddress)
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
func (it *GatewayEVMInboundTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedAddressIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedAddress)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayEVMInboundTestLogNamedAddress, error) {
	event := new(GatewayEVMInboundTestLogNamedAddress)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArrayIterator struct {
	Event *GatewayEVMInboundTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedArray)
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
		it.Event = new(GatewayEVMInboundTestLogNamedArray)
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
func (it *GatewayEVMInboundTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedArray represents a LogNamedArray event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedArrayIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedArray)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayEVMInboundTestLogNamedArray, error) {
	event := new(GatewayEVMInboundTestLogNamedArray)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray0Iterator struct {
	Event *GatewayEVMInboundTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedArray0)
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
		it.Event = new(GatewayEVMInboundTestLogNamedArray0)
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
func (it *GatewayEVMInboundTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedArray0Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedArray0)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayEVMInboundTestLogNamedArray0, error) {
	event := new(GatewayEVMInboundTestLogNamedArray0)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray1Iterator struct {
	Event *GatewayEVMInboundTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedArray1)
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
		it.Event = new(GatewayEVMInboundTestLogNamedArray1)
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
func (it *GatewayEVMInboundTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedArray1Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedArray1)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayEVMInboundTestLogNamedArray1, error) {
	event := new(GatewayEVMInboundTestLogNamedArray1)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedBytesIterator struct {
	Event *GatewayEVMInboundTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedBytes)
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
		it.Event = new(GatewayEVMInboundTestLogNamedBytes)
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
func (it *GatewayEVMInboundTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedBytesIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedBytes)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayEVMInboundTestLogNamedBytes, error) {
	event := new(GatewayEVMInboundTestLogNamedBytes)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedBytes32Iterator struct {
	Event *GatewayEVMInboundTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedBytes32)
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
		it.Event = new(GatewayEVMInboundTestLogNamedBytes32)
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
func (it *GatewayEVMInboundTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedBytes32Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedBytes32)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayEVMInboundTestLogNamedBytes32, error) {
	event := new(GatewayEVMInboundTestLogNamedBytes32)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedDecimalIntIterator struct {
	Event *GatewayEVMInboundTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedDecimalInt)
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
		it.Event = new(GatewayEVMInboundTestLogNamedDecimalInt)
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
func (it *GatewayEVMInboundTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedDecimalIntIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedDecimalInt)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayEVMInboundTestLogNamedDecimalInt, error) {
	event := new(GatewayEVMInboundTestLogNamedDecimalInt)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedDecimalUintIterator struct {
	Event *GatewayEVMInboundTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedDecimalUint)
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
		it.Event = new(GatewayEVMInboundTestLogNamedDecimalUint)
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
func (it *GatewayEVMInboundTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedDecimalUintIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedDecimalUint)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayEVMInboundTestLogNamedDecimalUint, error) {
	event := new(GatewayEVMInboundTestLogNamedDecimalUint)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedIntIterator struct {
	Event *GatewayEVMInboundTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedInt)
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
		it.Event = new(GatewayEVMInboundTestLogNamedInt)
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
func (it *GatewayEVMInboundTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedInt represents a LogNamedInt event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedIntIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedInt)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayEVMInboundTestLogNamedInt, error) {
	event := new(GatewayEVMInboundTestLogNamedInt)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedStringIterator struct {
	Event *GatewayEVMInboundTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedString)
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
		it.Event = new(GatewayEVMInboundTestLogNamedString)
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
func (it *GatewayEVMInboundTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedString represents a LogNamedString event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedStringIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedString)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedString(log types.Log) (*GatewayEVMInboundTestLogNamedString, error) {
	event := new(GatewayEVMInboundTestLogNamedString)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedUintIterator struct {
	Event *GatewayEVMInboundTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedUint)
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
		it.Event = new(GatewayEVMInboundTestLogNamedUint)
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
func (it *GatewayEVMInboundTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedUint represents a LogNamedUint event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedUintIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedUint)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayEVMInboundTestLogNamedUint, error) {
	event := new(GatewayEVMInboundTestLogNamedUint)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogStringIterator struct {
	Event *GatewayEVMInboundTestLogString // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogString)
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
		it.Event = new(GatewayEVMInboundTestLogString)
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
func (it *GatewayEVMInboundTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogString represents a LogString event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogStringIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogStringIterator{contract: _GatewayEVMInboundTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogString)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogString(log types.Log) (*GatewayEVMInboundTestLogString, error) {
	event := new(GatewayEVMInboundTestLogString)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogUintIterator struct {
	Event *GatewayEVMInboundTestLogUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogUint)
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
		it.Event = new(GatewayEVMInboundTestLogUint)
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
func (it *GatewayEVMInboundTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogUint represents a LogUint event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogUintIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogUintIterator{contract: _GatewayEVMInboundTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogUint)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogUint(log types.Log) (*GatewayEVMInboundTestLogUint, error) {
	event := new(GatewayEVMInboundTestLogUint)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogsIterator struct {
	Event *GatewayEVMInboundTestLogs // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogs)
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
		it.Event = new(GatewayEVMInboundTestLogs)
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
func (it *GatewayEVMInboundTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogs represents a Logs event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogsIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogsIterator{contract: _GatewayEVMInboundTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogs)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "logs", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogs(log types.Log) (*GatewayEVMInboundTestLogs, error) {
	event := new(GatewayEVMInboundTestLogs)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
