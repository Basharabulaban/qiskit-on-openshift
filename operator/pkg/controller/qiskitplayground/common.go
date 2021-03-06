package qiskitplayground

import (
	"context"

	dobtechv1 "github.com/jdob/qiskit-operator/pkg/apis/dobtech/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	routev1 "github.com/openshift/api/route/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func (r *ReconcileQiskitPlayground) ensureDeployment(request reconcile.Request,
	instance *dobtechv1.QiskitPlayground,
	dep *appsv1.Deployment,
) (*reconcile.Result, error) {

	// See if deployment already exists and create if it doesn't
	found := &appsv1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
		Name:      dep.Name,
		Namespace: instance.Namespace,
	}, found)
	if err != nil && errors.IsNotFound(err) {

		// Create the deployment
		log.Info("Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
		err = r.client.Create(context.TODO(), dep)

		if err != nil {
			// Deployment failed
			log.Error(err, "Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
			return &reconcile.Result{}, err
		} else {
			// Deployment was successful
			return nil, nil
		}
	} else if err != nil {
		// Error that isn't due to the deployment not existing
		log.Error(err, "Failed to get Deployment")
		return &reconcile.Result{}, err
	}

	return nil, nil
}

func (r *ReconcileQiskitPlayground) ensureService(request reconcile.Request,
	instance *dobtechv1.QiskitPlayground,
	s *corev1.Service,
) (*reconcile.Result, error) {
	found := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
		Name:      s.Name,
		Namespace: instance.Namespace,
	}, found)
	if err != nil && errors.IsNotFound(err) {

		// Create the service
		log.Info("Creating a new Service", "Service.Namespace", s.Namespace, "Service.Name", s.Name)
		err = r.client.Create(context.TODO(), s)

		if err != nil {
			// Creation failed
			log.Error(err, "Failed to create new Service", "Service.Namespace", s.Namespace, "Service.Name", s.Name)
			return &reconcile.Result{}, err
		} else {
			// Creation was successful
			return nil, nil
		}
	} else if err != nil {
		// Error that isn't due to the service not existing
		log.Error(err, "Failed to get Service")
		return &reconcile.Result{}, err
	}

	return nil, nil
}

func (r *ReconcileQiskitPlayground) ensureRoute(request reconcile.Request,
	instance *dobtechv1.QiskitPlayground,
	rte *routev1.Route,
) (*reconcile.Result, error) {
	found := &routev1.Route{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
		Name: 		rte.Name,
		Namespace: 	instance.Namespace,
	}, found)
	if err != nil && errors.IsNotFound(err) {

		// Create the route
		log.Info("Creating a new Route", "Route.Namespace", rte.Namespace, "Route.Name", rte.Name)
		err = r.client.Create(context.TODO(), rte)

		if err != nil {
			// Creation failed
			log.Error(err, "Failed to create new Route")
			return &reconcile.Result{}, err
		} else {
			// Creation successful
			return nil, nil
		}
	} else if err != nil {
		// Error that isn't the route not existing
		log.Error(err, "Failed to create Route")
		return &reconcile.Result{}, err
	}

	return nil, nil
}

func labels(q *dobtechv1.QiskitPlayground) map[string]string {
	return map[string]string{
		"app":                "playground",
		"qiskitplayground_cr": q.Name,
	}
}
