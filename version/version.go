package version

type OrbitVersion struct {
	version string
}

var handler = &OrbitVersion{}

func SetVersion(version string) {
	handler.version = version
}

func GetVersion() string {
	return handler.version
}