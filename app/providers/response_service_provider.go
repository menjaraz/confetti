package providers

import (
	"github.com/lanvard/contract/inter"
	decorator "github.com/lanvard/foundation/decorator/response_decorator"
	"github.com/lanvard/foundation/encoder"
	"lanvard/resources/views"
	net "net/http"
)

type ResponseServiceProvider struct{}

func (c ResponseServiceProvider) Register(container inter.Container) inter.Container {
	// Response decorators are responsible for modifying the response object.
	// All these decorators will be used.
	container.Bind("response_decorators", []inter.ResponseDecorator{
		decorator.LogError{},
		decorator.FilterSensitiveError{},
		decorator.HttpStatus{ErrorDefault: net.StatusInternalServerError},
		// add your custom decorators here
	})

	// Outcome encoders are responsible for converting an object
	// to a string. One encoder will be used.
	container.Bind("outcome_html_encoders", []inter.Encoder{
		// add your custom HTML encoders here
		encoder.ViewToHtml{},
		encoder.ErrorToHtml{View: views.Error},
		encoder.StringerToHtml{},
		encoder.RawToHtml{},
		encoder.InterfaceToHtml{},
	})

	container.Bind("outcome_json_encoders", []inter.Encoder{
		// add your custom JSON encoders here
		encoder.JsonReaderToJson{},
		encoder.ErrorToJson{},
		encoder.RawToJson{},
		encoder.JsonToJson{},
		encoder.InterfaceToJson{},
	})

	return container
}
