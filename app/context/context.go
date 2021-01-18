/*
Package context helps to populate the application context.

The main goal of the application context is to gather all the data which will be applied to a data-driven template.
*/
package context

import (
	OrbitError "github.com/gulien/orbit/app/error"
	"github.com/gulien/orbit/app/helpers"
	"github.com/gulien/orbit/app/logger"
)

// OrbitContext contains the data necessary for executing a data-driven template.
type OrbitContext struct {
	// TemplateFilePath is the path of a data-driven template.
	TemplateFilePath string

	// Payload map contains data from various entries.
	Payload map[string]interface{}

	// Templates array contains the list of additional templates to parse.
	Templates []string

	// Optional pair of template delimiters (used to override go defaults "{{" and "}}")
	TemplateDelimiters []string
}

// NewOrbitContext creates an instance of OrbitContext.
func NewOrbitContext(templateFilePath string, payload string, templates string, templateDelimiters []string) (*OrbitContext, error) {
	// as the data-driven template is mandatory, we must check its validity.
	if templateFilePath == "" {
		return nil, OrbitError.NewOrbitErrorf("no data-driven template given")
	}

	if !helpers.FileExists(templateFilePath) {
		return nil, OrbitError.NewOrbitErrorf("the data-driven template %s does not exist", templateFilePath)
	}

	// let's instantiates our OrbitContext!
	ctx := &OrbitContext{
		TemplateFilePath: templateFilePath,
	}

	logger.Debugf("context has been instantiated with the data-driven template %s", ctx.TemplateFilePath)

	// last but not least, instantiates an orbitPayload which will allow us
	// to retrieves the data provided by the entries given by the user.
	p := &orbitPayload{}

	if err := p.populateFromFile(""); err != nil {
		return nil, err
	}

	if err := p.populateFromString(payload, templates); err != nil {
		return nil, err
	}

	payloadData, err := p.retrievePayloadData()
	if err != nil {
		return nil, err
	}

	ctx.Payload = payloadData
	logger.Debugf("context has been populated with payload %s", ctx.Payload)

	ctx.Templates = p.TemplatesEntries
	logger.Debugf("context has been populated with templates %s", ctx.Templates)

	if templateDelimiters == nil {
		ctx.TemplateDelimiters = make([]string, 2)
	} else if len(templateDelimiters) == 2 {
		ctx.TemplateDelimiters = templateDelimiters
		logger.Debugf("context has been instantiated with the template delimiters %+v", ctx.TemplateDelimiters)
	} else {
		return nil, OrbitError.NewOrbitErrorf("%d delimiter(s) specified: %+v. Exactly two (left,right) must be specified", len(templateDelimiters), templateDelimiters)
	}

	return ctx, nil
}
