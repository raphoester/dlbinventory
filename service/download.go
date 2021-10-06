package service

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(URL, fileName string) error {
	response, err := http.Get(URL)
	if err != nil {
		return fmt.Errorf("error downloading resource | %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed creating downloaded file | %s", err.Error())
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("failed copying downloaded content | %s", err.Error())
	}

	return nil
}
