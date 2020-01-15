package main

import (
	"os"

	v1 "k8s.io/api/core/v1"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("example-controller")

func main() {
	logf.SetLogger(zap.Logger(false))
	mLog := log.WithName("main.go")

	cfg, err := config.GetConfig()
	if err != nil {
		mLog.Error(err, "failed to get config")
		os.Exit(1)
	}

	mgr, err := manager.New(cfg, manager.Options{})
	if err != nil {
		mLog.Error(err, "failed to create manager")
		os.Exit(1)
	}

	ctrl, err := controller.New("hello-controller", mgr, controller.Options{
		Reconciler: &helloReconciler{
			log: log.WithName("hello-reconciler"),
		},
	})
	if err != nil {
		mLog.Error(err, "failed to create controller")
		os.Exit(1)
	}

	mLog.Info("start to watch")
	if err := ctrl.Watch(&source.Kind{Type: &v1.Pod{}}, &handler.EnqueueRequestForObject{}); err != nil {
		mLog.Error(err, "failed to watch pods")
		os.Exit(1)
	}

	mLog.Info("start manager")
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		mLog.Error(err, "failed to start manager")
		os.Exit(1)
	}
}
