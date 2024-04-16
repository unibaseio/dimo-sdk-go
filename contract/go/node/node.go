// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package node

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

// INodePledgeInfo is an auto generated low-level Go binding around an user-defined struct.
type INodePledgeInfo struct {
	Time  *big.Int
	Value *big.Int
}

// NodeMetaData contains all meta data concerning the Node contract.
var NodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Pledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Punish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_detail\",\"type\":\"bytes\"}],\"name\":\"declare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"getMinPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"getPledge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structINode.PledgeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"getPledgeValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260646002553480156200001657600080fd5b5060405162001ee838038062001ee883398181016040528101906200003c9190620001da565b6200005c62000050620000a460201b60201c565b620000ac60201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200020c565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001a28262000175565b9050919050565b620001b48162000195565b8114620001c057600080fd5b50565b600081519050620001d481620001a9565b92915050565b600060208284031215620001f357620001f262000170565b5b60006200020384828501620001c3565b91505092915050565b611ccc806200021c6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638da5cb5b1161008c57806398c365d51161006657806398c365d51461020d5780639fa6a6e31461023d578063e3fdacb61461025b578063f2fde38b1461028b576100ea565b80638da5cb5b146101b757806390d8cf07146101d55780639708cab3146101f1576100ea565b806348ab5e6c116100c857806348ab5e6c1461015757806361e728b114610173578063715018a61461018f57806376cdb03b14610199576100ea565b806306f0b4f1146100ef5780633f4899141461010b5780634339ceca14610127575b600080fd5b61010960048036038101906101049190611134565b6102a7565b005b6101256004803603810190610120919061119b565b61051c565b005b610141600480360381019061013c91906111db565b61073f565b60405161014e919061122a565b60405180910390f35b610171600480360381019061016c919061119b565b6107a3565b005b61018d600480360381019061018891906111db565b610806565b005b6101976108bf565b005b6101a16108d3565b6040516101ae9190611254565b60405180910390f35b6101bf6108f9565b6040516101cc9190611254565b60405180910390f35b6101ef60048036038101906101ea91906113b5565b610922565b005b61020b6004803603810190610206919061119b565b610975565b005b610227600480360381019061022291906113fe565b610bb0565b604051610234919061122a565b60405180910390f35b610245610bd3565b604051610252919061144e565b60405180910390f35b610275600480360381019061027091906111db565b610bed565b60405161028291906114a7565b60405180910390f35b6102a560048036038101906102a091906114c2565b610cd0565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016103009061154c565b6020604051808303816000875af115801561031f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103439190611581565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a7906115fa565b60405180910390fd5b80600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008560ff1660ff16815260200190815260200160002060010160008282546104199190611649565b92505081905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166376890c5883836040518363ffffffff1660e01b815260040161047d92919061167d565b600060405180830381600087803b15801561049757600080fd5b505af11580156104ab573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32858460405161050e9291906116b5565b60405180910390a350505050565b610524610d53565b600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff16600254600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008560ff1660ff168152602001908152602001600020600001546105ac91906116de565b106105ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e39061175e565b60405180910390fd5b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008460ff1660ff16815260200190815260200160002060010160008282546106559190611649565b92505081905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166376890c5833836040518363ffffffff1660e01b81526004016106b992919061167d565b600060405180830381600087803b1580156106d357600080fd5b505af11580156106e7573d6000803e3d6000fd5b505050503373ffffffffffffffffffffffffffffffffffffffff167fe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d7483836040516107339291906116b5565b60405180910390a25050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008360ff1660ff16815260200190815260200160002060010154905092915050565b6107ab610eef565b80600360008460ff1660ff168152602001908152602001600020819055507fc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef6788182826040516107fa9291906116b5565b60405180910390a15050565b600360008260ff1660ff16815260200190815260200160002054600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008360ff1660ff1681526020019081526020016000206001015410156108bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b2906117ca565b60405180910390fd5b5050565b6108c7610eef565b6108d16000610f6d565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000190816109719190611a01565b5050565b61097d610d53565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e888891533836040518363ffffffff1660e01b81526004016109da92919061167d565b600060405180830381600087803b1580156109f457600080fd5b505af1158015610a08573d6000803e3d6000fd5b505050506000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008460ff1660ff1681526020019081526020016000206000015403610aec57600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff16600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008460ff1660ff168152602001908152602001600020600001819055505b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008460ff1660ff1681526020019081526020016000206001016000828254610b5591906116de565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167fedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee8383604051610ba49291906116b5565b60405180910390a25050565b6000600360008360ff1660ff168152602001908152602001600020549050919050565b600160149054906101000a900467ffffffffffffffff1681565b610bf5611039565b610bfd611039565b6040518060400160405280600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008660ff1660ff168152602001908152602001600020600001548152602001600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008660ff1660ff1681526020019081526020016000206001015481525090508091505092915050565b610cd8610eef565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610d47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3e90611b45565b60405180910390fd5b610d5081610f6d565b50565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610dae90611bb1565b6020604051808303816000875af1158015610dcd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df19190611581565b90508073ffffffffffffffffffffffffffffffffffffffff1663919840ad6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610e3b57600080fd5b505af1158015610e4f573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16639fa6a6e36040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610ea0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec49190611bfd565b600160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b610ef7611031565b73ffffffffffffffffffffffffffffffffffffffff16610f156108f9565b73ffffffffffffffffffffffffffffffffffffffff1614610f6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6290611c76565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b604051806040016040528060008152602001600081525090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061109282611067565b9050919050565b6110a281611087565b81146110ad57600080fd5b50565b6000813590506110bf81611099565b92915050565b600060ff82169050919050565b6110db816110c5565b81146110e657600080fd5b50565b6000813590506110f8816110d2565b92915050565b6000819050919050565b611111816110fe565b811461111c57600080fd5b50565b60008135905061112e81611108565b92915050565b6000806000806080858703121561114e5761114d61105d565b5b600061115c878288016110b0565b945050602061116d878288016110e9565b935050604061117e878288016110b0565b925050606061118f8782880161111f565b91505092959194509250565b600080604083850312156111b2576111b161105d565b5b60006111c0858286016110e9565b92505060206111d18582860161111f565b9150509250929050565b600080604083850312156111f2576111f161105d565b5b6000611200858286016110b0565b9250506020611211858286016110e9565b9150509250929050565b611224816110fe565b82525050565b600060208201905061123f600083018461121b565b92915050565b61124e81611087565b82525050565b60006020820190506112696000830184611245565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6112c282611279565b810181811067ffffffffffffffff821117156112e1576112e061128a565b5b80604052505050565b60006112f4611053565b905061130082826112b9565b919050565b600067ffffffffffffffff8211156113205761131f61128a565b5b61132982611279565b9050602081019050919050565b82818337600083830152505050565b600061135861135384611305565b6112ea565b90508281526020810184848401111561137457611373611274565b5b61137f848285611336565b509392505050565b600082601f83011261139c5761139b61126f565b5b81356113ac848260208601611345565b91505092915050565b6000602082840312156113cb576113ca61105d565b5b600082013567ffffffffffffffff8111156113e9576113e8611062565b5b6113f584828501611387565b91505092915050565b6000602082840312156114145761141361105d565b5b6000611422848285016110e9565b91505092915050565b600067ffffffffffffffff82169050919050565b6114488161142b565b82525050565b6000602082019050611463600083018461143f565b92915050565b611472816110fe565b82525050565b60408201600082015161148e6000850182611469565b5060208201516114a16020850182611469565b50505050565b60006040820190506114bc6000830184611478565b92915050565b6000602082840312156114d8576114d761105d565b5b60006114e6848285016110b0565b91505092915050565b600082825260208201905092915050565b7f636f6e74726f6c00000000000000000000000000000000000000000000000000600082015250565b60006115366007836114ef565b915061154182611500565b602082019050919050565b6000602082019050818103600083015261156581611529565b9050919050565b60008151905061157b81611099565b92915050565b6000602082840312156115975761159661105d565b5b60006115a58482850161156c565b91505092915050565b7f696e76616c696420630000000000000000000000000000000000000000000000600082015250565b60006115e46009836114ef565b91506115ef826115ae565b602082019050919050565b60006020820190508181036000830152611613816115d7565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611654826110fe565b915061165f836110fe565b92508282039050818111156116775761167661161a565b5b92915050565b60006040820190506116926000830185611245565b61169f602083018461121b565b9392505050565b6116af816110c5565b82525050565b60006040820190506116ca60008301856116a6565b6116d7602083018461121b565b9392505050565b60006116e9826110fe565b91506116f4836110fe565b925082820190508082111561170c5761170b61161a565b5b92915050565b7f6c6f636b65640000000000000000000000000000000000000000000000000000600082015250565b60006117486006836114ef565b915061175382611712565b602082019050919050565b600060208201905081810360008301526117778161173b565b9050919050565b7f696e73756620706c656467650000000000000000000000000000000000000000600082015250565b60006117b4600c836114ef565b91506117bf8261177e565b602082019050919050565b600060208201905081810360008301526117e3816117a7565b9050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061183c57607f821691505b60208210810361184f5761184e6117f5565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026118b77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261187a565b6118c1868361187a565b95508019841693508086168417925050509392505050565b6000819050919050565b60006118fe6118f96118f4846110fe565b6118d9565b6110fe565b9050919050565b6000819050919050565b611918836118e3565b61192c61192482611905565b848454611887565b825550505050565b600090565b611941611934565b61194c81848461190f565b505050565b5b8181101561197057611965600082611939565b600181019050611952565b5050565b601f8211156119b55761198681611855565b61198f8461186a565b8101602085101561199e578190505b6119b26119aa8561186a565b830182611951565b50505b505050565b600082821c905092915050565b60006119d8600019846008026119ba565b1980831691505092915050565b60006119f183836119c7565b9150826002028217905092915050565b611a0a826117ea565b67ffffffffffffffff811115611a2357611a2261128a565b5b611a2d8254611824565b611a38828285611974565b600060209050601f831160018114611a6b5760008415611a59578287015190505b611a6385826119e5565b865550611acb565b601f198416611a7986611855565b60005b82811015611aa157848901518255600182019150602085019450602081019050611a7c565b86831015611abe5784890151611aba601f8916826119c7565b8355505b6001600288020188555050505b505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611b2f6026836114ef565b9150611b3a82611ad3565b604082019050919050565b60006020820190508181036000830152611b5e81611b22565b9050919050565b7f65706f6368000000000000000000000000000000000000000000000000000000600082015250565b6000611b9b6005836114ef565b9150611ba682611b65565b602082019050919050565b60006020820190508181036000830152611bca81611b8e565b9050919050565b611bda8161142b565b8114611be557600080fd5b50565b600081519050611bf781611bd1565b92915050565b600060208284031215611c1357611c1261105d565b5b6000611c2184828501611be8565b91505092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611c606020836114ef565b9150611c6b82611c2a565b602082019050919050565b60006020820190508181036000830152611c8f81611c53565b905091905056fea26469706673582212202672723642deaa317e0ce4a8001e59a0955994c8d93619b07dc00bb4107252b764736f6c63430008180033",
}

// NodeABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeMetaData.ABI instead.
var NodeABI = NodeMetaData.ABI

// NodeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NodeMetaData.Bin instead.
var NodeBin = NodeMetaData.Bin

// DeployNode deploys a new Ethereum contract, binding an instance of Node to it.
func DeployNode(auth *bind.TransactOpts, backend bind.ContractBackend, _b common.Address) (common.Address, *types.Transaction, *Node, error) {
	parsed, err := NodeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NodeBin), backend, _b)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// Node is an auto generated Go binding around an Ethereum contract.
type Node struct {
	NodeCaller     // Read-only binding to the contract
	NodeTransactor // Write-only binding to the contract
	NodeFilterer   // Log filterer for contract events
}

// NodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeSession struct {
	Contract     *Node             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeCallerSession struct {
	Contract *NodeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeTransactorSession struct {
	Contract     *NodeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRaw struct {
	Contract *Node // Generic contract binding to access the raw methods on
}

// NodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeCallerRaw struct {
	Contract *NodeCaller // Generic read-only contract binding to access the raw methods on
}

// NodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeTransactorRaw struct {
	Contract *NodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNode creates a new instance of Node, bound to a specific deployed contract.
func NewNode(address common.Address, backend bind.ContractBackend) (*Node, error) {
	contract, err := bindNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// NewNodeCaller creates a new read-only instance of Node, bound to a specific deployed contract.
func NewNodeCaller(address common.Address, caller bind.ContractCaller) (*NodeCaller, error) {
	contract, err := bindNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeCaller{contract: contract}, nil
}

// NewNodeTransactor creates a new write-only instance of Node, bound to a specific deployed contract.
func NewNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeTransactor, error) {
	contract, err := bindNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeTransactor{contract: contract}, nil
}

// NewNodeFilterer creates a new log filterer instance of Node, bound to a specific deployed contract.
func NewNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFilterer, error) {
	contract, err := bindNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFilterer{contract: contract}, nil
}

