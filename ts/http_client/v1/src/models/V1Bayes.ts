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
 * The version of the OpenAPI document: 1.17.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    V1OptimizationMetric,
    V1OptimizationMetricFromJSON,
    V1OptimizationMetricFromJSONTyped,
    V1OptimizationMetricToJSON,
    V1Tuner,
    V1TunerFromJSON,
    V1TunerFromJSONTyped,
    V1TunerToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1Bayes
 */
export interface V1Bayes {
    /**
     * 
     * @type {string}
     * @memberof V1Bayes
     */
    kind?: string;
    /**
     * 
     * @type {{ [key: string]: object; }}
     * @memberof V1Bayes
     */
    params?: { [key: string]: object; };
    /**
     * 
     * @type {number}
     * @memberof V1Bayes
     */
    numInitialRuns?: number;
    /**
     * 
     * @type {number}
     * @memberof V1Bayes
     */
    maxIterations?: number;
    /**
     * 
     * @type {object}
     * @memberof V1Bayes
     */
    utilityFunction?: object;
    /**
     * 
     * @type {V1OptimizationMetric}
     * @memberof V1Bayes
     */
    metric?: V1OptimizationMetric;
    /**
     * 
     * @type {number}
     * @memberof V1Bayes
     */
    seed?: number;
    /**
     * 
     * @type {number}
     * @memberof V1Bayes
     */
    concurrency?: number;
    /**
     * 
     * @type {V1Tuner}
     * @memberof V1Bayes
     */
    tuner?: V1Tuner;
    /**
     * 
     * @type {Array<object>}
     * @memberof V1Bayes
     */
    earlyStopping?: Array<object>;
}

export function V1BayesFromJSON(json: any): V1Bayes {
    return V1BayesFromJSONTyped(json, false);
}

export function V1BayesFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Bayes {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'params': !exists(json, 'params') ? undefined : json['params'],
        'numInitialRuns': !exists(json, 'numInitialRuns') ? undefined : json['numInitialRuns'],
        'maxIterations': !exists(json, 'maxIterations') ? undefined : json['maxIterations'],
        'utilityFunction': !exists(json, 'utilityFunction') ? undefined : json['utilityFunction'],
        'metric': !exists(json, 'metric') ? undefined : V1OptimizationMetricFromJSON(json['metric']),
        'seed': !exists(json, 'seed') ? undefined : json['seed'],
        'concurrency': !exists(json, 'concurrency') ? undefined : json['concurrency'],
        'tuner': !exists(json, 'tuner') ? undefined : V1TunerFromJSON(json['tuner']),
        'earlyStopping': !exists(json, 'earlyStopping') ? undefined : json['earlyStopping'],
    };
}

export function V1BayesToJSON(value?: V1Bayes | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'kind': value.kind,
        'params': value.params,
        'numInitialRuns': value.numInitialRuns,
        'maxIterations': value.maxIterations,
        'utilityFunction': value.utilityFunction,
        'metric': V1OptimizationMetricToJSON(value.metric),
        'seed': value.seed,
        'concurrency': value.concurrency,
        'tuner': V1TunerToJSON(value.tuner),
        'earlyStopping': value.earlyStopping,
    };
}

