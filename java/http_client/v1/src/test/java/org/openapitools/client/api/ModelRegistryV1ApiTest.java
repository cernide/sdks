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
 * The version of the OpenAPI document: 1.13.0-rc3
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.RuntimeError;
import org.openapitools.client.model.V1EntityStageBodyRequest;
import org.openapitools.client.model.V1ListActivitiesResponse;
import org.openapitools.client.model.V1ListModelRegistriesResponse;
import org.openapitools.client.model.V1ListModelVersionsResponse;
import org.openapitools.client.model.V1ModelRegistry;
import org.openapitools.client.model.V1ModelRegistrySettings;
import org.openapitools.client.model.V1ModelVersion;
import org.openapitools.client.model.V1Stage;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ModelRegistryV1Api
 */
@Ignore
public class ModelRegistryV1ApiTest {

    private final ModelRegistryV1Api api = new ModelRegistryV1Api();


    /**
     * Archive registry model
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void archiveModelRegistryTest() throws ApiException {
        String owner = null;
        String name = null;
                api.archiveModelRegistry(owner, name);
        // TODO: test validations
    }

    /**
     * Bookmark registry model
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void bookmarkModelRegistryTest() throws ApiException {
        String owner = null;
        String name = null;
                api.bookmarkModelRegistry(owner, name);
        // TODO: test validations
    }

    /**
     * Create registry model
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createModelRegistryTest() throws ApiException {
        String owner = null;
        V1ModelRegistry body = null;
                V1ModelRegistry response = api.createModelRegistry(owner, body);
        // TODO: test validations
    }

    /**
     * Create model version
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createModelVersionTest() throws ApiException {
        String owner = null;
        String model = null;
        V1ModelVersion body = null;
                V1ModelVersion response = api.createModelVersion(owner, model, body);
        // TODO: test validations
    }

    /**
     * Create new model version stage
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createModelVersionStageTest() throws ApiException {
        String owner = null;
        String entity = null;
        String name = null;
        V1EntityStageBodyRequest body = null;
                V1Stage response = api.createModelVersionStage(owner, entity, name, body);
        // TODO: test validations
    }

    /**
     * Delete registry model
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteModelRegistryTest() throws ApiException {
        String owner = null;
        String name = null;
                api.deleteModelRegistry(owner, name);
        // TODO: test validations
    }

    /**
     * Delete model version
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteModelVersionTest() throws ApiException {
        String owner = null;
        String entity = null;
        String name = null;
                api.deleteModelVersion(owner, entity, name);
        // TODO: test validations
    }

    /**
     * Get registry model
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getModelRegistryTest() throws ApiException {
        String owner = null;
        String name = null;
                V1ModelRegistry response = api.getModelRegistry(owner, name);
        // TODO: test validations
    }

    /**
     * Get model activities
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getModelRegistryActivitiesTest() throws ApiException {
        String owner = null;
        String name = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
                V1ListActivitiesResponse response = api.getModelRegistryActivities(owner, name, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * Get registry model settings
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getModelRegistrySettingsTest() throws ApiException {
        String owner = null;
        String name = null;
                V1ModelRegistrySettings response = api.getModelRegistrySettings(owner, name);
        // TODO: test validations
    }

    /**
     * Get model version
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getModelVersionTest() throws ApiException {
        String owner = null;
        String entity = null;
        String name = null;
                V1ModelVersion response = api.getModelVersion(owner, entity, name);
        // TODO: test validations
    }

    /**
     * Get model version stages
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getModelVersionStagesTest() throws ApiException {
        String owner = null;
        String entity = null;
        String name = null;
                V1Stage response = api.getModelVersionStages(owner, entity, name);
        // TODO: test validations
    }

    /**
     * List registry models
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listModelRegistriesTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
                V1ListModelRegistriesResponse response = api.listModelRegistries(owner, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * List registry model names
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listModelRegistryNamesTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
                V1ListModelRegistriesResponse response = api.listModelRegistryNames(owner, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * List model versions names
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listModelVersionNamesTest() throws ApiException {
        String owner = null;
        String name = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
                V1ListModelVersionsResponse response = api.listModelVersionNames(owner, name, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * List model versions
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listModelVersionsTest() throws ApiException {
        String owner = null;
        String name = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
                V1ListModelVersionsResponse response = api.listModelVersions(owner, name, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * Patch registry model
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchModelRegistryTest() throws ApiException {
        String owner = null;
        String modelName = null;
        V1ModelRegistry body = null;
                V1ModelRegistry response = api.patchModelRegistry(owner, modelName, body);
        // TODO: test validations
    }

    /**
     * Patch registry model settings
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchModelRegistrySettingsTest() throws ApiException {
        String owner = null;
        String model = null;
        V1ModelRegistrySettings body = null;
                V1ModelRegistrySettings response = api.patchModelRegistrySettings(owner, model, body);
        // TODO: test validations
    }

    /**
     * Patch model version
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchModelVersionTest() throws ApiException {
        String owner = null;
        String model = null;
        String versionName = null;
        V1ModelVersion body = null;
                V1ModelVersion response = api.patchModelVersion(owner, model, versionName, body);
        // TODO: test validations
    }

    /**
     * Restore registry model
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void restoreModelRegistryTest() throws ApiException {
        String owner = null;
        String name = null;
                api.restoreModelRegistry(owner, name);
        // TODO: test validations
    }

    /**
     * Unbookmark registry model
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void unbookmarkModelRegistryTest() throws ApiException {
        String owner = null;
        String name = null;
                api.unbookmarkModelRegistry(owner, name);
        // TODO: test validations
    }

    /**
     * Update registry model
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateModelRegistryTest() throws ApiException {
        String owner = null;
        String modelName = null;
        V1ModelRegistry body = null;
                V1ModelRegistry response = api.updateModelRegistry(owner, modelName, body);
        // TODO: test validations
    }

    /**
     * Update registry model settings
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateModelRegistrySettingsTest() throws ApiException {
        String owner = null;
        String model = null;
        V1ModelRegistrySettings body = null;
                V1ModelRegistrySettings response = api.updateModelRegistrySettings(owner, model, body);
        // TODO: test validations
    }

    /**
     * Update model version
     *
     *
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateModelVersionTest() throws ApiException {
        String owner = null;
        String model = null;
        String versionName = null;
        V1ModelVersion body = null;
                V1ModelVersion response = api.updateModelVersion(owner, model, versionName, body);
        // TODO: test validations
    }

}