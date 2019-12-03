package main

import (
	"flag"
	"log"

	tektonclientset "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"k8s.io/client-go/rest"

	"github.com/mattmoor/knobots/pkg/gotool"
	"github.com/mattmoor/knobots/pkg/handler"
)

func main() {
	flag.Parse()

	cfg, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	bc, err := tektonclientset.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Unable to build Tekton client: %v", err)
	}

	handler.Main(gotool.New(bc))
}
