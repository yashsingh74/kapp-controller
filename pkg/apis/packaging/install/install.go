// Copyright 2024 The Carvel Authors.
// SPDX-License-Identifier: Apache-2.0

package install

import (
	"carvel.dev/kapp-controller/pkg/apis/packaging/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

func Install(scheme *runtime.Scheme) {
	v1alpha1.AddToScheme(scheme)
}
