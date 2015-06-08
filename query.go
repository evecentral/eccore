package eccore

// Define a basic query language for fetching orders

type QueryResult interface {
}

type QueryResultAggregate interface {
}

type QueryFetch interface {
	Exec(input QueryResult) (QueryResult, error)
}

type QueryAggregate interface {
	Exec(input QueryResult) (QueryResultAggregate, error)
}

type QueryCompute interface {
	Exec(input QueryResultAggregate) (QueryResultAggregate, error)
}
