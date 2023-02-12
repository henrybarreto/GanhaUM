package paths

const (
	CreateContributor  string = "/contributors"
	GetContributor     string = "/contributors/:campaign_id/:id"
	GetContributors    string = "/contributors/:campaign_id"
	ConfirmContributor string = "/contributors/confirm/:campaign_id/:id"
)
