// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyModelsProjectPlanStatus() models.ProjectPlanStatus {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.ProjectPlanStatus))(nil)).Elem()))
	var nullValue models.ProjectPlanStatus
	return nullValue
}

func EqModelsProjectPlanStatus(value models.ProjectPlanStatus) models.ProjectPlanStatus {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.ProjectPlanStatus
	return nullValue
}

func NotEqModelsProjectPlanStatus(value models.ProjectPlanStatus) models.ProjectPlanStatus {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue models.ProjectPlanStatus
	return nullValue
}

func ModelsProjectPlanStatusThat(matcher pegomock.ArgumentMatcher) models.ProjectPlanStatus {
	pegomock.RegisterMatcher(matcher)
	var nullValue models.ProjectPlanStatus
	return nullValue
}