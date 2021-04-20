// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package create

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (ExecutionConfig) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (ExecutionConfig) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in ExecutionConfig and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg ExecutionConfig) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("ExecutionConfig", pflag.ExitOnError)
	cmdFlags.StringVar(&(executionConfig.ExecFile), fmt.Sprintf("%v%v", prefix, "execFile"), executionConfig.ExecFile, "file for the execution params.If not specified defaults to <<workflow/task>_name>.execution_spec.yaml")
	cmdFlags.StringVar(&(executionConfig.TargetDomain), fmt.Sprintf("%v%v", prefix, "targetDomain"), executionConfig.TargetDomain, "project where execution needs to be created.If not specified configured domain would be used.")
	cmdFlags.StringVar(&(executionConfig.TargetProject), fmt.Sprintf("%v%v", prefix, "targetProject"), executionConfig.TargetProject, "project where execution needs to be created.If not specified configured project would be used.")
	cmdFlags.StringVar(&(executionConfig.KubeServiceAcct), fmt.Sprintf("%v%v", prefix, "kubeServiceAcct"), executionConfig.KubeServiceAcct, "kubernetes service account AuthRole for launching execution.")
	cmdFlags.StringVar(&(executionConfig.IamRoleARN), fmt.Sprintf("%v%v", prefix, "iamRoleARN"), executionConfig.IamRoleARN, "iam role ARN AuthRole for launching execution.")
	cmdFlags.StringVar(&(executionConfig.Relaunch), fmt.Sprintf("%v%v", prefix, "relaunch"), executionConfig.Relaunch, "execution id to be relaunched.")

	return cmdFlags
}