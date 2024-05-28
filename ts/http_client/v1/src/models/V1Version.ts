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
 * @interface V1Version
 */
export interface V1Version {
    /**
     *
     * @type {string}
     * @memberof V1Version
     */
    min?: string;
    /**
     *
     * @type {string}
     * @memberof V1Version
     */
    latest?: string;
}

/**
 * Check if a given object implements the V1Version interface.
 */
export function instanceOfV1Version(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1VersionFromJSON(json: any): V1Version {
    return V1VersionFromJSONTyped(json, false);
}

export function V1VersionFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Version {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'min': !exists(json, 'min') ? undefined : json['min'],
        'latest': !exists(json, 'latest') ? undefined : json['latest'],
    };
}

export function V1VersionToJSON(value?: V1Version | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'min': value.min,
        'latest': value.latest,
    };
}

