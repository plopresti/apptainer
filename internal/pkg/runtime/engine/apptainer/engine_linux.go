// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2018-2019, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package apptainer

import (
	"github.com/apptainer/apptainer/internal/pkg/runtime/engine"
	"github.com/apptainer/apptainer/internal/pkg/runtime/engine/apptainer/rpc/server"
	apptainerConfig "github.com/apptainer/apptainer/pkg/runtime/engine/apptainer/config"
	"github.com/apptainer/apptainer/pkg/runtime/engine/config"
)

// EngineOperations is a Apptainer runtime engine that implements engine.Operations.
// Basically, this is the core of `apptainer run/exec/shell/instance` commands.
type EngineOperations struct {
	CommonConfig *config.Common                `json:"-"`
	EngineConfig *apptainerConfig.EngineConfig `json:"engineConfig"`
}

// InitConfig stores the parsed config.Common inside the engine.
//
// Since this method simply stores config.Common, it does not matter
// whether or not there are any elevated privileges during this call.
func (e *EngineOperations) InitConfig(cfg *config.Common) {
	e.CommonConfig = cfg
}

// Config returns a pointer to a apptainerConfig.EngineConfig
// literal as a config.EngineConfig interface. This pointer
// gets stored in the engine.Engine.Common field.
//
// Since this method simply returns a zero value of the concrete
// EngineConfig, it does not matter whether or not there are any elevated
// privileges during this call.
func (e *EngineOperations) Config() config.EngineConfig {
	return e.EngineConfig
}

func init() {
	engine.RegisterOperations(
		apptainerConfig.Name,
		&EngineOperations{
			EngineConfig: apptainerConfig.NewConfig(),
		},
	)

	engine.RegisterRPCMethods(
		apptainerConfig.Name,
		new(server.Methods),
	)
}