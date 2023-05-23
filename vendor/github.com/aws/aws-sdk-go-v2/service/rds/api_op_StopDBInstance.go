// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops an Amazon RDS DB instance. When you stop a DB instance, Amazon RDS retains
// the DB instance's metadata, including its endpoint, DB parameter group, and
// option group membership. Amazon RDS also retains the transaction logs so you can
// do a point-in-time restore if necessary. For more information, see  Stopping an
// Amazon RDS DB Instance Temporarily
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_StopInstance.html)
// in the Amazon RDS User Guide. This command doesn't apply to RDS Custom, Aurora
// MySQL, and Aurora PostgreSQL. For Aurora clusters, use StopDBCluster instead.
func (c *Client) StopDBInstance(ctx context.Context, params *StopDBInstanceInput, optFns ...func(*Options)) (*StopDBInstanceOutput, error) {
	if params == nil {
		params = &StopDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopDBInstance", params, optFns, c.addOperationStopDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDBInstanceInput struct {

	// The user-supplied instance identifier.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The user-supplied instance identifier of the DB Snapshot created immediately
	// before the DB instance is stopped.
	DBSnapshotIdentifier *string

	noSmithyDocumentSerde
}

type StopDBInstanceOutput struct {

	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the operations CreateDBInstance,
	// CreateDBInstanceReadReplica, DeleteDBInstance, DescribeDBInstances,
	// ModifyDBInstance, PromoteReadReplica, RebootDBInstance,
	// RestoreDBInstanceFromDBSnapshot, RestoreDBInstanceFromS3,
	// RestoreDBInstanceToPointInTime, StartDBInstance, and StopDBInstance.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStopDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStopDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStopDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopDBInstance(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opStopDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StopDBInstance",
	}
}
