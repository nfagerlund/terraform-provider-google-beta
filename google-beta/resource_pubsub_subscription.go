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
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/googleapi"
)

func comparePubsubSubscriptionExpirationPolicy(_, old, new string, _ *schema.ResourceData) bool {
	trimmedNew := strings.TrimLeft(new, "0")
	trimmedOld := strings.TrimLeft(old, "0")
	if strings.Contains(trimmedNew, ".") {
		trimmedNew = strings.TrimRight(strings.TrimSuffix(trimmedNew, "s"), "0") + "s"
	}
	if strings.Contains(trimmedOld, ".") {
		trimmedOld = strings.TrimRight(strings.TrimSuffix(trimmedOld, "s"), "0") + "s"
	}
	return trimmedNew == trimmedOld
}

func resourcePubsubSubscription() *schema.Resource {
	return &schema.Resource{
		Create: resourcePubsubSubscriptionCreate,
		Read:   resourcePubsubSubscriptionRead,
		Update: resourcePubsubSubscriptionUpdate,
		Delete: resourcePubsubSubscriptionDelete,

		Importer: &schema.ResourceImporter{
			State: resourcePubsubSubscriptionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Update: schema.DefaultTimeout(6 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the subscription.`,
			},
			"topic": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `A reference to a Topic resource.`,
			},
			"ack_deadline_seconds": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				Description: `This value is the maximum time after a subscriber receives a message
before the subscriber should acknowledge the message. After message
delivery but before the ack deadline expires and before the message is
acknowledged, it is an outstanding message and will not be delivered
again during that time (on a best-effort basis).

For pull subscriptions, this value is used as the initial value for
the ack deadline. To override this value for a given message, call
subscriptions.modifyAckDeadline with the corresponding ackId if using
pull. The minimum custom deadline you can specify is 10 seconds. The
maximum custom deadline you can specify is 600 seconds (10 minutes).
If this parameter is 0, a default value of 10 seconds is used.

For push delivery, this value is also used to set the request timeout
for the call to the push endpoint.

If the subscriber never acknowledges the message, the Pub/Sub system
will eventually redeliver the message.`,
			},
			"dead_letter_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A policy that specifies the conditions for dead lettering messages in
this subscription. If dead_letter_policy is not set, dead lettering
is disabled.

The Cloud Pub/Sub service account associated with this subscriptions's
parent project (i.e.,
service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
permission to Acknowledge() messages on this subscription.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dead_letter_topic": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The name of the topic to which dead letter messages should be published.
Format is 'projects/{project}/topics/{topic}'.

The Cloud Pub/Sub service\naccount associated with the enclosing subscription's
parent project (i.e., 
service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
permission to Publish() to this topic.

The operation will fail if the topic does not exist.
Users should ensure that there is a subscription attached to this topic
since messages published to a topic with no subscriptions are lost.`,
						},
						"max_delivery_attempts": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `The maximum number of delivery attempts for any message. The value must be
between 5 and 100.

The number of delivery attempts is defined as 1 + (the sum of number of 
NACKs and number of times the acknowledgement deadline has been exceeded for the message).

A NACK is any call to ModifyAckDeadline with a 0 deadline. Note that
client libraries may automatically extend ack_deadlines.

This field will be honored on a best effort basis.

If this parameter is 0, a default value of 5 is used.`,
						},
					},
				},
			},
			"expiration_policy": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Description: `A policy that specifies the conditions for this subscription's expiration.
A subscription is considered active as long as any connected subscriber
is successfully consuming messages from the subscription or is issuing
operations on the subscription. If expirationPolicy is not set, a default
policy with ttl of 31 days will be used.  If it is set but ttl is "", the
resource never expires.  The minimum allowed value for expirationPolicy.ttl
is 1 day.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ttl": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: comparePubsubSubscriptionExpirationPolicy,
							Description: `Specifies the "time-to-live" duration for an associated resource. The
