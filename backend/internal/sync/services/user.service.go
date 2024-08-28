package services

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/victorlujan/tentacle/backend/models"
)

func GetUsers() (users models.UserEnvelope, err error) {

	var url string = os.Getenv("BC_HOST")
	url = url + "/USUARIOS_WS"
	var method string = "POST"
	var basicAuth string = os.Getenv("BC_BASIC_AUTH")

	payload := strings.NewReader(`<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
	<Body>
		<ReadMultiple xmlns="urn:microsoft-dynamics-schemas/page/usuarios_ws">
			<filter>
				<Field></Field>
				<Criteria></Criteria>
			</filter>
			<bookmarkKey></bookmarkKey>
			<setSize></setSize>
		</ReadMultiple>
	</Body>
  </Envelope>`)

	if url == "" || basicAuth == "" {
		return models.UserEnvelope{}, fmt.Errorf("error: url or basicAuth is empty")
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("SOAPAction", "#POST")
	req.Header.Add("Authorization", basicAuth)

	if err != nil {
		return models.UserEnvelope{}, err
	}

	res, err := client.Do(req)

	if err != nil {
		return models.UserEnvelope{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return models.UserEnvelope{}, errors.New("Error getting users from BC. Status Code " + res.Status)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return models.UserEnvelope{}, err
	}
	var envolveUsers models.UserEnvelope
	err = xml.Unmarshal(body, &envolveUsers)

	if err != nil {
		return models.UserEnvelope{}, err
	}

	return envolveUsers, nil

}
