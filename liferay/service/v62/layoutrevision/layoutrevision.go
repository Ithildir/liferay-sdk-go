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

package layoutrevision

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddLayoutRevision(userId int64, layoutSetBranchId int64, layoutBranchId int64, parentLayoutRevisionId int64, head bool, plid int64, portletPreferencesPlid int64, privateLayout bool, name string, title string, description string, keywords string, robots string, typeSettings string, iconImage bool, iconImageId int64, themeId string, colorSchemeId string, wapThemeId string, wapColorSchemeId string, css string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["layoutSetBranchId"] = layoutSetBranchId
	_params["layoutBranchId"] = layoutBranchId
	_params["parentLayoutRevisionId"] = parentLayoutRevisionId
	_params["head"] = head
	_params["plid"] = plid
	_params["portletPreferencesPlid"] = portletPreferencesPlid
	_params["privateLayout"] = privateLayout
	_params["name"] = name
	_params["title"] = title
	_params["description"] = description
	_params["keywords"] = keywords
	_params["robots"] = robots
	_params["typeSettings"] = typeSettings
	_params["iconImage"] = iconImage
	_params["iconImageId"] = iconImageId
	_params["themeId"] = themeId
	_params["colorSchemeId"] = colorSchemeId
	_params["wapThemeId"] = wapThemeId
	_params["wapColorSchemeId"] = wapColorSchemeId
	_params["css"] = css
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layoutrevision/add-layout-revision": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

