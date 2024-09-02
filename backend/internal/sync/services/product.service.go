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

func GetProducts() (products models.Products, err error) {
	var url string = os.Getenv("BC_HOST")
	url = url + "/PRODUCTOS_WS"
	var method string = "POST"
	var basicAuth string = os.Getenv("BC_BASIC_AUTH")

	payload := strings.NewReader(`<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
    <Body>
        <ReadMultiple xmlns="urn:microsoft-dynamics-schemas/page/productos_ws">
            <bookmarkKey></bookmarkKey>
            <setSize></setSize>
        </ReadMultiple>
    </Body>
</Envelope>`)

	if url == "" || basicAuth == "" {
		return models.Products{}, fmt.Errorf("error: url or basicAuth is empty")
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("SOAPAction", "#POST")
	req.Header.Add("Authorization", basicAuth)

	if err != nil {
		return models.Products{}, err
	}

	res, err := client.Do(req)

	if err != nil {
		return models.Products{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return models.Products{}, err
	}
	var envolveUserHalls models.Products
	err = xml.Unmarshal(body, &envolveUserHalls)

	if err != nil {
		return models.Products{}, err
	}

	return envolveUserHalls, nil

}
