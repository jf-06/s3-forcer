// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the settings for sending SMS messages from your Amazon Web Services
// account. These settings are set with the SetSMSAttributes action.
func (c *Client) GetSMSAttributes(ctx context.Context, params *GetSMSAttributesInput, optFns ...func(*Options)) (*GetSMSAttributesOutput, error) {
	if params == nil {
		params = &GetSMSAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSMSAttributes", params, optFns, c.addOperationGetSMSAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSMSAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the GetSMSAttributes request.
type GetSMSAttributesInput struct {

	// A list of the individual attribute names, such as MonthlySpendLimit , for which
	// you want values. For all attribute names, see SetSMSAttributes (https://docs.aws.amazon.com/sns/latest/api/API_SetSMSAttributes.html)
	// . If you don't use this parameter, Amazon SNS returns all SMS attributes.
	Attributes []string

	noSmithyDocumentSerde
}

// The response from the GetSMSAttributes request.
type GetSMSAttributesOutput struct {

	// The SMS attribute names and their values.
	Attributes map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSMSAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetSMSAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetSMSAttributes{}, middleware.After)
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSMSAttributes(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetSMSAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "GetSMSAttributes",
	}
}
