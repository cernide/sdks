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

import { exists, mapValues } from '../runtime';
import type { V1ProjectVersion } from './V1ProjectVersion';
import {
    V1ProjectVersionFromJSON,
    V1ProjectVersionFromJSONTyped,
    V1ProjectVersionToJSON,
} from './V1ProjectVersion';

/**
 *
 * @export
 * @interface V1ListProjectVersionsResponse
 */
export interface V1ListProjectVersionsResponse {
    /**
     *
     * @type {number}
     * @memberof V1ListProjectVersionsResponse
     */
    count?: number;
    /**
     *
     * @type {Array<V1ProjectVersion>}
     * @memberof V1ListProjectVersionsResponse
     */
    results?: Array<V1ProjectVersion>;
    /**
     *
     * @type {string}
     * @memberof V1ListProjectVersionsResponse
     */
    previous?: string;
    /**
     *
     * @type {string}
     * @memberof V1ListProjectVersionsResponse
     */
    next?: string;
}

/**
 * Check if a given object implements the V1ListProjectVersionsResponse interface.
 */
export function instanceOfV1ListProjectVersionsResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1ListProjectVersionsResponseFromJSON(json: any): V1ListProjectVersionsResponse {
    return V1ListProjectVersionsResponseFromJSONTyped(json, false);
}

export function V1ListProjectVersionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ListProjectVersionsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'count': !exists(json, 'count') ? undefined : json['count'],
        'results': !exists(json, 'results') ? undefined : ((json['results'] as Array<any>).map(V1ProjectVersionFromJSON)),
        'previous': !exists(json, 'previous') ? undefined : json['previous'],
        'next': !exists(json, 'next') ? undefined : json['next'],
    };
}

export function V1ListProjectVersionsResponseToJSON(value?: V1ListProjectVersionsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'count': value.count,
        'results': value.results === undefined ? undefined : ((value.results as Array<any>).map(V1ProjectVersionToJSON)),
        'previous': value.previous,
        'next': value.next,
    };
}

