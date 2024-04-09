/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.6-rc0
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
 * @interface V1Tag
 */
export interface V1Tag {
    /**
     * 
     * @type {string}
     * @memberof V1Tag
     */
    uuid?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Tag
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Tag
     */
    color?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Tag
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Tag
     */
    icon?: string;
    /**
     * 
     * @type {object}
     * @memberof V1Tag
     */
    stats?: object;
}

/**
 * Check if a given object implements the V1Tag interface.
 */
export function instanceOfV1Tag(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1TagFromJSON(json: any): V1Tag {
    return V1TagFromJSONTyped(json, false);
}

export function V1TagFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Tag {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'uuid': !exists(json, 'uuid') ? undefined : json['uuid'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'color': !exists(json, 'color') ? undefined : json['color'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'icon': !exists(json, 'icon') ? undefined : json['icon'],
        'stats': !exists(json, 'stats') ? undefined : json['stats'],
    };
}

export function V1TagToJSON(value?: V1Tag | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'uuid': value.uuid,
        'name': value.name,
        'color': value.color,
        'description': value.description,
        'icon': value.icon,
        'stats': value.stats,
    };
}

