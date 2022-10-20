// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ocr2dr_oracle

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

var OCR2DROracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"donPublicKey\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptyRequestData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySendersList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LowGasForConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToSetSenders\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"AuthorizedSendersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"OracleResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"fulfillRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizedSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDONPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isAuthorizedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subscriptionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"}],\"name\":\"setAuthorizedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620015a2380380620015a283398101604081905262000034916200023b565b818060006001600160a01b038216620000945760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600380546001600160a01b0319166001600160a01b0384811691909117909155811615620000c757620000c781620000e8565b50508151620000df9150600590602084019062000195565b5050506200038e565b6001600160a01b038116331415620001435760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200008b565b600480546001600160a01b0319166001600160a01b03838116918217909255600354604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b828054620001a3906200033b565b90600052602060002090601f016020900481019282620001c7576000855562000212565b82601f10620001e257805160ff191683800117855562000212565b8280016001018555821562000212579182015b8281111562000212578251825591602001919060010190620001f5565b506200022092915062000224565b5090565b5b8082111562000220576000815560010162000225565b600080604083850312156200024f57600080fd5b82516001600160a01b03811681146200026757600080fd5b602084810151919350906001600160401b03808211156200028757600080fd5b818601915086601f8301126200029c57600080fd5b815181811115620002b157620002b162000378565b604051601f8201601f19908116603f01168101908382118183101715620002dc57620002dc62000378565b816040528281528986848701011115620002f557600080fd5b600093505b82841015620003195784840186015181850187015292850192620002fa565b828411156200032b5760008684830101525b8096505050505050509250929050565b600181811c908216806200035057607f821691505b602082108114156200037257634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b611204806200039e6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c8063bb9fa3f511610076578063ee56997b1161005b578063ee56997b14610191578063f2fde38b146101a4578063fa00763a146101b757600080fd5b8063bb9fa3f514610168578063d328a91e1461018957600080fd5b806339b05122116100a757806339b051221461012357806379ba5097146101385780638da5cb5b1461014057600080fd5b8063181f5a77146100c35780632408afaa1461010e575b600080fd5b60408051808201909152601281527f4f43523244524f7261636c6520302e302e30000000000000000000000000000060208201525b60405161010591906110b3565b60405180910390f35b6101166101da565b6040516101059190610ffd565b610136610131366004610e0b565b610249565b005b610136610416565b60035460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610105565b61017b610176366004610e85565b61051c565b604051908152602001610105565b6100f861068b565b61013661019f366004610d96565b610714565b6101366101b2366004610d7b565b610883565b6101ca6101c5366004610d7b565b610897565b6040519015158152602001610105565b6060600280548060200260200160405190810160405280929190818152602001828054801561023f57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610214575b5050505050905090565b600085815260076020526040902054859073ffffffffffffffffffffffffffffffffffffffff166102a6576040517f803ed86300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102ae6108a9565b60008681526007602090815260409182902054915188815273ffffffffffffffffffffffffffffffffffffffff909216917f9e9bc7616d42c2835d05ae617e508454e63b30b934be8aa932ebc125e0e58a64910160405180910390a162061a805a1015610347576040517f566e3edd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f0ca7617500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821690630ca76175906103a1908a908a908a908a908a9060040161107a565b600060405180830381600087803b1580156103bb57600080fd5b505af11580156103cf573d6000803e3d6000fd5b50505060009788525050600760205250506040842080547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560010193909355505050565b60045473ffffffffffffffffffffffffffffffffffffffff16331461049c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560048054909116905560405173ffffffffffffffffffffffffffffffffffffffff909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b600081610554576040517ec1cfc000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805490600061056483611131565b90915550506006546040517fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003360601b1660208201526034810191909152600090605401604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152828252805160209182012083830183523384528184018981526000828152600790935291839020935184547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161784559051600190930192909255519091507f9dc59d1e6d6042f6c2d2af5a3f9f6502bf0abf7cf40f0eb80edf752ca396f2ba9061067b90839087908790611057565b60405180910390a1949350505050565b60606005805461069a906110dd565b80601f01602080910402602001604051908101604052809291908181526020018280546106c6906110dd565b801561023f5780601f106106e85761010080835404028352916020019161023f565b820191906000526020600020905b8154815290600101906020018083116106f657509395945050505050565b61071c6108ea565b610752576040517fad77f06100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610789576040517f75158c3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b6002548110156107e5576107d2600282815481106107ac576107ac6111c8565b600091825260208220015473ffffffffffffffffffffffffffffffffffffffff16610938565b50806107dd81611131565b91505061078c565b5060005b8181101561083657610823838383818110610806576108066111c8565b905060200201602081019061081b9190610d7b565b600090610961565b508061082e81611131565b9150506107e9565b5061084360028383610c6c565b507ff263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a082823360405161087793929190610f85565b60405180910390a15050565b61088b610983565b61089481610a04565b50565b60006108a38183610afb565b92915050565b6108b233610897565b6108e8576040517f0809490800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60006108f533610897565b8061093357503361091b60035473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16145b905090565b600061095a8373ffffffffffffffffffffffffffffffffffffffff8416610b2a565b9392505050565b600061095a8373ffffffffffffffffffffffffffffffffffffffff8416610c1d565b60035473ffffffffffffffffffffffffffffffffffffffff1633146108e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610493565b73ffffffffffffffffffffffffffffffffffffffff8116331415610a84576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610493565b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600354604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020526040812054151561095a565b60008181526001830160205260408120548015610c13576000610b4e6001836110c6565b8554909150600090610b62906001906110c6565b9050818114610bc7576000866000018281548110610b8257610b826111c8565b9060005260206000200154905080876000018481548110610ba557610ba56111c8565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080610bd857610bd8611199565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506108a3565b60009150506108a3565b6000818152600183016020526040812054610c64575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556108a3565b5060006108a3565b828054828255906000526020600020908101928215610ce4579160200282015b82811115610ce45781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190610c8c565b50610cf0929150610cf4565b5090565b5b80821115610cf05760008155600101610cf5565b803573ffffffffffffffffffffffffffffffffffffffff81168114610d2d57600080fd5b919050565b60008083601f840112610d4457600080fd5b50813567ffffffffffffffff811115610d5c57600080fd5b602083019150836020828501011115610d7457600080fd5b9250929050565b600060208284031215610d8d57600080fd5b61095a82610d09565b60008060208385031215610da957600080fd5b823567ffffffffffffffff80821115610dc157600080fd5b818501915085601f830112610dd557600080fd5b813581811115610de457600080fd5b8660208260051b8501011115610df957600080fd5b60209290920196919550909350505050565b600080600080600060608688031215610e2357600080fd5b85359450602086013567ffffffffffffffff80821115610e4257600080fd5b610e4e89838a01610d32565b90965094506040880135915080821115610e6757600080fd5b50610e7488828901610d32565b969995985093965092949392505050565b600080600060408486031215610e9a57600080fd5b83359250602084013567ffffffffffffffff811115610eb857600080fd5b610ec486828701610d32565b9497909650939450505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6000815180845260005b81811015610f4057602081850181015186830182015201610f24565b81811115610f52576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6040808252810183905260008460608301825b86811015610fd35773ffffffffffffffffffffffffffffffffffffffff610fbe84610d09565b16825260209283019290910190600101610f98565b50809250505073ffffffffffffffffffffffffffffffffffffffff83166020830152949350505050565b6020808252825182820181905260009190848201906040850190845b8181101561104b57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101611019565b50909695505050505050565b838152604060208201526000611071604083018486610ed1565b95945050505050565b858152606060208201526000611094606083018688610ed1565b82810360408401526110a7818587610ed1565b98975050505050505050565b60208152600061095a6020830184610f1a565b6000828210156110d8576110d861116a565b500390565b600181811c908216806110f157607f821691505b6020821081141561112b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156111635761116361116a565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c6343000806000a",
}

