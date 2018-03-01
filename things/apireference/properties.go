package apireference

import "github.com/dpb587/go-schemaorg"

var (
	// Library file name e.g., mscorlib.dll, system.web.dll.
	ExecutableLibraryName = schemaorg.NewProperty("executableLibraryName")

	// Associated product/technology version. e.g., .NET Framework 4.5.
	AssemblyVersion = schemaorg.NewProperty("assemblyVersion")

	// Indicates whether API is managed or unmanaged.
	ProgrammingModel = schemaorg.NewProperty("programmingModel")

	// Library file name e.g., mscorlib.dll, system.web.dll.
	Assembly = schemaorg.NewProperty("assembly")

	// Type of app development: phone, Metro style, desktop, XBox, etc.
	TargetPlatform = schemaorg.NewProperty("targetPlatform")
)
