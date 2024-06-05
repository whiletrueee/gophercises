package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YamlHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse yaml somehow
	var pathURLs []PathURL
	err := yaml.Unmarshal(yamlBytes, &pathURLs)
	if err != nil {
		return nil, err
	}
	pathsToUrls := make(map[string]string)

	for _, pu := range pathURLs {
		pathsToUrls[pu.Path] = pu.URL
	}

	return MapHandler(pathsToUrls, fallback), nil

}

type PathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
