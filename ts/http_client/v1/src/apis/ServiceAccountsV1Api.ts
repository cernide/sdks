/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.2.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  RuntimeError,
  V1ListServiceAccountsResponse,
  V1ListTokenResponse,
  V1ServiceAccount,
  V1Token,
} from '../models';
import {
    RuntimeErrorFromJSON,
    RuntimeErrorToJSON,
    V1ListServiceAccountsResponseFromJSON,
    V1ListServiceAccountsResponseToJSON,
    V1ListTokenResponseFromJSON,
    V1ListTokenResponseToJSON,
    V1ServiceAccountFromJSON,
    V1ServiceAccountToJSON,
    V1TokenFromJSON,
    V1TokenToJSON,
} from '../models';

export interface CreateServiceAccountRequest {
    owner: string;
    body: V1ServiceAccount;
}

export interface CreateServiceAccountTokenRequest {
    owner: string;
    entity: string;
    body: V1Token;
}

export interface DeleteServiceAccountRequest {
    owner: string;
    uuid: string;
}

export interface DeleteServiceAccountTokenRequest {
    owner: string;
    entity: string;
    uuid: string;
}

export interface GetServiceAccountRequest {
    owner: string;
    uuid: string;
}

export interface GetServiceAccountTokenRequest {
    owner: string;
    entity: string;
    uuid: string;
}

export interface ListServiceAccountNamesRequest {
    owner: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    bookmarks?: boolean;
    mode?: string;
    noPage?: boolean;
}

export interface ListServiceAccountTokensRequest {
    owner: string;
    uuid: string;
    entity?: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    noPage?: boolean;
}

export interface ListServiceAccountsRequest {
    owner: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    bookmarks?: boolean;
    mode?: string;
    noPage?: boolean;
}

export interface PatchServiceAccountRequest {
    owner: string;
    saUuid: string;
    body: V1ServiceAccount;
}

export interface PatchServiceAccountTokenRequest {
    owner: string;
    entity: string;
    tokenUuid: string;
    body: V1Token;
}

export interface UpdateServiceAccountRequest {
    owner: string;
    saUuid: string;
    body: V1ServiceAccount;
}

export interface UpdateServiceAccountTokenRequest {
    owner: string;
    entity: string;
    tokenUuid: string;
    body: V1Token;
}

/**
 *
 */
export class ServiceAccountsV1Api extends runtime.BaseAPI {

