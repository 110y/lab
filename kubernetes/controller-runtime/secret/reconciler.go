package main

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ reconcile.Reconciler = &helloReconciler{}

type helloReconciler struct {
	client.Client
	log logr.Logger
}

func (r *helloReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	key := client.ObjectKey{
		Name:      "foo",
		Namespace: "secret",
	}

	var se *corev1.Secret

	r.Get(context.Background(), key, se)
	r.log.Info("RECONCILE")
	return reconcile.Result{}, nil
}
