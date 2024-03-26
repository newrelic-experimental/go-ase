// SPDX-FileCopyrightText: 2020 SAP SE
// SPDX-FileCopyrightText: 2021 SAP SE
// SPDX-FileCopyrightText: 2022 SAP SE
// SPDX-FileCopyrightText: 2023 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

module github.com/newrelic-experimental/go-ase

replace github.com/SAP/go-dblib v0.0.0-20220825075032-c1f3f4d6e7b3 => github.com/newrelic-experimental/go-dblib v0.0.0-20240321215801-9bfaa7c8a147

go 1.19

require (
	github.com/SAP/go-dblib v0.0.0-20220825075032-c1f3f4d6e7b3
	github.com/spf13/pflag v1.0.5
)

require (
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
)
