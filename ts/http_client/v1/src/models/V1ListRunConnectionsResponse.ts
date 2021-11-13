// Copyright 2018-2021 Polyaxon, Inc.
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
 * The version of the OpenAPI document: 1.13.0-rc2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    V1RunConnection,
    V1RunConnectionFromJSON,
    V1RunConnectionFromJSONTyped,
    V1RunConnectionToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1ListRunConnectionsResponse
 */
export interface V1ListRunConnectionsResponse {
    /**
     * 
     * @type {number}
     * @memberof V1ListRunConnectionsResponse
     */
    count?: number;
    /**
     * 
     * @type {Array<V1RunConnection>}
     * @memberof V1ListRunConnectionsResponse
     */
    results?: Array<V1RunConnection>;
    /**
     * 
     * @type {string}
     * @memberof V1ListRunConnectionsResponse
     */
    previous?: string;
    /**
     * 
     * @type {string}
     * @memberof V1ListRunConnectionsResponse
     */
    next?: string;
}

export function V1ListRunConnectionsResponseFromJSON(json: any): V1ListRunConnectionsResponse {
    return V1ListRunConnectionsResponseFromJSONTyped(json, false);
}

export function V1ListRunConnectionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ListRunConnectionsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'count': !exists(json, 'count') ? undefined : json['count'],
        'results': !exists(json, 'results') ? undefined : ((json['results'] as Array<any>).map(V1RunConnectionFromJSON)),
        'previous': !exists(json, 'previous') ? undefined : json['previous'],
        'next': !exists(json, 'next') ? undefined : json['next'],
    };
}

export function V1ListRunConnectionsResponseToJSON(value?: V1ListRunConnectionsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'count': value.count,
        'results': value.results === undefined ? undefined : ((value.results as Array<any>).map(V1RunConnectionToJSON)),
        'previous': value.previous,
        'next': value.next,
    };
}

