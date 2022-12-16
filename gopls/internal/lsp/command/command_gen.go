// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Don't include this file during code generation, or it will break the build
// if existing interface methods have been modified.
//go:build !generate
// +build !generate

package command

// Code generated by generate.go. DO NOT EDIT.

import (
	"context"
	"fmt"

	"golang.org/x/tools/gopls/internal/lsp/protocol"
)

const (
	AddDependency         Command = "add_dependency"
	AddImport             Command = "add_import"
	ApplyFix              Command = "apply_fix"
	CheckUpgrades         Command = "check_upgrades"
	EditGoDirective       Command = "edit_go_directive"
	FetchVulncheckResult  Command = "fetch_vulncheck_result"
	GCDetails             Command = "gc_details"
	Generate              Command = "generate"
	GoGetPackage          Command = "go_get_package"
	ListImports           Command = "list_imports"
	ListKnownPackages     Command = "list_known_packages"
	RegenerateCgo         Command = "regenerate_cgo"
	RemoveDependency      Command = "remove_dependency"
	ResetGoModDiagnostics Command = "reset_go_mod_diagnostics"
	RunGovulncheck        Command = "run_govulncheck"
	RunTests              Command = "run_tests"
	StartDebugging        Command = "start_debugging"
	Test                  Command = "test"
	Tidy                  Command = "tidy"
	ToggleGCDetails       Command = "toggle_gc_details"
	UpdateGoSum           Command = "update_go_sum"
	UpgradeDependency     Command = "upgrade_dependency"
	Vendor                Command = "vendor"
)

var Commands = []Command{
	AddDependency,
	AddImport,
	ApplyFix,
	CheckUpgrades,
	EditGoDirective,
	FetchVulncheckResult,
	GCDetails,
	Generate,
	GoGetPackage,
	ListImports,
	ListKnownPackages,
	RegenerateCgo,
	RemoveDependency,
	ResetGoModDiagnostics,
	RunGovulncheck,
	RunTests,
	StartDebugging,
	Test,
	Tidy,
	ToggleGCDetails,
	UpdateGoSum,
	UpgradeDependency,
	Vendor,
}

func Dispatch(ctx context.Context, params *protocol.ExecuteCommandParams, s Interface) (interface{}, error) {
	switch params.Command {
	case "gopls.add_dependency":
		var a0 DependencyArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.AddDependency(ctx, a0)
	case "gopls.add_import":
		var a0 AddImportArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.AddImport(ctx, a0)
	case "gopls.apply_fix":
		var a0 ApplyFixArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.ApplyFix(ctx, a0)
	case "gopls.check_upgrades":
		var a0 CheckUpgradesArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.CheckUpgrades(ctx, a0)
	case "gopls.edit_go_directive":
		var a0 EditGoDirectiveArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.EditGoDirective(ctx, a0)
	case "gopls.fetch_vulncheck_result":
		var a0 URIArg
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return s.FetchVulncheckResult(ctx, a0)
	case "gopls.gc_details":
		var a0 protocol.DocumentURI
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.GCDetails(ctx, a0)
	case "gopls.generate":
		var a0 GenerateArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.Generate(ctx, a0)
	case "gopls.go_get_package":
		var a0 GoGetPackageArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.GoGetPackage(ctx, a0)
	case "gopls.list_imports":
		var a0 URIArg
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return s.ListImports(ctx, a0)
	case "gopls.list_known_packages":
		var a0 URIArg
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return s.ListKnownPackages(ctx, a0)
	case "gopls.regenerate_cgo":
		var a0 URIArg
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.RegenerateCgo(ctx, a0)
	case "gopls.remove_dependency":
		var a0 RemoveDependencyArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.RemoveDependency(ctx, a0)
	case "gopls.reset_go_mod_diagnostics":
		var a0 ResetGoModDiagnosticsArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.ResetGoModDiagnostics(ctx, a0)
	case "gopls.run_govulncheck":
		var a0 VulncheckArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return s.RunGovulncheck(ctx, a0)
	case "gopls.run_tests":
		var a0 RunTestsArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.RunTests(ctx, a0)
	case "gopls.start_debugging":
		var a0 DebuggingArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return s.StartDebugging(ctx, a0)
	case "gopls.test":
		var a0 protocol.DocumentURI
		var a1 []string
		var a2 []string
		if err := UnmarshalArgs(params.Arguments, &a0, &a1, &a2); err != nil {
			return nil, err
		}
		return nil, s.Test(ctx, a0, a1, a2)
	case "gopls.tidy":
		var a0 URIArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.Tidy(ctx, a0)
	case "gopls.toggle_gc_details":
		var a0 URIArg
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.ToggleGCDetails(ctx, a0)
	case "gopls.update_go_sum":
		var a0 URIArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.UpdateGoSum(ctx, a0)
	case "gopls.upgrade_dependency":
		var a0 DependencyArgs
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.UpgradeDependency(ctx, a0)
	case "gopls.vendor":
		var a0 URIArg
		if err := UnmarshalArgs(params.Arguments, &a0); err != nil {
			return nil, err
		}
		return nil, s.Vendor(ctx, a0)
	}
	return nil, fmt.Errorf("unsupported command %q", params.Command)
}

