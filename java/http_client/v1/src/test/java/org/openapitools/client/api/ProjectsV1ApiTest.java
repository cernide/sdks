// Copyright 2018-2023 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 2.0.0-rc5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import java.io.File;
import org.openapitools.client.model.RuntimeError;
import org.openapitools.client.model.V1EntityStageBodyRequest;
import org.openapitools.client.model.V1ListActivitiesResponse;
import org.openapitools.client.model.V1ListBookmarksResponse;
import org.openapitools.client.model.V1ListProjectVersionsResponse;
import org.openapitools.client.model.V1ListProjectsResponse;
import org.openapitools.client.model.V1Project;
import org.openapitools.client.model.V1ProjectSettings;
import org.openapitools.client.model.V1ProjectVersion;
import org.openapitools.client.model.V1Stage;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ProjectsV1Api
 */
@Disabled
public class ProjectsV1ApiTest {

    private final ProjectsV1Api api = new ProjectsV1Api();

    /**
     * Archive project
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void archiveProjectTest() throws ApiException {
        String owner = null;
        String name = null;
        api.archiveProject(owner, name);
        // TODO: test validations
    }

    /**
     * Bookmark project
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void bookmarkProjectTest() throws ApiException {
        String owner = null;
        String name = null;
        api.bookmarkProject(owner, name);
        // TODO: test validations
    }

    /**
     * Create new project
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createProjectTest() throws ApiException {
        String owner = null;
        V1Project body = null;
        V1Project response = api.createProject(owner, body);
        // TODO: test validations
    }

    /**
     * Create version
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createVersionTest() throws ApiException {
        String owner = null;
        String project = null;
        String versionKind = null;
        V1ProjectVersion body = null;
        V1ProjectVersion response = api.createVersion(owner, project, versionKind, body);
        // TODO: test validations
    }

    /**
     * Create new artifact version stage
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createVersionStageTest() throws ApiException {
        String owner = null;
        String entity = null;
        String kind = null;
        String name = null;
        V1EntityStageBodyRequest body = null;
        V1Stage response = api.createVersionStage(owner, entity, kind, name, body);
        // TODO: test validations
    }

    /**
     * Delete project
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteProjectTest() throws ApiException {
        String owner = null;
        String name = null;
        api.deleteProject(owner, name);
        // TODO: test validations
    }

    /**
     * Delete version
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteVersionTest() throws ApiException {
        String owner = null;
        String entity = null;
        String kind = null;
        String name = null;
        api.deleteVersion(owner, entity, kind, name);
        // TODO: test validations
    }

    /**
     * Disbale project CI
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void disableProjectCITest() throws ApiException {
        String owner = null;
        String name = null;
        api.disableProjectCI(owner, name);
        // TODO: test validations
    }

    /**
     * Enable project CI
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void enableProjectCITest() throws ApiException {
        String owner = null;
        String name = null;
        api.enableProjectCI(owner, name);
        // TODO: test validations
    }

    /**
     * Get project
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getProjectTest() throws ApiException {
        String owner = null;
        String name = null;
        V1Project response = api.getProject(owner, name);
        // TODO: test validations
    }

    /**
     * Get project activities
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getProjectActivitiesTest() throws ApiException {
        String owner = null;
        String name = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        V1ListActivitiesResponse response = api.getProjectActivities(owner, name, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * Get Project settings
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getProjectSettingsTest() throws ApiException {
        String owner = null;
        String name = null;
        V1ProjectSettings response = api.getProjectSettings(owner, name);
        // TODO: test validations
    }

    /**
     * Get project stats
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getProjectStatsTest() throws ApiException {
        String owner = null;
        String name = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String kind = null;
        String aggregate = null;
        String groupby = null;
        String trunc = null;
        Object response = api.getProjectStats(owner, name, offset, limit, sort, query, bookmarks, kind, aggregate, groupby, trunc);
        // TODO: test validations
    }

    /**
     * Get version
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getVersionTest() throws ApiException {
        String owner = null;
        String entity = null;
        String kind = null;
        String name = null;
        V1ProjectVersion response = api.getVersion(owner, entity, kind, name);
        // TODO: test validations
    }

    /**
     * Get version stages
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getVersionStagesTest() throws ApiException {
        String owner = null;
        String entity = null;
        String kind = null;
        String name = null;
        V1Stage response = api.getVersionStages(owner, entity, kind, name);
        // TODO: test validations
    }

    /**
     * List archived projects for user
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listArchivedProjectsTest() throws ApiException {
        String user = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean noPage = null;
        V1ListProjectsResponse response = api.listArchivedProjects(user, offset, limit, sort, query, noPage);
        // TODO: test validations
    }

    /**
     * List bookmarked projects for user
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listBookmarkedProjectsTest() throws ApiException {
        String user = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean noPage = null;
        V1ListBookmarksResponse response = api.listBookmarkedProjects(user, offset, limit, sort, query, noPage);
        // TODO: test validations
    }

    /**
     * List project names
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listProjectNamesTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        V1ListProjectsResponse response = api.listProjectNames(owner, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * List projects
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listProjectsTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        V1ListProjectsResponse response = api.listProjects(owner, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * List versions names
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listVersionNamesTest() throws ApiException {
        String owner = null;
        String entity = null;
        String kind = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean noPage = null;
        V1ListProjectVersionsResponse response = api.listVersionNames(owner, entity, kind, offset, limit, sort, query, noPage);
        // TODO: test validations
    }

    /**
     * List versions
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listVersionsTest() throws ApiException {
        String owner = null;
        String entity = null;
        String kind = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean noPage = null;
        V1ListProjectVersionsResponse response = api.listVersions(owner, entity, kind, offset, limit, sort, query, noPage);
        // TODO: test validations
    }

    /**
     * Patch project
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void patchProjectTest() throws ApiException {
        String owner = null;
        String projectName = null;
        V1Project body = null;
        V1Project response = api.patchProject(owner, projectName, body);
        // TODO: test validations
    }

    /**
     * Patch project settings
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void patchProjectSettingsTest() throws ApiException {
        String owner = null;
        String project = null;
        V1ProjectSettings body = null;
        V1ProjectSettings response = api.patchProjectSettings(owner, project, body);
        // TODO: test validations
    }

    /**
     * Patch version
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void patchVersionTest() throws ApiException {
        String owner = null;
        String project = null;
        String versionKind = null;
        String versionName = null;
        V1ProjectVersion body = null;
        V1ProjectVersion response = api.patchVersion(owner, project, versionKind, versionName, body);
        // TODO: test validations
    }

    /**
     * Restore project
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void restoreProjectTest() throws ApiException {
        String owner = null;
        String name = null;
        api.restoreProject(owner, name);
        // TODO: test validations
    }

    /**
     * Transfer version
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void transferVersionTest() throws ApiException {
        String owner = null;
        String project = null;
        String versionKind = null;
        String versionName = null;
        V1ProjectVersion body = null;
        api.transferVersion(owner, project, versionKind, versionName, body);
        // TODO: test validations
    }

    /**
     * Unbookmark project
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void unbookmarkProjectTest() throws ApiException {
        String owner = null;
        String name = null;
        api.unbookmarkProject(owner, name);
        // TODO: test validations
    }

    /**
     * Update project
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateProjectTest() throws ApiException {
        String owner = null;
        String projectName = null;
        V1Project body = null;
        V1Project response = api.updateProject(owner, projectName, body);
        // TODO: test validations
    }

    /**
     * Update project settings
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateProjectSettingsTest() throws ApiException {
        String owner = null;
        String project = null;
        V1ProjectSettings body = null;
        V1ProjectSettings response = api.updateProjectSettings(owner, project, body);
        // TODO: test validations
    }

    /**
     * Update version
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateVersionTest() throws ApiException {
        String owner = null;
        String project = null;
        String versionKind = null;
        String versionName = null;
        V1ProjectVersion body = null;
        V1ProjectVersion response = api.updateVersion(owner, project, versionKind, versionName, body);
        // TODO: test validations
    }

    /**
     * Upload artifact to a store via project access
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void uploadProjectArtifactTest() throws ApiException {
        String owner = null;
        String project = null;
        String uuid = null;
        File uploadfile = null;
        String path = null;
        Boolean overwrite = null;
        api.uploadProjectArtifact(owner, project, uuid, uploadfile, path, overwrite);
        // TODO: test validations
    }

}
