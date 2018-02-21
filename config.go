package main

const (
	ZoneGen  = "gen"
	ZoneCity = "city"
	ZoneDC   = "dc"
)

var config = &Config{
	LumCustomer: "axixmedtech",
	LumGenZone:  "5d9w46qlbhf2",
	LumCityZone: "w4sirl8ppbqs",
}

type Config struct {
	LumCustomer string
	LumGenZone  string
	LumCityZone string
}
