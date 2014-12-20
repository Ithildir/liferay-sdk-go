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

package pollsquestion

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddQuestion(titleMap map[string]interface{}, descriptionMap map[string]interface{}, expirationDateMonth int, expirationDateDay int, expirationDateYear int, expirationDateHour int, expirationDateMinute int, neverExpire bool, choices []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["expirationDateMonth"] = expirationDateMonth
	_params["expirationDateDay"] = expirationDateDay
	_params["expirationDateYear"] = expirationDateYear
	_params["expirationDateHour"] = expirationDateHour
	_params["expirationDateMinute"] = expirationDateMinute
	_params["neverExpire"] = neverExpire
	_params["choices"] = choices
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/pollsquestion/add-question": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteQuestion(questionId int64) error {
	_params := make(map[string]interface{})

	_params["questionId"] = questionId

	_cmd := map[string]interface{}{
		"/pollsquestion/delete-question": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetQuestion(questionId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["questionId"] = questionId

	_cmd := map[string]interface{}{
		"/pollsquestion/get-question": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateQuestion(questionId int64, titleMap map[string]interface{}, descriptionMap map[string]interface{}, expirationDateMonth int, expirationDateDay int, expirationDateYear int, expirationDateHour int, expirationDateMinute int, neverExpire bool, choices []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["questionId"] = questionId
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["expirationDateMonth"] = expirationDateMonth
	_params["expirationDateDay"] = expirationDateDay
	_params["expirationDateYear"] = expirationDateYear
	_params["expirationDateHour"] = expirationDateHour
	_params["expirationDateMinute"] = expirationDateMinute
	_params["neverExpire"] = neverExpire
	_params["choices"] = choices
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/pollsquestion/update-question": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

