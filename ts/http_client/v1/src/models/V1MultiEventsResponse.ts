/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.5
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
 * @interface V1MultiEventsResponse
 */
export interface V1MultiEventsResponse {
    /**
     * 
     * @type {object}
     * @memberof V1MultiEventsResponse
     */
    data?: object;
}

/**
 * Check if a given object implements the V1MultiEventsResponse interface.
 */
export function instanceOfV1MultiEventsResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1MultiEventsResponseFromJSON(json: any): V1MultiEventsResponse {
    return V1MultiEventsResponseFromJSONTyped(json, false);
}

export function V1MultiEventsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1MultiEventsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'data': !exists(json, 'data') ? undefined : json['data'],
    };
}

export function V1MultiEventsResponseToJSON(value?: V1MultiEventsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'data': value.data,
    };
}

