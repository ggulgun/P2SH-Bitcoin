package main

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcutil"
	. "github.com/btcsuite/btcd/txscript"
)

func main() {
	
	pubKey1Hex := "04cb9c3c222c5f7a7d3b9bd152f363a0b6d54c9eb312c4d4f9af1e8551b6c421a6a4ab0e29105f24de20ff463c1c91fcf3bf662cdde4783d4799f787cb7c08869b"
	pubKey2Hex := "04ccc588420deeebea22a7e900cc8b68620d2212c374604e3487ca08f1ff3ae12bdc639514d0ec8612a2d3c519f084d9a00cbbe3b53d071e9b09e71e610b036aa2"
	pubKey3Hex := "04ab47ad1939edcb3db65f7fedea62bbf781c5410d3f22a7a3a56ffefb2238af8627363bdf2ed97c1f89784a1aecdb43384f11d2acc64443c7fc299cef0400421a"
	pubKey1, _ := hex.DecodeString(pubKey1Hex)
	pubKey2, _ := hex.DecodeString(pubKey2Hex)
	pubKey3, _ := hex.DecodeString(pubKey3Hex)
	sig1Hex := "3045022100dedc2621f9ab11cd008efa4a6734f64336cc5217368cc5a7bb3c74d025812267022059a300963f634adad7c4b01bcb3834ed697df17b0b81c708f92142b81f7ae12a01"
	sig2Hex := "3045022063a7dec8db21f38d73a922b9dd40a5107d59edb912339828ffce2c28a006f9b6022100a57a49beee175c913a5504c7d411d67736435b5cf21993ebc74e350576a98b1b01"
	sig1, _ := hex.DecodeString(sig1Hex)
	sig2, _ := hex.DecodeString(sig2Hex)


	payScript, _ := NewScriptBuilder().AddInt64(2).AddData(pubKey1).AddData(pubKey2).AddData(pubKey3).AddInt64(3).AddOp(OP_CHECKMULTISIG).Script()
	payScriptDisasm, _ := DisasmString(payScript)
	fmt.Printf("payScript hex: %x\npayScript asm: %s\n", payScript, payScriptDisasm)

	payScriptHash := btcutil.Hash160(payScript)
	fmt.Printf("payScript hash: %x\n", payScriptHash)

	outScript, _ := NewScriptBuilder().AddOp(OP_HASH160).AddData(payScriptHash).AddOp(OP_EQUAL).Script()
	outScriptDisasm, _ := DisasmString(outScript)
	fmt.Printf("outScript hex: %x\noutScript asm: %s\n", outScript, outScriptDisasm)

	inScript, _ := NewScriptBuilder().AddOp(OP_0).AddData(sig1).AddData(sig2).AddData(payScript).Script()
	inScriptDisasm, _ := DisasmString(inScript)
	fmt.Printf("inScript hex: %x\ninScript asm: %s\n", inScript, inScriptDisasm)
}