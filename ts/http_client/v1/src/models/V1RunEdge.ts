/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.3.2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { V1Run } from './V1Run';
import {
    V1RunFromJSON,
    V1RunFromJSONTyped,
    V1RunToJSON,
} from './V1Run';
import type { V1RunEdgeKind } from './V1RunEdgeKind';
import {
    V1RunEdgeKindFromJSON,
    V1RunEdgeKindFromJSONTyped,
    V1RunEdgeKindToJSON,
} from './V1RunEdgeKind';
import type { V1Statuses } from './V1Statuses';
import {
    V1StatusesFromJSON,
    V1StatusesFromJSONTyped,
    V1StatusesToJSON,
} from './V1Statuses';

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

/**
 * Check if a given object implements the V1RunEdge interface.
 */
export function instanceOfV1RunEdge(value: object): boolean {
    let isInstance = true;

    return isInstance;
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

