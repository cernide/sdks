/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.0.6-rc0
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
 * @interface V1Credentials
 */
export interface V1Credentials {
    /**
     * 
     * @type {string}
     * @memberof V1Credentials
     */
    username?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Credentials
     */
    password?: string;
}

/**
 * Check if a given object implements the V1Credentials interface.
 */
export function instanceOfV1Credentials(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1CredentialsFromJSON(json: any): V1Credentials {
    return V1CredentialsFromJSONTyped(json, false);
}

export function V1CredentialsFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Credentials {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'username': !exists(json, 'username') ? undefined : json['username'],
        'password': !exists(json, 'password') ? undefined : json['password'],
    };
}

export function V1CredentialsToJSON(value?: V1Credentials | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'username': value.username,
        'password': value.password,
    };
}

