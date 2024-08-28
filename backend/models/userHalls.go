package models

import (
	"encoding/xml"
)

type UserHalls struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"Soap,attr"`
	Body    struct {
		Text               string `xml:",chardata"`
		ReadMultipleResult struct {
			Text               string `xml:",chardata"`
			Xmlns              string `xml:"xmlns,attr"`
			ReadMultipleResult struct {
				Text         string         `xml:",chardata"`
				USERSALON_WS []USERSALON_WS `xml:"USERSALON_WS"`
			} `xml:"ReadMultiple_Result"`
		} `xml:"ReadMultiple_Result"`
	} `xml:"Body"`
}

type USERSALON_WS struct {
	Text    string `xml:",chardata"`
	Key     string `xml:"Key"`
	Usuario string `xml:"Usuario"`
	Salon   string `xml:"Salon"`
}
