package observability

import "github.com/prometheus/client_golang/prometheus"

var OrdersCreated = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "orders_created_total",
		Help: "Total number of orders created",
	},
)

func Init() {
	prometheus.MustRegister(OrdersCreated)
}
