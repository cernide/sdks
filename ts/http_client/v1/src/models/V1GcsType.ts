/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.2
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
 * @interface V1GcsType
 */
export interface V1GcsType {
    /**
     *
     * @type {string}
     * @memberof V1GcsType
     */
    bucket?: string;
    /**
     *
     * @type {string}
     * @memberof V1GcsType
     */
    blob?: string;
}

/**
 * Check if a given object implements the V1GcsType interface.
 */
export function instanceOfV1GcsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1GcsTypeFromJSON(json: any): V1GcsType {
    return V1GcsTypeFromJSONTyped(json, false);
}

export function V1GcsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1GcsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'bucket': !exists(json, 'bucket') ? undefined : json['bucket'],
        'blob': !exists(json, 'blob') ? undefined : json['blob'],
    };
}

export function V1GcsTypeToJSON(value?: V1GcsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'bucket': value.bucket,
        'blob': value.blob,
    };
}

