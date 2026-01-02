// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornativeupgradetest

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

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender common.Address
}

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// ZetaConnectorNativeUpgradeTestMetaData contains all meta data concerning the ZetaConnectorNativeUpgradeTest contract.
var ZetaConnectorNativeUpgradeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"gateway_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"messageContext\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndRevert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedZetaConnectorTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516122e46100fd600039600081816112250152818161124e015261142401526122e46000f3fe6080604052600436106101965760003560e01c80638456cb59116100e1578063ad3cb1cc1161008a578063e63ab1e911610064578063e63ab1e914610528578063f3fef3a31461055c578063f61ff82a1461057c578063f8c8765e1461059c57600080fd5b8063ad3cb1cc14610492578063b6b55f25146104e8578063d547741f1461050857600080fd5b8063950837aa116100bb578063950837aa14610429578063a217fddf14610449578063a783c7891461045e57600080fd5b80638456cb591461037b57806385f438c11461039057806391d14854146103c457600080fd5b806336568abe1161014357806352d1902d1161011d57806352d1902d1461030f5780635b112591146103245780635c975abb1461034457600080fd5b806336568abe146102c75780633f4ba83a146102e75780634f1ef286146102fc57600080fd5b806321e093b11161017457806321e093b11461022a578063248a9ca31461024a5780632f2ff15d146102a757600080fd5b806301ffc9a71461019b578063116191b6146101d057806316b12bb614610208575b600080fd5b3480156101a757600080fd5b506101bb6101b6366004611c23565b6105bc565b60405190151581526020015b60405180910390f35b3480156101dc57600080fd5b506000546101f0906001600160a01b031681565b6040516001600160a01b0390911681526020016101c7565b34801561021457600080fd5b50610228610223366004611cca565b610655565b005b34801561023657600080fd5b506001546101f0906001600160a01b031681565b34801561025657600080fd5b50610299610265366004611d5a565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016101c7565b3480156102b357600080fd5b506102286102c2366004611d73565b6107ad565b3480156102d357600080fd5b506102286102e2366004611d73565b6107f7565b3480156102f357600080fd5b50610228610848565b61022861030a366004611dce565b61087d565b34801561031b57600080fd5b5061029961089c565b34801561033057600080fd5b506002546101f0906001600160a01b031681565b34801561035057600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166101bb565b34801561038757600080fd5b506102286108cb565b34801561039c57600080fd5b506102997f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b3480156103d057600080fd5b506101bb6103df366004611d73565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561043557600080fd5b50610228610444366004611ed5565b6108fd565b34801561045557600080fd5b50610299600081565b34801561046a57600080fd5b506102997f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b34801561049e57600080fd5b506104db6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101c79190611f14565b3480156104f457600080fd5b50610228610503366004611d5a565b610a90565b34801561051457600080fd5b50610228610523366004611d73565b610ab0565b34801561053457600080fd5b506102997f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b34801561056857600080fd5b50610228610577366004611f65565b610af4565b34801561058857600080fd5b50610228610597366004611f8f565b610bb2565b3480156105a857600080fd5b506102286105b7366004612006565b610cd0565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061064f57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b61065d610e56565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461068781610ed7565b61068f610ee1565b6000546001546106ac916001600160a01b03918216911687610f3f565b6000546001546040517faa0c0fc10000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263aa0c0fc192610703929116908a908a908a908a908a90600401612154565b600060405180830381600087803b15801561071d57600080fd5b505af1158015610731573d6000803e3d6000fd5b50505050856001600160a01b03167f5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff08686868660405161077494939291906121ab565b60405180910390a2506107a660017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546107e781610ed7565b6107f18383610fd9565b50505050565b6001600160a01b0381163314610839576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61084382826110c6565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61087281610ed7565b61087a61118a565b50565b61088561121a565b61088e826112ea565b61089882826112f5565b5050565b60006108a6611419565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6108f581610ed7565b61087a61147b565b600061090881610ed7565b6001600160a01b038216610948576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025461097f907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4906001600160a01b03166110c6565b506002546109b7907f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb906001600160a01b03166110c6565b506109e27f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610fd9565b50610a0d7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610fd9565b50600254604080516001600160a01b03928316815291841660208301527f33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e910160405180910390a150600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b610a98610ee1565b60015461087a906001600160a01b03163330846114f4565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610aea81610ed7565b6107f183836110c6565b610afc610e56565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610b2681610ed7565b610b2e610ee1565b600154610b45906001600160a01b03168484610f3f565b826001600160a01b03167f3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee983604051610b8091815260200190565b60405180910390a25061089860017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b610bba610e56565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610be481610ed7565b610bec610ee1565b600054600154610c09916001600160a01b03918216911686610f3f565b6000546001546040517f7bbe9afa0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692637bbe9afa92610c61928b92909116908a908a908a908a906004016121e2565b600060405180830381600087803b158015610c7b57600080fd5b505af1158015610c8f573d6000803e3d6000fd5b50505050846001600160a01b03167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d8585856040516107749392919061223d565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610d1b5750825b905060008267ffffffffffffffff166001148015610d385750303b155b905081158015610d46575080155b15610d7d576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610dde5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610dea8989898961152d565b8315610e4b5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01610ed1576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b61087a81336116a9565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615610f3d576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6040516001600160a01b0383811660248301526044820183905261084391859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611736565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166110bc576000848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556110723390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061064f565b600091505061064f565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16156110bc576000848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061064f565b6111926117b2565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806112b357507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166112a77f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15610f3d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061089881610ed7565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561136d575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261136a91810190612257565b60015b6113b3576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc811461140f576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016113aa565b610843838361180d565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610f3d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611483610ee1565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336111fc565b6040516001600160a01b0384811660248301528381166044830152606482018390526107f19186918216906323b872dd90608401610f6c565b611535611863565b6001600160a01b038416158061155257506001600160a01b038316155b8061156457506001600160a01b038216155b8061157657506001600160a01b038116155b156115ad576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115b56118ca565b6115bd6118d2565b6115c56118ca565b6115cd6118e2565b600080546001600160a01b038087167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161783556001805487831690841617905560028054918616919092161790556116289082610fd9565b506116537f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610fd9565b5061167e7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610fd9565b506107a67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610fd9565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610898576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044016113aa565b600061174b6001600160a01b038416836118f2565b9050805160001415801561177057508080602001905181019061176e9190612270565b155b15610843576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201526024016113aa565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610f3d576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61181682611907565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a280511561185b5761084382826119af565b610898611a25565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610f3d576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f3d611863565b6118da611863565b610f3d611a5d565b6118ea611863565b610f3d611a65565b606061190083836000611ab6565b9392505050565b806001600160a01b03163b600003611956576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016113aa565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516119cc9190612292565b600060405180830381855af49150503d8060008114611a07576040519150601f19603f3d011682016040523d82523d6000602084013e611a0c565b606091505b5091509150611a1c858383611b6c565b95945050505050565b3415610f3d576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fb3611863565b611a6d611863565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b606081471015611af4576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016113aa565b600080856001600160a01b03168486604051611b109190612292565b60006040518083038185875af1925050503d8060008114611b4d576040519150601f19603f3d011682016040523d82523d6000602084013e611b52565b606091505b5091509150611b62868383611b6c565b9695505050505050565b606082611b8157611b7c82611be1565b611900565b8151158015611b9857506001600160a01b0384163b155b15611bda576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016113aa565b5080611900565b805115611bf15780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215611c3557600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461190057600080fd5b80356001600160a01b0381168114611c7c57600080fd5b919050565b60008083601f840112611c9357600080fd5b50813567ffffffffffffffff811115611cab57600080fd5b602083019150836020828501011115611cc357600080fd5b9250929050565b600080600080600060808688031215611ce257600080fd5b611ceb86611c65565b945060208601359350604086013567ffffffffffffffff811115611d0e57600080fd5b611d1a88828901611c81565b909450925050606086013567ffffffffffffffff811115611d3a57600080fd5b860160808189031215611d4c57600080fd5b809150509295509295909350565b600060208284031215611d6c57600080fd5b5035919050565b60008060408385031215611d8657600080fd5b82359150611d9660208401611c65565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060408385031215611de157600080fd5b611dea83611c65565b9150602083013567ffffffffffffffff811115611e0657600080fd5b8301601f81018513611e1757600080fd5b803567ffffffffffffffff811115611e3157611e31611d9f565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715611e9d57611e9d611d9f565b604052818152828201602001871015611eb557600080fd5b816020840160208301376000602083830101528093505050509250929050565b600060208284031215611ee757600080fd5b61190082611c65565b60005b83811015611f0b578181015183820152602001611ef3565b50506000910152565b6020815260008251806020840152611f33816040850160208701611ef0565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60008060408385031215611f7857600080fd5b611f8183611c65565b946020939093013593505050565b60008060008060008587036080811215611fa857600080fd5b6020811215611fb657600080fd5b50859450611fc660208701611c65565b935060408601359250606086013567ffffffffffffffff811115611fe957600080fd5b611ff588828901611c81565b969995985093965092949392505050565b6000806000806080858703121561201c57600080fd5b61202585611c65565b935061203360208601611c65565b925061204160408601611c65565b915061204f60608601611c65565b905092959194509250565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b036120b482611c65565b1682526001600160a01b036120cb60208301611c65565b1660208301526040818101359083015260006060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261211357600080fd5b820160208101903567ffffffffffffffff81111561213057600080fd5b80360382131561213f57600080fd5b60806060860152611a1c60808601828461205a565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a06060820152600061218c60a08301858761205a565b828103608084015261219e81856120a3565b9998505050505050505050565b8481526060602082015260006121c560608301858761205a565b82810360408401526121d781856120a3565b979650505050505050565b6001600160a01b036121f388611c65565b1681526001600160a01b03861660208201526001600160a01b038516604082015283606082015260a06080820152600061223160a08301848661205a565b98975050505050505050565b838152604060208201526000611a1c60408301848661205a565b60006020828403121561226957600080fd5b5051919050565b60006020828403121561228257600080fd5b8151801515811461190057600080fd5b600082516122a4818460208701611ef0565b919091019291505056fea2646970667358221220c6bda2270a5113fed7206de068eb2393837852902737eb44ccb88c6010ea463064736f6c634300081a0033",
}

// ZetaConnectorNativeUpgradeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNativeUpgradeTestMetaData.ABI instead.
var ZetaConnectorNativeUpgradeTestABI = ZetaConnectorNativeUpgradeTestMetaData.ABI

// ZetaConnectorNativeUpgradeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNativeUpgradeTestMetaData.Bin instead.
var ZetaConnectorNativeUpgradeTestBin = ZetaConnectorNativeUpgradeTestMetaData.Bin

// DeployZetaConnectorNativeUpgradeTest deploys a new Ethereum contract, binding an instance of ZetaConnectorNativeUpgradeTest to it.
func DeployZetaConnectorNativeUpgradeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZetaConnectorNativeUpgradeTest, error) {
	parsed, err := ZetaConnectorNativeUpgradeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNativeUpgradeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNativeUpgradeTest{ZetaConnectorNativeUpgradeTestCaller: ZetaConnectorNativeUpgradeTestCaller{contract: contract}, ZetaConnectorNativeUpgradeTestTransactor: ZetaConnectorNativeUpgradeTestTransactor{contract: contract}, ZetaConnectorNativeUpgradeTestFilterer: ZetaConnectorNativeUpgradeTestFilterer{contract: contract}}, nil
}

// ZetaConnectorNativeUpgradeTest is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNativeUpgradeTest struct {
	ZetaConnectorNativeUpgradeTestCaller     // Read-only binding to the contract
	ZetaConnectorNativeUpgradeTestTransactor // Write-only binding to the contract
	ZetaConnectorNativeUpgradeTestFilterer   // Log filterer for contract events
}

// ZetaConnectorNativeUpgradeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNativeUpgradeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeUpgradeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNativeUpgradeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeUpgradeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNativeUpgradeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeUpgradeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNativeUpgradeTestSession struct {
	Contract     *ZetaConnectorNativeUpgradeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ZetaConnectorNativeUpgradeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNativeUpgradeTestCallerSession struct {
	Contract *ZetaConnectorNativeUpgradeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// ZetaConnectorNativeUpgradeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNativeUpgradeTestTransactorSession struct {
	Contract     *ZetaConnectorNativeUpgradeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ZetaConnectorNativeUpgradeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNativeUpgradeTestRaw struct {
	Contract *ZetaConnectorNativeUpgradeTest // Generic contract binding to access the raw methods on
}

// ZetaConnectorNativeUpgradeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNativeUpgradeTestCallerRaw struct {
	Contract *ZetaConnectorNativeUpgradeTestCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNativeUpgradeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNativeUpgradeTestTransactorRaw struct {
	Contract *ZetaConnectorNativeUpgradeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNativeUpgradeTest creates a new instance of ZetaConnectorNativeUpgradeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeUpgradeTest(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNativeUpgradeTest, error) {
	contract, err := bindZetaConnectorNativeUpgradeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTest{ZetaConnectorNativeUpgradeTestCaller: ZetaConnectorNativeUpgradeTestCaller{contract: contract}, ZetaConnectorNativeUpgradeTestTransactor: ZetaConnectorNativeUpgradeTestTransactor{contract: contract}, ZetaConnectorNativeUpgradeTestFilterer: ZetaConnectorNativeUpgradeTestFilterer{contract: contract}}, nil
}

// NewZetaConnectorNativeUpgradeTestCaller creates a new read-only instance of ZetaConnectorNativeUpgradeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeUpgradeTestCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNativeUpgradeTestCaller, error) {
	contract, err := bindZetaConnectorNativeUpgradeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestCaller{contract: contract}, nil
}

// NewZetaConnectorNativeUpgradeTestTransactor creates a new write-only instance of ZetaConnectorNativeUpgradeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeUpgradeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNativeUpgradeTestTransactor, error) {
	contract, err := bindZetaConnectorNativeUpgradeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestTransactor{contract: contract}, nil
}

