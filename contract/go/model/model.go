// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package model

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

// IModelInfo is an auto generated low-level Go binding around an user-defined struct.
type IModelInfo struct {
	Name   string
	Owner  common.Address
	Root   []byte
	Count  [32]byte
	Active bool
}

// ModelMetaData contains all meta data concerning the Model contract.
var ModelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_mn\",\"type\":\"string\"}],\"name\":\"AddModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_s\",\"type\":\"bool\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_mn\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_rt\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_cnt\",\"type\":\"bytes32\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"c4vk\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_mn\",\"type\":\"string\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"getModel\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"count\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structIModel.Info\",\"name\":\"_info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVK\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_vk\",\"type\":\"bytes\"}],\"name\":\"setVK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620020ad380380620020ad8339818101604052810190620000379190620001d5565b620000576200004b6200009f60201b60201c565b620000a760201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000207565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200019d8262000170565b9050919050565b620001af8162000190565b8114620001bb57600080fd5b50565b600081519050620001cf81620001a4565b92915050565b600060208284031215620001ee57620001ed6200016b565b5b6000620001fe84828501620001be565b91505092915050565b611e9680620002176000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063a35176d01161008c578063d48dca2d11610066578063d48dca2d1461023a578063ea1bbe3514610256578063f2fde38b14610286578063f804843d146102a2576100ea565b8063a35176d0146101d0578063d0b9178b146101ee578063d370a37d1461020a576100ea565b8063715018a6116100c8578063715018a61461015a57806376cdb03b146101645780638da5cb5b146101825780639f5591ab146101a0576100ea565b8063120a4734146100ef578063354bb9ba1461010b5780636c37609a14610129575b600080fd5b6101096004803603810190610104919061110a565b6102be565b005b61011361056b565b6040516101209190611214565b60405180910390f35b610143600480360381019061013e9190611276565b6105fd565b6040516101519291906112b2565b60405180910390f35b6101626106f5565b005b61016c610709565b6040516101799190611323565b60405180910390f35b61018a61072f565b6040516101979190611323565b60405180910390f35b6101ba60048036038101906101b59190611276565b610758565b6040516101c79190611493565b60405180910390f35b6101d86109f1565b6040516101e59190611214565b60405180910390f35b610208600480360381019061020391906114b5565b610a7f565b005b610224600480360381019061021f9190611276565b610a9a565b6040516102319190611323565b60405180910390f35b610254600480360381019061024f919061152a565b610af3565b005b610270600480360381019061026b919061156a565b610bfb565b60405161027d91906115c2565b60405180910390f35b6102a0600480360381019061029b9190611609565b610c43565b005b6102bc60048036038101906102b79190611276565b610cc6565b005b6000801b8111610303576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fa90611693565b60405180910390fd5b600060048460405161031591906116ef565b908152602001604051809103902060009054906101000a900467ffffffffffffffff1667ffffffffffffffff1614610382576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037990611752565b60405180910390fd5b61038a610e8f565b83816000018190525033816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050828160400181905250818160600181815250506001816080019015159081151581525050600381908060018154018082558091505060019003906000526020600020906005020160009091909190915060008201518160000190816104309190611988565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201908161048d9190611ab5565b506060820151816003015560808201518160040160006101000a81548160ff02191690831515021790555050506003805490506004856040516104d091906116ef565b908152602001604051809103902060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f1e0210396cbbc2a1ccae7d7b17e27a3fea0e56962c68454eda6fa9682bc7fe47600160038054905061054e9190611bb6565b8660405161055d929190611c2b565b60405180910390a250505050565b60606002805461057a906117a1565b80601f01602080910402602001604051908101604052809291908181526020018280546105a6906117a1565b80156105f35780601f106105c8576101008083540402835291602001916105f3565b820191906000526020600020905b8154815290600101906020018083116105d657829003601f168201915b5050505050905090565b6000606060038367ffffffffffffffff168154811061061f5761061e611c5b565b5b90600052602060002090600502016003015460038467ffffffffffffffff168154811061064f5761064e611c5b565b5b906000526020600020906005020160020180805461066c906117a1565b80601f0160208091040260200160405190810160405280929190818152602001828054610698906117a1565b80156106e55780601f106106ba576101008083540402835291602001916106e5565b820191906000526020600020905b8154815290600101906020018083116106c857829003601f168201915b5050505050905091509150915091565b6106fd610d45565b6107076000610dc3565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610760610e8f565b60038267ffffffffffffffff168154811061077e5761077d611c5b565b5b9060005260206000209060050201600001805461079a906117a1565b80601f01602080910402602001604051908101604052809291908181526020018280546107c6906117a1565b80156108135780601f106107e857610100808354040283529160200191610813565b820191906000526020600020905b8154815290600101906020018083116107f657829003601f168201915b5050505050816000018190525060038267ffffffffffffffff168154811061083e5761083d611c5b565b5b906000526020600020906005020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060038267ffffffffffffffff16815481106108c5576108c4611c5b565b5b906000526020600020906005020160020180546108e1906117a1565b80601f016020809104026020016040519081016040528092919081815260200182805461090d906117a1565b801561095a5780601f1061092f5761010080835404028352916020019161095a565b820191906000526020600020905b81548152906001019060200180831161093d57829003601f168201915b5050505050816040018190525060038267ffffffffffffffff168154811061098557610984611c5b565b5b90600052602060002090600502016003015481606001818152505060038267ffffffffffffffff16815481106109be576109bd611c5b565b5b906000526020600020906005020160040160009054906101000a900460ff16816080019015159081151581525050919050565b600280546109fe906117a1565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2a906117a1565b8015610a775780601f10610a4c57610100808354040283529160200191610a77565b820191906000526020600020905b815481529060010190602001808311610a5a57829003601f168201915b505050505081565b610a87610d45565b8060029081610a969190611ab5565b5050565b600060038267ffffffffffffffff1681548110610aba57610ab9611c5b565b5b906000526020600020906005020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b3373ffffffffffffffffffffffffffffffffffffffff1660038367ffffffffffffffff1681548110610b2857610b27611c5b565b5b906000526020600020906005020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610bb0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ba790611cd6565b60405180910390fd5b8060038367ffffffffffffffff1681548110610bcf57610bce611c5b565b5b906000526020600020906005020160040160006101000a81548160ff0219169083151502179055505050565b60006001600483604051610c0f91906116ef565b908152602001604051809103902060009054906101000a900467ffffffffffffffff16610c3c9190611bb6565b9050919050565b610c4b610d45565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610cba576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cb190611d68565b60405180910390fd5b610cc381610dc3565b50565b60038167ffffffffffffffff1681548110610ce457610ce3611c5b565b5b906000526020600020906005020160040160009054906101000a900460ff16610d42576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3990611dd4565b60405180910390fd5b50565b610d4d610e87565b73ffffffffffffffffffffffffffffffffffffffff16610d6b61072f565b73ffffffffffffffffffffffffffffffffffffffff1614610dc1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610db890611e40565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6040518060a0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600080191681526020016000151581525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f4082610ef7565b810181811067ffffffffffffffff82111715610f5f57610f5e610f08565b5b80604052505050565b6000610f72610ed9565b9050610f7e8282610f37565b919050565b600067ffffffffffffffff821115610f9e57610f9d610f08565b5b610fa782610ef7565b9050602081019050919050565b82818337600083830152505050565b6000610fd6610fd184610f83565b610f68565b905082815260208101848484011115610ff257610ff1610ef2565b5b610ffd848285610fb4565b509392505050565b600082601f83011261101a57611019610eed565b5b813561102a848260208601610fc3565b91505092915050565b600067ffffffffffffffff82111561104e5761104d610f08565b5b61105782610ef7565b9050602081019050919050565b600061107761107284611033565b610f68565b90508281526020810184848401111561109357611092610ef2565b5b61109e848285610fb4565b509392505050565b600082601f8301126110bb576110ba610eed565b5b81356110cb848260208601611064565b91505092915050565b6000819050919050565b6110e7816110d4565b81146110f257600080fd5b50565b600081359050611104816110de565b92915050565b60008060006060848603121561112357611122610ee3565b5b600084013567ffffffffffffffff81111561114157611140610ee8565b5b61114d86828701611005565b935050602084013567ffffffffffffffff81111561116e5761116d610ee8565b5b61117a868287016110a6565b925050604061118b868287016110f5565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b60005b838110156111cf5780820151818401526020810190506111b4565b60008484015250505050565b60006111e682611195565b6111f081856111a0565b93506112008185602086016111b1565b61120981610ef7565b840191505092915050565b6000602082019050818103600083015261122e81846111db565b905092915050565b600067ffffffffffffffff82169050919050565b61125381611236565b811461125e57600080fd5b50565b6000813590506112708161124a565b92915050565b60006020828403121561128c5761128b610ee3565b5b600061129a84828501611261565b91505092915050565b6112ac816110d4565b82525050565b60006040820190506112c760008301856112a3565b81810360208301526112d981846111db565b90509392505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061130d826112e2565b9050919050565b61131d81611302565b82525050565b60006020820190506113386000830184611314565b92915050565b600081519050919050565b600082825260208201905092915050565b60006113658261133e565b61136f8185611349565b935061137f8185602086016111b1565b61138881610ef7565b840191505092915050565b61139c81611302565b82525050565b600082825260208201905092915050565b60006113be82611195565b6113c881856113a2565b93506113d88185602086016111b1565b6113e181610ef7565b840191505092915050565b6113f5816110d4565b82525050565b60008115159050919050565b611410816113fb565b82525050565b600060a0830160008301518482036000860152611433828261135a565b91505060208301516114486020860182611393565b506040830151848203604086015261146082826113b3565b915050606083015161147560608601826113ec565b5060808301516114886080860182611407565b508091505092915050565b600060208201905081810360008301526114ad8184611416565b905092915050565b6000602082840312156114cb576114ca610ee3565b5b600082013567ffffffffffffffff8111156114e9576114e8610ee8565b5b6114f5848285016110a6565b91505092915050565b611507816113fb565b811461151257600080fd5b50565b600081359050611524816114fe565b92915050565b6000806040838503121561154157611540610ee3565b5b600061154f85828601611261565b925050602061156085828601611515565b9150509250929050565b6000602082840312156115805761157f610ee3565b5b600082013567ffffffffffffffff81111561159e5761159d610ee8565b5b6115aa84828501611005565b91505092915050565b6115bc81611236565b82525050565b60006020820190506115d760008301846115b3565b92915050565b6115e681611302565b81146115f157600080fd5b50565b600081359050611603816115dd565b92915050565b60006020828403121561161f5761161e610ee3565b5b600061162d848285016115f4565b91505092915050565b600082825260208201905092915050565b7f696e76616c696420636f756e7400000000000000000000000000000000000000600082015250565b600061167d600d83611636565b915061168882611647565b602082019050919050565b600060208201905081810360008301526116ac81611670565b9050919050565b600081905092915050565b60006116c98261133e565b6116d381856116b3565b93506116e38185602086016111b1565b80840191505092915050565b60006116fb82846116be565b915081905092915050565b7f6578697374000000000000000000000000000000000000000000000000000000600082015250565b600061173c600583611636565b915061174782611706565b602082019050919050565b6000602082019050818103600083015261176b8161172f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806117b957607f821691505b6020821081036117cc576117cb611772565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026118347fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826117f7565b61183e86836117f7565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061188561188061187b84611856565b611860565b611856565b9050919050565b6000819050919050565b61189f8361186a565b6118b36118ab8261188c565b848454611804565b825550505050565b600090565b6118c86118bb565b6118d3818484611896565b505050565b5b818110156118f7576118ec6000826118c0565b6001810190506118d9565b5050565b601f82111561193c5761190d816117d2565b611916846117e7565b81016020851015611925578190505b611939611931856117e7565b8301826118d8565b50505b505050565b600082821c905092915050565b600061195f60001984600802611941565b1980831691505092915050565b6000611978838361194e565b9150826002028217905092915050565b6119918261133e565b67ffffffffffffffff8111156119aa576119a9610f08565b5b6119b482546117a1565b6119bf8282856118fb565b600060209050601f8311600181146119f257600084156119e0578287015190505b6119ea858261196c565b865550611a52565b601f198416611a00866117d2565b60005b82811015611a2857848901518255600182019150602085019450602081019050611a03565b86831015611a455784890151611a41601f89168261194e565b8355505b6001600288020188555050505b505050505050565b60008190508160005260206000209050919050565b601f821115611ab057611a8181611a5a565b611a8a846117e7565b81016020851015611a99578190505b611aad611aa5856117e7565b8301826118d8565b50505b505050565b611abe82611195565b67ffffffffffffffff811115611ad757611ad6610f08565b5b611ae182546117a1565b611aec828285611a6f565b600060209050601f831160018114611b1f5760008415611b0d578287015190505b611b17858261196c565b865550611b7f565b601f198416611b2d86611a5a565b60005b82811015611b5557848901518255600182019150602085019450602081019050611b30565b86831015611b725784890151611b6e601f89168261194e565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611bc182611236565b9150611bcc83611236565b9250828203905067ffffffffffffffff811115611bec57611beb611b87565b5b92915050565b6000611bfd8261133e565b611c078185611636565b9350611c178185602086016111b1565b611c2081610ef7565b840191505092915050565b6000604082019050611c4060008301856115b3565b8181036020830152611c528184611bf2565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f696e76616c696400000000000000000000000000000000000000000000000000600082015250565b6000611cc0600783611636565b9150611ccb82611c8a565b602082019050919050565b60006020820190508181036000830152611cef81611cb3565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611d52602683611636565b9150611d5d82611cf6565b604082019050919050565b60006020820190508181036000830152611d8181611d45565b9050919050565b7f6e6f206163746976650000000000000000000000000000000000000000000000600082015250565b6000611dbe600983611636565b9150611dc982611d88565b602082019050919050565b60006020820190508181036000830152611ded81611db1565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611e2a602083611636565b9150611e3582611df4565b602082019050919050565b60006020820190508181036000830152611e5981611e1d565b905091905056fea26469706673582212208111f7c475688fc18e37353b107b010576193188f8fd19b8734b6f46796b3eb164736f6c63430008180033",
}

