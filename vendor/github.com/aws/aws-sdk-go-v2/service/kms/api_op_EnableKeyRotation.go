// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables automatic rotation of the key material
// (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) of the
// specified symmetric encryption KMS key. When you enable automatic rotation of
// acustomer managed KMS key
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk),
// KMS rotates the key material of the KMS key one year (approximately 365 days)
// from the enable date and every year thereafter. You can monitor rotation of the
// key material for your KMS keys in CloudTrail and Amazon CloudWatch. To disable
// rotation of the key material in a customer managed KMS key, use the
// DisableKeyRotation operation. Automatic key rotation is supported only on
// symmetric encryption KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#symmetric-cmks).
// You cannot enable or disable automatic rotation of asymmetric KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html),
// HMAC KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html),
// KMS keys with imported key material
// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html), or
// KMS keys in a custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
// The key rotation status of these KMS keys is always false. To enable or disable
// automatic rotation of a set of related multi-Region keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-manage.html#multi-region-rotate),
// set the property on the primary key. You cannot enable or disable automatic
// rotation Amazon Web Services managed KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
// KMS always rotates the key material of Amazon Web Services managed keys every
// year. Rotation of Amazon Web Services owned KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-cmk)
// varies. In May 2022, KMS changed the rotation schedule for Amazon Web Services
// managed keys from every three years (approximately 1,095 days) to every year
// (approximately 365 days). New Amazon Web Services managed keys are automatically
// rotated one year after they are created, and approximately every year
// thereafter. Existing Amazon Web Services managed keys are automatically rotated
// one year after their most recent rotation, and every year thereafter. The KMS
// key that you use for this operation must be in a compatible key state. For
// details, see Key states of KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on a KMS key in a different Amazon Web Services account.
// Required permissions: kms:EnableKeyRotation
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations:
//
// * DisableKeyRotation
//
// * GetKeyRotationStatus
func (c *Client) EnableKeyRotation(ctx context.Context, params *EnableKeyRotationInput, optFns ...func(*Options)) (*EnableKeyRotationOutput, error) {
	if params == nil {
		params = &EnableKeyRotationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableKeyRotation", params, optFns, c.addOperationEnableKeyRotationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableKeyRotationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableKeyRotationInput struct {

	// Identifies a symmetric encryption KMS key. You cannot enable or disable
	// automatic rotation of asymmetric KMS keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html),
	// HMAC KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html),
	// KMS keys with imported key material
	// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html), or
	// KMS keys in a custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
	// The key rotation status of these KMS keys is always false. To enable or disable
	// automatic rotation of a set of related multi-Region keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-manage.html#multi-region-rotate),
	// set the property on the primary key. Specify the key ID or key ARN of the KMS
	// key. For example:
	//
	// * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	noSmithyDocumentSerde
}

type EnableKeyRotationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableKeyRotationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableKeyRotation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableKeyRotation{}, middleware.After)
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
	if err = addOpEnableKeyRotationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableKeyRotation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableKeyRotation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "EnableKeyRotation",
	}
}
