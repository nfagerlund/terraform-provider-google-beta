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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOSLoginSSHPublicKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceOSLoginSSHPublicKeyCreate,
		Read:   resourceOSLoginSSHPublicKeyRead,
		Update: resourceOSLoginSSHPublicKeyUpdate,
		Delete: resourceOSLoginSSHPublicKeyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceOSLoginSSHPublicKeyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Public key text in SSH format, defined by RFC4253 section 6.6.`,
			},
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The user email.`,
			},
			"expiration_time_usec": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An expiration time in microseconds since epoch.`,
			},
			"fingerprint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The SHA-256 fingerprint of the SSH public key.`,
			},
		},
	}
}

func resourceOSLoginSSHPublicKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	keyProp, err := expandOSLoginSSHPublicKeyKey(d.Get("key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("key"); !isEmptyValue(reflect.ValueOf(keyProp)) && (ok || !reflect.DeepEqual(v, keyProp)) {
		obj["key"] = keyProp
	}
	expirationTimeUsecProp, err := expandOSLoginSSHPublicKeyExpirationTimeUsec(d.Get("expiration_time_usec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expiration_time_usec"); !isEmptyValue(reflect.ValueOf(expirationTimeUsecProp)) && (ok || !reflect.DeepEqual(v, expirationTimeUsecProp)) {
		obj["expirationTimeUsec"] = expirationTimeUsecProp
	}

	url, err := replaceVars(d, config, "{{OSLoginBasePath}}users/{{user}}:importSshPublicKey")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SSHPublicKey: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", "", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating SSHPublicKey: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "users/{{user}}/sshPublicKeys/{{fingerprint}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating SSHPublicKey %q: %#v", d.Id(), res)

	loginProfile, ok := res["loginProfile"]
	if !ok {
		return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
	}

	// `fingerprint` is autogenerated from the api so needs to be set post-create
	sshPublicKeys := loginProfile.(map[string]interface{})["sshPublicKeys"]
	for _, sshPublicKey := range sshPublicKeys.(map[string]interface{}) {
		if sshPublicKey.(map[string]interface{})["key"].(string) == d.Get("key") {
			d.Set("fingerprint", sshPublicKey.(map[string]interface{})["fingerprint"].(string))
			break
		}
	}

	// Store the ID now
	id, err = replaceVars(d, config, "users/{{user}}/sshPublicKeys/{{fingerprint}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceOSLoginSSHPublicKeyRead(d, meta)
}

func resourceOSLoginSSHPublicKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{OSLoginBasePath}}users/{{user}}/sshPublicKeys/{{fingerprint}}/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", "", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("OSLoginSSHPublicKey %q", d.Id()))
	}

	if err := d.Set("key", flattenOSLoginSSHPublicKeyKey(res["key"], d, config)); err != nil {
		return fmt.Errorf("Error reading SSHPublicKey: %s", err)
	}
	if err := d.Set("expiration_time_usec", flattenOSLoginSSHPublicKeyExpirationTimeUsec(res["expirationTimeUsec"], d, config)); err != nil {
		return fmt.Errorf("Error reading SSHPublicKey: %s", err)
	}
	if err := d.Set("fingerprint", flattenOSLoginSSHPublicKeyFingerprint(res["fingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading SSHPublicKey: %s", err)
	}

	return nil
}

func resourceOSLoginSSHPublicKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	expirationTimeUsecProp, err := expandOSLoginSSHPublicKeyExpirationTimeUsec(d.Get("expiration_time_usec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expiration_time_usec"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, expirationTimeUsecProp)) {
		obj["expirationTimeUsec"] = expirationTimeUsecProp
	}

	url, err := replaceVars(d, config, "{{OSLoginBasePath}}users/{{user}}/sshPublicKeys/{{fingerprint}}/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating SSHPublicKey %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("expiration_time_usec") {
		updateMask = append(updateMask, "expirationTimeUsec")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", "", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating SSHPublicKey %q: %s", d.Id(), err)
	}

	return resourceOSLoginSSHPublicKeyRead(d, meta)
}

func resourceOSLoginSSHPublicKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{OSLoginBasePath}}users/{{user}}/sshPublicKeys/{{fingerprint}}/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting SSHPublicKey %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", "", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "SSHPublicKey")
	}

	log.Printf("[DEBUG] Finished deleting SSHPublicKey %q: %#v", d.Id(), res)
	return nil
}

func resourceOSLoginSSHPublicKeyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"users/(?P<user>[^/]+)/sshPublicKeys/(?P<fingerprint>[^/]+)",
		"(?P<user>[^/]+)/(?P<fingerprint>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "users/{{user}}/sshPublicKeys/{{fingerprint}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenOSLoginSSHPublicKeyKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenOSLoginSSHPublicKeyExpirationTimeUsec(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenOSLoginSSHPublicKeyFingerprint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandOSLoginSSHPublicKeyKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandOSLoginSSHPublicKeyExpirationTimeUsec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
