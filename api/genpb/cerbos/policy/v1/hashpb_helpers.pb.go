// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.2.0

package policyv1

import (
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	protowire "google.golang.org/protobuf/encoding/protowire"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	hash "hash"
	math "math"
	sort "sort"
)

func cerbos_engine_v1_AuxData_hashpb_sum(m *v1.AuxData, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.AuxData.jwt"]; !ok {
		if len(m.Jwt) > 0 {
			keys := make([]string, len(m.Jwt))
			i := 0
			for k := range m.Jwt {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Jwt[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Jwt[k], hasher, ignore)
				}

			}
		}
	}
}

func cerbos_engine_v1_CheckInput_hashpb_sum(m *v1.CheckInput, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.CheckInput.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.engine.v1.CheckInput.resource"]; !ok {
		if m.Resource != nil {
			cerbos_engine_v1_Resource_hashpb_sum(m.Resource, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.CheckInput.principal"]; !ok {
		if m.Principal != nil {
			cerbos_engine_v1_Principal_hashpb_sum(m.Principal, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.CheckInput.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.CheckInput.aux_data"]; !ok {
		if m.AuxData != nil {
			cerbos_engine_v1_AuxData_hashpb_sum(m.AuxData, hasher, ignore)
		}

	}
}

func cerbos_engine_v1_Principal_hashpb_sum(m *v1.Principal, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.Principal.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Id))

	}
	if _, ok := ignore["cerbos.engine.v1.Principal.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.engine.v1.Principal.roles"]; !ok {
		if len(m.Roles) > 0 {
			for _, v := range m.Roles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.Principal.attr"]; !ok {
		if len(m.Attr) > 0 {
			keys := make([]string, len(m.Attr))
			i := 0
			for k := range m.Attr {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attr[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attr[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.Principal.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
}

func cerbos_engine_v1_Resource_hashpb_sum(m *v1.Resource, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.Resource.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Kind))

	}
	if _, ok := ignore["cerbos.engine.v1.Resource.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.engine.v1.Resource.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Id))

	}
	if _, ok := ignore["cerbos.engine.v1.Resource.attr"]; !ok {
		if len(m.Attr) > 0 {
			keys := make([]string, len(m.Attr))
			i := 0
			for k := range m.Attr {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attr[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attr[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.Resource.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
}

func cerbos_engine_v1_Trace_Component_Variable_hashpb_sum(m *v1.Trace_Component_Variable, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.Trace.Component.Variable.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.engine.v1.Trace.Component.Variable.expr"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Expr))

	}
}

func cerbos_engine_v1_Trace_Component_hashpb_sum(m *v1.Trace_Component, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.Trace.Component.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Kind)))

	}
	if m.Details != nil {
		if _, ok := ignore["cerbos.engine.v1.Trace.Component.details"]; !ok {
			switch t := m.Details.(type) {
			case *v1.Trace_Component_Action:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Action))

			case *v1.Trace_Component_DerivedRole:
				_, _ = hasher.Write(protowire.AppendString(nil, t.DerivedRole))

			case *v1.Trace_Component_Expr:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Expr))

			case *v1.Trace_Component_Index:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.Index)))

			case *v1.Trace_Component_Policy:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Policy))

			case *v1.Trace_Component_Resource:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Resource))

			case *v1.Trace_Component_Rule:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Rule))

			case *v1.Trace_Component_Scope:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Scope))

			case *v1.Trace_Component_Variable_:
				if t.Variable != nil {
					cerbos_engine_v1_Trace_Component_Variable_hashpb_sum(t.Variable, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_engine_v1_Trace_Event_hashpb_sum(m *v1.Trace_Event, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.Trace.Event.status"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Status)))

	}
	if _, ok := ignore["cerbos.engine.v1.Trace.Event.effect"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Effect)))

	}
	if _, ok := ignore["cerbos.engine.v1.Trace.Event.error"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Error))

	}
	if _, ok := ignore["cerbos.engine.v1.Trace.Event.message"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Message))

	}
	if _, ok := ignore["cerbos.engine.v1.Trace.Event.result"]; !ok {
		if m.Result != nil {
			google_protobuf_Value_hashpb_sum(m.Result, hasher, ignore)
		}

	}
}

