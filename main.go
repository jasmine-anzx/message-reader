package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"encoding/base64"

	originationv2 "github.com/anzx/apis-go/origination/service/account/v2"
	onboardingv1 "github.com/anzx/apis-go/ribbon/service/onboarding/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	//readCloudPubSubMessage()
	convertJsonToPayload()
}

func readCloudPubSubMessage() {
	content, err := os.ReadFile("message.txt")
	if err != nil {
		fmt.Println(err)
	}
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(content)))
	if _, err := base64.StdEncoding.Decode(dst, content); err != nil {
		fmt.Println(err)
	}
	customerState := &onboardingv1.CustomerOnboardingState{}
	if err := proto.Unmarshal(dst, customerState); err != nil {
		fmt.Println(err)
	}
	//taxResidencies := customerState.GetTaxResidencies()
	//for i, t := range taxResidencies {
	//	t.RegionCode = countries[i].regionCode
	//	t.RegionName = countries[i].regionName
	//}
	//cs, err := proto.Marshal(customerState)
	//result := b64.StdEncoding.EncodeToString(cs)
	//fmt.Println(result)
	//Converting to json
	customerStateJSON, err := json.MarshalIndent(customerState, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Message Data:\n %s\n", string(customerStateJSON))
}

func convertJsonToPayload() {
	content, err := os.ReadFile("payload.json")
	if err != nil {
		fmt.Println(err)
	}
	packageOrigData := &originationv2.PackageOriginationData{}
	err = proto.Unmarshal(content, packageOrigData)
	if err != nil {
		fmt.Println(err)
	}

	msgData, _ := protojson.Marshal(packageOrigData)
	fmt.Println(msgData)
}
