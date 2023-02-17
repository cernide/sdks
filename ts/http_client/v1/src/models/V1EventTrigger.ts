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
import {
    V1EventKind,
    V1EventKindFromJSON,
    V1EventKindFromJSONTyped,
    V1EventKindToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1EventTrigger
 */
export interface V1EventTrigger {
    /**
     * 
     * @type {Array<V1EventKind>}
     * @memberof V1EventTrigger
     */
    kinds?: Array<V1EventKind>;
    /**
     * 
     * @type {string}
     * @memberof V1EventTrigger
     */
    ref?: string;
}

export function V1EventTriggerFromJSON(json: any): V1EventTrigger {
    return V1EventTriggerFromJSONTyped(json, false);
}

export function V1EventTriggerFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1EventTrigger {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'kinds': !exists(json, 'kinds') ? undefined : ((json['kinds'] as Array<any>).map(V1EventKindFromJSON)),
        'ref': !exists(json, 'ref') ? undefined : json['ref'],
    };
}

export function V1EventTriggerToJSON(value?: V1EventTrigger | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'kinds': value.kinds === undefined ? undefined : ((value.kinds as Array<any>).map(V1EventKindToJSON)),
        'ref': value.ref,
    };
}

