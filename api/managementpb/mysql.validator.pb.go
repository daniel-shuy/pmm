// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/mysql.proto

package managementpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/percona/pmm/api/inventorypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AddMySQLRequest) Validate() error {
	if this.AddNode != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AddNode); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AddNode", err)
		}
	}
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.PmmAgentId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PmmAgentId", fmt.Errorf(`value '%v' must not be an empty string`, this.PmmAgentId))
	}
	if this.Username == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Username", fmt.Errorf(`value '%v' must not be an empty string`, this.Username))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddMySQLResponse) Validate() error {
	if this.Service != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Service); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Service", err)
		}
	}
	if this.MysqldExporter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MysqldExporter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MysqldExporter", err)
		}
	}
	if this.QanMysqlPerfschema != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.QanMysqlPerfschema); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("QanMysqlPerfschema", err)
		}
	}
	if this.QanMysqlSlowlog != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.QanMysqlSlowlog); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("QanMysqlSlowlog", err)
		}
	}
	return nil
}
