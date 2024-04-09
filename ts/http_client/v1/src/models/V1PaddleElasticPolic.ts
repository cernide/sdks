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
 * @interface V1PaddleElasticPolic
 */
export interface V1PaddleElasticPolic {
    /**
     * 
     * @type {number}
     * @memberof V1PaddleElasticPolic
     */
    minReplicas?: number;
    /**
     * 
     * @type {number}
     * @memberof V1PaddleElasticPolic
     */
    maxReplicas?: number;
    /**
     * 
     * @type {number}
     * @memberof V1PaddleElasticPolic
     */
    maxRestarts?: number;
    /**
     * 
     * @type {Array<object>}
     * @memberof V1PaddleElasticPolic
     */
    metrics?: Array<object>;
}

/**
 * Check if a given object implements the V1PaddleElasticPolic interface.
 */
export function instanceOfV1PaddleElasticPolic(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1PaddleElasticPolicFromJSON(json: any): V1PaddleElasticPolic {
    return V1PaddleElasticPolicFromJSONTyped(json, false);
}

export function V1PaddleElasticPolicFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1PaddleElasticPolic {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'minReplicas': !exists(json, 'minReplicas') ? undefined : json['minReplicas'],
        'maxReplicas': !exists(json, 'maxReplicas') ? undefined : json['maxReplicas'],
        'maxRestarts': !exists(json, 'maxRestarts') ? undefined : json['maxRestarts'],
        'metrics': !exists(json, 'metrics') ? undefined : json['metrics'],
    };
}

export function V1PaddleElasticPolicToJSON(value?: V1PaddleElasticPolic | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'minReplicas': value.minReplicas,
        'maxReplicas': value.maxReplicas,
        'maxRestarts': value.maxRestarts,
        'metrics': value.metrics,
    };
}

