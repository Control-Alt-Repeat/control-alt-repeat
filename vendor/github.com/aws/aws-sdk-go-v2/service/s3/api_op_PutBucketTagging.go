// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalChecksum "github.com/aws/aws-sdk-go-v2/service/internal/checksum"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is not supported by directory buckets.
//
// Sets the tags for a bucket.
//
// Use tags to organize your Amazon Web Services bill to reflect your own cost
// structure. To do this, sign up to get your Amazon Web Services account bill with
// tag key values included. Then, to see the cost of combined resources, organize
// your billing information according to resources with the same tag key values.
// For example, you can tag several resources with a specific application name, and
// then organize your billing information to see the total cost of that application
// across several services. For more information, see [Cost Allocation and Tagging]and [Using Cost Allocation in Amazon S3 Bucket Tags].
//
// When this operation sets the tags for a bucket, it will overwrite any current
// tags the bucket already has. You cannot use this operation to add tags to an
// existing list of tags.
//
// To use this operation, you must have permissions to perform the
// s3:PutBucketTagging action. The bucket owner has this permission by default and
// can grant this permission to others. For more information about permissions, see
// [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// PutBucketTagging has the following special errors. For more Amazon S3 errors
// see, [Error Responses].
//
//   - InvalidTag - The tag provided was not a valid tag. This error can occur if
//     the tag did not pass input validation. For more information, see [Using Cost Allocation in Amazon S3 Bucket Tags].
//
//   - MalformedXML - The XML provided does not match the schema.
//
//   - OperationAborted - A conflicting conditional action is currently in progress
//     against this resource. Please try again.
//
//   - InternalError - The service was unable to apply the provided tag to the
//     bucket.
//
// The following operations are related to PutBucketTagging :
//
// [GetBucketTagging]
//
// [DeleteBucketTagging]
//
// [Error Responses]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html
// [GetBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketTagging.html
// [Cost Allocation and Tagging]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [DeleteBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html
// [Using Cost Allocation in Amazon S3 Bucket Tags]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/CostAllocTagging.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
func (c *Client) PutBucketTagging(ctx context.Context, params *PutBucketTaggingInput, optFns ...func(*Options)) (*PutBucketTaggingOutput, error) {
	if params == nil {
		params = &PutBucketTaggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketTagging", params, optFns, c.addOperationPutBucketTaggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketTaggingInput struct {

	// The bucket name.
	//
	// This member is required.
	Bucket *string

	// Container for the TagSet and Tag elements.
	//
	// This member is required.
	Tagging *types.Tagging

	// Indicates the algorithm used to create the checksum for the object when you use
	// the SDK. This header will not provide any additional functionality if you don't
	// use the SDK. When you send this header, there must be a corresponding
	// x-amz-checksum or x-amz-trailer header sent. Otherwise, Amazon S3 fails the
	// request with the HTTP status code 400 Bad Request . For more information, see [Checking object integrity]
	// in the Amazon S3 User Guide.
	//
	// If you provide an individual checksum, Amazon S3 ignores any provided
	// ChecksumAlgorithm parameter.
	//
	// [Checking object integrity]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html
	ChecksumAlgorithm types.ChecksumAlgorithm

	// The base64-encoded 128-bit MD5 digest of the data. You must use this header as
	// a message integrity check to verify that the request body was not corrupted in
	// transit. For more information, see [RFC 1864].
	//
	// For requests made using the Amazon Web Services Command Line Interface (CLI) or
	// Amazon Web Services SDKs, this field is calculated automatically.
	//
	// [RFC 1864]: http://www.ietf.org/rfc/rfc1864.txt
	ContentMD5 *string

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

func (in *PutBucketTaggingInput) bindEndpointParams(p *EndpointParameters) {

	p.Bucket = in.Bucket
	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
}

type PutBucketTaggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketTaggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketTagging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketTagging{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBucketTagging"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addIsExpressUserAgent(stack); err != nil {
		return err
	}
	if err = addOpPutBucketTaggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketTagging(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutBucketTaggingInputChecksumMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addPutBucketTaggingUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3cust.AddExpressDefaultChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func (v *PutBucketTaggingInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opPutBucketTagging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBucketTagging",
	}
}

// getPutBucketTaggingRequestAlgorithmMember gets the request checksum algorithm
// value provided as input.
func getPutBucketTaggingRequestAlgorithmMember(input interface{}) (string, bool) {
	in := input.(*PutBucketTaggingInput)
	if len(in.ChecksumAlgorithm) == 0 {
		return "", false
	}
	return string(in.ChecksumAlgorithm), true
}

func addPutBucketTaggingInputChecksumMiddlewares(stack *middleware.Stack, options Options) error {
	return internalChecksum.AddInputMiddleware(stack, internalChecksum.InputMiddlewareOptions{
		GetAlgorithm:                     getPutBucketTaggingRequestAlgorithmMember,
		RequireChecksum:                  true,
		EnableTrailingChecksum:           false,
		EnableComputeSHA256PayloadHash:   true,
		EnableDecodedContentLengthHeader: true,
	})
}

// getPutBucketTaggingBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getPutBucketTaggingBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketTaggingInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketTaggingUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketTaggingBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
