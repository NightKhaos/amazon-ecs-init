// +build !suse,!ubuntu

// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package docker

const (
	// dockerClientAPIVersion specifies the minimum docker client API version
	// required by ECS Init
	// Version 1.25 is required for setting Init to true when constructing
	// the HostConfig for creating the ECS Agent to enable the Task
	// Networking with ENI capability
	dockerClientAPIVersion = "1.25"
)
