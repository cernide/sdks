/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.0.0-rc14
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
 * @interface V1Ray
 */
export interface V1Ray {
    /**
     * 
     * @type {string}
     * @memberof V1Ray
     */
    kind?: string;
    /**
     * 
     * @type {object}
     * @memberof V1Ray
     */
    spec?: object;
}

/**
 * Check if a given object implements the V1Ray interface.
 */
export function instanceOfV1Ray(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1RayFromJSON(json: any): V1Ray {
    return V1RayFromJSONTyped(json, false);
}

export function V1RayFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Ray {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'spec': !exists(json, 'spec') ? undefined : json['spec'],
    };
}

export function V1RayToJSON(value?: V1Ray | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'kind': value.kind,
        'spec': value.spec,
    };
}

