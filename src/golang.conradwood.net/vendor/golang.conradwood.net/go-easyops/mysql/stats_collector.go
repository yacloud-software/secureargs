package mysql

import (
	"github.com/prometheus/client_golang/prometheus"
)

type PoolSizeCollector struct {
	counterDesc *prometheus.Desc
}

func (c *PoolSizeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.counterDesc
}

func (c *PoolSizeCollector) Collect(ch chan<- prometheus.Metric) {
	for _, db := range databases {
		value := float64(db.dbcon.Stats().OpenConnections)
		ch <- prometheus.MustNewConstMetric(
			c.counterDesc,
			prometheus.CounterValue,
			value,
			db.dbname,
		)

	}
}
func NewPoolSizeCollector() *PoolSizeCollector {
	x := &PoolSizeCollector{
		counterDesc: prometheus.NewDesc("mysql_pool_size", "sql pool size",
			[]string{"database"},
			nil),
	}
	return x
}
