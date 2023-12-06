/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.0.6-rc0
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
 * @interface V1HpLogNormal
 */
export interface V1HpLogNormal {
    /**
     * 
     * @type {string}
     * @memberof V1HpLogNormal
     */
    kind?: string;
    /**
     * 
     * @type {object}
     * @memberof V1HpLogNormal
     */
    value?: object;
}

/**
 * Check if a given object implements the V1HpLogNormal interface.
 */
export function instanceOfV1HpLogNormal(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1HpLogNormalFromJSON(json: any): V1HpLogNormal {
    return V1HpLogNormalFromJSONTyped(json, false);
}

export function V1HpLogNormalFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1HpLogNormal {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'value': !exists(json, 'value') ? undefined : json['value'],
    };
}

export function V1HpLogNormalToJSON(value?: V1HpLogNormal | null): any {
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

