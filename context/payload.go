package context

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gulien/orbit/errors"
	"github.com/gulien/orbit/helpers"
	"github.com/gulien/orbit/logger"

	"gopkg.in/yaml.v2"
)

// default payload file path.
const payloadFilePath = "orbit-payload.yml"

type (
	// orbitPayload contains the various entries provided by the user.
	orbitPayload struct {
		// PayloadEntries is a simple array of orbitPayloadEntry.
		PayloadEntries []*orbitPayloadEntry `yaml:"payload"`
	}

	// orbitPayloadEntry is an entry from a file or from a string.
	orbitPayloadEntry struct {
		// Key is the unique identifier of a value
		Key string `yaml:"key"`

		// Value is a raw data or a file path.
		Value string `yaml:"value"`
	}
)

// populateFromFile populates the instance of orbitPayload
// with entries provided by a YAML file.
func (p *orbitPayload) populateFromFile(filePath string) error {
	if filePath == "" {
		filePath = payloadFilePath
	}

	// first, checks if the file exists.
	if !helpers.FileExists(filePath) {
		logger.Debugf("payload file %s does not exist, skipping", filePath)
		return nil
	}

	// alright, let's read it to retrieve its data...
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.NewOrbitErrorf("unable to read the payload file %s. Details:\n%s", filePath, err)
	}

	// ...then parses the YAML.
	if err := yaml.Unmarshal(data, &p); err != nil {
		return errors.NewOrbitErrorf("payload file %s is not a valid YAML file. Details:\n%s", filePath, err)
	}

	return nil
}

// populateFromString populates the instance of orbitPayload
// with entries provided by a string.
func (p *orbitPayload) populateFromString(payload string) error {
	// first, checks if a payload has been given.
	if payload == "" {
		logger.Debugf("no payload given, skipping")
		return nil
	}

	// the payload string should be in the following format:
	// key,value;key,value.
	entries := strings.Split(payload, ";")
	for _, entry := range entries {
		entry := strings.Split(entry, ",")
		if len(entry) != 2 {
			return errors.NewOrbitErrorf("unable to process the payload entry %s", entry)
		}

		p.PayloadEntries = append(p.PayloadEntries, &orbitPayloadEntry{
			Key:   entry[0],
			Value: entry[1],
		})
	}

	return nil
}

// retrieveData parses all the entries from the instance of orbitPayload
// to retrieve the data which will be applied to a data-driven template.
func (p *orbitPayload) retrieveData() (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for _, payloadEntry := range p.PayloadEntries {
		d := getDecoder(payloadEntry.Value)

		value, err := d.decode()
		if err != nil {
			return nil, err
		}

		result[payloadEntry.Key] = value
	}

	return result, nil
}

// getDecoder returns the correct decoder for a given value.
func getDecoder(value string) orbitDecoder {
	if !helpers.FileExists(value) {
		return &orbitDumbDecoder{value: value}
	}

	switch extension := filepath.Ext(value); extension {
	case ".yaml":
	case ".yml":
		return &orbitYAMLDecoder{value: value}
	case ".toml":
		return &orbitTOMLDecoder{value: value}
	case ".json":
		return &orbitJSONDecoder{value: value}
	}

	return &orbitEnvFileDecoder{value: value}
}