// bindNode binds a generic wrapper to an already deployed contract.
func bindNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.NodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.contract.Transact(opts, method, params...)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Node *NodeCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Node *NodeSession) Bank() (common.Address, error) {
	return _Node.Contract.Bank(&_Node.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Node *NodeCallerSession) Bank() (common.Address, error) {
	return _Node.Contract.Bank(&_Node.CallOpts)
}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns()
func (_Node *NodeCaller) Check(opts *bind.CallOpts, _a common.Address, _type uint8) error {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "check", _a, _type)

	if err != nil {
		return err
	}

	return err

}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns()
func (_Node *NodeSession) Check(_a common.Address, _type uint8) error {
	return _Node.Contract.Check(&_Node.CallOpts, _a, _type)
}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns()
func (_Node *NodeCallerSession) Check(_a common.Address, _type uint8) error {
	return _Node.Contract.Check(&_Node.CallOpts, _a, _type)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Node *NodeCaller) Current(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "current")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Node *NodeSession) Current() (uint64, error) {
	return _Node.Contract.Current(&_Node.CallOpts)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Node *NodeCallerSession) Current() (uint64, error) {
	return _Node.Contract.Current(&_Node.CallOpts)
}

// GetMinPledge is a free data retrieval call binding the contract method 0x98c365d5.
//
// Solidity: function getMinPledge(uint8 _type) view returns(uint256)
func (_Node *NodeCaller) GetMinPledge(opts *bind.CallOpts, _type uint8) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getMinPledge", _type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinPledge is a free data retrieval call binding the contract method 0x98c365d5.
//
// Solidity: function getMinPledge(uint8 _type) view returns(uint256)
func (_Node *NodeSession) GetMinPledge(_type uint8) (*big.Int, error) {
	return _Node.Contract.GetMinPledge(&_Node.CallOpts, _type)
}

