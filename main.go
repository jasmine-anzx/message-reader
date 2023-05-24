package main

import (
	"fmt"
	"log"
	"os"

	"encoding/base64"
	b64 "encoding/base64"
	"encoding/json"

	onboardingv1 "github.com/anzx/apis-go/ribbon/service/onboarding/v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	content, err := os.ReadFile("message.txt")
	if err != nil {
		fmt.Println(err)
	}
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(content)))
	if _, err := b64.StdEncoding.Decode(dst, content); err != nil {
		fmt.Println(err)
	}
	customerState := &onboardingv1.CustomerOnboardingState{}
	if err := proto.Unmarshal(dst, customerState); err != nil {
		fmt.Println(err)
	}
	//Converting to json
	customerStateJSON, err := json.MarshalIndent(customerState, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Message Data:\n %s\n", string(customerStateJSON))
}
