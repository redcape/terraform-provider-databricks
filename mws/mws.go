package mws

// StsRole is the object that contains cross account role arn and external app id
type StsRole struct {
	RoleArn    string `json:"role_arn,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
}

// AwsCredentials is the object that points to the cross account role
type AwsCredentials struct {
	StsRole *StsRole `json:"sts_role,omitempty"`
}

// Credentials is the object that contains all the information for the credentials to create a workspace
type Credentials struct {
	CredentialsID   string          `json:"credentials_id,omitempty"`
	CredentialsName string          `json:"credentials_name,omitempty"`
	AwsCredentials  *AwsCredentials `json:"aws_credentials,omitempty"`
	AccountID       string          `json:"account_id,omitempty"`
	CreationTime    int64           `json:"creation_time,omitempty"`
}

// RootBucketInfo points to a bucket name
type RootBucketInfo struct {
	BucketName string `json:"bucket_name,omitempty"`
}

// StorageConfiguration is the object that contains all the information for the root storage bucket
type StorageConfiguration struct {
	StorageConfigurationID   string          `json:"storage_configuration_id,omitempty"`
	StorageConfigurationName string          `json:"storage_configuration_name,omitempty"`
	RootBucketInfo           *RootBucketInfo `json:"root_bucket_info,omitempty"`
	AccountID                string          `json:"account_id,omitempty"`
	CreationTime             int64           `json:"creation_time,omitempty"`
}

// NetworkHealth is the object that contains all the error message when attaching a network to workspace
type NetworkHealth struct {
	ErrorType    string `json:"error_type,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// NetworkVPCEndpoints is the object that contains VPC endpoints of a network
type NetworkVPCEndpoints struct {
	RestAPI           []string `json:"rest_api" tf:"slice_set"`
	DataplaneRelayAPI []string `json:"dataplane_relay" tf:"slice_set"`
}

// Network is the object that contains all the information for BYOVPC
type Network struct {
	AccountID        string               `json:"account_id"`
	NetworkID        string               `json:"network_id,omitempty" tf:"computed"`
	NetworkName      string               `json:"network_name"`
	VPCID            string               `json:"vpc_id"`
	SubnetIds        []string             `json:"subnet_ids" tf:"slice_set"`
	VPCEndpoints     *NetworkVPCEndpoints `json:"vpc_endpoints,omitempty" tf:"computed,force_new"`
	SecurityGroupIds []string             `json:"security_group_ids" tf:"slice_set"`
	VPCStatus        string               `json:"vpc_status,omitempty" tf:"computed"`
	ErrorMessages    []NetworkHealth      `json:"error_messages,omitempty" tf:"computed"`
	WorkspaceID      int64                `json:"workspace_id,omitempty" tf:"computed"`
	CreationTime     int64                `json:"creation_time,omitempty" tf:"computed"`
}

// VPCEndpoint is the object that contains all the information for registering an VPC endpoint
type VPCEndpoint struct {
	VPCEndpointID           string `json:"vpc_endpoint_id,omitempty" tf:"computed"`
	AwsVPCEndpointID        string `json:"aws_vpc_endpoint_id"`
	AccountID               string `json:"account_id,omitempty"`
	VPCEndpointName         string `json:"vpc_endpoint_name"`
	AwsVPCEndpointServiceID string `json:"aws_endpoint_service_id,omitempty" tf:"computed"`
	AWSAccountID            string `json:"aws_account_id,omitempty" tf:"computed"`
	UseCase                 string `json:"use_case,omitempty" tf:"computed"`
	Region                  string `json:"region"`
	State                   string `json:"state,omitempty" tf:"computed"`
}

// PrivateAccessSettings (PAS) is the object that contains all the information for creating an PrivateAccessSettings (PAS)
type PrivateAccessSettings struct {
	AccountID             string   `json:"account_id,omitempty"`
	PasID                 string   `json:"private_access_settings_id,omitempty" tf:"computed"`
	PasName               string   `json:"private_access_settings_name"`
	Region                string   `json:"region"`
	Status                string   `json:"status,omitempty" tf:"computed"`
	PublicAccessEnabled   bool     `json:"public_access_enabled,omitempty"`
	PrivateAccessLevel    string   `json:"private_access_level,omitempty" tf:"default:ANY"`
	AllowedVpcEndpointIDS []string `json:"allowed_vpc_endpoint_ids,omitempty"`
}
