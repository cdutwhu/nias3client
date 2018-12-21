package xjy

import (
	"io/ioutil"
	"testing"
)

func TestJstr2Y(t *testing.T) {

	ystr := Jstr2Y(`{
		"id":"7ccd3322-e1a5-411a-a67d-6a735c76f119",
		"timestamp": "2015-12-18T12:17:00+00:00",
		"actor":{
			"objectType": "Agent",
			"name":"Example Learner",
			"mbox":"mailto:example.learner@adlnet.gov"
		},
		"verb":{
			"id":"http://adlnet.gov/expapi/verbs/attempted",
			"display":{
				"en-US":"attempted"
			}
		},
		"object":{
			"id":"http://example.adlnet.gov/xapi/example/simpleCBT",
			"definition":{
				"name":{
					"en-US":"simple CBT course"
				},
				"description":{
					"en-US":"A fictitious example CBT course."
				}
			}
		},
		"result":{
			"score":{
				"scaled":0.95
			},
			"success":true,
			"completion":true,
			"duration": "PT1234S"
		}
	}`)

	ioutil.WriteFile("test.yaml", []byte(ystr), 0666)
}
