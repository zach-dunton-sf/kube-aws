package main

import (
   "strings"
  "os"
  "github.com/clipperhouse/typewriter"
  _ "github.com/kubernetes-incubator/kube-aws-ng/_codegen/yamlregexvalidator"
  _ "github.com/kubernetes-incubator/kube-aws-ng/_codegen/yamlsimpleparse"
)

func skipTests(f os.FileInfo) bool {
	return !strings.HasSuffix(f.Name(),"_test.go")
}


func main() {
      app, err := typewriter.NewAppFiltered("+gen", skipTests)
      if err != nil {
         panic(err)
      }
      _, err = app.WriteAll()
      if err != nil {
         panic(err)
      }
}
