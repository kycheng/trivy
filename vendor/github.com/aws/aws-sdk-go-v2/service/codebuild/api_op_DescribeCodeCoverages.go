// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves one or more code coverage reports.
func (c *Client) DescribeCodeCoverages(ctx context.Context, params *DescribeCodeCoveragesInput, optFns ...func(*Options)) (*DescribeCodeCoveragesOutput, error) {
	if params == nil {
		params = &DescribeCodeCoveragesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCodeCoverages", params, optFns, c.addOperationDescribeCodeCoveragesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCodeCoveragesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCodeCoveragesInput struct {

	// The ARN of the report for which test cases are returned.
	//
	// This member is required.
	ReportArn *string

	// The maximum line coverage percentage to report.
	MaxLineCoveragePercentage *float64

	// The maximum number of results to return.
	MaxResults *int32

	// The minimum line coverage percentage to report.
	MinLineCoveragePercentage *float64

	// The nextToken value returned from a previous call to DescribeCodeCoverages. This
	// specifies the next item to return. To return the beginning of the list, exclude
	// this parameter.
	NextToken *string

	// Specifies how the results are sorted. Possible values are: FILE_PATH The results
	// are sorted by file path. LINE_COVERAGE_PERCENTAGE The results are sorted by the
	// percentage of lines that are covered.
	SortBy types.ReportCodeCoverageSortByType

	// Specifies if the results are sorted in ascending or descending order.
	SortOrder types.SortOrderType

	noSmithyDocumentSerde
}

type DescribeCodeCoveragesOutput struct {

	// An array of CodeCoverage objects that contain the results.
	CodeCoverages []types.CodeCoverage

	// If there are more items to return, this contains a token that is passed to a
	// subsequent call to DescribeCodeCoverages to retrieve the next set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCodeCoveragesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCodeCoverages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCodeCoverages{}, middleware.After)
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
	if err = addOpDescribeCodeCoveragesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCodeCoverages(options.Region), middleware.Before); err != nil {
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

// DescribeCodeCoveragesAPIClient is a client that implements the
// DescribeCodeCoverages operation.
type DescribeCodeCoveragesAPIClient interface {
	DescribeCodeCoverages(context.Context, *DescribeCodeCoveragesInput, ...func(*Options)) (*DescribeCodeCoveragesOutput, error)
}

var _ DescribeCodeCoveragesAPIClient = (*Client)(nil)

// DescribeCodeCoveragesPaginatorOptions is the paginator options for
// DescribeCodeCoverages
type DescribeCodeCoveragesPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCodeCoveragesPaginator is a paginator for DescribeCodeCoverages
type DescribeCodeCoveragesPaginator struct {
	options   DescribeCodeCoveragesPaginatorOptions
	client    DescribeCodeCoveragesAPIClient
	params    *DescribeCodeCoveragesInput
	nextToken *string
	firstPage bool
}

// NewDescribeCodeCoveragesPaginator returns a new DescribeCodeCoveragesPaginator
func NewDescribeCodeCoveragesPaginator(client DescribeCodeCoveragesAPIClient, params *DescribeCodeCoveragesInput, optFns ...func(*DescribeCodeCoveragesPaginatorOptions)) *DescribeCodeCoveragesPaginator {
	if params == nil {
		params = &DescribeCodeCoveragesInput{}
	}

	options := DescribeCodeCoveragesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCodeCoveragesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCodeCoveragesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeCodeCoverages page.
func (p *DescribeCodeCoveragesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCodeCoveragesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeCodeCoverages(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeCodeCoverages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "DescribeCodeCoverages",
	}
}
