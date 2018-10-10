package shuttle

/*
A shuttle is a front-end client that lives in a main process that is constantly
communicating with other elements. LocalSensors live on the client process but
remote sensors are network addresses. The network latency can be compensated for
by subracting the latency from the injected latency.
*/
type Shuttle struct {
	localSensors map[SensorType]*Sensor

	remoteSensors []*Sensor // Confidence cone?

	// Everything Below is theoretical until I can nail down the sensors.
	localOrdinance map[OrdinanceType]*Ordinance

	remote map[string]Controllable //(remote sensors could be nested under this interface.)

	manuvoring Thrusters // Probably ignore for now.
	main       Thrusters

	reactor Reactor
	fuel    map[FuelType]*Fuel

	// Buffs to certain elements, but add weight and can die.
	personelle []Personelle

	shielding []Shielding

	commandCenter string // In reality, its a grpc address.
}
