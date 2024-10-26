// Code generated by smithy-go-codegen DO NOT EDIT.

package sts

import (
	"context"
	"fmt"

	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the account identifier for the specified access key ID.
//
// Access keys consist of two parts: an access key ID (for example,
// AKIAIOSFODNN7EXAMPLE ) and a secret access key (for example,
// wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY ). For more information about access
// keys, see [Managing Access Keys for IAM Users]in the IAM User Guide.
//
// When you pass an access key ID to this operation, it returns the ID of the
// Amazon Web Services account to which the keys belong. Access key IDs beginning
// with AKIA are long-term credentials for an IAM user or the Amazon Web Services
// account root user. Access key IDs beginning with ASIA are temporary credentials
// that are created using STS operations. If the account in the response belongs to
// you, you can sign in as the root user and review your root user access keys.
// Then, you can pull a [credentials report]to learn which IAM user owns the keys. To learn who
// requested the temporary credentials for an ASIA access key, view the STS events
// in your [CloudTrail logs]in the IAM User Guide.
//
// This operation does not indicate the state of the access key. The key might be
// active, inactive, or deleted. Active keys might not have permissions to perform
// an operation. Providing a deleted access key might return an error that the key
// doesn't exist.
//
// [credentials report]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_getting-report.html
// [CloudTrail logs]: https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-integration.html
// [Managing Access Keys for IAM Users]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html
func (c *Client) GetAccessKeyInfo(ctx context.Context, params *GetAccessKeyInfoInput, optFns ...func(*Options)) (*GetAccessKeyInfoOutput, error) {
	if params == nil {
		params = &GetAccessKeyInfoInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessKeyInfo", params, optFns, c.addOperationGetAccessKeyInfoMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessKeyInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessKeyInfoInput struct {

	// The identifier of an access key.
	//
	// This parameter allows (through its regex pattern) a string of characters that
	// can consist of any upper- or lowercase letter or digit.
	//
	// This member is required.
	AccessKeyId *string

	noSmithyDocumentSerde
}

type GetAccessKeyInfoOutput struct {

	// The number used to identify the Amazon Web Services account.
	Account *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccessKeyInfoMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetAccessKeyInfo{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetAccessKeyInfo{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAccessKeyInfo"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpGetAccessKeyInfoValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessKeyInfo(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
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

func newServiceMetadataMiddleware_opGetAccessKeyInfo(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAccessKeyInfo",
	}
}
