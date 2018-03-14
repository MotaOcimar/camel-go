package camel

// Endpoint --
type Endpoint interface {
	Service

	Component() Component

	CreateProducer() (Producer, error)
	CreateConsumer() (Consumer, error)
}
