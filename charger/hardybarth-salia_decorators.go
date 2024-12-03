package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateSalia(base *Salia, meter func() (float64, error), meterEnergy func() (float64, error), phaseCurrents func() (float64, float64, float64, error), phaseSwitcher func(int) error, phaseGetter func() (int, error)) api.Charger {
	switch {
	case meter == nil && phaseSwitcher == nil:
		return base

	case meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseSwitcher == nil:
		return &struct {
			*Salia
			api.Meter
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseSwitcher == nil:
		return &struct {
			*Salia
			api.Meter
			api.MeterEnergy
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateSaliaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseSwitcher == nil:
		return &struct {
			*Salia
			api.Meter
			api.PhaseCurrents
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateSaliaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseSwitcher == nil:
		return &struct {
			*Salia
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateSaliaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateSaliaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case meter == nil && phaseGetter == nil && phaseSwitcher != nil:
		return &struct {
			*Salia
			api.PhaseSwitcher
		}{
			Salia: base,
			PhaseSwitcher: &decorateSaliaPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseGetter == nil && phaseSwitcher != nil:
		return &struct {
			*Salia
			api.Meter
			api.PhaseSwitcher
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			PhaseSwitcher: &decorateSaliaPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseGetter == nil && phaseSwitcher != nil:
		return &struct {
			*Salia
			api.Meter
			api.MeterEnergy
			api.PhaseSwitcher
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateSaliaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseSwitcher: &decorateSaliaPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseGetter == nil && phaseSwitcher != nil:
		return &struct {
			*Salia
			api.Meter
			api.PhaseCurrents
			api.PhaseSwitcher
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateSaliaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseSwitcher: &decorateSaliaPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseGetter == nil && phaseSwitcher != nil:
		return &struct {
			*Salia
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseSwitcher
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateSaliaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateSaliaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseSwitcher: &decorateSaliaPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter == nil && phaseGetter != nil && phaseSwitcher != nil:
		return &struct {
			*Salia
			api.PhaseGetter
			api.PhaseSwitcher
		}{
			Salia: base,
			PhaseGetter: &decorateSaliaPhaseGetterImpl{
				phaseGetter: phaseGetter,
			},
			PhaseSwitcher: &decorateSaliaPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseGetter != nil && phaseSwitcher != nil:
		return &struct {
			*Salia
			api.Meter
			api.PhaseGetter
			api.PhaseSwitcher
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			PhaseGetter: &decorateSaliaPhaseGetterImpl{
				phaseGetter: phaseGetter,
			},
			PhaseSwitcher: &decorateSaliaPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseGetter != nil && phaseSwitcher != nil:
		return &struct {
			*Salia
			api.Meter
			api.MeterEnergy
			api.PhaseGetter
			api.PhaseSwitcher
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateSaliaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseGetter: &decorateSaliaPhaseGetterImpl{
				phaseGetter: phaseGetter,
			},
			PhaseSwitcher: &decorateSaliaPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseGetter != nil && phaseSwitcher != nil:
		return &struct {
			*Salia
			api.Meter
			api.PhaseCurrents
			api.PhaseGetter
			api.PhaseSwitcher
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateSaliaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseGetter: &decorateSaliaPhaseGetterImpl{
				phaseGetter: phaseGetter,
			},
			PhaseSwitcher: &decorateSaliaPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseGetter != nil && phaseSwitcher != nil:
		return &struct {
			*Salia
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseGetter
			api.PhaseSwitcher
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateSaliaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateSaliaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseGetter: &decorateSaliaPhaseGetterImpl{
				phaseGetter: phaseGetter,
			},
			PhaseSwitcher: &decorateSaliaPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}
	}

	return nil
}

type decorateSaliaMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateSaliaMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decorateSaliaMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateSaliaMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decorateSaliaPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateSaliaPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}

type decorateSaliaPhaseGetterImpl struct {
	phaseGetter func() (int, error)
}

func (impl *decorateSaliaPhaseGetterImpl) GetPhases() (int, error) {
	return impl.phaseGetter()
}

type decorateSaliaPhaseSwitcherImpl struct {
	phaseSwitcher func(int) error
}

func (impl *decorateSaliaPhaseSwitcherImpl) Phases1p3p(p0 int) error {
	return impl.phaseSwitcher(p0)
}
