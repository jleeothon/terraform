// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package terraform

import "github.com/hashicorp/terraform/internal/tfdiags"

// NodeEvalableProvider represents a provider during an "eval" walk.
// This special provider node type just initializes a provider and
// fetches its schema, without configuring it or otherwise interacting
// with it.
type NodeEvalableProvider struct {
	*NodeAbstractProvider
}

var _ GraphNodeExecutable = (*NodeEvalableProvider)(nil)

// GraphNodeExecutable
func (n *NodeEvalableProvider) Execute(ctx EvalContext, op walkOperation) (diags tfdiags.Diagnostics) {
	return nil
}
