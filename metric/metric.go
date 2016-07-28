package metric

import "io"

type Counter interface {
	Name() string
	Add()
	AddN(int64)
	Reset()
}

type Histogram interface {
	Name() string
	Record(int64) error
	Reset()
}

type Metric interface {
	Name() string
	Counter(string) Counter
	NewCounter(string) (Counter, error)
	Histogram(string) Histogram
	NewHistogram(name string, min, max int64, sigfig int) (Histogram, error)
	Counters() []Counter
	Histograms() []Histogram
	RemoveCounter(Counter) error
	RemoveHistogram(Histogram) error
	Reset()
	Flush(io.Writer) error
}