var OCR2DROracleABI = OCR2DROracleMetaData.ABI

var OCR2DROracleBin = OCR2DROracleMetaData.Bin

func DeployOCR2DROracle(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address, donPublicKey []byte) (common.Address, *types.Transaction, *OCR2DROracle, error) {
	parsed, err := OCR2DROracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OCR2DROracleBin), backend, owner, donPublicKey)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OCR2DROracle{OCR2DROracleCaller: OCR2DROracleCaller{contract: contract}, OCR2DROracleTransactor: OCR2DROracleTransactor{contract: contract}, OCR2DROracleFilterer: OCR2DROracleFilterer{contract: contract}}, nil
}

type OCR2DROracle struct {
	address common.Address
	abi     abi.ABI
	OCR2DROracleCaller
	OCR2DROracleTransactor
	OCR2DROracleFilterer
}

type OCR2DROracleCaller struct {
	contract *bind.BoundContract
}

type OCR2DROracleTransactor struct {
	contract *bind.BoundContract
}

type OCR2DROracleFilterer struct {
	contract *bind.BoundContract
}

type OCR2DROracleSession struct {
	Contract     *OCR2DROracle
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OCR2DROracleCallerSession struct {
	Contract *OCR2DROracleCaller
	CallOpts bind.CallOpts
}

type OCR2DROracleTransactorSession struct {
	Contract     *OCR2DROracleTransactor
	TransactOpts bind.TransactOpts
}

type OCR2DROracleRaw struct {
	Contract *OCR2DROracle
}

type OCR2DROracleCallerRaw struct {
	Contract *OCR2DROracleCaller
}

type OCR2DROracleTransactorRaw struct {
	Contract *OCR2DROracleTransactor
}

func NewOCR2DROracle(address common.Address, backend bind.ContractBackend) (*OCR2DROracle, error) {
	abi, err := abi.JSON(strings.NewReader(OCR2DROracleABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOCR2DROracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OCR2DROracle{address: address, abi: abi, OCR2DROracleCaller: OCR2DROracleCaller{contract: contract}, OCR2DROracleTransactor: OCR2DROracleTransactor{contract: contract}, OCR2DROracleFilterer: OCR2DROracleFilterer{contract: contract}}, nil
}

func NewOCR2DROracleCaller(address common.Address, caller bind.ContractCaller) (*OCR2DROracleCaller, error) {
	contract, err := bindOCR2DROracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OCR2DROracleCaller{contract: contract}, nil
}

func NewOCR2DROracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OCR2DROracleTransactor, error) {
	contract, err := bindOCR2DROracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OCR2DROracleTransactor{contract: contract}, nil
}

func NewOCR2DROracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OCR2DROracleFilterer, error) {
	contract, err := bindOCR2DROracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OCR2DROracleFilterer{contract: contract}, nil
}

func bindOCR2DROracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OCR2DROracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_OCR2DROracle *OCR2DROracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OCR2DROracle.Contract.OCR2DROracleCaller.contract.Call(opts, result, method, params...)
}

