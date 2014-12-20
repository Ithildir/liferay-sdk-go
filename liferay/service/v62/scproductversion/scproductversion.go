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

package scproductversion

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddProductVersion(productEntryId int64, version string, changeLog string, downloadPageURL string, directDownloadURL string, testDirectDownloadURL bool, repoStoreArtifact bool, frameworkVersionIds []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["productEntryId"] = productEntryId
	_params["version"] = version
	_params["changeLog"] = changeLog
	_params["downloadPageURL"] = downloadPageURL
	_params["directDownloadURL"] = directDownloadURL
	_params["testDirectDownloadURL"] = testDirectDownloadURL
	_params["repoStoreArtifact"] = repoStoreArtifact
	_params["frameworkVersionIds"] = frameworkVersionIds
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/scproductversion/add-product-version": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteProductVersion(productVersionId int64) error {
	_params := make(map[string]interface{})

	_params["productVersionId"] = productVersionId

	_cmd := map[string]interface{}{
		"/scproductversion/delete-product-version": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetProductVersion(productVersionId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["productVersionId"] = productVersionId

	_cmd := map[string]interface{}{
		"/scproductversion/get-product-version": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetProductVersions(productEntryId int64, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["productEntryId"] = productEntryId
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/scproductversion/get-product-versions": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetProductVersionsCount(productEntryId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["productEntryId"] = productEntryId

	_cmd := map[string]interface{}{
		"/scproductversion/get-product-versions-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) UpdateProductVersion(productVersionId int64, version string, changeLog string, downloadPageURL string, directDownloadURL string, testDirectDownloadURL bool, repoStoreArtifact bool, frameworkVersionIds []interface{}) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["productVersionId"] = productVersionId
	_params["version"] = version
	_params["changeLog"] = changeLog
	_params["downloadPageURL"] = downloadPageURL
	_params["directDownloadURL"] = directDownloadURL
	_params["testDirectDownloadURL"] = testDirectDownloadURL
	_params["repoStoreArtifact"] = repoStoreArtifact
	_params["frameworkVersionIds"] = frameworkVersionIds

	_cmd := map[string]interface{}{
		"/scproductversion/update-product-version": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

