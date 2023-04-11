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
/**
 *
 * @export
 * @interface V1TrialStart
 */
export interface V1TrialStart {
    /**
     *
     * @type {string}
     * @memberof V1TrialStart
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof V1TrialStart
     */
    email?: string;
    /**
     *
     * @type {string}
     * @memberof V1TrialStart
     */
    organization?: string;
    /**
     *
     * @type {string}
     * @memberof V1TrialStart
     */
    plan?: string;
    /**
     *
     * @type {number}
     * @memberof V1TrialStart
     */
    seats?: number;
    /**
     *
     * @type {object}
     * @memberof V1TrialStart
     */
    details?: object;
}

/**
 * Check if a given object implements the V1TrialStart interface.
 */
export function instanceOfV1TrialStart(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1TrialStartFromJSON(json: any): V1TrialStart {
    return V1TrialStartFromJSONTyped(json, false);
}

export function V1TrialStartFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1TrialStart {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'name': !exists(json, 'name') ? undefined : json['name'],
        'email': !exists(json, 'email') ? undefined : json['email'],
        'organization': !exists(json, 'organization') ? undefined : json['organization'],
        'plan': !exists(json, 'plan') ? undefined : json['plan'],
        'seats': !exists(json, 'seats') ? undefined : json['seats'],
        'details': !exists(json, 'details') ? undefined : json['details'],
    };
}

export function V1TrialStartToJSON(value?: V1TrialStart | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'name': value.name,
        'email': value.email,
        'organization': value.organization,
        'plan': value.plan,
        'seats': value.seats,
        'details': value.details,
    };
}

