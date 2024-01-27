package metrics

import (
	"encoding/json"
	"time"
)

// results are private to the module. It will store all results from all Run() function invocations.
var results = make(Results)

// JSONBytes is a named type over []byte to allow String() helper method.
type JSONBytes []byte

func (j JSONBytes) String() string {
	return string(j)
}

// Results contains a mapping of all runs and their associated results.
type Results map[string]Result

// ToJSON will return a []byte array of Results.
func (r Results) ToJSON() JSONBytes {
	data, _ := json.Marshal(r)
	return data
}

// Result represents a set of metrics for a singular run.
type Result struct {
	StartTime time.Time `dynamodbav:"StartTime" json:"start_time,omitempty"`
	EndTime   time.Time `dynamodbav:"EndTime" json:"end_time,omitempty"`
	TotalTime int64     `dynamodbav:"TotalTime" json:"total_time,omitempty"`
}

// Run will execute a given function and measure/record metrics associated with that function for a given name.
func Run(name string, fn func()) {
	var result Result
	result.StartTime = time.Now()
	fn()
	result.EndTime = time.Now()
	result.TotalTime = result.EndTime.Sub(result.StartTime).Microseconds()
	results[name] = result
}

// GetResults will return the set of results as type Results from all Run() function invocations.
func GetResults() Results {
	return results
}
