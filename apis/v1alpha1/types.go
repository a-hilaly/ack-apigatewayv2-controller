// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

type API struct {
	APIEndpoint               *string            `json:"apiEndpoint,omitempty"`
	APIGatewayManaged         *bool              `json:"apiGatewayManaged,omitempty"`
	APIID                     *string            `json:"apiID,omitempty"`
	APIKeySelectionExpression *string            `json:"apiKeySelectionExpression,omitempty"`
	CreatedDate               *metav1.Time       `json:"createdDate,omitempty"`
	Description               *string            `json:"description,omitempty"`
	DisableExecuteAPIEndpoint *bool              `json:"disableExecuteAPIEndpoint,omitempty"`
	DisableSchemaValidation   *bool              `json:"disableSchemaValidation,omitempty"`
	ImportInfo                []*string          `json:"importInfo,omitempty"`
	Name                      *string            `json:"name,omitempty"`
	RouteSelectionExpression  *string            `json:"routeSelectionExpression,omitempty"`
	Tags                      map[string]*string `json:"tags,omitempty"`
	Version                   *string            `json:"version,omitempty"`
	Warnings                  []*string          `json:"warnings,omitempty"`
}

type APIMapping struct {
	APIID         *string `json:"apiID,omitempty"`
	APIMappingID  *string `json:"apiMappingID,omitempty"`
	APIMappingKey *string `json:"apiMappingKey,omitempty"`
	Stage         *string `json:"stage,omitempty"`
}

type AccessLogSettings struct {
	DestinationARN *string `json:"destinationARN,omitempty"`
	Format         *string `json:"format,omitempty"`
}

type Authorizer struct {
	AuthorizerCredentialsARN       *string `json:"authorizerCredentialsARN,omitempty"`
	AuthorizerID                   *string `json:"authorizerID,omitempty"`
	AuthorizerPayloadFormatVersion *string `json:"authorizerPayloadFormatVersion,omitempty"`
	AuthorizerURI                  *string `json:"authorizerURI,omitempty"`
	EnableSimpleResponses          *bool   `json:"enableSimpleResponses,omitempty"`
	IDentityValidationExpression   *string `json:"identityValidationExpression,omitempty"`
	Name                           *string `json:"name,omitempty"`
}

type Cors struct {
	AllowCredentials *bool `json:"allowCredentials,omitempty"`
}

type Deployment struct {
	AutoDeployed            *bool        `json:"autoDeployed,omitempty"`
	CreatedDate             *metav1.Time `json:"createdDate,omitempty"`
	DeploymentID            *string      `json:"deploymentID,omitempty"`
	DeploymentStatusMessage *string      `json:"deploymentStatusMessage,omitempty"`
	Description             *string      `json:"description,omitempty"`
}

type DomainNameConfiguration struct {
	APIGatewayDomainName    *string      `json:"apiGatewayDomainName,omitempty"`
	CertificateARN          *string      `json:"certificateARN,omitempty"`
	CertificateName         *string      `json:"certificateName,omitempty"`
	CertificateUploadDate   *metav1.Time `json:"certificateUploadDate,omitempty"`
	DomainNameStatus        *string      `json:"domainNameStatus,omitempty"`
	DomainNameStatusMessage *string      `json:"domainNameStatusMessage,omitempty"`
	EndpointType            *string      `json:"endpointType,omitempty"`
	HostedZoneID            *string      `json:"hostedZoneID,omitempty"`
	SecurityPolicy          *string      `json:"securityPolicy,omitempty"`
}

type DomainName_SDK struct {
	APIMappingSelectionExpression *string                    `json:"apiMappingSelectionExpression,omitempty"`
	DomainName                    *string                    `json:"domainName,omitempty"`
	DomainNameConfigurations      []*DomainNameConfiguration `json:"domainNameConfigurations,omitempty"`
	MutualTLSAuthentication       *MutualTLSAuthentication   `json:"mutualTLSAuthentication,omitempty"`
	Tags                          map[string]*string         `json:"tags,omitempty"`
}

type Integration struct {
	APIGatewayManaged                      *bool   `json:"apiGatewayManaged,omitempty"`
	ConnectionID                           *string `json:"connectionID,omitempty"`
	CredentialsARN                         *string `json:"credentialsARN,omitempty"`
	Description                            *string `json:"description,omitempty"`
	IntegrationID                          *string `json:"integrationID,omitempty"`
	IntegrationMethod                      *string `json:"integrationMethod,omitempty"`
	IntegrationResponseSelectionExpression *string `json:"integrationResponseSelectionExpression,omitempty"`
	IntegrationSubtype                     *string `json:"integrationSubtype,omitempty"`
	IntegrationURI                         *string `json:"integrationURI,omitempty"`
	PayloadFormatVersion                   *string `json:"payloadFormatVersion,omitempty"`
	TemplateSelectionExpression            *string `json:"templateSelectionExpression,omitempty"`
}

