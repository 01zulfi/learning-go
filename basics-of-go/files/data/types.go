package data

type integer = int
type distance float32
type distanceKm float32
type location string

func distancTo(dest location) distance {
	return distance(10)
}

// method on type
func (miles distance) toKM() distanceKm {
	return distanceKm(1.6 * miles)
}

func (km distanceKm) toMiles() distance {
	return distance(km / 1.6)
}

func test() {
	d := distance(2.4)
	d.toKM()
	print(d)
}