// ModelABI is the input ABI used to generate the binding from.
// Deprecated: Use ModelMetaData.ABI instead.
var ModelABI = ModelMetaData.ABI

// ModelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ModelMetaData.Bin instead.
var ModelBin = ModelMetaData.Bin

// DeployModel deploys a new Ethereum contract, binding an instance of Model to it.
func DeployModel(auth *bind.TransactOpts, backend bind.ContractBackend, _b common.Address) (common.Address, *types.Transaction, *Model, error) {
	parsed, err := ModelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ModelBin), backend, _b)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Model{ModelCaller: ModelCaller{contract: contract}, ModelTransactor: ModelTransactor{contract: contract}, ModelFilterer: ModelFilterer{contract: contract}}, nil
}

// Model is an auto generated Go binding around an Ethereum contract.
type Model struct {
	ModelCaller     // Read-only binding to the contract
	ModelTransactor // Write-only binding to the contract
	ModelFilterer   // Log filterer for contract events
}

// ModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModelSession struct {
	Contract     *Model            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModelCallerSession struct {
	Contract *ModelCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModelTransactorSession struct {
	Contract     *ModelTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModelRaw struct {
	Contract *Model // Generic contract binding to access the raw methods on
}

// ModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModelCallerRaw struct {
	Contract *ModelCaller // Generic read-only contract binding to access the raw methods on
}

// ModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModelTransactorRaw struct {
	Contract *ModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModel creates a new instance of Model, bound to a specific deployed contract.
func NewModel(address common.Address, backend bind.ContractBackend) (*Model, error) {
	contract, err := bindModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Model{ModelCaller: ModelCaller{contract: contract}, ModelTransactor: ModelTransactor{contract: contract}, ModelFilterer: ModelFilterer{contract: contract}}, nil
}

// NewModelCaller creates a new read-only instance of Model, bound to a specific deployed contract.
func NewModelCaller(address common.Address, caller bind.ContractCaller) (*ModelCaller, error) {
	contract, err := bindModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModelCaller{contract: contract}, nil
}

// NewModelTransactor creates a new write-only instance of Model, bound to a specific deployed contract.
func NewModelTransactor(address common.Address, transactor bind.ContractTransactor) (*ModelTransactor, error) {
	contract, err := bindModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModelTransactor{contract: contract}, nil
}

// NewModelFilterer creates a new log filterer instance of Model, bound to a specific deployed contract.
func NewModelFilterer(address common.Address, filterer bind.ContractFilterer) (*ModelFilterer, error) {
	contract, err := bindModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModelFilterer{contract: contract}, nil
}

// bindModel binds a generic wrapper to an already deployed contract.
func bindModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Model *ModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Model.Contract.ModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Model *ModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Model.Contract.ModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Model *ModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Model.Contract.ModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Model *ModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Model.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Model *ModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Model.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Model *ModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Model.Contract.contract.Transact(opts, method, params...)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Model *ModelCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Model *ModelSession) Bank() (common.Address, error) {
	return _Model.Contract.Bank(&_Model.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Model *ModelCallerSession) Bank() (common.Address, error) {
	return _Model.Contract.Bank(&_Model.CallOpts)
}

// C4vk is a free data retrieval call binding the contract method 0xa35176d0.
//
// Solidity: function c4vk() view returns(bytes)
func (_Model *ModelCaller) C4vk(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "c4vk")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// C4vk is a free data retrieval call binding the contract method 0xa35176d0.
//
// Solidity: function c4vk() view returns(bytes)
func (_Model *ModelSession) C4vk() ([]byte, error) {
	return _Model.Contract.C4vk(&_Model.CallOpts)
}

// C4vk is a free data retrieval call binding the contract method 0xa35176d0.
//
// Solidity: function c4vk() view returns(bytes)
func (_Model *ModelCallerSession) C4vk() ([]byte, error) {
	return _Model.Contract.C4vk(&_Model.CallOpts)
}

// Check is a free data retrieval call binding the contract method 0xf804843d.
//
// Solidity: function check(uint64 _mi) view returns()
func (_Model *ModelCaller) Check(opts *bind.CallOpts, _mi uint64) error {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "check", _mi)

	if err != nil {
		return err
	}

	return err

}

