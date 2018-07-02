package main

import (
	"encoding/xml"
	"encoding/json"
	"fmt"
)

type error struct {
	XMLName xml.Name `xml:"xml_name"`
	Code int `json:"code", xml:"message"`
	Message string `json:"message, xml:"message`
}

func main()  {
	x := error{xml.Name{"",""},200,"Error Message"}
	j,_ := json.Marshal(x)
	fmt.Printf("%v\n%v\n",x, string(j))
}