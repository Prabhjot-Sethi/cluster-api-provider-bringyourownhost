package algo

import (
	_ "embed"
	"os"
	"strings"
)

var (
	//go:embed config.toml.tpl
	config string
)

func GetConfigTomlEchoString() string {
	data := MakeRenderData()
	sandboxImg, found := os.LookupEnv("CKP_K8S_SANDBOX_IMAGE")
	if found {
		data.Data["SANDBOX_IMAGE"] = sandboxImg
	} else {
		data.Data["SANDBOX_IMAGE"] = ""
	}
	str, _ := RenderTemplate("config", config, &data)
	str = strings.Replace(str,"\"","\\\"", -1)
	str = strings.Replace(str,"\n","\\n", -1)
	str = "echo -e \"" + str + "\""
	return str
}
