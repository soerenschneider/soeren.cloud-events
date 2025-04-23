package pkg

type EventTypes string

const (
	TypeSecretReplicated    EventTypes = "cloud.soeren.sc-agent.secret.replicated.v1"
	TypeSystemRebooted      EventTypes = "cloud.soeren.sc-agent.system.rebooted.v1"
	TypeUpdateApplied       EventTypes = "cloud.soeren.sc-agent.update.applied.v1"
	TypeDyndnsNewIpDetected EventTypes = "cloud.soeren.dyndns.client.new-ip-detected.v1"
	TypeDyndnsNewIpApplied  EventTypes = "cloud.soeren.dyndns.server.new-ip-applied.v1"
)
