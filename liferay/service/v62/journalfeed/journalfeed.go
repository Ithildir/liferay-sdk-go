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

package journalfeed

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddFeed(groupId int64, feedId string, autoFeedId bool, name string, description string, type string, structureId string, templateId string, rendererTemplateId string, delta int, orderByCol string, orderByType string, targetLayoutFriendlyUrl string, targetPortletId string, contentField string, feedType string, feedVersion float64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["feedId"] = feedId
	_params["autoFeedId"] = autoFeedId
	_params["name"] = name
	_params["description"] = description
	_params["type"] = type
	_params["structureId"] = structureId
	_params["templateId"] = templateId
	_params["rendererTemplateId"] = rendererTemplateId
	_params["delta"] = delta
	_params["orderByCol"] = orderByCol
	_params["orderByType"] = orderByType
	_params["targetLayoutFriendlyUrl"] = targetLayoutFriendlyUrl
	_params["targetPortletId"] = targetPortletId
	_params["contentField"] = contentField
	_params["feedType"] = feedType
	_params["feedVersion"] = feedVersion
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalfeed/add-feed": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteFeed(feedId int64) error {
	_params := make(map[string]interface{})

	_params["feedId"] = feedId

	_cmd := map[string]interface{}{
		"/journalfeed/delete-feed": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteFeed2(groupId int64, feedId string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["feedId"] = feedId

	_cmd := map[string]interface{}{
		"/journalfeed/delete-feed": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetFeed(feedId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["feedId"] = feedId

	_cmd := map[string]interface{}{
		"/journalfeed/get-feed": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFeed2(groupId int64, feedId string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["feedId"] = feedId

	_cmd := map[string]interface{}{
		"/journalfeed/get-feed": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateFeed(groupId int64, feedId string, name string, description string, type string, structureId string, templateId string, rendererTemplateId string, delta int, orderByCol string, orderByType string, targetLayoutFriendlyUrl string, targetPortletId string, contentField string, feedType string, feedVersion float64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["feedId"] = feedId
	_params["name"] = name
	_params["description"] = description
	_params["type"] = type
	_params["structureId"] = structureId
	_params["templateId"] = templateId
	_params["rendererTemplateId"] = rendererTemplateId
	_params["delta"] = delta
	_params["orderByCol"] = orderByCol
	_params["orderByType"] = orderByType
	_params["targetLayoutFriendlyUrl"] = targetLayoutFriendlyUrl
	_params["targetPortletId"] = targetPortletId
	_params["contentField"] = contentField
	_params["feedType"] = feedType
	_params["feedVersion"] = feedVersion
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalfeed/update-feed": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

