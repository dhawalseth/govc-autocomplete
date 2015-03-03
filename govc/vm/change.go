/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vm

import (
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type change struct {
	*flags.VirtualMachineFlag

	types.VirtualMachineConfigSpec
}

func init() {
	cli.Register("vm.change", &change{})
}

func (cmd *change) Register(f *flag.FlagSet) {
	f.Int64Var(&cmd.MemoryMB, "m", 0, "Size in MB of memory")
	f.IntVar(&cmd.NumCPUs, "c", 0, "Number of CPUs")
	f.StringVar(&cmd.GuestId, "g", "", "Guest OS")
	f.StringVar(&cmd.Name, "name", "", "Display name")
}

func (cmd *change) Process() error { return nil }

func (cmd *change) Run(f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	task, err := vm.Reconfigure(cmd.VirtualMachineConfigSpec)
	if err != nil {
		return err
	}

	return task.Wait()
}
