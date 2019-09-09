package webhook

const (
	IncidentTrigger       = "incident.trigger"       // IncidentTrigger is sent when an incident is newly created/triggered
	IncidentAcknowledge   = "incident.acknowledge"   // IncidentAcknowledge is sent when an incident is acknowledged by a user
	IncidentUnAcknowledge = "incident.unacknowledge" // IncidentUnacknowledge is sent when an incident is unacknowledged due to its acknowledgement timing out
	IncidentResolve       = "incident.resolve"       // IncidentResolve is sent when an incident has been resolved
	IncidentAssign        = "incident.assign"        // IncidentAssign is sent when an incident has been assigned to another user. Often occurs in concert with an acknowledge
	IncidentEscalate      = "incident.escalate"      // IncidentEscalate is sent when an incident has been escalated to another user in the same escalation chain
	IncidentDelegate      = "incident.delegate"      // IncidentDelegate is sent when an incident has been reassigned to another escalation policy
	IncidentAnnotate      = "incident.annotate"      // IncidentAnnotate is sent when a note is created on an incident.
)
