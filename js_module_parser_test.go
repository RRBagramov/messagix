package messagix_test

import (
	"log"
	"os"
	"testing"

	"github.com/RRBagramov/messagix"
	"github.com/RRBagramov/messagix/debug"
	"github.com/RRBagramov/messagix/types"
)

func TestParseJS(t *testing.T) {
	cli, err := messagix.NewClient(types.Instagram, nil, debug.NewLogger(), "")
	if err != nil {
		log.Fatal(err)
	}
	parser := &messagix.ModuleParser{}
	testData, _ := os.ReadFile("test_files/res.html")
	parser.SetTestData(testData)
	parser.SetClientInstance(cli)
	parser.Load("")
	//parser.Load("https://www.instagram.com/accounts/login/")
}
