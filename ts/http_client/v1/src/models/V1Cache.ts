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
 * The version of the OpenAPI document: 1.13.0-rc3
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
 * @interface V1Cache
 */
export interface V1Cache {
    /**
     * 
     * @type {boolean}
     * @memberof V1Cache
     */
    disable?: boolean;
    /**
     * 
     * @type {number}
     * @memberof V1Cache
     */
    ttl?: number;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1Cache
     */
    io?: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1Cache
     */
    sections?: Array<string>;
}

export function V1CacheFromJSON(json: any): V1Cache {
    return V1CacheFromJSONTyped(json, false);
}

export function V1CacheFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Cache {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'disable': !exists(json, 'disable') ? undefined : json['disable'],
        'ttl': !exists(json, 'ttl') ? undefined : json['ttl'],
        'io': !exists(json, 'io') ? undefined : json['io'],
        'sections': !exists(json, 'sections') ? undefined : json['sections'],
    };
}

export function V1CacheToJSON(value?: V1Cache | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'disable': value.disable,
        'ttl': value.ttl,
        'io': value.io,
        'sections': value.sections,
    };
}
