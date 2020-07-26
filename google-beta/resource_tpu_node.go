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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// compareTpuNodeSchedulingConfig diff suppresses for the default
// scheduling, i.e. if preemptible is false, the API may either return no
// schedulingConfig or an empty schedulingConfig.
func compareTpuNodeSchedulingConfig(k, old, new string, d *schema.ResourceData) bool {
	if k == "scheduling_config.0.preemptible" {
		return old == "" && new == "false"
	}
	if k == "scheduling_config.#" {
		o, n := d.GetChange("scheduling_config.0.preemptible")
		return o.(bool) == n.(bool)
	}
	return false
}

func validateHttpHeaders() schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		headers := i.(map[string]interface{})
		if _, ok := headers["Content-Length"]; ok {
			es = append(es, fmt.Errorf("Cannot set the Content-Length header on %s", k))
			return
		}
		r := regexp.MustCompile(`(X-Google-|X-AppEngine-).*`)
		for key := range headers {
			if r.MatchString(key) {
				es = append(es, fmt.Errorf("Cannot set the %s header on %s", key, k))
				return
			}
		}

		return
	}
}

func resourceTPUNode() *schema.Resource {
	return &schema.Resource{
		Create: resourceTPUNodeCreate,
		Read:   resourceTPUNodeRead,
		Update: resourceTPUNodeUpdate,
		Delete: resourceTPUNodeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceTPUNodeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Update: schema.DefaultTimeout(15 * time.Minute),
			Delete: schema.DefaultTimeout(15 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"accelerator_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The type of hardware accelerators associated with this node.`,
			},
			"cidr_block": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The CIDR block that the TPU node will use when selecting an IP
address. This CIDR block must be a /29 block; the Compute Engine
networks API forbids a smaller block, and using a larger block would
be wasteful (a node can only consume one IP address).

Errors will occur if the CIDR block has already been used for a
currently existing TPU node, the CIDR block conflicts with any
subnetworks in the user's provided network, or the provided network
is peered with another network that is using that CIDR block.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The immutable name of the TPU.`,
			},
			"tensorflow_version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The version of Tensorflow running in the Node.`,
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The GCP location for the TPU.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The user-supplied description of the TPU. Maximum of 512 characters.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: `Resource labels to represent user provided metadata.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"network": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The name of a network to peer the TPU node to. It must be a
preexisting Compute Engine network inside of the project on which
this API has been activated. If none is provided, "default" will be
used.`,
			},
			"scheduling_config": {
				Type:             schema.TypeList,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareTpuNodeSchedulingConfig,
				Description:      `Sets the scheduling options for this TPU instance.`,
				MaxItems:         1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"preemptible": {
							Type:             schema.TypeBool,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: compareTpuNodeSchedulingConfig,
							Description:      `Defines whether the TPU instance is preemptible.`,
						},
					},
				},
			},
			"network_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `The network endpoints where TPU workers can be accessed and sent work.
It is recommended that Tensorflow clients of the node first reach out
to the first (index 0) entry.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The IP address of this network endpoint.`,
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `The port of this network endpoint.`,
						},
					},
				},
			},
			"service_account": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The service account used to run the tensor flow services within the
node. To share resources, including Google Cloud Storage data, with
the Tensorflow job running in the Node, this account must have
permissions to that data.`,
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

func resourceTPUNodeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandTPUNodeName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandTPUNodeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	acceleratorTypeProp, err := expandTPUNodeAcceleratorType(d.Get("accelerator_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("accelerator_type"); !isEmptyValue(reflect.ValueOf(acceleratorTypeProp)) && (ok || !reflect.DeepEqual(v, acceleratorTypeProp)) {
		obj["acceleratorType"] = acceleratorTypeProp
	}
	tensorflowVersionProp, err := expandTPUNodeTensorflowVersion(d.Get("tensorflow_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tensorflow_version"); !isEmptyValue(reflect.ValueOf(tensorflowVersionProp)) && (ok || !reflect.DeepEqual(v, tensorflowVersionProp)) {
		obj["tensorflowVersion"] = tensorflowVersionProp
	}
	networkProp, err := expandTPUNodeNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	cidrBlockProp, err := expandTPUNodeCidrBlock(d.Get("cidr_block"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cidr_block"); !isEmptyValue(reflect.ValueOf(cidrBlockProp)) && (ok || !reflect.DeepEqual(v, cidrBlockProp)) {
		obj["cidrBlock"] = cidrBlockProp
	}
	schedulingConfigProp, err := expandTPUNodeSchedulingConfig(d.Get("scheduling_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("scheduling_config"); !isEmptyValue(reflect.ValueOf(schedulingConfigProp)) && (ok || !reflect.DeepEqual(v, schedulingConfigProp)) {
		obj["schedulingConfig"] = schedulingConfigProp
	}
	labelsProp, err := expandTPUNodeLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := replaceVars(d, config, "{{TPUBasePath}}projects/{{project}}/locations/{{zone}}/nodes?nodeId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Node: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Node: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = tpuOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Node",
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Node: %s", err)
	}

	if err := d.Set("name", flattenTPUNodeName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Node %q: %#v", d.Id(), res)

	return resourceTPUNodeRead(d, meta)
}

func resourceTPUNodeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{TPUBasePath}}projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("TPUNode %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}

	if err := d.Set("name", flattenTPUNodeName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("description", flattenTPUNodeDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("accelerator_type", flattenTPUNodeAcceleratorType(res["acceleratorType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("tensorflow_version", flattenTPUNodeTensorflowVersion(res["tensorflowVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("network", flattenTPUNodeNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("cidr_block", flattenTPUNodeCidrBlock(res["cidrBlock"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("service_account", flattenTPUNodeServiceAccount(res["serviceAccount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("scheduling_config", flattenTPUNodeSchedulingConfig(res["schedulingConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("network_endpoints", flattenTPUNodeNetworkEndpoints(res["networkEndpoints"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("labels", flattenTPUNodeLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}

	return nil
}

func resourceTPUNodeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	d.Partial(true)

	if d.HasChange("tensorflow_version") {
		obj := make(map[string]interface{})

		tensorflowVersionProp, err := expandTPUNodeTensorflowVersion(d.Get("tensorflow_version"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("tensorflow_version"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, tensorflowVersionProp)) {
			obj["tensorflowVersion"] = tensorflowVersionProp
		}

		url, err := replaceVars(d, config, "{{TPUBasePath}}projects/{{project}}/locations/{{zone}}/nodes/{{name}}:reimage")
		if err != nil {
			return err
		}
		res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Node %q: %s", d.Id(), err)
		}

		err = tpuOperationWaitTime(
			config, res, project, "Updating Node",
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceTPUNodeRead(d, meta)
}

func resourceTPUNodeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{TPUBasePath}}projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Node %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Node")
	}

	err = tpuOperationWaitTime(
		config, res, project, "Deleting Node",
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Node %q: %#v", d.Id(), res)
	return nil
}

func resourceTPUNodeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<zone>[^/]+)/nodes/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenTPUNodeName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenTPUNodeDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTPUNodeAcceleratorType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTPUNodeTensorflowVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTPUNodeNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTPUNodeCidrBlock(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTPUNodeServiceAccount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTPUNodeSchedulingConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["preemptible"] =
		flattenTPUNodeSchedulingConfigPreemptible(original["preemptible"], d, config)
	return []interface{}{transformed}
}
func flattenTPUNodeSchedulingConfigPreemptible(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTPUNodeNetworkEndpoints(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"ip_address": flattenTPUNodeNetworkEndpointsIpAddress(original["ipAddress"], d, config),
			"port":       flattenTPUNodeNetworkEndpointsPort(original["port"], d, config),
		})
	}
	return transformed
}
func flattenTPUNodeNetworkEndpointsIpAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTPUNodeNetworkEndpointsPort(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenTPUNodeLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandTPUNodeName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeAcceleratorType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeTensorflowVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeCidrBlock(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeSchedulingConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPreemptible, err := expandTPUNodeSchedulingConfigPreemptible(original["preemptible"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPreemptible); val.IsValid() && !isEmptyValue(val) {
		transformed["preemptible"] = transformedPreemptible
	}

	return transformed, nil
}

func expandTPUNodeSchedulingConfigPreemptible(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
