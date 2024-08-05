/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.3.2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  RuntimeError,
  V1ListSearchesResponse,
  V1Search,
} from '../models';
import {
    RuntimeErrorFromJSON,
    RuntimeErrorToJSON,
    V1ListSearchesResponseFromJSON,
    V1ListSearchesResponseToJSON,
    V1SearchFromJSON,
    V1SearchToJSON,
} from '../models';

export interface CreateProjectSearchRequest {
    owner: string;
    project: string;
    body: V1Search;
}

export interface DeleteProjectSearchRequest {
    owner: string;
    entity: string;
    uuid: string;
}

export interface GetProjectSearchRequest {
    owner: string;
    entity: string;
    uuid: string;
}

export interface ListProjectSearchNamesRequest {
    owner: string;
    name: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    bookmarks?: boolean;
    mode?: string;
    noPage?: boolean;
}

export interface ListProjectSearchesRequest {
    owner: string;
    name: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    bookmarks?: boolean;
    mode?: string;
    noPage?: boolean;
}

export interface PatchProjectSearchRequest {
    owner: string;
    project: string;
    searchUuid: string;
    body: V1Search;
}

export interface PromoteProjectSearchRequest {
    owner: string;
    entity: string;
    uuid: string;
}

export interface UpdateProjectSearchRequest {
    owner: string;
    project: string;
    searchUuid: string;
    body: V1Search;
}

/**
 *
 */
export class ProjectSearchesV1Api extends runtime.BaseAPI {

    /**
     * Create project search
     */
    async createProjectSearchRaw(requestParameters: CreateProjectSearchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Search>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling createProjectSearch.');
        }

        if (requestParameters.project === null || requestParameters.project === undefined) {
            throw new runtime.RequiredError('project','Required parameter requestParameters.project was null or undefined when calling createProjectSearch.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createProjectSearch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/{owner}/{project}/searches`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"project"}}`, encodeURIComponent(String(requestParameters.project))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V1SearchToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1SearchFromJSON(jsonValue));
    }

    /**
     * Create project search
     */
    async createProjectSearch(requestParameters: CreateProjectSearchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Search> {
        const response = await this.createProjectSearchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete project search
     */
    async deleteProjectSearchRaw(requestParameters: DeleteProjectSearchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling deleteProjectSearch.');
        }

        if (requestParameters.entity === null || requestParameters.entity === undefined) {
            throw new runtime.RequiredError('entity','Required parameter requestParameters.entity was null or undefined when calling deleteProjectSearch.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling deleteProjectSearch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/{owner}/{entity}/searches/{uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"entity"}}`, encodeURIComponent(String(requestParameters.entity))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete project search
     */
    async deleteProjectSearch(requestParameters: DeleteProjectSearchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteProjectSearchRaw(requestParameters, initOverrides);
    }

    /**
     * Get project search
     */
    async getProjectSearchRaw(requestParameters: GetProjectSearchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Search>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling getProjectSearch.');
        }

        if (requestParameters.entity === null || requestParameters.entity === undefined) {
            throw new runtime.RequiredError('entity','Required parameter requestParameters.entity was null or undefined when calling getProjectSearch.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling getProjectSearch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/{owner}/{entity}/searches/{uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"entity"}}`, encodeURIComponent(String(requestParameters.entity))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1SearchFromJSON(jsonValue));
    }

    /**
     * Get project search
     */
    async getProjectSearch(requestParameters: GetProjectSearchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Search> {
        const response = await this.getProjectSearchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List project search names
     */
    async listProjectSearchNamesRaw(requestParameters: ListProjectSearchNamesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ListSearchesResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listProjectSearchNames.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling listProjectSearchNames.');
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
            path: `/api/v1/{owner}/{name}/searches/names`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListSearchesResponseFromJSON(jsonValue));
    }

    /**
     * List project search names
     */
    async listProjectSearchNames(requestParameters: ListProjectSearchNamesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ListSearchesResponse> {
        const response = await this.listProjectSearchNamesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List project searches
     */
    async listProjectSearchesRaw(requestParameters: ListProjectSearchesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ListSearchesResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listProjectSearches.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling listProjectSearches.');
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
            path: `/api/v1/{owner}/{name}/searches`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListSearchesResponseFromJSON(jsonValue));
    }

    /**
     * List project searches
     */
    async listProjectSearches(requestParameters: ListProjectSearchesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ListSearchesResponse> {
        const response = await this.listProjectSearchesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Patch project search
     */
    async patchProjectSearchRaw(requestParameters: PatchProjectSearchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Search>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling patchProjectSearch.');
        }

        if (requestParameters.project === null || requestParameters.project === undefined) {
            throw new runtime.RequiredError('project','Required parameter requestParameters.project was null or undefined when calling patchProjectSearch.');
        }

        if (requestParameters.searchUuid === null || requestParameters.searchUuid === undefined) {
            throw new runtime.RequiredError('searchUuid','Required parameter requestParameters.searchUuid was null or undefined when calling patchProjectSearch.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling patchProjectSearch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/{owner}/{project}/searches/{search.uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"project"}}`, encodeURIComponent(String(requestParameters.project))).replace(`{${"search.uuid"}}`, encodeURIComponent(String(requestParameters.searchUuid))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V1SearchToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1SearchFromJSON(jsonValue));
    }

    /**
     * Patch project search
     */
    async patchProjectSearch(requestParameters: PatchProjectSearchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Search> {
        const response = await this.patchProjectSearchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Promote project search
     */
    async promoteProjectSearchRaw(requestParameters: PromoteProjectSearchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling promoteProjectSearch.');
        }

        if (requestParameters.entity === null || requestParameters.entity === undefined) {
            throw new runtime.RequiredError('entity','Required parameter requestParameters.entity was null or undefined when calling promoteProjectSearch.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling promoteProjectSearch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/{owner}/{entity}/searches/{uuid}/promote`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"entity"}}`, encodeURIComponent(String(requestParameters.entity))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Promote project search
     */
    async promoteProjectSearch(requestParameters: PromoteProjectSearchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.promoteProjectSearchRaw(requestParameters, initOverrides);
    }

    /**
     * Update project search
     */
    async updateProjectSearchRaw(requestParameters: UpdateProjectSearchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Search>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling updateProjectSearch.');
        }

        if (requestParameters.project === null || requestParameters.project === undefined) {
            throw new runtime.RequiredError('project','Required parameter requestParameters.project was null or undefined when calling updateProjectSearch.');
        }

        if (requestParameters.searchUuid === null || requestParameters.searchUuid === undefined) {
            throw new runtime.RequiredError('searchUuid','Required parameter requestParameters.searchUuid was null or undefined when calling updateProjectSearch.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateProjectSearch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/{owner}/{project}/searches/{search.uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"project"}}`, encodeURIComponent(String(requestParameters.project))).replace(`{${"search.uuid"}}`, encodeURIComponent(String(requestParameters.searchUuid))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V1SearchToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1SearchFromJSON(jsonValue));
    }

    /**
     * Update project search
     */
    async updateProjectSearch(requestParameters: UpdateProjectSearchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Search> {
        const response = await this.updateProjectSearchRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
