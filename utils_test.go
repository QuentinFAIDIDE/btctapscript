package btctapscript

import (
	"reflect"
	"testing"
)

func TestGetWitnessForInput01(t *testing.T) {
	testTx := `
	{
		"txid": "700bf4b1d18f40c674a11d5a2c3a9c09666d24a8cfe47746b7e8148781b1f9e8",
		"hash": "e2cd2f906324066a10e5ed36df899a92ce9e299eac61d9c96119f4bc19196981",
		"version": 2,
		"size": 182,
		"vsize": 131,
		"weight": 524,
		"locktime": 0,
		"vin": [
		  {
			"txid": "2c3fd8ae40f4011712b2cd0eaa8431c82e15823fabdb304a4c3205a06fe5edf1",
			"vout": 0,
			"scriptSig": {
			  "asm": "",
			  "hex": ""
			},
			"txinwitness": [
			  "6570cfa568a062d2653b7cbdfb8bf2bdfd537fda7284a9ff00c455d3df9f140a94f9894d720d09a6e70c13329fc71c86245f8e66cef7dcce425de28d3ca1a8e6"
			],
			"prevout": {
			  "generated": false,
			  "height": 841476,
			  "value": 0.00005045,
			  "scriptPubKey": {
				"asm": "1 d96da555997d201c0ba18ebcb73e312571fb30f455907d55a153e7104f94cd29",
				"desc": "rawtr(d96da555997d201c0ba18ebcb73e312571fb30f455907d55a153e7104f94cd29)#as43zmld",
				"hex": "5120d96da555997d201c0ba18ebcb73e312571fb30f455907d55a153e7104f94cd29",
				"address": "bc1pm9k624ve05spczap367tw033y4clkv852kg864dp20n3qnu5e55sdkfwz9",
				"type": "witness_v1_taproot"
			  }
			},
			"sequence": 4294967295
		  }
		],
		"vout": [
		  {
			"value": 0.00000546,
			"n": 0,
			"scriptPubKey": {
			  "asm": "1 72a2e3409ac29a01dc70bb851dba2b4bc5e4c06a2db7a0eb90088ca65b896742",
			  "desc": "rawtr(72a2e3409ac29a01dc70bb851dba2b4bc5e4c06a2db7a0eb90088ca65b896742)#56ga255w",
			  "hex": "512072a2e3409ac29a01dc70bb851dba2b4bc5e4c06a2db7a0eb90088ca65b896742",
			  "address": "bc1pw23wxsy6c2dqrhrshwz3mw3tf0z7fsr29km6p6uspzx2vkufvapqhpf8z3",
			  "type": "witness_v1_taproot"
			}
		  },
		  {
			"value": 0.00000000,
			"n": 1,
			"scriptPubKey": {
			  "asm": "OP_RETURN 13 14b0a733141d1600",
			  "desc": "raw(6a5d0814b0a733141d1600)#p5vtrnze",
			  "hex": "6a5d0814b0a733141d1600",
			  "type": "nulldata"
			}
		  }
		],
		"fee": 0.00004499,
		"hex": "02000000000101f1ede56fa005324c4a30dbab3f82152ec83184aa0ecdb2121701f440aed83f2c0000000000ffffffff02220200000000000022512072a2e3409ac29a01dc70bb851dba2b4bc5e4c06a2db7a0eb90088ca65b89674200000000000000000b6a5d0814b0a733141d160001406570cfa568a062d2653b7cbdfb8bf2bdfd537fda7284a9ff00c455d3df9f140a94f9894d720d09a6e70c13329fc71c86245f8e66cef7dcce425de28d3ca1a8e600000000",
		"blockhash": "00000000000000000002e3e27ba27fe3d96803ee2e7615c7e6f3aea23fb7ca7e",
		"confirmations": 561,
		"time": 1714462823,
		"blocktime": 1714462823
	  }	  
`

	witness, err := getWitnessForInput(testTx, 0)
	if err != nil {
		t.Fatal(err)
	}

	expectedWitnessStr := []string{
		"6570cfa568a062d2653b7cbdfb8bf2bdfd537fda7284a9ff00c455d3df9f140a94f9894d720d09a6e70c13329fc71c86245f8e66cef7dcce425de28d3ca1a8e6",
	}

	if !reflect.DeepEqual(witness, hexaStringSliceToBytes(expectedWitnessStr)) {
		t.Fatal("witness retrieves has unexpected value")
	}
}