    /**
     * Create service account
     */
    async createServiceAccountRaw(requestParameters: CreateServiceAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ServiceAccount>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling createServiceAccount.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createServiceAccount.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/sa`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V1ServiceAccountToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ServiceAccountFromJSON(jsonValue));
    }

    /**
     * Create service account
     */
    async createServiceAccount(requestParameters: CreateServiceAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ServiceAccount> {
        const response = await this.createServiceAccountRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create service account token
     */
    async createServiceAccountTokenRaw(requestParameters: CreateServiceAccountTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Token>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling createServiceAccountToken.');
        }

        if (requestParameters.entity === null || requestParameters.entity === undefined) {
            throw new runtime.RequiredError('entity','Required parameter requestParameters.entity was null or undefined when calling createServiceAccountToken.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createServiceAccountToken.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/sa/{entity}/tokens`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"entity"}}`, encodeURIComponent(String(requestParameters.entity))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V1TokenToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TokenFromJSON(jsonValue));
    }

    /**
     * Create service account token
     */
    async createServiceAccountToken(requestParameters: CreateServiceAccountTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Token> {
        const response = await this.createServiceAccountTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete service account
     */
    async deleteServiceAccountRaw(requestParameters: DeleteServiceAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling deleteServiceAccount.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling deleteServiceAccount.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/sa/{uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete service account
     */
    async deleteServiceAccount(requestParameters: DeleteServiceAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteServiceAccountRaw(requestParameters, initOverrides);
    }

    /**
     * Delete service account token
     */
    async deleteServiceAccountTokenRaw(requestParameters: DeleteServiceAccountTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling deleteServiceAccountToken.');
        }

        if (requestParameters.entity === null || requestParameters.entity === undefined) {
            throw new runtime.RequiredError('entity','Required parameter requestParameters.entity was null or undefined when calling deleteServiceAccountToken.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling deleteServiceAccountToken.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/sa/{entity}/tokens/{uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"entity"}}`, encodeURIComponent(String(requestParameters.entity))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete service account token
     */
    async deleteServiceAccountToken(requestParameters: DeleteServiceAccountTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteServiceAccountTokenRaw(requestParameters, initOverrides);
    }

    /**
     * Get service account
     */
    async getServiceAccountRaw(requestParameters: GetServiceAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ServiceAccount>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling getServiceAccount.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling getServiceAccount.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/sa/{uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ServiceAccountFromJSON(jsonValue));
    }

    /**
     * Get service account
     */
    async getServiceAccount(requestParameters: GetServiceAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ServiceAccount> {
        const response = await this.getServiceAccountRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get service account token
     */
    async getServiceAccountTokenRaw(requestParameters: GetServiceAccountTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Token>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling getServiceAccountToken.');
        }

        if (requestParameters.entity === null || requestParameters.entity === undefined) {
            throw new runtime.RequiredError('entity','Required parameter requestParameters.entity was null or undefined when calling getServiceAccountToken.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling getServiceAccountToken.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/sa/{entity}/tokens/{uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"entity"}}`, encodeURIComponent(String(requestParameters.entity))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TokenFromJSON(jsonValue));
    }

    /**
     * Get service account token
     */
    async getServiceAccountToken(requestParameters: GetServiceAccountTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Token> {
        const response = await this.getServiceAccountTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List service accounts names
     */
    async listServiceAccountNamesRaw(requestParameters: ListServiceAccountNamesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ListServiceAccountsResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listServiceAccountNames.');
        }

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

        if (requestParameters.bookmarks !== undefined) {
            queryParameters['bookmarks'] = requestParameters.bookmarks;
        }

        if (requestParameters.mode !== undefined) {
            queryParameters['mode'] = requestParameters.mode;
        }

        if (requestParameters.noPage !== undefined) {
            queryParameters['no_page'] = requestParameters.noPage;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/sa/names`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListServiceAccountsResponseFromJSON(jsonValue));
    }

    /**
     * List service accounts names
     */
    async listServiceAccountNames(requestParameters: ListServiceAccountNamesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ListServiceAccountsResponse> {
        const response = await this.listServiceAccountNamesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List service account tokens
     */
    async listServiceAccountTokensRaw(requestParameters: ListServiceAccountTokensRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ListTokenResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listServiceAccountTokens.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling listServiceAccountTokens.');
        }

        const queryParameters: any = {};

        if (requestParameters.entity !== undefined) {
            queryParameters['entity'] = requestParameters.entity;
        }

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
            path: `/api/v1/orgs/{owner}/sa/{uuid}/tokens`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListTokenResponseFromJSON(jsonValue));
    }

    /**
     * List service account tokens
     */
    async listServiceAccountTokens(requestParameters: ListServiceAccountTokensRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ListTokenResponse> {
        const response = await this.listServiceAccountTokensRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List service accounts
     */
    async listServiceAccountsRaw(requestParameters: ListServiceAccountsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ListServiceAccountsResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listServiceAccounts.');
        }

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

        if (requestParameters.bookmarks !== undefined) {
            queryParameters['bookmarks'] = requestParameters.bookmarks;
        }

        if (requestParameters.mode !== undefined) {
            queryParameters['mode'] = requestParameters.mode;
        }

        if (requestParameters.noPage !== undefined) {
            queryParameters['no_page'] = requestParameters.noPage;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/sa`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListServiceAccountsResponseFromJSON(jsonValue));
    }

    /**
     * List service accounts
     */
    async listServiceAccounts(requestParameters: ListServiceAccountsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ListServiceAccountsResponse> {
        const response = await this.listServiceAccountsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Patch service account
     */
    async patchServiceAccountRaw(requestParameters: PatchServiceAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ServiceAccount>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling patchServiceAccount.');
        }

        if (requestParameters.saUuid === null || requestParameters.saUuid === undefined) {
            throw new runtime.RequiredError('saUuid','Required parameter requestParameters.saUuid was null or undefined when calling patchServiceAccount.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling patchServiceAccount.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/sa/{sa.uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"sa.uuid"}}`, encodeURIComponent(String(requestParameters.saUuid))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V1ServiceAccountToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ServiceAccountFromJSON(jsonValue));
    }

    /**
     * Patch service account
     */
    async patchServiceAccount(requestParameters: PatchServiceAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ServiceAccount> {
        const response = await this.patchServiceAccountRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Patch service account token
     */
    async patchServiceAccountTokenRaw(requestParameters: PatchServiceAccountTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Token>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling patchServiceAccountToken.');
        }

        if (requestParameters.entity === null || requestParameters.entity === undefined) {
            throw new runtime.RequiredError('entity','Required parameter requestParameters.entity was null or undefined when calling patchServiceAccountToken.');
        }

        if (requestParameters.tokenUuid === null || requestParameters.tokenUuid === undefined) {
            throw new runtime.RequiredError('tokenUuid','Required parameter requestParameters.tokenUuid was null or undefined when calling patchServiceAccountToken.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling patchServiceAccountToken.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/sa/{entity}/tokens/{token.uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"entity"}}`, encodeURIComponent(String(requestParameters.entity))).replace(`{${"token.uuid"}}`, encodeURIComponent(String(requestParameters.tokenUuid))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V1TokenToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TokenFromJSON(jsonValue));
    }

    /**
     * Patch service account token
     */
    async patchServiceAccountToken(requestParameters: PatchServiceAccountTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Token> {
        const response = await this.patchServiceAccountTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Update service account
     */
    async updateServiceAccountRaw(requestParameters: UpdateServiceAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ServiceAccount>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling updateServiceAccount.');
        }

        if (requestParameters.saUuid === null || requestParameters.saUuid === undefined) {
            throw new runtime.RequiredError('saUuid','Required parameter requestParameters.saUuid was null or undefined when calling updateServiceAccount.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateServiceAccount.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/sa/{sa.uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"sa.uuid"}}`, encodeURIComponent(String(requestParameters.saUuid))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V1ServiceAccountToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ServiceAccountFromJSON(jsonValue));
    }

    /**
     * Update service account
     */
    async updateServiceAccount(requestParameters: UpdateServiceAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ServiceAccount> {
        const response = await this.updateServiceAccountRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Update service account token
     */
    async updateServiceAccountTokenRaw(requestParameters: UpdateServiceAccountTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Token>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling updateServiceAccountToken.');
        }

        if (requestParameters.entity === null || requestParameters.entity === undefined) {
            throw new runtime.RequiredError('entity','Required parameter requestParameters.entity was null or undefined when calling updateServiceAccountToken.');
        }

        if (requestParameters.tokenUuid === null || requestParameters.tokenUuid === undefined) {
            throw new runtime.RequiredError('tokenUuid','Required parameter requestParameters.tokenUuid was null or undefined when calling updateServiceAccountToken.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateServiceAccountToken.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/sa/{entity}/tokens/{token.uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"entity"}}`, encodeURIComponent(String(requestParameters.entity))).replace(`{${"token.uuid"}}`, encodeURIComponent(String(requestParameters.tokenUuid))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V1TokenToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TokenFromJSON(jsonValue));
    }

    /**
     * Update service account token
     */
    async updateServiceAccountToken(requestParameters: UpdateServiceAccountTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Token> {
        const response = await this.updateServiceAccountTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
