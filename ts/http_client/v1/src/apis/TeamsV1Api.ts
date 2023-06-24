/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc22
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  RuntimeError,
  V1ListTeamMembersResponse,
  V1ListTeamsResponse,
  V1Team,
  V1TeamMember,
} from '../models';
import {
    RuntimeErrorFromJSON,
    RuntimeErrorToJSON,
    V1ListTeamMembersResponseFromJSON,
    V1ListTeamMembersResponseToJSON,
    V1ListTeamsResponseFromJSON,
    V1ListTeamsResponseToJSON,
    V1TeamFromJSON,
    V1TeamToJSON,
    V1TeamMemberFromJSON,
    V1TeamMemberToJSON,
} from '../models';

export interface CreateTeamRequest {
    owner: string;
    body: V1Team;
}

export interface CreateTeamMemberRequest {
    owner: string;
    team: string;
    body: V1TeamMember;
}

export interface DeleteTeamRequest {
    owner: string;
    name: string;
}

export interface DeleteTeamMemberRequest {
    owner: string;
    team: string;
    user: string;
}

export interface GetTeamRequest {
    owner: string;
    name: string;
}

export interface GetTeamMemberRequest {
    owner: string;
    team: string;
    user: string;
}

export interface ListTeamMembersRequest {
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

export interface ListTeamNamesRequest {
    owner: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    bookmarks?: boolean;
    mode?: string;
    noPage?: boolean;
}

export interface ListTeamsRequest {
    owner: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
    bookmarks?: boolean;
    mode?: string;
    noPage?: boolean;
}

export interface PatchTeamRequest {
    owner: string;
    teamName: string;
    body: V1Team;
}

export interface PatchTeamMemberRequest {
    owner: string;
    team: string;
    memberUser: string;
    body: V1TeamMember;
}

export interface UpdateTeamRequest {
    owner: string;
    teamName: string;
    body: V1Team;
}

export interface UpdateTeamMemberRequest {
    owner: string;
    team: string;
    memberUser: string;
    body: V1TeamMember;
}

/**
 *
 */
export class TeamsV1Api extends runtime.BaseAPI {

