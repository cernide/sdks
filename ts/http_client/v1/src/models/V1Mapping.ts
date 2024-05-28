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
/**
 *
 * @export
 * @interface V1Mapping
 */
export interface V1Mapping {
    /**
     *
     * @type {string}
     * @memberof V1Mapping
     */
    kind?: string;
    /**
     *
     * @type {Array<object>}
     * @memberof V1Mapping
     */
    values?: Array<object>;
    /**
     *
     * @type {number}
     * @memberof V1Mapping
     */
    concurrency?: number;
    /**
     *
     * @type {Array<object>}
     * @memberof V1Mapping
     */
    earlyStopping?: Array<object>;
}

/**
 * Check if a given object implements the V1Mapping interface.
 */
export function instanceOfV1Mapping(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1MappingFromJSON(json: any): V1Mapping {
    return V1MappingFromJSONTyped(json, false);
}

export function V1MappingFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Mapping {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'values': !exists(json, 'values') ? undefined : json['values'],
        'concurrency': !exists(json, 'concurrency') ? undefined : json['concurrency'],
        'earlyStopping': !exists(json, 'earlyStopping') ? undefined : json['earlyStopping'],
    };
}

export function V1MappingToJSON(value?: V1Mapping | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'kind': value.kind,
        'values': value.values,
        'concurrency': value.concurrency,
        'earlyStopping': value.earlyStopping,
    };
}