// NewZetaConnectorNativeUpgradeTestFilterer creates a new log filterer instance of ZetaConnectorNativeUpgradeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeUpgradeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNativeUpgradeTestFilterer, error) {
	contract, err := bindZetaConnectorNativeUpgradeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestFilterer{contract: contract}, nil
}

// bindZetaConnectorNativeUpgradeTest binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNativeUpgradeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNativeUpgradeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNativeUpgradeTest.Contract.ZetaConnectorNativeUpgradeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ZetaConnectorNativeUpgradeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ZetaConnectorNativeUpgradeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNativeUpgradeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.DEFAULTADMINROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.DEFAULTADMINROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.PAUSERROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.PAUSERROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.TSSROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.TSSROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.WITHDRAWERROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.WITHDRAWERROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Gateway(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Gateway(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.GetRoleAdmin(&_ZetaConnectorNativeUpgradeTest.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.GetRoleAdmin(&_ZetaConnectorNativeUpgradeTest.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.HasRole(&_ZetaConnectorNativeUpgradeTest.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.HasRole(&_ZetaConnectorNativeUpgradeTest.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Paused() (bool, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Paused(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) Paused() (bool, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Paused(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) ProxiableUUID() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ProxiableUUID(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ProxiableUUID(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.SupportsInterface(&_ZetaConnectorNativeUpgradeTest.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.SupportsInterface(&_ZetaConnectorNativeUpgradeTest.CallOpts, interfaceId)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.TssAddress(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.TssAddress(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ZetaToken(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ZetaToken(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Deposit(&_ZetaConnectorNativeUpgradeTest.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Deposit(&_ZetaConnectorNativeUpgradeTest.TransactOpts, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.GrantRole(&_ZetaConnectorNativeUpgradeTest.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.GrantRole(&_ZetaConnectorNativeUpgradeTest.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address gateway_, address zetaToken_, address tssAddress_, address admin_) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) Initialize(opts *bind.TransactOpts, gateway_ common.Address, zetaToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "initialize", gateway_, zetaToken_, tssAddress_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address gateway_, address zetaToken_, address tssAddress_, address admin_) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Initialize(gateway_ common.Address, zetaToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Initialize(&_ZetaConnectorNativeUpgradeTest.TransactOpts, gateway_, zetaToken_, tssAddress_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address gateway_, address zetaToken_, address tssAddress_, address admin_) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) Initialize(gateway_ common.Address, zetaToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Initialize(&_ZetaConnectorNativeUpgradeTest.TransactOpts, gateway_, zetaToken_, tssAddress_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Pause(&_ZetaConnectorNativeUpgradeTest.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Pause(&_ZetaConnectorNativeUpgradeTest.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.RenounceRole(&_ZetaConnectorNativeUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.RenounceRole(&_ZetaConnectorNativeUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.RevokeRole(&_ZetaConnectorNativeUpgradeTest.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.RevokeRole(&_ZetaConnectorNativeUpgradeTest.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Unpause(&_ZetaConnectorNativeUpgradeTest.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Unpause(&_ZetaConnectorNativeUpgradeTest.TransactOpts)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) UpdateTSSAddress(opts *bind.TransactOpts, newTSSAddress common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "updateTSSAddress", newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.UpdateTSSAddress(&_ZetaConnectorNativeUpgradeTest.TransactOpts, newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.UpdateTSSAddress(&_ZetaConnectorNativeUpgradeTest.TransactOpts, newTSSAddress)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.UpgradeToAndCall(&_ZetaConnectorNativeUpgradeTest.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.UpgradeToAndCall(&_ZetaConnectorNativeUpgradeTest.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Withdraw(&_ZetaConnectorNativeUpgradeTest.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Withdraw(&_ZetaConnectorNativeUpgradeTest.TransactOpts, to, amount)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0xf61ff82a.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, uint256 amount, bytes data) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) WithdrawAndCall(opts *bind.TransactOpts, messageContext MessageContext, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "withdrawAndCall", messageContext, to, amount, data)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0xf61ff82a.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, uint256 amount, bytes data) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) WithdrawAndCall(messageContext MessageContext, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.WithdrawAndCall(&_ZetaConnectorNativeUpgradeTest.TransactOpts, messageContext, to, amount, data)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0xf61ff82a.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, uint256 amount, bytes data) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) WithdrawAndCall(messageContext MessageContext, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.WithdrawAndCall(&_ZetaConnectorNativeUpgradeTest.TransactOpts, messageContext, to, amount, data)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x16b12bb6.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) WithdrawAndRevert(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "withdrawAndRevert", to, amount, data, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x16b12bb6.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.WithdrawAndRevert(&_ZetaConnectorNativeUpgradeTest.TransactOpts, to, amount, data, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x16b12bb6.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.WithdrawAndRevert(&_ZetaConnectorNativeUpgradeTest.TransactOpts, to, amount, data, revertContext)
}

// ZetaConnectorNativeUpgradeTestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestInitializedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestInitialized // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestInitialized)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestInitialized)
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
func (it *ZetaConnectorNativeUpgradeTestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestInitialized represents a Initialized event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZetaConnectorNativeUpgradeTestInitializedIterator, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestInitializedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestInitialized) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestInitialized)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseInitialized(log types.Log) (*ZetaConnectorNativeUpgradeTestInitialized, error) {
	event := new(ZetaConnectorNativeUpgradeTestInitialized)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestPausedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestPaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestPaused)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestPaused)
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
func (it *ZetaConnectorNativeUpgradeTestPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestPaused represents a Paused event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterPaused(opts *bind.FilterOpts) (*ZetaConnectorNativeUpgradeTestPausedIterator, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestPausedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestPaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestPaused)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParsePaused(log types.Log) (*ZetaConnectorNativeUpgradeTestPaused, error) {
	event := new(ZetaConnectorNativeUpgradeTestPaused)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestRoleAdminChanged)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestRoleAdminChanged)
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
func (it *ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestRoleAdminChanged represents a RoleAdminChanged event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestRoleAdminChanged)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseRoleAdminChanged(log types.Log) (*ZetaConnectorNativeUpgradeTestRoleAdminChanged, error) {
	event := new(ZetaConnectorNativeUpgradeTestRoleAdminChanged)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestRoleGrantedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestRoleGranted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestRoleGranted)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestRoleGranted)
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
func (it *ZetaConnectorNativeUpgradeTestRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestRoleGranted represents a RoleGranted event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZetaConnectorNativeUpgradeTestRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestRoleGrantedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestRoleGranted)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseRoleGranted(log types.Log) (*ZetaConnectorNativeUpgradeTestRoleGranted, error) {
	event := new(ZetaConnectorNativeUpgradeTestRoleGranted)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestRoleRevokedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestRoleRevoked)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestRoleRevoked)
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
func (it *ZetaConnectorNativeUpgradeTestRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestRoleRevoked represents a RoleRevoked event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZetaConnectorNativeUpgradeTestRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestRoleRevokedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestRoleRevoked)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseRoleRevoked(log types.Log) (*ZetaConnectorNativeUpgradeTestRoleRevoked, error) {
	event := new(ZetaConnectorNativeUpgradeTestRoleRevoked)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestUnpausedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestUnpaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestUnpaused)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestUnpaused)
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
func (it *ZetaConnectorNativeUpgradeTestUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestUnpaused represents a Unpaused event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ZetaConnectorNativeUpgradeTestUnpausedIterator, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestUnpausedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestUnpaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestUnpaused)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseUnpaused(log types.Log) (*ZetaConnectorNativeUpgradeTestUnpaused, error) {
	event := new(ZetaConnectorNativeUpgradeTestUnpaused)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator is returned from FilterUpdatedZetaConnectorTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedZetaConnectorTSSAddress events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator struct {
	Event *ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress)
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
func (it *ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress represents a UpdatedZetaConnectorTSSAddress event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedZetaConnectorTSSAddress is a free log retrieval operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterUpdatedZetaConnectorTSSAddress(opts *bind.FilterOpts) (*ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "UpdatedZetaConnectorTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedZetaConnectorTSSAddress is a free log subscription operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchUpdatedZetaConnectorTSSAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
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

// ParseUpdatedZetaConnectorTSSAddress is a log parse operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseUpdatedZetaConnectorTSSAddress(log types.Log) (*ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress, error) {
	event := new(ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestUpgradedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestUpgraded // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestUpgraded)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestUpgraded)
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
func (it *ZetaConnectorNativeUpgradeTestUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestUpgraded represents a Upgraded event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ZetaConnectorNativeUpgradeTestUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestUpgradedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestUpgraded)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseUpgraded(log types.Log) (*ZetaConnectorNativeUpgradeTestUpgraded, error) {
	event := new(ZetaConnectorNativeUpgradeTestUpgraded)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnIterator struct {
	Event *ZetaConnectorNativeUpgradeTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawn)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawn)
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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestWithdrawn represents a Withdrawn event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeUpgradeTestWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestWithdrawnIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestWithdrawn)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseWithdrawn(log types.Log) (*ZetaConnectorNativeUpgradeTestWithdrawn, error) {
	event := new(ZetaConnectorNativeUpgradeTestWithdrawn)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator struct {
	Event *ZetaConnectorNativeUpgradeTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawnAndCalled)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawnAndCalled)
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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnAndCalled struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestWithdrawnAndCalled, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestWithdrawnAndCalled)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*ZetaConnectorNativeUpgradeTestWithdrawnAndCalled, error) {
	event := new(ZetaConnectorNativeUpgradeTestWithdrawnAndCalled)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawnAndReverted)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawnAndReverted)
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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnAndReverted struct {
	To            common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestWithdrawnAndReverted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestWithdrawnAndReverted)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*ZetaConnectorNativeUpgradeTestWithdrawnAndReverted, error) {
	event := new(ZetaConnectorNativeUpgradeTestWithdrawnAndReverted)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator struct {
	Event *ZetaConnectorNativeUpgradeTestWithdrawnV2 // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawnV2)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawnV2)
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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnV2 represents a WithdrawnV2 event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnV2 struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnV2 is a free log retrieval operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "WithdrawnV2", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestWithdrawnV2, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "WithdrawnV2", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestWithdrawnV2)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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

// ParseWithdrawnV2 is a log parse operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseWithdrawnV2(log types.Log) (*ZetaConnectorNativeUpgradeTestWithdrawnV2, error) {
	event := new(ZetaConnectorNativeUpgradeTestWithdrawnV2)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
