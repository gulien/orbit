package context

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	OrbitError "github.com/gulien/orbit/app/error"
	"github.com/gulien/orbit/app/helpers"
	"github.com/gulien/orbit/app/logger"

	"gopkg.in/yaml.v2"
)

// default payload file path.
const payloadFilePath = "orbit-payload.yml"

type (
	// orbitPayload contains the various entries provided by the user.
	orbitPayload struct {
		// PayloadEntries is a simple array of orbitPayloadEntry.
		PayloadEntries []*orbitPayloadEntry `yaml:"payload,omitempty"`

		// TemplatesEntries is a simple array of string.
		TemplatesEntries []string `yaml:"templates,omitempty"`
	}

	// orbitPayloadEntry is an entry from a file or from a string.
	orbitPayloadEntry struct {
		// Key is the unique identifier of a value.
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
		return OrbitError.NewOrbitErrorf("unable to read the payload file %s. Details:\n%s", filePath, err)
	}

	// ...then parses the YAML.
	if err := yaml.Unmarshal(data, &p); err != nil {
		return OrbitError.NewOrbitErrorf("payload file %s is not a valid YAML file. Details:\n%s", filePath, err)
	}

	return nil
}

// populateFromString populates the instance of orbitPayload
// with entries provided by a string.
func (p *orbitPayload) populateFromString(payload string, templates string) error {
	// first, checks if a payload has been given.
	if payload == "" && templates == "" {
		logger.Debugf("no payload and templates flags given, skipping")
		return nil
	}

	if payload != "" {
		// the payload string should be in the following format:
		// key,value;key,value.
		entries := strings.Split(payload, ";")
		for _, entry := range entries {
			entry := strings.Split(entry, ",")
			if len(entry) == 1 {
				return OrbitError.NewOrbitErrorf("unable to process the payload entry %s", entry)
			}

			p.PayloadEntries = append(p.PayloadEntries, &orbitPayloadEntry{
				Key:   entry[0],
				Value: strings.Join(entry[1:], ","),
			})
		}
	}

	if templates != "" {
		// the templates string should be in the following format:
		// path,path,path.
		entries := strings.Split(templates, ",")
		for _, entry := range entries {
			p.TemplatesEntries = append(p.TemplatesEntries, entry)
		}
	}

	return nil
}

// retrievePayloadData parses all the payload entries from the instance of orbitPayload
// to retrieve the data which will be applied to a data-driven template.
func (p *orbitPayload) retrievePayloadData() (map[string]interface{}, error) {
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
	case ".yaml", ".yml":
		return &orbitYAMLDecoder{value: value}
	case ".toml":
		return &orbitTOMLDecoder{value: value}
	case ".json":
		return &orbitJSONDecoder{value: value}
	}

	return &orbitEnvFileDecoder{value: value}
}
