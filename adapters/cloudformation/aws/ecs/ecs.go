package ecs

import (
	"github.com/aquasecurity/defsec/provider/aws/ecs"
	"github.com/aquasecurity/trivy-config-parsers/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) (result ecs.ECS) {

	result.Clusters = getClusters(cfFile)
	result.TaskDefinitions = getTaskDefinitions(cfFile)
	return result

}