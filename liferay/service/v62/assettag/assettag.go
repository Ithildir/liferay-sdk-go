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

package assettag

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddTag(name string, tagProperties []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name
	_params["tagProperties"] = tagProperties
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assettag/add-tag": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteTag(tagId int64) error {
	_params := make(map[string]interface{})

	_params["tagId"] = tagId

	_cmd := map[string]interface{}{
		"/assettag/delete-tag": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteTags(tagIds []interface{}) error {
	_params := make(map[string]interface{})

	_params["tagIds"] = tagIds

	_cmd := map[string]interface{}{
		"/assettag/delete-tags": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetGroupTags(groupId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/assettag/get-group-tags": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupTags2(groupId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assettag/get-group-tags": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupTagsCount(groupId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/assettag/get-group-tags-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetGroupTagsDisplay(groupId int64, name string, start int, end int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/assettag/get-group-tags-display": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetGroupsTags(groupIds []interface{}) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupIds"] = groupIds

	_cmd := map[string]interface{}{
		"/assettag/get-groups-tags": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetJsonGroupTags(groupId int64, name string, start int, end int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/assettag/get-json-group-tags": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetTag(tagId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["tagId"] = tagId

	_cmd := map[string]interface{}{
		"/assettag/get-tag": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetTags(className string, classPK int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/assettag/get-tags": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTags2(groupId int64, classNameId int64, name string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/assettag/get-tags": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTags3(groupId int64, name string, tagProperties []interface{}, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["tagProperties"] = tagProperties
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/assettag/get-tags": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTags4(groupIds []interface{}, name string, tagProperties []interface{}, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupIds"] = groupIds
	_params["name"] = name
	_params["tagProperties"] = tagProperties
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/assettag/get-tags": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTags5(groupId int64, classNameId int64, name string, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["name"] = name
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/assettag/get-tags": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTagsCount(groupId int64, name string) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/assettag/get-tags-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetTagsCount2(groupId int64, classNameId int64, name string) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["name"] = name

	_cmd := map[string]interface{}{
		"/assettag/get-tags-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetTagsCount3(groupId int64, name string, tagProperties []interface{}) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["tagProperties"] = tagProperties

	_cmd := map[string]interface{}{
		"/assettag/get-tags-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) MergeTags(fromTagId int64, toTagId int64, overrideProperties bool) error {
	_params := make(map[string]interface{})

	_params["fromTagId"] = fromTagId
	_params["toTagId"] = toTagId
	_params["overrideProperties"] = overrideProperties

	_cmd := map[string]interface{}{
		"/assettag/merge-tags": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) MergeTags2(fromTagIds []interface{}, toTagId int64, overrideProperties bool) error {
	_params := make(map[string]interface{})

	_params["fromTagIds"] = fromTagIds
	_params["toTagId"] = toTagId
	_params["overrideProperties"] = overrideProperties

	_cmd := map[string]interface{}{
		"/assettag/merge-tags": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) Search(groupId int64, name string, tagProperties []interface{}, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["name"] = name
	_params["tagProperties"] = tagProperties
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/assettag/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search2(groupIds []interface{}, name string, tagProperties []interface{}, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupIds"] = groupIds
	_params["name"] = name
	_params["tagProperties"] = tagProperties
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/assettag/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdateTag(tagId int64, name string, tagProperties []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["tagId"] = tagId
	_params["name"] = name
	_params["tagProperties"] = tagProperties
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/assettag/update-tag": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

