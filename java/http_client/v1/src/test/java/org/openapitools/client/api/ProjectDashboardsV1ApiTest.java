/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.3.3
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.RuntimeError;
import org.openapitools.client.model.V1Dashboard;
import org.openapitools.client.model.V1ListDashboardsResponse;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ProjectDashboardsV1Api
 */
@Disabled
public class ProjectDashboardsV1ApiTest {

    private final ProjectDashboardsV1Api api = new ProjectDashboardsV1Api();

    /**
     * Create project dashboard
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createProjectDashboardTest() throws ApiException {
        String owner = null;
        String project = null;
        V1Dashboard body = null;
        V1Dashboard response = api.createProjectDashboard(owner, project, body);
        // TODO: test validations
    }

    /**
     * Delete project dashboard
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteProjectDashboardTest() throws ApiException {
        String owner = null;
        String entity = null;
        String uuid = null;
        api.deleteProjectDashboard(owner, entity, uuid);
        // TODO: test validations
    }

    /**
     * Get project dashboard
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getProjectDashboardTest() throws ApiException {
        String owner = null;
        String entity = null;
        String uuid = null;
        V1Dashboard response = api.getProjectDashboard(owner, entity, uuid);
        // TODO: test validations
    }

    /**
     * List project dashboard
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listProjectDashboardNamesTest() throws ApiException {
        String owner = null;
        String name = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        V1ListDashboardsResponse response = api.listProjectDashboardNames(owner, name, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * List project dashboards
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listProjectDashboardsTest() throws ApiException {
        String owner = null;
        String name = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        V1ListDashboardsResponse response = api.listProjectDashboards(owner, name, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * Patch project dashboard
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void patchProjectDashboardTest() throws ApiException {
        String owner = null;
        String project = null;
        String dashboardUuid = null;
        V1Dashboard body = null;
        V1Dashboard response = api.patchProjectDashboard(owner, project, dashboardUuid, body);
        // TODO: test validations
    }

    /**
     * Promote project dashboard
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void promoteProjectDashboardTest() throws ApiException {
        String owner = null;
        String entity = null;
        String uuid = null;
        api.promoteProjectDashboard(owner, entity, uuid);
        // TODO: test validations
    }

    /**
     * Update project dashboard
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateProjectDashboardTest() throws ApiException {
        String owner = null;
        String project = null;
        String dashboardUuid = null;
        V1Dashboard body = null;
        V1Dashboard response = api.updateProjectDashboard(owner, project, dashboardUuid, body);
        // TODO: test validations
    }

}
