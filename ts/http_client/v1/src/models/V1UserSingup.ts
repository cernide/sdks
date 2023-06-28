/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc23
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
 * @interface V1UserSingup
 */
export interface V1UserSingup {
    /**
     *
     * @type {string}
     * @memberof V1UserSingup
     */
    username?: string;
    /**
     *
     * @type {string}
     * @memberof V1UserSingup
     */
    email?: string;
    /**
     *
     * @type {string}
     * @memberof V1UserSingup
     */
    organization?: string;
    /**
     *
     * @type {string}
     * @memberof V1UserSingup
     */
    password?: string;
    /**
     *
     * @type {string}
     * @memberof V1UserSingup
     */
    invitation_key?: string;
}

/**
 * Check if a given object implements the V1UserSingup interface.
 */
export function instanceOfV1UserSingup(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1UserSingupFromJSON(json: any): V1UserSingup {
    return V1UserSingupFromJSONTyped(json, false);
}

export function V1UserSingupFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1UserSingup {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'username': !exists(json, 'username') ? undefined : json['username'],
        'email': !exists(json, 'email') ? undefined : json['email'],
        'organization': !exists(json, 'organization') ? undefined : json['organization'],
        'password': !exists(json, 'password') ? undefined : json['password'],
        'invitation_key': !exists(json, 'invitation_key') ? undefined : json['invitation_key'],
    };
}

export function V1UserSingupToJSON(value?: V1UserSingup | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'username': value.username,
        'email': value.email,
        'organization': value.organization,
        'password': value.password,
        'invitation_key': value.invitation_key,
    };
}

