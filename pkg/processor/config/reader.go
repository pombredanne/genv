/*
Copyright 2017 The Nuclio Authors.

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

package processorconfig

import (
	"io"
	"io/ioutil"

	"github.com/pmker/genv/pkg/errors"
	"github.com/pmker/genv/pkg/processor"

	"github.com/ghodss/yaml"
)

type Reader struct {
}

func NewReader() (*Reader, error) {
	return &Reader{}, nil
}

func (r *Reader) Read(reader io.Reader, processorConfiguration *processor.Configuration) error {
	bodyBytes, err := ioutil.ReadAll(reader)

	if err != nil {
		return errors.Wrap(err, "Failed to read processor configuration")
	}

	if err := yaml.Unmarshal(bodyBytes, processorConfiguration); err != nil {
		return errors.Wrap(err, "Failed to write configuration")
	}

	return nil
}
