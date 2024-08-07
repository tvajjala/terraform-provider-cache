package cache

// Copyright (c) tvajjal.github.io.
// Author: tvajjala@gmail.com
import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCacheCluster() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample data source in the Terraform provider cache.",

		ReadContext: dataSourceCacheClusterRead,

		Schema: map[string]*schema.Schema{},
	}
}

func dataSourceCacheClusterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	//client := meta.(*apiClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	idFromAPI := "cache.cluster.id.1234"
	d.SetId(idFromAPI)
	tflog.Trace(ctx, "Returnig Record")

	return diags
}
