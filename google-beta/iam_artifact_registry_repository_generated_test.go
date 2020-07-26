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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccArtifactRegistryRepositoryIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepositoryIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccArtifactRegistryRepositoryIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccArtifactRegistryRepositoryIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccArtifactRegistryRepositoryIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccArtifactRegistryRepositoryIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepositoryIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccArtifactRegistryRepositoryIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccArtifactRegistryRepositoryIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  provider = google-beta

  location = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description = "example docker repository%{random_suffix}"
  format = "DOCKER"
}

resource "google_artifact_registry_repository_iam_member" "foo" {
  provider = google-beta
  project = google_artifact_registry_repository.my-repo.project
  location = google_artifact_registry_repository.my-repo.location
  repository = google_artifact_registry_repository.my-repo.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccArtifactRegistryRepositoryIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  provider = google-beta

  location = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description = "example docker repository%{random_suffix}"
  format = "DOCKER"
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_artifact_registry_repository_iam_policy" "foo" {
  provider = google-beta
  project = google_artifact_registry_repository.my-repo.project
  location = google_artifact_registry_repository.my-repo.location
  repository = google_artifact_registry_repository.my-repo.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccArtifactRegistryRepositoryIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  provider = google-beta

  location = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description = "example docker repository%{random_suffix}"
  format = "DOCKER"
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_artifact_registry_repository_iam_policy" "foo" {
  provider = google-beta
  project = google_artifact_registry_repository.my-repo.project
  location = google_artifact_registry_repository.my-repo.location
  repository = google_artifact_registry_repository.my-repo.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccArtifactRegistryRepositoryIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  provider = google-beta

  location = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description = "example docker repository%{random_suffix}"
  format = "DOCKER"
}

resource "google_artifact_registry_repository_iam_binding" "foo" {
 
  provider = google-beta
  project = google_artifact_registry_repository.my-repo.project
  location = google_artifact_registry_repository.my-repo.location
  repository = google_artifact_registry_repository.my-repo.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccArtifactRegistryRepositoryIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  provider = google-beta

  location = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description = "example docker repository%{random_suffix}"
  format = "DOCKER"
}

resource "google_artifact_registry_repository_iam_binding" "foo" {
  provider = google-beta
  project = google_artifact_registry_repository.my-repo.project
  location = google_artifact_registry_repository.my-repo.location
  repository = google_artifact_registry_repository.my-repo.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
