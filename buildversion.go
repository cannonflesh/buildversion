package buildversion

var (
	// Version текущий git tag, v0.0.0 если не объявлена.
	Version string
	// Commit укороченный хеш текущего коммита.
	Commit string
	// Branch имя текущей ветки.
	Branch string
	// Env предназначение билда (staging или production).
	Env string
	// BuildTime время билда
	BuildTime string
)

func init() {
	if Version == "" {
		Version = "v0.0.0"
	}
}
