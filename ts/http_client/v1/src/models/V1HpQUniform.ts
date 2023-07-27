/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc28
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
 * @interface V1HpQUniform
 */
export interface V1HpQUniform {
    /**
     *
     * @type {string}
     * @memberof V1HpQUniform
     */
    kind?: string;
    /**
     *
     * @type {object}
     * @memberof V1HpQUniform
     */
    value?: object;
}

/**
 * Check if a given object implements the V1HpQUniform interface.
 */
export function instanceOfV1HpQUniform(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1HpQUniformFromJSON(json: any): V1HpQUniform {
    return V1HpQUniformFromJSONTyped(json, false);
}

export function V1HpQUniformFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1HpQUniform {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'value': !exists(json, 'value') ? undefined : json['value'],
    };
}

export function V1HpQUniformToJSON(value?: V1HpQUniform | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'kind': value.kind,
        'value': value.value,
    };
}

