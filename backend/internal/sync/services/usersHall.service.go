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

func GetUserHalls() (userHalls models.UserHalls, err error) {
	var url string = os.Getenv("BC_HOST")
	url = url + "/USERSALON_WS"
	var method string = "POST"
	var basicAuth string = os.Getenv("BC_BASIC_AUTH")

	payload := strings.NewReader(`<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
    <Body>
        <ReadMultiple xmlns="urn:microsoft-dynamics-schemas/page/usersalon_ws">
            <bookmarkKey></bookmarkKey>
            <setSize></setSize>
        </ReadMultiple>
    </Body>
</Envelope>`)

	if url == "" || basicAuth == "" {
		return models.UserHalls{}, fmt.Errorf("error: url or basicAuth is empty")
	}

	client := http.Client{}
	req, err := http.NewRequest(method, url, payload)
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("SOAPAction", "#POST")
	req.Header.Add("Authorization", basicAuth)

	if err != nil {
		return models.UserHalls{}, err
	}

	res, err := client.Do(req)

	if err != nil {
		return models.UserHalls{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return models.UserHalls{}, err
	}
	var envolveUserHalls models.UserHalls
	err = xml.Unmarshal(body, &envolveUserHalls)

	if err != nil {
		return models.UserHalls{}, err
	}

	return envolveUserHalls, nil

}
