package libmig_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/mazrean/libmig"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, libmig.Analyzer, "a")
}
