/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.2.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { V1EventCurveKind } from './V1EventCurveKind';
import {
    V1EventCurveKindFromJSON,
    V1EventCurveKindFromJSONTyped,
    V1EventCurveKindToJSON,
} from './V1EventCurveKind';

/**
 *
 * @export
 * @interface V1EventCurve
 */
export interface V1EventCurve {
    /**
     *
     * @type {V1EventCurveKind}
     * @memberof V1EventCurve
     */
    kind?: V1EventCurveKind;
    /**
     *
     * @type {Array<number>}
     * @memberof V1EventCurve
     */
    x?: Array<number>;
    /**
     *
     * @type {Array<number>}
     * @memberof V1EventCurve
     */
    y?: Array<number>;
    /**
     *
     * @type {string}
     * @memberof V1EventCurve
     */
    annotation?: string;
}

/**
 * Check if a given object implements the V1EventCurve interface.
 */
export function instanceOfV1EventCurve(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1EventCurveFromJSON(json: any): V1EventCurve {
    return V1EventCurveFromJSONTyped(json, false);
}

export function V1EventCurveFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1EventCurve {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'kind': !exists(json, 'kind') ? undefined : V1EventCurveKindFromJSON(json['kind']),
        'x': !exists(json, 'x') ? undefined : json['x'],
        'y': !exists(json, 'y') ? undefined : json['y'],
        'annotation': !exists(json, 'annotation') ? undefined : json['annotation'],
    };
}

export function V1EventCurveToJSON(value?: V1EventCurve | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'kind': V1EventCurveKindToJSON(value.kind),
        'x': value.x,
        'y': value.y,
        'annotation': value.annotation,
    };
}

