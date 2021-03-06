//Package ssm provides gucumber integration tests support.
package ssm

import (
	"github.com/aws/aws-sdk-go/internal/features/shared"
	"github.com/aws/aws-sdk-go/service/ssm"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@ssm", func() {
		World["client"] = ssm.New(nil)
	})
}
