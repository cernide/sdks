/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.2-rc1
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { V1ConnectionKind } from './V1ConnectionKind';
import {
    V1ConnectionKindFromJSON,
    V1ConnectionKindFromJSONTyped,
    V1ConnectionKindToJSON,
} from './V1ConnectionKind';
import type { V1ConnectionResource } from './V1ConnectionResource';
import {
    V1ConnectionResourceFromJSON,
    V1ConnectionResourceFromJSONTyped,
    V1ConnectionResourceToJSON,
} from './V1ConnectionResource';

/**
 *
 * @export
 * @interface V1ConnectionType
 */
export interface V1ConnectionType {
    /**
     *
     * @type {string}
     * @memberof V1ConnectionType
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof V1ConnectionType
     */
    description?: string;
    /**
     *
     * @type {string}
     * @memberof V1ConnectionType
     */
    tags?: string;
    /**
     *
     * @type {V1ConnectionKind}
     * @memberof V1ConnectionType
     */
    kind?: V1ConnectionKind;
    /**
     *
     * @type {object}
     * @memberof V1ConnectionType
     */
    schema?: object;
    /**
     *
     * @type {V1ConnectionResource}
     * @memberof V1ConnectionType
     */
    secret?: V1ConnectionResource;
    /**
     *
     * @type {V1ConnectionResource}
     * @memberof V1ConnectionType
     */
    configMap?: V1ConnectionResource;
    /**
     *
     * @type {Array<object>}
     * @memberof V1ConnectionType
     */
    env?: Array<object>;
    /**
     *
     * @type {{ [key: string]: string; }}
     * @memberof V1ConnectionType
     */
    annotations?: { [key: string]: string; };
}

/**
 * Check if a given object implements the V1ConnectionType interface.
 */
export function instanceOfV1ConnectionType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1ConnectionTypeFromJSON(json: any): V1ConnectionType {
    return V1ConnectionTypeFromJSONTyped(json, false);
}

export function V1ConnectionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ConnectionType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'name': !exists(json, 'name') ? undefined : json['name'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'kind': !exists(json, 'kind') ? undefined : V1ConnectionKindFromJSON(json['kind']),
        'schema': !exists(json, 'schema') ? undefined : json['schema'],
        'secret': !exists(json, 'secret') ? undefined : V1ConnectionResourceFromJSON(json['secret']),
        'configMap': !exists(json, 'configMap') ? undefined : V1ConnectionResourceFromJSON(json['configMap']),
        'env': !exists(json, 'env') ? undefined : json['env'],
        'annotations': !exists(json, 'annotations') ? undefined : json['annotations'],
    };
}

export function V1ConnectionTypeToJSON(value?: V1ConnectionType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'name': value.name,
        'description': value.description,
        'tags': value.tags,
        'kind': V1ConnectionKindToJSON(value.kind),
        'schema': value.schema,
        'secret': V1ConnectionResourceToJSON(value.secret),
        'configMap': V1ConnectionResourceToJSON(value.configMap),
        'env': value.env,
        'annotations': value.annotations,
    };
}

