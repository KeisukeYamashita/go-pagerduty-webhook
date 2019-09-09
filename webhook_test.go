package webhook

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

var (
	testIncidentTriggerPayloadStr = `
	{
		"messages": [
		  {
			"event": "incident.trigger",
			"log_entries": [
			  {
				"id": "R2XGXEI3W0FHMSDXHDIBQGBQ5E",
				"type": "trigger_log_entry",
				"summary": "Triggered through the website",
				"self": "https://api.pagerduty.com/log_entries/R2XGXEI3W0FHMSDXHDIBQGBQ5E",
				"html_url": "https://webdemo.pagerduty.com/incidents/PRORDTY/log_entries/R2XGXEI3W0FHMSDXHDIBQGBQ5E",
				"created_at": "2017-09-26T15:14:36Z",
				"agent": {
				  "id": "P553OPV",
				  "type": "user_reference",
				  "summary": "Laura Haley",
				  "self": "https://api.pagerduty.com/users/P553OPV",
				  "html_url": "https://webdemo.pagerduty.com/users/P553OPV"
				},
				"channel": {
				  "type": "web_trigger",
				  "summary": "My new incident",
				  "subject": "My new incident",
				  "details": "Oh my gosh"
				},
				"service": {
				  "id": "PN49J75",
				  "type": "service_reference",
				  "summary": "Production XDB Cluster",
				  "self": "https://api.pagerduty.com/services/PN49J75",
				  "html_url": "https://webdemo.pagerduty.com/services/PN49J75"
				},
				"incident": {
				  "id": "PRORDTY",
				  "type": "incident_reference",
				  "summary": "[#33] My new incident",
				  "self": "https://api.pagerduty.com/incidents/PRORDTY",
				  "html_url": "https://webdemo.pagerduty.com/incidents/PRORDTY"
				},
				"teams": [
				  {
					"id": "P4SI59S",
					"type": "team_reference",
					"summary": "Engineering",
					"self": "https://api.pagerduty.com/teams/P4SI59S",
					"html_url": "https://webdemo.pagerduty.com/teams/P4SI59S"
				  }
				],
				"contexts": [],
				"event_details": {
				  "description": "My new incident"
				}
			  }
			],
			"webhook": {
			  "endpoint_url": "https://requestb.in/18ao6fs1",
			  "name": "V2 wabhook",
			  "description": null,
			  "webhook_object": {
				"id": "PN49J75",
				"type": "service_reference",
				"summary": "Production XDB Cluster",
				"self": "https://api.pagerduty.com/services/PN49J75",
				"html_url": "https://webdemo.pagerduty.com/services/PN49J75"
			  },
			  "config": {},
			  "outbound_integration": {
				"id": "PJFWPEP",
				"type": "outbound_integration_reference",
				"summary": "Generic V2 Webhook",
				"self": "https://api.pagerduty.com/outbound_integrations/PJFWPEP",
				"html_url": null
			  },
			  "accounts_addon": null,
			  "id": "PKT9NNX",
			  "type": "webhook",
			  "summary": "V2 wabhook",
			  "self": "https://api.pagerduty.com/webhooks/PKT9NNX",
			  "html_url": null
			},
			"incident": {
			  "incident_number": 33,
			  "title": "My new incident",
			  "description": "My new incident",
			  "created_at": "2017-09-26T15:14:36Z",
			  "status": "triggered",
			  "pending_actions": [
				{
				  "type": "escalate",
				  "at": "2017-09-26T15:44:36Z"
				},
				{
				  "type": "resolve",
				  "at": "2017-09-26T19:14:36Z"
				}
			  ],
			  "incident_key": null,
			  "service": {
				"id": "PN49J75",
				"name": "Production XDB Cluster",
				"description": "This service was created during onboarding on July 5, 2017.",
				"auto_resolve_timeout": 14400,
				"acknowledgement_timeout": 1800,
				"created_at": "2017-07-05T17:33:09Z",
				"status": "critical",
				"last_incident_timestamp": "2017-09-26T15:14:36Z",
				"teams": [
				  {
					"id": "P4SI59S",
					"type": "team_reference",
					"summary": "Engineering",
					"self": "https://api.pagerduty.com/teams/P4SI59S",
					"html_url": "https://webdemo.pagerduty.com/teams/P4SI59S"
				  }
				],
				"incident_urgency_rule": {
				  "type": "constant",
				  "urgency": "high"
				},
				"scheduled_actions": [],
				"support_hours": null,
				"escalation_policy": {
				  "id": "PINYWEF",
				  "type": "escalation_policy_reference",
				  "summary": "Default",
				  "self": "https://api.pagerduty.com/escalation_policies/PINYWEF",
				  "html_url": "https://webdemo.pagerduty.com/escalation_policies/PINYWEF"
				},
				"addons": [],
				"privilege": null,
				"alert_creation": "create_alerts_and_incidents",
				"integrations": [
				  {
					"id": "PUAYF96",
					"type": "generic_events_api_inbound_integration_reference",
					"summary": "API",
					"self": "https://api.pagerduty.com/services/PN49J75/integrations/PUAYF96",
					"html_url": "https://webdemo.pagerduty.com/services/PN49J75/integrations/PUAYF96"
				  },
				  {
					"id": "P90GZUH",
					"type": "generic_email_inbound_integration_reference",
					"summary": "Email",
					"self": "https://api.pagerduty.com/services/PN49J75/integrations/P90GZUH",
					"html_url": "https://webdemo.pagerduty.com/services/PN49J75/integrations/P90GZUH"
				  }
				],
				"metadata": {},
				"type": "service",
				"summary": "Production XDB Cluster",
				"self": "https://api.pagerduty.com/services/PN49J75",
				"html_url": "https://webdemo.pagerduty.com/services/PN49J75"
			  },
			  "assignments": [
				{
				  "at": "2017-09-26T15:14:36Z",
				  "assignee": {
					"id": "P553OPV",
					"type": "user_reference",
					"summary": "Laura Haley",
					"self": "https://api.pagerduty.com/users/P553OPV",
					"html_url": "https://webdemo.pagerduty.com/users/P553OPV"
				  }
				}
			  ],
			  "acknowledgements": [],
			  "last_status_change_at": "2017-09-26T15:14:36Z",
			  "last_status_change_by": {
				"id": "PN49J75",
				"type": "service_reference",
				"summary": "Production XDB Cluster",
				"self": "https://api.pagerduty.com/services/PN49J75",
				"html_url": "https://webdemo.pagerduty.com/services/PN49J75"
			  },
			  "first_trigger_log_entry": {
				"id": "R2XGXEI3W0FHMSDXHDIBQGBQ5E",
				"type": "trigger_log_entry_reference",
				"summary": "Triggered through the website",
				"self": "https://api.pagerduty.com/log_entries/R2XGXEI3W0FHMSDXHDIBQGBQ5E",
				"html_url": "https://webdemo.pagerduty.com/incidents/PRORDTY/log_entries/R2XGXEI3W0FHMSDXHDIBQGBQ5E"
			  },
			  "escalation_policy": {
				"id": "PINYWEF",
				"type": "escalation_policy_reference",
				"summary": "Default",
				"self": "https://api.pagerduty.com/escalation_policies/PINYWEF",
				"html_url": "https://webdemo.pagerduty.com/escalation_policies/PINYWEF"
			  },
			  "privilege": null,
			  "teams": [
				{
				  "id": "P4SI59S",
				  "type": "team_reference",
				  "summary": "Engineering",
				  "self": "https://api.pagerduty.com/teams/P4SI59S",
				  "html_url": "https://webdemo.pagerduty.com/teams/P4SI59S"
				}
			  ],
			  "alert_counts": {
				"all": 0,
				"triggered": 0,
				"resolved": 0
			  },
			  "impacted_services": [
				{
				  "id": "PN49J75",
				  "type": "service_reference",
				  "summary": "Production XDB Cluster",
				  "self": "https://api.pagerduty.com/services/PN49J75",
				  "html_url": "https://webdemo.pagerduty.com/services/PN49J75"
				}
			  ],
			  "is_mergeable": true,
			  "basic_alert_grouping": null,
			  "alert_grouping": null,
			  "metadata": {},
			  "external_references": [],
			  "importance": null,
			  "incidents_responders": [],
			  "responder_requests": [],
			  "subscriber_requests": [],
			  "urgency": "high",
			  "id": "PRORDTY",
			  "type": "incident",
			  "summary": "[#33] My new incident",
			  "self": "https://api.pagerduty.com/incidents/PRORDTY",
			  "html_url": "https://webdemo.pagerduty.com/incidents/PRORDTY",
			  "alerts": [
				  {
					  "alert_key": "c24117fc42e44b44b4d6876190583378"
				  }
			  ]
			},
			"id": "69a7ced0-a2cd-11e7-a799-22000a15839c",
			"created_on": "2017-09-26T15:14:36Z"
		  }
		]
	  }`
	testIncidentTriggerPayload = Payload{
		Messages: []Message{
			Message{
				Event: "incident.trigger",
				LogEntries: []LogEntry{
					LogEntry{
						ID:        "R2XGXEI3W0FHMSDXHDIBQGBQ5E",
						Type:      "trigger_log_entry",
						Summary:   "Triggered through the website",
						Self:      "https://api.pagerduty.com/log_entries/R2XGXEI3W0FHMSDXHDIBQGBQ5E",
						HTMLURL:   "https://webdemo.pagerduty.com/incidents/PRORDTY/log_entries/R2XGXEI3W0FHMSDXHDIBQGBQ5E",
						CreatedAt: "2017-09-26T15:14:36Z",
						Agent: Agent{
							ID:      "P553OPV",
							Type:    "user_reference",
							Summary: "Laura Haley",
							Self:    "https://api.pagerduty.com/users/P553OPV",
							HTMLURL: "https://webdemo.pagerduty.com/users/P553OPV",
						},
						Channel: Channel{
							Type:    "web_trigger",
							Summary: "My new incident",
							Subject: "My new incident",
							Details: "Oh my gosh",
						},
						Service: Service{
							ID:      "PN49J75",
							Type:    "service_reference",
							Summary: "Production XDB Cluster",
							Self:    "https://api.pagerduty.com/services/PN49J75",
							HTMLURL: "https://webdemo.pagerduty.com/services/PN49J75",
						},
						Incident: Incident{
							ID:      "PRORDTY",
							Type:    "incident_reference",
							Summary: "[#33] My new incident",
							Self:    "https://api.pagerduty.com/incidents/PRORDTY",
							HTMLURL: "https://webdemo.pagerduty.com/incidents/PRORDTY",
						},
						Teams: []Team{
							Team{
								ID:      "P4SI59S",
								Type:    "team_reference",
								Summary: "Engineering",
								Self:    "https://api.pagerduty.com/teams/P4SI59S",
								HTMLURL: "https://webdemo.pagerduty.com/teams/P4SI59S",
							},
						},
						Contexts: []Context{},
						EventDetails: EventDetails{
							Description: "My new incident",
						},
					},
				},
				Webhook: Webhook{
					EndpointURL: "https://requestb.in/18ao6fs1",
					Name:        "V2 wabhook",
					Description: "",
					WebhookObject: WebhookObject{
						ID:      "PN49J75",
						Type:    "service_reference",
						Summary: "Production XDB Cluster",
						Self:    "https://api.pagerduty.com/services/PN49J75",
						HTMLURL: "https://webdemo.pagerduty.com/services/PN49J75",
					},
					Config: Config{},
					OutboundIntegration: OutboundIntegration{
						ID:      "PJFWPEP",
						Type:    "outbound_integration_reference",
						Summary: "Generic V2 Webhook",
						Self:    "https://api.pagerduty.com/outbound_integrations/PJFWPEP",
						HTMLURL: "",
					},
					AccountsAddon: "",
					ID:            "PKT9NNX",
					Type:          "webhook",
					Summary:       "V2 wabhook",
					Self:          "https://api.pagerduty.com/webhooks/PKT9NNX",
					HTMLURL:       "",
				},
				Incident: Incident{
					IncidentNumber: 33,
					Title:          "My new incident",
					Description:    "My new incident",
					CreatedAt:      "2017-09-26T15:14:36Z",
					Status:         "triggered",
					PendingActions: []PendingAction{
						PendingAction{
							Type: "escalate",
							At:   "2017-09-26T15:44:36Z",
						},
						PendingAction{
							Type: "resolve",
							At:   "2017-09-26T19:14:36Z",
						},
					},
					IncidentKey: "",
					Service: Service{
						ID:                     "PN49J75",
						Name:                   "Production XDB Cluster",
						Description:            "This service was created during onboarding on July 5, 2017.",
						AutoResolveTimeout:     14400,
						AcknowledgementTimeout: 1800,
						CreatedAt:              "2017-07-05T17:33:09Z",
						Status:                 "critical",
						LastIncidentTimestamp:  "2017-09-26T15:14:36Z",
						Teams: []Team{
							Team{
								ID:      "P4SI59S",
								Type:    "team_reference",
								Summary: "Engineering",
								Self:    "https://api.pagerduty.com/teams/P4SI59S",
								HTMLURL: "https://webdemo.pagerduty.com/teams/P4SI59S",
							},
						},
						IncidentUrgencyRule: IncidentUrgencyRule{
							Type:    "constant",
							Urgency: "high",
						},
						ScheduledActions: []ScheduledAction{},
						SupportHours:     "",
						EscalationPolicy: EscalationPolicy{
							ID:      "PINYWEF",
							Type:    "escalation_policy_reference",
							Summary: "Default",
							Self:    "https://api.pagerduty.com/escalation_policies/PINYWEF",
							HTMLURL: "https://webdemo.pagerduty.com/escalation_policies/PINYWEF",
						},
						Addons:        []Addon{},
						Privilege:     "",
						AlertCreation: "create_alerts_and_incidents",
						Integrations: []Integration{
							Integration{
								ID:      "PUAYF96",
								Type:    "generic_events_api_inbound_integration_reference",
								Summary: "API",
								Self:    "https://api.pagerduty.com/services/PN49J75/integrations/PUAYF96",
								HTMLURL: "https://webdemo.pagerduty.com/services/PN49J75/integrations/PUAYF96",
							},
							Integration{
								ID:      "P90GZUH",
								Type:    "generic_email_inbound_integration_reference",
								Summary: "Email",
								Self:    "https://api.pagerduty.com/services/PN49J75/integrations/P90GZUH",
								HTMLURL: "https://webdemo.pagerduty.com/services/PN49J75/integrations/P90GZUH",
							},
						},
						Metadata: Metadata{},
						Type:     "service",
						Summary:  "Production XDB Cluster",
						Self:     "https://api.pagerduty.com/services/PN49J75",
						HTMLURL:  "https://webdemo.pagerduty.com/services/PN49J75",
					},
					Assignments: []Assignment{
						Assignment{
							At: "2017-09-26T15:14:36Z",
							Assignee: Assignee{
								ID:      "P553OPV",
								Type:    "user_reference",
								Summary: "Laura Haley",
								Self:    "https://api.pagerduty.com/users/P553OPV",
								HTMLURL: "https://webdemo.pagerduty.com/users/P553OPV",
							},
						},
					},
					Acknowledgements:   []Acknowledgement{},
					LastStatusChangeAt: "2017-09-26T15:14:36Z",
					LastStatusChangeBy: LastStatusChangeBy{
						ID:      "PN49J75",
						Type:    "service_reference",
						Summary: "Production XDB Cluster",
						Self:    "https://api.pagerduty.com/services/PN49J75",
						HTMLURL: "https://webdemo.pagerduty.com/services/PN49J75",
					},
					FirstTriggerLogEntry: FirstTriggerLogEntry{
						ID:      "R2XGXEI3W0FHMSDXHDIBQGBQ5E",
						Type:    "trigger_log_entry_reference",
						Summary: "Triggered through the website",
						Self:    "https://api.pagerduty.com/log_entries/R2XGXEI3W0FHMSDXHDIBQGBQ5E",
						HTMLURL: "https://webdemo.pagerduty.com/incidents/PRORDTY/log_entries/R2XGXEI3W0FHMSDXHDIBQGBQ5E",
					},
					EscalationPolicy: EscalationPolicy{
						ID:      "PINYWEF",
						Type:    "escalation_policy_reference",
						Summary: "Default",
						Self:    "https://api.pagerduty.com/escalation_policies/PINYWEF",
						HTMLURL: "https://webdemo.pagerduty.com/escalation_policies/PINYWEF",
					},
					Privilege: "",
					Teams: []Team{
						Team{
							ID:      "P4SI59S",
							Type:    "team_reference",
							Summary: "Engineering",
							Self:    "https://api.pagerduty.com/teams/P4SI59S",
							HTMLURL: "https://webdemo.pagerduty.com/teams/P4SI59S",
						},
					},
					AlertCounts: AlertCounts{
						All:       0,
						Triggered: 0,
						Resolved:  0,
					},
					ImpactedServices: []ImpactedService{
						ImpactedService{
							ID:      "PN49J75",
							Type:    "service_reference",
							Summary: "Production XDB Cluster",
							Self:    "https://api.pagerduty.com/services/PN49J75",
							HTMLURL: "https://webdemo.pagerduty.com/services/PN49J75",
						},
					},
					IsMergeable:         true,
					BasicAlertGrouping:  "",
					Metadata:            Metadata{},
					ExternalReferences:  []ExternalReference{},
					Importance:          "",
					IncidentsResponders: []IncidentsResponder{},
					ResponderRequests:   []ResponderRequest{},
					SubscriberRequests:  []SubscriberRequest{},
					Urgency:             "high",
					ID:                  "PRORDTY",
					Type:                "incident",
					Summary:             "[#33] My new incident",
					Self:                "https://api.pagerduty.com/incidents/PRORDTY",
					HTMLURL:             "https://webdemo.pagerduty.com/incidents/PRORDTY",
					Alerts: []Alert{
						Alert{
							AlertKey: "c24117fc42e44b44b4d6876190583378",
						},
					},
				},
				ID:        "69a7ced0-a2cd-11e7-a799-22000a15839c",
				CreatedOn: "2017-09-26T15:14:36Z",
			},
		},
	}
)

