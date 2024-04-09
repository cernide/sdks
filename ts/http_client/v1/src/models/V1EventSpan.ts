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
import type { V1EventSpanKind } from './V1EventSpanKind';
import {
    V1EventSpanKindFromJSON,
    V1EventSpanKindFromJSONTyped,
    V1EventSpanKindToJSON,
} from './V1EventSpanKind';
import type { V1StatusCondition } from './V1StatusCondition';
import {
    V1StatusConditionFromJSON,
    V1StatusConditionFromJSONTyped,
    V1StatusConditionToJSON,
} from './V1StatusCondition';
import type { V1Statuses } from './V1Statuses';
import {
    V1StatusesFromJSON,
    V1StatusesFromJSONTyped,
    V1StatusesToJSON,
} from './V1Statuses';

/**
 * 
 * @export
 * @interface V1EventSpan
 */
export interface V1EventSpan {
    /**
     * 
     * @type {string}
     * @memberof V1EventSpan
     */
    uuid?: string;
    /**
     * 
     * @type {string}
     * @memberof V1EventSpan
     */
    name?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1EventSpan
     */
    tags?: Array<string>;
    /**
     * 
     * @type {Date}
     * @memberof V1EventSpan
     */
    started_at?: Date;
    /**
     * 
     * @type {Date}
     * @memberof V1EventSpan
     */
    finished_at?: Date;
    /**
     * 
     * @type {V1Statuses}
     * @memberof V1EventSpan
     */
    status?: V1Statuses;
    /**
     * 
     * @type {Array<V1StatusCondition>}
     * @memberof V1EventSpan
     */
    status_conditions?: Array<V1StatusCondition>;
    /**
     * 
     * @type {object}
     * @memberof V1EventSpan
     */
    metadata?: object;
    /**
     * 
     * @type {object}
     * @memberof V1EventSpan
     */
    inputs?: object;
    /**
     * 
     * @type {object}
     * @memberof V1EventSpan
     */
    outputs?: object;
    /**
     * 
     * @type {Array<V1EventSpan>}
     * @memberof V1EventSpan
     */
    children?: Array<V1EventSpan>;
    /**
     * 
     * @type {V1EventSpanKind}
     * @memberof V1EventSpan
     */
    kind?: V1EventSpanKind;
}

/**
 * Check if a given object implements the V1EventSpan interface.
 */
export function instanceOfV1EventSpan(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1EventSpanFromJSON(json: any): V1EventSpan {
    return V1EventSpanFromJSONTyped(json, false);
}

export function V1EventSpanFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1EventSpan {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'uuid': !exists(json, 'uuid') ? undefined : json['uuid'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'started_at': !exists(json, 'started_at') ? undefined : (new Date(json['started_at'])),
        'finished_at': !exists(json, 'finished_at') ? undefined : (new Date(json['finished_at'])),
        'status': !exists(json, 'status') ? undefined : V1StatusesFromJSON(json['status']),
        'status_conditions': !exists(json, 'status_conditions') ? undefined : ((json['status_conditions'] as Array<any>).map(V1StatusConditionFromJSON)),
        'metadata': !exists(json, 'metadata') ? undefined : json['metadata'],
        'inputs': !exists(json, 'inputs') ? undefined : json['inputs'],
        'outputs': !exists(json, 'outputs') ? undefined : json['outputs'],
        'children': !exists(json, 'children') ? undefined : ((json['children'] as Array<any>).map(V1EventSpanFromJSON)),
        'kind': !exists(json, 'kind') ? undefined : V1EventSpanKindFromJSON(json['kind']),
    };
}

export function V1EventSpanToJSON(value?: V1EventSpan | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'uuid': value.uuid,
        'name': value.name,
        'tags': value.tags,
        'started_at': value.started_at === undefined ? undefined : (value.started_at.toISOString()),
        'finished_at': value.finished_at === undefined ? undefined : (value.finished_at.toISOString()),
        'status': V1StatusesToJSON(value.status),
        'status_conditions': value.status_conditions === undefined ? undefined : ((value.status_conditions as Array<any>).map(V1StatusConditionToJSON)),
        'metadata': value.metadata,
        'inputs': value.inputs,
        'outputs': value.outputs,
        'children': value.children === undefined ? undefined : ((value.children as Array<any>).map(V1EventSpanToJSON)),
        'kind': V1EventSpanKindToJSON(value.kind),
    };
}