func (_OCR2DROracle *OCR2DROracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OCR2DROracle.Contract.OCR2DROracleTransactor.contract.Transfer(opts)
}

func (_OCR2DROracle *OCR2DROracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OCR2DROracle.Contract.OCR2DROracleTransactor.contract.Transact(opts, method, params...)
}

func (_OCR2DROracle *OCR2DROracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OCR2DROracle.Contract.contract.Call(opts, result, method, params...)
}

func (_OCR2DROracle *OCR2DROracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OCR2DROracle.Contract.contract.Transfer(opts)
}

func (_OCR2DROracle *OCR2DROracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OCR2DROracle.Contract.contract.Transact(opts, method, params...)
}

func (_OCR2DROracle *OCR2DROracleCaller) GetAuthorizedSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OCR2DROracle.contract.Call(opts, &out, "getAuthorizedSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OCR2DROracle *OCR2DROracleSession) GetAuthorizedSenders() ([]common.Address, error) {
	return _OCR2DROracle.Contract.GetAuthorizedSenders(&_OCR2DROracle.CallOpts)
}

func (_OCR2DROracle *OCR2DROracleCallerSession) GetAuthorizedSenders() ([]common.Address, error) {
	return _OCR2DROracle.Contract.GetAuthorizedSenders(&_OCR2DROracle.CallOpts)
}

