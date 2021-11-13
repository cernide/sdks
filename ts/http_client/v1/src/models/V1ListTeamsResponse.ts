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
    V1Team,
    V1TeamFromJSON,
    V1TeamFromJSONTyped,
    V1TeamToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1ListTeamsResponse
 */
export interface V1ListTeamsResponse {
    /**
     * 
     * @type {number}
     * @memberof V1ListTeamsResponse
     */
    count?: number;
    /**
     * 
     * @type {Array<V1Team>}
     * @memberof V1ListTeamsResponse
     */
    results?: Array<V1Team>;
    /**
     * 
     * @type {string}
     * @memberof V1ListTeamsResponse
     */
    previous?: string;
    /**
     * 
     * @type {string}
     * @memberof V1ListTeamsResponse
     */
    next?: string;
}

export function V1ListTeamsResponseFromJSON(json: any): V1ListTeamsResponse {
    return V1ListTeamsResponseFromJSONTyped(json, false);
}

export function V1ListTeamsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ListTeamsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'count': !exists(json, 'count') ? undefined : json['count'],
        'results': !exists(json, 'results') ? undefined : ((json['results'] as Array<any>).map(V1TeamFromJSON)),
        'previous': !exists(json, 'previous') ? undefined : json['previous'],
        'next': !exists(json, 'next') ? undefined : json['next'],
    };
}

export function V1ListTeamsResponseToJSON(value?: V1ListTeamsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'count': value.count,
        'results': value.results === undefined ? undefined : ((value.results as Array<any>).map(V1TeamToJSON)),
        'previous': value.previous,
        'next': value.next,
    };
}

