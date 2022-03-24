// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureCloudSpec AzureCloudSpec specifies access credentials to Azure cloud.
//
// swagger:model AzureCloudSpec
type AzureCloudSpec struct {

	// assign availability set
	AssignAvailabilitySet bool `json:"assignAvailabilitySet,omitempty"`

	// availability set
	AvailabilitySet string `json:"availabilitySet,omitempty"`

	// client ID
	ClientID string `json:"clientID,omitempty"`

	// client secret
	ClientSecret string `json:"clientSecret,omitempty"`

	// node ports allowed IP range
	NodePortsAllowedIPRange string `json:"nodePortsAllowedIPRange,omitempty"`

	// resource group
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// route table name
	RouteTableName string `json:"routeTable,omitempty"`

	// security group
	SecurityGroup string `json:"securityGroup,omitempty"`

	// subnet name
	SubnetName string `json:"subnet,omitempty"`

	// subscription ID
	SubscriptionID string `json:"subscriptionID,omitempty"`

	// tenant ID
	TenantID string `json:"tenantID,omitempty"`

	// v net name
	VNetName string `json:"vnet,omitempty"`

	// v net resource group
	VNetResourceGroup string `json:"vnetResourceGroup,omitempty"`

	// credentials reference
	CredentialsReference *GlobalSecretKeySelector `json:"credentialsReference,omitempty"`

	// load balancer s k u
	LoadBalancerSKU LBSKU `json:"loadBalancerSKU,omitempty"`
}

// Validate validates this azure cloud spec
func (m *AzureCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBalancerSKU(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureCloudSpec) validateCredentialsReference(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsReference) { // not required
		return nil
	}

	if m.CredentialsReference != nil {
		if err := m.CredentialsReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentialsReference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentialsReference")
			}
			return err
		}
	}

	return nil
}

func (m *AzureCloudSpec) validateLoadBalancerSKU(formats strfmt.Registry) error {
	if swag.IsZero(m.LoadBalancerSKU) { // not required
		return nil
	}

	if err := m.LoadBalancerSKU.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("loadBalancerSKU")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("loadBalancerSKU")
		}
		return err
	}

	return nil
}

// ContextValidate validate this azure cloud spec based on the context it is used
func (m *AzureCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentialsReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLoadBalancerSKU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureCloudSpec) contextValidateCredentialsReference(ctx context.Context, formats strfmt.Registry) error {

	if m.CredentialsReference != nil {
		if err := m.CredentialsReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentialsReference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentialsReference")
			}
			return err
		}
	}

	return nil
}

func (m *AzureCloudSpec) contextValidateLoadBalancerSKU(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LoadBalancerSKU.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("loadBalancerSKU")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("loadBalancerSKU")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCloudSpec) UnmarshalBinary(b []byte) error {
	var res AzureCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
