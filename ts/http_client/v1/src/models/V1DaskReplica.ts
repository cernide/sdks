/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.4.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { V1Environment } from './V1Environment';
import {
    V1EnvironmentFromJSON,
    V1EnvironmentFromJSONTyped,
    V1EnvironmentToJSON,
} from './V1Environment';
import type { V1Init } from './V1Init';
import {
    V1InitFromJSON,
    V1InitFromJSONTyped,
    V1InitToJSON,
} from './V1Init';

/**
 *
 * @export
 * @interface V1DaskReplica
 */
export interface V1DaskReplica {
    /**
     *
     * @type {number}
     * @memberof V1DaskReplica
     */
    replicas?: number;
    /**
     *
     * @type {V1Environment}
     * @memberof V1DaskReplica
     */
    environment?: V1Environment;
    /**
     *
     * @type {Array<string>}
     * @memberof V1DaskReplica
     */
    connections?: Array<string>;
    /**
     *
     * @type {Array<object>}
     * @memberof V1DaskReplica
     */
    volumes?: Array<object>;
    /**
     *
     * @type {Array<V1Init>}
     * @memberof V1DaskReplica
     */
    init?: Array<V1Init>;
    /**
     *
     * @type {Array<object>}
     * @memberof V1DaskReplica
     */
    sidecars?: Array<object>;
    /**
     *
     * @type {object}
     * @memberof V1DaskReplica
     */
    container?: object;
}

/**
 * Check if a given object implements the V1DaskReplica interface.
 */
export function instanceOfV1DaskReplica(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1DaskReplicaFromJSON(json: any): V1DaskReplica {
    return V1DaskReplicaFromJSONTyped(json, false);
}

export function V1DaskReplicaFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1DaskReplica {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'replicas': !exists(json, 'replicas') ? undefined : json['replicas'],
        'environment': !exists(json, 'environment') ? undefined : V1EnvironmentFromJSON(json['environment']),
        'connections': !exists(json, 'connections') ? undefined : json['connections'],
        'volumes': !exists(json, 'volumes') ? undefined : json['volumes'],
        'init': !exists(json, 'init') ? undefined : ((json['init'] as Array<any>).map(V1InitFromJSON)),
        'sidecars': !exists(json, 'sidecars') ? undefined : json['sidecars'],
        'container': !exists(json, 'container') ? undefined : json['container'],
    };
}

export function V1DaskReplicaToJSON(value?: V1DaskReplica | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'replicas': value.replicas,
        'environment': V1EnvironmentToJSON(value.environment),
        'connections': value.connections,
        'volumes': value.volumes,
        'init': value.init === undefined ? undefined : ((value.init as Array<any>).map(V1InitToJSON)),
        'sidecars': value.sidecars,
        'container': value.container,
    };
}