// GetMinPledge is a free data retrieval call binding the contract method 0x98c365d5.
//
// Solidity: function getMinPledge(uint8 _type) view returns(uint256)
func (_Node *NodeCallerSession) GetMinPledge(_type uint8) (*big.Int, error) {
	return _Node.Contract.GetMinPledge(&_Node.CallOpts, _type)
}

// GetPledge is a free data retrieval call binding the contract method 0xe3fdacb6.
//
// Solidity: function getPledge(address _a, uint8 _type) view returns((uint256,uint256))
func (_Node *NodeCaller) GetPledge(opts *bind.CallOpts, _a common.Address, _type uint8) (INodePledgeInfo, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getPledge", _a, _type)

	if err != nil {
		return *new(INodePledgeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(INodePledgeInfo)).(*INodePledgeInfo)

	return out0, err

}

// GetPledge is a free data retrieval call binding the contract method 0xe3fdacb6.
//
// Solidity: function getPledge(address _a, uint8 _type) view returns((uint256,uint256))
func (_Node *NodeSession) GetPledge(_a common.Address, _type uint8) (INodePledgeInfo, error) {
	return _Node.Contract.GetPledge(&_Node.CallOpts, _a, _type)
}

// GetPledge is a free data retrieval call binding the contract method 0xe3fdacb6.
//
// Solidity: function getPledge(address _a, uint8 _type) view returns((uint256,uint256))
func (_Node *NodeCallerSession) GetPledge(_a common.Address, _type uint8) (INodePledgeInfo, error) {
	return _Node.Contract.GetPledge(&_Node.CallOpts, _a, _type)
}

// GetPledgeValue is a free data retrieval call binding the contract method 0x4339ceca.
//
// Solidity: function getPledgeValue(address _a, uint8 _type) view returns(uint256)
func (_Node *NodeCaller) GetPledgeValue(opts *bind.CallOpts, _a common.Address, _type uint8) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getPledgeValue", _a, _type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledgeValue is a free data retrieval call binding the contract method 0x4339ceca.
//
// Solidity: function getPledgeValue(address _a, uint8 _type) view returns(uint256)
func (_Node *NodeSession) GetPledgeValue(_a common.Address, _type uint8) (*big.Int, error) {
	return _Node.Contract.GetPledgeValue(&_Node.CallOpts, _a, _type)
}

// GetPledgeValue is a free data retrieval call binding the contract method 0x4339ceca.
//
// Solidity: function getPledgeValue(address _a, uint8 _type) view returns(uint256)
func (_Node *NodeCallerSession) GetPledgeValue(_a common.Address, _type uint8) (*big.Int, error) {
	return _Node.Contract.GetPledgeValue(&_Node.CallOpts, _a, _type)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeSession) Owner() (common.Address, error) {
	return _Node.Contract.Owner(&_Node.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeCallerSession) Owner() (common.Address, error) {
	return _Node.Contract.Owner(&_Node.CallOpts)
}

// Declare is a paid mutator transaction binding the contract method 0x90d8cf07.
//
// Solidity: function declare(bytes _detail) returns()
func (_Node *NodeTransactor) Declare(opts *bind.TransactOpts, _detail []byte) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "declare", _detail)
}