// Check is a free data retrieval call binding the contract method 0xf804843d.
//
// Solidity: function check(uint64 _mi) view returns()
func (_Model *ModelSession) Check(_mi uint64) error {
	return _Model.Contract.Check(&_Model.CallOpts, _mi)
}

// Check is a free data retrieval call binding the contract method 0xf804843d.
//
// Solidity: function check(uint64 _mi) view returns()
func (_Model *ModelCallerSession) Check(_mi uint64) error {
	return _Model.Contract.Check(&_Model.CallOpts, _mi)
}

// GetIndex is a free data retrieval call binding the contract method 0xea1bbe35.
//
// Solidity: function getIndex(string _mn) view returns(uint64)
func (_Model *ModelCaller) GetIndex(opts *bind.CallOpts, _mn string) (uint64, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "getIndex", _mn)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetIndex is a free data retrieval call binding the contract method 0xea1bbe35.
//
// Solidity: function getIndex(string _mn) view returns(uint64)
func (_Model *ModelSession) GetIndex(_mn string) (uint64, error) {
	return _Model.Contract.GetIndex(&_Model.CallOpts, _mn)
}

// GetIndex is a free data retrieval call binding the contract method 0xea1bbe35.
//
// Solidity: function getIndex(string _mn) view returns(uint64)
func (_Model *ModelCallerSession) GetIndex(_mn string) (uint64, error) {
	return _Model.Contract.GetIndex(&_Model.CallOpts, _mn)
}

