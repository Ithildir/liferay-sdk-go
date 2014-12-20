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

package mbmessage

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddDiscussionMessage(groupId int64, className string, classPK int64, permissionClassName string, permissionClassPK int64, permissionOwnerId int64, threadId int64, parentMessageId int64, subject string, body string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className
	_params["classPK"] = classPK
	_params["permissionClassName"] = permissionClassName
	_params["permissionClassPK"] = permissionClassPK
	_params["permissionOwnerId"] = permissionOwnerId
	_params["threadId"] = threadId
	_params["parentMessageId"] = parentMessageId
	_params["subject"] = subject
	_params["body"] = body
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mbmessage/add-discussion-message": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddMessage(groupId int64, categoryId int64, subject string, body string, format string, inputStreamOVPs []interface{}, anonymous bool, priority float64, allowPingbacks bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["subject"] = subject
	_params["body"] = body
	_params["format"] = format
	_params["inputStreamOVPs"] = inputStreamOVPs
	_params["anonymous"] = anonymous
	_params["priority"] = priority
	_params["allowPingbacks"] = allowPingbacks
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mbmessage/add-message": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddMessage2(groupId int64, categoryId int64, threadId int64, parentMessageId int64, subject string, body string, format string, inputStreamOVPs []interface{}, anonymous bool, priority float64, allowPingbacks bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["threadId"] = threadId
	_params["parentMessageId"] = parentMessageId
	_params["subject"] = subject
	_params["body"] = body
	_params["format"] = format
	_params["inputStreamOVPs"] = inputStreamOVPs
	_params["anonymous"] = anonymous
	_params["priority"] = priority
	_params["allowPingbacks"] = allowPingbacks
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mbmessage/add-message": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddMessage3(categoryId int64, subject string, body string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["categoryId"] = categoryId
	_params["subject"] = subject
	_params["body"] = body
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mbmessage/add-message": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddMessage4(parentMessageId int64, subject string, body string, format string, inputStreamOVPs []interface{}, anonymous bool, priority float64, allowPingbacks bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["parentMessageId"] = parentMessageId
	_params["subject"] = subject
	_params["body"] = body
	_params["format"] = format
	_params["inputStreamOVPs"] = inputStreamOVPs
	_params["anonymous"] = anonymous
	_params["priority"] = priority
	_params["allowPingbacks"] = allowPingbacks
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mbmessage/add-message": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteDiscussionMessage(groupId int64, className string, classPK int64, permissionClassName string, permissionClassPK int64, permissionOwnerId int64, messageId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className
	_params["classPK"] = classPK
	_params["permissionClassName"] = permissionClassName
	_params["permissionClassPK"] = permissionClassPK
	_params["permissionOwnerId"] = permissionOwnerId
	_params["messageId"] = messageId

	_cmd := map[string]interface{}{
		"/mbmessage/delete-discussion-message": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteMessage(messageId int64) error {
	_params := make(map[string]interface{})

	_params["messageId"] = messageId

	_cmd := map[string]interface{}{
		"/mbmessage/delete-message": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteMessageAttachments(messageId int64) error {
	_params := make(map[string]interface{})

	_params["messageId"] = messageId

	_cmd := map[string]interface{}{
		"/mbmessage/delete-message-attachments": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetCategoryMessages(groupId int64, categoryId int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbmessage/get-category-messages": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetCategoryMessagesCount(groupId int64, categoryId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/mbmessage/get-category-messages-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetCategoryMessagesRss(groupId int64, categoryId int64, status int, max int, type string, version float64, displayStyle string, feedURL string, entryURL string, themeDisplay *liferay.ObjectWrapper) (string, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["status"] = status
	_params["max"] = max
	_params["type"] = type
	_params["version"] = version
	_params["displayStyle"] = displayStyle
	_params["feedURL"] = feedURL
	_params["entryURL"] = entryURL
	liferay.MangleObjectWrapper(_params, "themeDisplay", "com.liferay.portal.theme.ThemeDisplay", themeDisplay)

	_cmd := map[string]interface{}{
		"/mbmessage/get-category-messages-rss": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetCompanyMessagesRss(companyId int64, status int, max int, type string, version float64, displayStyle string, feedURL string, entryURL string, themeDisplay *liferay.ObjectWrapper) (string, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["status"] = status
	_params["max"] = max
	_params["type"] = type
	_params["version"] = version
	_params["displayStyle"] = displayStyle
	_params["feedURL"] = feedURL
	_params["entryURL"] = entryURL
	liferay.MangleObjectWrapper(_params, "themeDisplay", "com.liferay.portal.theme.ThemeDisplay", themeDisplay)

	_cmd := map[string]interface{}{
		"/mbmessage/get-company-messages-rss": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetGroupMessagesCount(groupId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/mbmessage/get-group-messages-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetGroupMessagesRss(groupId int64, userId int64, status int, max int, type string, version float64, displayStyle string, feedURL string, entryURL string, themeDisplay *liferay.ObjectWrapper) (string, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["status"] = status
	_params["max"] = max
	_params["type"] = type
	_params["version"] = version
	_params["displayStyle"] = displayStyle
	_params["feedURL"] = feedURL
	_params["entryURL"] = entryURL
	liferay.MangleObjectWrapper(_params, "themeDisplay", "com.liferay.portal.theme.ThemeDisplay", themeDisplay)

	_cmd := map[string]interface{}{
		"/mbmessage/get-group-messages-rss": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetGroupMessagesRss2(groupId int64, status int, max int, type string, version float64, displayStyle string, feedURL string, entryURL string, themeDisplay *liferay.ObjectWrapper) (string, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["status"] = status
	_params["max"] = max
	_params["type"] = type
	_params["version"] = version
	_params["displayStyle"] = displayStyle
	_params["feedURL"] = feedURL
	_params["entryURL"] = entryURL
	liferay.MangleObjectWrapper(_params, "themeDisplay", "com.liferay.portal.theme.ThemeDisplay", themeDisplay)

	_cmd := map[string]interface{}{
		"/mbmessage/get-group-messages-rss": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetMessage(messageId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["messageId"] = messageId

	_cmd := map[string]interface{}{
		"/mbmessage/get-message": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetMessageDisplay(messageId int64, status int, threadView string, includePrevAndNext bool) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["messageId"] = messageId
	_params["status"] = status
	_params["threadView"] = threadView
	_params["includePrevAndNext"] = includePrevAndNext

	_cmd := map[string]interface{}{
		"/mbmessage/get-message-display": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetThreadAnswersCount(groupId int64, categoryId int64, threadId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["threadId"] = threadId

	_cmd := map[string]interface{}{
		"/mbmessage/get-thread-answers-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetThreadMessages(groupId int64, categoryId int64, threadId int64, status int, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["threadId"] = threadId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/mbmessage/get-thread-messages": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetThreadMessagesCount(groupId int64, categoryId int64, threadId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["categoryId"] = categoryId
	_params["threadId"] = threadId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/mbmessage/get-thread-messages-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetThreadMessagesRss(threadId int64, status int, max int, type string, version float64, displayStyle string, feedURL string, entryURL string, themeDisplay *liferay.ObjectWrapper) (string, error) {
	_params := make(map[string]interface{})

	_params["threadId"] = threadId
	_params["status"] = status
	_params["max"] = max
	_params["type"] = type
	_params["version"] = version
	_params["displayStyle"] = displayStyle
	_params["feedURL"] = feedURL
	_params["entryURL"] = entryURL
	liferay.MangleObjectWrapper(_params, "themeDisplay", "com.liferay.portal.theme.ThemeDisplay", themeDisplay)

	_cmd := map[string]interface{}{
		"/mbmessage/get-thread-messages-rss": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) RestoreMessageAttachmentFromTrash(messageId int64, fileName string) error {
	_params := make(map[string]interface{})

	_params["messageId"] = messageId
	_params["fileName"] = fileName

	_cmd := map[string]interface{}{
		"/mbmessage/restore-message-attachment-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) SubscribeMessage(messageId int64) error {
	_params := make(map[string]interface{})

	_params["messageId"] = messageId

	_cmd := map[string]interface{}{
		"/mbmessage/subscribe-message": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UnsubscribeMessage(messageId int64) error {
	_params := make(map[string]interface{})

	_params["messageId"] = messageId

	_cmd := map[string]interface{}{
		"/mbmessage/unsubscribe-message": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateAnswer(messageId int64, answer bool, cascade bool) error {
	_params := make(map[string]interface{})

	_params["messageId"] = messageId
	_params["answer"] = answer
	_params["cascade"] = cascade

	_cmd := map[string]interface{}{
		"/mbmessage/update-answer": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateDiscussionMessage(className string, classPK int64, permissionClassName string, permissionClassPK int64, permissionOwnerId int64, messageId int64, subject string, body string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["className"] = className
	_params["classPK"] = classPK
	_params["permissionClassName"] = permissionClassName
	_params["permissionClassPK"] = permissionClassPK
	_params["permissionOwnerId"] = permissionOwnerId
	_params["messageId"] = messageId
	_params["subject"] = subject
	_params["body"] = body
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mbmessage/update-discussion-message": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateMessage(messageId int64, subject string, body string, inputStreamOVPs []interface{}, existingFiles []interface{}, priority float64, allowPingbacks bool, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["messageId"] = messageId
	_params["subject"] = subject
	_params["body"] = body
	_params["inputStreamOVPs"] = inputStreamOVPs
	_params["existingFiles"] = existingFiles
	_params["priority"] = priority
	_params["allowPingbacks"] = allowPingbacks
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/mbmessage/update-message": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

