package softwareapplication

import "github.com/dpb587/go-schemaorg"

var (
	// If the file can be downloaded, URL to download the binary.
	DownloadUrl = schemaorg.NewProperty("downloadUrl")

	// Component dependency requirements for application. This includes runtime
	// environments and shared libraries that are not included in the application
	// distribution package, but required to run the application (Examples: DirectX,
	// Java or .NET runtime).
	SoftwareRequirements = schemaorg.NewProperty("softwareRequirements")

	// Permission(s) required to run the app (for example, a mobile app may require
	// full internet access or may run only on wifi).
	Permissions = schemaorg.NewProperty("permissions")

	// Processor architecture required to run the application (e.g. IA64).
	ProcessorRequirements = schemaorg.NewProperty("processorRequirements")

	// Device required to run the application. Used in cases where a specific
	// make/model is required to run the application.
	AvailableOnDevice = schemaorg.NewProperty("availableOnDevice")

	// Features or modules provided by this application (and possibly required by
	// other applications).
	FeatureList = schemaorg.NewProperty("featureList")

	// Subcategory of the application, e.g. 'Arcade Game'.
	ApplicationSubCategory = schemaorg.NewProperty("applicationSubCategory")

	// Device required to run the application. Used in cases where a specific
	// make/model is required to run the application.
	Device = schemaorg.NewProperty("device")

	// Type of software application, e.g. 'Game, Multimedia'.
	ApplicationCategory = schemaorg.NewProperty("applicationCategory")

	// Version of the software instance.
	SoftwareVersion = schemaorg.NewProperty("softwareVersion")

	// Storage requirements (free space required).
	StorageRequirements = schemaorg.NewProperty("storageRequirements")

	// The name of the application suite to which the application belongs (e.g.
	// Excel belongs to Office).
	ApplicationSuite = schemaorg.NewProperty("applicationSuite")

	// Minimum memory requirements.
	MemoryRequirements = schemaorg.NewProperty("memoryRequirements")

	// A link to a screenshot image of the app.
	Screenshot = schemaorg.NewProperty("screenshot")

	// Countries for which the application is supported. You can also provide the
	// two-letter ISO 3166-1 alpha-2 country code.
	CountriesSupported = schemaorg.NewProperty("countriesSupported")

	// Software application help.
	SoftwareHelp = schemaorg.NewProperty("softwareHelp")

	// Additional content for a software application.
	SoftwareAddOn = schemaorg.NewProperty("softwareAddOn")

	// Description of what changed in this version.
	ReleaseNotes = schemaorg.NewProperty("releaseNotes")

	// Supporting data for a SoftwareApplication.
	SupportingData = schemaorg.NewProperty("supportingData")

	// Component dependency requirements for application. This includes runtime
	// environments and shared libraries that are not included in the application
	// distribution package, but required to run the application (Examples: DirectX,
	// Java or .NET runtime).
	Requirements = schemaorg.NewProperty("requirements")

	// Countries for which the application is not supported. You can also provide
	// the two-letter ISO 3166-1 alpha-2 country code.
	CountriesNotSupported = schemaorg.NewProperty("countriesNotSupported")

	// Operating systems supported (Windows 7, OSX 10.6, Android 1.6).
	OperatingSystem = schemaorg.NewProperty("operatingSystem")

	// Size of the application / package (e.g. 18MB). In the absence of a unit (MB,
	// KB etc.), KB will be assumed.
	FileSize = schemaorg.NewProperty("fileSize")

	// URL at which the app may be installed, if different from the URL of the item.
	InstallUrl = schemaorg.NewProperty("installUrl")
)