func (_OCR2DROracle *OCR2DROracleCaller) GetDONPublicKey(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _OCR2DROracle.contract.Call(opts, &out, "getDONPublicKey")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_OCR2DROracle *OCR2DROracleSession) GetDONPublicKey() ([]byte, error) {
	return _OCR2DROracle.Contract.GetDONPublicKey(&_OCR2DROracle.CallOpts)
}

func (_OCR2DROracle *OCR2DROracleCallerSession) GetDONPublicKey() ([]byte, error) {
	return _OCR2DROracle.Contract.GetDONPublicKey(&_OCR2DROracle.CallOpts)
}

func (_OCR2DROracle *OCR2DROracleCaller) IsAuthorizedSender(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _OCR2DROracle.contract.Call(opts, &out, "isAuthorizedSender", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OCR2DROracle *OCR2DROracleSession) IsAuthorizedSender(sender common.Address) (bool, error) {
	return _OCR2DROracle.Contract.IsAuthorizedSender(&_OCR2DROracle.CallOpts, sender)
}

func (_OCR2DROracle *OCR2DROracleCallerSession) IsAuthorizedSender(sender common.Address) (bool, error) {
	return _OCR2DROracle.Contract.IsAuthorizedSender(&_OCR2DROracle.CallOpts, sender)
}

func (_OCR2DROracle *OCR2DROracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OCR2DROracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OCR2DROracle *OCR2DROracleSession) Owner() (common.Address, error) {
	return _OCR2DROracle.Contract.Owner(&_OCR2DROracle.CallOpts)
}

func (_OCR2DROracle *OCR2DROracleCallerSession) Owner() (common.Address, error) {
	return _OCR2DROracle.Contract.Owner(&_OCR2DROracle.CallOpts)
}

func (_OCR2DROracle *OCR2DROracleCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OCR2DROracle.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OCR2DROracle *OCR2DROracleSession) TypeAndVersion() (string, error) {
	return _OCR2DROracle.Contract.TypeAndVersion(&_OCR2DROracle.CallOpts)
}

func (_OCR2DROracle *OCR2DROracleCallerSession) TypeAndVersion() (string, error) {
	return _OCR2DROracle.Contract.TypeAndVersion(&_OCR2DROracle.CallOpts)
}

func (_OCR2DROracle *OCR2DROracleTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OCR2DROracle.contract.Transact(opts, "acceptOwnership")
}

func (_OCR2DROracle *OCR2DROracleSession) AcceptOwnership() (*types.Transaction, error) {
	return _OCR2DROracle.Contract.AcceptOwnership(&_OCR2DROracle.TransactOpts)
}

func (_OCR2DROracle *OCR2DROracleTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OCR2DROracle.Contract.AcceptOwnership(&_OCR2DROracle.TransactOpts)
}

func (_OCR2DROracle *OCR2DROracleTransactor) FulfillRequest(opts *bind.TransactOpts, requestId [32]byte, response []byte, err []byte) (*types.Transaction, error) {
	return _OCR2DROracle.contract.Transact(opts, "fulfillRequest", requestId, response, err)
}

func (_OCR2DROracle *OCR2DROracleSession) FulfillRequest(requestId [32]byte, response []byte, err []byte) (*types.Transaction, error) {
	return _OCR2DROracle.Contract.FulfillRequest(&_OCR2DROracle.TransactOpts, requestId, response, err)
}

func (_OCR2DROracle *OCR2DROracleTransactorSession) FulfillRequest(requestId [32]byte, response []byte, err []byte) (*types.Transaction, error) {
	return _OCR2DROracle.Contract.FulfillRequest(&_OCR2DROracle.TransactOpts, requestId, response, err)
}

func (_OCR2DROracle *OCR2DROracleTransactor) SendRequest(opts *bind.TransactOpts, subscriptionId *big.Int, data []byte) (*types.Transaction, error) {
	return _OCR2DROracle.contract.Transact(opts, "sendRequest", subscriptionId, data)
}

func (_OCR2DROracle *OCR2DROracleSession) SendRequest(subscriptionId *big.Int, data []byte) (*types.Transaction, error) {
	return _OCR2DROracle.Contract.SendRequest(&_OCR2DROracle.TransactOpts, subscriptionId, data)
}

func (_OCR2DROracle *OCR2DROracleTransactorSession) SendRequest(subscriptionId *big.Int, data []byte) (*types.Transaction, error) {
	return _OCR2DROracle.Contract.SendRequest(&_OCR2DROracle.TransactOpts, subscriptionId, data)
}

func (_OCR2DROracle *OCR2DROracleTransactor) SetAuthorizedSenders(opts *bind.TransactOpts, senders []common.Address) (*types.Transaction, error) {
	return _OCR2DROracle.contract.Transact(opts, "setAuthorizedSenders", senders)
}

func (_OCR2DROracle *OCR2DROracleSession) SetAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _OCR2DROracle.Contract.SetAuthorizedSenders(&_OCR2DROracle.TransactOpts, senders)
}

