// Copyright (c) tvajjal.github.io.
// Author: tvajjala@gmail.com
package cache

import (
	"context"
	"log" //legacy

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCacheClusterPatching() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider for patching cache cluster.",
		Schema: map[string]*schema.Schema{
			"compartment_ocid": {
				// This description is used by the documentation generator and the language server.
				Description: "Compartment OCID.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cache_cluster_ocid": {
				// This description is used by the documentation generator and the language server.
				Description: "Cache Cluster ocid",
				Type:        schema.TypeString,
				Required:    true,
			},
			//https://medium.com/spaceapetech/creating-a-terraform-provider-part-1-ed12884e06d7
		},
		CreateContext: resourceCacheClusterPatchingCreate,
		ReadContext:   resourceCacheClusterPatchingRead,
		UpdateContext: resourceCacheClusterPatchingUpdate,
		DeleteContext: resourceCacheClusterPatchingDelete,
	}
}

func resourceCacheClusterPatchingCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	cacheClusterId := d.Get("cache_cluster_ocid").(string)
	log.Println("Received cacheClusterId ", cacheClusterId)
	tflog.Info(ctx, cacheClusterId)

	idFromAPI := "ocid1.cache.cluster.patching.1234"
	d.SetId(idFromAPI)
	//d.SetId(record.(string)) //you can set entire json response to Id
	// then use d.lifeCycleState in terraform

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, "Created a resource")

	//c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	/*diagnostic := diag.Diagnostic{
		Summary:  "Resource Creation failed",
		Severity: diag.Error,
		Detail:   "Resource creation failed",
	}
	diags = append(diags, diagnostic)*/ //one way

	return diags
}

func resourceCacheClusterPatchingRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	var diags diag.Diagnostics

	return diags
}

func resourceCacheClusterPatchingUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	var diags diag.Diagnostics
	diags = append(diags, diag.Errorf("")...) //another way

	return diags
}

func resourceCacheClusterPatchingDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	var diags diag.Diagnostics
	return diags
}
