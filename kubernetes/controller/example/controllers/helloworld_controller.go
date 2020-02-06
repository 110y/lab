/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	labolithv1 "github.com/110y/lab/api/v1"
)

// HelloWorldReconciler reconciles a HelloWorld object
type HelloWorldReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=labolith.labolith.com,resources=helloworlds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=labolith.labolith.com,resources=helloworlds/status,verbs=get;update;patch

func (r *HelloWorldReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	l := r.Log.WithValues("helloworld", req.NamespacedName)

	// your logic here

	var helloworld labolithv1.HelloWorld
	if err := r.Get(ctx, req.NamespacedName, &helloworld); err != nil {
		l.Error(err, "unable to fetch HelloWorld")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	l.Info(fmt.Sprintf("HELLOWORLD: %s", helloworld.Spec.Foo))

	return ctrl.Result{}, nil
}

func (r *HelloWorldReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&labolithv1.HelloWorld{}).
		Complete(r)
}