// Declare is a paid mutator transaction binding the contract method 0x90d8cf07.
//
// Solidity: function declare(bytes _detail) returns()
func (_Node *NodeSession) Declare(_detail []byte) (*types.Transaction, error) {
	return _Node.Contract.Declare(&_Node.TransactOpts, _detail)
}

// Declare is a paid mutator transaction binding the contract method 0x90d8cf07.
//
// Solidity: function declare(bytes _detail) returns()
func (_Node *NodeTransactorSession) Declare(_detail []byte) (*types.Transaction, error) {
	return _Node.Contract.Declare(&_Node.TransactOpts, _detail)
}

// Pledge is a paid mutator transaction binding the contract method 0x9708cab3.
//
// Solidity: function pledge(uint8 _typ, uint256 _m) returns()
func (_Node *NodeTransactor) Pledge(opts *bind.TransactOpts, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "pledge", _typ, _m)
}

// Pledge is a paid mutator transaction binding the contract method 0x9708cab3.
//
// Solidity: function pledge(uint8 _typ, uint256 _m) returns()
func (_Node *NodeSession) Pledge(_typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Pledge(&_Node.TransactOpts, _typ, _m)
}

// Pledge is a paid mutator transaction binding the contract method 0x9708cab3.
//
// Solidity: function pledge(uint8 _typ, uint256 _m) returns()
func (_Node *NodeTransactorSession) Pledge(_typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Pledge(&_Node.TransactOpts, _typ, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _from, uint8 _typ, address _to, uint256 _m) returns()
func (_Node *NodeTransactor) Punish(opts *bind.TransactOpts, _from common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "punish", _from, _typ, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _from, uint8 _typ, address _to, uint256 _m) returns()
func (_Node *NodeSession) Punish(_from common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Punish(&_Node.TransactOpts, _from, _typ, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _from, uint8 _typ, address _to, uint256 _m) returns()
func (_Node *NodeTransactorSession) Punish(_from common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Punish(&_Node.TransactOpts, _from, _typ, _to, _m)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Node.Contract.RenounceOwnership(&_Node.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Node.Contract.RenounceOwnership(&_Node.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x48ab5e6c.
//
// Solidity: function set(uint8 _type, uint256 _money) returns()
func (_Node *NodeTransactor) Set(opts *bind.TransactOpts, _type uint8, _money *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "set", _type, _money)
}

// Set is a paid mutator transaction binding the contract method 0x48ab5e6c.
//
// Solidity: function set(uint8 _type, uint256 _money) returns()
func (_Node *NodeSession) Set(_type uint8, _money *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Set(&_Node.TransactOpts, _type, _money)
}

// Set is a paid mutator transaction binding the contract method 0x48ab5e6c.
//
// Solidity: function set(uint8 _type, uint256 _money) returns()
func (_Node *NodeTransactorSession) Set(_type uint8, _money *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Set(&_Node.TransactOpts, _type, _money)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.TransferOwnership(&_Node.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.TransferOwnership(&_Node.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 _typ, uint256 _m) returns()
func (_Node *NodeTransactor) Withdraw(opts *bind.TransactOpts, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "withdraw", _typ, _m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 _typ, uint256 _m) returns()
func (_Node *NodeSession) Withdraw(_typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Withdraw(&_Node.TransactOpts, _typ, _m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 _typ, uint256 _m) returns()
func (_Node *NodeTransactorSession) Withdraw(_typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Withdraw(&_Node.TransactOpts, _typ, _m)
}

// NodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Node contract.
type NodeOwnershipTransferredIterator struct {
	Event *NodeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOwnershipTransferred)
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
		it.Event = new(NodeOwnershipTransferred)
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
func (it *NodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOwnershipTransferred represents a OwnershipTransferred event raised by the Node contract.
type NodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Node *NodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeOwnershipTransferredIterator{contract: _Node.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Node *NodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOwnershipTransferred)
				if err := _Node.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Node *NodeFilterer) ParseOwnershipTransferred(log types.Log) (*NodeOwnershipTransferred, error) {
	event := new(NodeOwnershipTransferred)
	if err := _Node.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodePledgeIterator is returned from FilterPledge and is used to iterate over the raw logs and unpacked data for Pledge events raised by the Node contract.
type NodePledgeIterator struct {
	Event *NodePledge // Event containing the contract specifics and raw log

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
func (it *NodePledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodePledge)
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
		it.Event = new(NodePledge)
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
func (it *NodePledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodePledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodePledge represents a Pledge event raised by the Node contract.
type NodePledge struct {
	A     common.Address
	Type  uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPledge is a free log retrieval operation binding the contract event 0xedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee.
//
// Solidity: event Pledge(address indexed _a, uint8 _type, uint256 _money)
func (_Node *NodeFilterer) FilterPledge(opts *bind.FilterOpts, _a []common.Address) (*NodePledgeIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Pledge", _aRule)
	if err != nil {
		return nil, err
	}
	return &NodePledgeIterator{contract: _Node.contract, event: "Pledge", logs: logs, sub: sub}, nil
}

// WatchPledge is a free log subscription operation binding the contract event 0xedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee.
//
// Solidity: event Pledge(address indexed _a, uint8 _type, uint256 _money)
func (_Node *NodeFilterer) WatchPledge(opts *bind.WatchOpts, sink chan<- *NodePledge, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Pledge", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodePledge)
				if err := _Node.contract.UnpackLog(event, "Pledge", log); err != nil {
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

// ParsePledge is a log parse operation binding the contract event 0xedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee.
//
// Solidity: event Pledge(address indexed _a, uint8 _type, uint256 _money)
func (_Node *NodeFilterer) ParsePledge(log types.Log) (*NodePledge, error) {
	event := new(NodePledge)
	if err := _Node.contract.UnpackLog(event, "Pledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodePunishIterator is returned from FilterPunish and is used to iterate over the raw logs and unpacked data for Punish events raised by the Node contract.
type NodePunishIterator struct {
	Event *NodePunish // Event containing the contract specifics and raw log

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
func (it *NodePunishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodePunish)
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
		it.Event = new(NodePunish)
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
func (it *NodePunishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodePunishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodePunish represents a Punish event raised by the Node contract.
type NodePunish struct {
	A     common.Address
	Typ   uint8
	To    common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPunish is a free log retrieval operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_Node *NodeFilterer) FilterPunish(opts *bind.FilterOpts, _a []common.Address, _to []common.Address) (*NodePunishIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Punish", _aRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &NodePunishIterator{contract: _Node.contract, event: "Punish", logs: logs, sub: sub}, nil
}

// WatchPunish is a free log subscription operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_Node *NodeFilterer) WatchPunish(opts *bind.WatchOpts, sink chan<- *NodePunish, _a []common.Address, _to []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Punish", _aRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodePunish)
				if err := _Node.contract.UnpackLog(event, "Punish", log); err != nil {
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

// ParsePunish is a log parse operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_Node *NodeFilterer) ParsePunish(log types.Log) (*NodePunish, error) {
	event := new(NodePunish)
	if err := _Node.contract.UnpackLog(event, "Punish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the Node contract.
type NodeSetIterator struct {
	Event *NodeSet // Event containing the contract specifics and raw log

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
func (it *NodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeSet)
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
		it.Event = new(NodeSet)
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
func (it *NodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeSet represents a Set event raised by the Node contract.
type NodeSet struct {
	Type  uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _money)
func (_Node *NodeFilterer) FilterSet(opts *bind.FilterOpts) (*NodeSetIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &NodeSetIterator{contract: _Node.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _money)
func (_Node *NodeFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *NodeSet) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeSet)
				if err := _Node.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _money)
func (_Node *NodeFilterer) ParseSet(log types.Log) (*NodeSet, error) {
	event := new(NodeSet)
	if err := _Node.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Node contract.
type NodeWithdrawIterator struct {
	Event *NodeWithdraw // Event containing the contract specifics and raw log

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
func (it *NodeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeWithdraw)
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
		it.Event = new(NodeWithdraw)
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
func (it *NodeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeWithdraw represents a Withdraw event raised by the Node contract.
type NodeWithdraw struct {
	A     common.Address
	Type  uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d74.
//
// Solidity: event Withdraw(address indexed _a, uint8 _type, uint256 _money)
func (_Node *NodeFilterer) FilterWithdraw(opts *bind.FilterOpts, _a []common.Address) (*NodeWithdrawIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return &NodeWithdrawIterator{contract: _Node.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d74.
//
// Solidity: event Withdraw(address indexed _a, uint8 _type, uint256 _money)
func (_Node *NodeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *NodeWithdraw, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeWithdraw)
				if err := _Node.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d74.
//
// Solidity: event Withdraw(address indexed _a, uint8 _type, uint256 _money)
func (_Node *NodeFilterer) ParseWithdraw(log types.Log) (*NodeWithdraw, error) {
	event := new(NodeWithdraw)
	if err := _Node.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
