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

package shoppingcoupon

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddCoupon(code string, autoCode bool, name string, description string, startDateMonth int, startDateDay int, startDateYear int, startDateHour int, startDateMinute int, endDateMonth int, endDateDay int, endDateYear int, endDateHour int, endDateMinute int, neverExpire bool, active bool, limitCategories string, limitSkus string, minOrder float64, discount float64, discountType string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["code"] = code
	_params["autoCode"] = autoCode
	_params["name"] = name
	_params["description"] = description
	_params["startDateMonth"] = startDateMonth
	_params["startDateDay"] = startDateDay
	_params["startDateYear"] = startDateYear
	_params["startDateHour"] = startDateHour
	_params["startDateMinute"] = startDateMinute
	_params["endDateMonth"] = endDateMonth
	_params["endDateDay"] = endDateDay
	_params["endDateYear"] = endDateYear
	_params["endDateHour"] = endDateHour
	_params["endDateMinute"] = endDateMinute
	_params["neverExpire"] = neverExpire
	_params["active"] = active
	_params["limitCategories"] = limitCategories
	_params["limitSkus"] = limitSkus
	_params["minOrder"] = minOrder
	_params["discount"] = discount
	_params["discountType"] = discountType
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/shoppingcoupon/add-coupon": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteCoupon(groupId int64, couponId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["couponId"] = couponId

	_cmd := map[string]interface{}{
		"/shoppingcoupon/delete-coupon": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetCoupon(groupId int64, couponId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["couponId"] = couponId

	_cmd := map[string]interface{}{
		"/shoppingcoupon/get-coupon": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) Search(groupId int64, companyId int64, code string, active bool, discountType string, andOperator bool, start int, end int) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["companyId"] = companyId
	_params["code"] = code
	_params["active"] = active
	_params["discountType"] = discountType
	_params["andOperator"] = andOperator
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/shoppingcoupon/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) UpdateCoupon(couponId int64, name string, description string, startDateMonth int, startDateDay int, startDateYear int, startDateHour int, startDateMinute int, endDateMonth int, endDateDay int, endDateYear int, endDateHour int, endDateMinute int, neverExpire bool, active bool, limitCategories string, limitSkus string, minOrder float64, discount float64, discountType string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["couponId"] = couponId
	_params["name"] = name
	_params["description"] = description
	_params["startDateMonth"] = startDateMonth
	_params["startDateDay"] = startDateDay
	_params["startDateYear"] = startDateYear
	_params["startDateHour"] = startDateHour
	_params["startDateMinute"] = startDateMinute
	_params["endDateMonth"] = endDateMonth
	_params["endDateDay"] = endDateDay
	_params["endDateYear"] = endDateYear
	_params["endDateHour"] = endDateHour
	_params["endDateMinute"] = endDateMinute
	_params["neverExpire"] = neverExpire
	_params["active"] = active
	_params["limitCategories"] = limitCategories
	_params["limitSkus"] = limitSkus
	_params["minOrder"] = minOrder
	_params["discount"] = discount
	_params["discountType"] = discountType
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/shoppingcoupon/update-coupon": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

