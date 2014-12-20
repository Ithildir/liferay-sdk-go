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

import com.liferay.mobile.sdk.BaseBuilder;
import com.liferay.mobile.sdk.http.Action;
import com.liferay.mobile.sdk.http.Discovery;
import com.liferay.mobile.sdk.util.CharPool;
import com.liferay.mobile.sdk.util.Validator;
import com.liferay.mobile.sdk.velocity.VelocityUtil;

import java.io.File;
import java.io.IOException;
import java.io.InputStream;

import java.net.JarURLConnection;
import java.net.URL;
import java.net.URLConnection;

import java.util.Enumeration;
import java.util.Iterator;
import java.util.List;
import java.util.jar.JarEntry;
import java.util.jar.JarFile;

import org.apache.commons.io.FileUtils;
import org.apache.commons.io.FilenameUtils;
import org.apache.commons.io.IOUtils;
import org.apache.velocity.VelocityContext;
import org.apache.velocity.tools.generic.EscapeTool;

/**
 * @author Andrea Di Giorgi
 */
public class GoSDKBuilder extends BaseBuilder {

	@Override
	public void build(
			Discovery discovery, List<Action> actions, String packageName,
			int version, String filter, String destination)
		throws Exception {

		copyResource("liferay", destination);
		generateService(discovery, actions, version, filter, destination);
	}

	protected void copyJarResource(
			JarURLConnection jarConnection, File destinationDir)
		throws IOException {

		String jarConnectionEntryName = jarConnection.getEntryName();
		JarFile jarFile = jarConnection.getJarFile();

		Enumeration<JarEntry> enu = jarFile.entries();

		while (enu.hasMoreElements()) {
			JarEntry jarEntry = enu.nextElement();
			String jarEntryName = jarEntry.getName();

			if (jarEntryName.startsWith(jarConnectionEntryName)) {
				String fileName = jarEntryName;

				if (fileName.startsWith(jarConnectionEntryName)) {
					fileName = fileName.substring(
						jarConnectionEntryName.length());
				}

				File file = new File(destinationDir, fileName);

				if (jarEntry.isDirectory()) {
					file.mkdirs();
				}
				else {
					InputStream is = null;

					try {
						is = jarFile.getInputStream(jarEntry);

						FileUtils.copyInputStreamToFile(is, file);
					}
					finally {
						IOUtils.closeQuietly(is);
					}
				}
			}
		}
	}

	protected void copyResource(String name, String destination)
		throws IOException {

		File destinationDir = new File(destination);

		destinationDir = new File(
			destinationDir.getAbsoluteFile(), CharPool.SLASH + name);

		URL sourceURL = getClass().getResource(CharPool.SLASH + name);
		URLConnection sourceConnection = sourceURL.openConnection();

		if (sourceConnection instanceof JarURLConnection) {
			copyJarResource((JarURLConnection)sourceConnection, destinationDir);
		}
		else {
			File sourceDir = new File(sourceURL.getPath());

			FileUtils.copyDirectory(sourceDir, destinationDir);
		}

		Iterator<File> itr = FileUtils.iterateFiles(
			destinationDir, new String[] {"copy"}, true);

		while (itr.hasNext()) {
			File file = itr.next();

			String cleanPath = FilenameUtils.removeExtension(
				file.getAbsolutePath());

			File cleanFile = new File(cleanPath);

			if (!cleanFile.exists()) {
				FileUtils.moveFile(file, new File(cleanPath));
			}
			else {
				file.delete();
			}
		}
	}

	protected void generateService(
			Discovery discovery, List<Action> actions, int version,
			String filter, String destination)
		throws Exception {

		StringBuilder sb = new StringBuilder();

		if (Validator.isNotNull(destination)) {
			sb.append(destination);

			if (!destination.endsWith("/")) {
				sb.append(CharPool.SLASH);
			}
		}

		sb.append("liferay/service");
		destination = sb.toString();

		VelocityContext context = getVelocityContext(
			discovery, actions, filter);

		String templatePath = "templates/go/service.vm";
		String filePath = getServiceFilePath(
			context, version, destination, filter);

		VelocityUtil.generate(context, templatePath, filePath, true);
	}

	protected String getServiceFilePath(
		VelocityContext context, int version, String destination,
		String filter) {

		StringBuilder sb = new StringBuilder();

		sb.append(destination);
		sb.append("/v");
		sb.append(version);
		sb.append(CharPool.SLASH);
		sb.append(filter);
		sb.append(CharPool.SLASH);

		File file = new File(sb.toString());
		file.mkdirs();

		sb.append(filter);
		sb.append(".go");

		return sb.toString();
	}

	protected VelocityContext getVelocityContext(
		Discovery discovery, List<Action> actions, String filter) {

		VelocityContext context = new VelocityContext();

		GoUtil goUtil = new GoUtil();

		context.put(BYTE_SLICE, GoUtil.BYTE_SLICE);
		context.put(ACTIONS, actions);
		context.put(DISCOVERY, discovery);
		context.put(ESCAPE_TOOL, new EscapeTool());
		context.put(JSON_OBJECT_WRAPPER, GoUtil.JSON_OBJECT_WRAPPER);
		context.put(LANGUAGE_UTIL, goUtil);
		context.put(PACKAGE, filter);
		context.put(RESPONSE_VARIABLE, GoUtil.RESPONSE_VARIABLE);
		context.put(VOID, GoUtil.VOID);

		return context;
	}

	protected static final String BYTE_SLICE = "BYTE_SLICE";

	protected static final String PACKAGE = "package";

	protected static final String RESPONSE_VARIABLE = "RESPONSE_VARIABLE";

}