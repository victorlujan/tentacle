package services

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/victorlujan/tentacle/backend/models"
)

func GetHalls() (halls models.Halls, err error) {
	var url string = os.Getenv("BC_HOST")
	url = url + "/CONFSALONES_WS"
	var method string = "POST"
	var basicAuth string = os.Getenv("BC_BASIC_AUTH")

	payload := strings.NewReader(`<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
    <Body>
        <ReadMultiple xmlns="urn:microsoft-dynamics-schemas/page/confsalones_ws">
            <bookmarkKey></bookmarkKey>
            <setSize></setSize>
        </ReadMultiple>
    </Body>
</Envelope>`)

	if url == "" || basicAuth == "" {
		return models.Halls{}, fmt.Errorf("error: url or basicAuth is empty")
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("SOAPAction", "#POST")
	req.Header.Add("Authorization", basicAuth)

	if err != nil {
		return models.Halls{}, err
	}

	res, err := client.Do(req)

	if err != nil {
		return models.Halls{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return models.Halls{}, err
	}
	var envolveHalls models.Halls
	err = xml.Unmarshal(body, &envolveHalls)

	if err != nil {
		return models.Halls{}, err
	}

	return envolveHalls, nil

}
