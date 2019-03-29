package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/aymerick/raymond"
)

type (
	KubeConfig struct {
		Ca        string
		Server    string
		Token     string
		Namespace string
		Template  string
	}
	Plugin struct {
		Template   string
		KubeConfig KubeConfig
	}
)

func (p Plugin) Exec() error {
	if p.KubeConfig.Server == "" {
		log.Fatal("KUBE_SERVER is not defined")
	}
	if p.KubeConfig.Token == "" {
		log.Fatal("KUBE_TOKEN is not defined")
	}
	if p.KubeConfig.Ca == "" {
		log.Fatal("KUBE_CA is not defined")
	}
	if p.KubeConfig.Namespace == "" {
		p.KubeConfig.Namespace = "default"
	}
	if p.Template == "" {
		log.Fatal("KUBE_TEMPLATE, or template must be defined")
	}

	// // connect to Kubernetes
	// clientset, err := p.createKubeClient()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	raw, err := ioutil.ReadFile(p.Template)
	if err != nil {
		return err
	}

	source := string(raw)

	ctx := make(map[string]string)
	ctx["KUBE_CA"] = p.KubeConfig.Ca
	ctx["KUBE_TOKEN"] = p.KubeConfig.Ca
	ctx["KUBE_SERVER"] = p.KubeConfig.Ca
	droneEnv := os.Environ()
	for _, value := range droneEnv {
		re := regexp.MustCompile(`^(DRONE_.*)=(.*)`)
		if re.MatchString(value) {
			matches := re.FindStringSubmatch(value)
			ctx[matches[1]] = matches[2]
		}
	}

	// parse template
	tpl, err := raymond.Parse(source)
	if err != nil {
		panic(err)
	}

	result, err := tpl.Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Print(result)

	return err
}
