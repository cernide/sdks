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
    V1Run,
    V1RunFromJSON,
    V1RunFromJSONTyped,
    V1RunToJSON,
    V1RunEdgeKind,
    V1RunEdgeKindFromJSON,
    V1RunEdgeKindFromJSONTyped,
    V1RunEdgeKindToJSON,
    V1Statuses,
    V1StatusesFromJSON,
    V1StatusesFromJSONTyped,
    V1StatusesToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1RunEdge
 */
export interface V1RunEdge {
    /**
     * 
     * @type {V1Run}
     * @memberof V1RunEdge
     */
    upstream?: V1Run;
    /**
     * 
     * @type {V1Run}
     * @memberof V1RunEdge
     */
    downstream?: V1Run;
    /**
     * 
     * @type {V1RunEdgeKind}
     * @memberof V1RunEdge
     */
    kind?: V1RunEdgeKind;
    /**
     * 
     * @type {object}
     * @memberof V1RunEdge
     */
    values?: object;
    /**
     * 
     * @type {Array<V1Statuses>}
     * @memberof V1RunEdge
     */
    statuses?: Array<V1Statuses>;
}

export function V1RunEdgeFromJSON(json: any): V1RunEdge {
    return V1RunEdgeFromJSONTyped(json, false);
}

export function V1RunEdgeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1RunEdge {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'upstream': !exists(json, 'upstream') ? undefined : V1RunFromJSON(json['upstream']),
        'downstream': !exists(json, 'downstream') ? undefined : V1RunFromJSON(json['downstream']),
        'kind': !exists(json, 'kind') ? undefined : V1RunEdgeKindFromJSON(json['kind']),
        'values': !exists(json, 'values') ? undefined : json['values'],
        'statuses': !exists(json, 'statuses') ? undefined : ((json['statuses'] as Array<any>).map(V1StatusesFromJSON)),
    };
}

export function V1RunEdgeToJSON(value?: V1RunEdge | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'upstream': V1RunToJSON(value.upstream),
        'downstream': V1RunToJSON(value.downstream),
        'kind': V1RunEdgeKindToJSON(value.kind),
        'values': value.values,
        'statuses': value.statuses === undefined ? undefined : ((value.statuses as Array<any>).map(V1StatusesToJSON)),
    };
}

