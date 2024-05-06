package btctapscript

import (
	"encoding/hex"
	"encoding/json"
	"errors"
)

// hexaStringSliceToBytes parse an array of string in hexadecimal format
// into bytes.
func hexaStringSliceToBytes(input []string) [][]byte {
	output := make([][]byte, len(input))
	for i := range input {
		hexVal, err := hex.DecodeString(input[i])
		if err != nil {
			panic("unable to convert str array to hex: " + err.Error())
		}
		output[i] = hexVal
	}
	return output
}

// getWitnessForInput will parse the raw transaction from the
// bitcoin node (with verbosity 2 or 3) and will extract the witness
// entries from the transaction input with the provided index.
func getWitnessForInput(txJson string, inputIndex int) ([][]byte, error) {

	var txRaw map[string]interface{}
	err := json.Unmarshal([]byte(txJson), &txRaw)
	if err != nil {
		return nil, err
	}

	vin, hasVin := txRaw["vin"].([]interface{})
	if !hasVin {
		return nil, errors.New("missing vin")
	}

	if inputIndex >= len(vin) || inputIndex < 0 {
		return nil, errors.New("input index out of range")
	}

	inputObject, inputObjectExists := vin[inputIndex].(map[string]interface{})
	if !inputObjectExists {
		return nil, errors.New("unable to acccess input in vin array as a JSON object")
	}

	witnessArray, hasWitnessArray := inputObject["txinwitness"].([]interface{})
	if !hasWitnessArray {
		return nil, errors.New("input is missing txinwitness field, is it a taproot input with a script ?")
	}

	output := make([][]byte, 0, len(witnessArray))
	for i := range witnessArray {
		if strVal, isAString := witnessArray[i].(string); !isAString {
			return nil, errors.New("witness array has items which are not strings")
		} else {
			byteVal, err := hex.DecodeString(strVal)
			if err != nil {
				return nil, errors.New("unable to decode hex witness: " + err.Error())
			}
			output = append(output, byteVal)
		}
	}

	return output, nil
}
