package param

type paramInfo struct {
	hash string
	size int64
}

var paramMap map[string]paramInfo

func init() {
	paramMap = make(map[string]paramInfo)
	// kzg
	paramMap["kzgsrs-bls12_377-kzg"] = paramInfo{
		size: 3221274340,
		hash: "ecabcd204af80676d39594bb2a1fa7b649f7115a13ac143f5b378f5b3d6d49e2",
	}
	paramMap["kzgsrs-bls12_377-kzg-vk"] = paramInfo{
		size: 48624,
		hash: "3f042498c2fe3cb2bc15816fccd64868c638a3d0a90a81eb05d515b727ddfc37",
	}

	// bn254
	paramMap["kzgsrs-bn254-canonical"] = paramInfo{
		size: 2147517956,
		hash: "4c73afc8e6c59579dd18bf68bf0312ebc975f8a7067b852f6fb616d4acc1d65f",
	}
	paramMap["kzgsrs-bn254-lagrange-33554432"] = paramInfo{
		size: 2147517764,
		hash: "066dd595dd4f89a00e0d67a49b71843b7c5feed7c13a50eee4ad581b9119c5ba",
	}

	// bw6_761
	paramMap["kzgsrs-bw6_761-canonical"] = paramInfo{
		size: 402799492,
		hash: "0476886cd15dedc190229f01003b5cd0a23ec32ed2293e135721a813a80b69e6",
	}
	paramMap["kzgsrs-bw6_761-lagrange-128"] = paramInfo{
		size: 170308,
		hash: "04ca05ea12f3a8cf362f9a2120a4f2a6470b8cd06f6c3e37bc0dd849457799e4",
	}
	paramMap["kzgsrs-bw6_761-lagrange-256"] = paramInfo{
		size: 194884,
		hash: "9d2442c2dd73bcf93278e963a17e2250690cf2117423ec79a29c7db9c4c21f06",
	}
	paramMap["kzgsrs-bw6_761-lagrange-512"] = paramInfo{
		size: 244036,
		hash: "6745164cac368d9b0cc1a0b740eee56efd490b0840165acde6555626191718a8",
	}
	paramMap["kzgsrs-bw6_761-lagrange-1024"] = paramInfo{
		size: 342340,
		hash: "7a48faa2df260322dc9b06da3d89f2810a5169f1d10bdeba193378f84f9765a1",
	}
	paramMap["kzgsrs-bw6_761-lagrange-2048"] = paramInfo{
		size: 538948,
		hash: "d332086d0e5d941d057df4576eaeeeb772205f2e69cba3a0caf16718ffa8f9de",
	}
	paramMap["kzgsrs-bw6_761-lagrange-4096"] = paramInfo{
		size: 932164,
		hash: "547a9fb1b3df2afbf9647a08fe8b172d6b58415b197bda2b95bda40539ff73a1",
	}
	paramMap["kzgsrs-bw6_761-lagrange-8192"] = paramInfo{
		size: 1718596,
		hash: "cc0ffbf2832d367588fd16b306c07e15c7443616dd2445d07c7caf98b5a32c99",
	}
	paramMap["kzgsrs-bw6_761-lagrange-16384"] = paramInfo{
		size: 3291460,
		hash: "671fd98367b4b1ffad20b7274045e8167234590ca5bb26db40cc97d086806049",
	}
	paramMap["kzgsrs-bw6_761-lagrange-32768"] = paramInfo{
		size: 6437188,
		hash: "19f9e4b7cfd7fb73f2b0e1db1a58bbf43bc41b4aa6a33412763c04b9514b193e",
	}
	paramMap["kzgsrs-bw6_761-lagrange-65536"] = paramInfo{
		size: 12728644,
		hash: "4ef974e7bf0173d9e458f61c0ed709364813461ff32f13becf6e7c7bf4268e98",
	}
	paramMap["kzgsrs-bw6_761-lagrange-131072"] = paramInfo{
		size: 25311556,
		hash: "8ce79b8bf53946e8308b8c6fda6302af264394cfe725e0048487984798d11f3b",
	}
	paramMap["kzgsrs-bw6_761-lagrange-262144"] = paramInfo{
		size: 50477380,
		hash: "97024814fea0fc4df09015a908cac0ffc18025d244a2654f0291ff1527ddae05",
	}
}