func NewAddDependencyCommand(title string, a0 DependencyArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.add_dependency",
		Arguments: args,
	}, nil
}

func NewAddImportCommand(title string, a0 AddImportArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.add_import",
		Arguments: args,
	}, nil
}

func NewApplyFixCommand(title string, a0 ApplyFixArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.apply_fix",
		Arguments: args,
	}, nil
}

func NewCheckUpgradesCommand(title string, a0 CheckUpgradesArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.check_upgrades",
		Arguments: args,
	}, nil
}

func NewEditGoDirectiveCommand(title string, a0 EditGoDirectiveArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.edit_go_directive",
		Arguments: args,
	}, nil
}

func NewFetchVulncheckResultCommand(title string, a0 URIArg) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.fetch_vulncheck_result",
		Arguments: args,
	}, nil
}

func NewGCDetailsCommand(title string, a0 protocol.DocumentURI) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.gc_details",
		Arguments: args,
	}, nil
}

func NewGenerateCommand(title string, a0 GenerateArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.generate",
		Arguments: args,
	}, nil
}

func NewGoGetPackageCommand(title string, a0 GoGetPackageArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.go_get_package",
		Arguments: args,
	}, nil
}

func NewListImportsCommand(title string, a0 URIArg) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.list_imports",
		Arguments: args,
	}, nil
}

func NewListKnownPackagesCommand(title string, a0 URIArg) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.list_known_packages",
		Arguments: args,
	}, nil
}

func NewRegenerateCgoCommand(title string, a0 URIArg) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.regenerate_cgo",
		Arguments: args,
	}, nil
}

func NewRemoveDependencyCommand(title string, a0 RemoveDependencyArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.remove_dependency",
		Arguments: args,
	}, nil
}

func NewResetGoModDiagnosticsCommand(title string, a0 ResetGoModDiagnosticsArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.reset_go_mod_diagnostics",
		Arguments: args,
	}, nil
}

func NewRunGovulncheckCommand(title string, a0 VulncheckArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.run_govulncheck",
		Arguments: args,
	}, nil
}

func NewRunTestsCommand(title string, a0 RunTestsArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.run_tests",
		Arguments: args,
	}, nil
}

func NewStartDebuggingCommand(title string, a0 DebuggingArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.start_debugging",
		Arguments: args,
	}, nil
}

func NewTestCommand(title string, a0 protocol.DocumentURI, a1 []string, a2 []string) (protocol.Command, error) {
	args, err := MarshalArgs(a0, a1, a2)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.test",
		Arguments: args,
	}, nil
}

func NewTidyCommand(title string, a0 URIArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.tidy",
		Arguments: args,
	}, nil
}

func NewToggleGCDetailsCommand(title string, a0 URIArg) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.toggle_gc_details",
		Arguments: args,
	}, nil
}

func NewUpdateGoSumCommand(title string, a0 URIArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.update_go_sum",
		Arguments: args,
	}, nil
}

func NewUpgradeDependencyCommand(title string, a0 DependencyArgs) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.upgrade_dependency",
		Arguments: args,
	}, nil
}

func NewVendorCommand(title string, a0 URIArg) (protocol.Command, error) {
	args, err := MarshalArgs(a0)
	if err != nil {
		return protocol.Command{}, err
	}
	return protocol.Command{
		Title:     title,
		Command:   "gopls.vendor",
		Arguments: args,
	}, nil
}
