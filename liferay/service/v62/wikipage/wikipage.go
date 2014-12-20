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

package wikipage

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddPage(nodeId int64, title string, content string, summary string, minorEdit bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["content"] = content
	_params["summary"] = summary
	_params["minorEdit"] = minorEdit
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/wikipage/add-page": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddPage2(nodeId int64, title string, content string, summary string, minorEdit bool, format string, parentTitle string, redirectTitle string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["content"] = content
	_params["summary"] = summary
	_params["minorEdit"] = minorEdit
	_params["format"] = format
	_params["parentTitle"] = parentTitle
	_params["redirectTitle"] = redirectTitle
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/wikipage/add-page": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddPageAttachment(nodeId int64, title string, fileName string, file io.Reader, mimeType string) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["fileName"] = fileName
	_params["file"] = file
	_params["mimeType"] = mimeType

	_cmd := map[string]interface{}{
		"/wikipage/add-page-attachment": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) AddPageAttachments(nodeId int64, title string, inputStreamOVPs []interface{}) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["inputStreamOVPs"] = inputStreamOVPs

	_cmd := map[string]interface{}{
		"/wikipage/add-page-attachments": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) ChangeParent(nodeId int64, title string, newParentTitle string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["newParentTitle"] = newParentTitle
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/wikipage/change-parent": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) CopyPageAttachments(templateNodeId int64, templateTitle string, nodeId int64, title string) error {
	_params := make(map[string]interface{})

	_params["templateNodeId"] = templateNodeId
	_params["templateTitle"] = templateTitle
	_params["nodeId"] = nodeId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/wikipage/copy-page-attachments": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeletePage(nodeId int64, title string) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/wikipage/delete-page": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeletePage2(nodeId int64, title string, version float64) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["version"] = version

	_cmd := map[string]interface{}{
		"/wikipage/delete-page": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeletePageAttachment(nodeId int64, title string, fileName string) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["fileName"] = fileName

	_cmd := map[string]interface{}{
		"/wikipage/delete-page-attachment": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeletePageAttachments(nodeId int64, title string) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/wikipage/delete-page-attachments": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteTempPageAttachment(nodeId int64, fileName string, tempFolderName string) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["fileName"] = fileName
	_params["tempFolderName"] = tempFolderName

	_cmd := map[string]interface{}{
		"/wikipage/delete-temp-page-attachment": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteTrashPageAttachments(nodeId int64, title string) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/wikipage/delete-trash-page-attachments": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DiscardDraft(nodeId int64, title string, version float64) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["version"] = version

	_cmd := map[string]interface{}{
		"/wikipage/discard-draft": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetChildren(groupId int64, nodeId int64, head bool, parentTitle string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["nodeId"] = nodeId
	_params["head"] = head
	_params["parentTitle"] = parentTitle

	_cmd := map[string]interface{}{
		"/wikipage/get-children": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetDraftPage(nodeId int64, title string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/wikipage/get-draft-page": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetNodePages(nodeId int64, max int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["max"] = max

	_cmd := map[string]interface{}{
		"/wikipage/get-node-pages": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetNodePagesRss(nodeId int64, max int, type string, version float64, displayStyle string, feedURL string, entryURL string) (string, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["max"] = max
	_params["type"] = type
	_params["version"] = version
	_params["displayStyle"] = displayStyle
	_params["feedURL"] = feedURL
	_params["entryURL"] = entryURL

	_cmd := map[string]interface{}{
		"/wikipage/get-node-pages-rss": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetNodePagesRss2(nodeId int64, max int, type string, version float64, displayStyle string, feedURL string, entryURL string, attachmentURLPrefix string) (string, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["max"] = max
	_params["type"] = type
	_params["version"] = version
	_params["displayStyle"] = displayStyle
	_params["feedURL"] = feedURL
	_params["entryURL"] = entryURL
	_params["attachmentURLPrefix"] = attachmentURLPrefix

	_cmd := map[string]interface{}{
		"/wikipage/get-node-pages-rss": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetOrphans(groupId int64, nodeId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["nodeId"] = nodeId

	_cmd := map[string]interface{}{
		"/wikipage/get-orphans": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetPage(nodeId int64, title string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/wikipage/get-page": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetPage2(groupId int64, nodeId int64, title string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["nodeId"] = nodeId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/wikipage/get-page": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetPage3(nodeId int64, title string, head *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	liferay.MangleObjectWrapper(_params, "head", "java.lang.Boolean", head)

	_cmd := map[string]interface{}{
		"/wikipage/get-page": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetPage4(nodeId int64, title string, version float64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["version"] = version

	_cmd := map[string]interface{}{
		"/wikipage/get-page": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetPages(groupId int64, userId int64, nodeId int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["nodeId"] = nodeId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/wikipage/get-pages": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetPages2(groupId int64, nodeId int64, head bool, status int, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["nodeId"] = nodeId
	_params["head"] = head
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/wikipage/get-pages": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetPagesCount(groupId int64, nodeId int64, head bool) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["nodeId"] = nodeId
	_params["head"] = head

	_cmd := map[string]interface{}{
		"/wikipage/get-pages-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetPagesCount2(groupId int64, userId int64, nodeId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["nodeId"] = nodeId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/wikipage/get-pages-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetPagesRss(companyId int64, nodeId int64, title string, max int, type string, version float64, displayStyle string, feedURL string, entryURL string, locale string) (string, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["max"] = max
	_params["type"] = type
	_params["version"] = version
	_params["displayStyle"] = displayStyle
	_params["feedURL"] = feedURL
	_params["entryURL"] = entryURL
	_params["locale"] = locale

	_cmd := map[string]interface{}{
		"/wikipage/get-pages-rss": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetPagesRss2(companyId int64, nodeId int64, title string, max int, type string, version float64, displayStyle string, feedURL string, entryURL string, attachmentURLPrefix string, locale string) (string, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["max"] = max
	_params["type"] = type
	_params["version"] = version
	_params["displayStyle"] = displayStyle
	_params["feedURL"] = feedURL
	_params["entryURL"] = entryURL
	_params["attachmentURLPrefix"] = attachmentURLPrefix
	_params["locale"] = locale

	_cmd := map[string]interface{}{
		"/wikipage/get-pages-rss": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetRecentChanges(groupId int64, nodeId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["nodeId"] = nodeId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/wikipage/get-recent-changes": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetRecentChangesCount(groupId int64, nodeId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["nodeId"] = nodeId

	_cmd := map[string]interface{}{
		"/wikipage/get-recent-changes-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64)
	}

	return v, err
}

func (s *Service) GetTempPageAttachmentNames(nodeId int64, tempFolderName string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["tempFolderName"] = tempFolderName

	_cmd := map[string]interface{}{
		"/wikipage/get-temp-page-attachment-names": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) MovePage(nodeId int64, title string, newTitle string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["newTitle"] = newTitle
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/wikipage/move-page": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) MovePageAttachmentToTrash(nodeId int64, title string, fileName string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["fileName"] = fileName

	_cmd := map[string]interface{}{
		"/wikipage/move-page-attachment-to-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MovePageToTrash(nodeId int64, title string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/wikipage/move-page-to-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MovePageToTrash2(nodeId int64, title string, version float64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["version"] = version

	_cmd := map[string]interface{}{
		"/wikipage/move-page-to-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RestorePageAttachmentFromTrash(nodeId int64, title string, fileName string) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["fileName"] = fileName

	_cmd := map[string]interface{}{
		"/wikipage/restore-page-attachment-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RestorePageFromTrash(resourcePrimKey int64) error {
	_params := make(map[string]interface{})

	_params["resourcePrimKey"] = resourcePrimKey

	_cmd := map[string]interface{}{
		"/wikipage/restore-page-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RevertPage(nodeId int64, title string, version float64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["version"] = version
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/wikipage/revert-page": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) SubscribePage(nodeId int64, title string) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/wikipage/subscribe-page": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsubscribePage(nodeId int64, title string) error {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title

	_cmd := map[string]interface{}{
		"/wikipage/unsubscribe-page": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdatePage(nodeId int64, title string, version float64, content string, summary string, minorEdit bool, format string, parentTitle string, redirectTitle string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["nodeId"] = nodeId
	_params["title"] = title
	_params["version"] = version
	_params["content"] = content
	_params["summary"] = summary
	_params["minorEdit"] = minorEdit
	_params["format"] = format
	_params["parentTitle"] = parentTitle
	_params["redirectTitle"] = redirectTitle
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/wikipage/update-page": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

