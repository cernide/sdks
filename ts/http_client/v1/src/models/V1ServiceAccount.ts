/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc33
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
 * @interface V1ServiceAccount
 */
export interface V1ServiceAccount {
    /**
     *
     * @type {string}
     * @memberof V1ServiceAccount
     */
    uuid?: string;
    /**
     *
     * @type {string}
     * @memberof V1ServiceAccount
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof V1ServiceAccount
     */
    description?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof V1ServiceAccount
     */
    tags?: Array<string>;
    /**
     *
     * @type {number}
     * @memberof V1ServiceAccount
     */
    live_state?: number;
    /**
     *
     * @type {Date}
     * @memberof V1ServiceAccount
     */
    created_at?: Date;
    /**
     *
     * @type {Date}
     * @memberof V1ServiceAccount
     */
    updated_at?: Date;
    /**
     *
     * @type {Array<string>}
     * @memberof V1ServiceAccount
     */
    scopes?: Array<string>;
    /**
     *
     * @type {Array<string>}
     * @memberof V1ServiceAccount
     */
    services?: Array<string>;
}

/**
 * Check if a given object implements the V1ServiceAccount interface.
 */
export function instanceOfV1ServiceAccount(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1ServiceAccountFromJSON(json: any): V1ServiceAccount {
    return V1ServiceAccountFromJSONTyped(json, false);
}

export function V1ServiceAccountFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ServiceAccount {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'uuid': !exists(json, 'uuid') ? undefined : json['uuid'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'live_state': !exists(json, 'live_state') ? undefined : json['live_state'],
        'created_at': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'updated_at': !exists(json, 'updated_at') ? undefined : (new Date(json['updated_at'])),
        'scopes': !exists(json, 'scopes') ? undefined : json['scopes'],
        'services': !exists(json, 'services') ? undefined : json['services'],
    };
}

export function V1ServiceAccountToJSON(value?: V1ServiceAccount | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'uuid': value.uuid,
        'name': value.name,
        'description': value.description,
        'tags': value.tags,
        'live_state': value.live_state,
        'created_at': value.created_at === undefined ? undefined : (value.created_at.toISOString()),
        'updated_at': value.updated_at === undefined ? undefined : (value.updated_at.toISOString()),
        'scopes': value.scopes,
        'services': value.services,
    };
}

