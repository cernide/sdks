/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc21
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
 * @interface V1EventDataframe
 */
export interface V1EventDataframe {
    /**
     *
     * @type {string}
     * @memberof V1EventDataframe
     */
    path?: string;
    /**
     *
     * @type {string}
     * @memberof V1EventDataframe
     */
    content_type?: string;
}

/**
 * Check if a given object implements the V1EventDataframe interface.
 */
export function instanceOfV1EventDataframe(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1EventDataframeFromJSON(json: any): V1EventDataframe {
    return V1EventDataframeFromJSONTyped(json, false);
}

export function V1EventDataframeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1EventDataframe {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'path': !exists(json, 'path') ? undefined : json['path'],
        'content_type': !exists(json, 'content_type') ? undefined : json['content_type'],
    };
}

export function V1EventDataframeToJSON(value?: V1EventDataframe | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'path': value.path,
        'content_type': value.content_type,
    };
}

