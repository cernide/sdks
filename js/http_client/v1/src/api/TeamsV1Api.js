/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.0.0-rc16
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import RuntimeError from '../model/RuntimeError';
import V1ListTeamMembersResponse from '../model/V1ListTeamMembersResponse';
import V1ListTeamsResponse from '../model/V1ListTeamsResponse';
import V1Team from '../model/V1Team';
import V1TeamMember from '../model/V1TeamMember';

/**
* TeamsV1 service.
* @module api/TeamsV1Api
* @version 2.0.0-rc16
*/
export default class TeamsV1Api {

    /**
    * Constructs a new TeamsV1Api. 
    * Polyaxon sdk
    * @alias module:api/TeamsV1Api
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the createTeam operation.
     * @callback module:api/TeamsV1Api~createTeamCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Team} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create team
     * @param {String} owner Owner of the namespace
     * @param {module:model/V1Team} body Team body
     * @param {module:api/TeamsV1Api~createTeamCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Team}
     */
    createTeam(owner, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling createTeam");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createTeam");
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
      let returnType = V1Team;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the createTeamMember operation.
     * @callback module:api/TeamsV1Api~createTeamMemberCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1TeamMember} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create team member
     * @param {String} owner Owner of the namespace
     * @param {String} team Team
     * @param {module:model/V1TeamMember} body Team body
     * @param {module:api/TeamsV1Api~createTeamMemberCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1TeamMember}
     */
    createTeamMember(owner, team, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling createTeamMember");
      }
      // verify the required parameter 'team' is set
      if (team === undefined || team === null) {
        throw new Error("Missing the required parameter 'team' when calling createTeamMember");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createTeamMember");
      }

      let pathParams = {
        'owner': owner,
        'team': team
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
      let returnType = V1TeamMember;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams/{team}/members', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteTeam operation.
     * @callback module:api/TeamsV1Api~deleteTeamCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete team
     * @param {String} owner Owner of the namespace
     * @param {String} name Component under namesapce
     * @param {module:api/TeamsV1Api~deleteTeamCallback} callback The callback function, accepting three arguments: error, data, response
     */
    deleteTeam(owner, name, callback) {
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling deleteTeam");
      }
      // verify the required parameter 'name' is set
      if (name === undefined || name === null) {
        throw new Error("Missing the required parameter 'name' when calling deleteTeam");
      }

      let pathParams = {
        'owner': owner,
        'name': name
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
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams/{name}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteTeamMember operation.
     * @callback module:api/TeamsV1Api~deleteTeamMemberCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete team member details
     * @param {String} owner Owner of the namespace
     * @param {String} team Team under namesapce
     * @param {String} user Member under team
     * @param {module:api/TeamsV1Api~deleteTeamMemberCallback} callback The callback function, accepting three arguments: error, data, response
     */
    deleteTeamMember(owner, team, user, callback) {
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling deleteTeamMember");
      }
      // verify the required parameter 'team' is set
      if (team === undefined || team === null) {
        throw new Error("Missing the required parameter 'team' when calling deleteTeamMember");
      }
      // verify the required parameter 'user' is set
      if (user === undefined || user === null) {
        throw new Error("Missing the required parameter 'user' when calling deleteTeamMember");
      }

      let pathParams = {
        'owner': owner,
        'team': team,
        'user': user
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
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams/{team}/members/{user}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getTeam operation.
     * @callback module:api/TeamsV1Api~getTeamCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Team} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get team
     * @param {String} owner Owner of the namespace
     * @param {String} name Component under namesapce
     * @param {module:api/TeamsV1Api~getTeamCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Team}
     */
    getTeam(owner, name, callback) {
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getTeam");
      }
      // verify the required parameter 'name' is set
      if (name === undefined || name === null) {
        throw new Error("Missing the required parameter 'name' when calling getTeam");
      }

      let pathParams = {
        'owner': owner,
        'name': name
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
      let returnType = V1Team;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams/{name}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getTeamMember operation.
     * @callback module:api/TeamsV1Api~getTeamMemberCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1TeamMember} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get team member details
     * @param {String} owner Owner of the namespace
     * @param {String} team Team under namesapce
     * @param {String} user Member under team
     * @param {module:api/TeamsV1Api~getTeamMemberCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1TeamMember}
     */
    getTeamMember(owner, team, user, callback) {
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getTeamMember");
      }
      // verify the required parameter 'team' is set
      if (team === undefined || team === null) {
        throw new Error("Missing the required parameter 'team' when calling getTeamMember");
      }
      // verify the required parameter 'user' is set
      if (user === undefined || user === null) {
        throw new Error("Missing the required parameter 'user' when calling getTeamMember");
      }

      let pathParams = {
        'owner': owner,
        'team': team,
        'user': user
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
      let returnType = V1TeamMember;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams/{team}/members/{user}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listTeamMembers operation.
     * @callback module:api/TeamsV1Api~listTeamMembersCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListTeamMembersResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get team members
     * @param {String} owner Owner of the namespace
     * @param {String} name Entity managing the resource
     * @param {Object} opts Optional parameters
     * @param {Number} [offset] Pagination offset.
     * @param {Number} [limit] Limit size.
     * @param {String} [sort] Sort to order the search.
     * @param {String} [query] Query filter the search.
     * @param {Boolean} [bookmarks] Filter by bookmarks.
     * @param {String} [mode] Mode of the search.
     * @param {Boolean} [no_page] No pagination.
     * @param {module:api/TeamsV1Api~listTeamMembersCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListTeamMembersResponse}
     */
    listTeamMembers(owner, name, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listTeamMembers");
      }
      // verify the required parameter 'name' is set
      if (name === undefined || name === null) {
        throw new Error("Missing the required parameter 'name' when calling listTeamMembers");
      }

      let pathParams = {
        'owner': owner,
        'name': name
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
      let returnType = V1ListTeamMembersResponse;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams/{name}/members', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listTeamNames operation.
     * @callback module:api/TeamsV1Api~listTeamNamesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListTeamsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List teams names
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} [offset] Pagination offset.
     * @param {Number} [limit] Limit size.
     * @param {String} [sort] Sort to order the search.
     * @param {String} [query] Query filter the search.
     * @param {Boolean} [bookmarks] Filter by bookmarks.
     * @param {String} [mode] Mode of the search.
     * @param {Boolean} [no_page] No pagination.
     * @param {module:api/TeamsV1Api~listTeamNamesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListTeamsResponse}
     */
    listTeamNames(owner, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listTeamNames");
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
      let returnType = V1ListTeamsResponse;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams/names', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listTeams operation.
     * @callback module:api/TeamsV1Api~listTeamsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListTeamsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List teams
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} [offset] Pagination offset.
     * @param {Number} [limit] Limit size.
     * @param {String} [sort] Sort to order the search.
     * @param {String} [query] Query filter the search.
     * @param {Boolean} [bookmarks] Filter by bookmarks.
     * @param {String} [mode] Mode of the search.
     * @param {Boolean} [no_page] No pagination.
     * @param {module:api/TeamsV1Api~listTeamsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListTeamsResponse}
     */
    listTeams(owner, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listTeams");
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
      let returnType = V1ListTeamsResponse;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the patchTeam operation.
     * @callback module:api/TeamsV1Api~patchTeamCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Team} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch team
     * @param {String} owner Owner of the namespace
     * @param {String} team_name Name
     * @param {module:model/V1Team} body Team body
     * @param {module:api/TeamsV1Api~patchTeamCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Team}
     */
    patchTeam(owner, team_name, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling patchTeam");
      }
      // verify the required parameter 'team_name' is set
      if (team_name === undefined || team_name === null) {
        throw new Error("Missing the required parameter 'team_name' when calling patchTeam");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling patchTeam");
      }

      let pathParams = {
        'owner': owner,
        'team.name': team_name
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
      let returnType = V1Team;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams/{team.name}', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the patchTeamMember operation.
     * @callback module:api/TeamsV1Api~patchTeamMemberCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1TeamMember} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch team member
     * @param {String} owner Owner of the namespace
     * @param {String} team Team
     * @param {String} member_user User
     * @param {module:model/V1TeamMember} body Team body
     * @param {module:api/TeamsV1Api~patchTeamMemberCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1TeamMember}
     */
    patchTeamMember(owner, team, member_user, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling patchTeamMember");
      }
      // verify the required parameter 'team' is set
      if (team === undefined || team === null) {
        throw new Error("Missing the required parameter 'team' when calling patchTeamMember");
      }
      // verify the required parameter 'member_user' is set
      if (member_user === undefined || member_user === null) {
        throw new Error("Missing the required parameter 'member_user' when calling patchTeamMember");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling patchTeamMember");
      }

      let pathParams = {
        'owner': owner,
        'team': team,
        'member.user': member_user
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
      let returnType = V1TeamMember;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams/{team}/members/{member.user}', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateTeam operation.
     * @callback module:api/TeamsV1Api~updateTeamCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Team} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update team
     * @param {String} owner Owner of the namespace
     * @param {String} team_name Name
     * @param {module:model/V1Team} body Team body
     * @param {module:api/TeamsV1Api~updateTeamCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Team}
     */
    updateTeam(owner, team_name, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling updateTeam");
      }
      // verify the required parameter 'team_name' is set
      if (team_name === undefined || team_name === null) {
        throw new Error("Missing the required parameter 'team_name' when calling updateTeam");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateTeam");
      }

      let pathParams = {
        'owner': owner,
        'team.name': team_name
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
      let returnType = V1Team;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams/{team.name}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateTeamMember operation.
     * @callback module:api/TeamsV1Api~updateTeamMemberCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1TeamMember} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update team member
     * @param {String} owner Owner of the namespace
     * @param {String} team Team
     * @param {String} member_user User
     * @param {module:model/V1TeamMember} body Team body
     * @param {module:api/TeamsV1Api~updateTeamMemberCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1TeamMember}
     */
    updateTeamMember(owner, team, member_user, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling updateTeamMember");
      }
      // verify the required parameter 'team' is set
      if (team === undefined || team === null) {
        throw new Error("Missing the required parameter 'team' when calling updateTeamMember");
      }
      // verify the required parameter 'member_user' is set
      if (member_user === undefined || member_user === null) {
        throw new Error("Missing the required parameter 'member_user' when calling updateTeamMember");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateTeamMember");
      }

      let pathParams = {
        'owner': owner,
        'team': team,
        'member.user': member_user
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
      let returnType = V1TeamMember;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/teams/{team}/members/{member.user}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
