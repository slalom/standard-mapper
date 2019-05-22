package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

func unmarshal(input string, result interface{}) error {
	source, err := ioutil.ReadFile(input)
	if err != nil {
		logrus.Errorf("could not open file %s. Error %s", input, err.Error())
		return err
	}

	err = json.Unmarshal(source, result)

	if jsonError, ok := err.(*json.SyntaxError); ok {
		line, character, lcErr := lineAndCharacter(string(source), int(jsonError.Offset))
		logrus.Errorf("Cannot parse JSON schema due to a syntax error at line %d, character %d: %v\n", line, character, jsonError.Error())
		if lcErr != nil {
			logrus.Errorf("Couldn't find the line and character position of the error due to error %v\n", lcErr)
		}
	}
	if jsonError, ok := err.(*json.UnmarshalTypeError); ok {
		line, character, lcErr := lineAndCharacter(string(source), int(jsonError.Offset))
		logrus.Errorf("The JSON type '%v' cannot be converted into the Go '%v' type on struct '%s', field '%v'. See input file line %d, character %d\n", jsonError.Value, jsonError.Type.Name(), jsonError.Struct, jsonError.Field, line, character)
		if lcErr != nil {
			fmt.Fprintf(os.Stderr, "Couldn't find the line and character position of the error due to error %v\n", lcErr)
		}
	}

	if err != nil {
		logrus.Errorf("could not unmarshall file %s. Error %s", input, err.Error())
		return err
	}

	return nil
}

func lineAndCharacter(input string, offset int) (line int, character int, err error) {
	lf := rune(0x0A)

	if offset > len(input) || offset < 0 {
		return 0, 0, fmt.Errorf("Couldn't find offset %d within the input.", offset)
	}

	// Humans tend to count from 1.
	line = 1

	for i, b := range input {
		if b == lf {
			line++
			character = 0
		}
		character++
		if i == offset {
			break
		}
	}

	return line, character, nil
}

func serializeToDisk(outputFilename string, t interface{}) error {
	file, err := os.Create(outputFilename)
	defer file.Close()
	if err != nil {
		return err
	}

	js, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	file.Write(js)

	return nil
}
