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
 * @interface V1SettingsCatalog
 */
export interface V1SettingsCatalog {
    /**
     *
     * @type {string}
     * @memberof V1SettingsCatalog
     */
    uuid?: string;
    /**
     *
     * @type {string}
     * @memberof V1SettingsCatalog
     */
    name?: string;
}

/**
 * Check if a given object implements the V1SettingsCatalog interface.
 */
export function instanceOfV1SettingsCatalog(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1SettingsCatalogFromJSON(json: any): V1SettingsCatalog {
    return V1SettingsCatalogFromJSONTyped(json, false);
}

export function V1SettingsCatalogFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1SettingsCatalog {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'uuid': !exists(json, 'uuid') ? undefined : json['uuid'],
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}

export function V1SettingsCatalogToJSON(value?: V1SettingsCatalog | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'uuid': value.uuid,
        'name': value.name,
    };
}

