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

package scproductentry

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddProductEntry(name string, _type string, tags string, shortDescription string, longDescription string, pageURL string, author string, repoGroupId string, repoArtifactId string, licenseIds []interface{}, thumbnails []interface{}, fullImages []interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["name"] = name
	_params["type"] = _type
	_params["tags"] = tags
	_params["shortDescription"] = shortDescription
	_params["longDescription"] = longDescription
	_params["pageURL"] = pageURL
	_params["author"] = author
	_params["repoGroupId"] = repoGroupId
	_params["repoArtifactId"] = repoArtifactId
	_params["licenseIds"] = licenseIds
	_params["thumbnails"] = thumbnails
	_params["fullImages"] = fullImages
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/scproductentry/add-product-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteProductEntry(productEntryId int64) error {
	_params := make(map[string]interface{})

	_params["productEntryId"] = productEntryId

	_cmd := map[string]interface{}{
		"/scproductentry/delete-product-entry": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetProductEntry(productEntryId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["productEntryId"] = productEntryId

	_cmd := map[string]interface{}{
		"/scproductentry/get-product-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateProductEntry(productEntryId int64, name string, _type string, tags string, shortDescription string, longDescription string, pageURL string, author string, repoGroupId string, repoArtifactId string, licenseIds []interface{}, thumbnails []interface{}, fullImages []interface{}) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["productEntryId"] = productEntryId
	_params["name"] = name
	_params["type"] = _type
	_params["tags"] = tags
	_params["shortDescription"] = shortDescription
	_params["longDescription"] = longDescription
	_params["pageURL"] = pageURL
	_params["author"] = author
	_params["repoGroupId"] = repoGroupId
	_params["repoArtifactId"] = repoArtifactId
	_params["licenseIds"] = licenseIds
	_params["thumbnails"] = thumbnails
	_params["fullImages"] = fullImages

	_cmd := map[string]interface{}{
		"/scproductentry/update-product-entry": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

