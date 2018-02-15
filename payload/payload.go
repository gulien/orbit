/*
TODO
*/
package payload

import (
	"github.com/gulien/orbit/helpers"
	"github.com/gulien/orbit/logger"
	"strings"
	"github.com/gulien/orbit/errors"
)

const payloadFilePath = "orbit-payload.yml"

type (
	OrbitPayload struct {
		PayloadEntries []*OrbitPayloadEntry
	}

	OrbitPayloadEntry struct {
		Key string

		Value string
	}
)

func NewOrbitPayload(payloadDataSources string) (*OrbitPayload, error) {
	p := &OrbitPayload{}

	p, err := parsePayloadFile(p)
	if err != nil {
		return p, err
	}

	p, err = parsePayloadEntries(p, payloadDataSources)
	if err != nil {
		return p, err
	}

	return p, nil
}

func parsePayloadFile(payload *OrbitPayload) (*OrbitPayload, error) {
	if !helpers.FileExists(payloadFilePath) {
		logger.Debugf("no %s file, skipping", payloadFilePath)
		return nil, nil
	}
}

func parsePayloadEntries(payload *OrbitPayload, payloadDataSources string) (*OrbitPayload, error) {
	var payloadEntries *[]OrbitPayloadEntry


	// checks if the given string is a map of files:
	// if not, considers the string as a path.
	// otherwise tries to populate an array of OrbitFileMap instances.
	entries := strings.Split(payloadDataSources, ";")
	for _, entry := range entries {
		data := strings.Split(entry, ",")
		if len(data) != 2 {
			return payload, errors.NewOrbitErrorf("unable to process the payload entry %s", entry)
		}

		payloadEntries = append(payloadEntries, &OrbitPayloadEntry{
			Key: data[0],
			Value: data[1],
		})
	}
}

func unmarshalYAML(in []byte, out interface{}) error {
	return nil
}

func unmarshalJSON(in []byte, out interface{}) error {
	return nil
}

func unmarshalTOML(in []byte, out interface{}) error {
	return nil
}

func unmarshalEnvFile(in []byte, out interface{}) error {
	return nil
}