// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a source identifier to an existing event notification subscription.
func (c *Client) AddSourceIdentifierToSubscription(ctx context.Context, params *AddSourceIdentifierToSubscriptionInput, optFns ...func(*Options)) (*AddSourceIdentifierToSubscriptionOutput, error) {
	if params == nil {
		params = &AddSourceIdentifierToSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddSourceIdentifierToSubscription", params, optFns, c.addOperationAddSourceIdentifierToSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddSourceIdentifierToSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddSourceIdentifierToSubscriptionInput struct {

	// The identifier of the event source to be added. Constraints:
	//
	// * If the source
	// type is a DB instance, then a DBInstanceIdentifier must be supplied.
	//
	// * If the
	// source type is a DB security group, a DBSecurityGroupName must be supplied.
	//
	// *
	// If the source type is a DB parameter group, a DBParameterGroupName must be
	// supplied.
	//
	// * If the source type is a DB snapshot, a DBSnapshotIdentifier must be
	// supplied.
	//
	// This member is required.
	SourceIdentifier *string

	// The name of the event notification subscription you want to add a source
	// identifier to.
	//
	// This member is required.
	SubscriptionName *string

	noSmithyDocumentSerde
}

type AddSourceIdentifierToSubscriptionOutput struct {

	// Contains the results of a successful invocation of the
	// DescribeEventSubscriptions action.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddSourceIdentifierToSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddSourceIdentifierToSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddSourceIdentifierToSubscription{}, middleware.After)
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
	if err = addOpAddSourceIdentifierToSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddSourceIdentifierToSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddSourceIdentifierToSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "AddSourceIdentifierToSubscription",
	}
}
