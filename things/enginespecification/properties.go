package enginespecification

import "github.com/dpb587/go-schemaorg"

var (
	// The type of fuel suitable for the engine or engines of the vehicle. If the
	// vehicle has only one engine, this property can be attached directly to the
	// vehicle.
	FuelType = schemaorg.NewProperty("fuelType")
)
