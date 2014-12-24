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

package layout

import (
	"io"

	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddLayout(groupId int64, privateLayout bool, parentLayoutId int64, name string, title string, description string, _type string, hidden bool, friendlyURL string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["parentLayoutId"] = parentLayoutId
	_params["name"] = name
	_params["title"] = title
	_params["description"] = description
	_params["type"] = _type
	_params["hidden"] = hidden
	_params["friendlyURL"] = friendlyURL
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layout/add-layout": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddLayout2(groupId int64, privateLayout bool, parentLayoutId int64, localeNamesMap map[string]interface{}, localeTitlesMap map[string]interface{}, descriptionMap map[string]interface{}, keywordsMap map[string]interface{}, robotsMap map[string]interface{}, _type string, hidden bool, friendlyURL string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["parentLayoutId"] = parentLayoutId
	_params["localeNamesMap"] = localeNamesMap
	_params["localeTitlesMap"] = localeTitlesMap
	_params["descriptionMap"] = descriptionMap
	_params["keywordsMap"] = keywordsMap
	_params["robotsMap"] = robotsMap
	_params["type"] = _type
	_params["hidden"] = hidden
	_params["friendlyURL"] = friendlyURL
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layout/add-layout": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddLayout3(groupId int64, privateLayout bool, parentLayoutId int64, localeNamesMap map[string]interface{}, localeTitlesMap map[string]interface{}, descriptionMap map[string]interface{}, keywordsMap map[string]interface{}, robotsMap map[string]interface{}, _type string, typeSettings string, hidden bool, friendlyURLMap map[string]interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["parentLayoutId"] = parentLayoutId
	_params["localeNamesMap"] = localeNamesMap
	_params["localeTitlesMap"] = localeTitlesMap
	_params["descriptionMap"] = descriptionMap
	_params["keywordsMap"] = keywordsMap
	_params["robotsMap"] = robotsMap
	_params["type"] = _type
	_params["typeSettings"] = typeSettings
	_params["hidden"] = hidden
	_params["friendlyURLMap"] = friendlyURLMap
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layout/add-layout": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteLayout(plid int64, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["plid"] = plid
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layout/delete-layout": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteLayout2(groupId int64, privateLayout bool, layoutId int64, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutId"] = layoutId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layout/delete-layout": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteTempFileEntry(groupId int64, fileName string, tempFolderName string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["fileName"] = fileName
	_params["tempFolderName"] = tempFolderName

	_cmd := map[string]interface{}{
		"/layout/delete-temp-file-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) ExportLayouts(groupId int64, privateLayout bool, parameterMap map[string]interface{}, startDate int64, endDate int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["parameterMap"] = parameterMap
	_params["startDate"] = startDate
	_params["endDate"] = endDate

	_cmd := map[string]interface{}{
		"/layout/export-layouts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) ExportLayouts2(groupId int64, privateLayout bool, layoutIds []interface{}, parameterMap map[string]interface{}, startDate int64, endDate int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutIds"] = layoutIds
	_params["parameterMap"] = parameterMap
	_params["startDate"] = startDate
	_params["endDate"] = endDate

	_cmd := map[string]interface{}{
		"/layout/export-layouts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) ExportLayoutsAsFile(groupId int64, privateLayout bool, layoutIds []interface{}, parameterMap map[string]interface{}, startDate int64, endDate int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutIds"] = layoutIds
	_params["parameterMap"] = parameterMap
	_params["startDate"] = startDate
	_params["endDate"] = endDate

	_cmd := map[string]interface{}{
		"/layout/export-layouts-as-file": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) ExportLayoutsAsFileInBackground(taskName string, groupId int64, privateLayout bool, layoutIds []interface{}, parameterMap map[string]interface{}, startDate int64, endDate int64, fileName string) (int64, error) {
	_params := make(map[string]interface{})

	_params["taskName"] = taskName
	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutIds"] = layoutIds
	_params["parameterMap"] = parameterMap
	_params["startDate"] = startDate
	_params["endDate"] = endDate
	_params["fileName"] = fileName

	_cmd := map[string]interface{}{
		"/layout/export-layouts-as-file-in-background": _params,
	}

	var v int64

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64))
	}

	return v, err
}

