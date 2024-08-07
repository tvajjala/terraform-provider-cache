// Copyright (c) tvajjal.github.io.
// Author: tvajjala@gmail.com

package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	cache "github.io/tvajjala/terraform-provider-cache/internal/provider"
)

//FIXME: importing sub module fails please run `go get github.io/tvajjala/terraform-provider-cache`

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but it's suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	var debugMode bool
	// don't make debug true here, it fails with below error
	// TF_REATTACH_PROVIDERS environment variable with the following
	flag.BoolVar(&debugMode, "debug", false, "set to true")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug: debugMode,
		// TODO: update this string with the full name of your provider as used in your configs
		//NOTE: this is optional, required in local testing/
		ProviderAddr: "tvajjala.github.io/service/cache", // this plugin under ~/.terraform.d
		ProviderFunc: cache.New(version),
	}

	plugin.Serve(opts)
}
