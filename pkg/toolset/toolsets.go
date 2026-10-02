package toolset

import (
	"context"

	"github.com/containers/kubernetes-mcp-server/pkg/api"
	"github.com/containers/kubernetes-mcp-server/pkg/toolsets"
	"github.com/openshift/cluster-health-analyzer/pkg/toolset/incidents"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/rhobs/obs-mcp/pkg/logs"
	"github.com/rhobs/obs-mcp/pkg/metrics"
	"github.com/rhobs/obs-mcp/pkg/otelcol"
	"github.com/rhobs/obs-mcp/pkg/traces"
)

type incidentsToolset struct{ incidents.Toolset }

func (t *incidentsToolset) GetTools(p api.FilteringProvider) []api.ServerTool {
	if p == nil || !p.AnyTargetHasGVKs(context.TODO(), []schema.GroupVersionKind{{
		Group: "observability.openshift.io", Version: "v1alpha1", Kind: "UIPlugin",
	}}) {
		return nil
	}
	return t.Toolset.GetTools(p)
}

func init() {
	toolsets.Register(&metrics.Toolset{})
	toolsets.Register(&logs.Toolset{})
	toolsets.Register(&traces.Toolset{})
	toolsets.Register(&otelcol.Toolset{})
	toolsets.Register(&incidentsToolset{})
}
