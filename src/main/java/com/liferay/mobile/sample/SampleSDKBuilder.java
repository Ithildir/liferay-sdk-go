/**
 * Copyright (c) 2000-2014 Liferay, Inc. All rights reserved.
 *
 * This library is free software; you can redistribute it and/or modify it under
 * the terms of the GNU Lesser General Public License as published by the Free
 * Software Foundation; either version 2.1 of the License, or (at your option)
 * any later version.
 *
 * This library is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
 * FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more
 * details.
 */

package com.liferay.mobile.sample;

import com.liferay.mobile.sdk.BaseBuilder;
import com.liferay.mobile.sdk.http.Discovery;

/**
 * @author Bruno Farache
 */
public class SampleSDKBuilder extends BaseBuilder {

	@Override
	public void build(
			Discovery discovery, String packageName, int version, String filter,
			String destination)
		throws Exception {

		System.out.println(
			"I'm just a dummy SDK Builder, I don' generate anything.");
	}

}