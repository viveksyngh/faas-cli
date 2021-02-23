// Copyright (c) Alex Ellis 2017. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package proxy

import (
	"context"
	"net/http"

	"testing"

	"regexp"

	"github.com/openfaas/faas-cli/test"
)

func Test_DeleteFunction(t *testing.T) {
	s := test.MockHttpServerStatus(t, http.StatusOK)
	defer s.Close()

	cliAuth := NewTestAuth(nil)
	proxyClient, _ := NewClient(cliAuth, s.URL, nil, &defaultCommandTimeout)

	err := proxyClient.DeleteFunction(context.Background(), "function-to-delete", "")

	if err != nil {
		t.Fatalf("Got error: %s", err.Error())
	}
}

func Test_DeleteFunction_404(t *testing.T) {
	s := test.MockHttpServerStatus(t, http.StatusNotFound)
	defer s.Close()

	cliAuth := NewTestAuth(nil)
	proxyClient, _ := NewClient(cliAuth, s.URL, nil, &defaultCommandTimeout)

	err := proxyClient.DeleteFunction(context.Background(), "function-to-delete", "")

	r := regexp.MustCompile(`(?m:No existing function to remove)`)
	if !r.MatchString(err.Error()) {
		t.Fatalf("Want: %s, got: %s", "No existing function to remove", err.Error())
	}
}

func Test_DeleteFunction_Not2xxAnd404(t *testing.T) {
	s := test.MockHttpServerStatus(t, http.StatusInternalServerError)
	defer s.Close()

	cliAuth := NewTestAuth(nil)
	proxyClient, _ := NewClient(cliAuth, s.URL, nil, &defaultCommandTimeout)

	err := proxyClient.DeleteFunction(context.Background(), "function-to-delete", "")

	r := regexp.MustCompile(`(?m:Server returned unexpected status code)`)
	if !r.MatchString(err.Error()) {
		t.Fatalf("Output not matched: %s", err.Error())
	}
}
