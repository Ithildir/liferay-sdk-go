/**
 * Copyright (c) 2014 Andrea Di Giorgi. All rights reserved.
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

package com.github.ithildir.liferay.mobile.go;

import com.liferay.mobile.sdk.http.Action;
import com.liferay.mobile.sdk.http.Parameter;
import com.liferay.mobile.sdk.util.CharPool;
import com.liferay.mobile.sdk.util.LanguageUtil;

import java.util.List;

import org.apache.commons.lang.ArrayUtils;

/**
 * @author Andrea Di Giorgi
 */
public class GoUtil extends LanguageUtil {

	public static final String BOOL = "bool";

	public static final String BYTE_SLICE = "[]byte";

	public static final String FLOAT64 = "float64";

	public static final String INT64 = "int64";

	public static final String IO_READER = "io.Reader";

	public static final String JSON_ARRAY = "[]interface{}";

	public static final String JSON_OBJECT = "map[string]interface{}";

	public static final String JSON_OBJECT_WRAPPER = "*liferay.ObjectWrapper";

	public static final String RESPONSE_VARIABLE = "res";

	public String getMethodName(List<Action> actions, int index) {
		Action action = actions.get(index);

		String last = getMethodURL(action.getPath());

		String[] methodName = last.split("-");

		StringBuilder sb = new StringBuilder();

		for (int i = 0; i < methodName.length; i++) {
			String word = capitalize(methodName[i]);

			sb.append(word);
		}

		int methodOverloadIndex = getMethodOverloadIndex(actions, index);

		if (methodOverloadIndex > 0) {
			sb.append(methodOverloadIndex);
		}

		return sb.toString();
	}

	public String getParameterName(String parameterName) {
		if (ArrayUtils.contains(_RESERVED_KEYWORDS, parameterName)) {
			return CharPool.UNDERLINE + parameterName;
		}

		return parameterName;
	}

	public String getResponseVariableConversion(String returnType) {
		if (returnType.equals(INT) || returnType.equals(INT64)) {
			StringBuilder sb = new StringBuilder();

			sb.append(returnType);
			sb.append(CharPool.OPEN_PARENTHESIS);
			sb.append(RESPONSE_VARIABLE);
			sb.append(".(float64))");

			return sb.toString();
		}

		return
			RESPONSE_VARIABLE + ".(" + returnType + CharPool.CLOSE_PARENTHESIS;
	}

	public String getReturnType(String type) {
		type = getType(type);

		if (type.equals(VOID)) {
			return type;
		}

		if (type.equals(IO_READER) || type.equals(JSON_OBJECT_WRAPPER)) {
			return JSON_OBJECT;
		}

		if (type.equals(BYTE_SLICE)) {
			return JSON_ARRAY;
		}

		return type;
	}

	public String getType(String type) {
		if (type.equals(INT) || type.equals(STRING) || type.equals(VOID)) {
			return type;
		}

		if (type.equals(BOOLEAN)) {
			return BOOL;
		}

		if (type.equals(BYTE_ARRAY)) {
			return BYTE_SLICE;
		}

		if (type.equals(DOUBLE)) {
			return FLOAT64;
		}

		if (type.equals(FILE)) {
			return IO_READER;
		}

		if (type.equals(LONG)) {
			return INT64;
		}

		if (isArray(type)) {
			return JSON_ARRAY;
		}

		if (type.startsWith(OBJECT_PREFIX)) {
			return JSON_OBJECT_WRAPPER;
		}

		return JSON_OBJECT;
	}

	public boolean isIOPackageRequired(List<Action> actions) {
		if (_hasParameterType(actions, FILE)) {
			return true;
		}
		else {
			return false;
		}
	}

	protected int getMethodOverloadIndex(List<Action> actions, int index) {
		Action action = actions.get(index);

		String path = action.getPath();

		int overloadIndex = 0;

		for (int i = index - 1; i >= 0; i--) {
			Action curAction = actions.get(i);

			if (!path.equals(curAction.getPath())) {
				break;
			}

			overloadIndex++;
		}

		if (overloadIndex > 0) {
			overloadIndex++;
		}

		return overloadIndex;
	}

	private boolean _hasParameterType(List<Action> actions, String type) {
		for (Action action : actions) {
			for (Parameter parameter : action.getParameters()) {
				if (type.equals(parameter.getType())) {
					return true;
				}
			}
		}

		return false;
	}

	private static final String[] _RESERVED_KEYWORDS = {
		"break", "case", "chan", "const", "continue", "default", "defer",
		"else", "fallthrough", "for", "func", "go", "goto", "if", "import",
		"interface", "map", "package", "range", "return", "select", "struct",
		"switch", "type", "var"};

}