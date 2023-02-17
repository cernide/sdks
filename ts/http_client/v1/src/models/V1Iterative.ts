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
    V1Tuner,
    V1TunerFromJSON,
    V1TunerFromJSONTyped,
    V1TunerToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1Iterative
 */
export interface V1Iterative {
    /**
     * 
     * @type {string}
     * @memberof V1Iterative
     */
    kind?: string;
    /**
     * 
     * @type {{ [key: string]: object; }}
     * @memberof V1Iterative
     */
    params?: { [key: string]: object; };
    /**
     * 
     * @type {number}
     * @memberof V1Iterative
     */
    maxIterations?: number;
    /**
     * 
     * @type {number}
     * @memberof V1Iterative
     */
    seed?: number;
    /**
     * 
     * @type {number}
     * @memberof V1Iterative
     */
    concurrency?: number;
    /**
     * 
     * @type {V1Tuner}
     * @memberof V1Iterative
     */
    tuner?: V1Tuner;
    /**
     * 
     * @type {Array<object>}
     * @memberof V1Iterative
     */
    earlyStopping?: Array<object>;
}

export function V1IterativeFromJSON(json: any): V1Iterative {
    return V1IterativeFromJSONTyped(json, false);
}

export function V1IterativeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Iterative {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'params': !exists(json, 'params') ? undefined : json['params'],
        'maxIterations': !exists(json, 'maxIterations') ? undefined : json['maxIterations'],
        'seed': !exists(json, 'seed') ? undefined : json['seed'],
        'concurrency': !exists(json, 'concurrency') ? undefined : json['concurrency'],
        'tuner': !exists(json, 'tuner') ? undefined : V1TunerFromJSON(json['tuner']),
        'earlyStopping': !exists(json, 'earlyStopping') ? undefined : json['earlyStopping'],
    };
}

export function V1IterativeToJSON(value?: V1Iterative | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'kind': value.kind,
        'params': value.params,
        'maxIterations': value.maxIterations,
        'seed': value.seed,
        'concurrency': value.concurrency,
        'tuner': V1TunerToJSON(value.tuner),
        'earlyStopping': value.earlyStopping,
    };
}

