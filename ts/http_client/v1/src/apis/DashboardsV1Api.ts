/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc28
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  RuntimeError,
  V1Dashboard,
  V1ListDashboardsResponse,
} from '../models';
import {
    RuntimeErrorFromJSON,
    RuntimeErrorToJSON,
    V1DashboardFromJSON,
    V1DashboardToJSON,
    V1ListDashboardsResponseFromJSON,
    V1ListDashboardsResponseToJSON,
} from '../models';

export interface CreateDashboardRequest {
    owner: string;
    body: V1Dashboard;
}

export interface DeleteDashboardRequest {
    owner: string;
    uuid: string;
}

export interface GetDashboardRequest {
    owner: string;
    uuid: string;
}

export interface ListDashboardNamesRequest {
    owner: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    bookmarks?: boolean;
    mode?: string;
    noPage?: boolean;
}

export interface ListDashboardsRequest {
    owner: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    bookmarks?: boolean;
    mode?: string;
    noPage?: boolean;
}

export interface PatchDashboardRequest {
    owner: string;
    dashboardUuid: string;
    body: V1Dashboard;
}

export interface UpdateDashboardRequest {
    owner: string;
    dashboardUuid: string;
    body: V1Dashboard;
}

/**
 *
 */
export class DashboardsV1Api extends runtime.BaseAPI {

    /**
     * Create dashboard
     */
    async createDashboardRaw(requestParameters: CreateDashboardRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Dashboard>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling createDashboard.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createDashboard.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/dashboards`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V1DashboardToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1DashboardFromJSON(jsonValue));
    }

    /**
     * Create dashboard
     */
    async createDashboard(requestParameters: CreateDashboardRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Dashboard> {
        const response = await this.createDashboardRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete dashboard
     */
    async deleteDashboardRaw(requestParameters: DeleteDashboardRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling deleteDashboard.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling deleteDashboard.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/dashboards/{uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete dashboard
     */
    async deleteDashboard(requestParameters: DeleteDashboardRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteDashboardRaw(requestParameters, initOverrides);
    }

    /**
     * Get dashboard
     */
    async getDashboardRaw(requestParameters: GetDashboardRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Dashboard>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling getDashboard.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling getDashboard.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/dashboards/{uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1DashboardFromJSON(jsonValue));
    }

    /**
     * Get dashboard
     */
    async getDashboard(requestParameters: GetDashboardRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Dashboard> {
        const response = await this.getDashboardRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List dashboard names
     */
    async listDashboardNamesRaw(requestParameters: ListDashboardNamesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ListDashboardsResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listDashboardNames.');
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
            path: `/api/v1/orgs/{owner}/dashboards/names`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListDashboardsResponseFromJSON(jsonValue));
    }

    /**
     * List dashboard names
     */
    async listDashboardNames(requestParameters: ListDashboardNamesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ListDashboardsResponse> {
        const response = await this.listDashboardNamesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List dashboards
     */
    async listDashboardsRaw(requestParameters: ListDashboardsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ListDashboardsResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listDashboards.');
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
            path: `/api/v1/orgs/{owner}/dashboards`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListDashboardsResponseFromJSON(jsonValue));
    }

    /**
     * List dashboards
     */
    async listDashboards(requestParameters: ListDashboardsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ListDashboardsResponse> {
        const response = await this.listDashboardsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Patch dashboard
     */
    async patchDashboardRaw(requestParameters: PatchDashboardRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Dashboard>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling patchDashboard.');
        }

        if (requestParameters.dashboardUuid === null || requestParameters.dashboardUuid === undefined) {
            throw new runtime.RequiredError('dashboardUuid','Required parameter requestParameters.dashboardUuid was null or undefined when calling patchDashboard.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling patchDashboard.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/dashboards/{dashboard.uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"dashboard.uuid"}}`, encodeURIComponent(String(requestParameters.dashboardUuid))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V1DashboardToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1DashboardFromJSON(jsonValue));
    }

    /**
     * Patch dashboard
     */
    async patchDashboard(requestParameters: PatchDashboardRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Dashboard> {
        const response = await this.patchDashboardRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Update dashboard
     */
    async updateDashboardRaw(requestParameters: UpdateDashboardRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Dashboard>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling updateDashboard.');
        }

        if (requestParameters.dashboardUuid === null || requestParameters.dashboardUuid === undefined) {
            throw new runtime.RequiredError('dashboardUuid','Required parameter requestParameters.dashboardUuid was null or undefined when calling updateDashboard.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateDashboard.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/dashboards/{dashboard.uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"dashboard.uuid"}}`, encodeURIComponent(String(requestParameters.dashboardUuid))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V1DashboardToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1DashboardFromJSON(jsonValue));
    }

    /**
     * Update dashboard
     */
    async updateDashboard(requestParameters: UpdateDashboardRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Dashboard> {
        const response = await this.updateDashboardRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
