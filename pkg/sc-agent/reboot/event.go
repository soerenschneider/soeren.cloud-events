package secrets

type RebootData struct {
	SecretName string `json:"secretName"`
	Path       string `json:"path"`
	Status     string `json:"status"`
}