func (s *Service) ExportPortletInfo(companyId int64, portletId string, parameterMap map[string]interface{}, startDate int64, endDate int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["portletId"] = portletId
	_params["parameterMap"] = parameterMap
	_params["startDate"] = startDate
	_params["endDate"] = endDate

	_cmd := map[string]interface{}{
		"/layout/export-portlet-info": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) ExportPortletInfo2(plid int64, groupId int64, portletId string, parameterMap map[string]interface{}, startDate int64, endDate int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["plid"] = plid
	_params["groupId"] = groupId
	_params["portletId"] = portletId
	_params["parameterMap"] = parameterMap
	_params["startDate"] = startDate
	_params["endDate"] = endDate

	_cmd := map[string]interface{}{
		"/layout/export-portlet-info": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) ExportPortletInfoAsFile(portletId string, parameterMap map[string]interface{}, startDate int64, endDate int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["portletId"] = portletId
	_params["parameterMap"] = parameterMap
	_params["startDate"] = startDate
	_params["endDate"] = endDate

	_cmd := map[string]interface{}{
		"/layout/export-portlet-info-as-file": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) ExportPortletInfoAsFile2(plid int64, groupId int64, portletId string, parameterMap map[string]interface{}, startDate int64, endDate int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["plid"] = plid
	_params["groupId"] = groupId
	_params["portletId"] = portletId
	_params["parameterMap"] = parameterMap
	_params["startDate"] = startDate
	_params["endDate"] = endDate

	_cmd := map[string]interface{}{
		"/layout/export-portlet-info-as-file": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) ExportPortletInfoAsFileInBackground(taskName string, portletId string, parameterMap map[string]interface{}, startDate int64, endDate int64, fileName string) (int64, error) {
	_params := make(map[string]interface{})

	_params["taskName"] = taskName
	_params["portletId"] = portletId
	_params["parameterMap"] = parameterMap
	_params["startDate"] = startDate
	_params["endDate"] = endDate
	_params["fileName"] = fileName

	_cmd := map[string]interface{}{
		"/layout/export-portlet-info-as-file-in-background": _params,
	}

	var v int64

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64))
	}

	return v, err
}

func (s *Service) ExportPortletInfoAsFileInBackground2(taskName string, plid int64, groupId int64, portletId string, parameterMap map[string]interface{}, startDate int64, endDate int64, fileName string) (int64, error) {
	_params := make(map[string]interface{})

	_params["taskName"] = taskName
	_params["plid"] = plid
	_params["groupId"] = groupId
	_params["portletId"] = portletId
	_params["parameterMap"] = parameterMap
	_params["startDate"] = startDate
	_params["endDate"] = endDate
	_params["fileName"] = fileName

	_cmd := map[string]interface{}{
		"/layout/export-portlet-info-as-file-in-background": _params,
	}

	var v int64

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64))
	}

	return v, err
}

func (s *Service) GetAncestorLayouts(plid int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["plid"] = plid

	_cmd := map[string]interface{}{
		"/layout/get-ancestor-layouts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetDefaultPlid(groupId int64, scopeGroupId int64, portletId string) (int64, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["scopeGroupId"] = scopeGroupId
	_params["portletId"] = portletId

	_cmd := map[string]interface{}{
		"/layout/get-default-plid": _params,
	}

	var v int64

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64))
	}

	return v, err
}

func (s *Service) GetDefaultPlid2(groupId int64, scopeGroupId int64, privateLayout bool, portletId string) (int64, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["scopeGroupId"] = scopeGroupId
	_params["privateLayout"] = privateLayout
	_params["portletId"] = portletId

	_cmd := map[string]interface{}{
		"/layout/get-default-plid": _params,
	}

	var v int64

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64))
	}

	return v, err
}

