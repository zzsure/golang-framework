package util

import (
	"encoding/json"
	"fmt"
	"golang-framework/library/util/net"
)

func GenServerUUID() string {
	ip, mac := net.NewLAN().NetInfo()
	return fmt.Sprintf("%s-%s", ip, mac)
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
