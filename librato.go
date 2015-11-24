// Go client for Librato Metrics
//
// <https://github.com/rcrowley/go-librato>
package librato

type Metrics interface {
	Close()
	GetCounter(name string) chan float64
	GetCustomCounter(name string) chan map[string]float64
	GetCustomGauge(name string) chan map[string]float64
	GetGauge(name string) chan float64
	NewCounter(name string) chan float64
	NewCustomCounter(name string) chan map[string]float64
	NewCustomGauge(name string) chan map[string]float64
	NewGauge(name string) chan float64
	Wait()
}

func handle(i interface{}, bodyMetric tmetric) bool {
	var obj map[string]float64
	var ok bool
	switch ch := i.(type) {
	case chan float64:
		bodyMetric["value"], ok = <-ch
	case chan map[string]float64:
		obj, ok = <-ch
		for k, v := range obj {
			bodyMetric[k] = v
		}
	}
	return ok
}

// models http://dev.librato.com/v1/post/metrics (3) Array format (JSON only)
type tbody map[string]tibody
type tibody []tmetric
type tmetric map[string]interface{}
