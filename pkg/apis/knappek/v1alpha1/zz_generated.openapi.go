// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasCluster":       schema_pkg_apis_knappek_v1alpha1_MongoDBAtlasCluster(ref),
		"github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasClusterSpec":   schema_pkg_apis_knappek_v1alpha1_MongoDBAtlasClusterSpec(ref),
		"github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasClusterStatus": schema_pkg_apis_knappek_v1alpha1_MongoDBAtlasClusterStatus(ref),
		"github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasProject":       schema_pkg_apis_knappek_v1alpha1_MongoDBAtlasProject(ref),
		"github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasProjectSpec":   schema_pkg_apis_knappek_v1alpha1_MongoDBAtlasProjectSpec(ref),
		"github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasProjectStatus": schema_pkg_apis_knappek_v1alpha1_MongoDBAtlasProjectStatus(ref),
	}
}

func schema_pkg_apis_knappek_v1alpha1_MongoDBAtlasCluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MongoDBAtlasCluster is the Schema for the mongodbatlasclusters API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasClusterSpec", "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasClusterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_knappek_v1alpha1_MongoDBAtlasClusterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MongoDBAtlasClusterSpec defines the desired state of MongoDBAtlasCluster",
				Properties: map[string]spec.Schema{
					"publicKey": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"privateKey": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.PrivateKey"),
						},
					},
					"projectName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mongoDBVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mongoDBMajorVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"diskSizeGB": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"number"},
							Format: "double",
						},
					},
					"backupEnabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"providerBackupEnabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"replicationFactor": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"replicationSpec": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/akshaykarle/go-mongodbatlas/mongodbatlas.ReplicationSpec"),
									},
								},
							},
						},
					},
					"numShards": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"autoScaling": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/akshaykarle/go-mongodbatlas/mongodbatlas.AutoScaling"),
						},
					},
					"providerSettings": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/akshaykarle/go-mongodbatlas/mongodbatlas.ProviderSettings"),
						},
					},
				},
				Required: []string{"publicKey", "privateKey", "projectName", "backupEnabled", "providerBackupEnabled"},
			},
		},
		Dependencies: []string{
			"github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.PrivateKey", "github.com/akshaykarle/go-mongodbatlas/mongodbatlas.AutoScaling", "github.com/akshaykarle/go-mongodbatlas/mongodbatlas.ProviderSettings", "github.com/akshaykarle/go-mongodbatlas/mongodbatlas.ReplicationSpec"},
	}
}

func schema_pkg_apis_knappek_v1alpha1_MongoDBAtlasClusterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MongoDBAtlasClusterStatus defines the observed state of MongoDBAtlasCluster",
				Properties: map[string]spec.Schema{
					"id": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"groupID": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mongoDBVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mongoDBMajorVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mongoURI": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mongoURIUpdated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mongoURIWithOptions": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"srvAddress": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"diskSizeGB": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"number"},
							Format: "double",
						},
					},
					"backupEnabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"providerBackupEnabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"stateName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"replicationFactor": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"replicationSpec": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/akshaykarle/go-mongodbatlas/mongodbatlas.ReplicationSpec"),
									},
								},
							},
						},
					},
					"numShards": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"paused": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"autoScaling": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/akshaykarle/go-mongodbatlas/mongodbatlas.AutoScaling"),
						},
					},
					"providerSettings": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/akshaykarle/go-mongodbatlas/mongodbatlas.ProviderSettings"),
						},
					},
				},
				Required: []string{"backupEnabled", "providerBackupEnabled", "paused"},
			},
		},
		Dependencies: []string{
			"github.com/akshaykarle/go-mongodbatlas/mongodbatlas.AutoScaling", "github.com/akshaykarle/go-mongodbatlas/mongodbatlas.ProviderSettings", "github.com/akshaykarle/go-mongodbatlas/mongodbatlas.ReplicationSpec"},
	}
}

func schema_pkg_apis_knappek_v1alpha1_MongoDBAtlasProject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MongoDBAtlasProject is the Schema for the mongodbatlasprojects API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasProjectSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasProjectStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasProjectSpec", "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.MongoDBAtlasProjectStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_knappek_v1alpha1_MongoDBAtlasProjectSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MongoDBAtlasProjectSpec defines the desired state of MongoDBAtlasProject",
				Properties: map[string]spec.Schema{
					"publicKey": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"privateKey": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.PrivateKey"),
						},
					},
					"orgID": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"publicKey", "privateKey"},
			},
		},
		Dependencies: []string{
			"github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1.PrivateKey"},
	}
}

func schema_pkg_apis_knappek_v1alpha1_MongoDBAtlasProjectStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MongoDBAtlasProjectStatus defines the observed state of MongoDBAtlasProject",
				Properties: map[string]spec.Schema{
					"id": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"orgID": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"created": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"clusterCount": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
				Required: []string{"id", "name", "orgID", "created", "clusterCount"},
			},
		},
		Dependencies: []string{},
	}
}
