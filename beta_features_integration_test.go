// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

func TestBetaFeaturesService_Integration_List(t *testing.T) {
	_, res, err := integrationClient.BetaFeatures.List(context.TODO(), integrationUserId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}
