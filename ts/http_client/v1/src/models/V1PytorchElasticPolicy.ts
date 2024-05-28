/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.2.0
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
 * @interface V1PytorchElasticPolicy
 */
export interface V1PytorchElasticPolicy {
    /**
     * minReplicas is the lower limit for the number of replicas to which the training job can scale down.
     * @type {number}
     * @memberof V1PytorchElasticPolicy
     */
    minReplicas?: number;
    /**
     * upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas, defaults to null.
     * @type {number}
     * @memberof V1PytorchElasticPolicy
     */
    maxReplicas?: number;
    /**
     *
     * @type {string}
     * @memberof V1PytorchElasticPolicy
     */
    rdvzBackend?: string;
    /**
     *
     * @type {number}
     * @memberof V1PytorchElasticPolicy
     */
    rdvzPort?: number;
    /**
     *
     * @type {string}
     * @memberof V1PytorchElasticPolicy
     */
    rdvzHost?: string;
    /**
     *
     * @type {string}
     * @memberof V1PytorchElasticPolicy
     */
    rdvzId?: string;
    /**
     * RDZVConf contains additional rendezvous configuration (<key1>=<value1>,<key2>=<value2>,...).
     * @type {Array<object>}
     * @memberof V1PytorchElasticPolicy
     */
    rdvzConf?: Array<object>;
    /**
     * Start a local standalone rendezvous backend that is represented by a C10d TCP store
     * on port 29400. Useful when launching single-node, multi-worker job. If specified
     * --rdzv_backend, --rdzv_endpoint, --rdzv_id are auto-assigned; any explicitly set values
     * are ignored.
     * @type {boolean}
     * @memberof V1PytorchElasticPolicy
     */
    standalone?: boolean;
    /**
     * Number of workers per node; supported values: [auto, cpu, gpu, int].
     * Deprecated: This API is deprecated in v1.7+
     * Use .spec.nprocPerNode instead.
     * @type {number}
     * @memberof V1PytorchElasticPolicy
     */
    nProcPerNode?: number;
    /**
     *
     * @type {number}
     * @memberof V1PytorchElasticPolicy
     */
    maxRestarts?: number;
    /**
     * Metrics contains the specifications which are used to calculate the
     * desired replica count (the maximum replica count across all metrics will
     * be used).  The desired replica count is calculated with multiplying the
     * ratio between the target value and the current value by the current
     * number of pods. Ergo, metrics used must decrease as the pod count is
     * increased, and vice-versa.  See the individual metric source types for
     * more information about how each type of metric must respond.
     * If not set, the HPA will not be created.
     * @type {Array<object>}
     * @memberof V1PytorchElasticPolicy
     */
    Metrics?: Array<object>;
}

/**
 * Check if a given object implements the V1PytorchElasticPolicy interface.
 */
export function instanceOfV1PytorchElasticPolicy(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1PytorchElasticPolicyFromJSON(json: any): V1PytorchElasticPolicy {
    return V1PytorchElasticPolicyFromJSONTyped(json, false);
}

export function V1PytorchElasticPolicyFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1PytorchElasticPolicy {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'minReplicas': !exists(json, 'minReplicas') ? undefined : json['minReplicas'],
        'maxReplicas': !exists(json, 'maxReplicas') ? undefined : json['maxReplicas'],
        'rdvzBackend': !exists(json, 'rdvzBackend') ? undefined : json['rdvzBackend'],
        'rdvzPort': !exists(json, 'rdvzPort') ? undefined : json['rdvzPort'],
        'rdvzHost': !exists(json, 'rdvzHost') ? undefined : json['rdvzHost'],
        'rdvzId': !exists(json, 'rdvzId') ? undefined : json['rdvzId'],
        'rdvzConf': !exists(json, 'rdvzConf') ? undefined : json['rdvzConf'],
        'standalone': !exists(json, 'standalone') ? undefined : json['standalone'],
        'nProcPerNode': !exists(json, 'nProcPerNode') ? undefined : json['nProcPerNode'],
        'maxRestarts': !exists(json, 'maxRestarts') ? undefined : json['maxRestarts'],
        'Metrics': !exists(json, 'Metrics') ? undefined : json['Metrics'],
    };
}

export function V1PytorchElasticPolicyToJSON(value?: V1PytorchElasticPolicy | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'minReplicas': value.minReplicas,
        'maxReplicas': value.maxReplicas,
        'rdvzBackend': value.rdvzBackend,
        'rdvzPort': value.rdvzPort,
        'rdvzHost': value.rdvzHost,
        'rdvzId': value.rdvzId,
        'rdvzConf': value.rdvzConf,
        'standalone': value.standalone,
        'nProcPerNode': value.nProcPerNode,
        'maxRestarts': value.maxRestarts,
        'Metrics': value.Metrics,
    };
}

