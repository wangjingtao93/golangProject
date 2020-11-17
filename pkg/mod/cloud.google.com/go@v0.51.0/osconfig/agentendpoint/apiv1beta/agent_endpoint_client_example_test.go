// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package agentendpoint_test

import (
	"context"

	agentendpoint "cloud.google.com/go/osconfig/agentendpoint/apiv1beta"
	agentendpointpb "google.golang.org/genproto/googleapis/cloud/osconfig/agentendpoint/v1beta"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := agentendpoint.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_StartNextTask() {
	// import agentendpointpb "google.golang.org/genproto/googleapis/cloud/osconfig/agentendpoint/v1beta"

	ctx := context.Background()
	c, err := agentendpoint.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &agentendpointpb.StartNextTaskRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.StartNextTask(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ReportTaskProgress() {
	// import agentendpointpb "google.golang.org/genproto/googleapis/cloud/osconfig/agentendpoint/v1beta"

	ctx := context.Background()
	c, err := agentendpoint.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &agentendpointpb.ReportTaskProgressRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ReportTaskProgress(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ReportTaskComplete() {
	// import agentendpointpb "google.golang.org/genproto/googleapis/cloud/osconfig/agentendpoint/v1beta"

	ctx := context.Background()
	c, err := agentendpoint.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &agentendpointpb.ReportTaskCompleteRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ReportTaskComplete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_LookupEffectiveGuestPolicy() {
	// import agentendpointpb "google.golang.org/genproto/googleapis/cloud/osconfig/agentendpoint/v1beta"

	ctx := context.Background()
	c, err := agentendpoint.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &agentendpointpb.LookupEffectiveGuestPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.LookupEffectiveGuestPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
