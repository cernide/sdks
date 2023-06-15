/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc19
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
 * @interface V1ProjectUserAccess
 */
export interface V1ProjectUserAccess {
    /**
     *
     * @type {string}
     * @memberof V1ProjectUserAccess
     */
    user?: string;
    /**
     *
     * @type {string}
     * @memberof V1ProjectUserAccess
     */
    queue?: string;
    /**
     *
     * @type {string}
     * @memberof V1ProjectUserAccess
     */
    preset?: string;
}

/**
 * Check if a given object implements the V1ProjectUserAccess interface.
 */
export function instanceOfV1ProjectUserAccess(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1ProjectUserAccessFromJSON(json: any): V1ProjectUserAccess {
    return V1ProjectUserAccessFromJSONTyped(json, false);
}

export function V1ProjectUserAccessFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ProjectUserAccess {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'user': !exists(json, 'user') ? undefined : json['user'],
        'queue': !exists(json, 'queue') ? undefined : json['queue'],
        'preset': !exists(json, 'preset') ? undefined : json['preset'],
    };
}

export function V1ProjectUserAccessToJSON(value?: V1ProjectUserAccess | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'user': value.user,
        'queue': value.queue,
        'preset': value.preset,
    };
}

