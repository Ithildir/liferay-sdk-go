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

package passwordpolicy

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddPasswordPolicy(name string, description string, changeable bool, changeRequired bool, minAge int64, checkSyntax bool, allowDictionaryWords bool, minAlphanumeric int, minLength int, minLowerCase int, minNumbers int, minSymbols int, minUpperCase int, history bool, historyCount int, expireable bool, maxAge int64, warningTime int64, graceLimit int, lockout bool, maxFailure int, lockoutDuration int64, resetFailureCount int64, resetTicketMaxAge int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name
	_params["description"] = description
	_params["changeable"] = changeable
	_params["changeRequired"] = changeRequired
	_params["minAge"] = minAge
	_params["checkSyntax"] = checkSyntax
	_params["allowDictionaryWords"] = allowDictionaryWords
	_params["minAlphanumeric"] = minAlphanumeric
	_params["minLength"] = minLength
	_params["minLowerCase"] = minLowerCase
	_params["minNumbers"] = minNumbers
	_params["minSymbols"] = minSymbols
	_params["minUpperCase"] = minUpperCase
	_params["history"] = history
	_params["historyCount"] = historyCount
	_params["expireable"] = expireable
	_params["maxAge"] = maxAge
	_params["warningTime"] = warningTime
	_params["graceLimit"] = graceLimit
	_params["lockout"] = lockout
	_params["maxFailure"] = maxFailure
	_params["lockoutDuration"] = lockoutDuration
	_params["resetFailureCount"] = resetFailureCount
	_params["resetTicketMaxAge"] = resetTicketMaxAge

	_cmd := map[string]interface{}{
		"/passwordpolicy/add-password-policy": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddPasswordPolicy2(name string, description string, changeable bool, changeRequired bool, minAge int64, checkSyntax bool, allowDictionaryWords bool, minAlphanumeric int, minLength int, minLowerCase int, minNumbers int, minSymbols int, minUpperCase int, regex string, history bool, historyCount int, expireable bool, maxAge int64, warningTime int64, graceLimit int, lockout bool, maxFailure int, lockoutDuration int64, resetFailureCount int64, resetTicketMaxAge int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name
	_params["description"] = description
	_params["changeable"] = changeable
	_params["changeRequired"] = changeRequired
	_params["minAge"] = minAge
	_params["checkSyntax"] = checkSyntax
	_params["allowDictionaryWords"] = allowDictionaryWords
	_params["minAlphanumeric"] = minAlphanumeric
	_params["minLength"] = minLength
	_params["minLowerCase"] = minLowerCase
	_params["minNumbers"] = minNumbers
	_params["minSymbols"] = minSymbols
	_params["minUpperCase"] = minUpperCase
	_params["regex"] = regex
	_params["history"] = history
	_params["historyCount"] = historyCount
	_params["expireable"] = expireable
	_params["maxAge"] = maxAge
	_params["warningTime"] = warningTime
	_params["graceLimit"] = graceLimit
	_params["lockout"] = lockout
	_params["maxFailure"] = maxFailure
	_params["lockoutDuration"] = lockoutDuration
	_params["resetFailureCount"] = resetFailureCount
	_params["resetTicketMaxAge"] = resetTicketMaxAge
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/passwordpolicy/add-password-policy": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeletePasswordPolicy(passwordPolicyId int64) error {
	_params := make(map[string]interface{})

	_params["passwordPolicyId"] = passwordPolicyId

	_cmd := map[string]interface{}{
		"/passwordpolicy/delete-password-policy": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdatePasswordPolicy(passwordPolicyId int64, name string, description string, changeable bool, changeRequired bool, minAge int64, checkSyntax bool, allowDictionaryWords bool, minAlphanumeric int, minLength int, minLowerCase int, minNumbers int, minSymbols int, minUpperCase int, history bool, historyCount int, expireable bool, maxAge int64, warningTime int64, graceLimit int, lockout bool, maxFailure int, lockoutDuration int64, resetFailureCount int64, resetTicketMaxAge int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["passwordPolicyId"] = passwordPolicyId
	_params["name"] = name
	_params["description"] = description
	_params["changeable"] = changeable
	_params["changeRequired"] = changeRequired
	_params["minAge"] = minAge
	_params["checkSyntax"] = checkSyntax
	_params["allowDictionaryWords"] = allowDictionaryWords
	_params["minAlphanumeric"] = minAlphanumeric
	_params["minLength"] = minLength
	_params["minLowerCase"] = minLowerCase
	_params["minNumbers"] = minNumbers
	_params["minSymbols"] = minSymbols
	_params["minUpperCase"] = minUpperCase
	_params["history"] = history
	_params["historyCount"] = historyCount
	_params["expireable"] = expireable
	_params["maxAge"] = maxAge
	_params["warningTime"] = warningTime
	_params["graceLimit"] = graceLimit
	_params["lockout"] = lockout
	_params["maxFailure"] = maxFailure
	_params["lockoutDuration"] = lockoutDuration
	_params["resetFailureCount"] = resetFailureCount
	_params["resetTicketMaxAge"] = resetTicketMaxAge

	_cmd := map[string]interface{}{
		"/passwordpolicy/update-password-policy": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdatePasswordPolicy2(passwordPolicyId int64, name string, description string, changeable bool, changeRequired bool, minAge int64, checkSyntax bool, allowDictionaryWords bool, minAlphanumeric int, minLength int, minLowerCase int, minNumbers int, minSymbols int, minUpperCase int, regex string, history bool, historyCount int, expireable bool, maxAge int64, warningTime int64, graceLimit int, lockout bool, maxFailure int, lockoutDuration int64, resetFailureCount int64, resetTicketMaxAge int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["passwordPolicyId"] = passwordPolicyId
	_params["name"] = name
	_params["description"] = description
	_params["changeable"] = changeable
	_params["changeRequired"] = changeRequired
	_params["minAge"] = minAge
	_params["checkSyntax"] = checkSyntax
	_params["allowDictionaryWords"] = allowDictionaryWords
	_params["minAlphanumeric"] = minAlphanumeric
	_params["minLength"] = minLength
	_params["minLowerCase"] = minLowerCase
	_params["minNumbers"] = minNumbers
	_params["minSymbols"] = minSymbols
	_params["minUpperCase"] = minUpperCase
	_params["regex"] = regex
	_params["history"] = history
	_params["historyCount"] = historyCount
	_params["expireable"] = expireable
	_params["maxAge"] = maxAge
	_params["warningTime"] = warningTime
	_params["graceLimit"] = graceLimit
	_params["lockout"] = lockout
	_params["maxFailure"] = maxFailure
	_params["lockoutDuration"] = lockoutDuration
	_params["resetFailureCount"] = resetFailureCount
	_params["resetTicketMaxAge"] = resetTicketMaxAge
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/passwordpolicy/update-password-policy": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