/*
	TestPayload
*/

func TestPayload(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		if reflect.DeepEqual(&payload, &tc.want) == false {
			t.Fatalf("payload failed => testCase: %s", n)
		}
	}
}

/*
	TestMessages
*/

func TestMessages(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		if reflect.DeepEqual(&payload.Messages, &tc.want.Messages) == false {
			t.Fatalf("payload failed => testCase: %s", n)
		}
	}
}

/*
	TestMessage
*/

func TestMessage(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i], &tc.want.Messages[i]) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

/*
	TestIncident
*/

func TestGetCreatedOn(t *testing.T) {
	type TestCase struct {
		message Message
		pass    bool
	}

	testCases := map[string]TestCase{
		"ok": TestCase{
			message: Message{
				CreatedOn: "2017-09-26T15:14:36Z",
			},
			pass: true,
		},
		"parse error": TestCase{
			message: Message{
				CreatedOn: "keke-is-engineer",
			},
			pass: false,
		},
	}

	for n, tc := range testCases {
		_, err := tc.message.GetCreatedOn()
		if err != nil {
			if tc.pass {
				t.Fatalf("getCreatedOn failed => testCase:%s, error:%v", n, err)
			}
		}
	}
}

func TestIncident(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident, &tc.want.Messages[i].Incident) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}
func TestIncidentAlerts(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Alerts, &tc.want.Messages[i].Incident.Alerts) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentPendingActions(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.PendingActions, &tc.want.Messages[i].Incident.PendingActions) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentService(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Service, &tc.want.Messages[i].Incident.Service) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

