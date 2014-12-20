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

package portletpreferences

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) DeleteArchivedPreferences(portletItemId int64) error {
	_params := make(map[string]interface{})

	_params["portletItemId"] = portletItemId

	_cmd := map[string]interface{}{
		"/portletpreferences/delete-archived-preferences": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RestoreArchivedPreferences(groupId int64, layout *liferay.ObjectWrapper, portletId string, portletItem *liferay.ObjectWrapper, preferences *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	liferay.MangleObjectWrapper(_params, "layout", "com.liferay.portal.model.Layout", layout)
	_params["portletId"] = portletId
	liferay.MangleObjectWrapper(_params, "portletItem", "com.liferay.portal.model.PortletItem", portletItem)
	liferay.MangleObjectWrapper(_params, "preferences", "javax.portlet.PortletPreferences", preferences)

	_cmd := map[string]interface{}{
		"/portletpreferences/restore-archived-preferences": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RestoreArchivedPreferences2(groupId int64, layout *liferay.ObjectWrapper, portletId string, portletItemId int64, preferences *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	liferay.MangleObjectWrapper(_params, "layout", "com.liferay.portal.model.Layout", layout)
	_params["portletId"] = portletId
	_params["portletItemId"] = portletItemId
	liferay.MangleObjectWrapper(_params, "preferences", "javax.portlet.PortletPreferences", preferences)

	_cmd := map[string]interface{}{
		"/portletpreferences/restore-archived-preferences": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RestoreArchivedPreferences3(groupId int64, name string, layout *liferay.ObjectWrapper, portletId string, preferences *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	liferay.MangleObjectWrapper(_params, "layout", "com.liferay.portal.model.Layout", layout)
	_params["portletId"] = portletId
	liferay.MangleObjectWrapper(_params, "preferences", "javax.portlet.PortletPreferences", preferences)

	_cmd := map[string]interface{}{
		"/portletpreferences/restore-archived-preferences": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateArchivePreferences(userId int64, groupId int64, name string, portletId string, preferences *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["groupId"] = groupId
	_params["name"] = name
	_params["portletId"] = portletId
	liferay.MangleObjectWrapper(_params, "preferences", "javax.portlet.PortletPreferences", preferences)

	_cmd := map[string]interface{}{
		"/portletpreferences/update-archive-preferences": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

