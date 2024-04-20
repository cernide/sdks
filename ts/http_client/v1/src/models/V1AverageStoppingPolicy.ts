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
 * @interface V1AverageStoppingPolicy
 */
export interface V1AverageStoppingPolicy {
    /**
     * 
     * @type {object}
     * @memberof V1AverageStoppingPolicy
     */
    kind?: object;
}

/**
 * Check if a given object implements the V1AverageStoppingPolicy interface.
 */
export function instanceOfV1AverageStoppingPolicy(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1AverageStoppingPolicyFromJSON(json: any): V1AverageStoppingPolicy {
    return V1AverageStoppingPolicyFromJSONTyped(json, false);
}

export function V1AverageStoppingPolicyFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1AverageStoppingPolicy {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
    };
}

export function V1AverageStoppingPolicyToJSON(value?: V1AverageStoppingPolicy | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'kind': value.kind,
    };
}

