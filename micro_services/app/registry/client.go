package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buff := new(bytes.Buffer)
	enc := json.NewEncoder(buff)
	err := enc.Encode(r)

	if err != nil {
		return err
	}

	res, err := http.Post(ServiceUrl, "application/json", buff)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service with code %v", res.StatusCode)
	}
	return nil
}

func ShutdownService(serviceUrl string) error {
	fmt.Println("serviceUrl >>>>", serviceUrl)
	req, err := http.NewRequest(http.MethodDelete, ServiceUrl, bytes.NewBuffer([]byte(serviceUrl)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to deregister the service. Registry responded with code : %v ", res.StatusCode)
	}
	return err
}
