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

package journaltemplate

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddTemplate(groupId int64, templateId string, autoTemplateId bool, structureId string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, xsl string, formatXsl bool, langType string, cacheable bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["templateId"] = templateId
	_params["autoTemplateId"] = autoTemplateId
	_params["structureId"] = structureId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["xsl"] = xsl
	_params["formatXsl"] = formatXsl
	_params["langType"] = langType
	_params["cacheable"] = cacheable
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journaltemplate/add-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddTemplate2(groupId int64, templateId string, autoTemplateId bool, structureId string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, xsl string, formatXsl bool, langType string, cacheable bool, smallImage bool, smallImageURL string, smallFile io.Reader, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["templateId"] = templateId
	_params["autoTemplateId"] = autoTemplateId
	_params["structureId"] = structureId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["xsl"] = xsl
	_params["formatXsl"] = formatXsl
	_params["langType"] = langType
	_params["cacheable"] = cacheable
	_params["smallImage"] = smallImage
	_params["smallImageURL"] = smallImageURL
	_params["smallFile"] = smallFile
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journaltemplate/add-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CopyTemplate(groupId int64, oldTemplateId string, newTemplateId string, autoTemplateId bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["oldTemplateId"] = oldTemplateId
	_params["newTemplateId"] = newTemplateId
	_params["autoTemplateId"] = autoTemplateId

	_cmd := map[string]interface{}{
		"/journaltemplate/copy-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteTemplate(groupId int64, templateId string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["templateId"] = templateId

	_cmd := map[string]interface{}{
		"/journaltemplate/delete-template": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetStructureTemplates(groupId int64, structureId string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["structureId"] = structureId

	_cmd := map[string]interface{}{
		"/journaltemplate/get-structure-templates": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTemplate(groupId int64, templateId string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["templateId"] = templateId

	_cmd := map[string]interface{}{
		"/journaltemplate/get-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetTemplate2(groupId int64, templateId string, includeGlobalTemplates bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["templateId"] = templateId
	_params["includeGlobalTemplates"] = includeGlobalTemplates

	_cmd := map[string]interface{}{
		"/journaltemplate/get-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) Search(companyId int64, groupIds []interface{}, templateId string, structureId string, structureIdComparator string, name string, description string, andOperator bool, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["templateId"] = templateId
	_params["structureId"] = structureId
	_params["structureIdComparator"] = structureIdComparator
	_params["name"] = name
	_params["description"] = description
	_params["andOperator"] = andOperator
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journaltemplate/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search2(companyId int64, groupIds []interface{}, keywords string, structureId string, structureIdComparator string, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["keywords"] = keywords
	_params["structureId"] = structureId
	_params["structureIdComparator"] = structureIdComparator
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journaltemplate/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) SearchCount(companyId int64, groupIds []interface{}, keywords string, structureId string, structureIdComparator string) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["keywords"] = keywords
	_params["structureId"] = structureId
	_params["structureIdComparator"] = structureIdComparator

	_cmd := map[string]interface{}{
		"/journaltemplate/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) SearchCount2(companyId int64, groupIds []interface{}, templateId string, structureId string, structureIdComparator string, name string, description string, andOperator bool) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["templateId"] = templateId
	_params["structureId"] = structureId
	_params["structureIdComparator"] = structureIdComparator
	_params["name"] = name
	_params["description"] = description
	_params["andOperator"] = andOperator

	_cmd := map[string]interface{}{
		"/journaltemplate/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) UpdateTemplate(groupId int64, templateId string, structureId string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, xsl string, formatXsl bool, langType string, cacheable bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["templateId"] = templateId
	_params["structureId"] = structureId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["xsl"] = xsl
	_params["formatXsl"] = formatXsl
	_params["langType"] = langType
	_params["cacheable"] = cacheable
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journaltemplate/update-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateTemplate2(groupId int64, templateId string, structureId string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, xsl string, formatXsl bool, langType string, cacheable bool, smallImage bool, smallImageURL string, smallFile io.Reader, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["templateId"] = templateId
	_params["structureId"] = structureId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["xsl"] = xsl
	_params["formatXsl"] = formatXsl
	_params["langType"] = langType
	_params["cacheable"] = cacheable
	_params["smallImage"] = smallImage
	_params["smallImageURL"] = smallImageURL
	_params["smallFile"] = smallFile
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journaltemplate/update-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

