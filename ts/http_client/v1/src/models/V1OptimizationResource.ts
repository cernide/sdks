/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { V1ResourceType } from './V1ResourceType';
import {
    V1ResourceTypeFromJSON,
    V1ResourceTypeFromJSONTyped,
    V1ResourceTypeToJSON,
} from './V1ResourceType';

/**
 *
 * @export
 * @interface V1OptimizationResource
 */
export interface V1OptimizationResource {
    /**
     *
     * @type {string}
     * @memberof V1OptimizationResource
     */
    name?: string;
    /**
     *
     * @type {V1ResourceType}
     * @memberof V1OptimizationResource
     */
    type?: V1ResourceType;
}

/**
 * Check if a given object implements the V1OptimizationResource interface.
 */
export function instanceOfV1OptimizationResource(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1OptimizationResourceFromJSON(json: any): V1OptimizationResource {
    return V1OptimizationResourceFromJSONTyped(json, false);
}

export function V1OptimizationResourceFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1OptimizationResource {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'name': !exists(json, 'name') ? undefined : json['name'],
        'type': !exists(json, 'type') ? undefined : V1ResourceTypeFromJSON(json['type']),
    };
}

export function V1OptimizationResourceToJSON(value?: V1OptimizationResource | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'name': value.name,
        'type': V1ResourceTypeToJSON(value.type),
    };
}

