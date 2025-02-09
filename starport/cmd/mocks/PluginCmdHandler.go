// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	cobra "github.com/spf13/cobra"
	mock "github.com/stretchr/testify/mock"
)

// PluginCmdHandler is an autogenerated mock type for the PluginCmdHandler type
type PluginCmdHandler struct {
	mock.Mock
}

// HandleInstall provides a mock function with given fields: cmd, args
func (_m *PluginCmdHandler) HandleInstall(cmd *cobra.Command, args []string) error {
	ret := _m.Called(cmd, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(*cobra.Command, []string) error); ok {
		r0 = rf(cmd, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HandleList provides a mock function with given fields: cmd, args
func (_m *PluginCmdHandler) HandleList(cmd *cobra.Command, args []string) error {
	ret := _m.Called(cmd, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(*cobra.Command, []string) error); ok {
		r0 = rf(cmd, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
