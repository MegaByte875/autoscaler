// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroups

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// The request includes one or more parameters that violate validation rules.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	//
	// The caller isn't authorized to make the request. Check permissions.
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeInternalServerErrorException for service response error code
	// "InternalServerErrorException".
	//
	// An internal error occurred while processing the request. Try again later.
	ErrCodeInternalServerErrorException = "InternalServerErrorException"

	// ErrCodeMethodNotAllowedException for service response error code
	// "MethodNotAllowedException".
	//
	// The request uses an HTTP method that isn't allowed for the specified resource.
	ErrCodeMethodNotAllowedException = "MethodNotAllowedException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// One or more of the specified resources don't exist.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// You've exceeded throttling limits by making too many requests in a period
	// of time.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// The request was rejected because it doesn't have valid credentials for the
	// target resource.
	ErrCodeUnauthorizedException = "UnauthorizedException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":          newErrorBadRequestException,
	"ForbiddenException":           newErrorForbiddenException,
	"InternalServerErrorException": newErrorInternalServerErrorException,
	"MethodNotAllowedException":    newErrorMethodNotAllowedException,
	"NotFoundException":            newErrorNotFoundException,
	"TooManyRequestsException":     newErrorTooManyRequestsException,
	"UnauthorizedException":        newErrorUnauthorizedException,
}
