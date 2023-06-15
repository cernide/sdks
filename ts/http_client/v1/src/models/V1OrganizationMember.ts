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
 * @interface V1OrganizationMember
 */
export interface V1OrganizationMember {
    /**
     *
     * @type {string}
     * @memberof V1OrganizationMember
     */
    user?: string;
    /**
     *
     * @type {string}
     * @memberof V1OrganizationMember
     */
    user_email?: string;
    /**
     *
     * @type {string}
     * @memberof V1OrganizationMember
     */
    role?: string;
    /**
     *
     * @type {string}
     * @memberof V1OrganizationMember
     */
    kind?: string;
    /**
     *
     * @type {Date}
     * @memberof V1OrganizationMember
     */
    created_at?: Date;
    /**
     *
     * @type {Date}
     * @memberof V1OrganizationMember
     */
    updated_at?: Date;
}

/**
 * Check if a given object implements the V1OrganizationMember interface.
 */
export function instanceOfV1OrganizationMember(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1OrganizationMemberFromJSON(json: any): V1OrganizationMember {
    return V1OrganizationMemberFromJSONTyped(json, false);
}

export function V1OrganizationMemberFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1OrganizationMember {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'user': !exists(json, 'user') ? undefined : json['user'],
        'user_email': !exists(json, 'user_email') ? undefined : json['user_email'],
        'role': !exists(json, 'role') ? undefined : json['role'],
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'created_at': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'updated_at': !exists(json, 'updated_at') ? undefined : (new Date(json['updated_at'])),
    };
}

export function V1OrganizationMemberToJSON(value?: V1OrganizationMember | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'user': value.user,
        'user_email': value.user_email,
        'role': value.role,
        'kind': value.kind,
        'created_at': value.created_at === undefined ? undefined : (value.created_at.toISOString()),
        'updated_at': value.updated_at === undefined ? undefined : (value.updated_at.toISOString()),
    };
}

