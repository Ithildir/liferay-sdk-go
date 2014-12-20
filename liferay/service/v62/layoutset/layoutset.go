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

package layoutset

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) UpdateLayoutSetPrototypeLinkEnabled(groupId int64, privateLayout bool, layoutSetPrototypeLinkEnabled bool, layoutSetPrototypeUuid string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutSetPrototypeLinkEnabled"] = layoutSetPrototypeLinkEnabled
	_params["layoutSetPrototypeUuid"] = layoutSetPrototypeUuid

	_cmd := map[string]interface{}{
		"/layoutset/update-layout-set-prototype-link-enabled": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateLogo(groupId int64, privateLayout bool, logo bool, bytes []byte) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["logo"] = logo
	_params["bytes"] = bytes

	_cmd := map[string]interface{}{
		"/layoutset/update-logo": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateLogo2(groupId int64, privateLayout bool, logo bool, file io.Reader) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["logo"] = logo
	_params["file"] = file

	_cmd := map[string]interface{}{
		"/layoutset/update-logo": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateLookAndFeel(groupId int64, privateLayout bool, themeId string, colorSchemeId string, css string, wapTheme bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["themeId"] = themeId
	_params["colorSchemeId"] = colorSchemeId
	_params["css"] = css
	_params["wapTheme"] = wapTheme

	_cmd := map[string]interface{}{
		"/layoutset/update-look-and-feel": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateSettings(groupId int64, privateLayout bool, settings string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["settings"] = settings

	_cmd := map[string]interface{}{
		"/layoutset/update-settings": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateVirtualHost(groupId int64, privateLayout bool, virtualHost string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["virtualHost"] = virtualHost

	_cmd := map[string]interface{}{
		"/layoutset/update-virtual-host": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

