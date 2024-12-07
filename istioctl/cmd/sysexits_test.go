// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"testing"

	"istio.io/istio/123/istioctl/pkg/analyze"
	"istio.io/istio/123/istioctl/pkg/util"
)

var KnownErrorCode = map[error]int{
	errors.New("unknown command"):                                  ExitIncorrectUsage,
	errors.New("unexpected error"):                                 ExitUnknownError,
	util.CommandParseError{Err: errors.New("command parse error")}: ExitIncorrectUsage,
	analyze.FileParseError{}:                                       ExitDataError,
	analyze.AnalyzerFoundIssuesError{}:                             ExitAnalyzerFoundIssues,
}

func TestKnownExitStrings(t *testing.T) {
	for err, wantCode := range KnownErrorCode {
		if code := GetExitCode(err); code != wantCode {
			t.Errorf("For %v want %v, but is %v", err, wantCode, code)
		}
	}
}
