// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list that describes one or more specified images, if the image
// identifiers are provided. Otherwise, all images in the account are described.
func (c *Client) DescribeWorkspaceImages(ctx context.Context, params *DescribeWorkspaceImagesInput, optFns ...func(*Options)) (*DescribeWorkspaceImagesOutput, error) {
	if params == nil {
		params = &DescribeWorkspaceImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkspaceImages", params, optFns, c.addOperationDescribeWorkspaceImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkspaceImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspaceImagesInput struct {

	// The identifier of the image.
	ImageIds []string

	// The type (owned or shared) of the image.
	ImageType types.ImageType

	// The maximum number of items to return.
	MaxResults *int32

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeWorkspaceImagesOutput struct {

	// Information about the images.
	Images []types.WorkspaceImage

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWorkspaceImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspaceImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspaceImages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaceImages(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeWorkspaceImages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeWorkspaceImages",
	}
}
