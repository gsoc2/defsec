package mq

import (
	"github.com/aquasecurity/defsec/provider/aws/mq"
	"github.com/aquasecurity/trivy-config-parsers/cloudformation/parser"
)

func getBrokers(ctx parser.FileContext) (brokers []mq.Broker) {
	for _, r := range ctx.GetResourceByType("AWS::AmazonMQ::Broker") {

		broker := mq.Broker{
			Metadata:     r.Metadata(),
			PublicAccess: r.GetBoolProperty("PubliclyAccessible"),
			Logging: mq.Logging{
				General: r.GetBoolProperty("Logs.General"),
				Audit:   r.GetBoolProperty("Logs.Audit"),
			},
		}

		brokers = append(brokers, broker)
	}
	return brokers
}