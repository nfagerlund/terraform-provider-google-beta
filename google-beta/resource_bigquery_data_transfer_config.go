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
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBigqueryDataTransferConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigqueryDataTransferConfigCreate,
		Read:   resourceBigqueryDataTransferConfigRead,
		Update: resourceBigqueryDataTransferConfigUpdate,
		Delete: resourceBigqueryDataTransferConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigqueryDataTransferConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"data_source_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The data source id. Cannot be changed once the transfer config is created.`,
			},
			"destination_dataset_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The BigQuery target dataset id.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The user specified display name for the transfer config.`,
			},
			"params": {
				Type:        schema.TypeMap,
				Required:    true,
				Description: `These parameters are specific to each data source.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"data_refresh_window_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `The number of days to look back to automatically refresh the data.
For example, if dataRefreshWindowDays = 10, then every day BigQuery
reingests data for [today-10, today-1], rather than ingesting data for
just [today-1]. Only valid if the data source supports the feature.
Set the value to 0 to use the default value.`,
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `When set to true, no runs are scheduled for a given transfer.`,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The geographic location where the transfer config should reside.
Examples: US, EU, asia-northeast1. The default value is US.`,
				Default: "US",
			},
			"schedule": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Data transfer schedule. If the data source does not support a custom
schedule, this should be empty. If it is empty, the default value for
the data source will be used. The specified times are in UTC. Examples
of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
jun 13:15, and first sunday of quarter 00:00. See more explanation
about the format here:
https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
NOTE: the granularity should be at least 8 hours, or less frequent.`,
			},
			"service_account_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Optional service account name. If this field is set, transfer config will
be created with this service account credentials. It requires that
requesting user calling this API has permissions to act as this service account.`,
				Default: "",
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the transfer config. Transfer config names have the
form projects/{projectId}/locations/{location}/transferConfigs/{configId}.
Where configId is usually a uuid, but this is not required.
The name is ignored when creating a transfer config.`,
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

func resourceBigqueryDataTransferConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	displayNameProp, err := expandBigqueryDataTransferConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	destinationDatasetIdProp, err := expandBigqueryDataTransferConfigDestinationDatasetId(d.Get("destination_dataset_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("destination_dataset_id"); !isEmptyValue(reflect.ValueOf(destinationDatasetIdProp)) && (ok || !reflect.DeepEqual(v, destinationDatasetIdProp)) {
		obj["destinationDatasetId"] = destinationDatasetIdProp
	}
	dataSourceIdProp, err := expandBigqueryDataTransferConfigDataSourceId(d.Get("data_source_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_source_id"); !isEmptyValue(reflect.ValueOf(dataSourceIdProp)) && (ok || !reflect.DeepEqual(v, dataSourceIdProp)) {
		obj["dataSourceId"] = dataSourceIdProp
	}
	scheduleProp, err := expandBigqueryDataTransferConfigSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("schedule"); !isEmptyValue(reflect.ValueOf(scheduleProp)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	dataRefreshWindowDaysProp, err := expandBigqueryDataTransferConfigDataRefreshWindowDays(d.Get("data_refresh_window_days"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_refresh_window_days"); !isEmptyValue(reflect.ValueOf(dataRefreshWindowDaysProp)) && (ok || !reflect.DeepEqual(v, dataRefreshWindowDaysProp)) {
		obj["dataRefreshWindowDays"] = dataRefreshWindowDaysProp
	}
	disabledProp, err := expandBigqueryDataTransferConfigDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !isEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	paramsProp, err := expandBigqueryDataTransferConfigParams(d.Get("params"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("params"); !isEmptyValue(reflect.ValueOf(paramsProp)) && (ok || !reflect.DeepEqual(v, paramsProp)) {
		obj["params"] = paramsProp
	}

	url, err := replaceVars(d, config, "{{BigqueryDataTransferBasePath}}projects/{{project}}/locations/{{location}}/transferConfigs?serviceAccountName={{service_account_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Config: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate), iamMemberMissing)
	if err != nil {
		return fmt.Errorf("Error creating Config: %s", err)
	}
	if err := d.Set("name", flattenBigqueryDataTransferConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Config %q: %#v", d.Id(), res)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		respBody, ok := res["response"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}

		name, ok = respBody.(map[string]interface{})["name"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}
	}
	d.Set("name", name.(string))
	d.SetId(name.(string))

	return resourceBigqueryDataTransferConfigRead(d, meta)
}

func resourceBigqueryDataTransferConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{BigqueryDataTransferBasePath}}{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil, iamMemberMissing)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BigqueryDataTransferConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}

	if err := d.Set("display_name", flattenBigqueryDataTransferConfigDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("name", flattenBigqueryDataTransferConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("destination_dataset_id", flattenBigqueryDataTransferConfigDestinationDatasetId(res["destinationDatasetId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("data_source_id", flattenBigqueryDataTransferConfigDataSourceId(res["dataSourceId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("schedule", flattenBigqueryDataTransferConfigSchedule(res["schedule"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("data_refresh_window_days", flattenBigqueryDataTransferConfigDataRefreshWindowDays(res["dataRefreshWindowDays"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("disabled", flattenBigqueryDataTransferConfigDisabled(res["disabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("params", flattenBigqueryDataTransferConfigParams(res["params"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}

	return nil
}

func resourceBigqueryDataTransferConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	destinationDatasetIdProp, err := expandBigqueryDataTransferConfigDestinationDatasetId(d.Get("destination_dataset_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("destination_dataset_id"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, destinationDatasetIdProp)) {
		obj["destinationDatasetId"] = destinationDatasetIdProp
	}
	scheduleProp, err := expandBigqueryDataTransferConfigSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("schedule"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	dataRefreshWindowDaysProp, err := expandBigqueryDataTransferConfigDataRefreshWindowDays(d.Get("data_refresh_window_days"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_refresh_window_days"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, dataRefreshWindowDaysProp)) {
		obj["dataRefreshWindowDays"] = dataRefreshWindowDaysProp
	}
	disabledProp, err := expandBigqueryDataTransferConfigDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	paramsProp, err := expandBigqueryDataTransferConfigParams(d.Get("params"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("params"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, paramsProp)) {
		obj["params"] = paramsProp
	}

	url, err := replaceVars(d, config, "{{BigqueryDataTransferBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Config %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("destination_dataset_id") {
		updateMask = append(updateMask, "destinationDatasetId")
	}

	if d.HasChange("schedule") {
		updateMask = append(updateMask, "schedule")
	}

	if d.HasChange("data_refresh_window_days") {
		updateMask = append(updateMask, "dataRefreshWindowDays")
	}

	if d.HasChange("disabled") {
		updateMask = append(updateMask, "disabled")
	}

	if d.HasChange("params") {
		updateMask = append(updateMask, "params")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate), iamMemberMissing)

	if err != nil {
		return fmt.Errorf("Error updating Config %q: %s", d.Id(), err)
	}

	return resourceBigqueryDataTransferConfigRead(d, meta)
}

func resourceBigqueryDataTransferConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{BigqueryDataTransferBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Config %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete), iamMemberMissing)
	if err != nil {
		return handleNotFoundError(err, d, "Config")
	}

	log.Printf("[DEBUG] Finished deleting Config %q: %#v", d.Id(), res)
	return nil
}

func resourceBigqueryDataTransferConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenBigqueryDataTransferConfigDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigDestinationDatasetId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigDataSourceId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigSchedule(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigDataRefreshWindowDays(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenBigqueryDataTransferConfigDisabled(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigParams(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}

	kv := v.(map[string]interface{})

	res := make(map[string]string)
	for key, value := range kv {
		res[key] = fmt.Sprintf("%v", value)
	}
	return res
}

func expandBigqueryDataTransferConfigDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDestinationDatasetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDataSourceId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigSchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDataRefreshWindowDays(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDisabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigParams(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
