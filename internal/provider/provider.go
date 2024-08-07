package cache

// Copyright (c) tvajjal.github.io.
// Author: tvajjala@gmail.com

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"auth": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("AUTH", nil),
					Description: "Cache API Key from https://cache.tvajjala.githubt.io",
				},
				"endpoint_url": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("ENDPOINT_URL", nil),
					Description: "Cache API URL",
				},
			},

			//NOTE: must start with sfcp
			DataSourcesMap: map[string]*schema.Resource{
				"cache_cluster_datasource": dataSourceCacheCluster(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"cache_cluster_resource":          resourceCacheCluster(),
				"cache_cluster_patching_resource": resourceCacheClusterPatching(),
			},
			//api endpoint

		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
}

func configure(version string, p *schema.Provider) func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		// Setup a User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-scaffolding", version)
		// TODO: myClient.UserAgent = userAgent
		BaseURL, ok := d.Get("endpoint_url").(string)
		tflog.Info(ctx, "Received BaseURL", map[string]interface{}{
			"BaseURL": BaseURL,
		})

		if !ok {
			tflog.Warn(ctx, "Failed fetch endpoint endpoint_url")
		}

		return &apiClient{}, nil
	}
}
