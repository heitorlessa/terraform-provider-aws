// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fsx

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeActiveDirectoryError for service response error code
	// "ActiveDirectoryError".
	//
	// An Active Directory error.
	ErrCodeActiveDirectoryError = "ActiveDirectoryError"

	// ErrCodeBackupInProgress for service response error code
	// "BackupInProgress".
	//
	// Another backup is already under way. Wait for completion before initiating
	// additional backups of this file system.
	ErrCodeBackupInProgress = "BackupInProgress"

	// ErrCodeBackupNotFound for service response error code
	// "BackupNotFound".
	//
	// No Amazon FSx backups were found based upon the supplied parameters.
	ErrCodeBackupNotFound = "BackupNotFound"

	// ErrCodeBackupRestoring for service response error code
	// "BackupRestoring".
	//
	// You can't delete a backup while it's being used to restore a file system.
	ErrCodeBackupRestoring = "BackupRestoring"

	// ErrCodeBadRequest for service response error code
	// "BadRequest".
	//
	// A generic error indicating a failure with a client request.
	ErrCodeBadRequest = "BadRequest"

	// ErrCodeDataRepositoryTaskEnded for service response error code
	// "DataRepositoryTaskEnded".
	//
	// The data repository task could not be canceled because the task has already
	// ended.
	ErrCodeDataRepositoryTaskEnded = "DataRepositoryTaskEnded"

	// ErrCodeDataRepositoryTaskExecuting for service response error code
	// "DataRepositoryTaskExecuting".
	//
	// An existing data repository task is currently executing on the file system.
	// Wait until the existing task has completed, then create the new task.
	ErrCodeDataRepositoryTaskExecuting = "DataRepositoryTaskExecuting"

	// ErrCodeDataRepositoryTaskNotFound for service response error code
	// "DataRepositoryTaskNotFound".
	//
	// The data repository task or tasks you specified could not be found.
	ErrCodeDataRepositoryTaskNotFound = "DataRepositoryTaskNotFound"

	// ErrCodeFileSystemNotFound for service response error code
	// "FileSystemNotFound".
	//
	// No Amazon FSx file systems were found based upon supplied parameters.
	ErrCodeFileSystemNotFound = "FileSystemNotFound"

	// ErrCodeIncompatibleParameterError for service response error code
	// "IncompatibleParameterError".
	//
	// The error returned when a second request is received with the same client
	// request token but different parameters settings. A client request token should
	// always uniquely identify a single request.
	ErrCodeIncompatibleParameterError = "IncompatibleParameterError"

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// A generic error indicating a server-side failure.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeInvalidExportPath for service response error code
	// "InvalidExportPath".
	//
	// The path provided for data repository export isn't valid.
	ErrCodeInvalidExportPath = "InvalidExportPath"

	// ErrCodeInvalidImportPath for service response error code
	// "InvalidImportPath".
	//
	// The path provided for data repository import isn't valid.
	ErrCodeInvalidImportPath = "InvalidImportPath"

	// ErrCodeInvalidNetworkSettings for service response error code
	// "InvalidNetworkSettings".
	//
	// One or more network settings specified in the request are invalid. InvalidVpcId
	// means that the ID passed for the virtual private cloud (VPC) is invalid.
	// InvalidSubnetIds returns the list of IDs for subnets that are either invalid
	// or not part of the VPC specified. InvalidSecurityGroupIds returns the list
	// of IDs for security groups that are either invalid or not part of the VPC
	// specified.
	ErrCodeInvalidNetworkSettings = "InvalidNetworkSettings"

	// ErrCodeMissingFileSystemConfiguration for service response error code
	// "MissingFileSystemConfiguration".
	//
	// File system configuration is required for this operation.
	ErrCodeMissingFileSystemConfiguration = "MissingFileSystemConfiguration"

	// ErrCodeNotServiceResourceError for service response error code
	// "NotServiceResourceError".
	//
	// The resource specified for the tagging operation is not a resource type owned
	// by Amazon FSx. Use the API of the relevant service to perform the operation.
	ErrCodeNotServiceResourceError = "NotServiceResourceError"

	// ErrCodeResourceDoesNotSupportTagging for service response error code
	// "ResourceDoesNotSupportTagging".
	//
	// The resource specified does not support tagging.
	ErrCodeResourceDoesNotSupportTagging = "ResourceDoesNotSupportTagging"

	// ErrCodeResourceNotFound for service response error code
	// "ResourceNotFound".
	//
	// The resource specified by the Amazon Resource Name (ARN) can't be found.
	ErrCodeResourceNotFound = "ResourceNotFound"

	// ErrCodeServiceLimitExceeded for service response error code
	// "ServiceLimitExceeded".
	//
	// An error indicating that a particular service limit was exceeded. You can
	// increase some service limits by contacting AWS Support.
	ErrCodeServiceLimitExceeded = "ServiceLimitExceeded"

	// ErrCodeUnsupportedOperation for service response error code
	// "UnsupportedOperation".
	//
	// The requested operation is not supported for this resource or API.
	ErrCodeUnsupportedOperation = "UnsupportedOperation"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ActiveDirectoryError":           newErrorActiveDirectoryError,
	"BackupInProgress":               newErrorBackupInProgress,
	"BackupNotFound":                 newErrorBackupNotFound,
	"BackupRestoring":                newErrorBackupRestoring,
	"BadRequest":                     newErrorBadRequest,
	"DataRepositoryTaskEnded":        newErrorDataRepositoryTaskEnded,
	"DataRepositoryTaskExecuting":    newErrorDataRepositoryTaskExecuting,
	"DataRepositoryTaskNotFound":     newErrorDataRepositoryTaskNotFound,
	"FileSystemNotFound":             newErrorFileSystemNotFound,
	"IncompatibleParameterError":     newErrorIncompatibleParameterError,
	"InternalServerError":            newErrorInternalServerError,
	"InvalidExportPath":              newErrorInvalidExportPath,
	"InvalidImportPath":              newErrorInvalidImportPath,
	"InvalidNetworkSettings":         newErrorInvalidNetworkSettings,
	"MissingFileSystemConfiguration": newErrorMissingFileSystemConfiguration,
	"NotServiceResourceError":        newErrorNotServiceResourceError,
	"ResourceDoesNotSupportTagging":  newErrorResourceDoesNotSupportTagging,
	"ResourceNotFound":               newErrorResourceNotFound,
	"ServiceLimitExceeded":           newErrorServiceLimitExceeded,
	"UnsupportedOperation":           newErrorUnsupportedOperation,
}
