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

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 2.0.0-rc5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  RuntimeError,
  V1Compatibility,
  V1Installation,
  V1LogHandler,
} from '../models';
import {
    RuntimeErrorFromJSON,
    RuntimeErrorToJSON,
    V1CompatibilityFromJSON,
    V1CompatibilityToJSON,
    V1InstallationFromJSON,
    V1InstallationToJSON,
    V1LogHandlerFromJSON,
    V1LogHandlerToJSON,
} from '../models';

export interface GetCompatibilityRequest {
    uuid: string;
    version: string;
    service: string;
}

export interface GetInstallationRequest {
    auth?: boolean;
}

/**
 *
 */
export class VersionsV1Api extends runtime.BaseAPI {

    /**
     * Get compatibility versions
     */
    async getCompatibilityRaw(requestParameters: GetCompatibilityRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Compatibility>> {
        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling getCompatibility.');
        }

        if (requestParameters.version === null || requestParameters.version === undefined) {
            throw new runtime.RequiredError('version','Required parameter requestParameters.version was null or undefined when calling getCompatibility.');
        }

        if (requestParameters.service === null || requestParameters.service === undefined) {
            throw new runtime.RequiredError('service','Required parameter requestParameters.service was null or undefined when calling getCompatibility.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/compatibility/{uuid}/{version}/{service}`.replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))).replace(`{${"version"}}`, encodeURIComponent(String(requestParameters.version))).replace(`{${"service"}}`, encodeURIComponent(String(requestParameters.service))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1CompatibilityFromJSON(jsonValue));
    }

    /**
     * Get compatibility versions
     */
    async getCompatibility(requestParameters: GetCompatibilityRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Compatibility> {
        const response = await this.getCompatibilityRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get installation versions
     */
    async getInstallationRaw(requestParameters: GetInstallationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1Installation>> {
        const queryParameters: any = {};

        if (requestParameters.auth !== undefined) {
            queryParameters['auth'] = requestParameters.auth;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/installation`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1InstallationFromJSON(jsonValue));
    }

    /**
     * Get installation versions
     */
    async getInstallation(requestParameters: GetInstallationRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1Installation> {
        const response = await this.getInstallationRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get log handler versions
     */
    async getLogHandlerRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V1LogHandler>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/log_handler`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V1LogHandlerFromJSON(jsonValue));
    }

    /**
     * Get log handler versions
     */
    async getLogHandler(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V1LogHandler> {
        const response = await this.getLogHandlerRaw(initOverrides);
        return await response.value();
    }

}
