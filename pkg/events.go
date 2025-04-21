package pkg

type EventTypes string

const (
	TypeSecretReplicated EventTypes = "cloud.soeren.sc-agent.secret.replicated.v1"
	TypeSystemRebooted   EventTypes = "cloud.soeren.sc-agent.system.rebooted.v1"
	TypeUpdateApplied    EventTypes = "cloud.soeren.sc-agent.update.applied.v1"
)
