package auth

// Encrypter interface for different encryption methods
type Encrypter interface {
	Encrypt(string) (string, string, error)
	Validate(string, string, string) (bool, error)
}
