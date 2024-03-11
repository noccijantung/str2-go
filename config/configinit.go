package config

type PushConfig struct {
	Path         string
	ApiToken     string
	ApiUrl       string
	PushOnlyFile bool
}

var StrConfig *PushConfig

func SetConfig(path string, token string, url string, onlyfile bool) {
	StrConfig = &PushConfig{Path: path, ApiToken: token, ApiUrl: url, PushOnlyFile: onlyfile}
}