func cerbos_engine_v1_Trace_hashpb_sum(m *v1.Trace, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.Trace.components"]; !ok {
		if len(m.Components) > 0 {
			for _, v := range m.Components {
				if v != nil {
					cerbos_engine_v1_Trace_Component_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.Trace.event"]; !ok {
		if m.Event != nil {
			cerbos_engine_v1_Trace_Event_hashpb_sum(m.Event, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_Condition_hashpb_sum(m *Condition, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Condition != nil {
		if _, ok := ignore["cerbos.policy.v1.Condition.condition"]; !ok {
			switch t := m.Condition.(type) {
			case *Condition_Match:
				if t.Match != nil {
					cerbos_policy_v1_Match_hashpb_sum(t.Match, hasher, ignore)
				}

			case *Condition_Script:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Script))

			}
		}
	}
}

func cerbos_policy_v1_DerivedRoles_hashpb_sum(m *DerivedRoles, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.DerivedRoles.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.policy.v1.DerivedRoles.definitions"]; !ok {
		if len(m.Definitions) > 0 {
			for _, v := range m.Definitions {
				if v != nil {
					cerbos_policy_v1_RoleDef_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_Match_ExprList_hashpb_sum(m *Match_ExprList, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Match.ExprList.of"]; !ok {
		if len(m.Of) > 0 {
			for _, v := range m.Of {
				if v != nil {
					cerbos_policy_v1_Match_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_Match_hashpb_sum(m *Match, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Op != nil {
		if _, ok := ignore["cerbos.policy.v1.Match.op"]; !ok {
			switch t := m.Op.(type) {
			case *Match_All:
				if t.All != nil {
					cerbos_policy_v1_Match_ExprList_hashpb_sum(t.All, hasher, ignore)
				}

			case *Match_Any:
				if t.Any != nil {
					cerbos_policy_v1_Match_ExprList_hashpb_sum(t.Any, hasher, ignore)
				}

			case *Match_None:
				if t.None != nil {
					cerbos_policy_v1_Match_ExprList_hashpb_sum(t.None, hasher, ignore)
				}

			case *Match_Expr:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Expr))

			}
		}
	}
}

func cerbos_policy_v1_Metadata_hashpb_sum(m *Metadata, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Metadata.source_file"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.SourceFile))

	}
	if _, ok := ignore["cerbos.policy.v1.Metadata.annotations"]; !ok {
		if len(m.Annotations) > 0 {
			keys := make([]string, len(m.Annotations))
			i := 0
			for k := range m.Annotations {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendString(nil, m.Annotations[k]))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.Metadata.hash"]; !ok {
		if m.Hash != nil {
			google_protobuf_UInt64Value_hashpb_sum(m.Hash, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.Metadata.store_identifer"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.StoreIdentifer))

	}
	if _, ok := ignore["cerbos.policy.v1.Metadata.store_identifier"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.StoreIdentifier))

	}
}

func cerbos_policy_v1_Policy_hashpb_sum(m *Policy, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Policy.api_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.ApiVersion))

	}
	if _, ok := ignore["cerbos.policy.v1.Policy.disabled"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.Disabled)))

	}
	if _, ok := ignore["cerbos.policy.v1.Policy.description"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Description))

	}
	if _, ok := ignore["cerbos.policy.v1.Policy.metadata"]; !ok {
		if m.Metadata != nil {
			cerbos_policy_v1_Metadata_hashpb_sum(m.Metadata, hasher, ignore)
		}

	}
	if m.PolicyType != nil {
		if _, ok := ignore["cerbos.policy.v1.Policy.policy_type"]; !ok {
			switch t := m.PolicyType.(type) {
			case *Policy_ResourcePolicy:
				if t.ResourcePolicy != nil {
					cerbos_policy_v1_ResourcePolicy_hashpb_sum(t.ResourcePolicy, hasher, ignore)
				}

			case *Policy_PrincipalPolicy:
				if t.PrincipalPolicy != nil {
					cerbos_policy_v1_PrincipalPolicy_hashpb_sum(t.PrincipalPolicy, hasher, ignore)
				}

			case *Policy_DerivedRoles:
				if t.DerivedRoles != nil {
					cerbos_policy_v1_DerivedRoles_hashpb_sum(t.DerivedRoles, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.Policy.variables"]; !ok {
		if len(m.Variables) > 0 {
			keys := make([]string, len(m.Variables))
			i := 0
			for k := range m.Variables {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendString(nil, m.Variables[k]))

			}
		}
	}
}

func cerbos_policy_v1_PrincipalPolicy_hashpb_sum(m *PrincipalPolicy, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.principal"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Principal))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Version))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.rules"]; !ok {
		if len(m.Rules) > 0 {
			for _, v := range m.Rules {
				if v != nil {
					cerbos_policy_v1_PrincipalRule_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
}

func cerbos_policy_v1_PrincipalRule_Action_hashpb_sum(m *PrincipalRule_Action, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.action"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Action))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.condition"]; !ok {
		if m.Condition != nil {
			cerbos_policy_v1_Condition_hashpb_sum(m.Condition, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.effect"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Effect)))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
}

func cerbos_policy_v1_PrincipalRule_hashpb_sum(m *PrincipalRule, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.resource"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Resource))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				if v != nil {
					cerbos_policy_v1_PrincipalRule_Action_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_ResourcePolicy_hashpb_sum(m *ResourcePolicy, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.resource"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Resource))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Version))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.import_derived_roles"]; !ok {
		if len(m.ImportDerivedRoles) > 0 {
			for _, v := range m.ImportDerivedRoles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.rules"]; !ok {
		if len(m.Rules) > 0 {
			for _, v := range m.Rules {
				if v != nil {
					cerbos_policy_v1_ResourceRule_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.schemas"]; !ok {
		if m.Schemas != nil {
			cerbos_policy_v1_Schemas_hashpb_sum(m.Schemas, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_ResourceRule_hashpb_sum(m *ResourceRule, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.derived_roles"]; !ok {
		if len(m.DerivedRoles) > 0 {
			for _, v := range m.DerivedRoles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.roles"]; !ok {
		if len(m.Roles) > 0 {
			for _, v := range m.Roles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.condition"]; !ok {
		if m.Condition != nil {
			cerbos_policy_v1_Condition_hashpb_sum(m.Condition, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.effect"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Effect)))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
}

func cerbos_policy_v1_RoleDef_hashpb_sum(m *RoleDef, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.RoleDef.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.policy.v1.RoleDef.parent_roles"]; !ok {
		if len(m.ParentRoles) > 0 {
			for _, v := range m.ParentRoles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.RoleDef.condition"]; !ok {
		if m.Condition != nil {
			cerbos_policy_v1_Condition_hashpb_sum(m.Condition, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_Schemas_IgnoreWhen_hashpb_sum(m *Schemas_IgnoreWhen, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Schemas.IgnoreWhen.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_policy_v1_Schemas_Schema_hashpb_sum(m *Schemas_Schema, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Schemas.Schema.ref"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Ref))

	}
	if _, ok := ignore["cerbos.policy.v1.Schemas.Schema.ignore_when"]; !ok {
		if m.IgnoreWhen != nil {
			cerbos_policy_v1_Schemas_IgnoreWhen_hashpb_sum(m.IgnoreWhen, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_Schemas_hashpb_sum(m *Schemas, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Schemas.principal_schema"]; !ok {
		if m.PrincipalSchema != nil {
			cerbos_policy_v1_Schemas_Schema_hashpb_sum(m.PrincipalSchema, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.Schemas.resource_schema"]; !ok {
		if m.ResourceSchema != nil {
			cerbos_policy_v1_Schemas_Schema_hashpb_sum(m.ResourceSchema, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_TestFixture_hashpb_sum(m *TestFixture, hasher hash.Hash, ignore map[string]struct{}) {
}

func cerbos_policy_v1_TestOptions_hashpb_sum(m *TestOptions, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestOptions.now"]; !ok {
		if m.Now != nil {
			google_protobuf_Timestamp_hashpb_sum(m.Now, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_TestResults_Action_hashpb_sum(m *TestResults_Action, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestResults.Action.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Action.details"]; !ok {
		if m.Details != nil {
			cerbos_policy_v1_TestResults_Details_hashpb_sum(m.Details, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_TestResults_Details_hashpb_sum(m *TestResults_Details, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestResults.Details.result"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Result)))

	}
	if m.Outcome != nil {
		if _, ok := ignore["cerbos.policy.v1.TestResults.Details.outcome"]; !ok {
			switch t := m.Outcome.(type) {
			case *TestResults_Details_Failure:
				if t.Failure != nil {
					cerbos_policy_v1_TestResults_Failure_hashpb_sum(t.Failure, hasher, ignore)
				}

			case *TestResults_Details_Error:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Error))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Details.engine_trace"]; !ok {
		if len(m.EngineTrace) > 0 {
			for _, v := range m.EngineTrace {
				if v != nil {
					cerbos_engine_v1_Trace_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_TestResults_Failure_hashpb_sum(m *TestResults_Failure, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestResults.Failure.expected"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Expected)))

	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Failure.actual"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Actual)))

	}
}

func cerbos_policy_v1_TestResults_Principal_hashpb_sum(m *TestResults_Principal, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestResults.Principal.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Principal.resources"]; !ok {
		if len(m.Resources) > 0 {
			for _, v := range m.Resources {
				if v != nil {
					cerbos_policy_v1_TestResults_Resource_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_TestResults_Resource_hashpb_sum(m *TestResults_Resource, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestResults.Resource.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Resource.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				if v != nil {
					cerbos_policy_v1_TestResults_Action_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_TestResults_Suite_hashpb_sum(m *TestResults_Suite, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestResults.Suite.file"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.File))

	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Suite.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Suite.principals"]; !ok {
		if len(m.Principals) > 0 {
			for _, v := range m.Principals {
				if v != nil {
					cerbos_policy_v1_TestResults_Principal_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Suite.summary"]; !ok {
		if m.Summary != nil {
			cerbos_policy_v1_TestResults_Summary_hashpb_sum(m.Summary, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Suite.error"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Error))

	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Suite.test_cases"]; !ok {
		if len(m.TestCases) > 0 {
			for _, v := range m.TestCases {
				if v != nil {
					cerbos_policy_v1_TestResults_TestCase_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_TestResults_Summary_hashpb_sum(m *TestResults_Summary, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestResults.Summary.overall_result"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.OverallResult)))

	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Summary.tests_count"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.TestsCount)))

	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Summary.result_counts"]; !ok {
		if len(m.ResultCounts) > 0 {
			for _, v := range m.ResultCounts {
				if v != nil {
					cerbos_policy_v1_TestResults_Tally_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_TestResults_Tally_hashpb_sum(m *TestResults_Tally, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestResults.Tally.result"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Result)))

	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.Tally.count"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Count)))

	}
}

func cerbos_policy_v1_TestResults_TestCase_hashpb_sum(m *TestResults_TestCase, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestResults.TestCase.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.TestCase.principals"]; !ok {
		if len(m.Principals) > 0 {
			for _, v := range m.Principals {
				if v != nil {
					cerbos_policy_v1_TestResults_Principal_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_TestResults_hashpb_sum(m *TestResults, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestResults.suites"]; !ok {
		if len(m.Suites) > 0 {
			for _, v := range m.Suites {
				if v != nil {
					cerbos_policy_v1_TestResults_Suite_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.TestResults.summary"]; !ok {
		if m.Summary != nil {
			cerbos_policy_v1_TestResults_Summary_hashpb_sum(m.Summary, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_TestSuite_hashpb_sum(m *TestSuite, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestSuite.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.policy.v1.TestSuite.description"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Description))

	}
	if _, ok := ignore["cerbos.policy.v1.TestSuite.skip"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.Skip)))

	}
	if _, ok := ignore["cerbos.policy.v1.TestSuite.skip_reason"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.SkipReason))

	}
	if _, ok := ignore["cerbos.policy.v1.TestSuite.tests"]; !ok {
		if len(m.Tests) > 0 {
			for _, v := range m.Tests {
				if v != nil {
					cerbos_policy_v1_TestTable_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.TestSuite.principals"]; !ok {
		if len(m.Principals) > 0 {
			keys := make([]string, len(m.Principals))
			i := 0
			for k := range m.Principals {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Principals[k] != nil {
					cerbos_engine_v1_Principal_hashpb_sum(m.Principals[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.TestSuite.resources"]; !ok {
		if len(m.Resources) > 0 {
			keys := make([]string, len(m.Resources))
			i := 0
			for k := range m.Resources {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Resources[k] != nil {
					cerbos_engine_v1_Resource_hashpb_sum(m.Resources[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.TestSuite.aux_data"]; !ok {
		if len(m.AuxData) > 0 {
			keys := make([]string, len(m.AuxData))
			i := 0
			for k := range m.AuxData {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.AuxData[k] != nil {
					cerbos_engine_v1_AuxData_hashpb_sum(m.AuxData[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.TestSuite.options"]; !ok {
		if m.Options != nil {
			cerbos_policy_v1_TestOptions_hashpb_sum(m.Options, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_TestTable_Expectation_hashpb_sum(m *TestTable_Expectation, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestTable.Expectation.principal"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Principal))

	}
	if _, ok := ignore["cerbos.policy.v1.TestTable.Expectation.resource"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Resource))

	}
	if _, ok := ignore["cerbos.policy.v1.TestTable.Expectation.actions"]; !ok {
		if len(m.Actions) > 0 {
			keys := make([]string, len(m.Actions))
			i := 0
			for k := range m.Actions {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Actions[k])))

			}
		}
	}
}

func cerbos_policy_v1_TestTable_Input_hashpb_sum(m *TestTable_Input, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestTable.Input.principals"]; !ok {
		if len(m.Principals) > 0 {
			for _, v := range m.Principals {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.TestTable.Input.resources"]; !ok {
		if len(m.Resources) > 0 {
			for _, v := range m.Resources {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.TestTable.Input.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.TestTable.Input.aux_data"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.AuxData))

	}
}

func cerbos_policy_v1_TestTable_hashpb_sum(m *TestTable, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.TestTable.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.policy.v1.TestTable.description"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Description))

	}
	if _, ok := ignore["cerbos.policy.v1.TestTable.skip"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.Skip)))

	}
	if _, ok := ignore["cerbos.policy.v1.TestTable.skip_reason"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.SkipReason))

	}
	if _, ok := ignore["cerbos.policy.v1.TestTable.input"]; !ok {
		if m.Input != nil {
			cerbos_policy_v1_TestTable_Input_hashpb_sum(m.Input, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.TestTable.expected"]; !ok {
		if len(m.Expected) > 0 {
			for _, v := range m.Expected {
				if v != nil {
					cerbos_policy_v1_TestTable_Expectation_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.TestTable.options"]; !ok {
		if m.Options != nil {
			cerbos_policy_v1_TestOptions_hashpb_sum(m.Options, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_Test_TestName_hashpb_sum(m *Test_TestName, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Test.TestName.test_table_name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.TestTableName))

	}
	if _, ok := ignore["cerbos.policy.v1.Test.TestName.principal_key"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PrincipalKey))

	}
	if _, ok := ignore["cerbos.policy.v1.Test.TestName.resource_key"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.ResourceKey))

	}
}

func cerbos_policy_v1_Test_hashpb_sum(m *Test, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Test.name"]; !ok {
		if m.Name != nil {
			cerbos_policy_v1_Test_TestName_hashpb_sum(m.Name, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.Test.description"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Description))

	}
	if _, ok := ignore["cerbos.policy.v1.Test.skip"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.Skip)))

	}
	if _, ok := ignore["cerbos.policy.v1.Test.skip_reason"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.SkipReason))

	}
	if _, ok := ignore["cerbos.policy.v1.Test.input"]; !ok {
		if m.Input != nil {
			cerbos_engine_v1_CheckInput_hashpb_sum(m.Input, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.Test.expected"]; !ok {
		if len(m.Expected) > 0 {
			keys := make([]string, len(m.Expected))
			i := 0
			for k := range m.Expected {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Expected[k])))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.Test.options"]; !ok {
		if m.Options != nil {
			cerbos_policy_v1_TestOptions_hashpb_sum(m.Options, hasher, ignore)
		}

	}
}

func google_protobuf_ListValue_hashpb_sum(m *structpb.ListValue, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.ListValue.values"]; !ok {
		if len(m.Values) > 0 {
			for _, v := range m.Values {
				if v != nil {
					google_protobuf_Value_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Struct_hashpb_sum(m *structpb.Struct, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Struct.fields"]; !ok {
		if len(m.Fields) > 0 {
			keys := make([]string, len(m.Fields))
			i := 0
			for k := range m.Fields {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Fields[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Fields[k], hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Timestamp_hashpb_sum(m *timestamppb.Timestamp, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Timestamp.seconds"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Seconds)))

	}
	if _, ok := ignore["google.protobuf.Timestamp.nanos"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Nanos)))

	}
}

func google_protobuf_UInt64Value_hashpb_sum(m *wrapperspb.UInt64Value, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.UInt64Value.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, m.Value))

	}
}

func google_protobuf_Value_hashpb_sum(m *structpb.Value, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Kind != nil {
		if _, ok := ignore["google.protobuf.Value.kind"]; !ok {
			switch t := m.Kind.(type) {
			case *structpb.Value_NullValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.NullValue)))

			case *structpb.Value_NumberValue:
				_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(t.NumberValue)))

			case *structpb.Value_StringValue:
				_, _ = hasher.Write(protowire.AppendString(nil, t.StringValue))

			case *structpb.Value_BoolValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(t.BoolValue)))

			case *structpb.Value_StructValue:
				if t.StructValue != nil {
					google_protobuf_Struct_hashpb_sum(t.StructValue, hasher, ignore)
				}

			case *structpb.Value_ListValue:
				if t.ListValue != nil {
					google_protobuf_ListValue_hashpb_sum(t.ListValue, hasher, ignore)
				}

			}
		}
	}
}
