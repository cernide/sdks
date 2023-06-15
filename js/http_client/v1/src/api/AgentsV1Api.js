/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc19
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import RuntimeError from '../model/RuntimeError';
import V1Agent from '../model/V1Agent';
import V1AgentStateResponse from '../model/V1AgentStateResponse';
import V1AgentStatusBodyRequest from '../model/V1AgentStatusBodyRequest';
import V1ListAgentsResponse from '../model/V1ListAgentsResponse';
import V1Status from '../model/V1Status';
import V1Token from '../model/V1Token';

/**
* AgentsV1 service.
* @module api/AgentsV1Api
* @version 2.0.0-rc19
*/
export default class AgentsV1Api {

    /**
    * Constructs a new AgentsV1Api.
    * Polyaxon sdk
    * @alias module:api/AgentsV1Api
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the createAgent operation.
     * @callback module:api/AgentsV1Api~createAgentCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Agent} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create agent
     * @param {String} owner Owner of the namespace
     * @param {module:model/V1Agent} body Agent body
     * @param {module:api/AgentsV1Api~createAgentCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Agent}
     */
    createAgent(owner, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling createAgent");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createAgent");
      }

      let pathParams = {
        'owner': owner
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1Agent;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the createAgentStatus operation.
     * @callback module:api/AgentsV1Api~createAgentStatusCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Status} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create new agent status
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:model/V1AgentStatusBodyRequest} body
     * @param {module:api/AgentsV1Api~createAgentStatusCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Status}
     */
    createAgentStatus(owner, uuid, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling createAgentStatus");
      }
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling createAgentStatus");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createAgentStatus");
      }

      let pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1Status;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{uuid}/statuses', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteAgent operation.
     * @callback module:api/AgentsV1Api~deleteAgentCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete agent
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the sub-entity
     * @param {Object} opts Optional parameters
     * @param {String} [entity] Entity: project name, hub name, registry name, ...
     * @param {module:api/AgentsV1Api~deleteAgentCallback} callback The callback function, accepting three arguments: error, data, response
     */
    deleteAgent(owner, uuid, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling deleteAgent");
      }
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling deleteAgent");
      }

      let pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      let queryParams = {
        'entity': opts['entity']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{uuid}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getAgent operation.
     * @callback module:api/AgentsV1Api~getAgentCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Agent} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get agent
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the sub-entity
     * @param {Object} opts Optional parameters
     * @param {String} [entity] Entity: project name, hub name, registry name, ...
     * @param {module:api/AgentsV1Api~getAgentCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Agent}
     */
    getAgent(owner, uuid, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getAgent");
      }
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling getAgent");
      }

      let pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      let queryParams = {
        'entity': opts['entity']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1Agent;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{uuid}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getAgentConfig operation.
     * @callback module:api/AgentsV1Api~getAgentConfigCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Agent} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get agent config
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the sub-entity
     * @param {Object} opts Optional parameters
     * @param {String} [entity] Entity: project name, hub name, registry name, ...
     * @param {module:api/AgentsV1Api~getAgentConfigCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Agent}
     */
    getAgentConfig(owner, uuid, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getAgentConfig");
      }
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling getAgentConfig");
      }

      let pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      let queryParams = {
        'entity': opts['entity']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1Agent;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{uuid}/config', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getAgentState operation.
     * @callback module:api/AgentsV1Api~getAgentStateCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1AgentStateResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get State (queues/runs)
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/AgentsV1Api~getAgentStateCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1AgentStateResponse}
     */
    getAgentState(owner, uuid, callback) {
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getAgentState");
      }
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling getAgentState");
      }

      let pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1AgentStateResponse;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{uuid}/state', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getAgentToken operation.
     * @callback module:api/AgentsV1Api~getAgentTokenCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Token} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get agent token
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/AgentsV1Api~getAgentTokenCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Token}
     */
    getAgentToken(owner, uuid, callback) {
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getAgentToken");
      }
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling getAgentToken");
      }

      let pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1Token;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{uuid}/token', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getGlobalState operation.
     * @callback module:api/AgentsV1Api~getGlobalStateCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1AgentStateResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get Global State (queues/runs)
     * @param {String} owner Owner of the namespace
     * @param {module:api/AgentsV1Api~getGlobalStateCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1AgentStateResponse}
     */
    getGlobalState(owner, callback) {
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getGlobalState");
      }

      let pathParams = {
        'owner': owner
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1AgentStateResponse;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/state', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listAgentNames operation.
     * @callback module:api/AgentsV1Api~listAgentNamesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListAgentsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List agents names
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} [offset] Pagination offset.
     * @param {Number} [limit] Limit size.
     * @param {String} [sort] Sort to order the search.
     * @param {String} [query] Query filter the search.
     * @param {Boolean} [bookmarks] Filter by bookmarks.
     * @param {String} [mode] Mode of the search.
     * @param {Boolean} [no_page] No pagination.
     * @param {module:api/AgentsV1Api~listAgentNamesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListAgentsResponse}
     */
    listAgentNames(owner, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listAgentNames");
      }

      let pathParams = {
        'owner': owner
      };
      let queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query'],
        'bookmarks': opts['bookmarks'],
        'mode': opts['mode'],
        'no_page': opts['no_page']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1ListAgentsResponse;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/names', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listAgents operation.
     * @callback module:api/AgentsV1Api~listAgentsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListAgentsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List agents
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} [offset] Pagination offset.
     * @param {Number} [limit] Limit size.
     * @param {String} [sort] Sort to order the search.
     * @param {String} [query] Query filter the search.
     * @param {Boolean} [bookmarks] Filter by bookmarks.
     * @param {String} [mode] Mode of the search.
     * @param {Boolean} [no_page] No pagination.
     * @param {module:api/AgentsV1Api~listAgentsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListAgentsResponse}
     */
    listAgents(owner, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listAgents");
      }

      let pathParams = {
        'owner': owner
      };
      let queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query'],
        'bookmarks': opts['bookmarks'],
        'mode': opts['mode'],
        'no_page': opts['no_page']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1ListAgentsResponse;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the patchAgent operation.
     * @callback module:api/AgentsV1Api~patchAgentCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Agent} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch agent
     * @param {String} owner Owner of the namespace
     * @param {String} agent_uuid UUID
     * @param {module:model/V1Agent} body Agent body
     * @param {module:api/AgentsV1Api~patchAgentCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Agent}
     */
    patchAgent(owner, agent_uuid, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling patchAgent");
      }
      // verify the required parameter 'agent_uuid' is set
      if (agent_uuid === undefined || agent_uuid === null) {
        throw new Error("Missing the required parameter 'agent_uuid' when calling patchAgent");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling patchAgent");
      }

      let pathParams = {
        'owner': owner,
        'agent.uuid': agent_uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1Agent;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{agent.uuid}', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the patchAgentToken operation.
     * @callback module:api/AgentsV1Api~patchAgentTokenCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Token} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch agent token
     * @param {String} owner Owner of the namespace
     * @param {String} entity Rntity
     * @param {module:model/V1Token} body Token body
     * @param {module:api/AgentsV1Api~patchAgentTokenCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Token}
     */
    patchAgentToken(owner, entity, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling patchAgentToken");
      }
      // verify the required parameter 'entity' is set
      if (entity === undefined || entity === null) {
        throw new Error("Missing the required parameter 'entity' when calling patchAgentToken");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling patchAgentToken");
      }

      let pathParams = {
        'owner': owner,
        'entity': entity
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1Token;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{entity}/token', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the syncAgent operation.
     * @callback module:api/AgentsV1Api~syncAgentCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Sync agent
     * @param {String} owner Owner of the namespace
     * @param {String} agent_uuid UUID
     * @param {module:model/V1Agent} body Agent body
     * @param {module:api/AgentsV1Api~syncAgentCallback} callback The callback function, accepting three arguments: error, data, response
     */
    syncAgent(owner, agent_uuid, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling syncAgent");
      }
      // verify the required parameter 'agent_uuid' is set
      if (agent_uuid === undefined || agent_uuid === null) {
        throw new Error("Missing the required parameter 'agent_uuid' when calling syncAgent");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling syncAgent");
      }

      let pathParams = {
        'owner': owner,
        'agent.uuid': agent_uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{agent.uuid}/sync', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateAgent operation.
     * @callback module:api/AgentsV1Api~updateAgentCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Agent} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update agent
     * @param {String} owner Owner of the namespace
     * @param {String} agent_uuid UUID
     * @param {module:model/V1Agent} body Agent body
     * @param {module:api/AgentsV1Api~updateAgentCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Agent}
     */
    updateAgent(owner, agent_uuid, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling updateAgent");
      }
      // verify the required parameter 'agent_uuid' is set
      if (agent_uuid === undefined || agent_uuid === null) {
        throw new Error("Missing the required parameter 'agent_uuid' when calling updateAgent");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateAgent");
      }

      let pathParams = {
        'owner': owner,
        'agent.uuid': agent_uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1Agent;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{agent.uuid}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateAgentConfig operation.
     * @callback module:api/AgentsV1Api~updateAgentConfigCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Agent} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update agent config
     * @param {String} owner Owner of the namespace
     * @param {String} agent_uuid UUID
     * @param {module:model/V1Agent} body Agent body
     * @param {module:api/AgentsV1Api~updateAgentConfigCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Agent}
     */
    updateAgentConfig(owner, agent_uuid, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling updateAgentConfig");
      }
      // verify the required parameter 'agent_uuid' is set
      if (agent_uuid === undefined || agent_uuid === null) {
        throw new Error("Missing the required parameter 'agent_uuid' when calling updateAgentConfig");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateAgentConfig");
      }

      let pathParams = {
        'owner': owner,
        'agent.uuid': agent_uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1Agent;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{agent.uuid}/config', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateAgentToken operation.
     * @callback module:api/AgentsV1Api~updateAgentTokenCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Token} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update agent token
     * @param {String} owner Owner of the namespace
     * @param {String} entity Rntity
     * @param {module:model/V1Token} body Token body
     * @param {module:api/AgentsV1Api~updateAgentTokenCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Token}
     */
    updateAgentToken(owner, entity, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling updateAgentToken");
      }
      // verify the required parameter 'entity' is set
      if (entity === undefined || entity === null) {
        throw new Error("Missing the required parameter 'entity' when calling updateAgentToken");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateAgentToken");
      }

      let pathParams = {
        'owner': owner,
        'entity': entity
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1Token;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{entity}/token', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
