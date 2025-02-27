package clusters

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClusterData(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=abc",
				Response: ClusterInfo{
					ClusterID:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeID:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  ClusterStateRunning,
					AutoScale: &AutoScale{
						MaxWorkers: 4,
					},
				},
			},
		},
		Resource:    DataSourceCluster(),
		HCL:         `cluster_id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "abc",
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, 15, d.Get("cluster_info.0.autotermination_minutes"))
	assert.Equal(t, "Shared Autoscaling", d.Get("cluster_info.0.cluster_name"))
	assert.Equal(t, "i3.xlarge", d.Get("cluster_info.0.node_type_id"))
	assert.Equal(t, 4, d.Get("cluster_info.0.autoscale.0.max_workers"))
	assert.Equal(t, "RUNNING", d.Get("cluster_info.0.state"))

	for k, v := range d.State().Attributes {
		fmt.Printf("assert.Equal(t, %#v, d.Get(%#v))\n", v, k)
	}
}

func TestClusterData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceCluster(),
		Read:        true,
		NonWritable: true,
		HCL:         `cluster_id = "abc"`,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}
