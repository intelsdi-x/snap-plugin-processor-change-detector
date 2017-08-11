//
// +build small

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
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestChangeDetectorProcessor(t *testing.T) {
	proc := NewProcessor()
	Convey("Create change detector processor", t, func() {
		Convey("So proc should not be nil", func() {
			So(proc, ShouldNotBeNil)
		})
		Convey("So proc should be of type ChangeDetectorProcessor", func() {
			So(proc, ShouldHaveSameTypeAs, &ChangeDetectorProcessor{})
		})
		Convey("proc.GetConfigPolicy should return a config policy", func() {
			configPolicy, _ := proc.GetConfigPolicy()
			Convey("So config policy should be a plugin.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, plugin.ConfigPolicy{})
			})
		})
	})

	Convey("Test Change Detector Processor", t, func() {
		Convey("Process metrics with invalid config", func() {
			config := plugin.Config{
				"sometext": "sometext",
			}
			metrics := []plugin.Metric{
				{
					Namespace: plugin.NewNamespace("foo"),
					Data:      456,
					Tags:      map[string]string{"test1": "fooval", "test2": "barval"},
				},
			}
			mts, err := proc.Process(metrics, config)
			So(mts, ShouldBeNil)
			So(fmt.Sprint(err), ShouldEqual, "config item not found")
		})

		Convey("Process metrics with empty rules", func() {
			config := plugin.Config{
				"rules": "",
			}
			metrics := []plugin.Metric{
				{
					Namespace: plugin.NewNamespace("foo"),
					Data:      456,
					Tags:      map[string]string{"test1": "fooval", "test2": "barval"},
				},
			}
			mts, err := proc.Process(metrics, config)
			So(mts, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(mts, ShouldBeEmpty)
		})
	})
}
