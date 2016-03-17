package cluster

import (
	"testing"

	"github.com/docker/engine-api/types"
	"github.com/stretchr/testify/assert"
)

func TestNetworksFilter(t *testing.T) {
	engine := &Engine{ID: "id"}
	networks := Networks{
		{types.NetworkResource{
			ID:   "ababababab",
			Name: "something",
		}, engine},
		{types.NetworkResource{
			ID:   "aaaaaaaaaa1",
			Name: "network_name",
		}, engine},
		{types.NetworkResource{
			ID:   "bbbbbbbbbb",
			Name: "somethingelse",
		}, engine},
		{types.NetworkResource{
			ID:   "aaaaaaaaa2",
			Name: "foo",
		}, engine},
	}

	filtered := networks.Filter([]string{"network_name"}, []string{"abababab"}, nil)
	assert.Equal(t, len(filtered), 2)
	for _, network := range filtered {
		assert.True(t, network.ID == "aaaaaaaaaa1" || network.ID == "ababababab")
	}
}

func TestNetworkUniq(t *testing.T) {
	engine1 := &Engine{ID: "id1"}
	engine2 := &Engine{ID: "id2"}
	networks := Networks{
		{types.NetworkResource{
			ID:   "global",
			Name: "global",
			Containers: map[string]types.EndpointResource{
				"c1": {},
			},
		}, engine1},
		{types.NetworkResource{
			ID:   "global",
			Name: "global",
			Containers: map[string]types.EndpointResource{
				"c2": {},
			},
		}, engine2},
		{types.NetworkResource{
			ID:   "local1",
			Name: "local",
			Containers: map[string]types.EndpointResource{
				"c3": {},
			},
		}, engine1},
		{types.NetworkResource{
			ID:   "local2",
			Name: "local",
			Containers: map[string]types.EndpointResource{
				"c4": {},
			},
		}, engine2},
	}

	global := networks.Uniq().Get("global")
	assert.NotNil(t, global)
	assert.Equal(t, 2, len(global.Containers))

	local1 := networks.Uniq().Get("local1")
	assert.NotNil(t, local1)
	assert.Equal(t, 1, len(local1.Containers))

	local3 := networks.Uniq().Get("local3")
	assert.Nil(t, local3)
}
