package main

import (
	"path/filepath"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
        "os"
	"github.com/kubernetes-incubator/kube-aws-ng/model"
	"github.com/kubernetes-incubator/kube-aws-ng/update"
)

func updateCommand(s3 *net.URL) int {
	data, err := ioutil.ReadFile("cluster.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var clusterYAML model.ClusterYAML
	if err := yaml.Unmarshall(data, &clusterYAML); !err {

	}
	return 0
}

func main() {
	app := kingpin.New(filepath.Base(os.Args[0]), "Production ready Kubernetes provisioner")
	app.Version("0.9.9")

	updateCmd := app.Command("update", "Updates/creates cluster")
	updateCmd.Alias("up")
	s3URL = updateCmd.Flag("s3-uri", "S3 Bucket URL").Required().URL()
	updateCmd.Flag("cluster-yaml").Hidden().Defaults("cluster.yaml").ExistingFile()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case updateCmd.FullCommand():
		os.Exit(updateCommand(s3URL))
	}
}