// GetModel is a free data retrieval call binding the contract method 0x9f5591ab.
//
// Solidity: function getModel(uint64 _mi) view returns((string,address,bytes,bytes32,bool) _info)
func (_Model *ModelCaller) GetModel(opts *bind.CallOpts, _mi uint64) (IModelInfo, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "getModel", _mi)

	if err != nil {
		return *new(IModelInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IModelInfo)).(*IModelInfo)

	return out0, err

}

// GetModel is a free data retrieval call binding the contract method 0x9f5591ab.
//
// Solidity: function getModel(uint64 _mi) view returns((string,address,bytes,bytes32,bool) _info)
func (_Model *ModelSession) GetModel(_mi uint64) (IModelInfo, error) {
	return _Model.Contract.GetModel(&_Model.CallOpts, _mi)
}

// GetModel is a free data retrieval call binding the contract method 0x9f5591ab.
//
// Solidity: function getModel(uint64 _mi) view returns((string,address,bytes,bytes32,bool) _info)
func (_Model *ModelCallerSession) GetModel(_mi uint64) (IModelInfo, error) {
	return _Model.Contract.GetModel(&_Model.CallOpts, _mi)
}

// GetOwner is a free data retrieval call binding the contract method 0xd370a37d.
//
// Solidity: function getOwner(uint64 _mi) view returns(address)
func (_Model *ModelCaller) GetOwner(opts *bind.CallOpts, _mi uint64) (common.Address, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "getOwner", _mi)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0xd370a37d.
//
// Solidity: function getOwner(uint64 _mi) view returns(address)
func (_Model *ModelSession) GetOwner(_mi uint64) (common.Address, error) {
	return _Model.Contract.GetOwner(&_Model.CallOpts, _mi)
}

