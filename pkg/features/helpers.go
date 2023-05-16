package features

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"fmt"
	"github.com/crumbhole/argocd-lovely-plugin/pkg/config"
	yaml "sigs.k8s.io/yaml"
	"strings"
)

func getPlugins(envname string) []string {
	pluginsText := config.GetStringParam(envname, ``)
	if pluginsText == `` {
		return make([]string, 0)
	}
	plugins := strings.Split(pluginsText, `,`)
	for i, plugin := range plugins {
		plugins[i] = strings.TrimSpace(plugin)
	}
	return plugins
}

type pluginYaml map[string][]string

func getYamlPlugins(envname string) (pluginYaml, error) {
	pluginsText := config.GetStringParam(envname, ``)
	if pluginsText == `` {
		return make(pluginYaml), nil
	}
	var plugins pluginYaml
	err := yaml.Unmarshal([]byte(pluginsText), &plugins)
	if err != nil {
		return make(pluginYaml), err
	}
	return plugins, nil
}

// Path here must be a relative path
func pluginsForPath(path string, yamlEnv string, plainEnv string) ([]string, error) {
	plugins, err := getYamlPlugins(yamlEnv)
	if err != nil {
		return make([]string, 0), fmt.Errorf("%s is invalid: %s", yamlEnv, err)
	}
	if list, contains := plugins[path]; contains {
		return list, nil
	}
	return getPlugins(plainEnv), nil
}
