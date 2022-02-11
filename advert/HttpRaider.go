package advert

type HttpRaider interface {
	StartRaid(maxConcurrent int, maxTasks int) error
	StopRaid() error
}
