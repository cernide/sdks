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
 * The version of the OpenAPI document: 2.0.0-rc2
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
 * @interface V1ClaimConnection
 */
export interface V1ClaimConnection {
    /**
     *
     * @type {string}
     * @memberof V1ClaimConnection
     */
    volumeClaim?: string;
    /**
     *
     * @type {string}
     * @memberof V1ClaimConnection
     */
    mountPath?: string;
    /**
     *
     * @type {boolean}
     * @memberof V1ClaimConnection
     */
    readOnly?: boolean;
    /**
     *
     * @type {object}
     * @memberof V1ClaimConnection
     */
    kind?: object;
}

/**
 * Check if a given object implements the V1ClaimConnection interface.
 */
export function instanceOfV1ClaimConnection(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1ClaimConnectionFromJSON(json: any): V1ClaimConnection {
    return V1ClaimConnectionFromJSONTyped(json, false);
}

export function V1ClaimConnectionFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ClaimConnection {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'volumeClaim': !exists(json, 'volumeClaim') ? undefined : json['volumeClaim'],
        'mountPath': !exists(json, 'mountPath') ? undefined : json['mountPath'],
        'readOnly': !exists(json, 'readOnly') ? undefined : json['readOnly'],
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
    };
}

export function V1ClaimConnectionToJSON(value?: V1ClaimConnection | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'volumeClaim': value.volumeClaim,
        'mountPath': value.mountPath,
        'readOnly': value.readOnly,
        'kind': value.kind,
    };
}

