/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc28
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
 * @interface V1Cache
 */
export interface V1Cache {
    /**
     *
     * @type {boolean}
     * @memberof V1Cache
     */
    disable?: boolean;
    /**
     *
     * @type {number}
     * @memberof V1Cache
     */
    ttl?: number;
    /**
     *
     * @type {Array<string>}
     * @memberof V1Cache
     */
    io?: Array<string>;
    /**
     *
     * @type {Array<string>}
     * @memberof V1Cache
     */
    sections?: Array<string>;
}

/**
 * Check if a given object implements the V1Cache interface.
 */
export function instanceOfV1Cache(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1CacheFromJSON(json: any): V1Cache {
    return V1CacheFromJSONTyped(json, false);
}

export function V1CacheFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Cache {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'disable': !exists(json, 'disable') ? undefined : json['disable'],
        'ttl': !exists(json, 'ttl') ? undefined : json['ttl'],
        'io': !exists(json, 'io') ? undefined : json['io'],
        'sections': !exists(json, 'sections') ? undefined : json['sections'],
    };
}

export function V1CacheToJSON(value?: V1Cache | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'disable': value.disable,
        'ttl': value.ttl,
        'io': value.io,
        'sections': value.sections,
    };
}

