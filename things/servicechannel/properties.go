package servicechannel

import "github.com/dpb587/go-schemaorg"

var (
	// The service provided by this channel.
	ProvidesService = schemaorg.NewProperty("providesService")

	// The location (e.g. civic structure, local business, etc.) where a person can
	// go to access the service.
	ServiceLocation = schemaorg.NewProperty("serviceLocation")

	// The address for accessing the service by mail.
	ServicePostalAddress = schemaorg.NewProperty("servicePostalAddress")

	// A language someone may use with or at the item, service or place. Please use
	// one of the language codes from the <a
	// href="http://tools.ietf.org/html/bcp47">IETF BCP 47 standard</a>. See also <a
	// class="localLink" href="http://schema.org/inLanguage">inLanguage</a>
	AvailableLanguage = schemaorg.NewProperty("availableLanguage")

	// The number to access the service by text message.
	ServiceSmsNumber = schemaorg.NewProperty("serviceSmsNumber")

	// The website to access the service.
	ServiceUrl = schemaorg.NewProperty("serviceUrl")

	// The phone number to use to access the service.
	ServicePhone = schemaorg.NewProperty("servicePhone")

	// Estimated processing time for the service using this channel.
	ProcessingTime = schemaorg.NewProperty("processingTime")
)
