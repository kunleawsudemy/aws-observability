// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package plog

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlplogs "go.opentelemetry.io/collector/pdata/internal/data/protogen/logs/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceLogs is a collection of logs from a Resource.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceLogs function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceLogs struct {
	orig  *otlplogs.ResourceLogs
	state *internal.State
}

func newResourceLogs(orig *otlplogs.ResourceLogs, state *internal.State) ResourceLogs {
	return ResourceLogs{orig: orig, state: state}
}

// NewResourceLogs creates a new empty ResourceLogs.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewResourceLogs() ResourceLogs {
	state := internal.StateMutable
	return newResourceLogs(&otlplogs.ResourceLogs{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ResourceLogs) MoveTo(dest ResourceLogs) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlplogs.ResourceLogs{}
}

// Resource returns the resource associated with this ResourceLogs.
func (ms ResourceLogs) Resource() pcommon.Resource {
	return pcommon.Resource(internal.NewResource(&ms.orig.Resource, ms.state))
}

// SchemaUrl returns the schemaurl associated with this ResourceLogs.
func (ms ResourceLogs) SchemaUrl() string {
	return ms.orig.SchemaUrl
}

// SetSchemaUrl replaces the schemaurl associated with this ResourceLogs.
func (ms ResourceLogs) SetSchemaUrl(v string) {
	ms.state.AssertMutable()
	ms.orig.SchemaUrl = v
}

// ScopeLogs returns the ScopeLogs associated with this ResourceLogs.
func (ms ResourceLogs) ScopeLogs() ScopeLogsSlice {
	return newScopeLogsSlice(&ms.orig.ScopeLogs, ms.state)
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ResourceLogs) CopyTo(dest ResourceLogs) {
	dest.state.AssertMutable()
	ms.Resource().CopyTo(dest.Resource())
	dest.SetSchemaUrl(ms.SchemaUrl())
	ms.ScopeLogs().CopyTo(dest.ScopeLogs())
}