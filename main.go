package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"encoding/base64"
	b64 "encoding/base64"

	onboardingv1 "github.com/anzx/apis-go/ribbon/service/onboarding/v1"
	"google.golang.org/protobuf/proto"
)

type country struct {
	regionCode string
	regionName string
}

var countries = []*country{
	{
		regionCode: "EGY",
		regionName: "Egypt",
	},
	{
		regionCode: "SIN",
		regionName: "Singapore",
	},
	{
		regionCode: "AUS",
		regionName: "Australia",
	},
	{
		regionCode: "USA",
		regionName: "United States",
	},
	{
		regionCode: "FJI",
		regionName: "Fiji",
	},
	{
		regionCode: "VNA",
		regionName: "Vietnam",
	},
	{
		regionCode: "THA",
		regionName: "Thailand",
	},
	{
		regionCode: "IND",
		regionName: "India",
	},
}

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
