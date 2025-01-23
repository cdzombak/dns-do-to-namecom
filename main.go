package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"

	ec "github.com/cdzombak/exitcode_go"
	"github.com/digitalocean/godo"
	_ "github.com/joho/godotenv/autoload"
	"github.com/namedotcom/go/namecom"
)

const (
	EnvDigitalOceanAPIToken = "DO_API_TOKEN"
	EnvNamecomUsername      = "NC_USERNAME"
	EnvNamecomToken         = "NC_API_TOKEN"
)

func main() {
	domain := flag.String("domain", "", "Domain to migrate. Required.")
	dryRun := flag.Bool("dry-run", true, "Dry run.")
	flag.Parse()

	if *domain == "" {
		Eprintln("Domain is required.")
		os.Exit(ec.Usage)
	}
	doToken := os.Getenv(EnvDigitalOceanAPIToken)
	if doToken == "" {
		Eprintln("DigitalOcean API token is required.")
		os.Exit(ec.NotConfigured)
	}
	ncUsername := os.Getenv(EnvNamecomUsername)
	if ncUsername == "" {
		Eprintln("Name.com username is required.")
		os.Exit(ec.NotConfigured)
	}
	ncToken := os.Getenv(EnvNamecomToken)
	if ncToken == "" {
		Eprintln("Name.com API token is required.")
		os.Exit(ec.NotConfigured)
	}

	ctx := context.Background()
	doClient := godo.NewFromToken(doToken)
	ncClient := namecom.New(ncUsername, ncToken)

	// call into migrate func, w/clients and dry run
	if err := Migrate(ctx, doClient, ncClient, strings.ToLower(*domain), *dryRun); err != nil {
		log.Fatal(err)
	}
}
