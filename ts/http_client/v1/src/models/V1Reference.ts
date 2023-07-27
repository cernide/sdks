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
import type { V1DagRef } from './V1DagRef';
import {
    V1DagRefFromJSON,
    V1DagRefFromJSONTyped,
    V1DagRefToJSON,
} from './V1DagRef';
import type { V1HubRef } from './V1HubRef';
import {
    V1HubRefFromJSON,
    V1HubRefFromJSONTyped,
    V1HubRefToJSON,
} from './V1HubRef';
import type { V1PathRef } from './V1PathRef';
import {
    V1PathRefFromJSON,
    V1PathRefFromJSONTyped,
    V1PathRefToJSON,
} from './V1PathRef';
import type { V1UrlRef } from './V1UrlRef';
import {
    V1UrlRefFromJSON,
    V1UrlRefFromJSONTyped,
    V1UrlRefToJSON,
} from './V1UrlRef';

/**
 *
 * @export
 * @interface V1Reference
 */
export interface V1Reference {
    /**
     *
     * @type {V1HubRef}
     * @memberof V1Reference
     */
    hubRef?: V1HubRef;
    /**
     *
     * @type {V1DagRef}
     * @memberof V1Reference
     */
    dagRef?: V1DagRef;
    /**
     *
     * @type {V1UrlRef}
     * @memberof V1Reference
     */
    urlRef?: V1UrlRef;
    /**
     *
     * @type {V1PathRef}
     * @memberof V1Reference
     */
    pathRef?: V1PathRef;
}

/**
 * Check if a given object implements the V1Reference interface.
 */
export function instanceOfV1Reference(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1ReferenceFromJSON(json: any): V1Reference {
    return V1ReferenceFromJSONTyped(json, false);
}

export function V1ReferenceFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Reference {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'hubRef': !exists(json, 'hubRef') ? undefined : V1HubRefFromJSON(json['hubRef']),
        'dagRef': !exists(json, 'dagRef') ? undefined : V1DagRefFromJSON(json['dagRef']),
        'urlRef': !exists(json, 'urlRef') ? undefined : V1UrlRefFromJSON(json['urlRef']),
        'pathRef': !exists(json, 'pathRef') ? undefined : V1PathRefFromJSON(json['pathRef']),
    };
}

export function V1ReferenceToJSON(value?: V1Reference | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'hubRef': V1HubRefToJSON(value.hubRef),
        'dagRef': V1DagRefToJSON(value.dagRef),
        'urlRef': V1UrlRefToJSON(value.urlRef),
        'pathRef': V1PathRefToJSON(value.pathRef),
    };
}

