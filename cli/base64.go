package cli

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(s string) error {
	if len(s) != 0 {
		encode := base64.StdEncoding.EncodeToString([]byte(s))
		encodeMessage()
		cm.NeuronSaysItsInfo(encode)
		return nil
	}
	return fmt.Errorf("could not get the string to encode")
}

func base64Dcode(s string) error {
	if len(s) != 0 {
		decode, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			return err
		}
		decodeMessage()
		cm.NeuronSaysItsInfo(string(decode))
		return nil
	}
	return fmt.Errorf("could not get the string to encode")
}

func getStringOfMessage(g interface{}) string {
	switch g.(type) {
	case string:
		return g.(string)
	case error:
		return g.(error).Error()
	default:
		return "unknown messagetype"
	}
}

func encodeMessage() {
	cm.NeuronSaysItsInfo("The encoded message is:")
	printInfo()
}

func decodeMessage() {
	cm.NeuronSaysItsInfo("The decoded message is:")
	printInfo()
}

func printInfo() {
	cm.NeuronSaysItsInfo("")
	cm.NeuronSaysItsInfo("")
}
