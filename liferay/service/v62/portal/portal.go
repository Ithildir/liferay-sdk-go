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

package portal

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) GetAutoDeployDirectory() (string, error) {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/portal/get-auto-deploy-directory": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetBuildNumber() (int, error) {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/portal/get-build-number": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) TestAddClassNameAndTestTransactionPortletBar_PortalRollback(transactionPortletBarText string) error {
	_params := make(map[string]interface{})

	_params["transactionPortletBarText"] = transactionPortletBarText

	_cmd := map[string]interface{}{
		"/portal/test-add-class-name-and-test-transaction-portlet-bar_-portal-rollback": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) TestAddClassNameAndTestTransactionPortletBar_PortletRollback(transactionPortletBarText string) error {
	_params := make(map[string]interface{})

	_params["transactionPortletBarText"] = transactionPortletBarText

	_cmd := map[string]interface{}{
		"/portal/test-add-class-name-and-test-transaction-portlet-bar_-portlet-rollback": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) TestAddClassNameAndTestTransactionPortletBar_Success(transactionPortletBarText string) error {
	_params := make(map[string]interface{})

	_params["transactionPortletBarText"] = transactionPortletBarText

	_cmd := map[string]interface{}{
		"/portal/test-add-class-name-and-test-transaction-portlet-bar_-success": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) TestAddClassName_Rollback(classNameValue string) error {
	_params := make(map[string]interface{})

	_params["classNameValue"] = classNameValue

	_cmd := map[string]interface{}{
		"/portal/test-add-class-name_-rollback": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) TestAddClassName_Success(classNameValue string) error {
	_params := make(map[string]interface{})

	_params["classNameValue"] = classNameValue

	_cmd := map[string]interface{}{
		"/portal/test-add-class-name_-success": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) TestAutoSyncHibernateSessionStateOnTxCreation() error {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/portal/test-auto-sync-hibernate-session-state-on-tx-creation": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) TestDeleteClassName() error {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/portal/test-delete-class-name": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) TestGetBuildNumber() (int, error) {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/portal/test-get-build-number": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) TestGetUserId() error {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/portal/test-get-user-id": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) TestHasClassName() (bool, error) {
	_params := make(map[string]interface{})

	_cmd := map[string]interface{}{
		"/portal/test-has-class-name": _params,
	}

	var v bool

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(bool)
	}

	return v, err
}

