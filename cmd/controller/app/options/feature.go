/*
Copyright 2021 KubeSphere Authors

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

package options

import (
	"strings"

	"github.com/spf13/pflag"
	cliflag "k8s.io/component-base/cli/flag"
	"kubesphere.io/devops/pkg/utils/reflectutils"
)

// FeatureOptions provide some feature options, such as specifying the controller to be enabled.
type FeatureOptions struct {
	EnabledControllers map[string]bool
}

// NewFeatureOptions provide default options
func NewFeatureOptions() *FeatureOptions {
	return &FeatureOptions{
		EnabledControllers: map[string]bool{
			"s2ibinary":        true,
			"s2irun":           true,
			"pipeline":         true,
			"devopsprojects":   true,
			"devopscredential": true,
			"jenkinsconfig":    true,
		},
	}
}

// Validate checks validation of FeatureOptions.
func (f *FeatureOptions) Validate() []error {
	return []error{}
}

// ApplyTo fills up FeatureOptions config with options
func (f *FeatureOptions) ApplyTo(options *FeatureOptions) {
	reflectutils.Override(options, f)
}

// AddFlags adds flags related to FeatureOptions for controller manager to the feature FlagSet.
func (f *FeatureOptions) AddFlags(fs *pflag.FlagSet, c *FeatureOptions) {
	fs.Var(cliflag.NewMapStringBool(&f.EnabledControllers), "enabled-controllers", "A set of key=value pairs that describe feature options for controllers. "+
		"Options are:\n"+strings.Join(c.knownControllers(), "\n"))
}

func (f *FeatureOptions) knownControllers() []string {
	controllers := make([]string, 0, len(f.EnabledControllers))
	for name := range f.EnabledControllers {
		controllers = append(controllers, name)
	}
	return controllers
}
