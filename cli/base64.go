package cli

import (
	"encoding/base64"
	"fmt"

	"github.com/nikhilsbhat/neuron/cli/ui"
)

func base64Encode(s string) error {
	if len(s) != 0 {
		encode := base64.StdEncoding.EncodeToString([]byte(s))
		cm.NeuronSaysItsInfo("")
		fmt.Println("The encoded message is: " + ui.Info(encode) + "\n")

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
		cm.NeuronSaysItsInfo("")
		fmt.Println("The decoded message is: " + ui.Info(string(decode)) + "\n")

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