/*
	TestIncidentService
*/

func TestIncidentServiceAddons(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Service.Addons, &tc.want.Messages[i].Incident.Service.Addons) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentServiceIntegrations(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Service.Integrations, &tc.want.Messages[i].Incident.Service.Integrations) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentServiceScheduledActions(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Service.ScheduledActions, &tc.want.Messages[i].Incident.Service.ScheduledActions) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentServiceIncidentUrgencyRule(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Service.IncidentUrgencyRule, &tc.want.Messages[i].Incident.Service.IncidentUrgencyRule) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentServiceEscalationPolicy(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Service.EscalationPolicy, &tc.want.Messages[i].Incident.Service.EscalationPolicy) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentServiceMetadata(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Service.Metadata, &tc.want.Messages[i].Incident.Service.Metadata) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentAssignments(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Assignments, &tc.want.Messages[i].Incident.Assignments) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentAcknowledgements(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Acknowledgements, &tc.want.Messages[i].Incident.Acknowledgements) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentLastStatusChangeBy(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.LastStatusChangeBy, &tc.want.Messages[i].Incident.LastStatusChangeBy) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentFirstTriggerLogEntry(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.FirstTriggerLogEntry, &tc.want.Messages[i].Incident.FirstTriggerLogEntry) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentEscalationPolicy(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.EscalationPolicy, &tc.want.Messages[i].Incident.EscalationPolicy) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentTeams(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Teams, &tc.want.Messages[i].Incident.Teams) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentPriority(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Priority, &tc.want.Messages[i].Incident.Priority) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentAlertCounts(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.AlertCounts, &tc.want.Messages[i].Incident.AlertCounts) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentMetadata(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.Metadata, &tc.want.Messages[i].Incident.Metadata) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentImpactedService(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.ImpactedServices, &tc.want.Messages[i].Incident.ImpactedServices) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentExternalReference(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.ExternalReferences, &tc.want.Messages[i].Incident.ExternalReferences) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentIncidentsResponders(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.IncidentsResponders, &tc.want.Messages[i].Incident.IncidentsResponders) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentResponderRequests(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.ResponderRequests, &tc.want.Messages[i].Incident.ResponderRequests) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestIncidentSubscriberRequests(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Incident.SubscriberRequests, &tc.want.Messages[i].Incident.SubscriberRequests) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

