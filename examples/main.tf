//This block required when you do local testing, if plugin is not publicly available. 
terraform {
  required_version = ">= 0.13.0" # terraform core version
  required_providers {
    cache = {
      source  = "tvajjala.github.io/service/cache"
      version = ">=0.0.1" # our plugin version which we mentioned in Makefile
    }
  }
}

// name of the provider is cache
provider "cache" {
  endpoint_url     = "localhost"     // control plane api endpoint (required)
   auth      =     "Bearer Oracle" // optional auth token (optional)
}

/**
  Resource name must be ${provider}_resource_xxxx format
  Supported resources for this provider are
  1. cache_cluster_resource
  2. cache_cluster_patching_resource
*/

resource "cache_cluster_resource" "redis-cluster" {
  display_name     = "redis-cluster" //optional
  compartment_ocid = "ocid1.compartment.id"
}

resource "cache_cluster_patching_resource" "redis-patch" {
  // populated from above resource
  cache_cluster_ocid   =  cache_cluster_resource.redis-cluster.id
  compartment_ocid     = "ocid1.compartment.id"
}

output "cache_cluster_ocid" {
  value = cache_cluster_resource.redis-cluster.id
}

output "cache_cluster_patching_ocid" {
  value = cache_cluster_patching_resource.redis-patch.id
}
