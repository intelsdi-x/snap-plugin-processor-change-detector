/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2017 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package changedetector

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

const (
	// Name of plugin
	Name = "change-detector"
	// Version of plugin
	Version = 1
)

// ChangeDetectorProcessor implements plugin interface
type ChangeDetectorProcessor struct {
	values map[string]interface{}
}

// NewProcessor returns a pointer to plugin instance
func NewProcessor() *ChangeDetectorProcessor {
	return &ChangeDetectorProcessor{values: make(map[string]interface{}, 0)}
}

// Process method to detect changes
func (p *ChangeDetectorProcessor) Process(mts []plugin.Metric, cfg plugin.Config) ([]plugin.Metric, error) {
	cfgItem, err := cfg.GetString("rules")
	if err != nil {
		return nil, err
	}
	rules := strings.Split(cfgItem, "|")

	metrics := []plugin.Metric{}
	for _, m := range mts {
		ns := "/" + strings.Join(m.Namespace.Strings(), "/")

		for _, rule := range rules {
			matched, err := regexp.MatchString(rule, ns)
			if err != nil {
				return nil, err
			}
			if matched {
				prevVal, ok := p.values[ns]
				if ok {
					if prevVal != m.Data {
						m.Tags["previous_value"] = fmt.Sprintf("%v", prevVal)
						metrics = append(metrics, m)
					}
				}
				p.values[ns] = m.Data
			}
		}
	}
	return metrics, nil
}

// GetConfigPolicy returns plugin config
func (p ChangeDetectorProcessor) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{""}, "rules", true)
	return *policy, nil
}