    /**
     * Create team
     */
    async createTeamRaw(requestParameters: CreateTeamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Team>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling createTeam.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createTeam.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/teams`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V1TeamToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TeamFromJSON(jsonValue));
    }

    /**
     * Create team
     */
    async createTeam(requestParameters: CreateTeamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Team> {
        const response = await this.createTeamRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create team member
     */
    async createTeamMemberRaw(requestParameters: CreateTeamMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1TeamMember>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling createTeamMember.');
        }

        if (requestParameters.team === null || requestParameters.team === undefined) {
            throw new runtime.RequiredError('team','Required parameter requestParameters.team was null or undefined when calling createTeamMember.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createTeamMember.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/teams/{team}/members`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"team"}}`, encodeURIComponent(String(requestParameters.team))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V1TeamMemberToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TeamMemberFromJSON(jsonValue));
    }

    /**
     * Create team member
     */
    async createTeamMember(requestParameters: CreateTeamMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1TeamMember> {
        const response = await this.createTeamMemberRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete team
     */
    async deleteTeamRaw(requestParameters: DeleteTeamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling deleteTeam.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling deleteTeam.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/teams/{name}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete team
     */
    async deleteTeam(requestParameters: DeleteTeamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteTeamRaw(requestParameters, initOverrides);
    }

    /**
     * Delete team member details
     */
    async deleteTeamMemberRaw(requestParameters: DeleteTeamMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling deleteTeamMember.');
        }

        if (requestParameters.team === null || requestParameters.team === undefined) {
            throw new runtime.RequiredError('team','Required parameter requestParameters.team was null or undefined when calling deleteTeamMember.');
        }

        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling deleteTeamMember.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/teams/{team}/members/{user}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"team"}}`, encodeURIComponent(String(requestParameters.team))).replace(`{${"user"}}`, encodeURIComponent(String(requestParameters.user))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete team member details
     */
    async deleteTeamMember(requestParameters: DeleteTeamMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteTeamMemberRaw(requestParameters, initOverrides);
    }

    /**
     * Get team
     */
    async getTeamRaw(requestParameters: GetTeamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Team>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling getTeam.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling getTeam.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/teams/{name}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TeamFromJSON(jsonValue));
    }

    /**
     * Get team
     */
    async getTeam(requestParameters: GetTeamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Team> {
        const response = await this.getTeamRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get team member details
     */
    async getTeamMemberRaw(requestParameters: GetTeamMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1TeamMember>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling getTeamMember.');
        }

        if (requestParameters.team === null || requestParameters.team === undefined) {
            throw new runtime.RequiredError('team','Required parameter requestParameters.team was null or undefined when calling getTeamMember.');
        }

        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling getTeamMember.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/teams/{team}/members/{user}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"team"}}`, encodeURIComponent(String(requestParameters.team))).replace(`{${"user"}}`, encodeURIComponent(String(requestParameters.user))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TeamMemberFromJSON(jsonValue));
    }

    /**
     * Get team member details
     */
    async getTeamMember(requestParameters: GetTeamMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1TeamMember> {
        const response = await this.getTeamMemberRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get team members
     */
    async listTeamMembersRaw(requestParameters: ListTeamMembersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ListTeamMembersResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listTeamMembers.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling listTeamMembers.');
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
            path: `/api/v1/orgs/{owner}/teams/{name}/members`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListTeamMembersResponseFromJSON(jsonValue));
    }

    /**
     * Get team members
     */
    async listTeamMembers(requestParameters: ListTeamMembersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ListTeamMembersResponse> {
        const response = await this.listTeamMembersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List teams names
     */
    async listTeamNamesRaw(requestParameters: ListTeamNamesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ListTeamsResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listTeamNames.');
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
            path: `/api/v1/orgs/{owner}/teams/names`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListTeamsResponseFromJSON(jsonValue));
    }

    /**
     * List teams names
     */
    async listTeamNames(requestParameters: ListTeamNamesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ListTeamsResponse> {
        const response = await this.listTeamNamesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List teams
     */
    async listTeamsRaw(requestParameters: ListTeamsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1ListTeamsResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listTeams.');
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
            path: `/api/v1/orgs/{owner}/teams`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListTeamsResponseFromJSON(jsonValue));
    }

    /**
     * List teams
     */
    async listTeams(requestParameters: ListTeamsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1ListTeamsResponse> {
        const response = await this.listTeamsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Patch team
     */
    async patchTeamRaw(requestParameters: PatchTeamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Team>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling patchTeam.');
        }

        if (requestParameters.teamName === null || requestParameters.teamName === undefined) {
            throw new runtime.RequiredError('teamName','Required parameter requestParameters.teamName was null or undefined when calling patchTeam.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling patchTeam.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/teams/{team.name}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"team.name"}}`, encodeURIComponent(String(requestParameters.teamName))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V1TeamToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TeamFromJSON(jsonValue));
    }

    /**
     * Patch team
     */
    async patchTeam(requestParameters: PatchTeamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Team> {
        const response = await this.patchTeamRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Patch team member
     */
    async patchTeamMemberRaw(requestParameters: PatchTeamMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1TeamMember>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling patchTeamMember.');
        }

        if (requestParameters.team === null || requestParameters.team === undefined) {
            throw new runtime.RequiredError('team','Required parameter requestParameters.team was null or undefined when calling patchTeamMember.');
        }

        if (requestParameters.memberUser === null || requestParameters.memberUser === undefined) {
            throw new runtime.RequiredError('memberUser','Required parameter requestParameters.memberUser was null or undefined when calling patchTeamMember.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling patchTeamMember.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/teams/{team}/members/{member.user}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"team"}}`, encodeURIComponent(String(requestParameters.team))).replace(`{${"member.user"}}`, encodeURIComponent(String(requestParameters.memberUser))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V1TeamMemberToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TeamMemberFromJSON(jsonValue));
    }

    /**
     * Patch team member
     */
    async patchTeamMember(requestParameters: PatchTeamMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1TeamMember> {
        const response = await this.patchTeamMemberRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Update team
     */
    async updateTeamRaw(requestParameters: UpdateTeamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Team>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling updateTeam.');
        }

        if (requestParameters.teamName === null || requestParameters.teamName === undefined) {
            throw new runtime.RequiredError('teamName','Required parameter requestParameters.teamName was null or undefined when calling updateTeam.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateTeam.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/teams/{team.name}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"team.name"}}`, encodeURIComponent(String(requestParameters.teamName))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V1TeamToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TeamFromJSON(jsonValue));
    }

    /**
     * Update team
     */
    async updateTeam(requestParameters: UpdateTeamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Team> {
        const response = await this.updateTeamRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Update team member
     */
    async updateTeamMemberRaw(requestParameters: UpdateTeamMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1TeamMember>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling updateTeamMember.');
        }

        if (requestParameters.team === null || requestParameters.team === undefined) {
            throw new runtime.RequiredError('team','Required parameter requestParameters.team was null or undefined when calling updateTeamMember.');
        }

        if (requestParameters.memberUser === null || requestParameters.memberUser === undefined) {
            throw new runtime.RequiredError('memberUser','Required parameter requestParameters.memberUser was null or undefined when calling updateTeamMember.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateTeamMember.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/teams/{team}/members/{member.user}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"team"}}`, encodeURIComponent(String(requestParameters.team))).replace(`{${"member.user"}}`, encodeURIComponent(String(requestParameters.memberUser))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V1TeamMemberToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1TeamMemberFromJSON(jsonValue));
    }

    /**
     * Update team member
     */
    async updateTeamMember(requestParameters: UpdateTeamMemberRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1TeamMember> {
        const response = await this.updateTeamMemberRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
