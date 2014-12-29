// Copyright (c) 2014 Andrea Di Giorgi. All rights reserved.
//
// This library is free software; you can redistribute it and/or modify it under
// the terms of the GNU Lesser General Public License as published by the Free
// Software Foundation; either version 2.1 of the License, or (at your option)
// any later version.
//
// This library is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more
// details.

package liferay

const (
	orderByComparator string = "com.liferay.portal.kernel.util.OrderByComparator"
	serviceContext    string = "com.liferay.portal.service.ServiceContext"
)

type ObjectWrapper struct {
	ClassName string
	Object    map[string]interface{}
}

func (w ObjectWrapper) addClassName(params map[string]interface{}, name string, className string) {
	if len(w.ClassName) > 0 {
		className = w.ClassName
	}

	params[("+" + name)] = className
}

func (w ObjectWrapper) addObject(params map[string]interface{}, name string) {
	for k, v := range w.Object {
		params[(name + "." + k)] = v
	}
}

func (w ObjectWrapper) mangle(params map[string]interface{}, name string, className string) {
	w.addClassName(params, name, className)
	w.addObject(params, name)
}

func MangleObjectWrapper(params map[string]interface{}, name string, className string, wrapper *ObjectWrapper) {
	if wrapper == nil {
		if className != serviceContext {
			params[name] = nil
		}

		return
	}

	wrapper.mangle(params, name, className)
}
