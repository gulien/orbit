/*
Package context helps to populate the application context.

The main goal of the application context is to gather all the data which will be applied to a data-driven template.
*/
package context

import (
	"github.com/gulien/orbit/errors"
	"github.com/gulien/orbit/helpers"
	"github.com/gulien/orbit/logger"
)

// OrbitContext contains the data necessary for executing a data-driven template.
type OrbitContext struct {
	// TemplateFilePath is the path of a data-driven template.
	TemplateFilePath string

	// Payload map contains data from various entries.
	Payload map[string]interface{}
}

// NewOrbitContext creates an instance of OrbitContext.
func NewOrbitContext(templateFilePath string, payload string) (*OrbitContext, error) {
	// as the data-driven template is mandatory, we must check its validity.
	if templateFilePath == "" {
		return nil, errors.NewOrbitErrorf("no data-driven template given")
	}

	if !helpers.FileExists(templateFilePath) {
		return nil, errors.NewOrbitErrorf("the data-driven template %s does not exist", templateFilePath)
	}

	// let's instantiates our OrbitContext!
	ctx := &OrbitContext{
		TemplateFilePath: templateFilePath,
	}

	logger.Debugf("context has been instantiated with the data-driven template %s", ctx.TemplateFilePath)

	// last but not least, instantiates an orbitPayload which will allow us
	// to retrieves the data provided by the entries given by the user.
	p := &orbitPayload{}
	p.populateFromFile("")
	p.populateFromString(payload)

	data, err := p.retrieveData()
	if err != nil {
		return nil, err
	}

	ctx.Payload = data

	logger.Debugf("context has been populated with payload %s", ctx.Payload)

	return ctx, nil
}
