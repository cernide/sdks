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
 * @interface V1UserEmail
 */
export interface V1UserEmail {
    /**
     *
     * @type {string}
     * @memberof V1UserEmail
     */
    email?: string;
}

/**
 * Check if a given object implements the V1UserEmail interface.
 */
export function instanceOfV1UserEmail(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1UserEmailFromJSON(json: any): V1UserEmail {
    return V1UserEmailFromJSONTyped(json, false);
}

export function V1UserEmailFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1UserEmail {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'email': !exists(json, 'email') ? undefined : json['email'],
    };
}

export function V1UserEmailToJSON(value?: V1UserEmail | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'email': value.email,
    };
}

