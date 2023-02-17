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
 * @interface V1Param
 */
export interface V1Param {
    /**
     * 
     * @type {object}
     * @memberof V1Param
     */
    value?: object;
    /**
     * 
     * @type {string}
     * @memberof V1Param
     */
    ref?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Param
     */
    connection?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V1Param
     */
    toInit?: boolean;
    /**
     * 
     * @type {string}
     * @memberof V1Param
     */
    toEnv?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V1Param
     */
    contextOnly?: boolean;
}

export function V1ParamFromJSON(json: any): V1Param {
    return V1ParamFromJSONTyped(json, false);
}

export function V1ParamFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Param {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'value': !exists(json, 'value') ? undefined : json['value'],
        'ref': !exists(json, 'ref') ? undefined : json['ref'],
        'connection': !exists(json, 'connection') ? undefined : json['connection'],
        'toInit': !exists(json, 'toInit') ? undefined : json['toInit'],
        'toEnv': !exists(json, 'toEnv') ? undefined : json['toEnv'],
        'contextOnly': !exists(json, 'contextOnly') ? undefined : json['contextOnly'],
    };
}

export function V1ParamToJSON(value?: V1Param | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'value': value.value,
        'ref': value.ref,
        'connection': value.connection,
        'toInit': value.toInit,
        'toEnv': value.toEnv,
        'contextOnly': value.contextOnly,
    };
}