resource expires if it is not active for a period of ttl.
If ttl is not set, the associated resource never expires.
A duration in seconds with up to nine fractional digits, terminated by 's'.
Example - "3.5s".`,
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `A set of key/value label pairs to assign to this Subscription.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"message_retention_duration": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `How long to retain unacknowledged messages in the subscription's
backlog, from the moment a message is published. If
retainAckedMessages is true, then this also configures the retention
of acknowledged messages, and thus configures how far back in time a
subscriptions.seek can be done. Defaults to 7 days. Cannot be more
than 7 days ('"604800s"') or less than 10 minutes ('"600s"').

A duration in seconds with up to nine fractional digits, terminated
by 's'. Example: '"600.5s"'.`,
				Default: "604800s",
			},
			"push_config": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `If push delivery is used with this subscription, this field is used to
configure it. An empty pushConfig signifies that the subscriber will
pull and ack messages using API methods.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"push_endpoint": {
							Type:     schema.TypeString,
							Required: true,
							Description: `A URL locating the endpoint to which messages should be pushed.
For example, a Webhook endpoint might use
"https://example.com/push".`,
						},
						"attributes": {
							Type:     schema.TypeMap,
							Optional: true,
							Description: `Endpoint configuration attributes.

Every endpoint has a set of API supported attributes that can
be used to control different aspects of the message delivery.

The currently supported attribute is x-goog-version, which you
can use to change the format of the pushed message. This
attribute indicates the version of the data expected by
the endpoint. This controls the shape of the pushed message
(i.e., its fields and metadata). The endpoint version is
based on the version of the Pub/Sub API.

If not present during the subscriptions.create call,
it will default to the version of the API used to make
such call. If not present during a subscriptions.modifyPushConfig
call, its value will not be changed. subscriptions.get
calls will always return a valid version, even if the
subscription was created without this attribute.

The possible values for this attribute are:

- v1beta1: uses the push format defined in the v1beta1 Pub/Sub API.
- v1 or v1beta2: uses the push format defined in the v1 Pub/Sub API.`,
							Elem: &schema.Schema{Type: schema.TypeString},
						},
						"oidc_token": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `If specified, Pub/Sub will generate and attach an OIDC JWT token as
an Authorization header in the HTTP request for every pushed message.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_account_email": {
										Type:     schema.TypeString,
										Required: true,
										Description: `Service account email to be used for generating the OIDC token.
The caller (for subscriptions.create, subscriptions.patch, and
subscriptions.modifyPushConfig RPCs) must have the
iam.serviceAccounts.actAs permission for the service account.`,
									},
									"audience": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `Audience to be used when generating OIDC token. The audience claim
identifies the recipients that the JWT is intended for. The audience
value is a single case-sensitive string. Having multiple values (array)
for the audience field is not supported. More info about the OIDC JWT
token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
Note: if not specified, the Push endpoint URL will be used.`,
									},
								},
							},
						},
					},
				},
			},
			"retain_acked_messages": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Indicates whether to retain acknowledged messages. If 'true', then