// GetOwner is a free data retrieval call binding the contract method 0xd370a37d.
//
// Solidity: function getOwner(uint64 _mi) view returns(address)
func (_Model *ModelCallerSession) GetOwner(_mi uint64) (common.Address, error) {
	return _Model.Contract.GetOwner(&_Model.CallOpts, _mi)
}

// GetRoot is a free data retrieval call binding the contract method 0x6c37609a.
//
// Solidity: function getRoot(uint64 _mi) view returns(bytes32, bytes)
func (_Model *ModelCaller) GetRoot(opts *bind.CallOpts, _mi uint64) ([32]byte, []byte, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "getRoot", _mi)

	if err != nil {
		return *new([32]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetRoot is a free data retrieval call binding the contract method 0x6c37609a.
//
// Solidity: function getRoot(uint64 _mi) view returns(bytes32, bytes)
func (_Model *ModelSession) GetRoot(_mi uint64) ([32]byte, []byte, error) {
	return _Model.Contract.GetRoot(&_Model.CallOpts, _mi)
}

// GetRoot is a free data retrieval call binding the contract method 0x6c37609a.
//
// Solidity: function getRoot(uint64 _mi) view returns(bytes32, bytes)
func (_Model *ModelCallerSession) GetRoot(_mi uint64) ([32]byte, []byte, error) {
	return _Model.Contract.GetRoot(&_Model.CallOpts, _mi)
}

// GetVK is a free data retrieval call binding the contract method 0x354bb9ba.
//
// Solidity: function getVK() view returns(bytes)
func (_Model *ModelCaller) GetVK(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "getVK")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetVK is a free data retrieval call binding the contract method 0x354bb9ba.
//
// Solidity: function getVK() view returns(bytes)
func (_Model *ModelSession) GetVK() ([]byte, error) {
	return _Model.Contract.GetVK(&_Model.CallOpts)
}

// GetVK is a free data retrieval call binding the contract method 0x354bb9ba.
//
// Solidity: function getVK() view returns(bytes)
func (_Model *ModelCallerSession) GetVK() ([]byte, error) {
	return _Model.Contract.GetVK(&_Model.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Model *ModelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Model *ModelSession) Owner() (common.Address, error) {
	return _Model.Contract.Owner(&_Model.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Model *ModelCallerSession) Owner() (common.Address, error) {
	return _Model.Contract.Owner(&_Model.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0xd48dca2d.
//
// Solidity: function activate(uint64 _mi, bool _s) returns()
func (_Model *ModelTransactor) Activate(opts *bind.TransactOpts, _mi uint64, _s bool) (*types.Transaction, error) {
	return _Model.contract.Transact(opts, "activate", _mi, _s)
}

// Activate is a paid mutator transaction binding the contract method 0xd48dca2d.
//
// Solidity: function activate(uint64 _mi, bool _s) returns()
func (_Model *ModelSession) Activate(_mi uint64, _s bool) (*types.Transaction, error) {
	return _Model.Contract.Activate(&_Model.TransactOpts, _mi, _s)
}

// Activate is a paid mutator transaction binding the contract method 0xd48dca2d.
//
// Solidity: function activate(uint64 _mi, bool _s) returns()
func (_Model *ModelTransactorSession) Activate(_mi uint64, _s bool) (*types.Transaction, error) {
	return _Model.Contract.Activate(&_Model.TransactOpts, _mi, _s)
}

// Add is a paid mutator transaction binding the contract method 0x120a4734.
//
// Solidity: function add(string _mn, bytes _rt, bytes32 _cnt) returns()
func (_Model *ModelTransactor) Add(opts *bind.TransactOpts, _mn string, _rt []byte, _cnt [32]byte) (*types.Transaction, error) {
	return _Model.contract.Transact(opts, "add", _mn, _rt, _cnt)
}

// Add is a paid mutator transaction binding the contract method 0x120a4734.
//
// Solidity: function add(string _mn, bytes _rt, bytes32 _cnt) returns()
func (_Model *ModelSession) Add(_mn string, _rt []byte, _cnt [32]byte) (*types.Transaction, error) {
	return _Model.Contract.Add(&_Model.TransactOpts, _mn, _rt, _cnt)
}

// Add is a paid mutator transaction binding the contract method 0x120a4734.
//
// Solidity: function add(string _mn, bytes _rt, bytes32 _cnt) returns()
func (_Model *ModelTransactorSession) Add(_mn string, _rt []byte, _cnt [32]byte) (*types.Transaction, error) {
	return _Model.Contract.Add(&_Model.TransactOpts, _mn, _rt, _cnt)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Model *ModelTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Model.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Model *ModelSession) RenounceOwnership() (*types.Transaction, error) {
	return _Model.Contract.RenounceOwnership(&_Model.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Model *ModelTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Model.Contract.RenounceOwnership(&_Model.TransactOpts)
}

// SetVK is a paid mutator transaction binding the contract method 0xd0b9178b.
//
// Solidity: function setVK(bytes _vk) returns()
func (_Model *ModelTransactor) SetVK(opts *bind.TransactOpts, _vk []byte) (*types.Transaction, error) {
	return _Model.contract.Transact(opts, "setVK", _vk)
}

// SetVK is a paid mutator transaction binding the contract method 0xd0b9178b.
//
// Solidity: function setVK(bytes _vk) returns()
func (_Model *ModelSession) SetVK(_vk []byte) (*types.Transaction, error) {
	return _Model.Contract.SetVK(&_Model.TransactOpts, _vk)
}

// SetVK is a paid mutator transaction binding the contract method 0xd0b9178b.
//
// Solidity: function setVK(bytes _vk) returns()
func (_Model *ModelTransactorSession) SetVK(_vk []byte) (*types.Transaction, error) {
	return _Model.Contract.SetVK(&_Model.TransactOpts, _vk)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Model *ModelTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Model.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Model *ModelSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Model.Contract.TransferOwnership(&_Model.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Model *ModelTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Model.Contract.TransferOwnership(&_Model.TransactOpts, newOwner)
}

// ModelAddModelIterator is returned from FilterAddModel and is used to iterate over the raw logs and unpacked data for AddModel events raised by the Model contract.
type ModelAddModelIterator struct {
	Event *ModelAddModel // Event containing the contract specifics and raw log

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
func (it *ModelAddModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelAddModel)
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
		it.Event = new(ModelAddModel)
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
func (it *ModelAddModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelAddModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelAddModel represents a AddModel event raised by the Model contract.
type ModelAddModel struct {
	A   common.Address
	Mi  uint64
	Mn  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddModel is a free log retrieval operation binding the contract event 0x1e0210396cbbc2a1ccae7d7b17e27a3fea0e56962c68454eda6fa9682bc7fe47.
//
// Solidity: event AddModel(address indexed _a, uint64 _mi, string _mn)
func (_Model *ModelFilterer) FilterAddModel(opts *bind.FilterOpts, _a []common.Address) (*ModelAddModelIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Model.contract.FilterLogs(opts, "AddModel", _aRule)
	if err != nil {
		return nil, err
	}
	return &ModelAddModelIterator{contract: _Model.contract, event: "AddModel", logs: logs, sub: sub}, nil
}

// WatchAddModel is a free log subscription operation binding the contract event 0x1e0210396cbbc2a1ccae7d7b17e27a3fea0e56962c68454eda6fa9682bc7fe47.
//
// Solidity: event AddModel(address indexed _a, uint64 _mi, string _mn)
func (_Model *ModelFilterer) WatchAddModel(opts *bind.WatchOpts, sink chan<- *ModelAddModel, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Model.contract.WatchLogs(opts, "AddModel", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelAddModel)
				if err := _Model.contract.UnpackLog(event, "AddModel", log); err != nil {
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

// ParseAddModel is a log parse operation binding the contract event 0x1e0210396cbbc2a1ccae7d7b17e27a3fea0e56962c68454eda6fa9682bc7fe47.
//
// Solidity: event AddModel(address indexed _a, uint64 _mi, string _mn)
func (_Model *ModelFilterer) ParseAddModel(log types.Log) (*ModelAddModel, error) {
	event := new(ModelAddModel)
	if err := _Model.contract.UnpackLog(event, "AddModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Model contract.
type ModelOwnershipTransferredIterator struct {
	Event *ModelOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ModelOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelOwnershipTransferred)
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
		it.Event = new(ModelOwnershipTransferred)
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
func (it *ModelOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelOwnershipTransferred represents a OwnershipTransferred event raised by the Model contract.
type ModelOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Model *ModelFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ModelOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Model.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ModelOwnershipTransferredIterator{contract: _Model.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Model *ModelFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ModelOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Model.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelOwnershipTransferred)
				if err := _Model.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Model *ModelFilterer) ParseOwnershipTransferred(log types.Log) (*ModelOwnershipTransferred, error) {
	event := new(ModelOwnershipTransferred)
	if err := _Model.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
