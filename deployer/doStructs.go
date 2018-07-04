package deployer

type Token struct {
	AccessToken string
}

type DOConfigWrapper struct {
	ModuleName  string
	Image       string
	PrivateKey  string
	Fingerprint string
	Size        string
	DefaultUser string
	RegionMap   map[string]int
}