messages are not expunged from the subscription's backlog, even if
they are acknowledged, until they fall out of the
messageRetentionDuration window.`,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePubsubSubscriptionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandPubsubSubscriptionName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	topicProp, err := expandPubsubSubscriptionTopic(d.Get("topic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("topic"); !isEmptyValue(reflect.ValueOf(topicProp)) && (ok || !reflect.DeepEqual(v, topicProp)) {
		obj["topic"] = topicProp
	}
	labelsProp, err := expandPubsubSubscriptionLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	pushConfigProp, err := expandPubsubSubscriptionPushConfig(d.Get("push_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("push_config"); !isEmptyValue(reflect.ValueOf(pushConfigProp)) && (ok || !reflect.DeepEqual(v, pushConfigProp)) {
		obj["pushConfig"] = pushConfigProp
	}
	ackDeadlineSecondsProp, err := expandPubsubSubscriptionAckDeadlineSeconds(d.Get("ack_deadline_seconds"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ack_deadline_seconds"); !isEmptyValue(reflect.ValueOf(ackDeadlineSecondsProp)) && (ok || !reflect.DeepEqual(v, ackDeadlineSecondsProp)) {
		obj["ackDeadlineSeconds"] = ackDeadlineSecondsProp
	}
	messageRetentionDurationProp, err := expandPubsubSubscriptionMessageRetentionDuration(d.Get("message_retention_duration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("message_retention_duration"); !isEmptyValue(reflect.ValueOf(messageRetentionDurationProp)) && (ok || !reflect.DeepEqual(v, messageRetentionDurationProp)) {
		obj["messageRetentionDuration"] = messageRetentionDurationProp
	}
	retainAckedMessagesProp, err := expandPubsubSubscriptionRetainAckedMessages(d.Get("retain_acked_messages"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retain_acked_messages"); !isEmptyValue(reflect.ValueOf(retainAckedMessagesProp)) && (ok || !reflect.DeepEqual(v, retainAckedMessagesProp)) {
		obj["retainAckedMessages"] = retainAckedMessagesProp
	}
	expirationPolicyProp, err := expandPubsubSubscriptionExpirationPolicy(d.Get("expiration_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expiration_policy"); ok || !reflect.DeepEqual(v, expirationPolicyProp) {
		obj["expirationPolicy"] = expirationPolicyProp
	}
	deadLetterPolicyProp, err := expandPubsubSubscriptionDeadLetterPolicy(d.Get("dead_letter_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dead_letter_policy"); ok || !reflect.DeepEqual(v, deadLetterPolicyProp) {
		obj["deadLetterPolicy"] = deadLetterPolicyProp
	}

	obj, err = resourcePubsubSubscriptionEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Subscription: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Subscription: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = PollingWaitTime(resourcePubsubSubscriptionPollRead(d, meta), PollCheckForExistence, "Creating Subscription", d.Timeout(schema.TimeoutCreate))
	if err != nil {
		log.Printf("[ERROR] Unable to confirm eventually consistent Subscription %q finished updating: %q", d.Id(), err)
	}

	log.Printf("[DEBUG] Finished creating Subscription %q: %#v", d.Id(), res)

	return resourcePubsubSubscriptionRead(d, meta)
}

func resourcePubsubSubscriptionPollRead(d *schema.ResourceData, meta interface{}) PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*Config)

		url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/subscriptions/{{name}}")
		if err != nil {
			return nil, err
		}

		project, err := getProject(d, config)
		if err != nil {
			return nil, err
		}
		res, err := sendRequest(config, "GET", project, url, nil)
		if err != nil {
			return res, err
		}
		res, err = resourcePubsubSubscriptionDecoder(d, meta, res)
		if err != nil {
			return nil, err
		}
		if res == nil {
			// Decoded object not found, spoof a 404 error for poll
			return nil, &googleapi.Error{
				Code:    404,
				Message: "could not find object PubsubSubscription",
			}
		}

		return res, nil
	}
}

func resourcePubsubSubscriptionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("PubsubSubscription %q", d.Id()))
	}

	res, err = resourcePubsubSubscriptionDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing PubsubSubscription because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}

	if err := d.Set("name", flattenPubsubSubscriptionName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("topic", flattenPubsubSubscriptionTopic(res["topic"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("labels", flattenPubsubSubscriptionLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("push_config", flattenPubsubSubscriptionPushConfig(res["pushConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("ack_deadline_seconds", flattenPubsubSubscriptionAckDeadlineSeconds(res["ackDeadlineSeconds"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("message_retention_duration", flattenPubsubSubscriptionMessageRetentionDuration(res["messageRetentionDuration"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("retain_acked_messages", flattenPubsubSubscriptionRetainAckedMessages(res["retainAckedMessages"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("expiration_policy", flattenPubsubSubscriptionExpirationPolicy(res["expirationPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("dead_letter_policy", flattenPubsubSubscriptionDeadLetterPolicy(res["deadLetterPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}

	return nil
}

func resourcePubsubSubscriptionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandPubsubSubscriptionLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	pushConfigProp, err := expandPubsubSubscriptionPushConfig(d.Get("push_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("push_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, pushConfigProp)) {
		obj["pushConfig"] = pushConfigProp
	}
	ackDeadlineSecondsProp, err := expandPubsubSubscriptionAckDeadlineSeconds(d.Get("ack_deadline_seconds"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ack_deadline_seconds"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, ackDeadlineSecondsProp)) {
		obj["ackDeadlineSeconds"] = ackDeadlineSecondsProp
	}
	messageRetentionDurationProp, err := expandPubsubSubscriptionMessageRetentionDuration(d.Get("message_retention_duration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("message_retention_duration"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, messageRetentionDurationProp)) {
		obj["messageRetentionDuration"] = messageRetentionDurationProp
	}
	retainAckedMessagesProp, err := expandPubsubSubscriptionRetainAckedMessages(d.Get("retain_acked_messages"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retain_acked_messages"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, retainAckedMessagesProp)) {
		obj["retainAckedMessages"] = retainAckedMessagesProp
	}
	expirationPolicyProp, err := expandPubsubSubscriptionExpirationPolicy(d.Get("expiration_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expiration_policy"); ok || !reflect.DeepEqual(v, expirationPolicyProp) {
		obj["expirationPolicy"] = expirationPolicyProp
	}
	deadLetterPolicyProp, err := expandPubsubSubscriptionDeadLetterPolicy(d.Get("dead_letter_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dead_letter_policy"); ok || !reflect.DeepEqual(v, deadLetterPolicyProp) {
		obj["deadLetterPolicy"] = deadLetterPolicyProp
	}

	obj, err = resourcePubsubSubscriptionUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Subscription %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("push_config") {
		updateMask = append(updateMask, "pushConfig")
	}

	if d.HasChange("ack_deadline_seconds") {
		updateMask = append(updateMask, "ackDeadlineSeconds")
	}

	if d.HasChange("message_retention_duration") {
		updateMask = append(updateMask, "messageRetentionDuration")
	}

	if d.HasChange("retain_acked_messages") {
		updateMask = append(updateMask, "retainAckedMessages")
	}

	if d.HasChange("expiration_policy") {
		updateMask = append(updateMask, "expirationPolicy")
	}

	if d.HasChange("dead_letter_policy") {
		updateMask = append(updateMask, "deadLetterPolicy")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Subscription %q: %s", d.Id(), err)
	}

	return resourcePubsubSubscriptionRead(d, meta)
}

func resourcePubsubSubscriptionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Subscription %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Subscription")
	}

	log.Printf("[DEBUG] Finished deleting Subscription %q: %#v", d.Id(), res)
	return nil
}

func resourcePubsubSubscriptionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/subscriptions/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenPubsubSubscriptionName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenPubsubSubscriptionTopic(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenPubsubSubscriptionLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenPubsubSubscriptionPushConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["oidc_token"] =
		flattenPubsubSubscriptionPushConfigOidcToken(original["oidcToken"], d, config)
	transformed["push_endpoint"] =
		flattenPubsubSubscriptionPushConfigPushEndpoint(original["pushEndpoint"], d, config)
	transformed["attributes"] =
		flattenPubsubSubscriptionPushConfigAttributes(original["attributes"], d, config)
	return []interface{}{transformed}
}
func flattenPubsubSubscriptionPushConfigOidcToken(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["service_account_email"] =
		flattenPubsubSubscriptionPushConfigOidcTokenServiceAccountEmail(original["serviceAccountEmail"], d, config)
	transformed["audience"] =
		flattenPubsubSubscriptionPushConfigOidcTokenAudience(original["audience"], d, config)
	return []interface{}{transformed}
}
func flattenPubsubSubscriptionPushConfigOidcTokenServiceAccountEmail(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenPubsubSubscriptionPushConfigOidcTokenAudience(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenPubsubSubscriptionPushConfigPushEndpoint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenPubsubSubscriptionPushConfigAttributes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenPubsubSubscriptionAckDeadlineSeconds(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenPubsubSubscriptionMessageRetentionDuration(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenPubsubSubscriptionRetainAckedMessages(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenPubsubSubscriptionExpirationPolicy(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["ttl"] =
		flattenPubsubSubscriptionExpirationPolicyTtl(original["ttl"], d, config)
	return []interface{}{transformed}
}
func flattenPubsubSubscriptionExpirationPolicyTtl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenPubsubSubscriptionDeadLetterPolicy(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dead_letter_topic"] =
		flattenPubsubSubscriptionDeadLetterPolicyDeadLetterTopic(original["deadLetterTopic"], d, config)
	transformed["max_delivery_attempts"] =
		flattenPubsubSubscriptionDeadLetterPolicyMaxDeliveryAttempts(original["maxDeliveryAttempts"], d, config)
	return []interface{}{transformed}
}
func flattenPubsubSubscriptionDeadLetterPolicyDeadLetterTopic(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenPubsubSubscriptionDeadLetterPolicyMaxDeliveryAttempts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandPubsubSubscriptionName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/subscriptions/{{name}}")
}

func expandPubsubSubscriptionTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	project, err := getProject(d, config)
	if err != nil {
		return "", err
	}

	topic := d.Get("topic").(string)

	re := regexp.MustCompile(`projects\/(.*)\/topics\/(.*)`)
	match := re.FindStringSubmatch(topic)
	if len(match) == 3 {
		return topic, nil
	} else {
		// If no full topic given, we expand it to a full topic on the same project
		fullTopic := fmt.Sprintf("projects/%s/topics/%s", project, topic)
		d.Set("topic", fullTopic)
		return fullTopic, nil
	}
}

func expandPubsubSubscriptionLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandPubsubSubscriptionPushConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedOidcToken, err := expandPubsubSubscriptionPushConfigOidcToken(original["oidc_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOidcToken); val.IsValid() && !isEmptyValue(val) {
		transformed["oidcToken"] = transformedOidcToken
	}

	transformedPushEndpoint, err := expandPubsubSubscriptionPushConfigPushEndpoint(original["push_endpoint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPushEndpoint); val.IsValid() && !isEmptyValue(val) {
		transformed["pushEndpoint"] = transformedPushEndpoint
	}

	transformedAttributes, err := expandPubsubSubscriptionPushConfigAttributes(original["attributes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAttributes); val.IsValid() && !isEmptyValue(val) {
		transformed["attributes"] = transformedAttributes
	}

	return transformed, nil
}

func expandPubsubSubscriptionPushConfigOidcToken(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceAccountEmail, err := expandPubsubSubscriptionPushConfigOidcTokenServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !isEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	transformedAudience, err := expandPubsubSubscriptionPushConfigOidcTokenAudience(original["audience"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAudience); val.IsValid() && !isEmptyValue(val) {
		transformed["audience"] = transformedAudience
	}

	return transformed, nil
}

func expandPubsubSubscriptionPushConfigOidcTokenServiceAccountEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionPushConfigOidcTokenAudience(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionPushConfigPushEndpoint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionPushConfigAttributes(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandPubsubSubscriptionAckDeadlineSeconds(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionMessageRetentionDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionRetainAckedMessages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionExpirationPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTtl, err := expandPubsubSubscriptionExpirationPolicyTtl(original["ttl"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTtl); val.IsValid() && !isEmptyValue(val) {
		transformed["ttl"] = transformedTtl
	}

	return transformed, nil
}

func expandPubsubSubscriptionExpirationPolicyTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionDeadLetterPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDeadLetterTopic, err := expandPubsubSubscriptionDeadLetterPolicyDeadLetterTopic(original["dead_letter_topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDeadLetterTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["deadLetterTopic"] = transformedDeadLetterTopic
	}

	transformedMaxDeliveryAttempts, err := expandPubsubSubscriptionDeadLetterPolicyMaxDeliveryAttempts(original["max_delivery_attempts"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDeliveryAttempts); val.IsValid() && !isEmptyValue(val) {
		transformed["maxDeliveryAttempts"] = transformedMaxDeliveryAttempts
	}

	return transformed, nil
}

func expandPubsubSubscriptionDeadLetterPolicyDeadLetterTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionDeadLetterPolicyMaxDeliveryAttempts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourcePubsubSubscriptionEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "name")
	return obj, nil
}

func resourcePubsubSubscriptionUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	newObj := make(map[string]interface{})
	newObj["subscription"] = obj
	return newObj, nil
}

func resourcePubsubSubscriptionDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {

	// path is a derived field from the API-side `name`
	path := fmt.Sprintf("projects/%s/subscriptions/%s", d.Get("project"), d.Get("name"))
	d.Set("path", path)

	return res, nil
}
