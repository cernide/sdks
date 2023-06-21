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
 * @interface V1AuthType
 */
export interface V1AuthType {
    /**
     *
     * @type {string}
     * @memberof V1AuthType
     */
    user?: string;
    /**
     *
     * @type {string}
     * @memberof V1AuthType
     */
    password?: string;
}

/**
 * Check if a given object implements the V1AuthType interface.
 */
export function instanceOfV1AuthType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1AuthTypeFromJSON(json: any): V1AuthType {
    return V1AuthTypeFromJSONTyped(json, false);
}

export function V1AuthTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1AuthType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'user': !exists(json, 'user') ? undefined : json['user'],
        'password': !exists(json, 'password') ? undefined : json['password'],
    };
}

export function V1AuthTypeToJSON(value?: V1AuthType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'user': value.user,
        'password': value.password,
    };
}

