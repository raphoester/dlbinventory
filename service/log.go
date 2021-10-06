package service

import (
	"fmt"
	"os"
)

func WriteToLog(line string) error {
	f, err := os.OpenFile("output.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return fmt.Errorf("failed opening log file | %s", err.Error())
	}
	defer f.Close()

	if _, err = f.WriteString(line); err != nil {
		return fmt.Errorf("failed writing text to log | %s", err.Error())
	}
	return nil
}
