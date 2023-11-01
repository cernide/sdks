/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0
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
 * @interface V1Queue
 */
export interface V1Queue {
    /**
     *
     * @type {string}
     * @memberof V1Queue
     */
    uuid?: string;
    /**
     *
     * @type {string}
     * @memberof V1Queue
     */
    agent?: string;
    /**
     *
     * @type {string}
     * @memberof V1Queue
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof V1Queue
     */
    description?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof V1Queue
     */
    tags?: Array<string>;
    /**
     *
     * @type {number}
     * @memberof V1Queue
     */
    priority?: number;
    /**
     *
     * @type {number}
     * @memberof V1Queue
     */
    concurrency?: number;
    /**
     *
     * @type {string}
     * @memberof V1Queue
     */
    resource?: string;
    /**
     *
     * @type {string}
     * @memberof V1Queue
     */
    quota?: string;
    /**
     *
     * @type {object}
     * @memberof V1Queue
     */
    stats?: object;
    /**
     *
     * @type {Date}
     * @memberof V1Queue
     */
    created_at?: Date;
    /**
     *
     * @type {Date}
     * @memberof V1Queue
     */
    updated_at?: Date;
}

/**
 * Check if a given object implements the V1Queue interface.
 */
export function instanceOfV1Queue(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1QueueFromJSON(json: any): V1Queue {
    return V1QueueFromJSONTyped(json, false);
}

export function V1QueueFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Queue {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'uuid': !exists(json, 'uuid') ? undefined : json['uuid'],
        'agent': !exists(json, 'agent') ? undefined : json['agent'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'priority': !exists(json, 'priority') ? undefined : json['priority'],
        'concurrency': !exists(json, 'concurrency') ? undefined : json['concurrency'],
        'resource': !exists(json, 'resource') ? undefined : json['resource'],
        'quota': !exists(json, 'quota') ? undefined : json['quota'],
        'stats': !exists(json, 'stats') ? undefined : json['stats'],
        'created_at': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'updated_at': !exists(json, 'updated_at') ? undefined : (new Date(json['updated_at'])),
    };
}

export function V1QueueToJSON(value?: V1Queue | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'uuid': value.uuid,
        'agent': value.agent,
        'name': value.name,
        'description': value.description,
        'tags': value.tags,
        'priority': value.priority,
        'concurrency': value.concurrency,
        'resource': value.resource,
        'quota': value.quota,
        'stats': value.stats,
        'created_at': value.created_at === undefined ? undefined : (value.created_at.toISOString()),
        'updated_at': value.updated_at === undefined ? undefined : (value.updated_at.toISOString()),
    };
}

