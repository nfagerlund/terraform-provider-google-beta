// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPubsubSubscription_pubsubSubscriptionPullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPubsubSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubSubscription_pubsubSubscriptionPullExample(context),
			},
			{
				ResourceName:      "google_pubsub_subscription.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPubsubSubscription_pubsubSubscriptionPullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"
}

resource "google_pubsub_subscription" "example" {
  name  = "tf-test-example-subscription%{random_suffix}"
  topic = google_pubsub_topic.example.name

  labels = {
    foo = "bar"
  }

  # 20 minutes
  message_retention_duration = "1200s"
  retain_acked_messages      = true

  ack_deadline_seconds = 20

  expiration_policy {
    ttl = "300000.5s"
  }
}
`, context)
}

func TestAccPubsubSubscription_pubsubSubscriptionDeadLetterExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPubsubSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubSubscription_pubsubSubscriptionDeadLetterExample(context),
			},
			{
				ResourceName:      "google_pubsub_subscription.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPubsubSubscription_pubsubSubscriptionDeadLetterExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"
}

resource "google_pubsub_topic" "example_dead_letter" {
  name = "tf-test-example-topic%{random_suffix}-dead-letter"
}

resource "google_pubsub_subscription" "example" {
  name  = "tf-test-example-subscription%{random_suffix}"
  topic = google_pubsub_topic.example.name

  dead_letter_policy {
    dead_letter_topic = google_pubsub_topic.example_dead_letter.id
    max_delivery_attempts = 10
  }
}
`, context)
}

func testAccCheckPubsubSubscriptionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_pubsub_subscription" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{PubsubBasePath}}projects/{{project}}/subscriptions/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("PubsubSubscription still exists at %s", url)
			}
		}

		return nil
	}
}