func (_OCR2DROracle *OCR2DROracleTransactorSession) SetAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _OCR2DROracle.Contract.SetAuthorizedSenders(&_OCR2DROracle.TransactOpts, senders)
}

func (_OCR2DROracle *OCR2DROracleTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OCR2DROracle.contract.Transact(opts, "transferOwnership", to)
}

func (_OCR2DROracle *OCR2DROracleSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OCR2DROracle.Contract.TransferOwnership(&_OCR2DROracle.TransactOpts, to)
}

func (_OCR2DROracle *OCR2DROracleTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OCR2DROracle.Contract.TransferOwnership(&_OCR2DROracle.TransactOpts, to)
}

type OCR2DROracleAuthorizedSendersChangedIterator struct {
	Event *OCR2DROracleAuthorizedSendersChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2DROracleAuthorizedSendersChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2DROracleAuthorizedSendersChanged)
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

	select {
	case log := <-it.logs:
		it.Event = new(OCR2DROracleAuthorizedSendersChanged)
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

func (it *OCR2DROracleAuthorizedSendersChangedIterator) Error() error {
	return it.fail
}

func (it *OCR2DROracleAuthorizedSendersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2DROracleAuthorizedSendersChanged struct {
	Senders   []common.Address
	ChangedBy common.Address
	Raw       types.Log
}

func (_OCR2DROracle *OCR2DROracleFilterer) FilterAuthorizedSendersChanged(opts *bind.FilterOpts) (*OCR2DROracleAuthorizedSendersChangedIterator, error) {

	logs, sub, err := _OCR2DROracle.contract.FilterLogs(opts, "AuthorizedSendersChanged")
	if err != nil {
		return nil, err
	}
	return &OCR2DROracleAuthorizedSendersChangedIterator{contract: _OCR2DROracle.contract, event: "AuthorizedSendersChanged", logs: logs, sub: sub}, nil
}

func (_OCR2DROracle *OCR2DROracleFilterer) WatchAuthorizedSendersChanged(opts *bind.WatchOpts, sink chan<- *OCR2DROracleAuthorizedSendersChanged) (event.Subscription, error) {

	logs, sub, err := _OCR2DROracle.contract.WatchLogs(opts, "AuthorizedSendersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2DROracleAuthorizedSendersChanged)
				if err := _OCR2DROracle.contract.UnpackLog(event, "AuthorizedSendersChanged", log); err != nil {
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

func (_OCR2DROracle *OCR2DROracleFilterer) ParseAuthorizedSendersChanged(log types.Log) (*OCR2DROracleAuthorizedSendersChanged, error) {
	event := new(OCR2DROracleAuthorizedSendersChanged)
	if err := _OCR2DROracle.contract.UnpackLog(event, "AuthorizedSendersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2DROracleOracleRequestIterator struct {
	Event *OCR2DROracleOracleRequest

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2DROracleOracleRequestIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2DROracleOracleRequest)
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

	select {
	case log := <-it.logs:
		it.Event = new(OCR2DROracleOracleRequest)
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

func (it *OCR2DROracleOracleRequestIterator) Error() error {
	return it.fail
}

func (it *OCR2DROracleOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2DROracleOracleRequest struct {
	RequestId [32]byte
	Data      []byte
	Raw       types.Log
}

func (_OCR2DROracle *OCR2DROracleFilterer) FilterOracleRequest(opts *bind.FilterOpts) (*OCR2DROracleOracleRequestIterator, error) {

	logs, sub, err := _OCR2DROracle.contract.FilterLogs(opts, "OracleRequest")
	if err != nil {
		return nil, err
	}
	return &OCR2DROracleOracleRequestIterator{contract: _OCR2DROracle.contract, event: "OracleRequest", logs: logs, sub: sub}, nil
}

func (_OCR2DROracle *OCR2DROracleFilterer) WatchOracleRequest(opts *bind.WatchOpts, sink chan<- *OCR2DROracleOracleRequest) (event.Subscription, error) {

	logs, sub, err := _OCR2DROracle.contract.WatchLogs(opts, "OracleRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2DROracleOracleRequest)
				if err := _OCR2DROracle.contract.UnpackLog(event, "OracleRequest", log); err != nil {
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

func (_OCR2DROracle *OCR2DROracleFilterer) ParseOracleRequest(log types.Log) (*OCR2DROracleOracleRequest, error) {
	event := new(OCR2DROracleOracleRequest)
	if err := _OCR2DROracle.contract.UnpackLog(event, "OracleRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2DROracleOracleResponseIterator struct {
	Event *OCR2DROracleOracleResponse

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2DROracleOracleResponseIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2DROracleOracleResponse)
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

	select {
	case log := <-it.logs:
		it.Event = new(OCR2DROracleOracleResponse)
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

func (it *OCR2DROracleOracleResponseIterator) Error() error {
	return it.fail
}

func (it *OCR2DROracleOracleResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2DROracleOracleResponse struct {
	RequestId [32]byte
	Raw       types.Log
}

func (_OCR2DROracle *OCR2DROracleFilterer) FilterOracleResponse(opts *bind.FilterOpts) (*OCR2DROracleOracleResponseIterator, error) {

	logs, sub, err := _OCR2DROracle.contract.FilterLogs(opts, "OracleResponse")
	if err != nil {
		return nil, err
	}
	return &OCR2DROracleOracleResponseIterator{contract: _OCR2DROracle.contract, event: "OracleResponse", logs: logs, sub: sub}, nil
}

func (_OCR2DROracle *OCR2DROracleFilterer) WatchOracleResponse(opts *bind.WatchOpts, sink chan<- *OCR2DROracleOracleResponse) (event.Subscription, error) {

	logs, sub, err := _OCR2DROracle.contract.WatchLogs(opts, "OracleResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2DROracleOracleResponse)
				if err := _OCR2DROracle.contract.UnpackLog(event, "OracleResponse", log); err != nil {
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

func (_OCR2DROracle *OCR2DROracleFilterer) ParseOracleResponse(log types.Log) (*OCR2DROracleOracleResponse, error) {
	event := new(OCR2DROracleOracleResponse)
	if err := _OCR2DROracle.contract.UnpackLog(event, "OracleResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2DROracleOwnershipTransferRequestedIterator struct {
	Event *OCR2DROracleOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2DROracleOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2DROracleOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(OCR2DROracleOwnershipTransferRequested)
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

func (it *OCR2DROracleOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OCR2DROracleOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2DROracleOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OCR2DROracle *OCR2DROracleFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OCR2DROracleOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OCR2DROracle.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OCR2DROracleOwnershipTransferRequestedIterator{contract: _OCR2DROracle.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OCR2DROracle *OCR2DROracleFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OCR2DROracleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OCR2DROracle.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2DROracleOwnershipTransferRequested)
				if err := _OCR2DROracle.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OCR2DROracle *OCR2DROracleFilterer) ParseOwnershipTransferRequested(log types.Log) (*OCR2DROracleOwnershipTransferRequested, error) {
	event := new(OCR2DROracleOwnershipTransferRequested)
	if err := _OCR2DROracle.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2DROracleOwnershipTransferredIterator struct {
	Event *OCR2DROracleOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2DROracleOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2DROracleOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(OCR2DROracleOwnershipTransferred)
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

func (it *OCR2DROracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OCR2DROracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2DROracleOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OCR2DROracle *OCR2DROracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OCR2DROracleOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OCR2DROracle.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OCR2DROracleOwnershipTransferredIterator{contract: _OCR2DROracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OCR2DROracle *OCR2DROracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OCR2DROracleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OCR2DROracle.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2DROracleOwnershipTransferred)
				if err := _OCR2DROracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OCR2DROracle *OCR2DROracleFilterer) ParseOwnershipTransferred(log types.Log) (*OCR2DROracleOwnershipTransferred, error) {
	event := new(OCR2DROracleOwnershipTransferred)
	if err := _OCR2DROracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_OCR2DROracle *OCR2DROracle) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OCR2DROracle.abi.Events["AuthorizedSendersChanged"].ID:
		return _OCR2DROracle.ParseAuthorizedSendersChanged(log)
	case _OCR2DROracle.abi.Events["OracleRequest"].ID:
		return _OCR2DROracle.ParseOracleRequest(log)
	case _OCR2DROracle.abi.Events["OracleResponse"].ID:
		return _OCR2DROracle.ParseOracleResponse(log)
	case _OCR2DROracle.abi.Events["OwnershipTransferRequested"].ID:
		return _OCR2DROracle.ParseOwnershipTransferRequested(log)
	case _OCR2DROracle.abi.Events["OwnershipTransferred"].ID:
		return _OCR2DROracle.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OCR2DROracleAuthorizedSendersChanged) Topic() common.Hash {
	return common.HexToHash("0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0")
}

func (OCR2DROracleOracleRequest) Topic() common.Hash {
	return common.HexToHash("0x9dc59d1e6d6042f6c2d2af5a3f9f6502bf0abf7cf40f0eb80edf752ca396f2ba")
}

func (OCR2DROracleOracleResponse) Topic() common.Hash {
	return common.HexToHash("0x9e9bc7616d42c2835d05ae617e508454e63b30b934be8aa932ebc125e0e58a64")
}

func (OCR2DROracleOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OCR2DROracleOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_OCR2DROracle *OCR2DROracle) Address() common.Address {
	return _OCR2DROracle.address
}

type OCR2DROracleInterface interface {
	GetAuthorizedSenders(opts *bind.CallOpts) ([]common.Address, error)

	GetDONPublicKey(opts *bind.CallOpts) ([]byte, error)

	IsAuthorizedSender(opts *bind.CallOpts, sender common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	FulfillRequest(opts *bind.TransactOpts, requestId [32]byte, response []byte, err []byte) (*types.Transaction, error)

	SendRequest(opts *bind.TransactOpts, subscriptionId *big.Int, data []byte) (*types.Transaction, error)

	SetAuthorizedSenders(opts *bind.TransactOpts, senders []common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAuthorizedSendersChanged(opts *bind.FilterOpts) (*OCR2DROracleAuthorizedSendersChangedIterator, error)

	WatchAuthorizedSendersChanged(opts *bind.WatchOpts, sink chan<- *OCR2DROracleAuthorizedSendersChanged) (event.Subscription, error)

	ParseAuthorizedSendersChanged(log types.Log) (*OCR2DROracleAuthorizedSendersChanged, error)

	FilterOracleRequest(opts *bind.FilterOpts) (*OCR2DROracleOracleRequestIterator, error)

	WatchOracleRequest(opts *bind.WatchOpts, sink chan<- *OCR2DROracleOracleRequest) (event.Subscription, error)

	ParseOracleRequest(log types.Log) (*OCR2DROracleOracleRequest, error)

	FilterOracleResponse(opts *bind.FilterOpts) (*OCR2DROracleOracleResponseIterator, error)

	WatchOracleResponse(opts *bind.WatchOpts, sink chan<- *OCR2DROracleOracleResponse) (event.Subscription, error)

	ParseOracleResponse(log types.Log) (*OCR2DROracleOracleResponse, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OCR2DROracleOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OCR2DROracleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OCR2DROracleOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OCR2DROracleOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OCR2DROracleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OCR2DROracleOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}