package toproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func GetProviderSchema_Request(in *tfprotov6.GetProviderSchemaRequest) (*tfplugin6.GetProviderSchema_Request, error) {
	return &tfplugin6.GetProviderSchema_Request{}, nil
}

func GetProviderSchema_Response(in *tfprotov6.GetProviderSchemaResponse) (*tfplugin6.GetProviderSchema_Response, error) {
	if in == nil {
		return nil, nil
	}
	resp := tfplugin6.GetProviderSchema_Response{
		ServerCapabilities: GetProviderSchema_ServerCapabilities(in.ServerCapabilities),
	}
	if in.Provider != nil {
		schema, err := Schema(in.Provider)
		if err != nil {
			return &resp, fmt.Errorf("error marshaling provider schema: %w", err)
		}
		resp.Provider = schema
	}
	if in.ProviderMeta != nil {
		schema, err := Schema(in.ProviderMeta)
		if err != nil {
			return &resp, fmt.Errorf("error marshaling provider_meta schema: %w", err)
		}
		resp.ProviderMeta = schema
	}
	resp.ResourceSchemas = make(map[string]*tfplugin6.Schema, len(in.ResourceSchemas))
	for k, v := range in.ResourceSchemas {
		if v == nil {
			resp.ResourceSchemas[k] = nil
			continue
		}
		schema, err := Schema(v)
		if err != nil {
			return &resp, fmt.Errorf("error marshaling resource schema for %q: %w", k, err)
		}
		resp.ResourceSchemas[k] = schema
	}
	resp.DataSourceSchemas = make(map[string]*tfplugin6.Schema, len(in.DataSourceSchemas))
	for k, v := range in.DataSourceSchemas {
		if v == nil {
			resp.DataSourceSchemas[k] = nil
			continue
		}
		schema, err := Schema(v)
		if err != nil {
			return &resp, fmt.Errorf("error marshaling data source schema for %q: %w", k, err)
		}
		resp.DataSourceSchemas[k] = schema
	}
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return &resp, err
	}
	resp.Diagnostics = diags
	return &resp, nil
}

func ValidateProviderConfig_Request(in *tfprotov6.ValidateProviderConfigRequest) (*tfplugin6.ValidateProviderConfig_Request, error) {
	resp := &tfplugin6.ValidateProviderConfig_Request{}
	if in.Config != nil {
		resp.Config = DynamicValue(in.Config)
	}
	return resp, nil
}

func ValidateProviderConfig_Response(in *tfprotov6.ValidateProviderConfigResponse) (*tfplugin6.ValidateProviderConfig_Response, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	resp := &tfplugin6.ValidateProviderConfig_Response{
		Diagnostics: diags,
	}
	return resp, nil
}

func Configure_Request(in *tfprotov6.ConfigureProviderRequest) (*tfplugin6.ConfigureProvider_Request, error) {
	resp := &tfplugin6.ConfigureProvider_Request{
		TerraformVersion: in.TerraformVersion,
	}
	if in.Config != nil {
		resp.Config = DynamicValue(in.Config)
	}
	return resp, nil
}

func Configure_Response(in *tfprotov6.ConfigureProviderResponse) (*tfplugin6.ConfigureProvider_Response, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	return &tfplugin6.ConfigureProvider_Response{
		Diagnostics: diags,
	}, nil
}

func Stop_Request(in *tfprotov6.StopProviderRequest) (*tfplugin6.StopProvider_Request, error) {
	return &tfplugin6.StopProvider_Request{}, nil
}

func Stop_Response(in *tfprotov6.StopProviderResponse) (*tfplugin6.StopProvider_Response, error) {
	return &tfplugin6.StopProvider_Response{
		Error: in.Error,
	}, nil
}

// we have to say this next thing to get golint to stop yelling at us about the
// underscores in the function names. We want the function names to match
// actually-generated code, so it feels like fair play. It's just a shame we
// lose golint for the entire file.
//
// This file is not actually generated. You can edit it. Ignore this next line.
// Code generated by hand ignore this next bit DO NOT EDIT.
