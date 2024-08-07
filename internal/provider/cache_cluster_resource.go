// Copyright (c) tvajjal.github.io.
// Author: tvajjala@gmail.com
package cache

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCacheCluster() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider to create cache cluster.",
		//IMP: schema should be first before CRUD method using
		Schema: map[string]*schema.Schema{
			"compartment_ocid": {
				// This description is used by the documentation generator and the language server.
				Description: "Resource CompartmentId.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"display_name": {
				// This description is used by the documentation generator and the language server.
				Description: "Resource DisplayName - Optional",
				Type:        schema.TypeString,
				Optional:    true,
			},
			//https://medium.com/spaceapetech/creating-a-terraform-provider-part-1-ed12884e06d7
		},
		CreateContext: resourceCacheClusterCreate,
		ReadContext:   resourceCacheClusterRead,
		UpdateContext: resourceCacheClusterUpdate,
		DeleteContext: resourceCacheClusterDelete,
	}
}

func resourceCacheClusterCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	compartmentOcid := d.Get("compartment_ocid").(string)

	tflog.Info(ctx, "Received compartment_ocid", map[string]interface{}{
		"compartment_ocid": compartmentOcid,
	})

	idFromAPI := "ocid1.cache.cluster.id.1234"
	d.SetId(idFromAPI)
	//d.SetId(record.(string)) //you can set entire json response to Id
	// then use d.lifeCycleState in terraform

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Info(ctx, "Created a cluster resource")

	//c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

func resourceCacheClusterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	var diags diag.Diagnostics
	return diags
}

func resourceCacheClusterUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	var diags diag.Diagnostics
	return diags
}

func resourceCacheClusterDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	var diags diag.Diagnostics
	return diags
}
