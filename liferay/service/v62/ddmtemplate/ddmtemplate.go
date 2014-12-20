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

package ddmtemplate

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddTemplate(groupId int64, classNameId int64, classPK int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, type string, mode string, language string, script string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["type"] = type
	_params["mode"] = mode
	_params["language"] = language
	_params["script"] = script
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmtemplate/add-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddTemplate2(groupId int64, classNameId int64, classPK int64, templateKey string, nameMap map[string]interface{}, descriptionMap map[string]interface{}, type string, mode string, language string, script string, cacheable bool, smallImage bool, smallImageURL string, smallImageFile io.Reader, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["templateKey"] = templateKey
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["type"] = type
	_params["mode"] = mode
	_params["language"] = language
	_params["script"] = script
	_params["cacheable"] = cacheable
	_params["smallImage"] = smallImage
	_params["smallImageURL"] = smallImageURL
	_params["smallImageFile"] = smallImageFile
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmtemplate/add-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CopyTemplate(templateId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["templateId"] = templateId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmtemplate/copy-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CopyTemplate2(templateId int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["templateId"] = templateId
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmtemplate/copy-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CopyTemplates(classNameId int64, classPK int64, newClassPK int64, type string, serviceContext *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["newClassPK"] = newClassPK
	_params["type"] = type
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmtemplate/copy-templates": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) DeleteTemplate(templateId int64) error {
	_params := make(map[string]interface{})

	_params["templateId"] = templateId

	_cmd := map[string]interface{}{
		"/ddmtemplate/delete-template": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) FetchTemplate(groupId int64, classNameId int64, templateKey string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["templateKey"] = templateKey

	_cmd := map[string]interface{}{
		"/ddmtemplate/fetch-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetTemplate(templateId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["templateId"] = templateId

	_cmd := map[string]interface{}{
		"/ddmtemplate/get-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetTemplate2(groupId int64, classNameId int64, templateKey string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["templateKey"] = templateKey

	_cmd := map[string]interface{}{
		"/ddmtemplate/get-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetTemplate3(groupId int64, classNameId int64, templateKey string, includeGlobalTemplates bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["templateKey"] = templateKey
	_params["includeGlobalTemplates"] = includeGlobalTemplates

	_cmd := map[string]interface{}{
		"/ddmtemplate/get-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetTemplates(groupId int64, classNameId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId

	_cmd := map[string]interface{}{
		"/ddmtemplate/get-templates": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTemplates2(groupId int64, classNameId int64, classPK int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/ddmtemplate/get-templates": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTemplates3(groupId int64, classNameId int64, classPK int64, type string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["type"] = type

	_cmd := map[string]interface{}{
		"/ddmtemplate/get-templates": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTemplates4(groupId int64, classNameId int64, classPK int64, type string, mode string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["type"] = type
	_params["mode"] = mode

	_cmd := map[string]interface{}{
		"/ddmtemplate/get-templates": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTemplatesByClassPk(groupId int64, classPK int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/ddmtemplate/get-templates-by-class-pk": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTemplatesByStructureClassNameId(groupId int64, structureClassNameId int64, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["structureClassNameId"] = structureClassNameId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/ddmtemplate/get-templates-by-structure-class-name-id": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTemplatesByStructureClassNameIdCount(groupId int64, structureClassNameId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["structureClassNameId"] = structureClassNameId

	_cmd := map[string]interface{}{
		"/ddmtemplate/get-templates-by-structure-class-name-id-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) Search(companyId int64, groupId int64, classNameId int64, classPK int64, keywords string, type string, mode string, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["keywords"] = keywords
	_params["type"] = type
	_params["mode"] = mode
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/ddmtemplate/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search2(companyId int64, groupIds []interface{}, classNameIds []interface{}, classPKs []interface{}, keywords string, type string, mode string, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["classNameIds"] = classNameIds
	_params["classPKs"] = classPKs
	_params["keywords"] = keywords
	_params["type"] = type
	_params["mode"] = mode
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/ddmtemplate/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search3(companyId int64, groupId int64, classNameId int64, classPK int64, name string, description string, type string, mode string, language string, andOperator bool, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["name"] = name
	_params["description"] = description
	_params["type"] = type
	_params["mode"] = mode
	_params["language"] = language
	_params["andOperator"] = andOperator
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/ddmtemplate/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search4(companyId int64, groupIds []interface{}, classNameIds []interface{}, classPKs []interface{}, name string, description string, type string, mode string, language string, andOperator bool, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["classNameIds"] = classNameIds
	_params["classPKs"] = classPKs
	_params["name"] = name
	_params["description"] = description
	_params["type"] = type
	_params["mode"] = mode
	_params["language"] = language
	_params["andOperator"] = andOperator
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/ddmtemplate/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) SearchCount(companyId int64, groupId int64, classNameId int64, classPK int64, name string, description string, type string, mode string, language string, andOperator bool) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["name"] = name
	_params["description"] = description
	_params["type"] = type
	_params["mode"] = mode
	_params["language"] = language
	_params["andOperator"] = andOperator

	_cmd := map[string]interface{}{
		"/ddmtemplate/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) SearchCount2(companyId int64, groupIds []interface{}, classNameIds []interface{}, classPKs []interface{}, name string, description string, type string, mode string, language string, andOperator bool) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["classNameIds"] = classNameIds
	_params["classPKs"] = classPKs
	_params["name"] = name
	_params["description"] = description
	_params["type"] = type
	_params["mode"] = mode
	_params["language"] = language
	_params["andOperator"] = andOperator

	_cmd := map[string]interface{}{
		"/ddmtemplate/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) SearchCount3(companyId int64, groupId int64, classNameId int64, classPK int64, keywords string, type string, mode string) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["keywords"] = keywords
	_params["type"] = type
	_params["mode"] = mode

	_cmd := map[string]interface{}{
		"/ddmtemplate/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) SearchCount4(companyId int64, groupIds []interface{}, classNameIds []interface{}, classPKs []interface{}, keywords string, type string, mode string) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupIds"] = groupIds
	_params["classNameIds"] = classNameIds
	_params["classPKs"] = classPKs
	_params["keywords"] = keywords
	_params["type"] = type
	_params["mode"] = mode

	_cmd := map[string]interface{}{
		"/ddmtemplate/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) UpdateTemplate(templateId int64, classPK int64, nameMap map[string]interface{}, descriptionMap map[string]interface{}, type string, mode string, language string, script string, cacheable bool, smallImage bool, smallImageURL string, smallImageFile io.Reader, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["templateId"] = templateId
	_params["classPK"] = classPK
	_params["nameMap"] = nameMap
	_params["descriptionMap"] = descriptionMap
	_params["type"] = type
	_params["mode"] = mode
	_params["language"] = language
	_params["script"] = script
	_params["cacheable"] = cacheable
	_params["smallImage"] = smallImage
	_params["smallImageURL"] = smallImageURL
	_params["smallImageFile"] = smallImageFile
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/ddmtemplate/update-template": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

