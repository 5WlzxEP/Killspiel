package ResultCollector

type Collector interface {
	// Begin blocks until the Condition for the Collector to realize the voting event should start
	Begin()
	// Result blocks until a result is present and returns it
	Result() int
}
