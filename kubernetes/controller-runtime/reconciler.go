package main

import (
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ reconcile.Reconciler = &helloReconciler{}

type helloReconciler struct {
	log logr.Logger
}

func (r *helloReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	r.log.Info("RECONCILE")
	return reconcile.Result{}, nil
}
