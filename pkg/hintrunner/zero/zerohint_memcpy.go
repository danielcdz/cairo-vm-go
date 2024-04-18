package zero

import (
	"github.com/NethermindEth/cairo-vm-go/pkg/hintrunner/core"
	"github.com/NethermindEth/cairo-vm-go/pkg/hintrunner/hinter"
	"github.com/NethermindEth/cairo-vm-go/pkg/utils"
	"github.com/NethermindEth/cairo-vm-go/pkg/vm/memory"
	"github.com/consensys/gnark-crypto/ecc/stark-curve/fp"

	VM "github.com/NethermindEth/cairo-vm-go/pkg/vm"
)

func newMemcpyContinueCopyingHinter(n, dst hinter.ResOperander) hinter.Hinter {
	return &GenericZeroHinter{
		Name: "MemcpyContinueCopying",
		Op: func(vm *VM.VirtualMachine, ctx *hinter.HintRunnerContext) error {
			//n -= 1
			//ids.continue_copying = 1 if n > 0 else 0
			var lhs fp.Element

			// n-=1
			n, err := ctx.ScopeManager.GetVariableValue("n")
			if err != nil {
				return err
			}
			ctx.ScopeManager.AssignVariable("n", lhs.Sub(n.(*fp.Element), &utils.FeltOne))

			// ids.continue_copying = 1 if n > 0 else 0
			continue_copying, err := dst.GetAddress(vm)
			if err != nil {
				return err
			}

			var v memory.MemoryValue
			if !utils.FeltIsPositive(n.(*fp.Element)) {
				v = memory.MemoryValueFromFieldElement(&utils.FeltOne)
			} else {
				v = memory.MemoryValueFromFieldElement(&utils.FeltZero)
			}
			return vm.Memory.WriteToAddress(&continue_copying, &v)
		},
	}
}

func createMemcpyContinueCopyingHinter(resolver hintReferenceResolver) (hinter.Hinter, error) {
	 n, err := resolver.GetResOperander("n")
	if err != nil {
		return nil, err
	}
	output, err := resolver.GetResOperander("output")
	if err != nil {
		return nil, err
	}
	return newMemcpyContinueCopyingHinter(n, output), nil
}


func createAllocSegmentHinter(resolver hintReferenceResolver) (hinter.Hinter, error) {
	return &core.AllocSegment{Dst: hinter.ApCellRef(0)}, nil
}

func createVMEnterScopeHinter(resolver hintReferenceResolver) (hinter.Hinter, error) {
	return &GenericZeroHinter{
		Name: "VMEnterScope",
		Op: func(vm *VM.VirtualMachine, ctx *hinter.HintRunnerContext) error {
			ctx.ScopeManager.EnterScope(make(map[string]any))
			return nil
		},
	}, nil
}

func createVMExitScopeHinter(resolver hintReferenceResolver) (hinter.Hinter, error) {
	return &GenericZeroHinter{
		Name: "VMExitScope",
		Op: func(vm *VM.VirtualMachine, ctx *hinter.HintRunnerContext) error {
			return ctx.ScopeManager.ExitScope()
		},
	}, nil
}