type IntegrationResponse struct {
	IntegrationResponseID       *string `json:"integrationResponseID,omitempty"`
	IntegrationResponseKey      *string `json:"integrationResponseKey,omitempty"`
	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty"`
}

type JWTConfiguration struct {
	Audience []*string `json:"audience,omitempty"`
	Issuer   *string   `json:"issuer,omitempty"`
}

type Model struct {
	Description *string `json:"description,omitempty"`
	ModelID     *string `json:"modelID,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type MutualTLSAuthentication struct {
	TruststoreURI      *string   `json:"truststoreURI,omitempty"`
	TruststoreVersion  *string   `json:"truststoreVersion,omitempty"`
	TruststoreWarnings []*string `json:"truststoreWarnings,omitempty"`
}

type MutualTLSAuthenticationInput struct {
	TruststoreURI     *string `json:"truststoreURI,omitempty"`
	TruststoreVersion *string `json:"truststoreVersion,omitempty"`
}

type ParameterConstraints struct {
	Required *bool `json:"required,omitempty"`
}

type RouteResponse_SDK struct {
	ModelSelectionExpression *string                            `json:"modelSelectionExpression,omitempty"`
	ResponseModels           map[string]*string                 `json:"responseModels,omitempty"`
	ResponseParameters       []map[string]*ParameterConstraints `json:"responseParameters,omitempty"`
	RouteResponseID          *string                            `json:"routeResponseID,omitempty"`
	RouteResponseKey         *string                            `json:"routeResponseKey,omitempty"`
}

type RouteSettings struct {
	DataTraceEnabled       *bool    `json:"dataTraceEnabled,omitempty"`
	DetailedMetricsEnabled *bool    `json:"detailedMetricsEnabled,omitempty"`
	LoggingLevel           *string  `json:"loggingLevel,omitempty"`
	ThrottlingBurstLimit   *int64   `json:"throttlingBurstLimit,omitempty"`
	ThrottlingRateLimit    *float64 `json:"throttlingRateLimit,omitempty"`
}

type Route_SDK struct {
	APIGatewayManaged                *bool                              `json:"apiGatewayManaged,omitempty"`
	APIKeyRequired                   *bool                              `json:"apiKeyRequired,omitempty"`
	AuthorizationScopes              []*string                          `json:"authorizationScopes,omitempty"`
	AuthorizationType                *string                            `json:"authorizationType,omitempty"`
	AuthorizerID                     *string                            `json:"authorizerID,omitempty"`
	ModelSelectionExpression         *string                            `json:"modelSelectionExpression,omitempty"`
	OperationName                    *string                            `json:"operationName,omitempty"`
	RequestModels                    map[string]*string                 `json:"requestModels,omitempty"`
	RequestParameters                []map[string]*ParameterConstraints `json:"requestParameters,omitempty"`
	RouteID                          *string                            `json:"routeID,omitempty"`
	RouteKey                         *string                            `json:"routeKey,omitempty"`
	RouteResponseSelectionExpression *string                            `json:"routeResponseSelectionExpression,omitempty"`
	Target                           *string                            `json:"target,omitempty"`
}

type Stage_SDK struct {
	AccessLogSettings           *AccessLogSettings          `json:"accessLogSettings,omitempty"`
	APIGatewayManaged           *bool                       `json:"apiGatewayManaged,omitempty"`
	AutoDeploy                  *bool                       `json:"autoDeploy,omitempty"`
	ClientCertificateID         *string                     `json:"clientCertificateID,omitempty"`
	CreatedDate                 *metav1.Time                `json:"createdDate,omitempty"`
	DefaultRouteSettings        *RouteSettings              `json:"defaultRouteSettings,omitempty"`
	DeploymentID                *string                     `json:"deploymentID,omitempty"`
	Description                 *string                     `json:"description,omitempty"`
	LastDeploymentStatusMessage *string                     `json:"lastDeploymentStatusMessage,omitempty"`
	LastUpdatedDate             *metav1.Time                `json:"lastUpdatedDate,omitempty"`
	RouteSettings               []map[string]*RouteSettings `json:"routeSettings,omitempty"`
	StageName                   *string                     `json:"stageName,omitempty"`
	StageVariables              map[string]*string          `json:"stageVariables,omitempty"`
	Tags                        map[string]*string          `json:"tags,omitempty"`
}

type TLSConfig struct {
	ServerNameToVerify *string `json:"serverNameToVerify,omitempty"`
}

type TLSConfigInput struct {
	ServerNameToVerify *string `json:"serverNameToVerify,omitempty"`
}

type VPCLink struct {
	CreatedDate          *metav1.Time       `json:"createdDate,omitempty"`
	Name                 *string            `json:"name,omitempty"`
	Tags                 map[string]*string `json:"tags,omitempty"`
	VPCLinkID            *string            `json:"vpcLinkID,omitempty"`
	VPCLinkStatusMessage *string            `json:"vpcLinkStatusMessage,omitempty"`
}
