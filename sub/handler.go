package function

import (
	"fmt"
	"handler/function/util"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("This is your hostname: %s", util.GetHostname())
}
