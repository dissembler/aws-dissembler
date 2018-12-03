package controller

import (
	"github.com/temp/aws-dissembler/pkg/controller/s3"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, s3.Add)
}
