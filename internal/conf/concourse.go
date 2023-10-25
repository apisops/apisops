package conf

import (
	"github.com/apisops/apisops/internal/mongo"
	"github.com/apisops/apisops/internal/websvr"
)

// ConcourseConf is the configuration object of Concourse
type ConcourseConf struct {
	WebSvr websvr.Configuration `yaml:"webSvr"`
	Mongo  mongo.Configuration  `yaml:"mongo"`
}
