package app

import (
	"github.com/sirupsen/logrus"
	"github.com/tredfield/standard-mapper/pkg/mappers"
	"github.com/tredfield/standard-mapper/pkg/types"
)

// Start processes input and generates output
func Start(input string, output string) error {
	var source types.Gdpr

	// parse json to Gdpr type
	err := unmarshal(input, &source)

	if err != nil {
		return err
	}

	logrus.Infof("unmarshal %s", input)

	// write result for debugging
	serializeToDisk("unmarshal-debug.json", source)

	// map from gdbr to standard
	standard, err := mappers.MapGdpr(source)
	if err != nil {
		return err
	}

	// write standard to output path
	err = serializeToDisk(output, standard)
	if err != nil {
		return err
	}

	return nil
}
