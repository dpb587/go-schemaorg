package entrypoint

import "github.com/dpb587/go-schemaorg"

var (
	// An url template (RFC6570) that will be used to construct the target of the
	// execution of the action.
	UrlTemplate = schemaorg.NewProperty("urlTemplate")

	// An application that can complete the request.
	ActionApplication = schemaorg.NewProperty("actionApplication")

	// An application that can complete the request.
	Application = schemaorg.NewProperty("application")

	// The high level platform(s) where the Action can be performed for the given
	// URL. To specify a specific application or operating system instance, use
	// actionApplication.
	ActionPlatform = schemaorg.NewProperty("actionPlatform")

	// An HTTP method that specifies the appropriate HTTP method for a request to an
	// HTTP EntryPoint. Values are capitalized strings as used in HTTP.
	HttpMethod = schemaorg.NewProperty("httpMethod")

	// The supported encoding type(s) for an EntryPoint request.
	EncodingType = schemaorg.NewProperty("encodingType")

	// The supported content type(s) for an EntryPoint response.
	ContentType = schemaorg.NewProperty("contentType")
)
