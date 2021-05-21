package webAnalyzer

import (
	"strings"
	"testing"
)

func TestGetPageInformation(t *testing.T) {
	for _, tc := range testCases {
		got, _ := GetPageInformation(strings.NewReader(tc.in))
		if got != tc.want {
			t.Errorf("GetPageInformation(strings.NewReader(%v) == %+v, want %+v", tc.in, got, tc.want)
		}
	}
}