func TestGetWitnessForInput02(t *testing.T) {
	testTx := `
	{
		"txid": "55750249356628f5b6367cec82b092fb69b1b80a1f55a4cdbf9041d92a101a30",
		"hash": "69515398243bfc19521903a6711ee86f23e7882b88a49393ff748ddf4dcf36b1",
		"version": 2,
		"size": 355,
		"vsize": 174,
		"weight": 694,
		"locktime": 0,
		"vin": [
		  {
			"txid": "632a62b8177bfa55d7eb9e757a2037f128bb7bed668fc0efeb08db24b3ae9dc7",
			"vout": 0,
			"scriptSig": {
			  "asm": "",
			  "hex": ""
			},
			"txinwitness": [
			  "521d4fa4df0bf2f0933ecdbb33e11fa288346e89ef82e568cb28b0d8a528fff135c6bdcba29a8321c6566ed446be6d75226a1a415c0b0b19701660c3f90c0829",
			  "2057109034c99ed544e7bc5285920eaccc05e67556403a7da82a633c8e6cd623f0ac0063036f7264010118746578742f706c61696e3b636861727365743d7574662d3800457b2270223a226272632d3230222c226f70223a226465706c6f79222c227469636b223a2273636467222c226c696d223a2231222c226d6178223a223231303030303030227d68",
			  "c1e8de26eb76709ec2eceef1512b6d1ca4c99c24e13509bd250cbb62d8ed422fcc"
			],
			"prevout": {
			  "generated": false,
			  "height": 841521,
			  "value": 0.00005000,
			  "scriptPubKey": {
				"asm": "1 17133286836c00bf5d24f8499441300534a52d25acc5597860f09dd7ac69e35d",
				"desc": "rawtr(17133286836c00bf5d24f8499441300534a52d25acc5597860f09dd7ac69e35d)#2sk3dtkv",
				"hex": "512017133286836c00bf5d24f8499441300534a52d25acc5597860f09dd7ac69e35d",
				"address": "bc1pzufn9p5rdsqt7hfylpyegsfsq5622tf94nz4j7rq7zwa0trfudwsuz9tpv",
				"type": "witness_v1_taproot"
			  }
			},
			"sequence": 4294967293
		  }
		],
		"vout": [
		  {
			"value": 0.00000546,
			"n": 0,
			"scriptPubKey": {
			  "asm": "0 8c4cab55ab911fad95542a1cd1739fd09df3029a",
			  "desc": "addr(bc1q33x2k4dtjy06m9259gwdzuul6zwlxq56cm479r)#dg8ye5ta",
			  "hex": "00148c4cab55ab911fad95542a1cd1739fd09df3029a",
			  "address": "bc1q33x2k4dtjy06m9259gwdzuul6zwlxq56cm479r",
			  "type": "witness_v0_keyhash"
			}
		  },
		  {
			"value": 0.00001496,
			"n": 1,
			"scriptPubKey": {
			  "asm": "0 e2fde28f96e76d7618f0e80b6c9a3db04292a509",
			  "desc": "addr(bc1qut779rukuakhvx8saq9kex3akppf9fgf6jp2g5)#93vq3yze",
			  "hex": "0014e2fde28f96e76d7618f0e80b6c9a3db04292a509",
			  "address": "bc1qut779rukuakhvx8saq9kex3akppf9fgf6jp2g5",
			  "type": "witness_v0_keyhash"
			}
		  }
		],
		"fee": 0.00002958,
		"hex": "02000000000101c79daeb324db08ebefc08f66ed7bbb28f137207a759eebd755fa7b17b8622a630000000000fdffffff0222020000000000001600148c4cab55ab911fad95542a1cd1739fd09df3029ad805000000000000160014e2fde28f96e76d7618f0e80b6c9a3db04292a5090340521d4fa4df0bf2f0933ecdbb33e11fa288346e89ef82e568cb28b0d8a528fff135c6bdcba29a8321c6566ed446be6d75226a1a415c0b0b19701660c3f90c08298b2057109034c99ed544e7bc5285920eaccc05e67556403a7da82a633c8e6cd623f0ac0063036f7264010118746578742f706c61696e3b636861727365743d7574662d3800457b2270223a226272632d3230222c226f70223a226465706c6f79222c227469636b223a2273636467222c226c696d223a2231222c226d6178223a223231303030303030227d6821c1e8de26eb76709ec2eceef1512b6d1ca4c99c24e13509bd250cbb62d8ed422fcc00000000",
		"blockhash": "000000000000000000031f43f7052bfa753b78f0c33d396055aa03c5e384e1df",
		"confirmations": 517,
		"time": 1714486187,
		"blocktime": 1714486187
	  }  
`

	witness, err := getWitnessForInput(testTx, 0)
	if err != nil {
		t.Fatal(err)
	}

	expectedWitnessStr := []string{
		"521d4fa4df0bf2f0933ecdbb33e11fa288346e89ef82e568cb28b0d8a528fff135c6bdcba29a8321c6566ed446be6d75226a1a415c0b0b19701660c3f90c0829",
		"2057109034c99ed544e7bc5285920eaccc05e67556403a7da82a633c8e6cd623f0ac0063036f7264010118746578742f706c61696e3b636861727365743d7574662d3800457b2270223a226272632d3230222c226f70223a226465706c6f79222c227469636b223a2273636467222c226c696d223a2231222c226d6178223a223231303030303030227d68",
		"c1e8de26eb76709ec2eceef1512b6d1ca4c99c24e13509bd250cbb62d8ed422fcc",
	}

	if !reflect.DeepEqual(witness, hexaStringSliceToBytes(expectedWitnessStr)) {
		t.Log(witness)
		t.Log(hexaStringSliceToBytes(expectedWitnessStr))
		t.Fatal("witness retrieves has unexpected value")
	}
}
