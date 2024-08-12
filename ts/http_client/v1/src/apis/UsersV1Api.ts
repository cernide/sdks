/* tslint:disable */
/* eslint-disable */
/**
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


import * as runtime from '../runtime';
import type {
  RuntimeError,
  V1ListActivitiesResponse,
  V1ListTokenResponse,
  V1Token,
  V1User,
} from '../models';
import {
    RuntimeErrorFromJSON,
    RuntimeErrorToJSON,
    V1ListActivitiesResponseFromJSON,
    V1ListActivitiesResponseToJSON,
    V1ListTokenResponseFromJSON,
    V1ListTokenResponseToJSON,
    V1TokenFromJSON,
    V1TokenToJSON,
    V1UserFromJSON,
    V1UserToJSON,
} from '../models';

export interface CreateTokenRequest {
    body: V1Token;
}

export interface DeleteTokenRequest {
    uuid: string;
}

export interface GetHistoryRequest {
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    noPage?: boolean;
}

export interface GetSuggestionsRequest {
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    noPage?: boolean;
}

export interface GetTokenRequest {
    uuid: string;
}

export interface GetWorkspacesRequest {
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    noPage?: boolean;
}

export interface ListTokensRequest {
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    noPage?: boolean;
}

export interface PatchTokenRequest {
    tokenUuid: string;
    body: V1Token;
}

export interface PatchUserRequest {
    body: V1User;
}

export interface UpdateTokenRequest {
    tokenUuid: string;
    body: V1Token;
}

export interface UpdateUserRequest {
    body: V1User;
}

/**
 *
 */
export class UsersV1Api extends runtime.BaseAPI {

    /**
     * Create token
     */
    async createTokenRaw(requestParameters: CreateTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Token>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createToken.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/users/tokens`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V1TokenToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TokenFromJSON(jsonValue));
    }

    /**
     * Create token
     */
    async createToken(requestParameters: CreateTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Token> {
        const response = await this.createTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete token
     */
    async deleteTokenRaw(requestParameters: DeleteTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling deleteToken.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/users/tokens/{uuid}`.replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete token
     */
    async deleteToken(requestParameters: DeleteTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteTokenRaw(requestParameters, initOverrides);
    }

    /**
     * User History
     */
    async getHistoryRaw(requestParameters: GetHistoryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ListActivitiesResponse>> {
        const queryParameters: any = {};

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.sort !== undefined) {
            queryParameters['sort'] = requestParameters.sort;
        }

        if (requestParameters.query !== undefined) {
            queryParameters['query'] = requestParameters.query;
        }

        if (requestParameters.noPage !== undefined) {
            queryParameters['no_page'] = requestParameters.noPage;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/users/history`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListActivitiesResponseFromJSON(jsonValue));
    }

    /**
     * User History
     */
    async getHistory(requestParameters: GetHistoryRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ListActivitiesResponse> {
        const response = await this.getHistoryRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * User suggestions
     */
    async getSuggestionsRaw(requestParameters: GetSuggestionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        const queryParameters: any = {};

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.sort !== undefined) {
            queryParameters['sort'] = requestParameters.sort;
        }

        if (requestParameters.query !== undefined) {
            queryParameters['query'] = requestParameters.query;
        }

        if (requestParameters.noPage !== undefined) {
            queryParameters['no_page'] = requestParameters.noPage;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/users/suggestions`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * User suggestions
     */
    async getSuggestions(requestParameters: GetSuggestionsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.getSuggestionsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get token
     */
    async getTokenRaw(requestParameters: GetTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Token>> {
        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling getToken.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/users/tokens/{uuid}`.replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TokenFromJSON(jsonValue));
    }

    /**
     * Get token
     */
    async getToken(requestParameters: GetTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Token> {
        const response = await this.getTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get current user
     */
    async getUserRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1User>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/users`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1UserFromJSON(jsonValue));
    }

    /**
     * Get current user
     */
    async getUser(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1User> {
        const response = await this.getUserRaw(initOverrides);
        return await response.value();
    }

    /**
     * User workspaces
     */
    async getWorkspacesRaw(requestParameters: GetWorkspacesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        const queryParameters: any = {};

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.sort !== undefined) {
            queryParameters['sort'] = requestParameters.sort;
        }

        if (requestParameters.query !== undefined) {
            queryParameters['query'] = requestParameters.query;
        }

        if (requestParameters.noPage !== undefined) {
            queryParameters['no_page'] = requestParameters.noPage;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/users/workspaces`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * User workspaces
     */
    async getWorkspaces(requestParameters: GetWorkspacesRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.getWorkspacesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List tokens
     */
    async listTokensRaw(requestParameters: ListTokensRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ListTokenResponse>> {
        const queryParameters: any = {};

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.sort !== undefined) {
            queryParameters['sort'] = requestParameters.sort;
        }

        if (requestParameters.query !== undefined) {
            queryParameters['query'] = requestParameters.query;
        }

        if (requestParameters.noPage !== undefined) {
            queryParameters['no_page'] = requestParameters.noPage;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/users/tokens`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListTokenResponseFromJSON(jsonValue));
    }

    /**
     * List tokens
     */
    async listTokens(requestParameters: ListTokensRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ListTokenResponse> {
        const response = await this.listTokensRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Patch token
     */
    async patchTokenRaw(requestParameters: PatchTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Token>> {
        if (requestParameters.tokenUuid === null || requestParameters.tokenUuid === undefined) {
            throw new runtime.RequiredError('tokenUuid','Required parameter requestParameters.tokenUuid was null or undefined when calling patchToken.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling patchToken.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/users/tokens/{token.uuid}`.replace(`{${"token.uuid"}}`, encodeURIComponent(String(requestParameters.tokenUuid))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V1TokenToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TokenFromJSON(jsonValue));
    }

    /**
     * Patch token
     */
    async patchToken(requestParameters: PatchTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Token> {
        const response = await this.patchTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Patch current user
     */
    async patchUserRaw(requestParameters: PatchUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1User>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling patchUser.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/users`,
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V1UserToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1UserFromJSON(jsonValue));
    }

    /**
     * Patch current user
     */
    async patchUser(requestParameters: PatchUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1User> {
        const response = await this.patchUserRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Update token
     */
    async updateTokenRaw(requestParameters: UpdateTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Token>> {
        if (requestParameters.tokenUuid === null || requestParameters.tokenUuid === undefined) {
            throw new runtime.RequiredError('tokenUuid','Required parameter requestParameters.tokenUuid was null or undefined when calling updateToken.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateToken.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/users/tokens/{token.uuid}`.replace(`{${"token.uuid"}}`, encodeURIComponent(String(requestParameters.tokenUuid))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V1TokenToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TokenFromJSON(jsonValue));
    }

    /**
     * Update token
     */
    async updateToken(requestParameters: UpdateTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Token> {
        const response = await this.updateTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Update current user
     */
    async updateUserRaw(requestParameters: UpdateUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1User>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateUser.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/users`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V1UserToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1UserFromJSON(jsonValue));
    }

    /**
     * Update current user
     */
    async updateUser(requestParameters: UpdateUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1User> {
        const response = await this.updateUserRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
