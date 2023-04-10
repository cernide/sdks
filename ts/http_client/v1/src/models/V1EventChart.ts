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
import type { V1EventChartKind } from './V1EventChartKind';
import {
    V1EventChartKindFromJSON,
    V1EventChartKindFromJSONTyped,
    V1EventChartKindToJSON,
} from './V1EventChartKind';

/**
 *
 * @export
 * @interface V1EventChart
 */
export interface V1EventChart {
    /**
     *
     * @type {V1EventChartKind}
     * @memberof V1EventChart
     */
    kind?: V1EventChartKind;
    /**
     *
     * @type {object}
     * @memberof V1EventChart
     */
    figure?: object;
}

/**
 * Check if a given object implements the V1EventChart interface.
 */
export function instanceOfV1EventChart(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1EventChartFromJSON(json: any): V1EventChart {
    return V1EventChartFromJSONTyped(json, false);
}

export function V1EventChartFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1EventChart {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'kind': !exists(json, 'kind') ? undefined : V1EventChartKindFromJSON(json['kind']),
        'figure': !exists(json, 'figure') ? undefined : json['figure'],
    };
}

export function V1EventChartToJSON(value?: V1EventChart | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'kind': V1EventChartKindToJSON(value.kind),
        'figure': value.figure,
    };
}

