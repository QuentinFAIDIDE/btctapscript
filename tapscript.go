package btctapscript

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"slices"
	"strconv"
	"strings"

	"github.com/QuentinFAIDIDE/btctapscript/opcodes"
)

// WitnessAnnexPrefixByte is the prefix in the witness array of the
// annex that is to be removed in the parsing process.
const WitnessAnnexPrefixByte = 0x50

// RemoveWitnessAnnex removes the eventual annex from the slice of witnesses and return a new one.
func RemoveWitnessAnnex(witnesses [][]byte) ([][]byte, error) {
	// condition for validating tapscript: https://bips.xyz/342 (Bitcoin BIP 342)
	// If there are at least two witness elements, and the first byte of the last element is 0x50[4], this last element is called annex a[5] and is removed from the witness stack.
	if len(witnesses) < 2 {
		return nil, errors.New("trying to remove annex from a witness slice with less than 2 elements")
	}
	witnessWithoutAnnex := slices.Clone(witnesses)
	if witnessWithoutAnnex[len(witnessWithoutAnnex)-1][0] == WitnessAnnexPrefixByte {
		witnessWithoutAnnex = witnessWithoutAnnex[:len(witnessWithoutAnnex)-1]
	}
	return witnessWithoutAnnex, nil
}

// WitnessLeafVersionIs0xc0 tells if the witness (after annex is removed before passing it)
// has its leaf version being 0xc0.
func WitnessLeafVersionIs0xc0(witnessWithoutAnnex [][]byte) bool {
	// (and) The leaf version is 0xc0 (i.e. the first byte of the last witness element after removing the optional annex is 0xc0 or 0xc1), marking it as a tapscript spend.
	lastWitness := witnessWithoutAnnex[len(witnessWithoutAnnex)-1]
	startsWith0xc0 := lastWitness[0] == 0xc0
	startsWithOxc1 := lastWitness[0] == 0xc1
	return startsWith0xc0 || startsWithOxc1
}

// WitnessContainsTapscript, given the array of txinwitness from the bitcoin
// transaction input, will tell if it has a tapscript in the right version.
func WitnessContainsCompatibleTapscript(witnessWithAnnexe [][]byte) bool {
	if len(witnessWithAnnexe) < 2 {
		return false
	}
	witnessWithoutAnnexe, err := RemoveWitnessAnnex(witnessWithAnnexe)
	if err != nil {
		return false
	}
	if len(witnessWithAnnexe) < 2 {
		return false
	}
	if !WitnessLeafVersionIs0xc0(witnessWithoutAnnexe) {
		return false
	}
	return true
}

// GetWitnessesTaprootScriptAsm return the part of the witness entries
// that contains the script and errors out if input the witnesses are from does
// not contains tapscript with the right version.
func GetWitnessesTaprootScriptAsm(witnessWithAnnexe [][]byte) ([]byte, error) {
	if !WitnessContainsCompatibleTapscript(witnessWithAnnexe) {
		return nil, errors.New("incompatible witness array sent, is it from a pay to taproot input in the right version ?")
	}
	witnessWithoutAnnexe, err := RemoveWitnessAnnex(witnessWithAnnexe)
	if err != nil {
		return nil, err
	}

	return witnessWithoutAnnexe[len(witnessWithoutAnnexe)-2], nil
}

// DisassembleAsmTaprootScript tries to disassemble the taproot script
// from its ASM format into human readable script with OP codes names.
func DisassembleAsmTaprootScript(scriptAsm []byte) (string, error) {

	// inspired by: https://github.com/mempool/mempool/blob/dbd4d152ce831859375753fb4ca32ac0e5b1aff8/backend/src/api/transaction-utils.ts#L253

	// output disasembled script
	scriptStr := make([]string, 0)

	// the position in the tapscript
	i := 0

	for i < len(scriptAsm) {
		op := scriptAsm[i]
		if op >= 0x01 && op <= 0x4e {
			i++
			var nbBytesPayload int

			// handling of the pushdata/pushbyte op codes and payload size parsing
			if op == 0x4c {
				lenUint8 := uint8(scriptAsm[i])
				nbBytesPayload = int(lenUint8)
				scriptStr = append(scriptStr, "OP_PUSHDATA1")
				i++

			} else if op == 0x4d {
				lenUint16 := binary.LittleEndian.Uint16(scriptAsm[i : i+2])
				nbBytesPayload = int(lenUint16)
				scriptStr = append(scriptStr, "OP_PUSHDATA2")
				i += 2

			} else if op == 0x4e {
				lenUint32 := binary.LittleEndian.Uint32(scriptAsm[i : i+4])
				nbBytesPayload = int(lenUint32)
				scriptStr = append(scriptStr, "OP_PUSHDATA4")
				i += 4

			} else {
				nbBytesPayload = int(op)
				scriptStr = append(scriptStr, "OP_PUSHBYTES_"+strconv.Itoa(nbBytesPayload))

			}

			// handling of the payload reading
			if len(scriptAsm)-i < nbBytesPayload {
				return strings.Join(scriptStr, "\n"), errors.New("not enough data remaining for expected number of payload bytes of this opcode")
			}
			hexadecimalPayload := hex.EncodeToString(scriptAsm[i : i+nbBytesPayload])
			scriptStr[len(scriptStr)-1] = scriptStr[len(scriptStr)-1] + " " + hexadecimalPayload
			i += nbBytesPayload

			// handle other tapscript specific opcodes
		} else if op == 0x00 {
			scriptStr = append(scriptStr, "OP_0")
			i++

		} else if op == 0x4f {
			scriptStr = append(scriptStr, "OP_PUSHNUM_NEG1")
			i++

		} else if op == 0xb1 {
			scriptStr = append(scriptStr, "OP_CLTV")
			i++

		} else if op == 0xb2 {
			scriptStr = append(scriptStr, "OP_CSV")
			i++

		} else if op == 0xba {
			scriptStr = append(scriptStr, "OP_CHECKSIGADD")
			i++

			// This branch will resort to vanilla Bitcoin script
			// op codes and eventually modify some when neeeded.
		} else {
			opcodeInfo := opcodes.OpcodeArray[op]
			if opcodeInfo.Length != 1 {
				return strings.Join(scriptStr, "\n"), errors.New("Vanilla opcode had a payload (non 1 length) and we didn't expect that: " + opcodeInfo.Name + " with size " + strconv.Itoa(opcodeInfo.Length))
			}
			i = i + opcodeInfo.Length

			if op < 0xfd {
				// if the code is OP_[NUMBER], tapscript should read OP_PUSHNUM_[NUMBER]
				opWithoutPrefix := strings.Replace(opcodeInfo.Name, "OP_", "", 1)
				_, numberParseErr := strconv.Atoi(opWithoutPrefix)
				if numberParseErr == nil {
					newOpcodeNameFormat := "OP_PUSHNUM_" + opWithoutPrefix
					scriptStr = append(scriptStr, newOpcodeNameFormat)
				} else {
					scriptStr = append(scriptStr, opcodeInfo.Name)
				}
			} else {
				scriptStr = append(scriptStr, "OP_RETURN_"+strconv.Itoa(int(op)))
			}
		}
	}

	return strings.Join(scriptStr, "\n"), nil
}
