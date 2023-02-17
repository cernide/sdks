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
 * The version of the OpenAPI document: 1.22.0-rc0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface V1JoinParam
 */
export interface V1JoinParam {
    /**
     * 
     * @type {string}
     * @memberof V1JoinParam
     */
    value?: string;
    /**
     * 
     * @type {string}
     * @memberof V1JoinParam
     */
    connection?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V1JoinParam
     */
    toInit?: boolean;
    /**
     * 
     * @type {string}
     * @memberof V1JoinParam
     */
    toEnv?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V1JoinParam
     */
    contextOnly?: boolean;
}

export function V1JoinParamFromJSON(json: any): V1JoinParam {
    return V1JoinParamFromJSONTyped(json, false);
}

export function V1JoinParamFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1JoinParam {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'value': !exists(json, 'value') ? undefined : json['value'],
        'connection': !exists(json, 'connection') ? undefined : json['connection'],
        'toInit': !exists(json, 'toInit') ? undefined : json['toInit'],
        'toEnv': !exists(json, 'toEnv') ? undefined : json['toEnv'],
        'contextOnly': !exists(json, 'contextOnly') ? undefined : json['contextOnly'],
    };
}

export function V1JoinParamToJSON(value?: V1JoinParam | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'value': value.value,
        'connection': value.connection,
        'toInit': value.toInit,
        'toEnv': value.toEnv,
        'contextOnly': value.contextOnly,
    };
}

