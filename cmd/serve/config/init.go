package config

import (
	"github.com/SquadcastHub/auxpkg/v2/env"
)

func Init(configFiles ...string) {
	env.LoadFromJSON(configFiles...)

	env.AddNewEnvEntryV2("HOST", &HostName)
	env.AddNewEnvEntryV2("PORT", &Port)
	env.AddNewEnvEntryV2("DATACENTER", &Datacenter)

	env.AddNewEnvEntryV2("PUBSUB_PROJECT_ID", &PubSubProjectId)
	env.AddNewEnvEntryV2("PUBSUB_EVENTS_TOPIC_NAME", &PubSubEventsTopicName)
	env.AddNewEnvEntryV2("PUBSUB_INCIDENTS_TOPIC_NAME", &PubSubIncidentsTopicName)
}
