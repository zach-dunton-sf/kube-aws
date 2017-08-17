package main

import (
	"path/filepath"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
        "os"
	"net/url"
	"fmt"
	"io/ioutil"
	"github.com/kubernetes-incubator/kube-aws-ng/config"
	//"github.com/kubernetes-incubator/kube-aws-ng/update"
)

func updateCommand(s3 *url.URL) int {
	data, err := ioutil.ReadFile("cluster.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var clusterYAML config.ClusterYAML
	if err := yaml.Unmarshal(data, &clusterYAML); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

func main() {
	app := kingpin.New(filepath.Base(os.Args[0]), "Production ready Kubernetes provisioner")
	app.Version("0.9.9")

	updateCmd := app.Command("update", "Updates/creates cluster")
	updateCmd.Alias("up")
	s3URL := updateCmd.Flag("s3-uri", "S3 Bucket URL").Required().URL()
	updateCmd.Flag("cluster-yaml", "Path to cluster.yaml").Hidden().Default("cluster.yaml").ExistingFile()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case updateCmd.FullCommand():
		os.Exit(updateCommand(*s3URL))
	}
}
