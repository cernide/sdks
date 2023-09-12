/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc41
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.RuntimeError;
import org.openapitools.client.model.V1Agent;
import org.openapitools.client.model.V1AgentStateResponse;
import org.openapitools.client.model.V1AgentStatusBodyRequest;
import org.openapitools.client.model.V1ListAgentsResponse;
import org.openapitools.client.model.V1Status;
import org.openapitools.client.model.V1Token;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for AgentsV1Api
 */
@Disabled
public class AgentsV1ApiTest {

    private final AgentsV1Api api = new AgentsV1Api();

    /**
     * Create agent
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createAgentTest() throws ApiException {
        String owner = null;
        V1Agent body = null;
        V1Agent response = api.createAgent(owner, body);
        // TODO: test validations
    }

    /**
     * Create new agent status
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createAgentStatusTest() throws ApiException {
        String owner = null;
        String uuid = null;
        V1AgentStatusBodyRequest body = null;
        V1Status response = api.createAgentStatus(owner, uuid, body);
        // TODO: test validations
    }

    /**
     * Global Cron
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void cronAgentTest() throws ApiException {
        String owner = null;
        api.cronAgent(owner);
        // TODO: test validations
    }

    /**
     * Delete agent
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteAgentTest() throws ApiException {
        String owner = null;
        String uuid = null;
        String entity = null;
        api.deleteAgent(owner, uuid, entity);
        // TODO: test validations
    }

    /**
     * Get agent
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getAgentTest() throws ApiException {
        String owner = null;
        String uuid = null;
        String entity = null;
        V1Agent response = api.getAgent(owner, uuid, entity);
        // TODO: test validations
    }

    /**
     * Get agent config
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getAgentConfigTest() throws ApiException {
        String owner = null;
        String uuid = null;
        String entity = null;
        V1Agent response = api.getAgentConfig(owner, uuid, entity);
        // TODO: test validations
    }

    /**
     * Get State (queues/runs)
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getAgentStateTest() throws ApiException {
        String owner = null;
        String uuid = null;
        V1AgentStateResponse response = api.getAgentState(owner, uuid);
        // TODO: test validations
    }

    /**
     * Get agent token
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getAgentTokenTest() throws ApiException {
        String owner = null;
        String uuid = null;
        V1Token response = api.getAgentToken(owner, uuid);
        // TODO: test validations
    }

    /**
     * Get Global State (queues/runs)
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getGlobalStateTest() throws ApiException {
        String owner = null;
        V1AgentStateResponse response = api.getGlobalState(owner);
        // TODO: test validations
    }

    /**
     * List agents names
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listAgentNamesTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        V1ListAgentsResponse response = api.listAgentNames(owner, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * List agents
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listAgentsTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        V1ListAgentsResponse response = api.listAgents(owner, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * Patch agent
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void patchAgentTest() throws ApiException {
        String owner = null;
        String agentUuid = null;
        V1Agent body = null;
        V1Agent response = api.patchAgent(owner, agentUuid, body);
        // TODO: test validations
    }

    /**
     * Patch agent token
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void patchAgentTokenTest() throws ApiException {
        String owner = null;
        String entity = null;
        V1Token body = null;
        V1Token response = api.patchAgentToken(owner, entity, body);
        // TODO: test validations
    }

    /**
     * Sync agent
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void syncAgentTest() throws ApiException {
        String owner = null;
        String agentUuid = null;
        V1Agent body = null;
        api.syncAgent(owner, agentUuid, body);
        // TODO: test validations
    }

    /**
     * Update agent
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateAgentTest() throws ApiException {
        String owner = null;
        String agentUuid = null;
        V1Agent body = null;
        V1Agent response = api.updateAgent(owner, agentUuid, body);
        // TODO: test validations
    }

    /**
     * Update agent config
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateAgentConfigTest() throws ApiException {
        String owner = null;
        String agentUuid = null;
        V1Agent body = null;
        V1Agent response = api.updateAgentConfig(owner, agentUuid, body);
        // TODO: test validations
    }

    /**
     * Update agent token
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateAgentTokenTest() throws ApiException {
        String owner = null;
        String entity = null;
        V1Token body = null;
        V1Token response = api.updateAgentToken(owner, entity, body);
        // TODO: test validations
    }

}
