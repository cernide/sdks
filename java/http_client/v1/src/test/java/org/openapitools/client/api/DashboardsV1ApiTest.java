// Copyright 2018-2021 Polyaxon, Inc.
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
 * The version of the OpenAPI document: 1.17.0
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
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for DashboardsV1Api
 */
@Ignore
public class DashboardsV1ApiTest {

    private final DashboardsV1Api api = new DashboardsV1Api();

    
    /**
     * Create dashboard
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createDashboardTest() throws ApiException {
        String owner = null;
        V1Dashboard body = null;
                V1Dashboard response = api.createDashboard(owner, body);
        // TODO: test validations
    }
    
    /**
     * Delete dashboard
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteDashboardTest() throws ApiException {
        String owner = null;
        String uuid = null;
                api.deleteDashboard(owner, uuid);
        // TODO: test validations
    }
    
    /**
     * Get dashboard
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getDashboardTest() throws ApiException {
        String owner = null;
        String uuid = null;
                V1Dashboard response = api.getDashboard(owner, uuid);
        // TODO: test validations
    }
    
    /**
     * List dashboard names
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listDashboardNamesTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
                V1ListDashboardsResponse response = api.listDashboardNames(owner, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }
    
    /**
     * List dashboards
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listDashboardsTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
                V1ListDashboardsResponse response = api.listDashboards(owner, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }
    
    /**
     * Patch dashboard
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchDashboardTest() throws ApiException {
        String owner = null;
        String dashboardUuid = null;
        V1Dashboard body = null;
                V1Dashboard response = api.patchDashboard(owner, dashboardUuid, body);
        // TODO: test validations
    }
    
    /**
     * Update dashboard
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateDashboardTest() throws ApiException {
        String owner = null;
        String dashboardUuid = null;
        V1Dashboard body = null;
                V1Dashboard response = api.updateDashboard(owner, dashboardUuid, body);
        // TODO: test validations
    }
    
}
