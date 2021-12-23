package advert

type Advert interface {
	StartRaid() error
	StopRaid() error
}

type MutableFormsOutput interface {
}
