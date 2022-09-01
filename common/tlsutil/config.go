package tlsutil

type TLSConfig struct {
    TLSDisable    bool
    TLSCaFile     string
    TLSCertFile   string
    TLSKeyFile    string
    TLSMinVersion string
    ServerName    string
}
