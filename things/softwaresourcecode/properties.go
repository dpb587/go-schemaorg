package softwaresourcecode

import "github.com/dpb587/go-schemaorg"

var (
	// Target Operating System / Product to which the code applies.  If applies to
	// several versions, just the product name can be used.
	TargetProduct = schemaorg.NewProperty("targetProduct")

	// Link to the repository where the un-compiled, human readable code and related
	// code is located (SVN, github, CodePlex).
	CodeRepository = schemaorg.NewProperty("codeRepository")

	// The computer programming language.
	ProgrammingLanguage = schemaorg.NewProperty("programmingLanguage")

	// What type of code sample: full (compile ready) solution, code snippet, inline
	// code, scripts, template.
	CodeSampleType = schemaorg.NewProperty("codeSampleType")

	// Runtime platform or script interpreter dependencies (Example - Java v1,
	// Python2.3, .Net Framework 3.0).
	RuntimePlatform = schemaorg.NewProperty("runtimePlatform")

	// What type of code sample: full (compile ready) solution, code snippet, inline
	// code, scripts, template.
	SampleType = schemaorg.NewProperty("sampleType")

	// Runtime platform or script interpreter dependencies (Example - Java v1,
	// Python2.3, .Net Framework 3.0).
	Runtime = schemaorg.NewProperty("runtime")
)
