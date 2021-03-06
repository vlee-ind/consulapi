package consulapitest

import consulapi "github.com/shoenig/consulapi"
import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Datacenters provides a mock function with given fields:
func (_m *Client) Datacenters() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: dc, path
func (_m *Client) Delete(dc string, path string) error {
	ret := _m.Called(dc, path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(dc, path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: dc, path
func (_m *Client) Get(dc string, path string) (string, error) {
	ret := _m.Called(dc, path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(dc, path)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(dc, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Keys provides a mock function with given fields: dc, path
func (_m *Client) Keys(dc string, path string) ([]string, error) {
	ret := _m.Called(dc, path)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(dc, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(dc, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaintenanceMode provides a mock function with given fields: enabled, reason
func (_m *Client) MaintenanceMode(enabled bool, reason string) error {
	ret := _m.Called(enabled, reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, string) error); ok {
		r0 = rf(enabled, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Members provides a mock function with given fields: wan
func (_m *Client) Members(wan bool) ([]consulapi.AgentInfo, error) {
	ret := _m.Called(wan)

	var r0 []consulapi.AgentInfo
	if rf, ok := ret.Get(0).(func(bool) []consulapi.AgentInfo); ok {
		r0 = rf(wan)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]consulapi.AgentInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(wan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Node provides a mock function with given fields: dc, name
func (_m *Client) Node(dc string, name string) (consulapi.NodeInfo, error) {
	ret := _m.Called(dc, name)

	var r0 consulapi.NodeInfo
	if rf, ok := ret.Get(0).(func(string, string) consulapi.NodeInfo); ok {
		r0 = rf(dc, name)
	} else {
		r0 = ret.Get(0).(consulapi.NodeInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(dc, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Nodes provides a mock function with given fields: dc
func (_m *Client) Nodes(dc string) ([]consulapi.Node, error) {
	ret := _m.Called(dc)

	var r0 []consulapi.Node
	if rf, ok := ret.Get(0).(func(string) []consulapi.Node); ok {
		r0 = rf(dc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]consulapi.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: dc, path, value
func (_m *Client) Put(dc string, path string, value string) error {
	ret := _m.Called(dc, path, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(dc, path, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Recurse provides a mock function with given fields: dc, path
func (_m *Client) Recurse(dc string, path string) ([][2]string, error) {
	ret := _m.Called(dc, path)

	var r0 [][2]string
	if rf, ok := ret.Get(0).(func(string, string) [][2]string); ok {
		r0 = rf(dc, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][2]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(dc, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reload provides a mock function with given fields:
func (_m *Client) Reload() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service provides a mock function with given fields: dc, service, tags
func (_m *Client) Service(dc string, service string, tags ...string) ([]consulapi.Service, error) {
	_va := make([]interface{}, len(tags))
	for _i := range tags {
		_va[_i] = tags[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, dc, service)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []consulapi.Service
	if rf, ok := ret.Get(0).(func(string, string, ...string) []consulapi.Service); ok {
		r0 = rf(dc, service, tags...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]consulapi.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, ...string) error); ok {
		r1 = rf(dc, service, tags...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Services provides a mock function with given fields: dc
func (_m *Client) Services(dc string) (map[string][]string, error) {
	ret := _m.Called(dc)

	var r0 map[string][]string
	if rf, ok := ret.Get(0).(func(string) map[string][]string); ok {
		r0 = rf(dc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
