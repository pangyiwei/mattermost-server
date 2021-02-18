// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// RetentionPolicyStore is an autogenerated mock type for the RetentionPolicyStore type
type RetentionPolicyStore struct {
	mock.Mock
}

// AddChannels provides a mock function with given fields: policyId, channelIds
func (_m *RetentionPolicyStore) AddChannels(policyId string, channelIds []string) error {
	ret := _m.Called(policyId, channelIds)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(policyId, channelIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddTeams provides a mock function with given fields: policyId, teamIds
func (_m *RetentionPolicyStore) AddTeams(policyId string, teamIds []string) error {
	ret := _m.Called(policyId, teamIds)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(policyId, teamIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *RetentionPolicyStore) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *RetentionPolicyStore) Get(id string) (*model.RetentionPolicyWithTeamsAndChannels, error) {
	ret := _m.Called(id)

	var r0 *model.RetentionPolicyWithTeamsAndChannels
	if rf, ok := ret.Get(0).(func(string) *model.RetentionPolicyWithTeamsAndChannels); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RetentionPolicyWithTeamsAndChannels)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: offset, limit
func (_m *RetentionPolicyStore) GetAll(offset uint64, limit uint64) ([]*model.RetentionPolicyWithTeamsAndChannels, error) {
	ret := _m.Called(offset, limit)

	var r0 []*model.RetentionPolicyWithTeamsAndChannels
	if rf, ok := ret.Get(0).(func(uint64, uint64) []*model.RetentionPolicyWithTeamsAndChannels); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RetentionPolicyWithTeamsAndChannels)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllWithCounts provides a mock function with given fields: offset, limit
func (_m *RetentionPolicyStore) GetAllWithCounts(offset uint64, limit uint64) ([]*model.RetentionPolicyWithTeamAndChannelCounts, error) {
	ret := _m.Called(offset, limit)

	var r0 []*model.RetentionPolicyWithTeamAndChannelCounts
	if rf, ok := ret.Get(0).(func(uint64, uint64) []*model.RetentionPolicyWithTeamAndChannelCounts); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RetentionPolicyWithTeamAndChannelCounts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: patch
func (_m *RetentionPolicyStore) Patch(patch *model.RetentionPolicyWithTeamAndChannelIds) (*model.RetentionPolicyWithTeamsAndChannels, error) {
	ret := _m.Called(patch)

	var r0 *model.RetentionPolicyWithTeamsAndChannels
	if rf, ok := ret.Get(0).(func(*model.RetentionPolicyWithTeamAndChannelIds) *model.RetentionPolicyWithTeamsAndChannels); ok {
		r0 = rf(patch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RetentionPolicyWithTeamsAndChannels)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.RetentionPolicyWithTeamAndChannelIds) error); ok {
		r1 = rf(patch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveChannels provides a mock function with given fields: policyId, channelIds
func (_m *RetentionPolicyStore) RemoveChannels(policyId string, channelIds []string) error {
	ret := _m.Called(policyId, channelIds)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(policyId, channelIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveOrphanedRows provides a mock function with given fields: limit
func (_m *RetentionPolicyStore) RemoveOrphanedRows(limit int64) (int64, error) {
	ret := _m.Called(limit)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveTeams provides a mock function with given fields: policyId, teamIds
func (_m *RetentionPolicyStore) RemoveTeams(policyId string, teamIds []string) error {
	ret := _m.Called(policyId, teamIds)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(policyId, teamIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: policy
func (_m *RetentionPolicyStore) Save(policy *model.RetentionPolicyWithTeamAndChannelIds) (*model.RetentionPolicyWithTeamsAndChannels, error) {
	ret := _m.Called(policy)

	var r0 *model.RetentionPolicyWithTeamsAndChannels
	if rf, ok := ret.Get(0).(func(*model.RetentionPolicyWithTeamAndChannelIds) *model.RetentionPolicyWithTeamsAndChannels); ok {
		r0 = rf(policy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RetentionPolicyWithTeamsAndChannels)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.RetentionPolicyWithTeamAndChannelIds) error); ok {
		r1 = rf(policy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: update
func (_m *RetentionPolicyStore) Update(update *model.RetentionPolicyWithTeamAndChannelIds) (*model.RetentionPolicyWithTeamsAndChannels, error) {
	ret := _m.Called(update)

	var r0 *model.RetentionPolicyWithTeamsAndChannels
	if rf, ok := ret.Get(0).(func(*model.RetentionPolicyWithTeamAndChannelIds) *model.RetentionPolicyWithTeamsAndChannels); ok {
		r0 = rf(update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RetentionPolicyWithTeamsAndChannels)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.RetentionPolicyWithTeamAndChannelIds) error); ok {
		r1 = rf(update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}