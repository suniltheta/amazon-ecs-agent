// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package version

import (
	"fmt"

	"github.com/aws/amazon-ecs-agent/agent/sighandlers/exitcodes"
)

// PrintVersions prints the version information on stdout as a multi-line
// string. The output will look similar to the following:
//
//	Amazon ECS Agent:
//	    Version: 0.0.3
//	    Commit: 4bdc7fc
func PrintVersion() int {
	cleanliness := ""
	if GitDirty {
		cleanliness = "\tDirty: true\n"
	}

	fmt.Printf(`Amazon ECS Agent:
	Version: %v
	Commit: %v
%v`, Version, GitShortHash, cleanliness)

	return exitcodes.ExitSuccess
}

// String produces a human-readable string showing the agent version.
func String() string {
	ret := "Amazon ECS Agent - v" + Version + " ("
	if GitDirty {
		ret += "*"
	}
	return ret + GitShortHash + ")"
}

func GitHashString() string {
	if GitDirty {
		return "*" + GitShortHash
	}
	return GitShortHash
}
