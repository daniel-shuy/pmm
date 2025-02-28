// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package management

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm/api/inventorypb"
	"github.com/percona/pmm/api/managementpb"
	"github.com/percona/pmm/managed/models"
	"github.com/percona/pmm/managed/services"
)

// ExternalService External Management Service.
//
//nolint:unused
type ExternalService struct {
	db    *reform.DB
	vmdb  prometheusService
	state agentsStateUpdater
	cc    connectionChecker

	managementpb.UnimplementedExternalServer
}

// NewExternalService creates new External Management Service.
func NewExternalService(db *reform.DB, vmdb prometheusService, state agentsStateUpdater, cc connectionChecker) *ExternalService {
	return &ExternalService{
		db:    db,
		vmdb:  vmdb,
		state: state,
		cc:    cc,
	}
}

func (e *ExternalService) AddExternal(ctx context.Context, req *managementpb.AddExternalRequest) (*managementpb.AddExternalResponse, error) {
	res := &managementpb.AddExternalResponse{}
	var pmmAgentID *string
	if e := e.db.InTransaction(func(tx *reform.TX) error {
		if (req.NodeId == "") != (req.RunsOnNodeId == "") {
			return status.Error(codes.InvalidArgument, "runs_on_node_id and node_id should be specified together.")
		}
		if req.Address == "" && req.AddNode != nil {
			return status.Error(codes.InvalidArgument, "address can't be empty for add node request.")
		}
		nodeID, err := nodeID(tx, req.NodeId, req.NodeName, req.AddNode, req.Address)
		if err != nil {
			return err
		}

		runsOnNodeId := req.RunsOnNodeId
		if req.AddNode != nil && runsOnNodeId == "" {
			runsOnNodeId = nodeID
		}

		service, err := models.AddNewService(tx.Querier, models.ExternalServiceType, &models.AddDBMSServiceParams{
			ServiceName:    req.ServiceName,
			NodeID:         nodeID,
			Environment:    req.Environment,
			Cluster:        req.Cluster,
			ReplicationSet: req.ReplicationSet,
			CustomLabels:   req.CustomLabels,
			ExternalGroup:  req.Group,
		})
		if err != nil {
			return err
		}

		invService, err := services.ToAPIService(service)
		if err != nil {
			return err
		}
		res.Service = invService.(*inventorypb.ExternalService)

		if req.MetricsMode == managementpb.MetricsMode_AUTO {
			agentIDs, err := models.FindPMMAgentsRunningOnNode(tx.Querier, req.RunsOnNodeId)
			switch {
			case err != nil || len(agentIDs) != 1:
				req.MetricsMode = managementpb.MetricsMode_PULL
			default:
				req.MetricsMode, err = supportedMetricsMode(tx.Querier, req.MetricsMode, agentIDs[0].AgentID)
				if err != nil {
					return err
				}
			}
		}

		params := &models.CreateExternalExporterParams{
			RunsOnNodeID: runsOnNodeId,
			ServiceID:    service.ServiceID,
			Username:     req.Username,
			Password:     req.Password,
			Scheme:       req.Scheme,
			MetricsPath:  req.MetricsPath,
			ListenPort:   req.ListenPort,
			CustomLabels: req.CustomLabels,
			PushMetrics:  isPushMode(req.MetricsMode),
		}
		row, err := models.CreateExternalExporter(tx.Querier, params)
		if err != nil {
			return err
		}

		if !req.SkipConnectionCheck {
			if err = e.cc.CheckConnectionToService(ctx, tx.Querier, service, row); err != nil {
				return err
			}
		}

		agent, err := services.ToAPIAgent(tx.Querier, row)
		if err != nil {
			return err
		}
		res.ExternalExporter = agent.(*inventorypb.ExternalExporter)
		pmmAgentID = row.PMMAgentID

		return nil
	}); e != nil {
		return nil, e
	}
	// we have to trigger after transaction
	if pmmAgentID != nil {
		// It's required to regenerate victoriametrics config file.
		e.state.RequestStateUpdate(ctx, *pmmAgentID)
	} else {
		e.vmdb.RequestConfigurationUpdate()
	}
	return res, nil
}
