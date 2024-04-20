/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.8
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
 * @interface V1HpRange
 */
export interface V1HpRange {
    /**
     * 
     * @type {string}
     * @memberof V1HpRange
     */
    kind?: string;
    /**
     * 
     * @type {object}
     * @memberof V1HpRange
     */
    value?: object;
}

/**
 * Check if a given object implements the V1HpRange interface.
 */
export function instanceOfV1HpRange(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1HpRangeFromJSON(json: any): V1HpRange {
    return V1HpRangeFromJSONTyped(json, false);
}

export function V1HpRangeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1HpRange {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'value': !exists(json, 'value') ? undefined : json['value'],
    };
}

export function V1HpRangeToJSON(value?: V1HpRange | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'kind': value.kind,
        'value': value.value,
    };
}

