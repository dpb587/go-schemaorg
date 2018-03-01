package propertyvaluespecification

import "github.com/dpb587/go-schemaorg"

var (
	// The default value of the input.  For properties that expect a literal, the
	// default is a literal value, for properties that expect an object, it's an ID
	// reference to one of the current values.
	DefaultValue = schemaorg.NewProperty("defaultValue")

	// The lower value of some characteristic or property.
	MinValue = schemaorg.NewProperty("minValue")

	// Specifies a regular expression for testing literal values according to the
	// HTML spec.
	ValuePattern = schemaorg.NewProperty("valuePattern")

	// The stepValue attribute indicates the granularity that is expected (and
	// required) of the value in a PropertyValueSpecification.
	StepValue = schemaorg.NewProperty("stepValue")

	// Specifies the allowed range for number of characters in a literal value.
	ValueMaxLength = schemaorg.NewProperty("valueMaxLength")

	// Whether or not a property is mutable.  Default is false. Specifying this for
	// a property that also has a value makes it act similar to a "hidden" input in
	// an HTML form.
	ReadonlyValue = schemaorg.NewProperty("readonlyValue")

	// Specifies the minimum allowed range for number of characters in a literal
	// value.
	ValueMinLength = schemaorg.NewProperty("valueMinLength")

	// The upper value of some characteristic or property.
	MaxValue = schemaorg.NewProperty("maxValue")

	// Indicates the name of the PropertyValueSpecification to be used in URL
	// templates and form encoding in a manner analogous to HTML's input@name.
	ValueName = schemaorg.NewProperty("valueName")

	// Whether multiple values are allowed for the property.  Default is false.
	MultipleValues = schemaorg.NewProperty("multipleValues")

	// Whether the property must be filled in to complete the action.  Default is
	// false.
	ValueRequired = schemaorg.NewProperty("valueRequired")
)
