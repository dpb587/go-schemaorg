package vehicle

import "github.com/dpb587/go-schemaorg"

var (
	// Indicates whether the vehicle has been used for special purposes, like
	// commercial rental, driving school, or as a taxi. The legislation in many
	// countries requires this information to be revealed when offering a car for
	// sale.
	VehicleSpecialUsage = schemaorg.NewProperty("vehicleSpecialUsage")

	// The total number of forward gears available for the transmission system of
	// the vehicle.</p>
	// 
	// <p>Typical unit code(s): C62
	NumberOfForwardGears = schemaorg.NewProperty("numberOfForwardGears")

	// The total distance travelled by the particular vehicle since its initial
	// production, as read from its odometer.</p>
	// 
	// <p>Typical unit code(s): KMT for kilometers, SMI for statute miles
	MileageFromOdometer = schemaorg.NewProperty("mileageFromOdometer")

	// The color or color combination of the interior of the vehicle.
	VehicleInteriorColor = schemaorg.NewProperty("vehicleInteriorColor")

	// The position of the steering wheel or similar device (mostly for cars).
	SteeringPosition = schemaorg.NewProperty("steeringPosition")

	// Information about the engine or engines of the vehicle.
	VehicleEngine = schemaorg.NewProperty("vehicleEngine")

	// The release date of a vehicle model (often used to differentiate versions of
	// the same make and model).
	VehicleModelDate = schemaorg.NewProperty("vehicleModelDate")

	// The number of doors.</p>
	// 
	// <p>Typical unit code(s): C62
	NumberOfDoors = schemaorg.NewProperty("numberOfDoors")

	// A short text indicating the configuration of the vehicle, e.g. '5dr hatchback
	// ST 2.5 MT 225 hp' or 'limited edition'.
	VehicleConfiguration = schemaorg.NewProperty("vehicleConfiguration")

	// The type of fuel suitable for the engine or engines of the vehicle. If the
	// vehicle has only one engine, this property can be attached directly to the
	// vehicle.
	FuelType = schemaorg.NewProperty("fuelType")

	// The Vehicle Identification Number (VIN) is a unique serial number used by the
	// automotive industry to identify individual motor vehicles.
	VehicleIdentificationNumber = schemaorg.NewProperty("vehicleIdentificationNumber")

	// <p>The amount of fuel consumed for traveling a particular distance or
	// temporal duration with the given vehicle (e.g. liters per 100 km).</p>
	// 
	// <ul>
	// <li>Note 1: There are unfortunately no standard unit codes for liters per 100
	// km.  Use <a class="localLink" href="http://schema.org/unitText">unitText</a>
	// to indicate the unit of measurement, e.g. L/100 km.</li>
	// <li>Note 2: There are two ways of indicating the fuel consumption, <a
	// class="localLink"
	// href="http://schema.org/fuelConsumption">fuelConsumption</a> (e.g. 8 liters
	// per 100 km) and <a class="localLink"
	// href="http://schema.org/fuelEfficiency">fuelEfficiency</a> (e.g. 30 miles per
	// gallon). They are reciprocal.</li>
	// <li>Note 3: Often, the absolute value is useful only when related to driving
	// speed ("at 80 km/h") or usage pattern ("city traffic"). You can use <a
	// class="localLink" href="http://schema.org/valueReference">valueReference</a>
	// to link the value for the fuel consumption to another value.</li>
	// </ul>
	// 
	FuelConsumption = schemaorg.NewProperty("fuelConsumption")

	// The number of owners of the vehicle, including the current one.</p>
	// 
	// <p>Typical unit code(s): C62
	NumberOfPreviousOwners = schemaorg.NewProperty("numberOfPreviousOwners")

	// <p>The distance traveled per unit of fuel used; most commonly miles per
	// gallon (mpg) or kilometers per liter (km/L).</p>
	// 
	// <ul>
	// <li>Note 1: There are unfortunately no standard unit codes for miles per
	// gallon or kilometers per liter. Use <a class="localLink"
	// href="http://schema.org/unitText">unitText</a> to indicate the unit of
	// measurement, e.g. mpg or km/L.</li>
	// <li>Note 2: There are two ways of indicating the fuel consumption, <a
	// class="localLink"
	// href="http://schema.org/fuelConsumption">fuelConsumption</a> (e.g. 8 liters
	// per 100 km) and <a class="localLink"
	// href="http://schema.org/fuelEfficiency">fuelEfficiency</a> (e.g. 30 miles per
	// gallon). They are reciprocal.</li>
	// <li>Note 3: Often, the absolute value is useful only when related to driving
	// speed ("at 80 km/h") or usage pattern ("city traffic"). You can use <a
	// class="localLink" href="http://schema.org/valueReference">valueReference</a>
	// to link the value for the fuel economy to another value.</li>
	// </ul>
	// 
	FuelEfficiency = schemaorg.NewProperty("fuelEfficiency")

	// The number of axles.</p>
	// 
	// <p>Typical unit code(s): C62
	NumberOfAxles = schemaorg.NewProperty("numberOfAxles")

	// The type or material of the interior of the vehicle (e.g. synthetic fabric,
	// leather, wood, etc.). While most interior types are characterized by the
	// material used, an interior type can also be based on vehicle usage or target
	// audience.
	VehicleInteriorType = schemaorg.NewProperty("vehicleInteriorType")

	// A textual description of known damages, both repaired and unrepaired.
	KnownVehicleDamages = schemaorg.NewProperty("knownVehicleDamages")

	// The number or type of airbags in the vehicle.
	NumberOfAirbags = schemaorg.NewProperty("numberOfAirbags")

	// The number of passengers that can be seated in the vehicle, both in terms of
	// the physical space available, and in terms of limitations set by law.</p>
	// 
	// <p>Typical unit code(s): C62 for persons.
	VehicleSeatingCapacity = schemaorg.NewProperty("vehicleSeatingCapacity")

	// The available volume for cargo or luggage. For automobiles, this is usually
	// the trunk volume.</p>
	// 
	// <p>Typical unit code(s): LTR for liters, FTQ for cubic foot/feet</p>
	// 
	// <p>Note: You can use <a class="localLink"
	// href="http://schema.org/minValue">minValue</a> and <a class="localLink"
	// href="http://schema.org/maxValue">maxValue</a> to indicate ranges.
	CargoVolume = schemaorg.NewProperty("cargoVolume")

	// The type of component used for transmitting the power from a rotating power
	// source to the wheels or other relevant component(s) ("gearbox" for cars).
	VehicleTransmission = schemaorg.NewProperty("vehicleTransmission")

	// The date of the first registration of the vehicle with the respective public
	// authorities.
	DateVehicleFirstRegistered = schemaorg.NewProperty("dateVehicleFirstRegistered")

	// The date the item e.g. vehicle was purchased by the current owner.
	PurchaseDate = schemaorg.NewProperty("purchaseDate")

	// The date of production of the item, e.g. vehicle.
	ProductionDate = schemaorg.NewProperty("productionDate")

	// The drive wheel configuration, i.e. which roadwheels will receive torque from
	// the vehicle's engine via the drivetrain.
	DriveWheelConfiguration = schemaorg.NewProperty("driveWheelConfiguration")
)