/*
	TestWebhook
*/

func TestWebhook(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Webhook, &tc.want.Messages[i].Webhook) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestWebhookWebhookObject(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Webhook.WebhookObject, &tc.want.Messages[i].Webhook.WebhookObject) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestWebhookConfig(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Webhook.Config, &tc.want.Messages[i].Webhook.Config) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestWebhookOutboundIntegration(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].Webhook.OutboundIntegration, &tc.want.Messages[i].Webhook.OutboundIntegration) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

/*
	TestLogEntries
*/

func TestLogEntries(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries, &tc.want.Messages[i].LogEntries) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

/*
	TestLogEntry
*/

func TestLogEntry(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0], &tc.want.Messages[i].LogEntries[0]) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryAgent(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Agent, &tc.want.Messages[i].LogEntries[0].Agent) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryChannel(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Channel, &tc.want.Messages[i].LogEntries[0].Channel) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryContexts(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Contexts, &tc.want.Messages[i].LogEntries[0].Contexts) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncident(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident, &tc.want.Messages[i].LogEntries[0].Incident) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

/*
	TestLogEntryIncident
*/

func TestLogEntryIncidentAlerts(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.Alerts, &tc.want.Messages[i].LogEntries[0].Incident.Alerts) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentPendingActions(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.PendingActions, &tc.want.Messages[i].LogEntries[0].Incident.PendingActions) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentService(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.Service, &tc.want.Messages[i].LogEntries[0].Incident.Service) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentAssignments(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.Assignments, &tc.want.Messages[i].LogEntries[0].Incident.Assignments) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentAcknowledgements(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.Acknowledgements, &tc.want.Messages[i].LogEntries[0].Incident.Acknowledgements) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentLastStatusChangeBy(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.LastStatusChangeBy, &tc.want.Messages[i].LogEntries[0].Incident.LastStatusChangeBy) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentFirstTriggerLogEntry(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.FirstTriggerLogEntry, &tc.want.Messages[i].LogEntries[0].Incident.FirstTriggerLogEntry) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentEscalationPolicy(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.EscalationPolicy, &tc.want.Messages[i].LogEntries[0].Incident.EscalationPolicy) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentTeams(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.Teams, &tc.want.Messages[i].LogEntries[0].Incident.Teams) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentPriority(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.Priority, &tc.want.Messages[i].LogEntries[0].Incident.Priority) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentAlertCounts(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.AlertCounts, &tc.want.Messages[i].LogEntries[0].Incident.AlertCounts) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentMetadata(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.Metadata, &tc.want.Messages[i].LogEntries[0].Incident.Metadata) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentImpactedServices(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.ImpactedServices, &tc.want.Messages[i].LogEntries[0].Incident.ImpactedServices) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentExternalReferences(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.ExternalReferences, &tc.want.Messages[i].LogEntries[0].Incident.ExternalReferences) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentIncidentsResponders(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.IncidentsResponders, &tc.want.Messages[i].LogEntries[0].Incident.IncidentsResponders) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentResponderRequests(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.ResponderRequests, &tc.want.Messages[i].LogEntries[0].Incident.ResponderRequests) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryIncidentSubscriberRequests(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Incident.SubscriberRequests, &tc.want.Messages[i].LogEntries[0].Incident.SubscriberRequests) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryService(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Service, &tc.want.Messages[i].LogEntries[0].Service) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryTeams(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].Teams, &tc.want.Messages[i].LogEntries[0].Teams) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestLogEntryEventDetails(t *testing.T) {
	type Testcase struct {
		payload string
		want    Payload
	}

	tcs := map[string]Testcase{
		"incident.trigger": Testcase{
			payload: testIncidentTriggerPayloadStr,
			want:    testIncidentTriggerPayload,
		},
	}

	for n, tc := range tcs {
		var payload Payload
		err := json.Unmarshal([]byte(tc.payload), &payload)

		if err != nil {
			t.Fatalf("payload failed => testCase: %s, error: %v", n, err)
		}

		for i := range payload.Messages {
			if reflect.DeepEqual(&payload.Messages[i].LogEntries[0].EventDetails, &tc.want.Messages[i].LogEntries[0].EventDetails) == false {
				t.Fatalf("payload failed => testCase: %s", n)
			}
		}
	}
}

func TestGetCreatedAt(t *testing.T) {
	format := "2006-01-02T15:04:05Z"
	validTime, _ := time.Parse(format, "2017-09-26T15:14:36Z")

	type Testcase struct {
		incident *Incident
		want     time.Time
	}

	tcs := map[string]Testcase{
		"okay": Testcase{
			incident: &Incident{
				CreatedAt: "2017-09-26T15:14:36Z",
			},
			want: validTime,
		},
	}

	for n, tc := range tcs {
		result, err := tc.incident.GetCreatedAt()
		if err != nil {
			t.Fatalf("GetCreatedAt failed => testCase: %s error: %v", n, err)
		}

		if result != tc.want {
			t.Fatalf("GetCreatedAt not match => testCase: %s want: %v, got: %v", n, result, tc.want)
		}
	}
}
