package vra

import (
	"fmt"
	"regexp"
)

func validateBlueprintName(v interface{}, k string) (warnings []string, errors []error) {
	blueprintName := v.(string)
	match, err := regexp.MatchString("^[^*?:<>|/[\\]]+$", blueprintName)
	if err != nil {
		return returnError("Regex error: Report bug to terraform", err)
	}
	if !match {
		return returnError("Blueprint Name should not contain special chars (^*?:<>|/\\).", fmt.Errorf("Blueprint Name is not correct"))
	}
	return nil, nil
}

func validateFileName(v interface{}, k string) (warnings []string, errors []error) {
	fileName := v.(string)
	match, err := regexp.MatchString("^[^*?:<>|/[\\]]+$", fileName)
	if err != nil {
		return returnError("Regex error: Report bug to terraform", err)
	}
	if !match {
		return returnError("File Name should not contain special chars (^*?:<>|/\\).", fmt.Errorf("File Name is not correct"))
	}
	return nil, nil
}

func returnError(message string, err error) (warnings []string, errors []error) {
	var errorVar []error
	var warningVar []string
	return append(warningVar, message), append(errorVar, err)
}