func (s *Service) GetLayoutByUuidAndGroupId(uuid string, groupId int64, privateLayout bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["uuid"] = uuid
	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout

	_cmd := map[string]interface{}{
		"/layout/get-layout-by-uuid-and-group-id": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetLayoutName(groupId int64, privateLayout bool, layoutId int64, languageId string) (string, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutId"] = layoutId
	_params["languageId"] = languageId

	_cmd := map[string]interface{}{
		"/layout/get-layout-name": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetLayoutReferences(companyId int64, portletId string, preferencesKey string, preferencesValue string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["portletId"] = portletId
	_params["preferencesKey"] = preferencesKey
	_params["preferencesValue"] = preferencesValue

	_cmd := map[string]interface{}{
		"/layout/get-layout-references": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetLayouts(groupId int64, privateLayout bool) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout

	_cmd := map[string]interface{}{
		"/layout/get-layouts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetLayouts2(groupId int64, privateLayout bool, parentLayoutId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["parentLayoutId"] = parentLayoutId

	_cmd := map[string]interface{}{
		"/layout/get-layouts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetLayouts3(groupId int64, privateLayout bool, parentLayoutId int64, incomplete bool, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["parentLayoutId"] = parentLayoutId
	_params["incomplete"] = incomplete
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/layout/get-layouts": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetTempFileEntryNames(groupId int64, tempFolderName string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["tempFolderName"] = tempFolderName

	_cmd := map[string]interface{}{
		"/layout/get-temp-file-entry-names": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) ImportLayouts(groupId int64, privateLayout bool, parameterMap map[string]interface{}, bytes []byte) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["parameterMap"] = parameterMap
	_params["bytes"] = liferay.ToJSONString(bytes)

	_cmd := map[string]interface{}{
		"/layout/import-layouts": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) ImportLayouts2(groupId int64, privateLayout bool, parameterMap map[string]interface{}, file io.Reader) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["parameterMap"] = parameterMap
	_params["file"] = file

	_cmd := map[string]interface{}{
		"/layout/import-layouts": _params,
	}

	_, err := s.session.Upload(_cmd)

	return err
}

func (s *Service) ImportLayoutsInBackground(taskName string, groupId int64, privateLayout bool, parameterMap map[string]interface{}, file io.Reader) (int64, error) {
	_params := make(map[string]interface{})

	_params["taskName"] = taskName
	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["parameterMap"] = parameterMap
	_params["file"] = file

	_cmd := map[string]interface{}{
		"/layout/import-layouts-in-background": _params,
	}

	var v int64

	res, err := s.session.Upload(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64))
	}

	return v, err
}

func (s *Service) ImportPortletInfo(portletId string, parameterMap map[string]interface{}, file io.Reader) error {
	_params := make(map[string]interface{})

	_params["portletId"] = portletId
	_params["parameterMap"] = parameterMap
	_params["file"] = file

	_cmd := map[string]interface{}{
		"/layout/import-portlet-info": _params,
	}

	_, err := s.session.Upload(_cmd)

	return err
}

func (s *Service) ImportPortletInfo2(plid int64, groupId int64, portletId string, parameterMap map[string]interface{}, file io.Reader) error {
	_params := make(map[string]interface{})

	_params["plid"] = plid
	_params["groupId"] = groupId
	_params["portletId"] = portletId
	_params["parameterMap"] = parameterMap
	_params["file"] = file

	_cmd := map[string]interface{}{
		"/layout/import-portlet-info": _params,
	}

	_, err := s.session.Upload(_cmd)

	return err
}

func (s *Service) ImportPortletInfoInBackground(taskName string, portletId string, parameterMap map[string]interface{}, file io.Reader) error {
	_params := make(map[string]interface{})

	_params["taskName"] = taskName
	_params["portletId"] = portletId
	_params["parameterMap"] = parameterMap
	_params["file"] = file

	_cmd := map[string]interface{}{
		"/layout/import-portlet-info-in-background": _params,
	}

	_, err := s.session.Upload(_cmd)

	return err
}

func (s *Service) ImportPortletInfoInBackground2(taskName string, plid int64, groupId int64, portletId string, parameterMap map[string]interface{}, file io.Reader) (int64, error) {
	_params := make(map[string]interface{})

	_params["taskName"] = taskName
	_params["plid"] = plid
	_params["groupId"] = groupId
	_params["portletId"] = portletId
	_params["parameterMap"] = parameterMap
	_params["file"] = file

	_cmd := map[string]interface{}{
		"/layout/import-portlet-info-in-background": _params,
	}

	var v int64

	res, err := s.session.Upload(_cmd)

	if err == nil && res != nil {
		v = int64(res.(float64))
	}

	return v, err
}

func (s *Service) SchedulePublishToLive(sourceGroupId int64, targetGroupId int64, privateLayout bool, layoutIdMap map[string]interface{}, parameterMap map[string]interface{}, scope string, startDate int64, endDate int64, groupName string, cronText string, schedulerStartDate int64, schedulerEndDate int64, description string) error {
	_params := make(map[string]interface{})

	_params["sourceGroupId"] = sourceGroupId
	_params["targetGroupId"] = targetGroupId
	_params["privateLayout"] = privateLayout
	_params["layoutIdMap"] = layoutIdMap
	_params["parameterMap"] = parameterMap
	_params["scope"] = scope
	_params["startDate"] = startDate
	_params["endDate"] = endDate
	_params["groupName"] = groupName
	_params["cronText"] = cronText
	_params["schedulerStartDate"] = schedulerStartDate
	_params["schedulerEndDate"] = schedulerEndDate
	_params["description"] = description

	_cmd := map[string]interface{}{
		"/layout/schedule-publish-to-live": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SchedulePublishToRemote(sourceGroupId int64, privateLayout bool, layoutIdMap map[string]interface{}, parameterMap map[string]interface{}, remoteAddress string, remotePort int, remotePathContext string, secureConnection bool, remoteGroupId int64, remotePrivateLayout bool, startDate int64, endDate int64, groupName string, cronText string, schedulerStartDate int64, schedulerEndDate int64, description string) error {
	_params := make(map[string]interface{})

	_params["sourceGroupId"] = sourceGroupId
	_params["privateLayout"] = privateLayout
	_params["layoutIdMap"] = layoutIdMap
	_params["parameterMap"] = parameterMap
	_params["remoteAddress"] = remoteAddress
	_params["remotePort"] = remotePort
	_params["remotePathContext"] = remotePathContext
	_params["secureConnection"] = secureConnection
	_params["remoteGroupId"] = remoteGroupId
	_params["remotePrivateLayout"] = remotePrivateLayout
	_params["startDate"] = startDate
	_params["endDate"] = endDate
	_params["groupName"] = groupName
	_params["cronText"] = cronText
	_params["schedulerStartDate"] = schedulerStartDate
	_params["schedulerEndDate"] = schedulerEndDate
	_params["description"] = description

	_cmd := map[string]interface{}{
		"/layout/schedule-publish-to-remote": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SetLayouts(groupId int64, privateLayout bool, parentLayoutId int64, layoutIds []interface{}, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["parentLayoutId"] = parentLayoutId
	_params["layoutIds"] = layoutIds
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layout/set-layouts": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnschedulePublishToLive(groupId int64, jobName string, groupName string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["jobName"] = jobName
	_params["groupName"] = groupName

	_cmd := map[string]interface{}{
		"/layout/unschedule-publish-to-live": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnschedulePublishToRemote(groupId int64, jobName string, groupName string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["jobName"] = jobName
	_params["groupName"] = groupName

	_cmd := map[string]interface{}{
		"/layout/unschedule-publish-to-remote": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateLayout(groupId int64, privateLayout bool, layoutId int64, parentLayoutId int64, localeNamesMap map[string]interface{}, localeTitlesMap map[string]interface{}, descriptionMap map[string]interface{}, keywordsMap map[string]interface{}, robotsMap map[string]interface{}, _type string, hidden bool, friendlyURL string, iconImage *liferay.ObjectWrapper, iconBytes []byte, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutId"] = layoutId
	_params["parentLayoutId"] = parentLayoutId
	_params["localeNamesMap"] = localeNamesMap
	_params["localeTitlesMap"] = localeTitlesMap
	_params["descriptionMap"] = descriptionMap
	_params["keywordsMap"] = keywordsMap
	_params["robotsMap"] = robotsMap
	_params["type"] = _type
	_params["hidden"] = hidden
	_params["friendlyURL"] = friendlyURL
	liferay.MangleObjectWrapper(_params, "iconImage", "java.lang.Boolean", iconImage)
	_params["iconBytes"] = liferay.ToJSONString(iconBytes)
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layout/update-layout": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateLayout2(groupId int64, privateLayout bool, layoutId int64, parentLayoutId int64, localeNamesMap map[string]interface{}, localeTitlesMap map[string]interface{}, descriptionMap map[string]interface{}, keywordsMap map[string]interface{}, robotsMap map[string]interface{}, _type string, hidden bool, friendlyURLMap map[string]interface{}, iconImage *liferay.ObjectWrapper, iconBytes []byte, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutId"] = layoutId
	_params["parentLayoutId"] = parentLayoutId
	_params["localeNamesMap"] = localeNamesMap
	_params["localeTitlesMap"] = localeTitlesMap
	_params["descriptionMap"] = descriptionMap
	_params["keywordsMap"] = keywordsMap
	_params["robotsMap"] = robotsMap
	_params["type"] = _type
	_params["hidden"] = hidden
	_params["friendlyURLMap"] = friendlyURLMap
	liferay.MangleObjectWrapper(_params, "iconImage", "java.lang.Boolean", iconImage)
	_params["iconBytes"] = liferay.ToJSONString(iconBytes)
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/layout/update-layout": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateLayout3(groupId int64, privateLayout bool, layoutId int64, typeSettings string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutId"] = layoutId
	_params["typeSettings"] = typeSettings

	_cmd := map[string]interface{}{
		"/layout/update-layout": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateLookAndFeel(groupId int64, privateLayout bool, layoutId int64, themeId string, colorSchemeId string, css string, wapTheme bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutId"] = layoutId
	_params["themeId"] = themeId
	_params["colorSchemeId"] = colorSchemeId
	_params["css"] = css
	_params["wapTheme"] = wapTheme

	_cmd := map[string]interface{}{
		"/layout/update-look-and-feel": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateName(plid int64, name string, languageId string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["plid"] = plid
	_params["name"] = name
	_params["languageId"] = languageId

	_cmd := map[string]interface{}{
		"/layout/update-name": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateName2(groupId int64, privateLayout bool, layoutId int64, name string, languageId string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutId"] = layoutId
	_params["name"] = name
	_params["languageId"] = languageId

	_cmd := map[string]interface{}{
		"/layout/update-name": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateParentLayoutId(plid int64, parentPlid int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["plid"] = plid
	_params["parentPlid"] = parentPlid

	_cmd := map[string]interface{}{
		"/layout/update-parent-layout-id": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateParentLayoutId2(groupId int64, privateLayout bool, layoutId int64, parentLayoutId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutId"] = layoutId
	_params["parentLayoutId"] = parentLayoutId

	_cmd := map[string]interface{}{
		"/layout/update-parent-layout-id": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdatePriority(plid int64, priority int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["plid"] = plid
	_params["priority"] = priority

	_cmd := map[string]interface{}{
		"/layout/update-priority": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdatePriority2(groupId int64, privateLayout bool, layoutId int64, priority int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutId"] = layoutId
	_params["priority"] = priority

	_cmd := map[string]interface{}{
		"/layout/update-priority": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdatePriority3(groupId int64, privateLayout bool, layoutId int64, nextLayoutId int64, previousLayoutId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["layoutId"] = layoutId
	_params["nextLayoutId"] = nextLayoutId
	_params["previousLayoutId"] = previousLayoutId

	_cmd := map[string]interface{}{
		"/layout/update-priority": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) ValidateImportLayoutsFile(groupId int64, privateLayout bool, parameterMap map[string]interface{}, file io.Reader) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["privateLayout"] = privateLayout
	_params["parameterMap"] = parameterMap
	_params["file"] = file

	_cmd := map[string]interface{}{
		"/layout/validate-import-layouts-file": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Upload(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) ValidateImportPortletInfo(plid int64, groupId int64, portletId string, parameterMap map[string]interface{}, file io.Reader) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["plid"] = plid
	_params["groupId"] = groupId
	_params["portletId"] = portletId
	_params["parameterMap"] = parameterMap
	_params["file"] = file

	_cmd := map[string]interface{}{
		"/layout/validate-import-portlet-info": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Upload(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

