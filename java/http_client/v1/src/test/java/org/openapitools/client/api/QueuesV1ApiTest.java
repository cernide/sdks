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
 * The version of the OpenAPI document: 2.0.0-rc12
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.RuntimeError;
import org.openapitools.client.model.V1ListQueuesResponse;
import org.openapitools.client.model.V1Queue;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for QueuesV1Api
 */
@Disabled
public class QueuesV1ApiTest {

    private final QueuesV1Api api = new QueuesV1Api();

    /**
     * Create queue
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createQueueTest() throws ApiException {
        String owner = null;
        String agent = null;
        V1Queue body = null;
        V1Queue response = api.createQueue(owner, agent, body);
        // TODO: test validations
    }

    /**
     * Delete queue
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteQueueTest() throws ApiException {
        String owner = null;
        String entity = null;
        String uuid = null;
        api.deleteQueue(owner, entity, uuid);
        // TODO: test validations
    }

    /**
     * Get queue
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getQueueTest() throws ApiException {
        String owner = null;
        String entity = null;
        String uuid = null;
        V1Queue response = api.getQueue(owner, entity, uuid);
        // TODO: test validations
    }

    /**
     * List organization level queues names
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listOrganizationQueueNamesTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        V1ListQueuesResponse response = api.listOrganizationQueueNames(owner, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * List organization level queues
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listOrganizationQueuesTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        V1ListQueuesResponse response = api.listOrganizationQueues(owner, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * List queues names
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listQueueNamesTest() throws ApiException {
        String owner = null;
        String name = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        V1ListQueuesResponse response = api.listQueueNames(owner, name, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * List queues
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listQueuesTest() throws ApiException {
        String owner = null;
        String name = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        V1ListQueuesResponse response = api.listQueues(owner, name, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * Patch queue
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void patchQueueTest() throws ApiException {
        String owner = null;
        String agent = null;
        String queueUuid = null;
        V1Queue body = null;
        V1Queue response = api.patchQueue(owner, agent, queueUuid, body);
        // TODO: test validations
    }

    /**
     * Update queue
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateQueueTest() throws ApiException {
        String owner = null;
        String agent = null;
        String queueUuid = null;
        V1Queue body = null;
        V1Queue response = api.updateQueue(owner, agent, queueUuid, body);
        // TODO: test validations
    }

}
