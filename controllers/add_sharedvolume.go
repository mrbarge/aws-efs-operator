package controllers

import (
	"openshift/aws-efs-operator/controllers"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, controllers.Add)
}
