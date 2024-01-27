package metrics

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name            string
		inFunc          func()
		expectedResults Results
	}{
		{
			name: "test run",
			inFunc: func() {
				fmt.Println("testing")
			},
		},
		{
			name: "sleep run",
			inFunc: func() {
				time.Sleep(time.Second)
			},
		},
		{
			name: "test JSONBytes String()",
			inFunc: func() {
				fmt.Println(GetResults().ToJSON().String())
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			Run(testCase.name, testCase.inFunc)
			if _, ok := GetResults()[testCase.name]; !ok {
				t.Errorf("expected: [%s] to be in results but got nothing", testCase.name)
			}
		})
	}
}
